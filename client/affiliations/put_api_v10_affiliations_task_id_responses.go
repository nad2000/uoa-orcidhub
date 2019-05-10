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

// PutAPIV10AffiliationsTaskIDReader is a Reader for the PutAPIV10AffiliationsTaskID structure.
type PutAPIV10AffiliationsTaskIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAPIV10AffiliationsTaskIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutAPIV10AffiliationsTaskIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPutAPIV10AffiliationsTaskIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPutAPIV10AffiliationsTaskIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutAPIV10AffiliationsTaskIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutAPIV10AffiliationsTaskIDOK creates a PutAPIV10AffiliationsTaskIDOK with default headers values
func NewPutAPIV10AffiliationsTaskIDOK() *PutAPIV10AffiliationsTaskIDOK {
	return &PutAPIV10AffiliationsTaskIDOK{}
}

/*PutAPIV10AffiliationsTaskIDOK handles this case with default header values.

successful operation
*/
type PutAPIV10AffiliationsTaskIDOK struct {
	Payload *models.AffiliationTask
}

func (o *PutAPIV10AffiliationsTaskIDOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1.0/affiliations/{task_id}][%d] putApiV10AffiliationsTaskIdOK  %+v", 200, o.Payload)
}

func (o *PutAPIV10AffiliationsTaskIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AffiliationTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIV10AffiliationsTaskIDUnauthorized creates a PutAPIV10AffiliationsTaskIDUnauthorized with default headers values
func NewPutAPIV10AffiliationsTaskIDUnauthorized() *PutAPIV10AffiliationsTaskIDUnauthorized {
	return &PutAPIV10AffiliationsTaskIDUnauthorized{}
}

/*PutAPIV10AffiliationsTaskIDUnauthorized handles this case with default header values.

Unauthorized
*/
type PutAPIV10AffiliationsTaskIDUnauthorized struct {
	Payload *models.Error
}

func (o *PutAPIV10AffiliationsTaskIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1.0/affiliations/{task_id}][%d] putApiV10AffiliationsTaskIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PutAPIV10AffiliationsTaskIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIV10AffiliationsTaskIDForbidden creates a PutAPIV10AffiliationsTaskIDForbidden with default headers values
func NewPutAPIV10AffiliationsTaskIDForbidden() *PutAPIV10AffiliationsTaskIDForbidden {
	return &PutAPIV10AffiliationsTaskIDForbidden{}
}

/*PutAPIV10AffiliationsTaskIDForbidden handles this case with default header values.

Access Denied
*/
type PutAPIV10AffiliationsTaskIDForbidden struct {
	Payload *models.Error
}

func (o *PutAPIV10AffiliationsTaskIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1.0/affiliations/{task_id}][%d] putApiV10AffiliationsTaskIdForbidden  %+v", 403, o.Payload)
}

func (o *PutAPIV10AffiliationsTaskIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIV10AffiliationsTaskIDNotFound creates a PutAPIV10AffiliationsTaskIDNotFound with default headers values
func NewPutAPIV10AffiliationsTaskIDNotFound() *PutAPIV10AffiliationsTaskIDNotFound {
	return &PutAPIV10AffiliationsTaskIDNotFound{}
}

/*PutAPIV10AffiliationsTaskIDNotFound handles this case with default header values.

The specified resource was not found
*/
type PutAPIV10AffiliationsTaskIDNotFound struct {
	Payload *models.Error
}

func (o *PutAPIV10AffiliationsTaskIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1.0/affiliations/{task_id}][%d] putApiV10AffiliationsTaskIdNotFound  %+v", 404, o.Payload)
}

func (o *PutAPIV10AffiliationsTaskIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
