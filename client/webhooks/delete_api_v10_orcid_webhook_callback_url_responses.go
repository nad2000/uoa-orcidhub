// Code generated by go-swagger; DO NOT EDIT.

package webhooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/nad2000/uoa-orcidhub/models"
)

// DeleteAPIV10OrcidWebhookCallbackURLReader is a Reader for the DeleteAPIV10OrcidWebhookCallbackURL structure.
type DeleteAPIV10OrcidWebhookCallbackURLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIV10OrcidWebhookCallbackURLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteAPIV10OrcidWebhookCallbackURLNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteAPIV10OrcidWebhookCallbackURLNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewDeleteAPIV10OrcidWebhookCallbackURLUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAPIV10OrcidWebhookCallbackURLNoContent creates a DeleteAPIV10OrcidWebhookCallbackURLNoContent with default headers values
func NewDeleteAPIV10OrcidWebhookCallbackURLNoContent() *DeleteAPIV10OrcidWebhookCallbackURLNoContent {
	return &DeleteAPIV10OrcidWebhookCallbackURLNoContent{}
}

/*DeleteAPIV10OrcidWebhookCallbackURLNoContent handles this case with default header values.

A webhoook successfully unregistered.
*/
type DeleteAPIV10OrcidWebhookCallbackURLNoContent struct {
}

func (o *DeleteAPIV10OrcidWebhookCallbackURLNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1.0/{orcid}/webhook/{callback_url}][%d] deleteApiV10OrcidWebhookCallbackUrlNoContent ", 204)
}

func (o *DeleteAPIV10OrcidWebhookCallbackURLNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAPIV10OrcidWebhookCallbackURLNotFound creates a DeleteAPIV10OrcidWebhookCallbackURLNotFound with default headers values
func NewDeleteAPIV10OrcidWebhookCallbackURLNotFound() *DeleteAPIV10OrcidWebhookCallbackURLNotFound {
	return &DeleteAPIV10OrcidWebhookCallbackURLNotFound{}
}

/*DeleteAPIV10OrcidWebhookCallbackURLNotFound handles this case with default header values.

Invalid ORCID iD.
*/
type DeleteAPIV10OrcidWebhookCallbackURLNotFound struct {
	Payload *models.Error
}

func (o *DeleteAPIV10OrcidWebhookCallbackURLNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1.0/{orcid}/webhook/{callback_url}][%d] deleteApiV10OrcidWebhookCallbackUrlNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAPIV10OrcidWebhookCallbackURLNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIV10OrcidWebhookCallbackURLUnsupportedMediaType creates a DeleteAPIV10OrcidWebhookCallbackURLUnsupportedMediaType with default headers values
func NewDeleteAPIV10OrcidWebhookCallbackURLUnsupportedMediaType() *DeleteAPIV10OrcidWebhookCallbackURLUnsupportedMediaType {
	return &DeleteAPIV10OrcidWebhookCallbackURLUnsupportedMediaType{}
}

/*DeleteAPIV10OrcidWebhookCallbackURLUnsupportedMediaType handles this case with default header values.

Invalid call-back URL or missing ORCID iD.
*/
type DeleteAPIV10OrcidWebhookCallbackURLUnsupportedMediaType struct {
	Payload *models.Error
}

func (o *DeleteAPIV10OrcidWebhookCallbackURLUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v1.0/{orcid}/webhook/{callback_url}][%d] deleteApiV10OrcidWebhookCallbackUrlUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteAPIV10OrcidWebhookCallbackURLUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
