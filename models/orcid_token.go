// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// OrcidToken orcid token
// swagger:model OrcidToken
type OrcidToken struct {

	// ORCID API user profile access token
	AccessToken string `json:"access_token,omitempty"`

	// expires in
	ExpiresIn int64 `json:"expires_in,omitempty"`

	// issue time
	IssueTime string `json:"issue_time,omitempty"`

	// ORCID API user profile refresh token
	RefreshToken string `json:"refresh_token,omitempty"`

	// ORCID API user token scopes
	Scopes string `json:"scopes,omitempty"`
}

// Validate validates this orcid token
func (m *OrcidToken) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OrcidToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrcidToken) UnmarshalBinary(b []byte) error {
	var res OrcidToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}