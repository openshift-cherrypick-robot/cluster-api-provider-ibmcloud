# Copyright 2021 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

ROOT_DIR_RELATIVE := .

include $(ROOT_DIR_RELATIVE)/common.mk

# Image URL to use all building/pushing image targets
IMG ?= controller:latest
# Produce CRDs that work back to Kubernetes 1.11 (no version conversion)
CRD_OPTIONS ?= "crd:crdVersions=v1"

# Directories.
REPO_ROOT := $(shell git rev-parse --show-toplevel)
ARTIFACTS ?= $(REPO_ROOT)/_artifacts
TOOLS_DIR := hack/tools
TOOLS_BIN_DIR := $(TOOLS_DIR)/bin
GO_INSTALL = ./scripts/go_install.sh
E2E_CONF_FILE_ENVSUBST := $(REPO_ROOT)/test/e2e/config/ibmcloud-e2e-envsubst.yaml

GOLANGCI_LINT := $(TOOLS_BIN_DIR)/golangci-lint
KUSTOMIZE := $(TOOLS_BIN_DIR)/kustomize
GOJQ := $(TOOLS_BIN_DIR)/gojq
CONVERSION_GEN := $(TOOLS_BIN_DIR)/conversion-gen
GINKGO := $(TOOLS_BIN_DIR)/ginkgo
ENVSUBST := $(TOOLS_BIN_DIR)/envsubst
MOCKGEN := $(TOOLS_BIN_DIR)/mockgen
CONTROLLER_GEN := $(TOOLS_BIN_DIR)/controller-gen

STAGING_REGISTRY ?= gcr.io/k8s-staging-capi-ibmcloud
STAGING_BUCKET ?= artifacts.k8s-staging-capi-ibmcloud.appspot.com
BUCKET ?= $(STAGING_BUCKET)
PROD_REGISTRY := k8s.gcr.io/capi-ibmcloud
REGISTRY ?= $(STAGING_REGISTRY)
RELEASE_TAG ?= $(shell git describe --abbrev=0 2>/dev/null)
PULL_BASE_REF ?= $(RELEASE_TAG) # PULL_BASE_REF will be provided by Prow
RELEASE_ALIAS_TAG ?= $(PULL_BASE_REF)
RELEASE_DIR := out

TAG ?= dev
ARCH ?= amd64
ALL_ARCH ?= amd64 ppc64le
PULL_POLICY ?= Always

# main controller
CORE_IMAGE_NAME ?= cluster-api-ibmcloud-controller
CORE_CONTROLLER_IMG ?= $(REGISTRY)/$(CORE_IMAGE_NAME)
CORE_CONTROLLER_ORIGINAL_IMG := gcr.io/k8s-staging-capi-ibmcloud/cluster-api-ibmcloud-controller
CORE_CONTROLLER_NAME := controller-manager
CORE_MANIFEST_FILE := infrastructure-components
CORE_CONFIG_DIR := config/default
CORE_NAMESPACE := capi-ibmcloud-system

PATH := $(abspath $(TOOLS_BIN_DIR)):$(PATH)
# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# Set --output-base for conversion-gen if we are not within GOPATH
ifneq ($(abspath $(ROOT_DIR_RELATIVE)),$(shell go env GOPATH)/src/sigs.k8s.io/cluster-api-provider-ibmcloud)
	CONVERSION_GEN_OUTPUT_BASE := --output-base=$(ROOT_DIR_RELATIVE)
else
	export GOPATH := $(shell go env GOPATH)
endif

all: manager

# Build manager binary
manager: generate fmt vet
	go build -o bin/manager main.go

# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate fmt vet manifests
	go run ./main.go

# Install CRDs into a cluster
install: manifests $(KUSTOMIZE)
	$(KUSTOMIZE) build config/crd | kubectl apply -f -

# Uninstall CRDs from a cluster
uninstall: manifests $(KUSTOMIZE)
	$(KUSTOMIZE) build config/crd | kubectl delete -f -

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
deploy: manifests $(KUSTOMIZE)
	cd config/manager && $(KUSTOMIZE) edit set image controller=${IMG}
	$(KUSTOMIZE) build config/default | kubectl apply -f -

# Generate manifests e.g. CRD, RBAC etc.
manifests: $(CONTROLLER_GEN)
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role paths="./..." output:crd:artifacts:config=config/crd/bases

.PHONY: generate-go-conversions-core
generate-go-conversions-core: $(CONVERSION_GEN)
	$(MAKE) clean-generated-conversions SRC_DIRS="./api/v1alpha3"
	$(MAKE) clean-generated-conversions SRC_DIRS="./api/v1alpha4"
	$(CONVERSION_GEN) \
		--input-dirs=./api/v1alpha3 \
		--input-dirs=./api/v1alpha4 \
		--build-tag=ignore_autogenerated_core \
		--output-file-base=zz_generated.conversion $(CONVERSION_GEN_OUTPUT_BASE) \
		--go-header-file=./hack/boilerplate/boilerplate.generatego.txt

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

