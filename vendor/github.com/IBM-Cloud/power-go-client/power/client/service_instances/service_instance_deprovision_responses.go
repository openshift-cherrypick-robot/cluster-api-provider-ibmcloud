// Code generated by go-swagger; DO NOT EDIT.

package service_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// ServiceInstanceDeprovisionReader is a Reader for the ServiceInstanceDeprovision structure.
type ServiceInstanceDeprovisionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceInstanceDeprovisionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceInstanceDeprovisionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewServiceInstanceDeprovisionAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServiceInstanceDeprovisionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewServiceInstanceDeprovisionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewServiceInstanceDeprovisionGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewServiceInstanceDeprovisionUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewServiceInstanceDeprovisionOK creates a ServiceInstanceDeprovisionOK with default headers values
func NewServiceInstanceDeprovisionOK() *ServiceInstanceDeprovisionOK {
	return &ServiceInstanceDeprovisionOK{}
}

/*
ServiceInstanceDeprovisionOK describes a response with status code 200, with default header values.

OK
*/
type ServiceInstanceDeprovisionOK struct {
	Payload models.Object
}

// IsSuccess returns true when this service instance deprovision o k response has a 2xx status code
func (o *ServiceInstanceDeprovisionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service instance deprovision o k response has a 3xx status code
func (o *ServiceInstanceDeprovisionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service instance deprovision o k response has a 4xx status code
func (o *ServiceInstanceDeprovisionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service instance deprovision o k response has a 5xx status code
func (o *ServiceInstanceDeprovisionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service instance deprovision o k response a status code equal to that given
func (o *ServiceInstanceDeprovisionOK) IsCode(code int) bool {
	return code == 200
}

func (o *ServiceInstanceDeprovisionOK) Error() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}][%d] serviceInstanceDeprovisionOK  %+v", 200, o.Payload)
}

func (o *ServiceInstanceDeprovisionOK) String() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}][%d] serviceInstanceDeprovisionOK  %+v", 200, o.Payload)
}

func (o *ServiceInstanceDeprovisionOK) GetPayload() models.Object {
	return o.Payload
}

func (o *ServiceInstanceDeprovisionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceInstanceDeprovisionAccepted creates a ServiceInstanceDeprovisionAccepted with default headers values
func NewServiceInstanceDeprovisionAccepted() *ServiceInstanceDeprovisionAccepted {
	return &ServiceInstanceDeprovisionAccepted{}
}

/*
ServiceInstanceDeprovisionAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ServiceInstanceDeprovisionAccepted struct {
	Payload *models.AsyncOperation
}

// IsSuccess returns true when this service instance deprovision accepted response has a 2xx status code
func (o *ServiceInstanceDeprovisionAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service instance deprovision accepted response has a 3xx status code
func (o *ServiceInstanceDeprovisionAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service instance deprovision accepted response has a 4xx status code
func (o *ServiceInstanceDeprovisionAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this service instance deprovision accepted response has a 5xx status code
func (o *ServiceInstanceDeprovisionAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this service instance deprovision accepted response a status code equal to that given
func (o *ServiceInstanceDeprovisionAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *ServiceInstanceDeprovisionAccepted) Error() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}][%d] serviceInstanceDeprovisionAccepted  %+v", 202, o.Payload)
}

func (o *ServiceInstanceDeprovisionAccepted) String() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}][%d] serviceInstanceDeprovisionAccepted  %+v", 202, o.Payload)
}

func (o *ServiceInstanceDeprovisionAccepted) GetPayload() *models.AsyncOperation {
	return o.Payload
}

func (o *ServiceInstanceDeprovisionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AsyncOperation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceInstanceDeprovisionBadRequest creates a ServiceInstanceDeprovisionBadRequest with default headers values
func NewServiceInstanceDeprovisionBadRequest() *ServiceInstanceDeprovisionBadRequest {
	return &ServiceInstanceDeprovisionBadRequest{}
}

/*
ServiceInstanceDeprovisionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ServiceInstanceDeprovisionBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this service instance deprovision bad request response has a 2xx status code
func (o *ServiceInstanceDeprovisionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service instance deprovision bad request response has a 3xx status code
func (o *ServiceInstanceDeprovisionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service instance deprovision bad request response has a 4xx status code
func (o *ServiceInstanceDeprovisionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this service instance deprovision bad request response has a 5xx status code
func (o *ServiceInstanceDeprovisionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this service instance deprovision bad request response a status code equal to that given
func (o *ServiceInstanceDeprovisionBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ServiceInstanceDeprovisionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}][%d] serviceInstanceDeprovisionBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceInstanceDeprovisionBadRequest) String() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}][%d] serviceInstanceDeprovisionBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceInstanceDeprovisionBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceInstanceDeprovisionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceInstanceDeprovisionUnauthorized creates a ServiceInstanceDeprovisionUnauthorized with default headers values
func NewServiceInstanceDeprovisionUnauthorized() *ServiceInstanceDeprovisionUnauthorized {
	return &ServiceInstanceDeprovisionUnauthorized{}
}

/*
ServiceInstanceDeprovisionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ServiceInstanceDeprovisionUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this service instance deprovision unauthorized response has a 2xx status code
func (o *ServiceInstanceDeprovisionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service instance deprovision unauthorized response has a 3xx status code
func (o *ServiceInstanceDeprovisionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service instance deprovision unauthorized response has a 4xx status code
func (o *ServiceInstanceDeprovisionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this service instance deprovision unauthorized response has a 5xx status code
func (o *ServiceInstanceDeprovisionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this service instance deprovision unauthorized response a status code equal to that given
func (o *ServiceInstanceDeprovisionUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ServiceInstanceDeprovisionUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}][%d] serviceInstanceDeprovisionUnauthorized  %+v", 401, o.Payload)
}

func (o *ServiceInstanceDeprovisionUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}][%d] serviceInstanceDeprovisionUnauthorized  %+v", 401, o.Payload)
}

func (o *ServiceInstanceDeprovisionUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceInstanceDeprovisionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceInstanceDeprovisionGone creates a ServiceInstanceDeprovisionGone with default headers values
func NewServiceInstanceDeprovisionGone() *ServiceInstanceDeprovisionGone {
	return &ServiceInstanceDeprovisionGone{}
}

/*
ServiceInstanceDeprovisionGone describes a response with status code 410, with default header values.

Gone
*/
type ServiceInstanceDeprovisionGone struct {
	Payload *models.Error
}

// IsSuccess returns true when this service instance deprovision gone response has a 2xx status code
func (o *ServiceInstanceDeprovisionGone) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service instance deprovision gone response has a 3xx status code
func (o *ServiceInstanceDeprovisionGone) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service instance deprovision gone response has a 4xx status code
func (o *ServiceInstanceDeprovisionGone) IsClientError() bool {
	return true
}

// IsServerError returns true when this service instance deprovision gone response has a 5xx status code
func (o *ServiceInstanceDeprovisionGone) IsServerError() bool {
	return false
}

// IsCode returns true when this service instance deprovision gone response a status code equal to that given
func (o *ServiceInstanceDeprovisionGone) IsCode(code int) bool {
	return code == 410
}

func (o *ServiceInstanceDeprovisionGone) Error() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}][%d] serviceInstanceDeprovisionGone  %+v", 410, o.Payload)
}

func (o *ServiceInstanceDeprovisionGone) String() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}][%d] serviceInstanceDeprovisionGone  %+v", 410, o.Payload)
}

func (o *ServiceInstanceDeprovisionGone) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceInstanceDeprovisionGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceInstanceDeprovisionUnprocessableEntity creates a ServiceInstanceDeprovisionUnprocessableEntity with default headers values
func NewServiceInstanceDeprovisionUnprocessableEntity() *ServiceInstanceDeprovisionUnprocessableEntity {
	return &ServiceInstanceDeprovisionUnprocessableEntity{}
}

/*
ServiceInstanceDeprovisionUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type ServiceInstanceDeprovisionUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this service instance deprovision unprocessable entity response has a 2xx status code
func (o *ServiceInstanceDeprovisionUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service instance deprovision unprocessable entity response has a 3xx status code
func (o *ServiceInstanceDeprovisionUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service instance deprovision unprocessable entity response has a 4xx status code
func (o *ServiceInstanceDeprovisionUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this service instance deprovision unprocessable entity response has a 5xx status code
func (o *ServiceInstanceDeprovisionUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this service instance deprovision unprocessable entity response a status code equal to that given
func (o *ServiceInstanceDeprovisionUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *ServiceInstanceDeprovisionUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}][%d] serviceInstanceDeprovisionUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ServiceInstanceDeprovisionUnprocessableEntity) String() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}][%d] serviceInstanceDeprovisionUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ServiceInstanceDeprovisionUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceInstanceDeprovisionUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
