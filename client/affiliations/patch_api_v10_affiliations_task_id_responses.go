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

// PatchAPIV10AffiliationsTaskIDReader is a Reader for the PatchAPIV10AffiliationsTaskID structure.
type PatchAPIV10AffiliationsTaskIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPIV10AffiliationsTaskIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchAPIV10AffiliationsTaskIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPatchAPIV10AffiliationsTaskIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPatchAPIV10AffiliationsTaskIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchAPIV10AffiliationsTaskIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchAPIV10AffiliationsTaskIDOK creates a PatchAPIV10AffiliationsTaskIDOK with default headers values
func NewPatchAPIV10AffiliationsTaskIDOK() *PatchAPIV10AffiliationsTaskIDOK {
	return &PatchAPIV10AffiliationsTaskIDOK{}
}

/*PatchAPIV10AffiliationsTaskIDOK handles this case with default header values.

successful operation
*/
type PatchAPIV10AffiliationsTaskIDOK struct {
	Payload *models.AffiliationTask
}

func (o *PatchAPIV10AffiliationsTaskIDOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1.0/affiliations/{task_id}][%d] patchApiV10AffiliationsTaskIdOK  %+v", 200, o.Payload)
}

func (o *PatchAPIV10AffiliationsTaskIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AffiliationTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIV10AffiliationsTaskIDUnauthorized creates a PatchAPIV10AffiliationsTaskIDUnauthorized with default headers values
func NewPatchAPIV10AffiliationsTaskIDUnauthorized() *PatchAPIV10AffiliationsTaskIDUnauthorized {
	return &PatchAPIV10AffiliationsTaskIDUnauthorized{}
}

/*PatchAPIV10AffiliationsTaskIDUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchAPIV10AffiliationsTaskIDUnauthorized struct {
	Payload *models.Error
}

func (o *PatchAPIV10AffiliationsTaskIDUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v1.0/affiliations/{task_id}][%d] patchApiV10AffiliationsTaskIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchAPIV10AffiliationsTaskIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIV10AffiliationsTaskIDForbidden creates a PatchAPIV10AffiliationsTaskIDForbidden with default headers values
func NewPatchAPIV10AffiliationsTaskIDForbidden() *PatchAPIV10AffiliationsTaskIDForbidden {
	return &PatchAPIV10AffiliationsTaskIDForbidden{}
}

/*PatchAPIV10AffiliationsTaskIDForbidden handles this case with default header values.

Access Denied
*/
type PatchAPIV10AffiliationsTaskIDForbidden struct {
	Payload *models.Error
}

func (o *PatchAPIV10AffiliationsTaskIDForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1.0/affiliations/{task_id}][%d] patchApiV10AffiliationsTaskIdForbidden  %+v", 403, o.Payload)
}

func (o *PatchAPIV10AffiliationsTaskIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIV10AffiliationsTaskIDNotFound creates a PatchAPIV10AffiliationsTaskIDNotFound with default headers values
func NewPatchAPIV10AffiliationsTaskIDNotFound() *PatchAPIV10AffiliationsTaskIDNotFound {
	return &PatchAPIV10AffiliationsTaskIDNotFound{}
}

/*PatchAPIV10AffiliationsTaskIDNotFound handles this case with default header values.

The specified resource was not found
*/
type PatchAPIV10AffiliationsTaskIDNotFound struct {
	Payload *models.Error
}

func (o *PatchAPIV10AffiliationsTaskIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v1.0/affiliations/{task_id}][%d] patchApiV10AffiliationsTaskIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchAPIV10AffiliationsTaskIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