# Generate code
generate: $(CONTROLLER_GEN)
	$(MAKE) generate-go
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."

.PHONY: generate-go
generate-go: $(MOCKGEN)
	go generate ./...

images: docker-build

set-flavor: 
ifeq ($(E2E_FLAVOR), vpc)
	 $(eval E2E_CONF_FILE=$(REPO_ROOT)/test/e2e/config/ibmcloud-e2e-vpc.yaml)
else
	 $(eval E2E_CONF_FILE=$(REPO_ROOT)/test/e2e/config/ibmcloud-e2e-powervs.yaml)
endif
	@echo "Setting e2e test flavour to ${E2E_CONF_FILE}"

## --------------------------------------
## Linting
## --------------------------------------

.PHONY: lint
lint: $(GOLANGCI_LINT) ## Lint codebase
	$(GOLANGCI_LINT) run -v --fast=false

## --------------------------------------
## Testing
## --------------------------------------

# Run unit tests
test: generate fmt vet manifests
	go test ./... -coverprofile cover.out

# Allow overriding the e2e configurations
GINKGO_FOCUS ?= Workload cluster creation
GINKGO_NODES ?= 3
GINKGO_NOCOLOR ?= false
E2E_FLAVOR ?= powervs
GINKGO_ARGS ?= -v -trace -progress -v -tags=e2e -focus=$(GINKGO_FOCUS) -nodes=$(GINKGO_NODES) --noColor=$(GINKGO_NOCOLOR)
ARTIFACTS ?= $(REPO_ROOT)/_artifacts
SKIP_CLEANUP ?= false
SKIP_CREATE_MGMT_CLUSTER ?= false

# Run the end-to-end tests
.PHONY: test-e2e
test-e2e: $(KUBECTL) $(GINKGO) $(KUSTOMIZE) $(ENVSUBST) set-flavor e2e-image 
	$(ENVSUBST) < $(E2E_CONF_FILE) > $(E2E_CONF_FILE_ENVSUBST) 
	$(GINKGO) $(GINKGO_ARGS) ./test/e2e -- \
		-e2e.artifacts-folder="$(ARTIFACTS)" \
		-e2e.config="$(E2E_CONF_FILE_ENVSUBST)" \
		-e2e.skip-resource-cleanup=$(SKIP_CLEANUP) \
		-e2e.use-existing-cluster=$(SKIP_CREATE_MGMT_CLUSTER) \
		-e2e.flavor="$(E2E_FLAVOR)"

# Basic checks for deploying kind cluster and required providers
.PHONY: test-sanity
test-sanity:
	GINKGO_FOCUS="Run Sanity tests" $(MAKE) test-e2e

## --------------------------------------
## Docker
## --------------------------------------

.PHONY: docker-build
docker-build: docker-pull-prerequisites ## Build the docker image for controller-manager
	docker build --build-arg ARCH=$(ARCH) . -t $(CORE_CONTROLLER_IMG)-$(ARCH):$(TAG)

.PHONY: docker-push
docker-push: ## Push the docker image
	docker push $(CORE_CONTROLLER_IMG)-$(ARCH):$(TAG)

.PHONY: docker-pull-prerequisites
docker-pull-prerequisites:
	docker pull docker.io/docker/dockerfile:1.1-experimental
	docker pull gcr.io/distroless/static:latest

.PHONY: e2e-image
e2e-image: docker-pull-prerequisites
	docker build --tag=$(CORE_CONTROLLER_ORIGINAL_IMG):e2e .
	$(MAKE) set-manifest-image MANIFEST_IMG=$(CORE_CONTROLLER_ORIGINAL_IMG):e2e TARGET_RESOURCE="./config/default/manager_image_patch.yaml"
	$(MAKE) set-manifest-pull-policy PULL_POLICY=Never TARGET_RESOURCE="./config/default/manager_pull_policy.yaml"

.PHONY: set-manifest-image
set-manifest-image:
	$(info Updating kustomize image patch file for default resource)
	sed -i'' -e 's@image: .*@image: '"${MANIFEST_IMG}"'@' ./config/default/manager_image_patch.yaml

.PHONY: set-manifest-pull-policy
set-manifest-pull-policy:
	$(info Updating kustomize pull policy file for default resource)
	sed -i'' -e 's@imagePullPolicy: .*@imagePullPolicy: '"$(PULL_POLICY)"'@' ./config/default/manager_pull_policy.yaml	

## --------------------------------------
## Docker - All ARCH
## --------------------------------------

