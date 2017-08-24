// Code generated by go-swagger; DO NOT EDIT.

package character

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

// GetCharactersCharacterIDStandingsReader is a Reader for the GetCharactersCharacterIDStandings structure.
type GetCharactersCharacterIDStandingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDStandingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDStandingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCharactersCharacterIDStandingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDStandingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDStandingsOK creates a GetCharactersCharacterIDStandingsOK with default headers values
func NewGetCharactersCharacterIDStandingsOK() *GetCharactersCharacterIDStandingsOK {
	return &GetCharactersCharacterIDStandingsOK{}
}

/*GetCharactersCharacterIDStandingsOK handles this case with default header values.

A list of standings
*/
type GetCharactersCharacterIDStandingsOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []*GetCharactersCharacterIDStandingsOKBodyItems0
}

func (o *GetCharactersCharacterIDStandingsOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/standings/][%d] getCharactersCharacterIdStandingsOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDStandingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDStandingsForbidden creates a GetCharactersCharacterIDStandingsForbidden with default headers values
func NewGetCharactersCharacterIDStandingsForbidden() *GetCharactersCharacterIDStandingsForbidden {
	return &GetCharactersCharacterIDStandingsForbidden{}
}

/*GetCharactersCharacterIDStandingsForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDStandingsForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDStandingsForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/standings/][%d] getCharactersCharacterIdStandingsForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDStandingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDStandingsInternalServerError creates a GetCharactersCharacterIDStandingsInternalServerError with default headers values
func NewGetCharactersCharacterIDStandingsInternalServerError() *GetCharactersCharacterIDStandingsInternalServerError {
	return &GetCharactersCharacterIDStandingsInternalServerError{}
}

/*GetCharactersCharacterIDStandingsInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDStandingsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDStandingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/standings/][%d] getCharactersCharacterIdStandingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDStandingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDStandingsOKBodyItems0 get_characters_character_id_standings_200_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDStandingsOKBodyItems0
*/

type GetCharactersCharacterIDStandingsOKBodyItems0 struct {

	// get_characters_character_id_standings_from_id
	//
	// from_id integer
	// Required: true
	FromID *int32 `json:"from_id"`

	// get_characters_character_id_standings_from_type
	//
	// from_type string
	// Required: true
	FromType *string `json:"from_type"`

	// get_characters_character_id_standings_standing
	//
	// standing number
	// Required: true
	// Maximum: 10
	// Minimum: -10
	Standing *float32 `json:"standing"`
}

/* polymorph GetCharactersCharacterIDStandingsOKBodyItems0 from_id false */

/* polymorph GetCharactersCharacterIDStandingsOKBodyItems0 from_type false */

/* polymorph GetCharactersCharacterIDStandingsOKBodyItems0 standing false */

// Validate validates this get characters character ID standings o k body items0
func (o *GetCharactersCharacterIDStandingsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFromID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateFromType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateStanding(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCharactersCharacterIDStandingsOKBodyItems0) validateFromID(formats strfmt.Registry) error {

	if err := validate.Required("from_id", "body", o.FromID); err != nil {
		return err
	}

	return nil
}

var getCharactersCharacterIdStandingsOKBodyItems0TypeFromTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["agent","npc_corp","faction"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdStandingsOKBodyItems0TypeFromTypePropEnum = append(getCharactersCharacterIdStandingsOKBodyItems0TypeFromTypePropEnum, v)
	}
}

const (
	// GetCharactersCharacterIDStandingsOKBodyItems0FromTypeAgent captures enum value "agent"
	GetCharactersCharacterIDStandingsOKBodyItems0FromTypeAgent string = "agent"
	// GetCharactersCharacterIDStandingsOKBodyItems0FromTypeNpcCorp captures enum value "npc_corp"
	GetCharactersCharacterIDStandingsOKBodyItems0FromTypeNpcCorp string = "npc_corp"
	// GetCharactersCharacterIDStandingsOKBodyItems0FromTypeFaction captures enum value "faction"
	GetCharactersCharacterIDStandingsOKBodyItems0FromTypeFaction string = "faction"
)

// prop value enum
func (o *GetCharactersCharacterIDStandingsOKBodyItems0) validateFromTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCharactersCharacterIdStandingsOKBodyItems0TypeFromTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetCharactersCharacterIDStandingsOKBodyItems0) validateFromType(formats strfmt.Registry) error {

	if err := validate.Required("from_type", "body", o.FromType); err != nil {
		return err
	}

	// value enum
	if err := o.validateFromTypeEnum("from_type", "body", *o.FromType); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDStandingsOKBodyItems0) validateStanding(formats strfmt.Registry) error {

	if err := validate.Required("standing", "body", o.Standing); err != nil {
		return err
	}

	if err := validate.Minimum("standing", "body", float64(*o.Standing), -10, false); err != nil {
		return err
	}

	if err := validate.Maximum("standing", "body", float64(*o.Standing), 10, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDStandingsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDStandingsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDStandingsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}