package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// GetCharactersCharacterIDSkillqueueInternalServerError Internal server error
// swagger:model get_characters_character_id_skillqueue_internal_server_error
type GetCharactersCharacterIDSkillqueueInternalServerError struct {

	// get_characters_character_id_skillqueue_500_internal_server_error
	//
	// Internal server error message
	Error string `json:"error,omitempty"`
}

// Validate validates this get characters character id skillqueue internal server error
func (m *GetCharactersCharacterIDSkillqueueInternalServerError) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
