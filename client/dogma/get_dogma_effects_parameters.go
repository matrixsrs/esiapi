// Code generated by go-swagger; DO NOT EDIT.

package dogma

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDogmaEffectsParams creates a new GetDogmaEffectsParams object
// with the default values initialized.
func NewGetDogmaEffectsParams() *GetDogmaEffectsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetDogmaEffectsParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDogmaEffectsParamsWithTimeout creates a new GetDogmaEffectsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDogmaEffectsParamsWithTimeout(timeout time.Duration) *GetDogmaEffectsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetDogmaEffectsParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetDogmaEffectsParamsWithContext creates a new GetDogmaEffectsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDogmaEffectsParamsWithContext(ctx context.Context) *GetDogmaEffectsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetDogmaEffectsParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetDogmaEffectsParamsWithHTTPClient creates a new GetDogmaEffectsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDogmaEffectsParamsWithHTTPClient(client *http.Client) *GetDogmaEffectsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetDogmaEffectsParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetDogmaEffectsParams contains all the parameters to send to the API endpoint
for the get dogma effects operation typically these are written to a http.Request
*/
type GetDogmaEffectsParams struct {

	/*XUserAgent
	  Client identifier, takes precedence over User-Agent

	*/
	XUserAgent *string
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

// WithTimeout adds the timeout to the get dogma effects params
func (o *GetDogmaEffectsParams) WithTimeout(timeout time.Duration) *GetDogmaEffectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dogma effects params
func (o *GetDogmaEffectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dogma effects params
func (o *GetDogmaEffectsParams) WithContext(ctx context.Context) *GetDogmaEffectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dogma effects params
func (o *GetDogmaEffectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dogma effects params
func (o *GetDogmaEffectsParams) WithHTTPClient(client *http.Client) *GetDogmaEffectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dogma effects params
func (o *GetDogmaEffectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the get dogma effects params
func (o *GetDogmaEffectsParams) WithXUserAgent(xUserAgent *string) *GetDogmaEffectsParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the get dogma effects params
func (o *GetDogmaEffectsParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithDatasource adds the datasource to the get dogma effects params
func (o *GetDogmaEffectsParams) WithDatasource(datasource *string) *GetDogmaEffectsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get dogma effects params
func (o *GetDogmaEffectsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithUserAgent adds the userAgent to the get dogma effects params
func (o *GetDogmaEffectsParams) WithUserAgent(userAgent *string) *GetDogmaEffectsParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get dogma effects params
func (o *GetDogmaEffectsParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *GetDogmaEffectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
