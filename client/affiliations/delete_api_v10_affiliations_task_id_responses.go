// Code generated by go-swagger; DO NOT EDIT.

package affiliations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/nad2000/uoa-orcidhub/models"
)

// DeleteAPIV10AffiliationsTaskIDReader is a Reader for the DeleteAPIV10AffiliationsTaskID structure.
type DeleteAPIV10AffiliationsTaskIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIV10AffiliationsTaskIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteAPIV10AffiliationsTaskIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteAPIV10AffiliationsTaskIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteAPIV10AffiliationsTaskIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteAPIV10AffiliationsTaskIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAPIV10AffiliationsTaskIDOK creates a DeleteAPIV10AffiliationsTaskIDOK with default headers values
func NewDeleteAPIV10AffiliationsTaskIDOK() *DeleteAPIV10AffiliationsTaskIDOK {
	return &DeleteAPIV10AffiliationsTaskIDOK{}
}

/*DeleteAPIV10AffiliationsTaskIDOK handles this case with default header values.

Successful operation
*/
type DeleteAPIV10AffiliationsTaskIDOK struct {
}

func (o *DeleteAPIV10AffiliationsTaskIDOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1.0/affiliations/{task_id}][%d] deleteApiV10AffiliationsTaskIdOK ", 200)
}

func (o *DeleteAPIV10AffiliationsTaskIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAPIV10AffiliationsTaskIDUnauthorized creates a DeleteAPIV10AffiliationsTaskIDUnauthorized with default headers values
func NewDeleteAPIV10AffiliationsTaskIDUnauthorized() *DeleteAPIV10AffiliationsTaskIDUnauthorized {
	return &DeleteAPIV10AffiliationsTaskIDUnauthorized{}
}

/*DeleteAPIV10AffiliationsTaskIDUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteAPIV10AffiliationsTaskIDUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteAPIV10AffiliationsTaskIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1.0/affiliations/{task_id}][%d] deleteApiV10AffiliationsTaskIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteAPIV10AffiliationsTaskIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIV10AffiliationsTaskIDForbidden creates a DeleteAPIV10AffiliationsTaskIDForbidden with default headers values
func NewDeleteAPIV10AffiliationsTaskIDForbidden() *DeleteAPIV10AffiliationsTaskIDForbidden {
	return &DeleteAPIV10AffiliationsTaskIDForbidden{}
}

/*DeleteAPIV10AffiliationsTaskIDForbidden handles this case with default header values.

Access Denied
*/
type DeleteAPIV10AffiliationsTaskIDForbidden struct {
	Payload *models.Error
}

func (o *DeleteAPIV10AffiliationsTaskIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1.0/affiliations/{task_id}][%d] deleteApiV10AffiliationsTaskIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteAPIV10AffiliationsTaskIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIV10AffiliationsTaskIDNotFound creates a DeleteAPIV10AffiliationsTaskIDNotFound with default headers values
func NewDeleteAPIV10AffiliationsTaskIDNotFound() *DeleteAPIV10AffiliationsTaskIDNotFound {
	return &DeleteAPIV10AffiliationsTaskIDNotFound{}
}

/*DeleteAPIV10AffiliationsTaskIDNotFound handles this case with default header values.

The specified resource was not found
*/
type DeleteAPIV10AffiliationsTaskIDNotFound struct {
	Payload *models.Error
}

func (o *DeleteAPIV10AffiliationsTaskIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1.0/affiliations/{task_id}][%d] deleteApiV10AffiliationsTaskIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAPIV10AffiliationsTaskIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}