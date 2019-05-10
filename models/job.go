// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Job job
// swagger:model Job
type Job struct {

	// company
	Company string `json:"company,omitempty"`

	// cost centre
	CostCentre string `json:"costCentre,omitempty"`

	// department description
	DepartmentDescription string `json:"departmentDescription,omitempty"`

	// department ID
	DepartmentID string `json:"departmentID,omitempty"`

	// effective date
	// Format: date
	EffectiveDate strfmt.Date `json:"effectiveDate,omitempty"`

	// effective sequence
	EffectiveSequence int32 `json:"effectiveSequence,omitempty"`

	// employee record
	EmployeeRecord int32 `json:"employeeRecord,omitempty"`

	// employee status
	EmployeeStatus string `json:"employeeStatus,omitempty"`

	// employee type
	EmployeeType string `json:"employeeType,omitempty"`

	// full time equivalent
	FullTimeEquivalent int32 `json:"fullTimeEquivalent,omitempty"`

	// hr status
	HrStatus string `json:"hrStatus,omitempty"`

	// job code
	JobCode string `json:"jobCode,omitempty"`

	// job code description
	JobCodeDescription string `json:"jobCodeDescription,omitempty"`

	// job end date
	// Format: date
	JobEndDate strfmt.Date `json:"jobEndDate,omitempty"`

	// job grade
	JobGrade string `json:"jobGrade,omitempty"`

	// job indicator
	JobIndicator string `json:"jobIndicator,omitempty"`

	// job start date
	// Format: date
	JobStartDate strfmt.Date `json:"jobStartDate,omitempty"`

	// last h raction
	LastHRaction string `json:"lastHRaction,omitempty"`

	// location
	Location string `json:"location,omitempty"`

	// location description
	LocationDescription string `json:"locationDescription,omitempty"`

	// organizational relation
	OrganizationalRelation string `json:"organizationalRelation,omitempty"`

	// parent department description
	ParentDepartmentDescription string `json:"parentDepartmentDescription,omitempty"`

	// poi type
	PoiType string `json:"poiType,omitempty"`

	// position description
	PositionDescription string `json:"positionDescription,omitempty"`

	// position number
	PositionNumber string `json:"positionNumber,omitempty"`

	// primary activity centre dept description
	PrimaryActivityCentreDeptDescription string `json:"primaryActivityCentreDeptDescription,omitempty"`

	// primary activity centre dept ID
	PrimaryActivityCentreDeptID string `json:"primaryActivityCentreDeptID,omitempty"`

	// reports to position
	ReportsToPosition string `json:"reportsToPosition,omitempty"`

	// sal admin plan
	SalAdminPlan string `json:"salAdminPlan,omitempty"`

	// standard hours
	StandardHours float32 `json:"standardHours,omitempty"`

	// supervisor ID
	SupervisorID string `json:"supervisorID,omitempty"`

	// updated date time
	// Format: date-time
	UpdatedDateTime strfmt.DateTime `json:"updatedDateTime,omitempty"`
}

// Validate validates this job
func (m *Job) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEffectiveDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Job) validateEffectiveDate(formats strfmt.Registry) error {

	if swag.IsZero(m.EffectiveDate) { // not required
		return nil
	}

	if err := validate.FormatOf("effectiveDate", "body", "date", m.EffectiveDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Job) validateJobEndDate(formats strfmt.Registry) error {

	if swag.IsZero(m.JobEndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("jobEndDate", "body", "date", m.JobEndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Job) validateJobStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.JobStartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("jobStartDate", "body", "date", m.JobStartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Job) validateUpdatedDateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedDateTime", "body", "date-time", m.UpdatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Job) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Job) UnmarshalBinary(b []byte) error {
	var res Job
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}