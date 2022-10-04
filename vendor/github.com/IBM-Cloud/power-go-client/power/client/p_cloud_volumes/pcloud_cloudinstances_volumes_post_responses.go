// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudCloudinstancesVolumesPostReader is a Reader for the PcloudCloudinstancesVolumesPost structure.
type PcloudCloudinstancesVolumesPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudinstancesVolumesPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPcloudCloudinstancesVolumesPostAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudinstancesVolumesPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudinstancesVolumesPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudCloudinstancesVolumesPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudCloudinstancesVolumesPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudCloudinstancesVolumesPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudinstancesVolumesPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudCloudinstancesVolumesPostAccepted creates a PcloudCloudinstancesVolumesPostAccepted with default headers values
func NewPcloudCloudinstancesVolumesPostAccepted() *PcloudCloudinstancesVolumesPostAccepted {
	return &PcloudCloudinstancesVolumesPostAccepted{}
}

/* PcloudCloudinstancesVolumesPostAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudCloudinstancesVolumesPostAccepted struct {
	Payload *models.Volume
}

func (o *PcloudCloudinstancesVolumesPostAccepted) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes][%d] pcloudCloudinstancesVolumesPostAccepted  %+v", 202, o.Payload)
}
func (o *PcloudCloudinstancesVolumesPostAccepted) GetPayload() *models.Volume {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesPostAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Volume)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesPostBadRequest creates a PcloudCloudinstancesVolumesPostBadRequest with default headers values
func NewPcloudCloudinstancesVolumesPostBadRequest() *PcloudCloudinstancesVolumesPostBadRequest {
	return &PcloudCloudinstancesVolumesPostBadRequest{}
}

/* PcloudCloudinstancesVolumesPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudinstancesVolumesPostBadRequest struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesVolumesPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes][%d] pcloudCloudinstancesVolumesPostBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudCloudinstancesVolumesPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesPostUnauthorized creates a PcloudCloudinstancesVolumesPostUnauthorized with default headers values
func NewPcloudCloudinstancesVolumesPostUnauthorized() *PcloudCloudinstancesVolumesPostUnauthorized {
	return &PcloudCloudinstancesVolumesPostUnauthorized{}
}

/* PcloudCloudinstancesVolumesPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudinstancesVolumesPostUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesVolumesPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes][%d] pcloudCloudinstancesVolumesPostUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudCloudinstancesVolumesPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesPostForbidden creates a PcloudCloudinstancesVolumesPostForbidden with default headers values
func NewPcloudCloudinstancesVolumesPostForbidden() *PcloudCloudinstancesVolumesPostForbidden {
	return &PcloudCloudinstancesVolumesPostForbidden{}
}

/* PcloudCloudinstancesVolumesPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudCloudinstancesVolumesPostForbidden struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesVolumesPostForbidden) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes][%d] pcloudCloudinstancesVolumesPostForbidden  %+v", 403, o.Payload)
}
func (o *PcloudCloudinstancesVolumesPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesPostConflict creates a PcloudCloudinstancesVolumesPostConflict with default headers values
func NewPcloudCloudinstancesVolumesPostConflict() *PcloudCloudinstancesVolumesPostConflict {
	return &PcloudCloudinstancesVolumesPostConflict{}
}

/* PcloudCloudinstancesVolumesPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudCloudinstancesVolumesPostConflict struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesVolumesPostConflict) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes][%d] pcloudCloudinstancesVolumesPostConflict  %+v", 409, o.Payload)
}
func (o *PcloudCloudinstancesVolumesPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesPostUnprocessableEntity creates a PcloudCloudinstancesVolumesPostUnprocessableEntity with default headers values
func NewPcloudCloudinstancesVolumesPostUnprocessableEntity() *PcloudCloudinstancesVolumesPostUnprocessableEntity {
	return &PcloudCloudinstancesVolumesPostUnprocessableEntity{}
}

/* PcloudCloudinstancesVolumesPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudCloudinstancesVolumesPostUnprocessableEntity struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesVolumesPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes][%d] pcloudCloudinstancesVolumesPostUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *PcloudCloudinstancesVolumesPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesPostInternalServerError creates a PcloudCloudinstancesVolumesPostInternalServerError with default headers values
func NewPcloudCloudinstancesVolumesPostInternalServerError() *PcloudCloudinstancesVolumesPostInternalServerError {
	return &PcloudCloudinstancesVolumesPostInternalServerError{}
}

/* PcloudCloudinstancesVolumesPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudinstancesVolumesPostInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesVolumesPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes][%d] pcloudCloudinstancesVolumesPostInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudCloudinstancesVolumesPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
