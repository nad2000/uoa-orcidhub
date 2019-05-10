// Code generated by go-swagger; DO NOT EDIT.

package orcid_proxy

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
)

// NewPostOrcidAPIVersionOrcidPathParams creates a new PostOrcidAPIVersionOrcidPathParams object
// with the default values initialized.
func NewPostOrcidAPIVersionOrcidPathParams() *PostOrcidAPIVersionOrcidPathParams {
	var ()
	return &PostOrcidAPIVersionOrcidPathParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostOrcidAPIVersionOrcidPathParamsWithTimeout creates a new PostOrcidAPIVersionOrcidPathParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostOrcidAPIVersionOrcidPathParamsWithTimeout(timeout time.Duration) *PostOrcidAPIVersionOrcidPathParams {
	var ()
	return &PostOrcidAPIVersionOrcidPathParams{

		timeout: timeout,
	}
}

// NewPostOrcidAPIVersionOrcidPathParamsWithContext creates a new PostOrcidAPIVersionOrcidPathParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostOrcidAPIVersionOrcidPathParamsWithContext(ctx context.Context) *PostOrcidAPIVersionOrcidPathParams {
	var ()
	return &PostOrcidAPIVersionOrcidPathParams{

		Context: ctx,
	}
}

// NewPostOrcidAPIVersionOrcidPathParamsWithHTTPClient creates a new PostOrcidAPIVersionOrcidPathParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostOrcidAPIVersionOrcidPathParamsWithHTTPClient(client *http.Client) *PostOrcidAPIVersionOrcidPathParams {
	var ()
	return &PostOrcidAPIVersionOrcidPathParams{
		HTTPClient: client,
	}
}

/*PostOrcidAPIVersionOrcidPathParams contains all the parameters to send to the API endpoint
for the post orcid API version orcid path operation typically these are written to a http.Request
*/
type PostOrcidAPIVersionOrcidPathParams struct {

	/*Body*/
	Body interface{}
	/*Orcid
	  User ORCID ID.

	*/
	Orcid string
	/*Path
	  The rest of the ORCID API entry point URL.

	*/
	Path string
	/*Version
	  ORCID API version

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post orcid API version orcid path params
func (o *PostOrcidAPIVersionOrcidPathParams) WithTimeout(timeout time.Duration) *PostOrcidAPIVersionOrcidPathParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post orcid API version orcid path params
func (o *PostOrcidAPIVersionOrcidPathParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post orcid API version orcid path params
func (o *PostOrcidAPIVersionOrcidPathParams) WithContext(ctx context.Context) *PostOrcidAPIVersionOrcidPathParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post orcid API version orcid path params
func (o *PostOrcidAPIVersionOrcidPathParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post orcid API version orcid path params
func (o *PostOrcidAPIVersionOrcidPathParams) WithHTTPClient(client *http.Client) *PostOrcidAPIVersionOrcidPathParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post orcid API version orcid path params
func (o *PostOrcidAPIVersionOrcidPathParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post orcid API version orcid path params
func (o *PostOrcidAPIVersionOrcidPathParams) WithBody(body interface{}) *PostOrcidAPIVersionOrcidPathParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post orcid API version orcid path params
func (o *PostOrcidAPIVersionOrcidPathParams) SetBody(body interface{}) {
	o.Body = body
}

// WithOrcid adds the orcid to the post orcid API version orcid path params
func (o *PostOrcidAPIVersionOrcidPathParams) WithOrcid(orcid string) *PostOrcidAPIVersionOrcidPathParams {
	o.SetOrcid(orcid)
	return o
}

// SetOrcid adds the orcid to the post orcid API version orcid path params
func (o *PostOrcidAPIVersionOrcidPathParams) SetOrcid(orcid string) {
	o.Orcid = orcid
}

// WithPath adds the path to the post orcid API version orcid path params
func (o *PostOrcidAPIVersionOrcidPathParams) WithPath(path string) *PostOrcidAPIVersionOrcidPathParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the post orcid API version orcid path params
func (o *PostOrcidAPIVersionOrcidPathParams) SetPath(path string) {
	o.Path = path
}

// WithVersion adds the version to the post orcid API version orcid path params
func (o *PostOrcidAPIVersionOrcidPathParams) WithVersion(version string) *PostOrcidAPIVersionOrcidPathParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the post orcid API version orcid path params
func (o *PostOrcidAPIVersionOrcidPathParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PostOrcidAPIVersionOrcidPathParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param orcid
	if err := r.SetPathParam("orcid", o.Orcid); err != nil {
		return err
	}

	// path param path
	if err := r.SetPathParam("path", o.Path); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
