// Code generated by go-swagger; DO NOT EDIT.

package assets

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

// NewGetCharactersCharacterIDAssetsParams creates a new GetCharactersCharacterIDAssetsParams object
// with the default values initialized.
func NewGetCharactersCharacterIDAssetsParams() *GetCharactersCharacterIDAssetsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDAssetsParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDAssetsParamsWithTimeout creates a new GetCharactersCharacterIDAssetsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCharactersCharacterIDAssetsParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDAssetsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDAssetsParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDAssetsParamsWithContext creates a new GetCharactersCharacterIDAssetsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCharactersCharacterIDAssetsParamsWithContext(ctx context.Context) *GetCharactersCharacterIDAssetsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDAssetsParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetCharactersCharacterIDAssetsParamsWithHTTPClient creates a new GetCharactersCharacterIDAssetsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCharactersCharacterIDAssetsParamsWithHTTPClient(client *http.Client) *GetCharactersCharacterIDAssetsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDAssetsParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetCharactersCharacterIDAssetsParams contains all the parameters to send to the API endpoint
for the get characters character id assets operation typically these are written to a http.Request
*/
type GetCharactersCharacterIDAssetsParams struct {

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
	/*Token
	  Access token to use if unable to set a header

	*/
	Token *string
	/*UserAgent
	  Client identifier, takes precedence over headers

	*/
	UserAgent *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDAssetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) WithContext(ctx context.Context) *GetCharactersCharacterIDAssetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) WithHTTPClient(client *http.Client) *GetCharactersCharacterIDAssetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) WithXUserAgent(xUserAgent *string) *GetCharactersCharacterIDAssetsParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithCharacterID adds the characterID to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDAssetsParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) WithDatasource(datasource *string) *GetCharactersCharacterIDAssetsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithToken adds the token to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) WithToken(token *string) *GetCharactersCharacterIDAssetsParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) SetToken(token *string) {
	o.Token = token
}

// WithUserAgent adds the userAgent to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) WithUserAgent(userAgent *string) *GetCharactersCharacterIDAssetsParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDAssetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Token != nil {

		// query param token
		var qrToken string
		if o.Token != nil {
			qrToken = *o.Token
		}
		qToken := qrToken
		if qToken != "" {
			if err := r.SetQueryParam("token", qToken); err != nil {
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
