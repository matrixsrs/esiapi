// Code generated by go-swagger; DO NOT EDIT.

package mail

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

// NewGetCharactersCharacterIDMailMailIDParams creates a new GetCharactersCharacterIDMailMailIDParams object
// with the default values initialized.
func NewGetCharactersCharacterIDMailMailIDParams() *GetCharactersCharacterIDMailMailIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDMailMailIDParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDMailMailIDParamsWithTimeout creates a new GetCharactersCharacterIDMailMailIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCharactersCharacterIDMailMailIDParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDMailMailIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDMailMailIDParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDMailMailIDParamsWithContext creates a new GetCharactersCharacterIDMailMailIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCharactersCharacterIDMailMailIDParamsWithContext(ctx context.Context) *GetCharactersCharacterIDMailMailIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDMailMailIDParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetCharactersCharacterIDMailMailIDParamsWithHTTPClient creates a new GetCharactersCharacterIDMailMailIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCharactersCharacterIDMailMailIDParamsWithHTTPClient(client *http.Client) *GetCharactersCharacterIDMailMailIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDMailMailIDParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetCharactersCharacterIDMailMailIDParams contains all the parameters to send to the API endpoint
for the get characters character id mail mail id operation typically these are written to a http.Request
*/
type GetCharactersCharacterIDMailMailIDParams struct {

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
	/*MailID
	  An EVE mail ID

	*/
	MailID int32
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

// WithTimeout adds the timeout to the get characters character id mail mail id params
func (o *GetCharactersCharacterIDMailMailIDParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDMailMailIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id mail mail id params
func (o *GetCharactersCharacterIDMailMailIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id mail mail id params
func (o *GetCharactersCharacterIDMailMailIDParams) WithContext(ctx context.Context) *GetCharactersCharacterIDMailMailIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id mail mail id params
func (o *GetCharactersCharacterIDMailMailIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get characters character id mail mail id params
func (o *GetCharactersCharacterIDMailMailIDParams) WithHTTPClient(client *http.Client) *GetCharactersCharacterIDMailMailIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get characters character id mail mail id params
func (o *GetCharactersCharacterIDMailMailIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the get characters character id mail mail id params
func (o *GetCharactersCharacterIDMailMailIDParams) WithXUserAgent(xUserAgent *string) *GetCharactersCharacterIDMailMailIDParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the get characters character id mail mail id params
func (o *GetCharactersCharacterIDMailMailIDParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithCharacterID adds the characterID to the get characters character id mail mail id params
func (o *GetCharactersCharacterIDMailMailIDParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDMailMailIDParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id mail mail id params
func (o *GetCharactersCharacterIDMailMailIDParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id mail mail id params
func (o *GetCharactersCharacterIDMailMailIDParams) WithDatasource(datasource *string) *GetCharactersCharacterIDMailMailIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id mail mail id params
func (o *GetCharactersCharacterIDMailMailIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithMailID adds the mailID to the get characters character id mail mail id params
func (o *GetCharactersCharacterIDMailMailIDParams) WithMailID(mailID int32) *GetCharactersCharacterIDMailMailIDParams {
	o.SetMailID(mailID)
	return o
}

// SetMailID adds the mailId to the get characters character id mail mail id params
func (o *GetCharactersCharacterIDMailMailIDParams) SetMailID(mailID int32) {
	o.MailID = mailID
}

// WithToken adds the token to the get characters character id mail mail id params
func (o *GetCharactersCharacterIDMailMailIDParams) WithToken(token *string) *GetCharactersCharacterIDMailMailIDParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get characters character id mail mail id params
func (o *GetCharactersCharacterIDMailMailIDParams) SetToken(token *string) {
	o.Token = token
}

// WithUserAgent adds the userAgent to the get characters character id mail mail id params
func (o *GetCharactersCharacterIDMailMailIDParams) WithUserAgent(userAgent *string) *GetCharactersCharacterIDMailMailIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get characters character id mail mail id params
func (o *GetCharactersCharacterIDMailMailIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDMailMailIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param mail_id
	if err := r.SetPathParam("mail_id", swag.FormatInt32(o.MailID)); err != nil {
		return err
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
