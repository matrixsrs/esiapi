package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// GetCharactersCharacterIDKillmailsRecent200Ok 200 ok object
// swagger:model get_characters_character_id_killmails_recent_200_ok
type GetCharactersCharacterIDKillmailsRecent200Ok struct {

	// get_characters_character_id_killmails_recent_killmail_hash
	//
	// A hash of this killmail
	// Required: true
	KillmailHash *string `json:"killmail_hash"`

	// get_characters_character_id_killmails_recent_killmail_id
	//
	// ID of this killmail
	// Required: true
	KillmailID *int32 `json:"killmail_id"`
}

// Validate validates this get characters character id killmails recent 200 ok
func (m *GetCharactersCharacterIDKillmailsRecent200Ok) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKillmailHash(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateKillmailID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCharactersCharacterIDKillmailsRecent200Ok) validateKillmailHash(formats strfmt.Registry) error {

	if err := validate.Required("killmail_hash", "body", m.KillmailHash); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDKillmailsRecent200Ok) validateKillmailID(formats strfmt.Registry) error {

	if err := validate.Required("killmail_id", "body", m.KillmailID); err != nil {
		return err
	}

	return nil
}