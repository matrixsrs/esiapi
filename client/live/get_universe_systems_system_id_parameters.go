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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetUniverseSystemsSystemIDParams creates a new GetUniverseSystemsSystemIDParams object
// with the default values initialized.
func NewGetUniverseSystemsSystemIDParams() *GetUniverseSystemsSystemIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseSystemsSystemIDParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUniverseSystemsSystemIDParamsWithTimeout creates a new GetUniverseSystemsSystemIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUniverseSystemsSystemIDParamsWithTimeout(timeout time.Duration) *GetUniverseSystemsSystemIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseSystemsSystemIDParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetUniverseSystemsSystemIDParamsWithContext creates a new GetUniverseSystemsSystemIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUniverseSystemsSystemIDParamsWithContext(ctx context.Context) *GetUniverseSystemsSystemIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseSystemsSystemIDParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

/*GetUniverseSystemsSystemIDParams contains all the parameters to send to the API endpoint
for the get universe systems system id operation typically these are written to a http.Request
*/
type GetUniverseSystemsSystemIDParams struct {

	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*SystemID
	  An Eve solar system ID

	*/
	SystemID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get universe systems system id params
func (o *GetUniverseSystemsSystemIDParams) WithTimeout(timeout time.Duration) *GetUniverseSystemsSystemIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get universe systems system id params
func (o *GetUniverseSystemsSystemIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get universe systems system id params
func (o *GetUniverseSystemsSystemIDParams) WithContext(ctx context.Context) *GetUniverseSystemsSystemIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get universe systems system id params
func (o *GetUniverseSystemsSystemIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithDatasource adds the datasource to the get universe systems system id params
func (o *GetUniverseSystemsSystemIDParams) WithDatasource(datasource *string) *GetUniverseSystemsSystemIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get universe systems system id params
func (o *GetUniverseSystemsSystemIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithSystemID adds the systemID to the get universe systems system id params
func (o *GetUniverseSystemsSystemIDParams) WithSystemID(systemID int32) *GetUniverseSystemsSystemIDParams {
	o.SetSystemID(systemID)
	return o
}

// SetSystemID adds the systemId to the get universe systems system id params
func (o *GetUniverseSystemsSystemIDParams) SetSystemID(systemID int32) {
	o.SystemID = systemID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUniverseSystemsSystemIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param system_id
	if err := r.SetPathParam("system_id", swag.FormatInt32(o.SystemID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
