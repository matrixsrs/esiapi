// Code generated by go-swagger; DO NOT EDIT.

package character

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetCharactersCharacterIDParams creates a new GetCharactersCharacterIDParams object
// with the default values initialized.
func NewGetCharactersCharacterIDParams() *GetCharactersCharacterIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDParamsWithTimeout creates a new GetCharactersCharacterIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCharactersCharacterIDParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDParamsWithContext creates a new GetCharactersCharacterIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCharactersCharacterIDParamsWithContext(ctx context.Context) *GetCharactersCharacterIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetCharactersCharacterIDParamsWithHTTPClient creates a new GetCharactersCharacterIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCharactersCharacterIDParamsWithHTTPClient(client *http.Client) *GetCharactersCharacterIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetCharactersCharacterIDParams contains all the parameters to send to the API endpoint
for the get characters character id operation typically these are written to a http.Request
*/
type GetCharactersCharacterIDParams struct {

	/*XUserAgent
	  Client identifier, takes precedence over User-Agent

	*/
	XUserAgent *string
	/*CharacterID
	  An EVE character ID

	*/
	CharacterID int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*UserAgent
	  Client identifier, takes precedence over headers

	*/
	UserAgent *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get characters character id params
func (o *GetCharactersCharacterIDParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id params
func (o *GetCharactersCharacterIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id params
func (o *GetCharactersCharacterIDParams) WithContext(ctx context.Context) *GetCharactersCharacterIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id params
func (o *GetCharactersCharacterIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get characters character id params
func (o *GetCharactersCharacterIDParams) WithHTTPClient(client *http.Client) *GetCharactersCharacterIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get characters character id params
func (o *GetCharactersCharacterIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the get characters character id params
func (o *GetCharactersCharacterIDParams) WithXUserAgent(xUserAgent *string) *GetCharactersCharacterIDParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the get characters character id params
func (o *GetCharactersCharacterIDParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithCharacterID adds the characterID to the get characters character id params
func (o *GetCharactersCharacterIDParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id params
func (o *GetCharactersCharacterIDParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id params
func (o *GetCharactersCharacterIDParams) WithDatasource(datasource *string) *GetCharactersCharacterIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id params
func (o *GetCharactersCharacterIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithUserAgent adds the userAgent to the get characters character id params
func (o *GetCharactersCharacterIDParams) WithUserAgent(userAgent *string) *GetCharactersCharacterIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get characters character id params
func (o *GetCharactersCharacterIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XUserAgent != nil {

		// header param X-User-Agent
		if err := r.SetHeaderParam("X-User-Agent", *o.XUserAgent); err != nil {
			return err
		}

	}

	// path param character_id
	if err := r.SetPathParam("character_id", swag.FormatInt32(o.CharacterID)); err != nil {
		return err
	}

	if o.Datasource != nil {

		// query param datasource
		var qrDatasource string
		if o.Datasource != nil {
			qrDatasource = *o.Datasource
		}
		qDatasource := qrDatasource
		if qDatasource != "" {
			if err := r.SetQueryParam("datasource", qDatasource); err != nil {
				return err
			}
		}

	}

	if o.UserAgent != nil {

		// query param user_agent
		var qrUserAgent string
		if o.UserAgent != nil {
			qrUserAgent = *o.UserAgent
		}
		qUserAgent := qrUserAgent
		if qUserAgent != "" {
			if err := r.SetQueryParam("user_agent", qUserAgent); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
