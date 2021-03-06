package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// PostCharactersCharacterIDMailRecipient recipient object
// swagger:model post_characters_character_id_mail_recipient
type PostCharactersCharacterIDMailRecipient struct {

	// post_characters_character_id_mail_recipient_id
	//
	// recipient_id integer
	// Required: true
	RecipientID *int32 `json:"recipient_id"`

	// post_characters_character_id_mail_recipient_type
	//
	// recipient_type string
	// Required: true
	RecipientType *string `json:"recipient_type"`
}

// Validate validates this post characters character id mail recipient
func (m *PostCharactersCharacterIDMailRecipient) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecipientID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRecipientType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostCharactersCharacterIDMailRecipient) validateRecipientID(formats strfmt.Registry) error {

	if err := validate.Required("recipient_id", "body", m.RecipientID); err != nil {
		return err
	}

	return nil
}

var postCharactersCharacterIdMailRecipientTypeRecipientTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alliance","character","corporation","mailing_list"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postCharactersCharacterIdMailRecipientTypeRecipientTypePropEnum = append(postCharactersCharacterIdMailRecipientTypeRecipientTypePropEnum, v)
	}
}

const (
	PostCharactersCharacterIDMailRecipientRecipientTypeAlliance    string = "alliance"
	PostCharactersCharacterIDMailRecipientRecipientTypeCharacter   string = "character"
	PostCharactersCharacterIDMailRecipientRecipientTypeCorporation string = "corporation"
	PostCharactersCharacterIDMailRecipientRecipientTypeMailingList string = "mailing_list"
)

// prop value enum
func (m *PostCharactersCharacterIDMailRecipient) validateRecipientTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postCharactersCharacterIdMailRecipientTypeRecipientTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PostCharactersCharacterIDMailRecipient) validateRecipientType(formats strfmt.Registry) error {

	if err := validate.Required("recipient_type", "body", m.RecipientType); err != nil {
		return err
	}

	// value enum
	if err := m.validateRecipientTypeEnum("recipient_type", "body", *m.RecipientType); err != nil {
		return err
	}

	return nil
}
