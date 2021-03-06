// Code generated by go-swagger; DO NOT EDIT.

package insurance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new insurance API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for insurance API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetInsurancePrices lists insurance levels

Return available insurance levels for all ship types

---
Alternate route: `/v1/insurance/prices/`

Alternate route: `/legacy/insurance/prices/`

Alternate route: `/dev/insurance/prices/`

---
This route is cached for up to 3600 seconds
*/
func (a *Client) GetInsurancePrices(params *GetInsurancePricesParams) (*GetInsurancePricesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInsurancePricesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_insurance_prices",
		Method:             "GET",
		PathPattern:        "/insurance/prices/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInsurancePricesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetInsurancePricesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
