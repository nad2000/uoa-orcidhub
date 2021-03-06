// Code generated by go-swagger; DO NOT EDIT.

package token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new token API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for token API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetAPIV10TokensIdentifier retrieves user access token and refresh token
*/
func (a *Client) GetAPIV10TokensIdentifier(params *GetAPIV10TokensIdentifierParams, authInfo runtime.ClientAuthInfoWriter) (*GetAPIV10TokensIdentifierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV10TokensIdentifierParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAPIV10TokensIdentifier",
		Method:             "GET",
		PathPattern:        "/api/v1.0/tokens/{identifier}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIV10TokensIdentifierReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAPIV10TokensIdentifierOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
