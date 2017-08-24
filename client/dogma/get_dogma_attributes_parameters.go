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

// NewGetDogmaAttributesParams creates a new GetDogmaAttributesParams object
// with the default values initialized.
func NewGetDogmaAttributesParams() *GetDogmaAttributesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetDogmaAttributesParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDogmaAttributesParamsWithTimeout creates a new GetDogmaAttributesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDogmaAttributesParamsWithTimeout(timeout time.Duration) *GetDogmaAttributesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetDogmaAttributesParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetDogmaAttributesParamsWithContext creates a new GetDogmaAttributesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDogmaAttributesParamsWithContext(ctx context.Context) *GetDogmaAttributesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetDogmaAttributesParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetDogmaAttributesParamsWithHTTPClient creates a new GetDogmaAttributesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDogmaAttributesParamsWithHTTPClient(client *http.Client) *GetDogmaAttributesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetDogmaAttributesParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetDogmaAttributesParams contains all the parameters to send to the API endpoint
for the get dogma attributes operation typically these are written to a http.Request
*/
type GetDogmaAttributesParams struct {

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

// WithTimeout adds the timeout to the get dogma attributes params
func (o *GetDogmaAttributesParams) WithTimeout(timeout time.Duration) *GetDogmaAttributesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dogma attributes params
func (o *GetDogmaAttributesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dogma attributes params
func (o *GetDogmaAttributesParams) WithContext(ctx context.Context) *GetDogmaAttributesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dogma attributes params
func (o *GetDogmaAttributesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dogma attributes params
func (o *GetDogmaAttributesParams) WithHTTPClient(client *http.Client) *GetDogmaAttributesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dogma attributes params
func (o *GetDogmaAttributesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the get dogma attributes params
func (o *GetDogmaAttributesParams) WithXUserAgent(xUserAgent *string) *GetDogmaAttributesParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the get dogma attributes params
func (o *GetDogmaAttributesParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithDatasource adds the datasource to the get dogma attributes params
func (o *GetDogmaAttributesParams) WithDatasource(datasource *string) *GetDogmaAttributesParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get dogma attributes params
func (o *GetDogmaAttributesParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithUserAgent adds the userAgent to the get dogma attributes params
func (o *GetDogmaAttributesParams) WithUserAgent(userAgent *string) *GetDogmaAttributesParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get dogma attributes params
func (o *GetDogmaAttributesParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *GetDogmaAttributesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
