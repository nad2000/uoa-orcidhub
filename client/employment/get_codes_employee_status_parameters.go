// Code generated by go-swagger; DO NOT EDIT.

package employment

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

// NewGetCodesEmployeeStatusParams creates a new GetCodesEmployeeStatusParams object
// with the default values initialized.
func NewGetCodesEmployeeStatusParams() *GetCodesEmployeeStatusParams {

	return &GetCodesEmployeeStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCodesEmployeeStatusParamsWithTimeout creates a new GetCodesEmployeeStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCodesEmployeeStatusParamsWithTimeout(timeout time.Duration) *GetCodesEmployeeStatusParams {

	return &GetCodesEmployeeStatusParams{

		timeout: timeout,
	}
}

// NewGetCodesEmployeeStatusParamsWithContext creates a new GetCodesEmployeeStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCodesEmployeeStatusParamsWithContext(ctx context.Context) *GetCodesEmployeeStatusParams {

	return &GetCodesEmployeeStatusParams{

		Context: ctx,
	}
}

// NewGetCodesEmployeeStatusParamsWithHTTPClient creates a new GetCodesEmployeeStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCodesEmployeeStatusParamsWithHTTPClient(client *http.Client) *GetCodesEmployeeStatusParams {

	return &GetCodesEmployeeStatusParams{
		HTTPClient: client,
	}
}

/*GetCodesEmployeeStatusParams contains all the parameters to send to the API endpoint
for the get codes employee status operation typically these are written to a http.Request
*/
type GetCodesEmployeeStatusParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get codes employee status params
func (o *GetCodesEmployeeStatusParams) WithTimeout(timeout time.Duration) *GetCodesEmployeeStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get codes employee status params
func (o *GetCodesEmployeeStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get codes employee status params
func (o *GetCodesEmployeeStatusParams) WithContext(ctx context.Context) *GetCodesEmployeeStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get codes employee status params
func (o *GetCodesEmployeeStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get codes employee status params
func (o *GetCodesEmployeeStatusParams) WithHTTPClient(client *http.Client) *GetCodesEmployeeStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get codes employee status params
func (o *GetCodesEmployeeStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetCodesEmployeeStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}