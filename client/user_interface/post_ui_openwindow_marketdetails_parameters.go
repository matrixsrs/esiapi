// Code generated by go-swagger; DO NOT EDIT.

package user_interface

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

// NewPostUIOpenwindowMarketdetailsParams creates a new PostUIOpenwindowMarketdetailsParams object
// with the default values initialized.
func NewPostUIOpenwindowMarketdetailsParams() *PostUIOpenwindowMarketdetailsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostUIOpenwindowMarketdetailsParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostUIOpenwindowMarketdetailsParamsWithTimeout creates a new PostUIOpenwindowMarketdetailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostUIOpenwindowMarketdetailsParamsWithTimeout(timeout time.Duration) *PostUIOpenwindowMarketdetailsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostUIOpenwindowMarketdetailsParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewPostUIOpenwindowMarketdetailsParamsWithContext creates a new PostUIOpenwindowMarketdetailsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostUIOpenwindowMarketdetailsParamsWithContext(ctx context.Context) *PostUIOpenwindowMarketdetailsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostUIOpenwindowMarketdetailsParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewPostUIOpenwindowMarketdetailsParamsWithHTTPClient creates a new PostUIOpenwindowMarketdetailsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostUIOpenwindowMarketdetailsParamsWithHTTPClient(client *http.Client) *PostUIOpenwindowMarketdetailsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostUIOpenwindowMarketdetailsParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*PostUIOpenwindowMarketdetailsParams contains all the parameters to send to the API endpoint
for the post ui openwindow marketdetails operation typically these are written to a http.Request
*/
type PostUIOpenwindowMarketdetailsParams struct {

	/*XUserAgent
	  Client identifier, takes precedence over User-Agent

	*/
	XUserAgent *string
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*Token
	  Access token to use if unable to set a header

	*/
	Token *string
	/*TypeID
	  The item type to open in market window

	*/
	TypeID int32
	/*UserAgent
	  Client identifier, takes precedence over headers

	*/
	UserAgent *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post ui openwindow marketdetails params
func (o *PostUIOpenwindowMarketdetailsParams) WithTimeout(timeout time.Duration) *PostUIOpenwindowMarketdetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post ui openwindow marketdetails params
func (o *PostUIOpenwindowMarketdetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post ui openwindow marketdetails params
func (o *PostUIOpenwindowMarketdetailsParams) WithContext(ctx context.Context) *PostUIOpenwindowMarketdetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post ui openwindow marketdetails params
func (o *PostUIOpenwindowMarketdetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post ui openwindow marketdetails params
func (o *PostUIOpenwindowMarketdetailsParams) WithHTTPClient(client *http.Client) *PostUIOpenwindowMarketdetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post ui openwindow marketdetails params
func (o *PostUIOpenwindowMarketdetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the post ui openwindow marketdetails params
func (o *PostUIOpenwindowMarketdetailsParams) WithXUserAgent(xUserAgent *string) *PostUIOpenwindowMarketdetailsParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the post ui openwindow marketdetails params
func (o *PostUIOpenwindowMarketdetailsParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithDatasource adds the datasource to the post ui openwindow marketdetails params
func (o *PostUIOpenwindowMarketdetailsParams) WithDatasource(datasource *string) *PostUIOpenwindowMarketdetailsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the post ui openwindow marketdetails params
func (o *PostUIOpenwindowMarketdetailsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithToken adds the token to the post ui openwindow marketdetails params
func (o *PostUIOpenwindowMarketdetailsParams) WithToken(token *string) *PostUIOpenwindowMarketdetailsParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the post ui openwindow marketdetails params
func (o *PostUIOpenwindowMarketdetailsParams) SetToken(token *string) {
	o.Token = token
}

// WithTypeID adds the typeID to the post ui openwindow marketdetails params
func (o *PostUIOpenwindowMarketdetailsParams) WithTypeID(typeID int32) *PostUIOpenwindowMarketdetailsParams {
	o.SetTypeID(typeID)
	return o
}

// SetTypeID adds the typeId to the post ui openwindow marketdetails params
func (o *PostUIOpenwindowMarketdetailsParams) SetTypeID(typeID int32) {
	o.TypeID = typeID
}

// WithUserAgent adds the userAgent to the post ui openwindow marketdetails params
func (o *PostUIOpenwindowMarketdetailsParams) WithUserAgent(userAgent *string) *PostUIOpenwindowMarketdetailsParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the post ui openwindow marketdetails params
func (o *PostUIOpenwindowMarketdetailsParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *PostUIOpenwindowMarketdetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param type_id
	qrTypeID := o.TypeID
	qTypeID := swag.FormatInt32(qrTypeID)
	if qTypeID != "" {
		if err := r.SetQueryParam("type_id", qTypeID); err != nil {
			return err
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
