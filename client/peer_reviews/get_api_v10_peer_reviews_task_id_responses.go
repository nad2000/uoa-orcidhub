// Code generated by go-swagger; DO NOT EDIT.

package peer_reviews

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/nad2000/uoa-orcidhub/models"
)

// GetAPIV10PeerReviewsTaskIDReader is a Reader for the GetAPIV10PeerReviewsTaskID structure.
type GetAPIV10PeerReviewsTaskIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV10PeerReviewsTaskIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAPIV10PeerReviewsTaskIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetAPIV10PeerReviewsTaskIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAPIV10PeerReviewsTaskIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAPIV10PeerReviewsTaskIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAPIV10PeerReviewsTaskIDOK creates a GetAPIV10PeerReviewsTaskIDOK with default headers values
func NewGetAPIV10PeerReviewsTaskIDOK() *GetAPIV10PeerReviewsTaskIDOK {
	return &GetAPIV10PeerReviewsTaskIDOK{}
}

/*GetAPIV10PeerReviewsTaskIDOK handles this case with default header values.

successful operation
*/
type GetAPIV10PeerReviewsTaskIDOK struct {
	Payload *models.PeerReviewTask
}

func (o *GetAPIV10PeerReviewsTaskIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/peer-reviews/{task_id}][%d] getApiV10PeerReviewsTaskIdOK  %+v", 200, o.Payload)
}

func (o *GetAPIV10PeerReviewsTaskIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PeerReviewTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV10PeerReviewsTaskIDUnauthorized creates a GetAPIV10PeerReviewsTaskIDUnauthorized with default headers values
func NewGetAPIV10PeerReviewsTaskIDUnauthorized() *GetAPIV10PeerReviewsTaskIDUnauthorized {
	return &GetAPIV10PeerReviewsTaskIDUnauthorized{}
}

/*GetAPIV10PeerReviewsTaskIDUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAPIV10PeerReviewsTaskIDUnauthorized struct {
	Payload *models.Error
}

func (o *GetAPIV10PeerReviewsTaskIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/peer-reviews/{task_id}][%d] getApiV10PeerReviewsTaskIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAPIV10PeerReviewsTaskIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV10PeerReviewsTaskIDForbidden creates a GetAPIV10PeerReviewsTaskIDForbidden with default headers values
func NewGetAPIV10PeerReviewsTaskIDForbidden() *GetAPIV10PeerReviewsTaskIDForbidden {
	return &GetAPIV10PeerReviewsTaskIDForbidden{}
}

/*GetAPIV10PeerReviewsTaskIDForbidden handles this case with default header values.

Access Denied
*/
type GetAPIV10PeerReviewsTaskIDForbidden struct {
	Payload *models.Error
}

func (o *GetAPIV10PeerReviewsTaskIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/peer-reviews/{task_id}][%d] getApiV10PeerReviewsTaskIdForbidden  %+v", 403, o.Payload)
}

func (o *GetAPIV10PeerReviewsTaskIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV10PeerReviewsTaskIDNotFound creates a GetAPIV10PeerReviewsTaskIDNotFound with default headers values
func NewGetAPIV10PeerReviewsTaskIDNotFound() *GetAPIV10PeerReviewsTaskIDNotFound {
	return &GetAPIV10PeerReviewsTaskIDNotFound{}
}

/*GetAPIV10PeerReviewsTaskIDNotFound handles this case with default header values.

The specified resource was not found
*/
type GetAPIV10PeerReviewsTaskIDNotFound struct {
	Payload *models.Error
}

func (o *GetAPIV10PeerReviewsTaskIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/peer-reviews/{task_id}][%d] getApiV10PeerReviewsTaskIdNotFound  %+v", 404, o.Payload)
}

func (o *GetAPIV10PeerReviewsTaskIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
