package killmails

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

// NewGetKillmailsKillmailIDKillmailHashParams creates a new GetKillmailsKillmailIDKillmailHashParams object
// with the default values initialized.
func NewGetKillmailsKillmailIDKillmailHashParams() *GetKillmailsKillmailIDKillmailHashParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetKillmailsKillmailIDKillmailHashParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKillmailsKillmailIDKillmailHashParamsWithTimeout creates a new GetKillmailsKillmailIDKillmailHashParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKillmailsKillmailIDKillmailHashParamsWithTimeout(timeout time.Duration) *GetKillmailsKillmailIDKillmailHashParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetKillmailsKillmailIDKillmailHashParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetKillmailsKillmailIDKillmailHashParamsWithContext creates a new GetKillmailsKillmailIDKillmailHashParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKillmailsKillmailIDKillmailHashParamsWithContext(ctx context.Context) *GetKillmailsKillmailIDKillmailHashParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetKillmailsKillmailIDKillmailHashParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

/*GetKillmailsKillmailIDKillmailHashParams contains all the parameters to send to the API endpoint
for the get killmails killmail id killmail hash operation typically these are written to a http.Request
*/
type GetKillmailsKillmailIDKillmailHashParams struct {

	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*KillmailHash
	  The killmail hash for verification

	*/
	KillmailHash string
	/*KillmailID
	  The killmail ID to be queried

	*/
	KillmailID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get killmails killmail id killmail hash params
func (o *GetKillmailsKillmailIDKillmailHashParams) WithTimeout(timeout time.Duration) *GetKillmailsKillmailIDKillmailHashParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get killmails killmail id killmail hash params
func (o *GetKillmailsKillmailIDKillmailHashParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get killmails killmail id killmail hash params
func (o *GetKillmailsKillmailIDKillmailHashParams) WithContext(ctx context.Context) *GetKillmailsKillmailIDKillmailHashParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get killmails killmail id killmail hash params
func (o *GetKillmailsKillmailIDKillmailHashParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithDatasource adds the datasource to the get killmails killmail id killmail hash params
func (o *GetKillmailsKillmailIDKillmailHashParams) WithDatasource(datasource *string) *GetKillmailsKillmailIDKillmailHashParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get killmails killmail id killmail hash params
func (o *GetKillmailsKillmailIDKillmailHashParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithKillmailHash adds the killmailHash to the get killmails killmail id killmail hash params
func (o *GetKillmailsKillmailIDKillmailHashParams) WithKillmailHash(killmailHash string) *GetKillmailsKillmailIDKillmailHashParams {
	o.SetKillmailHash(killmailHash)
	return o
}

// SetKillmailHash adds the killmailHash to the get killmails killmail id killmail hash params
func (o *GetKillmailsKillmailIDKillmailHashParams) SetKillmailHash(killmailHash string) {
	o.KillmailHash = killmailHash
}

// WithKillmailID adds the killmailID to the get killmails killmail id killmail hash params
func (o *GetKillmailsKillmailIDKillmailHashParams) WithKillmailID(killmailID int32) *GetKillmailsKillmailIDKillmailHashParams {
	o.SetKillmailID(killmailID)
	return o
}

// SetKillmailID adds the killmailId to the get killmails killmail id killmail hash params
func (o *GetKillmailsKillmailIDKillmailHashParams) SetKillmailID(killmailID int32) {
	o.KillmailID = killmailID
}

// WriteToRequest writes these params to a swagger request
func (o *GetKillmailsKillmailIDKillmailHashParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param killmail_hash
	if err := r.SetPathParam("killmail_hash", o.KillmailHash); err != nil {
		return err
	}

	// path param killmail_id
	if err := r.SetPathParam("killmail_id", swag.FormatInt32(o.KillmailID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
