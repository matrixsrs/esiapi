// Code generated by go-swagger; DO NOT EDIT.

package universe

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

// GetUniverseBloodlinesReader is a Reader for the GetUniverseBloodlines structure.
type GetUniverseBloodlinesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseBloodlinesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUniverseBloodlinesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetUniverseBloodlinesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUniverseBloodlinesOK creates a GetUniverseBloodlinesOK with default headers values
func NewGetUniverseBloodlinesOK() *GetUniverseBloodlinesOK {
	return &GetUniverseBloodlinesOK{}
}

/*GetUniverseBloodlinesOK handles this case with default header values.

A list of bloodlines
*/
type GetUniverseBloodlinesOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*The language used in the response
	 */
	ContentLanguage string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []*GetUniverseBloodlinesOKBodyItems0
}

func (o *GetUniverseBloodlinesOK) Error() string {
	return fmt.Sprintf("[GET /universe/bloodlines/][%d] getUniverseBloodlinesOK  %+v", 200, o.Payload)
}

func (o *GetUniverseBloodlinesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Content-Language
	o.ContentLanguage = response.GetHeader("Content-Language")

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

// NewGetUniverseBloodlinesInternalServerError creates a GetUniverseBloodlinesInternalServerError with default headers values
func NewGetUniverseBloodlinesInternalServerError() *GetUniverseBloodlinesInternalServerError {
	return &GetUniverseBloodlinesInternalServerError{}
}

/*GetUniverseBloodlinesInternalServerError handles this case with default header values.

Internal server error
*/
type GetUniverseBloodlinesInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetUniverseBloodlinesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/bloodlines/][%d] getUniverseBloodlinesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseBloodlinesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetUniverseBloodlinesOKBodyItems0 get_universe_bloodlines_200_ok
//
// 200 ok object
swagger:model GetUniverseBloodlinesOKBodyItems0
*/

type GetUniverseBloodlinesOKBodyItems0 struct {

	// get_universe_bloodlines_bloodline_id
	//
	// bloodline_id integer
	// Required: true
	BloodlineID *int32 `json:"bloodline_id"`

	// get_universe_bloodlines_charisma
	//
	// charisma integer
	// Required: true
	Charisma *int32 `json:"charisma"`

	// get_universe_bloodlines_corporation_id
	//
	// corporation_id integer
	// Required: true
	CorporationID *int32 `json:"corporation_id"`

	// get_universe_bloodlines_description
	//
	// description string
	// Required: true
	Description *string `json:"description"`

	// get_universe_bloodlines_intelligence
	//
	// intelligence integer
	// Required: true
	Intelligence *int32 `json:"intelligence"`

	// get_universe_bloodlines_memory
	//
	// memory integer
	// Required: true
	Memory *int32 `json:"memory"`

	// get_universe_bloodlines_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// get_universe_bloodlines_perception
	//
	// perception integer
	// Required: true
	Perception *int32 `json:"perception"`

	// get_universe_bloodlines_race_id
	//
	// race_id integer
	// Required: true
	RaceID *int32 `json:"race_id"`

	// get_universe_bloodlines_ship_type_id
	//
	// ship_type_id integer
	// Required: true
	ShipTypeID *int32 `json:"ship_type_id"`

	// get_universe_bloodlines_willpower
	//
	// willpower integer
	// Required: true
	Willpower *int32 `json:"willpower"`
}

/* polymorph GetUniverseBloodlinesOKBodyItems0 bloodline_id false */

/* polymorph GetUniverseBloodlinesOKBodyItems0 charisma false */

/* polymorph GetUniverseBloodlinesOKBodyItems0 corporation_id false */

/* polymorph GetUniverseBloodlinesOKBodyItems0 description false */

/* polymorph GetUniverseBloodlinesOKBodyItems0 intelligence false */

/* polymorph GetUniverseBloodlinesOKBodyItems0 memory false */

/* polymorph GetUniverseBloodlinesOKBodyItems0 name false */

/* polymorph GetUniverseBloodlinesOKBodyItems0 perception false */

/* polymorph GetUniverseBloodlinesOKBodyItems0 race_id false */

/* polymorph GetUniverseBloodlinesOKBodyItems0 ship_type_id false */

/* polymorph GetUniverseBloodlinesOKBodyItems0 willpower false */

// Validate validates this get universe bloodlines o k body items0
func (o *GetUniverseBloodlinesOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBloodlineID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateCharisma(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateCorporationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateIntelligence(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateMemory(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validatePerception(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateRaceID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateShipTypeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateWillpower(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseBloodlinesOKBodyItems0) validateBloodlineID(formats strfmt.Registry) error {

	if err := validate.Required("bloodline_id", "body", o.BloodlineID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseBloodlinesOKBodyItems0) validateCharisma(formats strfmt.Registry) error {

	if err := validate.Required("charisma", "body", o.Charisma); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseBloodlinesOKBodyItems0) validateCorporationID(formats strfmt.Registry) error {

	if err := validate.Required("corporation_id", "body", o.CorporationID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseBloodlinesOKBodyItems0) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", o.Description); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseBloodlinesOKBodyItems0) validateIntelligence(formats strfmt.Registry) error {

	if err := validate.Required("intelligence", "body", o.Intelligence); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseBloodlinesOKBodyItems0) validateMemory(formats strfmt.Registry) error {

	if err := validate.Required("memory", "body", o.Memory); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseBloodlinesOKBodyItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseBloodlinesOKBodyItems0) validatePerception(formats strfmt.Registry) error {

	if err := validate.Required("perception", "body", o.Perception); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseBloodlinesOKBodyItems0) validateRaceID(formats strfmt.Registry) error {

	if err := validate.Required("race_id", "body", o.RaceID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseBloodlinesOKBodyItems0) validateShipTypeID(formats strfmt.Registry) error {

	if err := validate.Required("ship_type_id", "body", o.ShipTypeID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseBloodlinesOKBodyItems0) validateWillpower(formats strfmt.Registry) error {

	if err := validate.Required("willpower", "body", o.Willpower); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseBloodlinesOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseBloodlinesOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetUniverseBloodlinesOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
