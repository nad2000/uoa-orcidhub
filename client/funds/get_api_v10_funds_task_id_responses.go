// Code generated by go-swagger; DO NOT EDIT.

package funds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/nad2000/uoa-orcidhub/models"
)

// GetAPIV10FundsTaskIDReader is a Reader for the GetAPIV10FundsTaskID structure.
type GetAPIV10FundsTaskIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV10FundsTaskIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAPIV10FundsTaskIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetAPIV10FundsTaskIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAPIV10FundsTaskIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAPIV10FundsTaskIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAPIV10FundsTaskIDOK creates a GetAPIV10FundsTaskIDOK with default headers values
func NewGetAPIV10FundsTaskIDOK() *GetAPIV10FundsTaskIDOK {
	return &GetAPIV10FundsTaskIDOK{}
}

/*GetAPIV10FundsTaskIDOK handles this case with default header values.

successful operation
*/
type GetAPIV10FundsTaskIDOK struct {
	Payload *models.FundTask
}

func (o *GetAPIV10FundsTaskIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/funds/{task_id}][%d] getApiV10FundsTaskIdOK  %+v", 200, o.Payload)
}

func (o *GetAPIV10FundsTaskIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FundTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV10FundsTaskIDUnauthorized creates a GetAPIV10FundsTaskIDUnauthorized with default headers values
func NewGetAPIV10FundsTaskIDUnauthorized() *GetAPIV10FundsTaskIDUnauthorized {
	return &GetAPIV10FundsTaskIDUnauthorized{}
}

/*GetAPIV10FundsTaskIDUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAPIV10FundsTaskIDUnauthorized struct {
	Payload *models.Error
}

func (o *GetAPIV10FundsTaskIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/funds/{task_id}][%d] getApiV10FundsTaskIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAPIV10FundsTaskIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV10FundsTaskIDForbidden creates a GetAPIV10FundsTaskIDForbidden with default headers values
func NewGetAPIV10FundsTaskIDForbidden() *GetAPIV10FundsTaskIDForbidden {
	return &GetAPIV10FundsTaskIDForbidden{}
}

/*GetAPIV10FundsTaskIDForbidden handles this case with default header values.

Access Denied
*/
type GetAPIV10FundsTaskIDForbidden struct {
	Payload *models.Error
}

func (o *GetAPIV10FundsTaskIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/funds/{task_id}][%d] getApiV10FundsTaskIdForbidden  %+v", 403, o.Payload)
}

func (o *GetAPIV10FundsTaskIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV10FundsTaskIDNotFound creates a GetAPIV10FundsTaskIDNotFound with default headers values
func NewGetAPIV10FundsTaskIDNotFound() *GetAPIV10FundsTaskIDNotFound {
	return &GetAPIV10FundsTaskIDNotFound{}
}

/*GetAPIV10FundsTaskIDNotFound handles this case with default header values.

The specified resource was not found
*/
type GetAPIV10FundsTaskIDNotFound struct {
	Payload *models.Error
}

func (o *GetAPIV10FundsTaskIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/funds/{task_id}][%d] getApiV10FundsTaskIdNotFound  %+v", 404, o.Payload)
}

func (o *GetAPIV10FundsTaskIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
