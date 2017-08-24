// Code generated by go-swagger; DO NOT EDIT.

package contracts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/models"
)

// GetCharactersCharacterIDContractsReader is a Reader for the GetCharactersCharacterIDContracts structure.
type GetCharactersCharacterIDContractsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDContractsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDContractsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCharactersCharacterIDContractsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDContractsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDContractsOK creates a GetCharactersCharacterIDContractsOK with default headers values
func NewGetCharactersCharacterIDContractsOK() *GetCharactersCharacterIDContractsOK {
	return &GetCharactersCharacterIDContractsOK{}
}

/*GetCharactersCharacterIDContractsOK handles this case with default header values.

A list of contracts
*/
type GetCharactersCharacterIDContractsOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []*GetCharactersCharacterIDContractsOKBodyItems0
}

func (o *GetCharactersCharacterIDContractsOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contracts/][%d] getCharactersCharacterIdContractsOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDContractsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContractsForbidden creates a GetCharactersCharacterIDContractsForbidden with default headers values
func NewGetCharactersCharacterIDContractsForbidden() *GetCharactersCharacterIDContractsForbidden {
	return &GetCharactersCharacterIDContractsForbidden{}
}

/*GetCharactersCharacterIDContractsForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDContractsForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDContractsForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contracts/][%d] getCharactersCharacterIdContractsForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDContractsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContractsInternalServerError creates a GetCharactersCharacterIDContractsInternalServerError with default headers values
func NewGetCharactersCharacterIDContractsInternalServerError() *GetCharactersCharacterIDContractsInternalServerError {
	return &GetCharactersCharacterIDContractsInternalServerError{}
}

/*GetCharactersCharacterIDContractsInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDContractsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDContractsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contracts/][%d] getCharactersCharacterIdContractsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDContractsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDContractsOKBodyItems0 get_characters_character_id_contracts_200_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDContractsOKBodyItems0
*/

type GetCharactersCharacterIDContractsOKBodyItems0 struct {

	// get_characters_character_id_contracts_acceptor_id
	//
	// Who will accept the contract. If assignee_id is same as acceptorID then character ID else corporation ID (The contract accepted by the corporation)
	// Required: true
	AcceptorID *int32 `json:"acceptor_id"`

	// get_characters_character_id_contracts_assignee_id
	//
	// ID to whom the contract is assigned, can be corporation or character ID
	// Required: true
	AssigneeID *int32 `json:"assignee_id"`

	// get_characters_character_id_contracts_availability
	//
	// To whom the contract is available
	// Required: true
	Availability *string `json:"availability"`

	// get_characters_character_id_contracts_buyout
	//
	// Buyout price (for Auctions only)
	Buyout float32 `json:"buyout,omitempty"`

	// get_characters_character_id_contracts_collateral
	//
	// Collateral price (for Couriers only)
	Collateral float32 `json:"collateral,omitempty"`

	// get_characters_character_id_contracts_contract_id
	//
	// contract_id integer
	// Required: true
	ContractID *int32 `json:"contract_id"`

	// get_characters_character_id_contracts_date_accepted
	//
	// Date of confirmation of contract
	DateAccepted strfmt.DateTime `json:"date_accepted,omitempty"`

	// get_characters_character_id_contracts_date_completed
	//
	// Date of completed of contract
	DateCompleted strfmt.DateTime `json:"date_completed,omitempty"`

	// get_characters_character_id_contracts_date_expired
	//
	// Expiration date of the contract
	// Required: true
	DateExpired *strfmt.DateTime `json:"date_expired"`

	// get_characters_character_id_contracts_date_issued
	//
	// Сreation date of the contract
	// Required: true
	DateIssued *strfmt.DateTime `json:"date_issued"`

	// get_characters_character_id_contracts_days_to_complete
	//
	// Number of days to perform the contract
	DaysToComplete int32 `json:"days_to_complete,omitempty"`

	// get_characters_character_id_contracts_end_location_id
	//
	// End location ID (for Couriers contract)
	EndLocationID int64 `json:"end_location_id,omitempty"`

	// get_characters_character_id_contracts_for_corporation
	//
	// true if the contract was issued on behalf of the issuer's corporation
	// Required: true
	ForCorporation *bool `json:"for_corporation"`

	// get_characters_character_id_contracts_issuer_corporation_id
	//
	// Character's corporation ID for the issuer
	// Required: true
	IssuerCorporationID *int32 `json:"issuer_corporation_id"`

	// get_characters_character_id_contracts_issuer_id
	//
	// Character ID for the issuer
	// Required: true
	IssuerID *int32 `json:"issuer_id"`

	// get_characters_character_id_contracts_price
	//
	// Price of contract (for ItemsExchange and Auctions)
	Price float32 `json:"price,omitempty"`

	// get_characters_character_id_contracts_reward
	//
	// Remuneration for contract (for Couriers only)
	Reward float32 `json:"reward,omitempty"`

	// get_characters_character_id_contracts_start_location_id
	//
	// Start location ID (for Couriers contract)
	StartLocationID int64 `json:"start_location_id,omitempty"`

	// get_characters_character_id_contracts_status
	//
	// Status of the the contract
	// Required: true
	Status *string `json:"status"`

	// get_characters_character_id_contracts_title
	//
	// Title of the contract
	Title string `json:"title,omitempty"`

	// get_characters_character_id_contracts_type
	//
	// Type of the contract
	// Required: true
	Type *string `json:"type"`

	// get_characters_character_id_contracts_volume
	//
	// Volume of items in the contract
	Volume float32 `json:"volume,omitempty"`
}

