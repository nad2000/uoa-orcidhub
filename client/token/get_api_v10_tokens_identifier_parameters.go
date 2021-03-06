// Code generated by go-swagger; DO NOT EDIT.

package token

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

// NewGetAPIV10TokensIdentifierParams creates a new GetAPIV10TokensIdentifierParams object
// with the default values initialized.
func NewGetAPIV10TokensIdentifierParams() *GetAPIV10TokensIdentifierParams {
	var ()
	return &GetAPIV10TokensIdentifierParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV10TokensIdentifierParamsWithTimeout creates a new GetAPIV10TokensIdentifierParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIV10TokensIdentifierParamsWithTimeout(timeout time.Duration) *GetAPIV10TokensIdentifierParams {
	var ()
	return &GetAPIV10TokensIdentifierParams{

		timeout: timeout,
	}
}

// NewGetAPIV10TokensIdentifierParamsWithContext creates a new GetAPIV10TokensIdentifierParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIV10TokensIdentifierParamsWithContext(ctx context.Context) *GetAPIV10TokensIdentifierParams {
	var ()
	return &GetAPIV10TokensIdentifierParams{

		Context: ctx,
	}
}

// NewGetAPIV10TokensIdentifierParamsWithHTTPClient creates a new GetAPIV10TokensIdentifierParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIV10TokensIdentifierParamsWithHTTPClient(client *http.Client) *GetAPIV10TokensIdentifierParams {
	var ()
	return &GetAPIV10TokensIdentifierParams{
		HTTPClient: client,
	}
}

/*GetAPIV10TokensIdentifierParams contains all the parameters to send to the API endpoint
for the get API v10 tokens identifier operation typically these are written to a http.Request
*/
type GetAPIV10TokensIdentifierParams struct {

	/*Identifier
	  User identifier (either email, eppn or ORCID ID)

	*/
	Identifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get API v10 tokens identifier params
func (o *GetAPIV10TokensIdentifierParams) WithTimeout(timeout time.Duration) *GetAPIV10TokensIdentifierParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v10 tokens identifier params
func (o *GetAPIV10TokensIdentifierParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v10 tokens identifier params
func (o *GetAPIV10TokensIdentifierParams) WithContext(ctx context.Context) *GetAPIV10TokensIdentifierParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v10 tokens identifier params
func (o *GetAPIV10TokensIdentifierParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v10 tokens identifier params
func (o *GetAPIV10TokensIdentifierParams) WithHTTPClient(client *http.Client) *GetAPIV10TokensIdentifierParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v10 tokens identifier params
func (o *GetAPIV10TokensIdentifierParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIdentifier adds the identifier to the get API v10 tokens identifier params
func (o *GetAPIV10TokensIdentifierParams) WithIdentifier(identifier string) *GetAPIV10TokensIdentifierParams {
	o.SetIdentifier(identifier)
	return o
}

// SetIdentifier adds the identifier to the get API v10 tokens identifier params
func (o *GetAPIV10TokensIdentifierParams) SetIdentifier(identifier string) {
	o.Identifier = identifier
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV10TokensIdentifierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param identifier
	if err := r.SetPathParam("identifier", o.Identifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
