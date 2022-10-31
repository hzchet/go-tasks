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

// GetTasksReader is a Reader for the GetTasks structure.
type GetTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTasksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTasksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTasksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTasksOK creates a GetTasksOK with default headers values
func NewGetTasksOK() *GetTasksOK {
	return &GetTasksOK{}
}

/*
GetTasksOK describes a response with status code 200, with default header values.

OK
*/
type GetTasksOK struct {
}

// IsSuccess returns true when this get tasks o k response has a 2xx status code
func (o *GetTasksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get tasks o k response has a 3xx status code
func (o *GetTasksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tasks o k response has a 4xx status code
func (o *GetTasksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tasks o k response has a 5xx status code
func (o *GetTasksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get tasks o k response a status code equal to that given
func (o *GetTasksOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetTasksOK) Error() string {
	return fmt.Sprintf("[GET /][%d] getTasksOK ", 200)
}

func (o *GetTasksOK) String() string {
	return fmt.Sprintf("[GET /][%d] getTasksOK ", 200)
}

func (o *GetTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTasksBadRequest creates a GetTasksBadRequest with default headers values
func NewGetTasksBadRequest() *GetTasksBadRequest {
	return &GetTasksBadRequest{}
}

/*
GetTasksBadRequest describes a response with status code 400, with default header values.

bad request
*/
type GetTasksBadRequest struct {
	Payload string
}

// IsSuccess returns true when this get tasks bad request response has a 2xx status code
func (o *GetTasksBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tasks bad request response has a 3xx status code
func (o *GetTasksBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tasks bad request response has a 4xx status code
func (o *GetTasksBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tasks bad request response has a 5xx status code
func (o *GetTasksBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get tasks bad request response a status code equal to that given
func (o *GetTasksBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetTasksBadRequest) Error() string {
	return fmt.Sprintf("[GET /][%d] getTasksBadRequest  %+v", 400, o.Payload)
}

func (o *GetTasksBadRequest) String() string {
	return fmt.Sprintf("[GET /][%d] getTasksBadRequest  %+v", 400, o.Payload)
}

func (o *GetTasksBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetTasksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTasksForbidden creates a GetTasksForbidden with default headers values
func NewGetTasksForbidden() *GetTasksForbidden {
	return &GetTasksForbidden{}
}

/*
GetTasksForbidden describes a response with status code 403, with default header values.

forbidden
*/
type GetTasksForbidden struct {
	Payload string
}

// IsSuccess returns true when this get tasks forbidden response has a 2xx status code
func (o *GetTasksForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tasks forbidden response has a 3xx status code
func (o *GetTasksForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tasks forbidden response has a 4xx status code
func (o *GetTasksForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tasks forbidden response has a 5xx status code
func (o *GetTasksForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get tasks forbidden response a status code equal to that given
func (o *GetTasksForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetTasksForbidden) Error() string {
	return fmt.Sprintf("[GET /][%d] getTasksForbidden  %+v", 403, o.Payload)
}

func (o *GetTasksForbidden) String() string {
	return fmt.Sprintf("[GET /][%d] getTasksForbidden  %+v", 403, o.Payload)
}

func (o *GetTasksForbidden) GetPayload() string {
	return o.Payload
}

func (o *GetTasksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTasksInternalServerError creates a GetTasksInternalServerError with default headers values
func NewGetTasksInternalServerError() *GetTasksInternalServerError {
	return &GetTasksInternalServerError{}
}

/*
GetTasksInternalServerError describes a response with status code 500, with default header values.

internal error
*/
type GetTasksInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this get tasks internal server error response has a 2xx status code
func (o *GetTasksInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tasks internal server error response has a 3xx status code
func (o *GetTasksInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tasks internal server error response has a 4xx status code
func (o *GetTasksInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tasks internal server error response has a 5xx status code
func (o *GetTasksInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get tasks internal server error response a status code equal to that given
func (o *GetTasksInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetTasksInternalServerError) Error() string {
	return fmt.Sprintf("[GET /][%d] getTasksInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTasksInternalServerError) String() string {
	return fmt.Sprintf("[GET /][%d] getTasksInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTasksInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *GetTasksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
