// Code generated by go-swagger; DO NOT EDIT.

package corporation

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

// NewGetCorporationsCorporationIDMembertrackingParams creates a new GetCorporationsCorporationIDMembertrackingParams object
// with the default values initialized.
func NewGetCorporationsCorporationIDMembertrackingParams() *GetCorporationsCorporationIDMembertrackingParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDMembertrackingParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCorporationsCorporationIDMembertrackingParamsWithTimeout creates a new GetCorporationsCorporationIDMembertrackingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCorporationsCorporationIDMembertrackingParamsWithTimeout(timeout time.Duration) *GetCorporationsCorporationIDMembertrackingParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDMembertrackingParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCorporationsCorporationIDMembertrackingParamsWithContext creates a new GetCorporationsCorporationIDMembertrackingParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCorporationsCorporationIDMembertrackingParamsWithContext(ctx context.Context) *GetCorporationsCorporationIDMembertrackingParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDMembertrackingParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetCorporationsCorporationIDMembertrackingParamsWithHTTPClient creates a new GetCorporationsCorporationIDMembertrackingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCorporationsCorporationIDMembertrackingParamsWithHTTPClient(client *http.Client) *GetCorporationsCorporationIDMembertrackingParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDMembertrackingParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetCorporationsCorporationIDMembertrackingParams contains all the parameters to send to the API endpoint
for the get corporations corporation id membertracking operation typically these are written to a http.Request
*/
type GetCorporationsCorporationIDMembertrackingParams struct {

	/*XUserAgent
	  Client identifier, takes precedence over User-Agent

	*/
	XUserAgent *string
	/*CorporationID
	  An EVE corporation ID

	*/
	CorporationID int32
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

// WithTimeout adds the timeout to the get corporations corporation id membertracking params
func (o *GetCorporationsCorporationIDMembertrackingParams) WithTimeout(timeout time.Duration) *GetCorporationsCorporationIDMembertrackingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get corporations corporation id membertracking params
func (o *GetCorporationsCorporationIDMembertrackingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get corporations corporation id membertracking params
func (o *GetCorporationsCorporationIDMembertrackingParams) WithContext(ctx context.Context) *GetCorporationsCorporationIDMembertrackingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get corporations corporation id membertracking params
func (o *GetCorporationsCorporationIDMembertrackingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get corporations corporation id membertracking params
func (o *GetCorporationsCorporationIDMembertrackingParams) WithHTTPClient(client *http.Client) *GetCorporationsCorporationIDMembertrackingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get corporations corporation id membertracking params
func (o *GetCorporationsCorporationIDMembertrackingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the get corporations corporation id membertracking params
func (o *GetCorporationsCorporationIDMembertrackingParams) WithXUserAgent(xUserAgent *string) *GetCorporationsCorporationIDMembertrackingParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the get corporations corporation id membertracking params
func (o *GetCorporationsCorporationIDMembertrackingParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithCorporationID adds the corporationID to the get corporations corporation id membertracking params
func (o *GetCorporationsCorporationIDMembertrackingParams) WithCorporationID(corporationID int32) *GetCorporationsCorporationIDMembertrackingParams {
	o.SetCorporationID(corporationID)
	return o
}

// SetCorporationID adds the corporationId to the get corporations corporation id membertracking params
func (o *GetCorporationsCorporationIDMembertrackingParams) SetCorporationID(corporationID int32) {
	o.CorporationID = corporationID
}

// WithDatasource adds the datasource to the get corporations corporation id membertracking params
func (o *GetCorporationsCorporationIDMembertrackingParams) WithDatasource(datasource *string) *GetCorporationsCorporationIDMembertrackingParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get corporations corporation id membertracking params
func (o *GetCorporationsCorporationIDMembertrackingParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithToken adds the token to the get corporations corporation id membertracking params
func (o *GetCorporationsCorporationIDMembertrackingParams) WithToken(token *string) *GetCorporationsCorporationIDMembertrackingParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get corporations corporation id membertracking params
func (o *GetCorporationsCorporationIDMembertrackingParams) SetToken(token *string) {
	o.Token = token
}

// WithUserAgent adds the userAgent to the get corporations corporation id membertracking params
func (o *GetCorporationsCorporationIDMembertrackingParams) WithUserAgent(userAgent *string) *GetCorporationsCorporationIDMembertrackingParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get corporations corporation id membertracking params
func (o *GetCorporationsCorporationIDMembertrackingParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *GetCorporationsCorporationIDMembertrackingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param corporation_id
	if err := r.SetPathParam("corporation_id", swag.FormatInt32(o.CorporationID)); err != nil {
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
