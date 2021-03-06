// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAPIV10TasksParams creates a new GetAPIV10TasksParams object
// with the default values initialized.
func NewGetAPIV10TasksParams() *GetAPIV10TasksParams {
	var (
		pageDefault     = int64(1)
		pageSizeDefault = int64(20)
	)
	return &GetAPIV10TasksParams{
		Page:     &pageDefault,
		PageSize: &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV10TasksParamsWithTimeout creates a new GetAPIV10TasksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIV10TasksParamsWithTimeout(timeout time.Duration) *GetAPIV10TasksParams {
	var (
		pageDefault     = int64(1)
		pageSizeDefault = int64(20)
	)
	return &GetAPIV10TasksParams{
		Page:     &pageDefault,
		PageSize: &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetAPIV10TasksParamsWithContext creates a new GetAPIV10TasksParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIV10TasksParamsWithContext(ctx context.Context) *GetAPIV10TasksParams {
	var (
		pageDefault     = int64(1)
		pageSizeDefault = int64(20)
	)
	return &GetAPIV10TasksParams{
		Page:     &pageDefault,
		PageSize: &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetAPIV10TasksParamsWithHTTPClient creates a new GetAPIV10TasksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIV10TasksParamsWithHTTPClient(client *http.Client) *GetAPIV10TasksParams {
	var (
		pageDefault     = int64(1)
		pageSizeDefault = int64(20)
	)
	return &GetAPIV10TasksParams{
		Page:       &pageDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetAPIV10TasksParams contains all the parameters to send to the API endpoint
for the get API v10 tasks operation typically these are written to a http.Request
*/
type GetAPIV10TasksParams struct {

	/*Page
	  The number of the page of retrieved data starting counting from 1

	*/
	Page *int64
	/*PageSize
	  The size of the data page

	*/
	PageSize *int64
	/*Type
	  The task type: AFFILIATION, FUNDING, etc.

	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get API v10 tasks params
func (o *GetAPIV10TasksParams) WithTimeout(timeout time.Duration) *GetAPIV10TasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v10 tasks params
func (o *GetAPIV10TasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v10 tasks params
func (o *GetAPIV10TasksParams) WithContext(ctx context.Context) *GetAPIV10TasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v10 tasks params
func (o *GetAPIV10TasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v10 tasks params
func (o *GetAPIV10TasksParams) WithHTTPClient(client *http.Client) *GetAPIV10TasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v10 tasks params
func (o *GetAPIV10TasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get API v10 tasks params
func (o *GetAPIV10TasksParams) WithPage(page *int64) *GetAPIV10TasksParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get API v10 tasks params
func (o *GetAPIV10TasksParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the get API v10 tasks params
func (o *GetAPIV10TasksParams) WithPageSize(pageSize *int64) *GetAPIV10TasksParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get API v10 tasks params
func (o *GetAPIV10TasksParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithType adds the typeVar to the get API v10 tasks params
func (o *GetAPIV10TasksParams) WithType(typeVar *string) *GetAPIV10TasksParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get API v10 tasks params
func (o *GetAPIV10TasksParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV10TasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int64
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}

	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
