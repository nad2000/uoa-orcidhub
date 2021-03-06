// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Hremployee hremployee
// swagger:model Hremployee
type Hremployee struct {

	// academic staff f t e
	AcademicStaffFTE int32 `json:"academicStaffFTE,omitempty"`

	// employee ID
	EmployeeID string `json:"employeeID,omitempty"`

	// job
	Job []*Job `json:"job"`

	// professional staff f t e
	ProfessionalStaffFTE int32 `json:"professionalStaffFTE,omitempty"`

	// request time stamp
	// Format: date-time
	RequestTimeStamp strfmt.DateTime `json:"requestTimeStamp,omitempty"`

	// uni services f t e
	UniServicesFTE int64 `json:"uniServicesFTE,omitempty"`
}

// Validate validates this hremployee
func (m *Hremployee) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJob(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestTimeStamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Hremployee) validateJob(formats strfmt.Registry) error {

	if swag.IsZero(m.Job) { // not required
		return nil
	}

	for i := 0; i < len(m.Job); i++ {
		if swag.IsZero(m.Job[i]) { // not required
			continue
		}

		if m.Job[i] != nil {
			if err := m.Job[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("job" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Hremployee) validateRequestTimeStamp(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestTimeStamp) { // not required
		return nil
	}

	if err := validate.FormatOf("requestTimeStamp", "body", "date-time", m.RequestTimeStamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Hremployee) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Hremployee) UnmarshalBinary(b []byte) error {
	var res Hremployee
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
