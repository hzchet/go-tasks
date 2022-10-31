// Code generated by go-swagger; DO NOT EDIT.

package handler

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateReader is a Reader for the Create structure.
type CreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateCreated creates a CreateCreated with default headers values
func NewCreateCreated() *CreateCreated {
	return &CreateCreated{}
}

/*
CreateCreated describes a response with status code 201, with default header values.

Created
*/
type CreateCreated struct {
	Payload string
}

// IsSuccess returns true when this create created response has a 2xx status code
func (o *CreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create created response has a 3xx status code
func (o *CreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create created response has a 4xx status code
func (o *CreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create created response has a 5xx status code
func (o *CreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create created response a status code equal to that given
func (o *CreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateCreated) Error() string {
	return fmt.Sprintf("[POST /][%d] createCreated  %+v", 201, o.Payload)
}

func (o *CreateCreated) String() string {
	return fmt.Sprintf("[POST /][%d] createCreated  %+v", 201, o.Payload)
}

func (o *CreateCreated) GetPayload() string {
	return o.Payload
}

func (o *CreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBadRequest creates a CreateBadRequest with default headers values
func NewCreateBadRequest() *CreateBadRequest {
	return &CreateBadRequest{}
}

/*
CreateBadRequest describes a response with status code 400, with default header values.

bad request
*/
type CreateBadRequest struct {
	Payload string
}

// IsSuccess returns true when this create bad request response has a 2xx status code
func (o *CreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create bad request response has a 3xx status code
func (o *CreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create bad request response has a 4xx status code
func (o *CreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create bad request response has a 5xx status code
func (o *CreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create bad request response a status code equal to that given
func (o *CreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /][%d] createBadRequest  %+v", 400, o.Payload)
}

func (o *CreateBadRequest) String() string {
	return fmt.Sprintf("[POST /][%d] createBadRequest  %+v", 400, o.Payload)
}

func (o *CreateBadRequest) GetPayload() string {
	return o.Payload
}

func (o *CreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateForbidden creates a CreateForbidden with default headers values
func NewCreateForbidden() *CreateForbidden {
	return &CreateForbidden{}
}

/*
CreateForbidden describes a response with status code 403, with default header values.

forbidden
*/
type CreateForbidden struct {
	Payload string
}

// IsSuccess returns true when this create forbidden response has a 2xx status code
func (o *CreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create forbidden response has a 3xx status code
func (o *CreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create forbidden response has a 4xx status code
func (o *CreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create forbidden response has a 5xx status code
func (o *CreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create forbidden response a status code equal to that given
func (o *CreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateForbidden) Error() string {
	return fmt.Sprintf("[POST /][%d] createForbidden  %+v", 403, o.Payload)
}

func (o *CreateForbidden) String() string {
	return fmt.Sprintf("[POST /][%d] createForbidden  %+v", 403, o.Payload)
}

func (o *CreateForbidden) GetPayload() string {
	return o.Payload
}

func (o *CreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInternalServerError creates a CreateInternalServerError with default headers values
func NewCreateInternalServerError() *CreateInternalServerError {
	return &CreateInternalServerError{}
}

/*
CreateInternalServerError describes a response with status code 500, with default header values.

internal error
*/
type CreateInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this create internal server error response has a 2xx status code
func (o *CreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create internal server error response has a 3xx status code
func (o *CreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create internal server error response has a 4xx status code
func (o *CreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create internal server error response has a 5xx status code
func (o *CreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create internal server error response a status code equal to that given
func (o *CreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /][%d] createInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /][%d] createInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *CreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}