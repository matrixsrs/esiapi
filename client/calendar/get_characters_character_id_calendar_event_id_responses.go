// Code generated by go-swagger; DO NOT EDIT.

package calendar

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

// GetCharactersCharacterIDCalendarEventIDReader is a Reader for the GetCharactersCharacterIDCalendarEventID structure.
type GetCharactersCharacterIDCalendarEventIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDCalendarEventIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDCalendarEventIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCharactersCharacterIDCalendarEventIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDCalendarEventIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDCalendarEventIDOK creates a GetCharactersCharacterIDCalendarEventIDOK with default headers values
func NewGetCharactersCharacterIDCalendarEventIDOK() *GetCharactersCharacterIDCalendarEventIDOK {
	return &GetCharactersCharacterIDCalendarEventIDOK{}
}

/*GetCharactersCharacterIDCalendarEventIDOK handles this case with default header values.

Full details of a specific event
*/
type GetCharactersCharacterIDCalendarEventIDOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload GetCharactersCharacterIDCalendarEventIDOKBody
}

func (o *GetCharactersCharacterIDCalendarEventIDOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/calendar/{event_id}/][%d] getCharactersCharacterIdCalendarEventIdOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDCalendarEventIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDCalendarEventIDForbidden creates a GetCharactersCharacterIDCalendarEventIDForbidden with default headers values
func NewGetCharactersCharacterIDCalendarEventIDForbidden() *GetCharactersCharacterIDCalendarEventIDForbidden {
	return &GetCharactersCharacterIDCalendarEventIDForbidden{}
}

/*GetCharactersCharacterIDCalendarEventIDForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDCalendarEventIDForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDCalendarEventIDForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/calendar/{event_id}/][%d] getCharactersCharacterIdCalendarEventIdForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDCalendarEventIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDCalendarEventIDInternalServerError creates a GetCharactersCharacterIDCalendarEventIDInternalServerError with default headers values
func NewGetCharactersCharacterIDCalendarEventIDInternalServerError() *GetCharactersCharacterIDCalendarEventIDInternalServerError {
	return &GetCharactersCharacterIDCalendarEventIDInternalServerError{}
}

/*GetCharactersCharacterIDCalendarEventIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDCalendarEventIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDCalendarEventIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/calendar/{event_id}/][%d] getCharactersCharacterIdCalendarEventIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDCalendarEventIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDCalendarEventIDOKBody get_characters_character_id_calendar_event_id_ok
//
// Full details of a specific event
swagger:model GetCharactersCharacterIDCalendarEventIDOKBody
*/

type GetCharactersCharacterIDCalendarEventIDOKBody struct {

	// get_characters_character_id_calendar_event_id_date
	//
	// date string
	// Required: true
	Date *strfmt.DateTime `json:"date"`

	// get_characters_character_id_calendar_event_id_duration
	//
	// Length in minutes
	// Required: true
	Duration *int32 `json:"duration"`

	// get_characters_character_id_calendar_event_id_event_id
	//
	// event_id integer
	// Required: true
	EventID *int32 `json:"event_id"`

	// get_characters_character_id_calendar_event_id_importance
	//
	// importance integer
	// Required: true
	Importance *int32 `json:"importance"`

	// get_characters_character_id_calendar_event_id_owner_id
	//
	// owner_id integer
	// Required: true
	OwnerID *int32 `json:"owner_id"`

	// get_characters_character_id_calendar_event_id_owner_name
	//
	// owner_name string
	// Required: true
	OwnerName *string `json:"owner_name"`

	// get_characters_character_id_calendar_event_id_owner_type
	//
	// owner_type string
	// Required: true
	OwnerType *string `json:"owner_type"`

	// get_characters_character_id_calendar_event_id_response
	//
	// response string
	// Required: true
	Response *string `json:"response"`

	// get_characters_character_id_calendar_event_id_text
	//
	// text string
	// Required: true
	Text *string `json:"text"`

	// get_characters_character_id_calendar_event_id_title
	//
	// title string
	// Required: true
	Title *string `json:"title"`
}

/* polymorph GetCharactersCharacterIDCalendarEventIDOKBody date false */

/* polymorph GetCharactersCharacterIDCalendarEventIDOKBody duration false */

/* polymorph GetCharactersCharacterIDCalendarEventIDOKBody event_id false */

