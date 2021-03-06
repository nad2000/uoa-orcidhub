// Code generated by go-swagger; DO NOT EDIT.

package affiliations

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

// NewPostAPIV10AffiliationsTaskIDParams creates a new PostAPIV10AffiliationsTaskIDParams object
// with the default values initialized.
func NewPostAPIV10AffiliationsTaskIDParams() *PostAPIV10AffiliationsTaskIDParams {
	var ()
	return &PostAPIV10AffiliationsTaskIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV10AffiliationsTaskIDParamsWithTimeout creates a new PostAPIV10AffiliationsTaskIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPIV10AffiliationsTaskIDParamsWithTimeout(timeout time.Duration) *PostAPIV10AffiliationsTaskIDParams {
	var ()
	return &PostAPIV10AffiliationsTaskIDParams{

		timeout: timeout,
	}
}

// NewPostAPIV10AffiliationsTaskIDParamsWithContext creates a new PostAPIV10AffiliationsTaskIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPIV10AffiliationsTaskIDParamsWithContext(ctx context.Context) *PostAPIV10AffiliationsTaskIDParams {
	var ()
	return &PostAPIV10AffiliationsTaskIDParams{

		Context: ctx,
	}
}

// NewPostAPIV10AffiliationsTaskIDParamsWithHTTPClient creates a new PostAPIV10AffiliationsTaskIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPIV10AffiliationsTaskIDParamsWithHTTPClient(client *http.Client) *PostAPIV10AffiliationsTaskIDParams {
	var ()
	return &PostAPIV10AffiliationsTaskIDParams{
		HTTPClient: client,
	}
}

/*PostAPIV10AffiliationsTaskIDParams contains all the parameters to send to the API endpoint
for the post API v10 affiliations task ID operation typically these are written to a http.Request
*/
type PostAPIV10AffiliationsTaskIDParams struct {

	/*AffiliationTask
	  Affiliation task.

	*/
	AffiliationTask *models.AffiliationTask
	/*TaskID
	  Affiliation task ID.

	*/
	TaskID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API v10 affiliations task ID params
func (o *PostAPIV10AffiliationsTaskIDParams) WithTimeout(timeout time.Duration) *PostAPIV10AffiliationsTaskIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v10 affiliations task ID params
func (o *PostAPIV10AffiliationsTaskIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v10 affiliations task ID params
func (o *PostAPIV10AffiliationsTaskIDParams) WithContext(ctx context.Context) *PostAPIV10AffiliationsTaskIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v10 affiliations task ID params
func (o *PostAPIV10AffiliationsTaskIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v10 affiliations task ID params
func (o *PostAPIV10AffiliationsTaskIDParams) WithHTTPClient(client *http.Client) *PostAPIV10AffiliationsTaskIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v10 affiliations task ID params
func (o *PostAPIV10AffiliationsTaskIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAffiliationTask adds the affiliationTask to the post API v10 affiliations task ID params
func (o *PostAPIV10AffiliationsTaskIDParams) WithAffiliationTask(affiliationTask *models.AffiliationTask) *PostAPIV10AffiliationsTaskIDParams {
	o.SetAffiliationTask(affiliationTask)
	return o
}

// SetAffiliationTask adds the affiliationTask to the post API v10 affiliations task ID params
func (o *PostAPIV10AffiliationsTaskIDParams) SetAffiliationTask(affiliationTask *models.AffiliationTask) {
	o.AffiliationTask = affiliationTask
}

// WithTaskID adds the taskID to the post API v10 affiliations task ID params
func (o *PostAPIV10AffiliationsTaskIDParams) WithTaskID(taskID int64) *PostAPIV10AffiliationsTaskIDParams {
	o.SetTaskID(taskID)
	return o
}

// SetTaskID adds the taskId to the post API v10 affiliations task ID params
func (o *PostAPIV10AffiliationsTaskIDParams) SetTaskID(taskID int64) {
	o.TaskID = taskID
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV10AffiliationsTaskIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AffiliationTask != nil {
		if err := r.SetBodyParam(o.AffiliationTask); err != nil {
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
