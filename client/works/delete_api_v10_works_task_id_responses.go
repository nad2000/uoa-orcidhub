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

// DeleteAPIV10WorksTaskIDReader is a Reader for the DeleteAPIV10WorksTaskID structure.
type DeleteAPIV10WorksTaskIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIV10WorksTaskIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteAPIV10WorksTaskIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteAPIV10WorksTaskIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteAPIV10WorksTaskIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteAPIV10WorksTaskIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAPIV10WorksTaskIDOK creates a DeleteAPIV10WorksTaskIDOK with default headers values
func NewDeleteAPIV10WorksTaskIDOK() *DeleteAPIV10WorksTaskIDOK {
	return &DeleteAPIV10WorksTaskIDOK{}
}

/*DeleteAPIV10WorksTaskIDOK handles this case with default header values.

Successful operation
*/
type DeleteAPIV10WorksTaskIDOK struct {
}

func (o *DeleteAPIV10WorksTaskIDOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1.0/works/{task_id}][%d] deleteApiV10WorksTaskIdOK ", 200)
}

func (o *DeleteAPIV10WorksTaskIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAPIV10WorksTaskIDUnauthorized creates a DeleteAPIV10WorksTaskIDUnauthorized with default headers values
func NewDeleteAPIV10WorksTaskIDUnauthorized() *DeleteAPIV10WorksTaskIDUnauthorized {
	return &DeleteAPIV10WorksTaskIDUnauthorized{}
}

/*DeleteAPIV10WorksTaskIDUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteAPIV10WorksTaskIDUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteAPIV10WorksTaskIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1.0/works/{task_id}][%d] deleteApiV10WorksTaskIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteAPIV10WorksTaskIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIV10WorksTaskIDForbidden creates a DeleteAPIV10WorksTaskIDForbidden with default headers values
func NewDeleteAPIV10WorksTaskIDForbidden() *DeleteAPIV10WorksTaskIDForbidden {
	return &DeleteAPIV10WorksTaskIDForbidden{}
}

/*DeleteAPIV10WorksTaskIDForbidden handles this case with default header values.

Access Denied
*/
type DeleteAPIV10WorksTaskIDForbidden struct {
	Payload *models.Error
}

func (o *DeleteAPIV10WorksTaskIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1.0/works/{task_id}][%d] deleteApiV10WorksTaskIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteAPIV10WorksTaskIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIV10WorksTaskIDNotFound creates a DeleteAPIV10WorksTaskIDNotFound with default headers values
func NewDeleteAPIV10WorksTaskIDNotFound() *DeleteAPIV10WorksTaskIDNotFound {
	return &DeleteAPIV10WorksTaskIDNotFound{}
}

/*DeleteAPIV10WorksTaskIDNotFound handles this case with default header values.

The specified resource was not found
*/
type DeleteAPIV10WorksTaskIDNotFound struct {
	Payload *models.Error
}

func (o *DeleteAPIV10WorksTaskIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1.0/works/{task_id}][%d] deleteApiV10WorksTaskIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAPIV10WorksTaskIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
