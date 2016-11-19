package live

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

// NewGetUniverseStructuresParams creates a new GetUniverseStructuresParams object
// with the default values initialized.
func NewGetUniverseStructuresParams() *GetUniverseStructuresParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseStructuresParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUniverseStructuresParamsWithTimeout creates a new GetUniverseStructuresParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUniverseStructuresParamsWithTimeout(timeout time.Duration) *GetUniverseStructuresParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseStructuresParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetUniverseStructuresParamsWithContext creates a new GetUniverseStructuresParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUniverseStructuresParamsWithContext(ctx context.Context) *GetUniverseStructuresParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseStructuresParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

/*GetUniverseStructuresParams contains all the parameters to send to the API endpoint
for the get universe structures operation typically these are written to a http.Request
*/
type GetUniverseStructuresParams struct {

	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get universe structures params
func (o *GetUniverseStructuresParams) WithTimeout(timeout time.Duration) *GetUniverseStructuresParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get universe structures params
func (o *GetUniverseStructuresParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get universe structures params
func (o *GetUniverseStructuresParams) WithContext(ctx context.Context) *GetUniverseStructuresParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get universe structures params
func (o *GetUniverseStructuresParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithDatasource adds the datasource to the get universe structures params
func (o *GetUniverseStructuresParams) WithDatasource(datasource *string) *GetUniverseStructuresParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get universe structures params
func (o *GetUniverseStructuresParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *GetUniverseStructuresParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
