// Code generated by go-swagger; DO NOT EDIT.

package opportunities

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

// NewGetOpportunitiesTasksParams creates a new GetOpportunitiesTasksParams object
// with the default values initialized.
func NewGetOpportunitiesTasksParams() *GetOpportunitiesTasksParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetOpportunitiesTasksParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOpportunitiesTasksParamsWithTimeout creates a new GetOpportunitiesTasksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOpportunitiesTasksParamsWithTimeout(timeout time.Duration) *GetOpportunitiesTasksParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetOpportunitiesTasksParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetOpportunitiesTasksParamsWithContext creates a new GetOpportunitiesTasksParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOpportunitiesTasksParamsWithContext(ctx context.Context) *GetOpportunitiesTasksParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetOpportunitiesTasksParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetOpportunitiesTasksParamsWithHTTPClient creates a new GetOpportunitiesTasksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOpportunitiesTasksParamsWithHTTPClient(client *http.Client) *GetOpportunitiesTasksParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetOpportunitiesTasksParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetOpportunitiesTasksParams contains all the parameters to send to the API endpoint
for the get opportunities tasks operation typically these are written to a http.Request
*/
type GetOpportunitiesTasksParams struct {

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

// WithTimeout adds the timeout to the get opportunities tasks params
func (o *GetOpportunitiesTasksParams) WithTimeout(timeout time.Duration) *GetOpportunitiesTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get opportunities tasks params
func (o *GetOpportunitiesTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get opportunities tasks params
func (o *GetOpportunitiesTasksParams) WithContext(ctx context.Context) *GetOpportunitiesTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get opportunities tasks params
func (o *GetOpportunitiesTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get opportunities tasks params
func (o *GetOpportunitiesTasksParams) WithHTTPClient(client *http.Client) *GetOpportunitiesTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get opportunities tasks params
func (o *GetOpportunitiesTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the get opportunities tasks params
func (o *GetOpportunitiesTasksParams) WithXUserAgent(xUserAgent *string) *GetOpportunitiesTasksParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the get opportunities tasks params
func (o *GetOpportunitiesTasksParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithDatasource adds the datasource to the get opportunities tasks params
func (o *GetOpportunitiesTasksParams) WithDatasource(datasource *string) *GetOpportunitiesTasksParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get opportunities tasks params
func (o *GetOpportunitiesTasksParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithUserAgent adds the userAgent to the get opportunities tasks params
func (o *GetOpportunitiesTasksParams) WithUserAgent(userAgent *string) *GetOpportunitiesTasksParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get opportunities tasks params
func (o *GetOpportunitiesTasksParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *GetOpportunitiesTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
