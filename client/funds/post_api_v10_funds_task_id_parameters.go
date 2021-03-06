// Code generated by go-swagger; DO NOT EDIT.

package funds

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

	models "github.com/nad2000/uoa-orcidhub/models"
)

// NewPostAPIV10FundsTaskIDParams creates a new PostAPIV10FundsTaskIDParams object
// with the default values initialized.
func NewPostAPIV10FundsTaskIDParams() *PostAPIV10FundsTaskIDParams {
	var ()
	return &PostAPIV10FundsTaskIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV10FundsTaskIDParamsWithTimeout creates a new PostAPIV10FundsTaskIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPIV10FundsTaskIDParamsWithTimeout(timeout time.Duration) *PostAPIV10FundsTaskIDParams {
	var ()
	return &PostAPIV10FundsTaskIDParams{

		timeout: timeout,
	}
}

// NewPostAPIV10FundsTaskIDParamsWithContext creates a new PostAPIV10FundsTaskIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPIV10FundsTaskIDParamsWithContext(ctx context.Context) *PostAPIV10FundsTaskIDParams {
	var ()
	return &PostAPIV10FundsTaskIDParams{

		Context: ctx,
	}
}

// NewPostAPIV10FundsTaskIDParamsWithHTTPClient creates a new PostAPIV10FundsTaskIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPIV10FundsTaskIDParamsWithHTTPClient(client *http.Client) *PostAPIV10FundsTaskIDParams {
	var ()
	return &PostAPIV10FundsTaskIDParams{
		HTTPClient: client,
	}
}

/*PostAPIV10FundsTaskIDParams contains all the parameters to send to the API endpoint
for the post API v10 funds task ID operation typically these are written to a http.Request
*/
type PostAPIV10FundsTaskIDParams struct {

	/*FundTask
	  Fund task.

	*/
	FundTask *models.FundTask
	/*TaskID
	  Fund task ID.

	*/
	TaskID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API v10 funds task ID params
func (o *PostAPIV10FundsTaskIDParams) WithTimeout(timeout time.Duration) *PostAPIV10FundsTaskIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v10 funds task ID params
func (o *PostAPIV10FundsTaskIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v10 funds task ID params
func (o *PostAPIV10FundsTaskIDParams) WithContext(ctx context.Context) *PostAPIV10FundsTaskIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v10 funds task ID params
func (o *PostAPIV10FundsTaskIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v10 funds task ID params
func (o *PostAPIV10FundsTaskIDParams) WithHTTPClient(client *http.Client) *PostAPIV10FundsTaskIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v10 funds task ID params
func (o *PostAPIV10FundsTaskIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFundTask adds the fundTask to the post API v10 funds task ID params
func (o *PostAPIV10FundsTaskIDParams) WithFundTask(fundTask *models.FundTask) *PostAPIV10FundsTaskIDParams {
	o.SetFundTask(fundTask)
	return o
}

// SetFundTask adds the fundTask to the post API v10 funds task ID params
func (o *PostAPIV10FundsTaskIDParams) SetFundTask(fundTask *models.FundTask) {
	o.FundTask = fundTask
}

// WithTaskID adds the taskID to the post API v10 funds task ID params
func (o *PostAPIV10FundsTaskIDParams) WithTaskID(taskID int64) *PostAPIV10FundsTaskIDParams {
	o.SetTaskID(taskID)
	return o
}

// SetTaskID adds the taskId to the post API v10 funds task ID params
func (o *PostAPIV10FundsTaskIDParams) SetTaskID(taskID int64) {
	o.TaskID = taskID
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV10FundsTaskIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FundTask != nil {
		if err := r.SetBodyParam(o.FundTask); err != nil {
			return err
		}
	}

	// path param task_id
	if err := r.SetPathParam("task_id", swag.FormatInt64(o.TaskID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