/* polymorph GetCharactersCharacterIDCalendarEventIDOKBody importance false */

/* polymorph GetCharactersCharacterIDCalendarEventIDOKBody owner_id false */

/* polymorph GetCharactersCharacterIDCalendarEventIDOKBody owner_name false */

/* polymorph GetCharactersCharacterIDCalendarEventIDOKBody owner_type false */

/* polymorph GetCharactersCharacterIDCalendarEventIDOKBody response false */

/* polymorph GetCharactersCharacterIDCalendarEventIDOKBody text false */

/* polymorph GetCharactersCharacterIDCalendarEventIDOKBody title false */

// Validate validates this get characters character ID calendar event ID o k body
func (o *GetCharactersCharacterIDCalendarEventIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateDuration(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateEventID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateImportance(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateOwnerID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateOwnerName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateOwnerType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateResponse(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateText(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateTitle(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCharactersCharacterIDCalendarEventIDOKBody) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdCalendarEventIdOK"+"."+"date", "body", o.Date); err != nil {
		return err
	}

	if err := validate.FormatOf("getCharactersCharacterIdCalendarEventIdOK"+"."+"date", "body", "date-time", o.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDCalendarEventIDOKBody) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdCalendarEventIdOK"+"."+"duration", "body", o.Duration); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDCalendarEventIDOKBody) validateEventID(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdCalendarEventIdOK"+"."+"event_id", "body", o.EventID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDCalendarEventIDOKBody) validateImportance(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdCalendarEventIdOK"+"."+"importance", "body", o.Importance); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDCalendarEventIDOKBody) validateOwnerID(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdCalendarEventIdOK"+"."+"owner_id", "body", o.OwnerID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDCalendarEventIDOKBody) validateOwnerName(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdCalendarEventIdOK"+"."+"owner_name", "body", o.OwnerName); err != nil {
		return err
	}

	return nil
}

var getCharactersCharacterIdCalendarEventIdOKBodyTypeOwnerTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["eve_server","corporation","faction","character","alliance"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdCalendarEventIdOKBodyTypeOwnerTypePropEnum = append(getCharactersCharacterIdCalendarEventIdOKBodyTypeOwnerTypePropEnum, v)
	}
}

const (
	// GetCharactersCharacterIDCalendarEventIDOKBodyOwnerTypeEveServer captures enum value "eve_server"
	GetCharactersCharacterIDCalendarEventIDOKBodyOwnerTypeEveServer string = "eve_server"
	// GetCharactersCharacterIDCalendarEventIDOKBodyOwnerTypeCorporation captures enum value "corporation"
	GetCharactersCharacterIDCalendarEventIDOKBodyOwnerTypeCorporation string = "corporation"
	// GetCharactersCharacterIDCalendarEventIDOKBodyOwnerTypeFaction captures enum value "faction"
	GetCharactersCharacterIDCalendarEventIDOKBodyOwnerTypeFaction string = "faction"
	// GetCharactersCharacterIDCalendarEventIDOKBodyOwnerTypeCharacter captures enum value "character"
	GetCharactersCharacterIDCalendarEventIDOKBodyOwnerTypeCharacter string = "character"
	// GetCharactersCharacterIDCalendarEventIDOKBodyOwnerTypeAlliance captures enum value "alliance"
	GetCharactersCharacterIDCalendarEventIDOKBodyOwnerTypeAlliance string = "alliance"
)

// prop value enum
func (o *GetCharactersCharacterIDCalendarEventIDOKBody) validateOwnerTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCharactersCharacterIdCalendarEventIdOKBodyTypeOwnerTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetCharactersCharacterIDCalendarEventIDOKBody) validateOwnerType(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdCalendarEventIdOK"+"."+"owner_type", "body", o.OwnerType); err != nil {
		return err
	}

	// value enum
	if err := o.validateOwnerTypeEnum("getCharactersCharacterIdCalendarEventIdOK"+"."+"owner_type", "body", *o.OwnerType); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDCalendarEventIDOKBody) validateResponse(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdCalendarEventIdOK"+"."+"response", "body", o.Response); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDCalendarEventIDOKBody) validateText(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdCalendarEventIdOK"+"."+"text", "body", o.Text); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDCalendarEventIDOKBody) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdCalendarEventIdOK"+"."+"title", "body", o.Title); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDCalendarEventIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDCalendarEventIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDCalendarEventIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