.PHONY: docker-build-all ## Build all the architecture docker images
docker-build-all: $(addprefix docker-build-,$(ALL_ARCH))

docker-build-%:
	$(MAKE) ARCH=$* docker-build

.PHONY: docker-push-all ## Push all the architecture docker images
docker-push-all: $(addprefix docker-push-,$(ALL_ARCH))
	$(MAKE) docker-push-core-manifest

docker-push-%:
	$(MAKE) ARCH=$* docker-push

.PHONY: docker-push-core-manifest
docker-push-core-manifest: ## Push the fat manifest docker image.
	## Minimum docker version 18.06.0 is required for creating and pushing manifest images.
	$(MAKE) docker-push-manifest CONTROLLER_IMG=$(CORE_CONTROLLER_IMG) MANIFEST_FILE=$(CORE_MANIFEST_FILE)

.PHONY: docker-push-manifest
docker-push-manifest:
	docker manifest create --amend $(CONTROLLER_IMG):$(TAG) $(shell echo $(ALL_ARCH) | sed -e "s~[^ ]*~$(CONTROLLER_IMG)\-&:$(TAG)~g")
	@for arch in $(ALL_ARCH); do docker manifest annotate --arch $${arch} ${CONTROLLER_IMG}:${TAG} ${CONTROLLER_IMG}-$${arch}:${TAG}; done
	docker manifest push --purge ${CONTROLLER_IMG}:${TAG}

## --------------------------------------
## Release
## --------------------------------------

$(RELEASE_DIR):
	mkdir -p $@

$(ARTIFACTS):
	mkdir -p $@

.PHONY: list-staging-releases
list-staging-releases: ## List staging images for image promotion
	@echo $(CORE_IMAGE_NAME):
	$(MAKE) list-image RELEASE_TAG=$(RELEASE_TAG) IMAGE=$(CORE_IMAGE_NAME)

list-image:
	gcloud container images list-tags $(STAGING_REGISTRY)/$(IMAGE) --filter="tags=('$(RELEASE_TAG)')" --format=json

.PHONY: check-release-tag
check-release-tag:
	@if [ -z "${RELEASE_TAG}" ]; then echo "RELEASE_TAG is not set"; exit 1; fi
	@if ! [ -z "$$(git status --porcelain)" ]; then echo "Your local git repository contains uncommitted changes, use git clean before proceeding."; exit 1; fi

.PHONY: check-previous-release-tag
check-previous-release-tag:
	@if [ -z "${PREVIOUS_VERSION}" ]; then echo "PREVIOUS_VERSION is not set"; exit 1; fi

.PHONY: check-github-token
check-github-token:
	@if [ -z "${GITHUB_TOKEN}" ]; then echo "GITHUB_TOKEN is not set"; exit 1; fi

.PHONY: release
release: clean-release check-release-tag $(RELEASE_DIR)  ## Builds and push container images using the latest git tag for the commit.
	git checkout "${RELEASE_TAG}"
	CORE_CONTROLLER_IMG=$(PROD_REGISTRY)/$(CORE_IMAGE_NAME) $(MAKE) release-manifests
	$(MAKE) release-templates

.PHONY: release-manifests
release-manifests:
	$(MAKE) $(RELEASE_DIR)/$(CORE_MANIFEST_FILE).yaml TAG=$(RELEASE_TAG)
	# Add metadata to the release artifacts
	cp metadata.yaml $(RELEASE_DIR)/metadata.yaml

.PHONY: release-staging
release-staging: ## Builds and push container images and manifests to the staging bucket.
	$(MAKE) docker-build-all
	$(MAKE) docker-push-all
	$(MAKE) release-alias-tag
	$(MAKE) staging-manifests
	$(MAKE) release-templates
	$(MAKE) upload-staging-artifacts

.PHONY: staging-manifests
staging-manifests:
	$(MAKE) $(RELEASE_DIR)/$(CORE_MANIFEST_FILE).yaml TAG=$(RELEASE_ALIAS_TAG)
	cp metadata.yaml $(RELEASE_DIR)/metadata.yaml

