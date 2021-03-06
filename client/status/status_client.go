// Code generated by go-swagger; DO NOT EDIT.

package status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new status API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for status API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetStatus retrieves the uptime and player counts

EVE Server status

---
Alternate route: `/v1/status/`

Alternate route: `/legacy/status/`

Alternate route: `/dev/status/`

---
This route is cached for up to 30 seconds
*/
func (a *Client) GetStatus(params *GetStatusParams) (*GetStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_status",
		Method:             "GET",
		PathPattern:        "/status/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStatusOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
