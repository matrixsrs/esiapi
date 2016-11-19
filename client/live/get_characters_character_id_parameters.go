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

// NewGetCharactersCharacterIDParams creates a new GetCharactersCharacterIDParams object
// with the default values initialized.
func NewGetCharactersCharacterIDParams() *GetCharactersCharacterIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDParamsWithTimeout creates a new GetCharactersCharacterIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCharactersCharacterIDParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDParamsWithContext creates a new GetCharactersCharacterIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCharactersCharacterIDParamsWithContext(ctx context.Context) *GetCharactersCharacterIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

/*GetCharactersCharacterIDParams contains all the parameters to send to the API endpoint
for the get characters character id operation typically these are written to a http.Request
*/
type GetCharactersCharacterIDParams struct {

	/*CharacterID
	  An EVE character ID

	*/
	CharacterID int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get characters character id params
func (o *GetCharactersCharacterIDParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id params
func (o *GetCharactersCharacterIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id params
func (o *GetCharactersCharacterIDParams) WithContext(ctx context.Context) *GetCharactersCharacterIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id params
func (o *GetCharactersCharacterIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithCharacterID adds the characterID to the get characters character id params
func (o *GetCharactersCharacterIDParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id params
func (o *GetCharactersCharacterIDParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id params
func (o *GetCharactersCharacterIDParams) WithDatasource(datasource *string) *GetCharactersCharacterIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id params
func (o *GetCharactersCharacterIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param character_id
	if err := r.SetPathParam("character_id", swag.FormatInt32(o.CharacterID)); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
