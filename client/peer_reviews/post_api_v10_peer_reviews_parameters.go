// Code generated by go-swagger; DO NOT EDIT.

package peer_reviews

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/nad2000/uoa-orcidhub/models"
)

// NewPostAPIV10PeerReviewsParams creates a new PostAPIV10PeerReviewsParams object
// with the default values initialized.
func NewPostAPIV10PeerReviewsParams() *PostAPIV10PeerReviewsParams {
	var ()
	return &PostAPIV10PeerReviewsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV10PeerReviewsParamsWithTimeout creates a new PostAPIV10PeerReviewsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPIV10PeerReviewsParamsWithTimeout(timeout time.Duration) *PostAPIV10PeerReviewsParams {
	var ()
	return &PostAPIV10PeerReviewsParams{

		timeout: timeout,
	}
}

// NewPostAPIV10PeerReviewsParamsWithContext creates a new PostAPIV10PeerReviewsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPIV10PeerReviewsParamsWithContext(ctx context.Context) *PostAPIV10PeerReviewsParams {
	var ()
	return &PostAPIV10PeerReviewsParams{

		Context: ctx,
	}
}

// NewPostAPIV10PeerReviewsParamsWithHTTPClient creates a new PostAPIV10PeerReviewsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPIV10PeerReviewsParamsWithHTTPClient(client *http.Client) *PostAPIV10PeerReviewsParams {
	var ()
	return &PostAPIV10PeerReviewsParams{
		HTTPClient: client,
	}
}

/*PostAPIV10PeerReviewsParams contains all the parameters to send to the API endpoint
for the post API v10 peer reviews operation typically these are written to a http.Request
*/
type PostAPIV10PeerReviewsParams struct {

	/*Body
	  PeerReview task.

	*/
	Body *models.PeerReviewTask
	/*Filename
	  The batch process filename.

	*/
	Filename *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API v10 peer reviews params
func (o *PostAPIV10PeerReviewsParams) WithTimeout(timeout time.Duration) *PostAPIV10PeerReviewsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v10 peer reviews params
func (o *PostAPIV10PeerReviewsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v10 peer reviews params
func (o *PostAPIV10PeerReviewsParams) WithContext(ctx context.Context) *PostAPIV10PeerReviewsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v10 peer reviews params
func (o *PostAPIV10PeerReviewsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v10 peer reviews params
func (o *PostAPIV10PeerReviewsParams) WithHTTPClient(client *http.Client) *PostAPIV10PeerReviewsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v10 peer reviews params
func (o *PostAPIV10PeerReviewsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post API v10 peer reviews params
func (o *PostAPIV10PeerReviewsParams) WithBody(body *models.PeerReviewTask) *PostAPIV10PeerReviewsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post API v10 peer reviews params
func (o *PostAPIV10PeerReviewsParams) SetBody(body *models.PeerReviewTask) {
	o.Body = body
}

// WithFilename adds the filename to the post API v10 peer reviews params
func (o *PostAPIV10PeerReviewsParams) WithFilename(filename *string) *PostAPIV10PeerReviewsParams {
	o.SetFilename(filename)
	return o
}

// SetFilename adds the filename to the post API v10 peer reviews params
func (o *PostAPIV10PeerReviewsParams) SetFilename(filename *string) {
	o.Filename = filename
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV10PeerReviewsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Filename != nil {

		// query param filename
		var qrFilename string
		if o.Filename != nil {
			qrFilename = *o.Filename
		}
		qFilename := qrFilename
		if qFilename != "" {
			if err := r.SetQueryParam("filename", qFilename); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
