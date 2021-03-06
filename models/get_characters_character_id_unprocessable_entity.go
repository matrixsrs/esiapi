package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// GetCharactersCharacterIDUnprocessableEntity Is not a character ID
// swagger:model get_characters_character_id_unprocessable_entity
type GetCharactersCharacterIDUnprocessableEntity struct {

	// get_characters_character_id_error
	//
	// error message
	Error string `json:"error,omitempty"`
}

// Validate validates this get characters character id unprocessable entity
func (m *GetCharactersCharacterIDUnprocessableEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
