// Code generated by go-swagger; DO NOT EDIT.

package universe

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

// NewGetUniverseGroupsGroupIDParams creates a new GetUniverseGroupsGroupIDParams object
// with the default values initialized.
func NewGetUniverseGroupsGroupIDParams() *GetUniverseGroupsGroupIDParams {
	var (
		datasourceDefault = string("tranquility")
		languageDefault   = string("en-us")
	)
	return &GetUniverseGroupsGroupIDParams{
		Datasource: &datasourceDefault,
		Language:   &languageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUniverseGroupsGroupIDParamsWithTimeout creates a new GetUniverseGroupsGroupIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUniverseGroupsGroupIDParamsWithTimeout(timeout time.Duration) *GetUniverseGroupsGroupIDParams {
	var (
		datasourceDefault = string("tranquility")
		languageDefault   = string("en-us")
	)
	return &GetUniverseGroupsGroupIDParams{
		Datasource: &datasourceDefault,
		Language:   &languageDefault,

		timeout: timeout,
	}
}

// NewGetUniverseGroupsGroupIDParamsWithContext creates a new GetUniverseGroupsGroupIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUniverseGroupsGroupIDParamsWithContext(ctx context.Context) *GetUniverseGroupsGroupIDParams {
	var (
		datasourceDefault = string("tranquility")
		languageDefault   = string("en-us")
	)
	return &GetUniverseGroupsGroupIDParams{
		Datasource: &datasourceDefault,
		Language:   &languageDefault,

		Context: ctx,
	}
}

// NewGetUniverseGroupsGroupIDParamsWithHTTPClient creates a new GetUniverseGroupsGroupIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUniverseGroupsGroupIDParamsWithHTTPClient(client *http.Client) *GetUniverseGroupsGroupIDParams {
	var (
		datasourceDefault = string("tranquility")
		languageDefault   = string("en-us")
	)
	return &GetUniverseGroupsGroupIDParams{
		Datasource: &datasourceDefault,
		Language:   &languageDefault,
		HTTPClient: client,
	}
}

/*GetUniverseGroupsGroupIDParams contains all the parameters to send to the API endpoint
for the get universe groups group id operation typically these are written to a http.Request
*/
type GetUniverseGroupsGroupIDParams struct {

	/*XUserAgent
	  Client identifier, takes precedence over User-Agent

	*/
	XUserAgent *string
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*GroupID
	  An Eve item group ID

	*/
	GroupID int32
	/*Language
	  Language to use in the response

	*/
	Language *string
	/*UserAgent
	  Client identifier, takes precedence over headers

	*/
	UserAgent *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get universe groups group id params
func (o *GetUniverseGroupsGroupIDParams) WithTimeout(timeout time.Duration) *GetUniverseGroupsGroupIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get universe groups group id params
func (o *GetUniverseGroupsGroupIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get universe groups group id params
func (o *GetUniverseGroupsGroupIDParams) WithContext(ctx context.Context) *GetUniverseGroupsGroupIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get universe groups group id params
func (o *GetUniverseGroupsGroupIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get universe groups group id params
func (o *GetUniverseGroupsGroupIDParams) WithHTTPClient(client *http.Client) *GetUniverseGroupsGroupIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get universe groups group id params
func (o *GetUniverseGroupsGroupIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the get universe groups group id params
func (o *GetUniverseGroupsGroupIDParams) WithXUserAgent(xUserAgent *string) *GetUniverseGroupsGroupIDParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the get universe groups group id params
func (o *GetUniverseGroupsGroupIDParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithDatasource adds the datasource to the get universe groups group id params
func (o *GetUniverseGroupsGroupIDParams) WithDatasource(datasource *string) *GetUniverseGroupsGroupIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get universe groups group id params
func (o *GetUniverseGroupsGroupIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithGroupID adds the groupID to the get universe groups group id params
func (o *GetUniverseGroupsGroupIDParams) WithGroupID(groupID int32) *GetUniverseGroupsGroupIDParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the get universe groups group id params
func (o *GetUniverseGroupsGroupIDParams) SetGroupID(groupID int32) {
	o.GroupID = groupID
}

// WithLanguage adds the language to the get universe groups group id params
func (o *GetUniverseGroupsGroupIDParams) WithLanguage(language *string) *GetUniverseGroupsGroupIDParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the get universe groups group id params
func (o *GetUniverseGroupsGroupIDParams) SetLanguage(language *string) {
	o.Language = language
}

// WithUserAgent adds the userAgent to the get universe groups group id params
func (o *GetUniverseGroupsGroupIDParams) WithUserAgent(userAgent *string) *GetUniverseGroupsGroupIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get universe groups group id params
func (o *GetUniverseGroupsGroupIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *GetUniverseGroupsGroupIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param group_id
	if err := r.SetPathParam("group_id", swag.FormatInt32(o.GroupID)); err != nil {
		return err
	}

	if o.Language != nil {

		// query param language
		var qrLanguage string
		if o.Language != nil {
			qrLanguage = *o.Language
		}
		qLanguage := qrLanguage
		if qLanguage != "" {
			if err := r.SetQueryParam("language", qLanguage); err != nil {
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
