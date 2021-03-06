// Code generated by go-swagger; DO NOT EDIT.

package bookmarks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new bookmarks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for bookmarks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetCharactersCharacterIDBookmarks lists bookmarks

List your character's personal bookmarks

---
Alternate route: `/v1/characters/{character_id}/bookmarks/`

Alternate route: `/legacy/characters/{character_id}/bookmarks/`

Alternate route: `/dev/characters/{character_id}/bookmarks/`

---
This route is cached for up to 3600 seconds
*/
func (a *Client) GetCharactersCharacterIDBookmarks(params *GetCharactersCharacterIDBookmarksParams, authInfo runtime.ClientAuthInfoWriter) (*GetCharactersCharacterIDBookmarksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDBookmarksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_characters_character_id_bookmarks",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/bookmarks/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDBookmarksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCharactersCharacterIDBookmarksOK), nil

}

/*
GetCharactersCharacterIDBookmarksFolders lists bookmark folders

List your character's personal bookmark folders

---
Alternate route: `/v1/characters/{character_id}/bookmarks/folders/`

Alternate route: `/legacy/characters/{character_id}/bookmarks/folders/`

Alternate route: `/dev/characters/{character_id}/bookmarks/folders/`

---
This route is cached for up to 3600 seconds
*/
func (a *Client) GetCharactersCharacterIDBookmarksFolders(params *GetCharactersCharacterIDBookmarksFoldersParams, authInfo runtime.ClientAuthInfoWriter) (*GetCharactersCharacterIDBookmarksFoldersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDBookmarksFoldersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_characters_character_id_bookmarks_folders",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/bookmarks/folders/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDBookmarksFoldersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCharactersCharacterIDBookmarksFoldersOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
