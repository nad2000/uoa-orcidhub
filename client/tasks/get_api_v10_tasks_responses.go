// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/nad2000/uoa-orcidhub/models"
)

// GetAPIV10TasksReader is a Reader for the GetAPIV10Tasks structure.
type GetAPIV10TasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV10TasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAPIV10TasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetAPIV10TasksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAPIV10TasksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAPIV10TasksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAPIV10TasksOK creates a GetAPIV10TasksOK with default headers values
func NewGetAPIV10TasksOK() *GetAPIV10TasksOK {
	return &GetAPIV10TasksOK{}
}

/*GetAPIV10TasksOK handles this case with default header values.

successful operation
*/
type GetAPIV10TasksOK struct {
	Payload []*models.Task
}

func (o *GetAPIV10TasksOK) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/tasks][%d] getApiV10TasksOK  %+v", 200, o.Payload)
}

func (o *GetAPIV10TasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV10TasksUnauthorized creates a GetAPIV10TasksUnauthorized with default headers values
func NewGetAPIV10TasksUnauthorized() *GetAPIV10TasksUnauthorized {
	return &GetAPIV10TasksUnauthorized{}
}

/*GetAPIV10TasksUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAPIV10TasksUnauthorized struct {
	Payload *models.Error
}

func (o *GetAPIV10TasksUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/tasks][%d] getApiV10TasksUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAPIV10TasksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV10TasksForbidden creates a GetAPIV10TasksForbidden with default headers values
func NewGetAPIV10TasksForbidden() *GetAPIV10TasksForbidden {
	return &GetAPIV10TasksForbidden{}
}

/*GetAPIV10TasksForbidden handles this case with default header values.

Access Denied
*/
type GetAPIV10TasksForbidden struct {
	Payload *models.Error
}

func (o *GetAPIV10TasksForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/tasks][%d] getApiV10TasksForbidden  %+v", 403, o.Payload)
}

func (o *GetAPIV10TasksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV10TasksNotFound creates a GetAPIV10TasksNotFound with default headers values
func NewGetAPIV10TasksNotFound() *GetAPIV10TasksNotFound {
	return &GetAPIV10TasksNotFound{}
}

/*GetAPIV10TasksNotFound handles this case with default header values.

The specified resource was not found
*/
type GetAPIV10TasksNotFound struct {
	Payload *models.Error
}

func (o *GetAPIV10TasksNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/tasks][%d] getApiV10TasksNotFound  %+v", 404, o.Payload)
}

func (o *GetAPIV10TasksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
