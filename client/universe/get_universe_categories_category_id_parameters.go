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

// NewGetUniverseCategoriesCategoryIDParams creates a new GetUniverseCategoriesCategoryIDParams object
// with the default values initialized.
func NewGetUniverseCategoriesCategoryIDParams() *GetUniverseCategoriesCategoryIDParams {
	var (
		datasourceDefault = string("tranquility")
		languageDefault   = string("en-us")
	)
	return &GetUniverseCategoriesCategoryIDParams{
		Datasource: &datasourceDefault,
		Language:   &languageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUniverseCategoriesCategoryIDParamsWithTimeout creates a new GetUniverseCategoriesCategoryIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUniverseCategoriesCategoryIDParamsWithTimeout(timeout time.Duration) *GetUniverseCategoriesCategoryIDParams {
	var (
		datasourceDefault = string("tranquility")
		languageDefault   = string("en-us")
	)
	return &GetUniverseCategoriesCategoryIDParams{
		Datasource: &datasourceDefault,
		Language:   &languageDefault,

		timeout: timeout,
	}
}

// NewGetUniverseCategoriesCategoryIDParamsWithContext creates a new GetUniverseCategoriesCategoryIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUniverseCategoriesCategoryIDParamsWithContext(ctx context.Context) *GetUniverseCategoriesCategoryIDParams {
	var (
		datasourceDefault = string("tranquility")
		languageDefault   = string("en-us")
	)
	return &GetUniverseCategoriesCategoryIDParams{
		Datasource: &datasourceDefault,
		Language:   &languageDefault,

		Context: ctx,
	}
}

// NewGetUniverseCategoriesCategoryIDParamsWithHTTPClient creates a new GetUniverseCategoriesCategoryIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUniverseCategoriesCategoryIDParamsWithHTTPClient(client *http.Client) *GetUniverseCategoriesCategoryIDParams {
	var (
		datasourceDefault = string("tranquility")
		languageDefault   = string("en-us")
	)
	return &GetUniverseCategoriesCategoryIDParams{
		Datasource: &datasourceDefault,
		Language:   &languageDefault,
		HTTPClient: client,
	}
}

/*GetUniverseCategoriesCategoryIDParams contains all the parameters to send to the API endpoint
for the get universe categories category id operation typically these are written to a http.Request
*/
type GetUniverseCategoriesCategoryIDParams struct {

	/*XUserAgent
	  Client identifier, takes precedence over User-Agent

	*/
	XUserAgent *string
	/*CategoryID
	  An Eve item category ID

	*/
	CategoryID int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
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

// WithTimeout adds the timeout to the get universe categories category id params
func (o *GetUniverseCategoriesCategoryIDParams) WithTimeout(timeout time.Duration) *GetUniverseCategoriesCategoryIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get universe categories category id params
func (o *GetUniverseCategoriesCategoryIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get universe categories category id params
func (o *GetUniverseCategoriesCategoryIDParams) WithContext(ctx context.Context) *GetUniverseCategoriesCategoryIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get universe categories category id params
func (o *GetUniverseCategoriesCategoryIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get universe categories category id params
func (o *GetUniverseCategoriesCategoryIDParams) WithHTTPClient(client *http.Client) *GetUniverseCategoriesCategoryIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get universe categories category id params
func (o *GetUniverseCategoriesCategoryIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the get universe categories category id params
func (o *GetUniverseCategoriesCategoryIDParams) WithXUserAgent(xUserAgent *string) *GetUniverseCategoriesCategoryIDParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the get universe categories category id params
func (o *GetUniverseCategoriesCategoryIDParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithCategoryID adds the categoryID to the get universe categories category id params
func (o *GetUniverseCategoriesCategoryIDParams) WithCategoryID(categoryID int32) *GetUniverseCategoriesCategoryIDParams {
	o.SetCategoryID(categoryID)
	return o
}

// SetCategoryID adds the categoryId to the get universe categories category id params
func (o *GetUniverseCategoriesCategoryIDParams) SetCategoryID(categoryID int32) {
	o.CategoryID = categoryID
}

// WithDatasource adds the datasource to the get universe categories category id params
func (o *GetUniverseCategoriesCategoryIDParams) WithDatasource(datasource *string) *GetUniverseCategoriesCategoryIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get universe categories category id params
func (o *GetUniverseCategoriesCategoryIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithLanguage adds the language to the get universe categories category id params
func (o *GetUniverseCategoriesCategoryIDParams) WithLanguage(language *string) *GetUniverseCategoriesCategoryIDParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the get universe categories category id params
func (o *GetUniverseCategoriesCategoryIDParams) SetLanguage(language *string) {
	o.Language = language
}

// WithUserAgent adds the userAgent to the get universe categories category id params
func (o *GetUniverseCategoriesCategoryIDParams) WithUserAgent(userAgent *string) *GetUniverseCategoriesCategoryIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get universe categories category id params
func (o *GetUniverseCategoriesCategoryIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *GetUniverseCategoriesCategoryIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param category_id
	if err := r.SetPathParam("category_id", swag.FormatInt32(o.CategoryID)); err != nil {
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
