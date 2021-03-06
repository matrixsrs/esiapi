// Code generated by go-swagger; DO NOT EDIT.

package fittings

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

// PostCharactersCharacterIDFittingsReader is a Reader for the PostCharactersCharacterIDFittings structure.
type PostCharactersCharacterIDFittingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCharactersCharacterIDFittingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostCharactersCharacterIDFittingsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewPostCharactersCharacterIDFittingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostCharactersCharacterIDFittingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostCharactersCharacterIDFittingsCreated creates a PostCharactersCharacterIDFittingsCreated with default headers values
func NewPostCharactersCharacterIDFittingsCreated() *PostCharactersCharacterIDFittingsCreated {
	return &PostCharactersCharacterIDFittingsCreated{}
}

/*PostCharactersCharacterIDFittingsCreated handles this case with default header values.

A list of fittings
*/
type PostCharactersCharacterIDFittingsCreated struct {
	Payload PostCharactersCharacterIDFittingsCreatedBody
}

func (o *PostCharactersCharacterIDFittingsCreated) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/fittings/][%d] postCharactersCharacterIdFittingsCreated  %+v", 201, o.Payload)
}

func (o *PostCharactersCharacterIDFittingsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDFittingsForbidden creates a PostCharactersCharacterIDFittingsForbidden with default headers values
func NewPostCharactersCharacterIDFittingsForbidden() *PostCharactersCharacterIDFittingsForbidden {
	return &PostCharactersCharacterIDFittingsForbidden{}
}

/*PostCharactersCharacterIDFittingsForbidden handles this case with default header values.

Forbidden
*/
type PostCharactersCharacterIDFittingsForbidden struct {
	Payload *models.Forbidden
}

func (o *PostCharactersCharacterIDFittingsForbidden) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/fittings/][%d] postCharactersCharacterIdFittingsForbidden  %+v", 403, o.Payload)
}

func (o *PostCharactersCharacterIDFittingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDFittingsInternalServerError creates a PostCharactersCharacterIDFittingsInternalServerError with default headers values
func NewPostCharactersCharacterIDFittingsInternalServerError() *PostCharactersCharacterIDFittingsInternalServerError {
	return &PostCharactersCharacterIDFittingsInternalServerError{}
}

/*PostCharactersCharacterIDFittingsInternalServerError handles this case with default header values.

Internal server error
*/
type PostCharactersCharacterIDFittingsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *PostCharactersCharacterIDFittingsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/fittings/][%d] postCharactersCharacterIdFittingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostCharactersCharacterIDFittingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostCharactersCharacterIDFittingsBody post_characters_character_id_fittings_fitting
//
// fitting object
swagger:model PostCharactersCharacterIDFittingsBody
*/

type PostCharactersCharacterIDFittingsBody struct {

	// post_characters_character_id_fittings_description
	//
	// description string
	// Required: true
	// Max Length: 500
	// Min Length: 0
	Description *string `json:"description"`

	// post_characters_character_id_fittings_items
	//
	// items array
	// Required: true
	// Max Items: 255
	// Min Items: 1
	Items []*GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0 `json:"items"`

	// post_characters_character_id_fittings_name
	//
	// name string
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// post_characters_character_id_fittings_ship_type_id
	//
	// ship_type_id integer
	// Required: true
	ShipTypeID *int32 `json:"ship_type_id"`
}

/* polymorph PostCharactersCharacterIDFittingsBody description false */

/* polymorph PostCharactersCharacterIDFittingsBody items false */

/* polymorph PostCharactersCharacterIDFittingsBody name false */

/* polymorph PostCharactersCharacterIDFittingsBody ship_type_id false */

// MarshalBinary interface implementation
func (o *PostCharactersCharacterIDFittingsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCharactersCharacterIDFittingsBody) UnmarshalBinary(b []byte) error {
	var res PostCharactersCharacterIDFittingsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostCharactersCharacterIDFittingsCreatedBody post_characters_character_id_fittings_created
//
// 201 created object
swagger:model PostCharactersCharacterIDFittingsCreatedBody
*/

type PostCharactersCharacterIDFittingsCreatedBody struct {

	// post_characters_character_id_fittings_fitting_id
	//
	// fitting_id integer
	// Required: true
	FittingID *int32 `json:"fitting_id"`
}

/* polymorph PostCharactersCharacterIDFittingsCreatedBody fitting_id false */

// Validate validates this post characters character ID fittings created body
func (o *PostCharactersCharacterIDFittingsCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFittingID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostCharactersCharacterIDFittingsCreatedBody) validateFittingID(formats strfmt.Registry) error {

	if err := validate.Required("postCharactersCharacterIdFittingsCreated"+"."+"fitting_id", "body", o.FittingID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostCharactersCharacterIDFittingsCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCharactersCharacterIDFittingsCreatedBody) UnmarshalBinary(b []byte) error {
	var res PostCharactersCharacterIDFittingsCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
