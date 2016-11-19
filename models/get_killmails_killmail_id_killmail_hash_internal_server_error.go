package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// GetKillmailsKillmailIDKillmailHashInternalServerError Internal server error
// swagger:model get_killmails_killmail_id_killmail_hash_internal_server_error
type GetKillmailsKillmailIDKillmailHashInternalServerError struct {

	// get_killmails_killmail_id_killmail_hash_500_internal_server_error
	//
	// Internal server error message
	Error string `json:"error,omitempty"`
}

// Validate validates this get killmails killmail id killmail hash internal server error
func (m *GetKillmailsKillmailIDKillmailHashInternalServerError) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}