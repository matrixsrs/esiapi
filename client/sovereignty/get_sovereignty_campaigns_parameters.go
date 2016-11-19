package sovereignty

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

// NewGetSovereigntyCampaignsParams creates a new GetSovereigntyCampaignsParams object
// with the default values initialized.
func NewGetSovereigntyCampaignsParams() *GetSovereigntyCampaignsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetSovereigntyCampaignsParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSovereigntyCampaignsParamsWithTimeout creates a new GetSovereigntyCampaignsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSovereigntyCampaignsParamsWithTimeout(timeout time.Duration) *GetSovereigntyCampaignsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetSovereigntyCampaignsParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetSovereigntyCampaignsParamsWithContext creates a new GetSovereigntyCampaignsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSovereigntyCampaignsParamsWithContext(ctx context.Context) *GetSovereigntyCampaignsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetSovereigntyCampaignsParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

/*GetSovereigntyCampaignsParams contains all the parameters to send to the API endpoint
for the get sovereignty campaigns operation typically these are written to a http.Request
*/
type GetSovereigntyCampaignsParams struct {

	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sovereignty campaigns params
func (o *GetSovereigntyCampaignsParams) WithTimeout(timeout time.Duration) *GetSovereigntyCampaignsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sovereignty campaigns params
func (o *GetSovereigntyCampaignsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sovereignty campaigns params
func (o *GetSovereigntyCampaignsParams) WithContext(ctx context.Context) *GetSovereigntyCampaignsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sovereignty campaigns params
func (o *GetSovereigntyCampaignsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithDatasource adds the datasource to the get sovereignty campaigns params
func (o *GetSovereigntyCampaignsParams) WithDatasource(datasource *string) *GetSovereigntyCampaignsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get sovereignty campaigns params
func (o *GetSovereigntyCampaignsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *GetSovereigntyCampaignsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
