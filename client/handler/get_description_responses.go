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

// GetDescriptionReader is a Reader for the GetDescription structure.
type GetDescriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDescriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDescriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDescriptionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDescriptionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDescriptionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDescriptionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDescriptionOK creates a GetDescriptionOK with default headers values
func NewGetDescriptionOK() *GetDescriptionOK {
	return &GetDescriptionOK{}
}

/*
GetDescriptionOK describes a response with status code 200, with default header values.

OK
*/
type GetDescriptionOK struct {
	Payload string
}

// IsSuccess returns true when this get description o k response has a 2xx status code
func (o *GetDescriptionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get description o k response has a 3xx status code
func (o *GetDescriptionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get description o k response has a 4xx status code
func (o *GetDescriptionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get description o k response has a 5xx status code
func (o *GetDescriptionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get description o k response a status code equal to that given
func (o *GetDescriptionOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDescriptionOK) Error() string {
	return fmt.Sprintf("[GET /:taskId][%d] getDescriptionOK  %+v", 200, o.Payload)
}

func (o *GetDescriptionOK) String() string {
	return fmt.Sprintf("[GET /:taskId][%d] getDescriptionOK  %+v", 200, o.Payload)
}

func (o *GetDescriptionOK) GetPayload() string {
	return o.Payload
}

func (o *GetDescriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDescriptionBadRequest creates a GetDescriptionBadRequest with default headers values
func NewGetDescriptionBadRequest() *GetDescriptionBadRequest {
	return &GetDescriptionBadRequest{}
}

/*
GetDescriptionBadRequest describes a response with status code 400, with default header values.

bad request
*/
type GetDescriptionBadRequest struct {
	Payload string
}

// IsSuccess returns true when this get description bad request response has a 2xx status code
func (o *GetDescriptionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get description bad request response has a 3xx status code
func (o *GetDescriptionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get description bad request response has a 4xx status code
func (o *GetDescriptionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get description bad request response has a 5xx status code
func (o *GetDescriptionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get description bad request response a status code equal to that given
func (o *GetDescriptionBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetDescriptionBadRequest) Error() string {
	return fmt.Sprintf("[GET /:taskId][%d] getDescriptionBadRequest  %+v", 400, o.Payload)
}

func (o *GetDescriptionBadRequest) String() string {
	return fmt.Sprintf("[GET /:taskId][%d] getDescriptionBadRequest  %+v", 400, o.Payload)
}

func (o *GetDescriptionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetDescriptionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDescriptionForbidden creates a GetDescriptionForbidden with default headers values
func NewGetDescriptionForbidden() *GetDescriptionForbidden {
	return &GetDescriptionForbidden{}
}

/*
GetDescriptionForbidden describes a response with status code 403, with default header values.

forbidden
*/
type GetDescriptionForbidden struct {
	Payload string
}

// IsSuccess returns true when this get description forbidden response has a 2xx status code
func (o *GetDescriptionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get description forbidden response has a 3xx status code
func (o *GetDescriptionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get description forbidden response has a 4xx status code
func (o *GetDescriptionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get description forbidden response has a 5xx status code
func (o *GetDescriptionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get description forbidden response a status code equal to that given
func (o *GetDescriptionForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetDescriptionForbidden) Error() string {
	return fmt.Sprintf("[GET /:taskId][%d] getDescriptionForbidden  %+v", 403, o.Payload)
}

func (o *GetDescriptionForbidden) String() string {
	return fmt.Sprintf("[GET /:taskId][%d] getDescriptionForbidden  %+v", 403, o.Payload)
}

func (o *GetDescriptionForbidden) GetPayload() string {
	return o.Payload
}

func (o *GetDescriptionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDescriptionNotFound creates a GetDescriptionNotFound with default headers values
func NewGetDescriptionNotFound() *GetDescriptionNotFound {
	return &GetDescriptionNotFound{}
}

/*
GetDescriptionNotFound describes a response with status code 404, with default header values.

not found
*/
type GetDescriptionNotFound struct {
	Payload string
}

// IsSuccess returns true when this get description not found response has a 2xx status code
func (o *GetDescriptionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get description not found response has a 3xx status code
func (o *GetDescriptionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get description not found response has a 4xx status code
func (o *GetDescriptionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get description not found response has a 5xx status code
func (o *GetDescriptionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get description not found response a status code equal to that given
func (o *GetDescriptionNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetDescriptionNotFound) Error() string {
	return fmt.Sprintf("[GET /:taskId][%d] getDescriptionNotFound  %+v", 404, o.Payload)
}

func (o *GetDescriptionNotFound) String() string {
	return fmt.Sprintf("[GET /:taskId][%d] getDescriptionNotFound  %+v", 404, o.Payload)
}

func (o *GetDescriptionNotFound) GetPayload() string {
	return o.Payload
}

func (o *GetDescriptionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDescriptionInternalServerError creates a GetDescriptionInternalServerError with default headers values
func NewGetDescriptionInternalServerError() *GetDescriptionInternalServerError {
	return &GetDescriptionInternalServerError{}
}

/*
GetDescriptionInternalServerError describes a response with status code 500, with default header values.

internal error
*/
type GetDescriptionInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this get description internal server error response has a 2xx status code
func (o *GetDescriptionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get description internal server error response has a 3xx status code
func (o *GetDescriptionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get description internal server error response has a 4xx status code
func (o *GetDescriptionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get description internal server error response has a 5xx status code
func (o *GetDescriptionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get description internal server error response a status code equal to that given
func (o *GetDescriptionInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetDescriptionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /:taskId][%d] getDescriptionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDescriptionInternalServerError) String() string {
	return fmt.Sprintf("[GET /:taskId][%d] getDescriptionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDescriptionInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *GetDescriptionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}