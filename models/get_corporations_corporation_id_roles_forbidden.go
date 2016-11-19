package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// GetCorporationsCorporationIDRolesForbidden Forbidden
// swagger:model get_corporations_corporation_id_roles_forbidden
type GetCorporationsCorporationIDRolesForbidden struct {

	// get_corporations_corporation_id_roles_403_forbidden
	//
	// Forbidden message
	Error string `json:"error,omitempty"`
}

// Validate validates this get corporations corporation id roles forbidden
func (m *GetCorporationsCorporationIDRolesForbidden) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}