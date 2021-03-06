package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// GetMarketsRegionIDHistoryUnprocessableEntity bad region_id
// swagger:model get_markets_region_id_history_unprocessable_entity
type GetMarketsRegionIDHistoryUnprocessableEntity struct {

	// get_markets_region_id_history_error
	//
	// error message
	Error string `json:"error,omitempty"`
}

// Validate validates this get markets region id history unprocessable entity
func (m *GetMarketsRegionIDHistoryUnprocessableEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
