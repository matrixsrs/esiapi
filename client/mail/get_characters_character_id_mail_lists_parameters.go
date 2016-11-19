package mail

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

// NewGetCharactersCharacterIDMailListsParams creates a new GetCharactersCharacterIDMailListsParams object
// with the default values initialized.
func NewGetCharactersCharacterIDMailListsParams() *GetCharactersCharacterIDMailListsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDMailListsParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDMailListsParamsWithTimeout creates a new GetCharactersCharacterIDMailListsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCharactersCharacterIDMailListsParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDMailListsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDMailListsParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDMailListsParamsWithContext creates a new GetCharactersCharacterIDMailListsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCharactersCharacterIDMailListsParamsWithContext(ctx context.Context) *GetCharactersCharacterIDMailListsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDMailListsParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

/*GetCharactersCharacterIDMailListsParams contains all the parameters to send to the API endpoint
for the get characters character id mail lists operation typically these are written to a http.Request
*/
type GetCharactersCharacterIDMailListsParams struct {

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

// WithTimeout adds the timeout to the get characters character id mail lists params
func (o *GetCharactersCharacterIDMailListsParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDMailListsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id mail lists params
func (o *GetCharactersCharacterIDMailListsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id mail lists params
func (o *GetCharactersCharacterIDMailListsParams) WithContext(ctx context.Context) *GetCharactersCharacterIDMailListsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id mail lists params
func (o *GetCharactersCharacterIDMailListsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithCharacterID adds the characterID to the get characters character id mail lists params
func (o *GetCharactersCharacterIDMailListsParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDMailListsParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id mail lists params
func (o *GetCharactersCharacterIDMailListsParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id mail lists params
func (o *GetCharactersCharacterIDMailListsParams) WithDatasource(datasource *string) *GetCharactersCharacterIDMailListsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id mail lists params
func (o *GetCharactersCharacterIDMailListsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDMailListsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
