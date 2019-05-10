// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FundTask fund task
// swagger:model FundTask
type FundTask struct {

	// completed at
	// Format: date-time
	CompletedAt strfmt.DateTime `json:"completed-at,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created-at,omitempty"`

	// expires at
	// Format: date-time
	ExpiresAt strfmt.DateTime `json:"expires-at,omitempty"`

	// filename
	Filename string `json:"filename,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// records
	Records []FundTaskRecord `json:"records"`

	// task type
	// Enum: [fund FUNDING]
	TaskType string `json:"task-type,omitempty"`
}

// Validate validates this fund task
func (m *FundTask) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompletedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaskType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FundTask) validateCompletedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CompletedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("completed-at", "body", "date-time", m.CompletedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FundTask) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created-at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FundTask) validateExpiresAt(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpiresAt) { // not required
		return nil
	}

	if err := validate.FormatOf("expires-at", "body", "date-time", m.ExpiresAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var fundTaskTypeTaskTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["fund","FUNDING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fundTaskTypeTaskTypePropEnum = append(fundTaskTypeTaskTypePropEnum, v)
	}
}

const (

	// FundTaskTaskTypeFund captures enum value "fund"
	FundTaskTaskTypeFund string = "fund"

	// FundTaskTaskTypeFUNDING captures enum value "FUNDING"
	FundTaskTaskTypeFUNDING string = "FUNDING"
)

// prop value enum
func (m *FundTask) validateTaskTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, fundTaskTypeTaskTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FundTask) validateTaskType(formats strfmt.Registry) error {

	if swag.IsZero(m.TaskType) { // not required
		return nil
	}

	// value enum
	if err := m.validateTaskTypeEnum("task-type", "body", m.TaskType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FundTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FundTask) UnmarshalBinary(b []byte) error {
	var res FundTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}