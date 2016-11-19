package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// GetCharactersCharacterIDOk 200 ok object
// swagger:model get_characters_character_id_ok
type GetCharactersCharacterIDOk struct {

	// get_characters_character_id_ancestry_id
	//
	// ancestry_id integer
	// Required: true
	AncestryID *int32 `json:"ancestry_id"`

	// get_characters_character_id_birthday
	//
	// Creation date of the character
	// Required: true
	Birthday *strfmt.DateTime `json:"birthday"`

	// get_characters_character_id_bloodline_id
	//
	// bloodline_id integer
	// Required: true
	BloodlineID *int32 `json:"bloodline_id"`

	// get_characters_character_id_corporation_id
	//
	// The character's corporation ID
	// Required: true
	CorporationID *int32 `json:"corporation_id"`

	// get_characters_character_id_description
	//
	// description string
	// Required: true
	Description *string `json:"description"`

	// get_characters_character_id_gender
	//
	// gender string
	// Required: true
	Gender *string `json:"gender"`

	// get_characters_character_id_name
	//
	// The name of the character
	// Required: true
	Name *string `json:"name"`

	// get_characters_character_id_race_id
	//
	// race_id integer
	// Required: true
	RaceID *int32 `json:"race_id"`

	// get_characters_character_id_security_status
	//
	// security_status number
	// Maximum: 10
	// Minimum: -10
	SecurityStatus *float32 `json:"security_status,omitempty"`
}

// Validate validates this get characters character id ok
func (m *GetCharactersCharacterIDOk) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAncestryID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBirthday(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBloodlineID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCorporationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGender(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRaceID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSecurityStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCharactersCharacterIDOk) validateAncestryID(formats strfmt.Registry) error {

	if err := validate.Required("ancestry_id", "body", m.AncestryID); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDOk) validateBirthday(formats strfmt.Registry) error {

	if err := validate.Required("birthday", "body", m.Birthday); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDOk) validateBloodlineID(formats strfmt.Registry) error {

	if err := validate.Required("bloodline_id", "body", m.BloodlineID); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDOk) validateCorporationID(formats strfmt.Registry) error {

	if err := validate.Required("corporation_id", "body", m.CorporationID); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDOk) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

var getCharactersCharacterIdOkTypeGenderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["female","male"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdOkTypeGenderPropEnum = append(getCharactersCharacterIdOkTypeGenderPropEnum, v)
	}
}

const (
	GetCharactersCharacterIDOkGenderFemale string = "female"
	GetCharactersCharacterIDOkGenderMale   string = "male"
)

// prop value enum
func (m *GetCharactersCharacterIDOk) validateGenderEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCharactersCharacterIdOkTypeGenderPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetCharactersCharacterIDOk) validateGender(formats strfmt.Registry) error {

	if err := validate.Required("gender", "body", m.Gender); err != nil {
		return err
	}

	// value enum
	if err := m.validateGenderEnum("gender", "body", *m.Gender); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDOk) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDOk) validateRaceID(formats strfmt.Registry) error {

	if err := validate.Required("race_id", "body", m.RaceID); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDOk) validateSecurityStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityStatus) { // not required
		return nil
	}

	if err := validate.Minimum("security_status", "body", float64(*m.SecurityStatus), -10, false); err != nil {
		return err
	}

	if err := validate.Maximum("security_status", "body", float64(*m.SecurityStatus), 10, false); err != nil {
		return err
	}

	return nil
}
