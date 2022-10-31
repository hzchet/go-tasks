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

// ApproveReader is a Reader for the Approve structure.
type ApproveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApproveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApproveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewApproveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewApproveForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewApproveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewApproveOK creates a ApproveOK with default headers values
func NewApproveOK() *ApproveOK {
	return &ApproveOK{}
}

/*
ApproveOK describes a response with status code 200, with default header values.

OK
*/
type ApproveOK struct {
}

// IsSuccess returns true when this approve o k response has a 2xx status code
func (o *ApproveOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this approve o k response has a 3xx status code
func (o *ApproveOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this approve o k response has a 4xx status code
func (o *ApproveOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this approve o k response has a 5xx status code
func (o *ApproveOK) IsServerError() bool {
	return false
}

// IsCode returns true when this approve o k response a status code equal to that given
func (o *ApproveOK) IsCode(code int) bool {
	return code == 200
}

func (o *ApproveOK) Error() string {
	return fmt.Sprintf("[POST /approve/:taskId][%d] approveOK ", 200)
}

func (o *ApproveOK) String() string {
	return fmt.Sprintf("[POST /approve/:taskId][%d] approveOK ", 200)
}

func (o *ApproveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewApproveBadRequest creates a ApproveBadRequest with default headers values
func NewApproveBadRequest() *ApproveBadRequest {
	return &ApproveBadRequest{}
}

/*
ApproveBadRequest describes a response with status code 400, with default header values.

bad request
*/
type ApproveBadRequest struct {
	Payload string
}

// IsSuccess returns true when this approve bad request response has a 2xx status code
func (o *ApproveBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this approve bad request response has a 3xx status code
func (o *ApproveBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this approve bad request response has a 4xx status code
func (o *ApproveBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this approve bad request response has a 5xx status code
func (o *ApproveBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this approve bad request response a status code equal to that given
func (o *ApproveBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ApproveBadRequest) Error() string {
	return fmt.Sprintf("[POST /approve/:taskId][%d] approveBadRequest  %+v", 400, o.Payload)
}

func (o *ApproveBadRequest) String() string {
	return fmt.Sprintf("[POST /approve/:taskId][%d] approveBadRequest  %+v", 400, o.Payload)
}

func (o *ApproveBadRequest) GetPayload() string {
	return o.Payload
}

func (o *ApproveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApproveForbidden creates a ApproveForbidden with default headers values
func NewApproveForbidden() *ApproveForbidden {
	return &ApproveForbidden{}
}

/*
ApproveForbidden describes a response with status code 403, with default header values.

forbidden
*/
type ApproveForbidden struct {
	Payload string
}

// IsSuccess returns true when this approve forbidden response has a 2xx status code
func (o *ApproveForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this approve forbidden response has a 3xx status code
func (o *ApproveForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this approve forbidden response has a 4xx status code
func (o *ApproveForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this approve forbidden response has a 5xx status code
func (o *ApproveForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this approve forbidden response a status code equal to that given
func (o *ApproveForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ApproveForbidden) Error() string {
	return fmt.Sprintf("[POST /approve/:taskId][%d] approveForbidden  %+v", 403, o.Payload)
}

func (o *ApproveForbidden) String() string {
	return fmt.Sprintf("[POST /approve/:taskId][%d] approveForbidden  %+v", 403, o.Payload)
}

func (o *ApproveForbidden) GetPayload() string {
	return o.Payload
}

func (o *ApproveForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApproveInternalServerError creates a ApproveInternalServerError with default headers values
func NewApproveInternalServerError() *ApproveInternalServerError {
	return &ApproveInternalServerError{}
}

/*
ApproveInternalServerError describes a response with status code 500, with default header values.

internal error
*/
type ApproveInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this approve internal server error response has a 2xx status code
func (o *ApproveInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this approve internal server error response has a 3xx status code
func (o *ApproveInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this approve internal server error response has a 4xx status code
func (o *ApproveInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this approve internal server error response has a 5xx status code
func (o *ApproveInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this approve internal server error response a status code equal to that given
func (o *ApproveInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ApproveInternalServerError) Error() string {
	return fmt.Sprintf("[POST /approve/:taskId][%d] approveInternalServerError  %+v", 500, o.Payload)
}

func (o *ApproveInternalServerError) String() string {
	return fmt.Sprintf("[POST /approve/:taskId][%d] approveInternalServerError  %+v", 500, o.Payload)
}

func (o *ApproveInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *ApproveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
