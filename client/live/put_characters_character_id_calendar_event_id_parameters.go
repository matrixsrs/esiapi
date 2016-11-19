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

	"github.com/evecentral/esiapi/models"
)

// NewPutCharactersCharacterIDCalendarEventIDParams creates a new PutCharactersCharacterIDCalendarEventIDParams object
// with the default values initialized.
func NewPutCharactersCharacterIDCalendarEventIDParams() *PutCharactersCharacterIDCalendarEventIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PutCharactersCharacterIDCalendarEventIDParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCharactersCharacterIDCalendarEventIDParamsWithTimeout creates a new PutCharactersCharacterIDCalendarEventIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCharactersCharacterIDCalendarEventIDParamsWithTimeout(timeout time.Duration) *PutCharactersCharacterIDCalendarEventIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PutCharactersCharacterIDCalendarEventIDParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewPutCharactersCharacterIDCalendarEventIDParamsWithContext creates a new PutCharactersCharacterIDCalendarEventIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCharactersCharacterIDCalendarEventIDParamsWithContext(ctx context.Context) *PutCharactersCharacterIDCalendarEventIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PutCharactersCharacterIDCalendarEventIDParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

/*PutCharactersCharacterIDCalendarEventIDParams contains all the parameters to send to the API endpoint
for the put characters character id calendar event id operation typically these are written to a http.Request
*/
type PutCharactersCharacterIDCalendarEventIDParams struct {

	/*CharacterID
	  The character ID requesting the event

	*/
	CharacterID int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*EventID
	  The ID of the event requested

	*/
	EventID int32
	/*Response
	  The response value to set, overriding current value.

	*/
	Response *models.PutCharactersCharacterIDCalendarEventIDResponse

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put characters character id calendar event id params
func (o *PutCharactersCharacterIDCalendarEventIDParams) WithTimeout(timeout time.Duration) *PutCharactersCharacterIDCalendarEventIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put characters character id calendar event id params
func (o *PutCharactersCharacterIDCalendarEventIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put characters character id calendar event id params
func (o *PutCharactersCharacterIDCalendarEventIDParams) WithContext(ctx context.Context) *PutCharactersCharacterIDCalendarEventIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put characters character id calendar event id params
func (o *PutCharactersCharacterIDCalendarEventIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithCharacterID adds the characterID to the put characters character id calendar event id params
func (o *PutCharactersCharacterIDCalendarEventIDParams) WithCharacterID(characterID int32) *PutCharactersCharacterIDCalendarEventIDParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the put characters character id calendar event id params
func (o *PutCharactersCharacterIDCalendarEventIDParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the put characters character id calendar event id params
func (o *PutCharactersCharacterIDCalendarEventIDParams) WithDatasource(datasource *string) *PutCharactersCharacterIDCalendarEventIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the put characters character id calendar event id params
func (o *PutCharactersCharacterIDCalendarEventIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithEventID adds the eventID to the put characters character id calendar event id params
func (o *PutCharactersCharacterIDCalendarEventIDParams) WithEventID(eventID int32) *PutCharactersCharacterIDCalendarEventIDParams {
	o.SetEventID(eventID)
	return o
}

// SetEventID adds the eventId to the put characters character id calendar event id params
func (o *PutCharactersCharacterIDCalendarEventIDParams) SetEventID(eventID int32) {
	o.EventID = eventID
}

// WithResponse adds the response to the put characters character id calendar event id params
func (o *PutCharactersCharacterIDCalendarEventIDParams) WithResponse(response *models.PutCharactersCharacterIDCalendarEventIDResponse) *PutCharactersCharacterIDCalendarEventIDParams {
	o.SetResponse(response)
	return o
}

// SetResponse adds the response to the put characters character id calendar event id params
func (o *PutCharactersCharacterIDCalendarEventIDParams) SetResponse(response *models.PutCharactersCharacterIDCalendarEventIDResponse) {
	o.Response = response
}

// WriteToRequest writes these params to a swagger request
func (o *PutCharactersCharacterIDCalendarEventIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param event_id
	if err := r.SetPathParam("event_id", swag.FormatInt32(o.EventID)); err != nil {
		return err
	}

	if o.Response == nil {
		o.Response = new(models.PutCharactersCharacterIDCalendarEventIDResponse)
	}

	if err := r.SetBodyParam(o.Response); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
