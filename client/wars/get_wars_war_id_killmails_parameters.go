// Code generated by go-swagger; DO NOT EDIT.

package wars

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

// NewGetWarsWarIDKillmailsParams creates a new GetWarsWarIDKillmailsParams object
// with the default values initialized.
func NewGetWarsWarIDKillmailsParams() *GetWarsWarIDKillmailsParams {
	var (
		datasourceDefault = string("tranquility")
		pageDefault       = int32(1)
	)
	return &GetWarsWarIDKillmailsParams{
		Datasource: &datasourceDefault,
		Page:       &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWarsWarIDKillmailsParamsWithTimeout creates a new GetWarsWarIDKillmailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWarsWarIDKillmailsParamsWithTimeout(timeout time.Duration) *GetWarsWarIDKillmailsParams {
	var (
		datasourceDefault = string("tranquility")
		pageDefault       = int32(1)
	)
	return &GetWarsWarIDKillmailsParams{
		Datasource: &datasourceDefault,
		Page:       &pageDefault,

		timeout: timeout,
	}
}

// NewGetWarsWarIDKillmailsParamsWithContext creates a new GetWarsWarIDKillmailsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWarsWarIDKillmailsParamsWithContext(ctx context.Context) *GetWarsWarIDKillmailsParams {
	var (
		datasourceDefault = string("tranquility")
		pageDefault       = int32(1)
	)
	return &GetWarsWarIDKillmailsParams{
		Datasource: &datasourceDefault,
		Page:       &pageDefault,

		Context: ctx,
	}
}

// NewGetWarsWarIDKillmailsParamsWithHTTPClient creates a new GetWarsWarIDKillmailsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWarsWarIDKillmailsParamsWithHTTPClient(client *http.Client) *GetWarsWarIDKillmailsParams {
	var (
		datasourceDefault = string("tranquility")
		pageDefault       = int32(1)
	)
	return &GetWarsWarIDKillmailsParams{
		Datasource: &datasourceDefault,
		Page:       &pageDefault,
		HTTPClient: client,
	}
}

/*GetWarsWarIDKillmailsParams contains all the parameters to send to the API endpoint
for the get wars war id killmails operation typically these are written to a http.Request
*/
type GetWarsWarIDKillmailsParams struct {

	/*XUserAgent
	  Client identifier, takes precedence over User-Agent

	*/
	XUserAgent *string
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*Page
	  Which page of results to return

	*/
	Page *int32
	/*UserAgent
	  Client identifier, takes precedence over headers

	*/
	UserAgent *string
	/*WarID
	  A valid war ID

	*/
	WarID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get wars war id killmails params
func (o *GetWarsWarIDKillmailsParams) WithTimeout(timeout time.Duration) *GetWarsWarIDKillmailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get wars war id killmails params
func (o *GetWarsWarIDKillmailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get wars war id killmails params
func (o *GetWarsWarIDKillmailsParams) WithContext(ctx context.Context) *GetWarsWarIDKillmailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get wars war id killmails params
func (o *GetWarsWarIDKillmailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get wars war id killmails params
func (o *GetWarsWarIDKillmailsParams) WithHTTPClient(client *http.Client) *GetWarsWarIDKillmailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get wars war id killmails params
func (o *GetWarsWarIDKillmailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the get wars war id killmails params
func (o *GetWarsWarIDKillmailsParams) WithXUserAgent(xUserAgent *string) *GetWarsWarIDKillmailsParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the get wars war id killmails params
func (o *GetWarsWarIDKillmailsParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithDatasource adds the datasource to the get wars war id killmails params
func (o *GetWarsWarIDKillmailsParams) WithDatasource(datasource *string) *GetWarsWarIDKillmailsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get wars war id killmails params
func (o *GetWarsWarIDKillmailsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithPage adds the page to the get wars war id killmails params
func (o *GetWarsWarIDKillmailsParams) WithPage(page *int32) *GetWarsWarIDKillmailsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get wars war id killmails params
func (o *GetWarsWarIDKillmailsParams) SetPage(page *int32) {
	o.Page = page
}

// WithUserAgent adds the userAgent to the get wars war id killmails params
func (o *GetWarsWarIDKillmailsParams) WithUserAgent(userAgent *string) *GetWarsWarIDKillmailsParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get wars war id killmails params
func (o *GetWarsWarIDKillmailsParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithWarID adds the warID to the get wars war id killmails params
func (o *GetWarsWarIDKillmailsParams) WithWarID(warID int32) *GetWarsWarIDKillmailsParams {
	o.SetWarID(warID)
	return o
}

// SetWarID adds the warId to the get wars war id killmails params
func (o *GetWarsWarIDKillmailsParams) SetWarID(warID int32) {
	o.WarID = warID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWarsWarIDKillmailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Page != nil {

		// query param page
		var qrPage int32
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
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

	// path param war_id
	if err := r.SetPathParam("war_id", swag.FormatInt32(o.WarID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
