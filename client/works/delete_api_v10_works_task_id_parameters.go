// Code generated by go-swagger; DO NOT EDIT.

package works

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

// NewDeleteAPIV10WorksTaskIDParams creates a new DeleteAPIV10WorksTaskIDParams object
// with the default values initialized.
func NewDeleteAPIV10WorksTaskIDParams() *DeleteAPIV10WorksTaskIDParams {
	var ()
	return &DeleteAPIV10WorksTaskIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPIV10WorksTaskIDParamsWithTimeout creates a new DeleteAPIV10WorksTaskIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAPIV10WorksTaskIDParamsWithTimeout(timeout time.Duration) *DeleteAPIV10WorksTaskIDParams {
	var ()
	return &DeleteAPIV10WorksTaskIDParams{

		timeout: timeout,
	}
}

// NewDeleteAPIV10WorksTaskIDParamsWithContext creates a new DeleteAPIV10WorksTaskIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAPIV10WorksTaskIDParamsWithContext(ctx context.Context) *DeleteAPIV10WorksTaskIDParams {
	var ()
	return &DeleteAPIV10WorksTaskIDParams{

		Context: ctx,
	}
}

// NewDeleteAPIV10WorksTaskIDParamsWithHTTPClient creates a new DeleteAPIV10WorksTaskIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAPIV10WorksTaskIDParamsWithHTTPClient(client *http.Client) *DeleteAPIV10WorksTaskIDParams {
	var ()
	return &DeleteAPIV10WorksTaskIDParams{
		HTTPClient: client,
	}
}

/*DeleteAPIV10WorksTaskIDParams contains all the parameters to send to the API endpoint
for the delete API v10 works task ID operation typically these are written to a http.Request
*/
type DeleteAPIV10WorksTaskIDParams struct {

	/*TaskID
	  Work task ID.

	*/
	TaskID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete API v10 works task ID params
func (o *DeleteAPIV10WorksTaskIDParams) WithTimeout(timeout time.Duration) *DeleteAPIV10WorksTaskIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API v10 works task ID params
func (o *DeleteAPIV10WorksTaskIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API v10 works task ID params
func (o *DeleteAPIV10WorksTaskIDParams) WithContext(ctx context.Context) *DeleteAPIV10WorksTaskIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API v10 works task ID params
func (o *DeleteAPIV10WorksTaskIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API v10 works task ID params
func (o *DeleteAPIV10WorksTaskIDParams) WithHTTPClient(client *http.Client) *DeleteAPIV10WorksTaskIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API v10 works task ID params
func (o *DeleteAPIV10WorksTaskIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTaskID adds the taskID to the delete API v10 works task ID params
func (o *DeleteAPIV10WorksTaskIDParams) WithTaskID(taskID int64) *DeleteAPIV10WorksTaskIDParams {
	o.SetTaskID(taskID)
	return o
}

// SetTaskID adds the taskId to the delete API v10 works task ID params
func (o *DeleteAPIV10WorksTaskIDParams) SetTaskID(taskID int64) {
	o.TaskID = taskID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPIV10WorksTaskIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param task_id
	if err := r.SetPathParam("task_id", swag.FormatInt64(o.TaskID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}