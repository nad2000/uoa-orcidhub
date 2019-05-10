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

// PostAPIV10WorksReader is a Reader for the PostAPIV10Works structure.
type PostAPIV10WorksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIV10WorksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostAPIV10WorksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostAPIV10WorksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostAPIV10WorksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostAPIV10WorksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostAPIV10WorksOK creates a PostAPIV10WorksOK with default headers values
func NewPostAPIV10WorksOK() *PostAPIV10WorksOK {
	return &PostAPIV10WorksOK{}
}

/*PostAPIV10WorksOK handles this case with default header values.

successful operation
*/
type PostAPIV10WorksOK struct {
	Payload *models.WorkTask
}

func (o *PostAPIV10WorksOK) Error() string {
	return fmt.Sprintf("[POST /api/v1.0/works][%d] postApiV10WorksOK  %+v", 200, o.Payload)
}

func (o *PostAPIV10WorksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV10WorksUnauthorized creates a PostAPIV10WorksUnauthorized with default headers values
func NewPostAPIV10WorksUnauthorized() *PostAPIV10WorksUnauthorized {
	return &PostAPIV10WorksUnauthorized{}
}

/*PostAPIV10WorksUnauthorized handles this case with default header values.

Unauthorized
*/
type PostAPIV10WorksUnauthorized struct {
	Payload *models.Error
}

func (o *PostAPIV10WorksUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1.0/works][%d] postApiV10WorksUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAPIV10WorksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV10WorksForbidden creates a PostAPIV10WorksForbidden with default headers values
func NewPostAPIV10WorksForbidden() *PostAPIV10WorksForbidden {
	return &PostAPIV10WorksForbidden{}
}

/*PostAPIV10WorksForbidden handles this case with default header values.

Access Denied
*/
type PostAPIV10WorksForbidden struct {
	Payload *models.Error
}

func (o *PostAPIV10WorksForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1.0/works][%d] postApiV10WorksForbidden  %+v", 403, o.Payload)
}

func (o *PostAPIV10WorksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV10WorksNotFound creates a PostAPIV10WorksNotFound with default headers values
func NewPostAPIV10WorksNotFound() *PostAPIV10WorksNotFound {
	return &PostAPIV10WorksNotFound{}
}

/*PostAPIV10WorksNotFound handles this case with default header values.

The specified resource was not found
*/
type PostAPIV10WorksNotFound struct {
	Payload *models.Error
}

func (o *PostAPIV10WorksNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1.0/works][%d] postApiV10WorksNotFound  %+v", 404, o.Payload)
}

func (o *PostAPIV10WorksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