/* polymorph GetCharactersCharacterIDContractsOKBodyItems0 acceptor_id false */

/* polymorph GetCharactersCharacterIDContractsOKBodyItems0 assignee_id false */

/* polymorph GetCharactersCharacterIDContractsOKBodyItems0 availability false */

/* polymorph GetCharactersCharacterIDContractsOKBodyItems0 buyout false */

/* polymorph GetCharactersCharacterIDContractsOKBodyItems0 collateral false */

/* polymorph GetCharactersCharacterIDContractsOKBodyItems0 contract_id false */

/* polymorph GetCharactersCharacterIDContractsOKBodyItems0 date_accepted false */

/* polymorph GetCharactersCharacterIDContractsOKBodyItems0 date_completed false */

/* polymorph GetCharactersCharacterIDContractsOKBodyItems0 date_expired false */

/* polymorph GetCharactersCharacterIDContractsOKBodyItems0 date_issued false */

/* polymorph GetCharactersCharacterIDContractsOKBodyItems0 days_to_complete false */

/* polymorph GetCharactersCharacterIDContractsOKBodyItems0 end_location_id false */

/* polymorph GetCharactersCharacterIDContractsOKBodyItems0 for_corporation false */

/* polymorph GetCharactersCharacterIDContractsOKBodyItems0 issuer_corporation_id false */

/* polymorph GetCharactersCharacterIDContractsOKBodyItems0 issuer_id false */

/* polymorph GetCharactersCharacterIDContractsOKBodyItems0 price false */

/* polymorph GetCharactersCharacterIDContractsOKBodyItems0 reward false */

/* polymorph GetCharactersCharacterIDContractsOKBodyItems0 start_location_id false */

/* polymorph GetCharactersCharacterIDContractsOKBodyItems0 status false */

/* polymorph GetCharactersCharacterIDContractsOKBodyItems0 title false */

/* polymorph GetCharactersCharacterIDContractsOKBodyItems0 type false */

/* polymorph GetCharactersCharacterIDContractsOKBodyItems0 volume false */

