package dummy

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

// NewGetCharactersCharacterIDWalletsJournalParams creates a new GetCharactersCharacterIDWalletsJournalParams object
// with the default values initialized.
func NewGetCharactersCharacterIDWalletsJournalParams() *GetCharactersCharacterIDWalletsJournalParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDWalletsJournalParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDWalletsJournalParamsWithTimeout creates a new GetCharactersCharacterIDWalletsJournalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCharactersCharacterIDWalletsJournalParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDWalletsJournalParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDWalletsJournalParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDWalletsJournalParamsWithContext creates a new GetCharactersCharacterIDWalletsJournalParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCharactersCharacterIDWalletsJournalParamsWithContext(ctx context.Context) *GetCharactersCharacterIDWalletsJournalParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDWalletsJournalParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

/*GetCharactersCharacterIDWalletsJournalParams contains all the parameters to send to the API endpoint
for the get characters character id wallets journal operation typically these are written to a http.Request
*/
type GetCharactersCharacterIDWalletsJournalParams struct {

	/*CharacterID
	  An EVE character ID

	*/
	CharacterID int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*LastSeenID
	  A journal reference ID to paginate from

	*/
	LastSeenID *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get characters character id wallets journal params
func (o *GetCharactersCharacterIDWalletsJournalParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDWalletsJournalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id wallets journal params
func (o *GetCharactersCharacterIDWalletsJournalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id wallets journal params
func (o *GetCharactersCharacterIDWalletsJournalParams) WithContext(ctx context.Context) *GetCharactersCharacterIDWalletsJournalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id wallets journal params
func (o *GetCharactersCharacterIDWalletsJournalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithCharacterID adds the characterID to the get characters character id wallets journal params
func (o *GetCharactersCharacterIDWalletsJournalParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDWalletsJournalParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id wallets journal params
func (o *GetCharactersCharacterIDWalletsJournalParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id wallets journal params
func (o *GetCharactersCharacterIDWalletsJournalParams) WithDatasource(datasource *string) *GetCharactersCharacterIDWalletsJournalParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id wallets journal params
func (o *GetCharactersCharacterIDWalletsJournalParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithLastSeenID adds the lastSeenID to the get characters character id wallets journal params
func (o *GetCharactersCharacterIDWalletsJournalParams) WithLastSeenID(lastSeenID *int64) *GetCharactersCharacterIDWalletsJournalParams {
	o.SetLastSeenID(lastSeenID)
	return o
}

// SetLastSeenID adds the lastSeenId to the get characters character id wallets journal params
func (o *GetCharactersCharacterIDWalletsJournalParams) SetLastSeenID(lastSeenID *int64) {
	o.LastSeenID = lastSeenID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDWalletsJournalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.LastSeenID != nil {

		// query param last_seen_id
		var qrLastSeenID int64
		if o.LastSeenID != nil {
			qrLastSeenID = *o.LastSeenID
		}
		qLastSeenID := swag.FormatInt64(qrLastSeenID)
		if qLastSeenID != "" {
			if err := r.SetQueryParam("last_seen_id", qLastSeenID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
