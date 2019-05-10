// Code generated by go-swagger; DO NOT EDIT.

package employment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/nad2000/uoa-orcidhub/models"
)

// GetEmployeeUpiOrIDReader is a Reader for the GetEmployeeUpiOrID structure.
type GetEmployeeUpiOrIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEmployeeUpiOrIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEmployeeUpiOrIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetEmployeeUpiOrIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetEmployeeUpiOrIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetEmployeeUpiOrIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetEmployeeUpiOrIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 501:
		result := NewGetEmployeeUpiOrIDNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEmployeeUpiOrIDOK creates a GetEmployeeUpiOrIDOK with default headers values
func NewGetEmployeeUpiOrIDOK() *GetEmployeeUpiOrIDOK {
	return &GetEmployeeUpiOrIDOK{}
}

/*GetEmployeeUpiOrIDOK handles this case with default header values.

hremployee response
*/
type GetEmployeeUpiOrIDOK struct {
	Payload *models.Hremployee
}

func (o *GetEmployeeUpiOrIDOK) Error() string {
	return fmt.Sprintf("[GET /employee/{upiOrId}][%d] getEmployeeUpiOrIdOK  %+v", 200, o.Payload)
}

func (o *GetEmployeeUpiOrIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Hremployee)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEmployeeUpiOrIDBadRequest creates a GetEmployeeUpiOrIDBadRequest with default headers values
func NewGetEmployeeUpiOrIDBadRequest() *GetEmployeeUpiOrIDBadRequest {
	return &GetEmployeeUpiOrIDBadRequest{}
}

/*GetEmployeeUpiOrIDBadRequest handles this case with default header values.

Bad request (unsupported id)
*/
type GetEmployeeUpiOrIDBadRequest struct {
	Payload *models.Error
}

func (o *GetEmployeeUpiOrIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /employee/{upiOrId}][%d] getEmployeeUpiOrIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetEmployeeUpiOrIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEmployeeUpiOrIDUnauthorized creates a GetEmployeeUpiOrIDUnauthorized with default headers values
func NewGetEmployeeUpiOrIDUnauthorized() *GetEmployeeUpiOrIDUnauthorized {
	return &GetEmployeeUpiOrIDUnauthorized{}
}

/*GetEmployeeUpiOrIDUnauthorized handles this case with default header values.

Unauthorized
*/
type GetEmployeeUpiOrIDUnauthorized struct {
}

func (o *GetEmployeeUpiOrIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /employee/{upiOrId}][%d] getEmployeeUpiOrIdUnauthorized ", 401)
}

func (o *GetEmployeeUpiOrIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEmployeeUpiOrIDNotFound creates a GetEmployeeUpiOrIDNotFound with default headers values
func NewGetEmployeeUpiOrIDNotFound() *GetEmployeeUpiOrIDNotFound {
	return &GetEmployeeUpiOrIDNotFound{}
}

/*GetEmployeeUpiOrIDNotFound handles this case with default header values.

employee not found
*/
type GetEmployeeUpiOrIDNotFound struct {
}

func (o *GetEmployeeUpiOrIDNotFound) Error() string {
	return fmt.Sprintf("[GET /employee/{upiOrId}][%d] getEmployeeUpiOrIdNotFound ", 404)
}

func (o *GetEmployeeUpiOrIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEmployeeUpiOrIDInternalServerError creates a GetEmployeeUpiOrIDInternalServerError with default headers values
func NewGetEmployeeUpiOrIDInternalServerError() *GetEmployeeUpiOrIDInternalServerError {
	return &GetEmployeeUpiOrIDInternalServerError{}
}

/*GetEmployeeUpiOrIDInternalServerError handles this case with default header values.

Unexpected error
*/
type GetEmployeeUpiOrIDInternalServerError struct {
}

func (o *GetEmployeeUpiOrIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /employee/{upiOrId}][%d] getEmployeeUpiOrIdInternalServerError ", 500)
}

func (o *GetEmployeeUpiOrIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEmployeeUpiOrIDNotImplemented creates a GetEmployeeUpiOrIDNotImplemented with default headers values
func NewGetEmployeeUpiOrIDNotImplemented() *GetEmployeeUpiOrIDNotImplemented {
	return &GetEmployeeUpiOrIDNotImplemented{}
}

/*GetEmployeeUpiOrIDNotImplemented handles this case with default header values.

Unsupported operation (not implemented)
*/
type GetEmployeeUpiOrIDNotImplemented struct {
}

func (o *GetEmployeeUpiOrIDNotImplemented) Error() string {
	return fmt.Sprintf("[GET /employee/{upiOrId}][%d] getEmployeeUpiOrIdNotImplemented ", 501)
}

func (o *GetEmployeeUpiOrIDNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
