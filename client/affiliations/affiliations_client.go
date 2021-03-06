// Code generated by go-swagger; DO NOT EDIT.

package affiliations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new affiliations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for affiliations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteAPIV10AffiliationsTaskID deletes the specified affiliation task
*/
func (a *Client) DeleteAPIV10AffiliationsTaskID(params *DeleteAPIV10AffiliationsTaskIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteAPIV10AffiliationsTaskIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAPIV10AffiliationsTaskIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteAPIV10AffiliationsTaskID",
		Method:             "DELETE",
		PathPattern:        "/api/v1.0/affiliations/{task_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAPIV10AffiliationsTaskIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteAPIV10AffiliationsTaskIDOK), nil

}

/*
GetAPIV10AffiliationsTaskID retrieves the specified affiliation task
*/
func (a *Client) GetAPIV10AffiliationsTaskID(params *GetAPIV10AffiliationsTaskIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetAPIV10AffiliationsTaskIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV10AffiliationsTaskIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAPIV10AffiliationsTaskID",
		Method:             "GET",
		PathPattern:        "/api/v1.0/affiliations/{task_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIV10AffiliationsTaskIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAPIV10AffiliationsTaskIDOK), nil

}

/*
PatchAPIV10AffiliationsTaskID updates the affiliation task
*/
func (a *Client) PatchAPIV10AffiliationsTaskID(params *PatchAPIV10AffiliationsTaskIDParams, authInfo runtime.ClientAuthInfoWriter) (*PatchAPIV10AffiliationsTaskIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAPIV10AffiliationsTaskIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchAPIV10AffiliationsTaskID",
		Method:             "PATCH",
		PathPattern:        "/api/v1.0/affiliations/{task_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchAPIV10AffiliationsTaskIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchAPIV10AffiliationsTaskIDOK), nil

}

/*
PostAPIV10Affiliations uploads the affiliation task
*/
func (a *Client) PostAPIV10Affiliations(params *PostAPIV10AffiliationsParams, authInfo runtime.ClientAuthInfoWriter) (*PostAPIV10AffiliationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPIV10AffiliationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAPIV10Affiliations",
		Method:             "POST",
		PathPattern:        "/api/v1.0/affiliations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/csv", "text/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAPIV10AffiliationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAPIV10AffiliationsOK), nil

}

/*
PostAPIV10AffiliationsTaskID uploads the task and completely override the affiliation task
*/
func (a *Client) PostAPIV10AffiliationsTaskID(params *PostAPIV10AffiliationsTaskIDParams, authInfo runtime.ClientAuthInfoWriter) (*PostAPIV10AffiliationsTaskIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPIV10AffiliationsTaskIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAPIV10AffiliationsTaskID",
		Method:             "POST",
		PathPattern:        "/api/v1.0/affiliations/{task_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAPIV10AffiliationsTaskIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAPIV10AffiliationsTaskIDOK), nil

}

/*
PutAPIV10AffiliationsTaskID updates the affiliation task
*/
func (a *Client) PutAPIV10AffiliationsTaskID(params *PutAPIV10AffiliationsTaskIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutAPIV10AffiliationsTaskIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutAPIV10AffiliationsTaskIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutAPIV10AffiliationsTaskID",
		Method:             "PUT",
		PathPattern:        "/api/v1.0/affiliations/{task_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutAPIV10AffiliationsTaskIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutAPIV10AffiliationsTaskIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
