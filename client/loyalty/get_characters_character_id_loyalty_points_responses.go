// Code generated by go-swagger; DO NOT EDIT.

package loyalty

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/models"
)

// GetCharactersCharacterIDLoyaltyPointsReader is a Reader for the GetCharactersCharacterIDLoyaltyPoints structure.
type GetCharactersCharacterIDLoyaltyPointsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDLoyaltyPointsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDLoyaltyPointsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCharactersCharacterIDLoyaltyPointsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDLoyaltyPointsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDLoyaltyPointsOK creates a GetCharactersCharacterIDLoyaltyPointsOK with default headers values
func NewGetCharactersCharacterIDLoyaltyPointsOK() *GetCharactersCharacterIDLoyaltyPointsOK {
	return &GetCharactersCharacterIDLoyaltyPointsOK{}
}

/*GetCharactersCharacterIDLoyaltyPointsOK handles this case with default header values.

A list of loyalty points
*/
type GetCharactersCharacterIDLoyaltyPointsOK struct {
	Payload []*GetCharactersCharacterIDLoyaltyPointsOKBodyItems0
}

func (o *GetCharactersCharacterIDLoyaltyPointsOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/loyalty/points/][%d] getCharactersCharacterIdLoyaltyPointsOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDLoyaltyPointsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDLoyaltyPointsForbidden creates a GetCharactersCharacterIDLoyaltyPointsForbidden with default headers values
func NewGetCharactersCharacterIDLoyaltyPointsForbidden() *GetCharactersCharacterIDLoyaltyPointsForbidden {
	return &GetCharactersCharacterIDLoyaltyPointsForbidden{}
}

/*GetCharactersCharacterIDLoyaltyPointsForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDLoyaltyPointsForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDLoyaltyPointsForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/loyalty/points/][%d] getCharactersCharacterIdLoyaltyPointsForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDLoyaltyPointsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDLoyaltyPointsInternalServerError creates a GetCharactersCharacterIDLoyaltyPointsInternalServerError with default headers values
func NewGetCharactersCharacterIDLoyaltyPointsInternalServerError() *GetCharactersCharacterIDLoyaltyPointsInternalServerError {
	return &GetCharactersCharacterIDLoyaltyPointsInternalServerError{}
}

/*GetCharactersCharacterIDLoyaltyPointsInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDLoyaltyPointsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDLoyaltyPointsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/loyalty/points/][%d] getCharactersCharacterIdLoyaltyPointsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDLoyaltyPointsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDLoyaltyPointsOKBodyItems0 get_characters_character_id_loyalty_points_200_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDLoyaltyPointsOKBodyItems0
*/

type GetCharactersCharacterIDLoyaltyPointsOKBodyItems0 struct {

	// get_characters_character_id_loyalty_points_corporation_id
	//
	// corporation_id integer
	// Required: true
	CorporationID *int32 `json:"corporation_id"`

	// get_characters_character_id_loyalty_points_loyalty_points
	//
	// loyalty_points integer
	// Required: true
	LoyaltyPoints *int32 `json:"loyalty_points"`
}

/* polymorph GetCharactersCharacterIDLoyaltyPointsOKBodyItems0 corporation_id false */

/* polymorph GetCharactersCharacterIDLoyaltyPointsOKBodyItems0 loyalty_points false */

// Validate validates this get characters character ID loyalty points o k body items0
func (o *GetCharactersCharacterIDLoyaltyPointsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCorporationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateLoyaltyPoints(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCharactersCharacterIDLoyaltyPointsOKBodyItems0) validateCorporationID(formats strfmt.Registry) error {

	if err := validate.Required("corporation_id", "body", o.CorporationID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDLoyaltyPointsOKBodyItems0) validateLoyaltyPoints(formats strfmt.Registry) error {

	if err := validate.Required("loyalty_points", "body", o.LoyaltyPoints); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDLoyaltyPointsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDLoyaltyPointsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDLoyaltyPointsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
