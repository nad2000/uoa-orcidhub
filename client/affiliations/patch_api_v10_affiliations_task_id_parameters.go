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

// NewPatchAPIV10AffiliationsTaskIDParams creates a new PatchAPIV10AffiliationsTaskIDParams object
// with the default values initialized.
func NewPatchAPIV10AffiliationsTaskIDParams() *PatchAPIV10AffiliationsTaskIDParams {
	var ()
	return &PatchAPIV10AffiliationsTaskIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPIV10AffiliationsTaskIDParamsWithTimeout creates a new PatchAPIV10AffiliationsTaskIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchAPIV10AffiliationsTaskIDParamsWithTimeout(timeout time.Duration) *PatchAPIV10AffiliationsTaskIDParams {
	var ()
	return &PatchAPIV10AffiliationsTaskIDParams{

		timeout: timeout,
	}
}

// NewPatchAPIV10AffiliationsTaskIDParamsWithContext creates a new PatchAPIV10AffiliationsTaskIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchAPIV10AffiliationsTaskIDParamsWithContext(ctx context.Context) *PatchAPIV10AffiliationsTaskIDParams {
	var ()
	return &PatchAPIV10AffiliationsTaskIDParams{

		Context: ctx,
	}
}

// NewPatchAPIV10AffiliationsTaskIDParamsWithHTTPClient creates a new PatchAPIV10AffiliationsTaskIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchAPIV10AffiliationsTaskIDParamsWithHTTPClient(client *http.Client) *PatchAPIV10AffiliationsTaskIDParams {
	var ()
	return &PatchAPIV10AffiliationsTaskIDParams{
		HTTPClient: client,
	}
}

/*PatchAPIV10AffiliationsTaskIDParams contains all the parameters to send to the API endpoint
for the patch API v10 affiliations task ID operation typically these are written to a http.Request
*/
type PatchAPIV10AffiliationsTaskIDParams struct {

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

// WithTimeout adds the timeout to the patch API v10 affiliations task ID params
func (o *PatchAPIV10AffiliationsTaskIDParams) WithTimeout(timeout time.Duration) *PatchAPIV10AffiliationsTaskIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API v10 affiliations task ID params
func (o *PatchAPIV10AffiliationsTaskIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API v10 affiliations task ID params
func (o *PatchAPIV10AffiliationsTaskIDParams) WithContext(ctx context.Context) *PatchAPIV10AffiliationsTaskIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API v10 affiliations task ID params
func (o *PatchAPIV10AffiliationsTaskIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API v10 affiliations task ID params
func (o *PatchAPIV10AffiliationsTaskIDParams) WithHTTPClient(client *http.Client) *PatchAPIV10AffiliationsTaskIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API v10 affiliations task ID params
func (o *PatchAPIV10AffiliationsTaskIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAffiliationTask adds the affiliationTask to the patch API v10 affiliations task ID params
func (o *PatchAPIV10AffiliationsTaskIDParams) WithAffiliationTask(affiliationTask *models.AffiliationTask) *PatchAPIV10AffiliationsTaskIDParams {
	o.SetAffiliationTask(affiliationTask)
	return o
}

// SetAffiliationTask adds the affiliationTask to the patch API v10 affiliations task ID params
func (o *PatchAPIV10AffiliationsTaskIDParams) SetAffiliationTask(affiliationTask *models.AffiliationTask) {
	o.AffiliationTask = affiliationTask
}

// WithTaskID adds the taskID to the patch API v10 affiliations task ID params
func (o *PatchAPIV10AffiliationsTaskIDParams) WithTaskID(taskID int64) *PatchAPIV10AffiliationsTaskIDParams {
	o.SetTaskID(taskID)
	return o
}

// SetTaskID adds the taskId to the patch API v10 affiliations task ID params
func (o *PatchAPIV10AffiliationsTaskIDParams) SetTaskID(taskID int64) {
	o.TaskID = taskID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPIV10AffiliationsTaskIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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