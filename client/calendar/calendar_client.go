// Code generated by go-swagger; DO NOT EDIT.

package calendar

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new calendar API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for calendar API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetCharactersCharacterIDCalendar lists calendar event summaries

Get 50 event summaries from the calendar. If no event ID is given, the resource will return the next 50 chronological event summaries from now. If an event ID is specified, it will return the next 50 chronological event summaries from after that event.

---
Alternate route: `/v1/characters/{character_id}/calendar/`

Alternate route: `/legacy/characters/{character_id}/calendar/`

Alternate route: `/dev/characters/{character_id}/calendar/`

---
This route is cached for up to 5 seconds
*/
func (a *Client) GetCharactersCharacterIDCalendar(params *GetCharactersCharacterIDCalendarParams, authInfo runtime.ClientAuthInfoWriter) (*GetCharactersCharacterIDCalendarOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDCalendarParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_characters_character_id_calendar",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/calendar/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDCalendarReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCharactersCharacterIDCalendarOK), nil

}

/*
GetCharactersCharacterIDCalendarEventID gets an event

Get all the information for a specific event

---
Alternate route: `/v3/characters/{character_id}/calendar/{event_id}/`

Alternate route: `/dev/characters/{character_id}/calendar/{event_id}/`

---
This route is cached for up to 5 seconds
*/
func (a *Client) GetCharactersCharacterIDCalendarEventID(params *GetCharactersCharacterIDCalendarEventIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetCharactersCharacterIDCalendarEventIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDCalendarEventIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_characters_character_id_calendar_event_id",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/calendar/{event_id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDCalendarEventIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCharactersCharacterIDCalendarEventIDOK), nil

}

/*
PutCharactersCharacterIDCalendarEventID responds to an event

Set your response status to an event

---
Alternate route: `/v3/characters/{character_id}/calendar/{event_id}/`

Alternate route: `/dev/characters/{character_id}/calendar/{event_id}/`

*/
func (a *Client) PutCharactersCharacterIDCalendarEventID(params *PutCharactersCharacterIDCalendarEventIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutCharactersCharacterIDCalendarEventIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutCharactersCharacterIDCalendarEventIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "put_characters_character_id_calendar_event_id",
		Method:             "PUT",
		PathPattern:        "/characters/{character_id}/calendar/{event_id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutCharactersCharacterIDCalendarEventIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutCharactersCharacterIDCalendarEventIDNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
