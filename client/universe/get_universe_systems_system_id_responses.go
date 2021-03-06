// Code generated by go-swagger; DO NOT EDIT.

package universe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/models"
)

// GetUniverseSystemsSystemIDReader is a Reader for the GetUniverseSystemsSystemID structure.
type GetUniverseSystemsSystemIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseSystemsSystemIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUniverseSystemsSystemIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetUniverseSystemsSystemIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetUniverseSystemsSystemIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUniverseSystemsSystemIDOK creates a GetUniverseSystemsSystemIDOK with default headers values
func NewGetUniverseSystemsSystemIDOK() *GetUniverseSystemsSystemIDOK {
	return &GetUniverseSystemsSystemIDOK{}
}

/*GetUniverseSystemsSystemIDOK handles this case with default header values.

Information about a solar system
*/
type GetUniverseSystemsSystemIDOK struct {
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

	Payload GetUniverseSystemsSystemIDOKBody
}

func (o *GetUniverseSystemsSystemIDOK) Error() string {
	return fmt.Sprintf("[GET /universe/systems/{system_id}/][%d] getUniverseSystemsSystemIdOK  %+v", 200, o.Payload)
}

func (o *GetUniverseSystemsSystemIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseSystemsSystemIDNotFound creates a GetUniverseSystemsSystemIDNotFound with default headers values
func NewGetUniverseSystemsSystemIDNotFound() *GetUniverseSystemsSystemIDNotFound {
	return &GetUniverseSystemsSystemIDNotFound{}
}

/*GetUniverseSystemsSystemIDNotFound handles this case with default header values.

Solar system not found
*/
type GetUniverseSystemsSystemIDNotFound struct {
	Payload GetUniverseSystemsSystemIDNotFoundBody
}

func (o *GetUniverseSystemsSystemIDNotFound) Error() string {
	return fmt.Sprintf("[GET /universe/systems/{system_id}/][%d] getUniverseSystemsSystemIdNotFound  %+v", 404, o.Payload)
}

func (o *GetUniverseSystemsSystemIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseSystemsSystemIDInternalServerError creates a GetUniverseSystemsSystemIDInternalServerError with default headers values
func NewGetUniverseSystemsSystemIDInternalServerError() *GetUniverseSystemsSystemIDInternalServerError {
	return &GetUniverseSystemsSystemIDInternalServerError{}
}

/*GetUniverseSystemsSystemIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetUniverseSystemsSystemIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetUniverseSystemsSystemIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/systems/{system_id}/][%d] getUniverseSystemsSystemIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseSystemsSystemIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetUniverseSystemsSystemIDNotFoundBody get_universe_systems_system_id_not_found
//
// Not found
swagger:model GetUniverseSystemsSystemIDNotFoundBody
*/

type GetUniverseSystemsSystemIDNotFoundBody struct {

	// get_universe_systems_system_id_404_not_found
	//
	// Not found message
	// Required: true
	Error *string `json:"error"`
}

/* polymorph GetUniverseSystemsSystemIDNotFoundBody error false */

// Validate validates this get universe systems system ID not found body
func (o *GetUniverseSystemsSystemIDNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseSystemsSystemIDNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseSystemsSystemIdNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseSystemsSystemIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseSystemsSystemIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseSystemsSystemIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetUniverseSystemsSystemIDOKBody get_universe_systems_system_id_ok
//
// 200 ok object
swagger:model GetUniverseSystemsSystemIDOKBody
*/

type GetUniverseSystemsSystemIDOKBody struct {

	// get_universe_systems_system_id_constellation_id
	//
	// The constellation this solar system is in
	// Required: true
	ConstellationID *int32 `json:"constellation_id"`

	// get_universe_systems_system_id_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// get_universe_systems_system_id_planets
	//
	// planets array
	// Required: true
	// Max Items: 1000
	Planets []*PlanetsItems0 `json:"planets"`

	// position
	// Required: true
	Position *GetUniverseSystemsSystemIDOKBodyPosition `json:"position"`

	// get_universe_systems_system_id_security_class
	//
	// security_class string
	// Required: true
	SecurityClass *string `json:"security_class"`

	// get_universe_systems_system_id_security_status
	//
	// security_status number
	// Required: true
	SecurityStatus *float32 `json:"security_status"`

	// get_universe_systems_system_id_star_id
	//
	// star_id integer
	// Required: true
	StarID *int32 `json:"star_id"`

	// get_universe_systems_system_id_stargates
	//
	// stargates array
	// Required: true
	// Max Items: 25
	Stargates []int32 `json:"stargates"`

	// get_universe_systems_system_id_system_id
	//
	// system_id integer
	// Required: true
	SystemID *int32 `json:"system_id"`
}

/* polymorph GetUniverseSystemsSystemIDOKBody constellation_id false */

/* polymorph GetUniverseSystemsSystemIDOKBody name false */

/* polymorph GetUniverseSystemsSystemIDOKBody planets false */

/* polymorph GetUniverseSystemsSystemIDOKBody position false */

/* polymorph GetUniverseSystemsSystemIDOKBody security_class false */

/* polymorph GetUniverseSystemsSystemIDOKBody security_status false */

/* polymorph GetUniverseSystemsSystemIDOKBody star_id false */

/* polymorph GetUniverseSystemsSystemIDOKBody stargates false */

/* polymorph GetUniverseSystemsSystemIDOKBody system_id false */

// Validate validates this get universe systems system ID o k body
func (o *GetUniverseSystemsSystemIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateConstellationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validatePlanets(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validatePosition(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateSecurityClass(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateSecurityStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateStarID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateStargates(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateSystemID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseSystemsSystemIDOKBody) validateConstellationID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseSystemsSystemIdOK"+"."+"constellation_id", "body", o.ConstellationID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseSystemsSystemIDOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseSystemsSystemIdOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseSystemsSystemIDOKBody) validatePlanets(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseSystemsSystemIdOK"+"."+"planets", "body", o.Planets); err != nil {
		return err
	}

	iPlanetsSize := int64(len(o.Planets))

	if err := validate.MaxItems("getUniverseSystemsSystemIdOK"+"."+"planets", "body", iPlanetsSize, 1000); err != nil {
		return err
	}

	for i := 0; i < len(o.Planets); i++ {

		if swag.IsZero(o.Planets[i]) { // not required
			continue
		}

		if o.Planets[i] != nil {

			if err := o.Planets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getUniverseSystemsSystemIdOK" + "." + "planets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetUniverseSystemsSystemIDOKBody) validatePosition(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseSystemsSystemIdOK"+"."+"position", "body", o.Position); err != nil {
		return err
	}

	if o.Position != nil {

		if err := o.Position.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getUniverseSystemsSystemIdOK" + "." + "position")
			}
			return err
		}
	}

	return nil
}

func (o *GetUniverseSystemsSystemIDOKBody) validateSecurityClass(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseSystemsSystemIdOK"+"."+"security_class", "body", o.SecurityClass); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseSystemsSystemIDOKBody) validateSecurityStatus(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseSystemsSystemIdOK"+"."+"security_status", "body", o.SecurityStatus); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseSystemsSystemIDOKBody) validateStarID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseSystemsSystemIdOK"+"."+"star_id", "body", o.StarID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseSystemsSystemIDOKBody) validateStargates(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseSystemsSystemIdOK"+"."+"stargates", "body", o.Stargates); err != nil {
		return err
	}

	iStargatesSize := int64(len(o.Stargates))

	if err := validate.MaxItems("getUniverseSystemsSystemIdOK"+"."+"stargates", "body", iStargatesSize, 25); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseSystemsSystemIDOKBody) validateSystemID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseSystemsSystemIdOK"+"."+"system_id", "body", o.SystemID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseSystemsSystemIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseSystemsSystemIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseSystemsSystemIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetUniverseSystemsSystemIDOKBodyPosition get_universe_systems_system_id_position
//
// position object
swagger:model GetUniverseSystemsSystemIDOKBodyPosition
*/

type GetUniverseSystemsSystemIDOKBodyPosition struct {

	// get_universe_systems_system_id_x
	//
	// x number
	// Required: true
	X *float32 `json:"x"`

	// get_universe_systems_system_id_y
	//
	// y number
	// Required: true
	Y *float32 `json:"y"`

	// get_universe_systems_system_id_z
	//
	// z number
	// Required: true
	Z *float32 `json:"z"`
}

/* polymorph GetUniverseSystemsSystemIDOKBodyPosition x false */

/* polymorph GetUniverseSystemsSystemIDOKBodyPosition y false */

/* polymorph GetUniverseSystemsSystemIDOKBodyPosition z false */

// Validate validates this get universe systems system ID o k body position
func (o *GetUniverseSystemsSystemIDOKBodyPosition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateX(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateY(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateZ(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseSystemsSystemIDOKBodyPosition) validateX(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseSystemsSystemIdOK"+"."+"position"+"."+"x", "body", o.X); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseSystemsSystemIDOKBodyPosition) validateY(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseSystemsSystemIdOK"+"."+"position"+"."+"y", "body", o.Y); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseSystemsSystemIDOKBodyPosition) validateZ(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseSystemsSystemIdOK"+"."+"position"+"."+"z", "body", o.Z); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseSystemsSystemIDOKBodyPosition) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseSystemsSystemIDOKBodyPosition) UnmarshalBinary(b []byte) error {
	var res GetUniverseSystemsSystemIDOKBodyPosition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PlanetsItems0 get_universe_systems_system_id_planet
//
// planet object
swagger:model PlanetsItems0
*/

type PlanetsItems0 struct {

	// get_universe_systems_system_id_moons
	//
	// moons array
	// Max Items: 1000
	Moons []int32 `json:"moons"`

	// get_universe_systems_system_id_planet_id
	//
	// planet_id integer
	// Required: true
	PlanetID *int32 `json:"planet_id"`
}

/* polymorph PlanetsItems0 moons false */

/* polymorph PlanetsItems0 planet_id false */

// Validate validates this planets items0
func (o *PlanetsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMoons(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validatePlanetID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PlanetsItems0) validateMoons(formats strfmt.Registry) error {

	if swag.IsZero(o.Moons) { // not required
		return nil
	}

	iMoonsSize := int64(len(o.Moons))

	if err := validate.MaxItems("moons", "body", iMoonsSize, 1000); err != nil {
		return err
	}

	return nil
}

func (o *PlanetsItems0) validatePlanetID(formats strfmt.Registry) error {

	if err := validate.Required("planet_id", "body", o.PlanetID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PlanetsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PlanetsItems0) UnmarshalBinary(b []byte) error {
	var res PlanetsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
