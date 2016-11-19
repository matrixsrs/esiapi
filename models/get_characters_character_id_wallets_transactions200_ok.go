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

// GetCharactersCharacterIDWalletsTransactions200Ok wallet transaction
// swagger:model get_characters_character_id_wallets_transactions_200_ok
type GetCharactersCharacterIDWalletsTransactions200Ok struct {

	// get_characters_character_id_wallets_transactions_client_id
	//
	// client_id integer
	ClientID int32 `json:"client_id,omitempty"`

	// get_characters_character_id_wallets_transactions_client_type
	//
	// client_type string
	ClientType string `json:"client_type,omitempty"`

	// get_characters_character_id_wallets_transactions_journal_ref_id
	//
	// journal_ref_id integer
	JournalRefID int64 `json:"journal_ref_id,omitempty"`

	// get_characters_character_id_wallets_transactions_location_id
	//
	// location_id integer
	LocationID int64 `json:"location_id,omitempty"`

	// get_characters_character_id_wallets_transactions_location_type
	//
	// location_type string
	LocationType string `json:"location_type,omitempty"`

	// get_characters_character_id_wallets_transactions_price_per_unit
	//
	// price_per_unit integer
	PricePerUnit int64 `json:"price_per_unit,omitempty"`

	// get_characters_character_id_wallets_transactions_quantity
	//
	// quantity integer
	Quantity int32 `json:"quantity,omitempty"`

	// get_characters_character_id_wallets_transactions_transaction_date
	//
	// transaction_date string
	// Required: true
	TransactionDate *strfmt.DateTime `json:"transaction_date"`

	// get_characters_character_id_wallets_transactions_transaction_for
	//
	// transaction_for string
	TransactionFor string `json:"transaction_for,omitempty"`

	// get_characters_character_id_wallets_transactions_transaction_id
	//
	// transaction_id integer
	// Required: true
	TransactionID *int64 `json:"transaction_id"`

	// get_characters_character_id_wallets_transactions_transaction_type
	//
	// transaction_type string
	TransactionType string `json:"transaction_type,omitempty"`

	// get_characters_character_id_wallets_transactions_type_id
	//
	// type_id integer
	TypeID int32 `json:"type_id,omitempty"`
}

// Validate validates this get characters character id wallets transactions 200 ok
func (m *GetCharactersCharacterIDWalletsTransactions200Ok) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLocationType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTransactionDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTransactionFor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTransactionID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTransactionType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getCharactersCharacterIdWalletsTransactions200OkTypeClientTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["character","corporation","alliance"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdWalletsTransactions200OkTypeClientTypePropEnum = append(getCharactersCharacterIdWalletsTransactions200OkTypeClientTypePropEnum, v)
	}
}

const (
	GetCharactersCharacterIDWalletsTransactions200OkClientTypeCharacter   string = "character"
	GetCharactersCharacterIDWalletsTransactions200OkClientTypeCorporation string = "corporation"
	GetCharactersCharacterIDWalletsTransactions200OkClientTypeAlliance    string = "alliance"
)

// prop value enum
func (m *GetCharactersCharacterIDWalletsTransactions200Ok) validateClientTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCharactersCharacterIdWalletsTransactions200OkTypeClientTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetCharactersCharacterIDWalletsTransactions200Ok) validateClientType(formats strfmt.Registry) error {

	if swag.IsZero(m.ClientType) { // not required
		return nil
	}

	// value enum
	if err := m.validateClientTypeEnum("client_type", "body", m.ClientType); err != nil {
		return err
	}

	return nil
}

var getCharactersCharacterIdWalletsTransactions200OkTypeLocationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["station","structure"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdWalletsTransactions200OkTypeLocationTypePropEnum = append(getCharactersCharacterIdWalletsTransactions200OkTypeLocationTypePropEnum, v)
	}
}

const (
	GetCharactersCharacterIDWalletsTransactions200OkLocationTypeStation   string = "station"
	GetCharactersCharacterIDWalletsTransactions200OkLocationTypeStructure string = "structure"
)

// prop value enum
func (m *GetCharactersCharacterIDWalletsTransactions200Ok) validateLocationTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCharactersCharacterIdWalletsTransactions200OkTypeLocationTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetCharactersCharacterIDWalletsTransactions200Ok) validateLocationType(formats strfmt.Registry) error {

	if swag.IsZero(m.LocationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateLocationTypeEnum("location_type", "body", m.LocationType); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDWalletsTransactions200Ok) validateTransactionDate(formats strfmt.Registry) error {

	if err := validate.Required("transaction_date", "body", m.TransactionDate); err != nil {
		return err
	}

	return nil
}

var getCharactersCharacterIdWalletsTransactions200OkTypeTransactionForPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["personal","corporate"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdWalletsTransactions200OkTypeTransactionForPropEnum = append(getCharactersCharacterIdWalletsTransactions200OkTypeTransactionForPropEnum, v)
	}
}

const (
	GetCharactersCharacterIDWalletsTransactions200OkTransactionForPersonal  string = "personal"
	GetCharactersCharacterIDWalletsTransactions200OkTransactionForCorporate string = "corporate"
)

// prop value enum
func (m *GetCharactersCharacterIDWalletsTransactions200Ok) validateTransactionForEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCharactersCharacterIdWalletsTransactions200OkTypeTransactionForPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetCharactersCharacterIDWalletsTransactions200Ok) validateTransactionFor(formats strfmt.Registry) error {

	if swag.IsZero(m.TransactionFor) { // not required
		return nil
	}

	// value enum
	if err := m.validateTransactionForEnum("transaction_for", "body", m.TransactionFor); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDWalletsTransactions200Ok) validateTransactionID(formats strfmt.Registry) error {

	if err := validate.Required("transaction_id", "body", m.TransactionID); err != nil {
		return err
	}

	return nil
}

var getCharactersCharacterIdWalletsTransactions200OkTypeTransactionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["buy","sell"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdWalletsTransactions200OkTypeTransactionTypePropEnum = append(getCharactersCharacterIdWalletsTransactions200OkTypeTransactionTypePropEnum, v)
	}
}

const (
	GetCharactersCharacterIDWalletsTransactions200OkTransactionTypeBuy  string = "buy"
	GetCharactersCharacterIDWalletsTransactions200OkTransactionTypeSell string = "sell"
)

// prop value enum
func (m *GetCharactersCharacterIDWalletsTransactions200Ok) validateTransactionTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCharactersCharacterIdWalletsTransactions200OkTypeTransactionTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetCharactersCharacterIDWalletsTransactions200Ok) validateTransactionType(formats strfmt.Registry) error {

	if swag.IsZero(m.TransactionType) { // not required
		return nil
	}

	// value enum
	if err := m.validateTransactionTypeEnum("transaction_type", "body", m.TransactionType); err != nil {
		return err
	}

	return nil
}