// Validate validates this get characters character ID contracts o k body items0
func (o *GetCharactersCharacterIDContractsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAcceptorID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateAssigneeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateAvailability(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateContractID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateDateExpired(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateDateIssued(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateForCorporation(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateIssuerCorporationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateIssuerID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCharactersCharacterIDContractsOKBodyItems0) validateAcceptorID(formats strfmt.Registry) error {

	if err := validate.Required("acceptor_id", "body", o.AcceptorID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDContractsOKBodyItems0) validateAssigneeID(formats strfmt.Registry) error {

	if err := validate.Required("assignee_id", "body", o.AssigneeID); err != nil {
		return err
	}

	return nil
}

var getCharactersCharacterIdContractsOKBodyItems0TypeAvailabilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["public","personal","corporation","alliance"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdContractsOKBodyItems0TypeAvailabilityPropEnum = append(getCharactersCharacterIdContractsOKBodyItems0TypeAvailabilityPropEnum, v)
	}
}

const (
	// GetCharactersCharacterIDContractsOKBodyItems0AvailabilityPublic captures enum value "public"
	GetCharactersCharacterIDContractsOKBodyItems0AvailabilityPublic string = "public"
	// GetCharactersCharacterIDContractsOKBodyItems0AvailabilityPersonal captures enum value "personal"
	GetCharactersCharacterIDContractsOKBodyItems0AvailabilityPersonal string = "personal"
	// GetCharactersCharacterIDContractsOKBodyItems0AvailabilityCorporation captures enum value "corporation"
	GetCharactersCharacterIDContractsOKBodyItems0AvailabilityCorporation string = "corporation"
	// GetCharactersCharacterIDContractsOKBodyItems0AvailabilityAlliance captures enum value "alliance"
	GetCharactersCharacterIDContractsOKBodyItems0AvailabilityAlliance string = "alliance"
)

// prop value enum
func (o *GetCharactersCharacterIDContractsOKBodyItems0) validateAvailabilityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCharactersCharacterIdContractsOKBodyItems0TypeAvailabilityPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetCharactersCharacterIDContractsOKBodyItems0) validateAvailability(formats strfmt.Registry) error {

	if err := validate.Required("availability", "body", o.Availability); err != nil {
		return err
	}

	// value enum
	if err := o.validateAvailabilityEnum("availability", "body", *o.Availability); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDContractsOKBodyItems0) validateContractID(formats strfmt.Registry) error {

	if err := validate.Required("contract_id", "body", o.ContractID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDContractsOKBodyItems0) validateDateExpired(formats strfmt.Registry) error {

	if err := validate.Required("date_expired", "body", o.DateExpired); err != nil {
		return err
	}

	if err := validate.FormatOf("date_expired", "body", "date-time", o.DateExpired.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDContractsOKBodyItems0) validateDateIssued(formats strfmt.Registry) error {

	if err := validate.Required("date_issued", "body", o.DateIssued); err != nil {
		return err
	}

	if err := validate.FormatOf("date_issued", "body", "date-time", o.DateIssued.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDContractsOKBodyItems0) validateForCorporation(formats strfmt.Registry) error {

	if err := validate.Required("for_corporation", "body", o.ForCorporation); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDContractsOKBodyItems0) validateIssuerCorporationID(formats strfmt.Registry) error {

	if err := validate.Required("issuer_corporation_id", "body", o.IssuerCorporationID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDContractsOKBodyItems0) validateIssuerID(formats strfmt.Registry) error {

	if err := validate.Required("issuer_id", "body", o.IssuerID); err != nil {
		return err
	}

	return nil
}

var getCharactersCharacterIdContractsOKBodyItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["outstanding","in_progress","finished_issuer","finished_contractor","finished","cancelled","rejected","failed","deleted","reversed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdContractsOKBodyItems0TypeStatusPropEnum = append(getCharactersCharacterIdContractsOKBodyItems0TypeStatusPropEnum, v)
	}
}

const (
	// GetCharactersCharacterIDContractsOKBodyItems0StatusOutstanding captures enum value "outstanding"
	GetCharactersCharacterIDContractsOKBodyItems0StatusOutstanding string = "outstanding"
	// GetCharactersCharacterIDContractsOKBodyItems0StatusInProgress captures enum value "in_progress"
	GetCharactersCharacterIDContractsOKBodyItems0StatusInProgress string = "in_progress"
	// GetCharactersCharacterIDContractsOKBodyItems0StatusFinishedIssuer captures enum value "finished_issuer"
	GetCharactersCharacterIDContractsOKBodyItems0StatusFinishedIssuer string = "finished_issuer"
	// GetCharactersCharacterIDContractsOKBodyItems0StatusFinishedContractor captures enum value "finished_contractor"
	GetCharactersCharacterIDContractsOKBodyItems0StatusFinishedContractor string = "finished_contractor"
	// GetCharactersCharacterIDContractsOKBodyItems0StatusFinished captures enum value "finished"
	GetCharactersCharacterIDContractsOKBodyItems0StatusFinished string = "finished"
	// GetCharactersCharacterIDContractsOKBodyItems0StatusCancelled captures enum value "cancelled"
	GetCharactersCharacterIDContractsOKBodyItems0StatusCancelled string = "cancelled"
	// GetCharactersCharacterIDContractsOKBodyItems0StatusRejected captures enum value "rejected"
	GetCharactersCharacterIDContractsOKBodyItems0StatusRejected string = "rejected"
	// GetCharactersCharacterIDContractsOKBodyItems0StatusFailed captures enum value "failed"
	GetCharactersCharacterIDContractsOKBodyItems0StatusFailed string = "failed"
	// GetCharactersCharacterIDContractsOKBodyItems0StatusDeleted captures enum value "deleted"
	GetCharactersCharacterIDContractsOKBodyItems0StatusDeleted string = "deleted"
	// GetCharactersCharacterIDContractsOKBodyItems0StatusReversed captures enum value "reversed"
	GetCharactersCharacterIDContractsOKBodyItems0StatusReversed string = "reversed"
)

// prop value enum
func (o *GetCharactersCharacterIDContractsOKBodyItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCharactersCharacterIdContractsOKBodyItems0TypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetCharactersCharacterIDContractsOKBodyItems0) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", o.Status); err != nil {
		return err
	}

	// value enum
	if err := o.validateStatusEnum("status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

var getCharactersCharacterIdContractsOKBodyItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unknown","item_exchange","auction","courier","loan"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdContractsOKBodyItems0TypeTypePropEnum = append(getCharactersCharacterIdContractsOKBodyItems0TypeTypePropEnum, v)
	}
}

const (
	// GetCharactersCharacterIDContractsOKBodyItems0TypeUnknown captures enum value "unknown"
	GetCharactersCharacterIDContractsOKBodyItems0TypeUnknown string = "unknown"
	// GetCharactersCharacterIDContractsOKBodyItems0TypeItemExchange captures enum value "item_exchange"
	GetCharactersCharacterIDContractsOKBodyItems0TypeItemExchange string = "item_exchange"
	// GetCharactersCharacterIDContractsOKBodyItems0TypeAuction captures enum value "auction"
	GetCharactersCharacterIDContractsOKBodyItems0TypeAuction string = "auction"
	// GetCharactersCharacterIDContractsOKBodyItems0TypeCourier captures enum value "courier"
	GetCharactersCharacterIDContractsOKBodyItems0TypeCourier string = "courier"
	// GetCharactersCharacterIDContractsOKBodyItems0TypeLoan captures enum value "loan"
	GetCharactersCharacterIDContractsOKBodyItems0TypeLoan string = "loan"
)

// prop value enum
func (o *GetCharactersCharacterIDContractsOKBodyItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCharactersCharacterIdContractsOKBodyItems0TypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetCharactersCharacterIDContractsOKBodyItems0) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", o.Type); err != nil {
		return err
	}

	// value enum
	if err := o.validateTypeEnum("type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDContractsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDContractsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDContractsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
