// Code generated by go-swagger; DO NOT EDIT.

package orcid_proxy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/nad2000/uoa-orcidhub/models"
)

// GetOrcidAPIVersionOrcidReader is a Reader for the GetOrcidAPIVersionOrcid structure.
type GetOrcidAPIVersionOrcidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrcidAPIVersionOrcidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOrcidAPIVersionOrcidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetOrcidAPIVersionOrcidForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetOrcidAPIVersionOrcidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewGetOrcidAPIVersionOrcidUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrcidAPIVersionOrcidOK creates a GetOrcidAPIVersionOrcidOK with default headers values
func NewGetOrcidAPIVersionOrcidOK() *GetOrcidAPIVersionOrcidOK {
	return &GetOrcidAPIVersionOrcidOK{}
}

/*GetOrcidAPIVersionOrcidOK handles this case with default header values.

Successful operation
*/
type GetOrcidAPIVersionOrcidOK struct {
	Payload interface{}
}

func (o *GetOrcidAPIVersionOrcidOK) Error() string {
	return fmt.Sprintf("[GET /orcid/api/{version}/{orcid}][%d] getOrcidApiVersionOrcidOK  %+v", 200, o.Payload)
}

func (o *GetOrcidAPIVersionOrcidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrcidAPIVersionOrcidForbidden creates a GetOrcidAPIVersionOrcidForbidden with default headers values
func NewGetOrcidAPIVersionOrcidForbidden() *GetOrcidAPIVersionOrcidForbidden {
	return &GetOrcidAPIVersionOrcidForbidden{}
}

/*GetOrcidAPIVersionOrcidForbidden handles this case with default header values.

The user hasn't granted acceess to the profile.
*/
type GetOrcidAPIVersionOrcidForbidden struct {
	Payload *models.Error
}

func (o *GetOrcidAPIVersionOrcidForbidden) Error() string {
	return fmt.Sprintf("[GET /orcid/api/{version}/{orcid}][%d] getOrcidApiVersionOrcidForbidden  %+v", 403, o.Payload)
}

func (o *GetOrcidAPIVersionOrcidForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrcidAPIVersionOrcidNotFound creates a GetOrcidAPIVersionOrcidNotFound with default headers values
func NewGetOrcidAPIVersionOrcidNotFound() *GetOrcidAPIVersionOrcidNotFound {
	return &GetOrcidAPIVersionOrcidNotFound{}
}

/*GetOrcidAPIVersionOrcidNotFound handles this case with default header values.

Resource not found
*/
type GetOrcidAPIVersionOrcidNotFound struct {
	Payload *models.Error
}

func (o *GetOrcidAPIVersionOrcidNotFound) Error() string {
	return fmt.Sprintf("[GET /orcid/api/{version}/{orcid}][%d] getOrcidApiVersionOrcidNotFound  %+v", 404, o.Payload)
}

func (o *GetOrcidAPIVersionOrcidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrcidAPIVersionOrcidUnsupportedMediaType creates a GetOrcidAPIVersionOrcidUnsupportedMediaType with default headers values
func NewGetOrcidAPIVersionOrcidUnsupportedMediaType() *GetOrcidAPIVersionOrcidUnsupportedMediaType {
	return &GetOrcidAPIVersionOrcidUnsupportedMediaType{}
}

/*GetOrcidAPIVersionOrcidUnsupportedMediaType handles this case with default header values.

Missing or invalid ORCID iD.
*/
type GetOrcidAPIVersionOrcidUnsupportedMediaType struct {
	Payload *models.Error
}

func (o *GetOrcidAPIVersionOrcidUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /orcid/api/{version}/{orcid}][%d] getOrcidApiVersionOrcidUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOrcidAPIVersionOrcidUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