.PHONY: upload-staging-artifacts
upload-staging-artifacts: ## Upload release artifacts to the staging bucket
	gsutil cp $(RELEASE_DIR)/* gs://$(BUCKET)/components/$(RELEASE_ALIAS_TAG)

.PHONY: release-alias-tag
release-alias-tag: # Adds the tag to the last build tag.
	gcloud container images add-tag -q $(CORE_CONTROLLER_IMG):$(TAG) $(CORE_CONTROLLER_IMG):$(RELEASE_ALIAS_TAG)

.PHONY: release-templates
release-templates: $(RELEASE_DIR) ## Generate release templates
	cp templates/cluster-template*.yaml $(RELEASE_DIR)/

IMAGE_PATCH_DIR := $(ARTIFACTS)/image-patch

$(IMAGE_PATCH_DIR): $(ARTIFACTS)
	mkdir -p $@

.PHONY: $(RELEASE_DIR)/$(CORE_MANIFEST_FILE).yaml
$(RELEASE_DIR)/$(CORE_MANIFEST_FILE).yaml:
	$(MAKE) compiled-manifest \
		PROVIDER=$(CORE_MANIFEST_FILE) \
		OLD_IMG=$(CORE_CONTROLLER_ORIGINAL_IMG) \
		MANIFEST_IMG=$(CORE_CONTROLLER_IMG) \
		CONTROLLER_NAME=$(CORE_CONTROLLER_NAME) \
		PROVIDER_CONFIG_DIR=$(CORE_CONFIG_DIR) \
		NAMESPACE=$(CORE_NAMESPACE) \

.PHONY: compiled-manifest
compiled-manifest: $(RELEASE_DIR) $(KUSTOMIZE)
	$(MAKE) image-patch-source-manifest
	$(MAKE) image-patch-kustomization
	$(KUSTOMIZE) build $(IMAGE_PATCH_DIR)/$(PROVIDER) > $(RELEASE_DIR)/$(PROVIDER).yaml

.PHONY: image-patch-source-manifest
image-patch-source-manifest: $(IMAGE_PATCH_DIR) $(KUSTOMIZE)
	mkdir -p $(IMAGE_PATCH_DIR)/$(PROVIDER)
	$(KUSTOMIZE) build $(PROVIDER_CONFIG_DIR) > $(IMAGE_PATCH_DIR)/$(PROVIDER)/source-manifest.yaml

.PHONY: image-patch-kustomization
image-patch-kustomization: $(IMAGE_PATCH_DIR)
	mkdir -p $(IMAGE_PATCH_DIR)/$(PROVIDER)
	$(MAKE) image-patch-kustomization-without-webhook

.PHONY: image-patch-kustomization-without-webhook
image-patch-kustomization-without-webhook: $(IMAGE_PATCH_DIR) $(GOJQ)
	mkdir -p $(IMAGE_PATCH_DIR)/$(PROVIDER)
	$(GOJQ) --yaml-input --yaml-output '.images[0]={"name":"$(OLD_IMG)","newName":"$(MANIFEST_IMG)","newTag":"$(TAG)"}' \
		"hack/image-patch/kustomization.yaml" > $(IMAGE_PATCH_DIR)/$(PROVIDER)/kustomization.yaml

## --------------------------------------
## Cleanup / Verification
## --------------------------------------

.PHONY: clean
clean: ## Remove all generated files
	$(MAKE) -C hack/tools clean
	$(MAKE) clean-temporary

.PHONY: clean-temporary
clean-temporary: ## Remove all temporary files and folders
	rm -f minikube.kubeconfig
	rm -f kubeconfig
	rm -rf _artifacts

.PHONY: clean-release
clean-release: ## Remove the release folder
	rm -rf $(RELEASE_DIR)

.PHONY: clean-generated-conversions
clean-generated-conversions: ## Remove files generated by conversion-gen from the mentioned dirs
	(IFS=','; for i in $(SRC_DIRS); do find $$i -type f -name 'zz_generated.conversion*' -exec rm -f {} \;; done)

.PHONY: verify
verify:
	./hack/verify-boilerplate.sh
	./hack/verify-doctoc.sh
	./hack/verify-shellcheck.sh
	#TODO(mkumatag): looking for tiltfile, fix it when we use tilt
	#./hack/verify-starlark.sh
	$(MAKE) verify-modules
	$(MAKE) verify-gen
	$(MAKE) verify-conversions

.PHONY: modules
modules: ## Runs go mod to ensure modules are up to date.
	go mod tidy
	cd $(TOOLS_DIR); go mod tidy

.PHONY: verify-modules
verify-modules: modules
	@if !(git diff --quiet HEAD -- go.sum go.mod $(TOOLS_DIR)/go.mod $(TOOLS_DIR)/go.sum); then \
		git diff; \
		echo "go module files are out of date"; exit 1; \
	fi
	@if (find . -name 'go.mod' | xargs -n1 grep -q -i 'k8s.io/client-go.*+incompatible'); then \
		find . -name "go.mod" -exec grep -i 'k8s.io/client-go.*+incompatible' {} \; -print; \
		echo "go module contains an incompatible client-go version"; exit 1; \
	fi

.PHONY: verify-gen
verify-gen: generate
	@if !(git diff --quiet HEAD); then \
		git diff; \
		echo "generated files are out of date, run make generate"; exit 1; \
	fi

.PHONY: verify-conversions
verify-conversions: $(CONVERSION_VERIFIER)
	$(CONVERSION_VERIFIER)
