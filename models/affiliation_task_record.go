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

// AffiliationTaskRecord affiliation task record
// swagger:model AffiliationTaskRecord
type AffiliationTaskRecord struct {

	// affiliation type
	// Required: true
	// Enum: [student staff]
	AffiliationType *string `json:"affiliation-type"`

	// city
	City string `json:"city,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// department
	Department string `json:"department,omitempty"`

	// disambiguated id
	DisambiguatedID string `json:"disambiguated-id,omitempty"`

	// disambiguated source
	DisambiguatedSource string `json:"disambiguated-source,omitempty"`

	// email
	// Required: true
	Email *string `json:"email"`

	// end date
	EndDate string `json:"end-date,omitempty"`

	// external id
	ExternalID string `json:"external-id,omitempty"`

	// first name
	// Required: true
	FirstName *string `json:"first-name"`

	// id
	ID int64 `json:"id,omitempty"`

	// is active
	IsActive bool `json:"is-active,omitempty"`

	// last name
	// Required: true
	LastName *string `json:"last-name"`

	// User ORCID ID
	Orcid string `json:"orcid,omitempty"`

	// organisation
	Organisation string `json:"organisation,omitempty"`

	// processed at
	// Format: date-time
	ProcessedAt strfmt.DateTime `json:"processed-at,omitempty"`

	// put code
	PutCode string `json:"put-code,omitempty"`

	// role
	Role string `json:"role,omitempty"`

	// start date
	StartDate string `json:"start-date,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this affiliation task record
func (m *AffiliationTaskRecord) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAffiliationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var affiliationTaskRecordTypeAffiliationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["student","staff"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		affiliationTaskRecordTypeAffiliationTypePropEnum = append(affiliationTaskRecordTypeAffiliationTypePropEnum, v)
	}
}

const (

	// AffiliationTaskRecordAffiliationTypeStudent captures enum value "student"
	AffiliationTaskRecordAffiliationTypeStudent string = "student"

	// AffiliationTaskRecordAffiliationTypeStaff captures enum value "staff"
	AffiliationTaskRecordAffiliationTypeStaff string = "staff"
)

// prop value enum
func (m *AffiliationTaskRecord) validateAffiliationTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, affiliationTaskRecordTypeAffiliationTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AffiliationTaskRecord) validateAffiliationType(formats strfmt.Registry) error {

	if err := validate.Required("affiliation-type", "body", m.AffiliationType); err != nil {
		return err
	}

	// value enum
	if err := m.validateAffiliationTypeEnum("affiliation-type", "body", *m.AffiliationType); err != nil {
		return err
	}

	return nil
}

func (m *AffiliationTaskRecord) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *AffiliationTaskRecord) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("first-name", "body", m.FirstName); err != nil {
		return err
	}

	return nil
}

func (m *AffiliationTaskRecord) validateLastName(formats strfmt.Registry) error {

	if err := validate.Required("last-name", "body", m.LastName); err != nil {
		return err
	}

	return nil
}

func (m *AffiliationTaskRecord) validateProcessedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.ProcessedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("processed-at", "body", "date-time", m.ProcessedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AffiliationTaskRecord) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AffiliationTaskRecord) UnmarshalBinary(b []byte) error {
	var res AffiliationTaskRecord
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
