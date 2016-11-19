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

// NewGetCharactersCharacterIDBookmarksFoldersParams creates a new GetCharactersCharacterIDBookmarksFoldersParams object
// with the default values initialized.
func NewGetCharactersCharacterIDBookmarksFoldersParams() *GetCharactersCharacterIDBookmarksFoldersParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDBookmarksFoldersParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDBookmarksFoldersParamsWithTimeout creates a new GetCharactersCharacterIDBookmarksFoldersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCharactersCharacterIDBookmarksFoldersParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDBookmarksFoldersParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDBookmarksFoldersParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDBookmarksFoldersParamsWithContext creates a new GetCharactersCharacterIDBookmarksFoldersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCharactersCharacterIDBookmarksFoldersParamsWithContext(ctx context.Context) *GetCharactersCharacterIDBookmarksFoldersParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDBookmarksFoldersParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

/*GetCharactersCharacterIDBookmarksFoldersParams contains all the parameters to send to the API endpoint
for the get characters character id bookmarks folders operation typically these are written to a http.Request
*/
type GetCharactersCharacterIDBookmarksFoldersParams struct {

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

// WithTimeout adds the timeout to the get characters character id bookmarks folders params
func (o *GetCharactersCharacterIDBookmarksFoldersParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDBookmarksFoldersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id bookmarks folders params
func (o *GetCharactersCharacterIDBookmarksFoldersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id bookmarks folders params
func (o *GetCharactersCharacterIDBookmarksFoldersParams) WithContext(ctx context.Context) *GetCharactersCharacterIDBookmarksFoldersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id bookmarks folders params
func (o *GetCharactersCharacterIDBookmarksFoldersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithCharacterID adds the characterID to the get characters character id bookmarks folders params
func (o *GetCharactersCharacterIDBookmarksFoldersParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDBookmarksFoldersParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id bookmarks folders params
func (o *GetCharactersCharacterIDBookmarksFoldersParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id bookmarks folders params
func (o *GetCharactersCharacterIDBookmarksFoldersParams) WithDatasource(datasource *string) *GetCharactersCharacterIDBookmarksFoldersParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id bookmarks folders params
func (o *GetCharactersCharacterIDBookmarksFoldersParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDBookmarksFoldersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
