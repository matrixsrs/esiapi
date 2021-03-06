package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// GetCharactersCharacterIDClonesForbidden Forbidden
// swagger:model get_characters_character_id_clones_forbidden
type GetCharactersCharacterIDClonesForbidden struct {

	// get_characters_character_id_clones_403_forbidden
	//
	// Forbidden message
	Error string `json:"error,omitempty"`
}

// Validate validates this get characters character id clones forbidden
func (m *GetCharactersCharacterIDClonesForbidden) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
