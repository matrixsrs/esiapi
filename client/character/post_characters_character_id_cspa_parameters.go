package character

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/models"
)

// NewPostCharactersCharacterIDCspaParams creates a new PostCharactersCharacterIDCspaParams object
// with the default values initialized.
func NewPostCharactersCharacterIDCspaParams() *PostCharactersCharacterIDCspaParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostCharactersCharacterIDCspaParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCharactersCharacterIDCspaParamsWithTimeout creates a new PostCharactersCharacterIDCspaParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCharactersCharacterIDCspaParamsWithTimeout(timeout time.Duration) *PostCharactersCharacterIDCspaParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostCharactersCharacterIDCspaParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewPostCharactersCharacterIDCspaParamsWithContext creates a new PostCharactersCharacterIDCspaParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCharactersCharacterIDCspaParamsWithContext(ctx context.Context) *PostCharactersCharacterIDCspaParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostCharactersCharacterIDCspaParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

/*PostCharactersCharacterIDCspaParams contains all the parameters to send to the API endpoint
for the post characters character id cspa operation typically these are written to a http.Request
*/
type PostCharactersCharacterIDCspaParams struct {

	/*CharacterID
	  An EVE character ID

	*/
	CharacterID int32
	/*Characters
	  The target characters to calculate the charge for

	*/
	Characters *models.PostCharactersCharacterIDCspaCharacters
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post characters character id cspa params
func (o *PostCharactersCharacterIDCspaParams) WithTimeout(timeout time.Duration) *PostCharactersCharacterIDCspaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post characters character id cspa params
func (o *PostCharactersCharacterIDCspaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post characters character id cspa params
func (o *PostCharactersCharacterIDCspaParams) WithContext(ctx context.Context) *PostCharactersCharacterIDCspaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post characters character id cspa params
func (o *PostCharactersCharacterIDCspaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithCharacterID adds the characterID to the post characters character id cspa params
func (o *PostCharactersCharacterIDCspaParams) WithCharacterID(characterID int32) *PostCharactersCharacterIDCspaParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the post characters character id cspa params
func (o *PostCharactersCharacterIDCspaParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithCharacters adds the characters to the post characters character id cspa params
func (o *PostCharactersCharacterIDCspaParams) WithCharacters(characters *models.PostCharactersCharacterIDCspaCharacters) *PostCharactersCharacterIDCspaParams {
	o.SetCharacters(characters)
	return o
}

// SetCharacters adds the characters to the post characters character id cspa params
func (o *PostCharactersCharacterIDCspaParams) SetCharacters(characters *models.PostCharactersCharacterIDCspaCharacters) {
	o.Characters = characters
}

// WithDatasource adds the datasource to the post characters character id cspa params
func (o *PostCharactersCharacterIDCspaParams) WithDatasource(datasource *string) *PostCharactersCharacterIDCspaParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the post characters character id cspa params
func (o *PostCharactersCharacterIDCspaParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *PostCharactersCharacterIDCspaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param character_id
	if err := r.SetPathParam("character_id", swag.FormatInt32(o.CharacterID)); err != nil {
		return err
	}

	if o.Characters == nil {
		o.Characters = new(models.PostCharactersCharacterIDCspaCharacters)
	}

	if err := r.SetBodyParam(o.Characters); err != nil {
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