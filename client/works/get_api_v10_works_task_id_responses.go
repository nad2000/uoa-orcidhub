// Code generated by go-swagger; DO NOT EDIT.

package works

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/nad2000/uoa-orcidhub/models"
)

// GetAPIV10WorksTaskIDReader is a Reader for the GetAPIV10WorksTaskID structure.
type GetAPIV10WorksTaskIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV10WorksTaskIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAPIV10WorksTaskIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetAPIV10WorksTaskIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAPIV10WorksTaskIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAPIV10WorksTaskIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAPIV10WorksTaskIDOK creates a GetAPIV10WorksTaskIDOK with default headers values
func NewGetAPIV10WorksTaskIDOK() *GetAPIV10WorksTaskIDOK {
	return &GetAPIV10WorksTaskIDOK{}
}

/*GetAPIV10WorksTaskIDOK handles this case with default header values.

successful operation
*/
type GetAPIV10WorksTaskIDOK struct {
	Payload *models.WorkTask
}

func (o *GetAPIV10WorksTaskIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/works/{task_id}][%d] getApiV10WorksTaskIdOK  %+v", 200, o.Payload)
}

func (o *GetAPIV10WorksTaskIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV10WorksTaskIDUnauthorized creates a GetAPIV10WorksTaskIDUnauthorized with default headers values
func NewGetAPIV10WorksTaskIDUnauthorized() *GetAPIV10WorksTaskIDUnauthorized {
	return &GetAPIV10WorksTaskIDUnauthorized{}
}

/*GetAPIV10WorksTaskIDUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAPIV10WorksTaskIDUnauthorized struct {
	Payload *models.Error
}

func (o *GetAPIV10WorksTaskIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/works/{task_id}][%d] getApiV10WorksTaskIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAPIV10WorksTaskIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV10WorksTaskIDForbidden creates a GetAPIV10WorksTaskIDForbidden with default headers values
func NewGetAPIV10WorksTaskIDForbidden() *GetAPIV10WorksTaskIDForbidden {
	return &GetAPIV10WorksTaskIDForbidden{}
}

/*GetAPIV10WorksTaskIDForbidden handles this case with default header values.

Access Denied
*/
type GetAPIV10WorksTaskIDForbidden struct {
	Payload *models.Error
}

func (o *GetAPIV10WorksTaskIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/works/{task_id}][%d] getApiV10WorksTaskIdForbidden  %+v", 403, o.Payload)
}

func (o *GetAPIV10WorksTaskIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV10WorksTaskIDNotFound creates a GetAPIV10WorksTaskIDNotFound with default headers values
func NewGetAPIV10WorksTaskIDNotFound() *GetAPIV10WorksTaskIDNotFound {
	return &GetAPIV10WorksTaskIDNotFound{}
}

/*GetAPIV10WorksTaskIDNotFound handles this case with default header values.

The specified resource was not found
*/
type GetAPIV10WorksTaskIDNotFound struct {
	Payload *models.Error
}

func (o *GetAPIV10WorksTaskIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/works/{task_id}][%d] getApiV10WorksTaskIdNotFound  %+v", 404, o.Payload)
}

func (o *GetAPIV10WorksTaskIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
