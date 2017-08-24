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

// GetUniverseTypesTypeIDReader is a Reader for the GetUniverseTypesTypeID structure.
type GetUniverseTypesTypeIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseTypesTypeIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUniverseTypesTypeIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetUniverseTypesTypeIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetUniverseTypesTypeIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUniverseTypesTypeIDOK creates a GetUniverseTypesTypeIDOK with default headers values
func NewGetUniverseTypesTypeIDOK() *GetUniverseTypesTypeIDOK {
	return &GetUniverseTypesTypeIDOK{}
}

/*GetUniverseTypesTypeIDOK handles this case with default header values.

Information about a type
*/
type GetUniverseTypesTypeIDOK struct {
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

	Payload GetUniverseTypesTypeIDOKBody
}

func (o *GetUniverseTypesTypeIDOK) Error() string {
	return fmt.Sprintf("[GET /universe/types/{type_id}/][%d] getUniverseTypesTypeIdOK  %+v", 200, o.Payload)
}

func (o *GetUniverseTypesTypeIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseTypesTypeIDNotFound creates a GetUniverseTypesTypeIDNotFound with default headers values
func NewGetUniverseTypesTypeIDNotFound() *GetUniverseTypesTypeIDNotFound {
	return &GetUniverseTypesTypeIDNotFound{}
}

/*GetUniverseTypesTypeIDNotFound handles this case with default header values.

Type not found
*/
type GetUniverseTypesTypeIDNotFound struct {
	Payload GetUniverseTypesTypeIDNotFoundBody
}

func (o *GetUniverseTypesTypeIDNotFound) Error() string {
	return fmt.Sprintf("[GET /universe/types/{type_id}/][%d] getUniverseTypesTypeIdNotFound  %+v", 404, o.Payload)
}

func (o *GetUniverseTypesTypeIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseTypesTypeIDInternalServerError creates a GetUniverseTypesTypeIDInternalServerError with default headers values
func NewGetUniverseTypesTypeIDInternalServerError() *GetUniverseTypesTypeIDInternalServerError {
	return &GetUniverseTypesTypeIDInternalServerError{}
}

/*GetUniverseTypesTypeIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetUniverseTypesTypeIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetUniverseTypesTypeIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/types/{type_id}/][%d] getUniverseTypesTypeIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseTypesTypeIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DogmaAttributesItems0 get_universe_types_type_id_dogma_attribute
//
// dogma_attribute object
swagger:model DogmaAttributesItems0
*/

type DogmaAttributesItems0 struct {

	// get_universe_types_type_id_attribute_id
	//
	// attribute_id integer
	// Required: true
	AttributeID *int32 `json:"attribute_id"`

	// get_universe_types_type_id_value
	//
	// value number
	// Required: true
	Value *float32 `json:"value"`
}

/* polymorph DogmaAttributesItems0 attribute_id false */

/* polymorph DogmaAttributesItems0 value false */

// Validate validates this dogma attributes items0
func (o *DogmaAttributesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAttributeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateValue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DogmaAttributesItems0) validateAttributeID(formats strfmt.Registry) error {

	if err := validate.Required("attribute_id", "body", o.AttributeID); err != nil {
		return err
	}

	return nil
}

func (o *DogmaAttributesItems0) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DogmaAttributesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DogmaAttributesItems0) UnmarshalBinary(b []byte) error {
	var res DogmaAttributesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DogmaEffectsItems0 get_universe_types_type_id_dogma_effect
//
// dogma_effect object
swagger:model DogmaEffectsItems0
*/

type DogmaEffectsItems0 struct {

	// get_universe_types_type_id_effect_id
	//
	// effect_id integer
	// Required: true
	EffectID *int32 `json:"effect_id"`

	// get_universe_types_type_id_is_default
	//
	// is_default boolean
	// Required: true
	IsDefault *bool `json:"is_default"`
}

/* polymorph DogmaEffectsItems0 effect_id false */

/* polymorph DogmaEffectsItems0 is_default false */

// Validate validates this dogma effects items0
func (o *DogmaEffectsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEffectID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateIsDefault(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DogmaEffectsItems0) validateEffectID(formats strfmt.Registry) error {

	if err := validate.Required("effect_id", "body", o.EffectID); err != nil {
		return err
	}

	return nil
}

func (o *DogmaEffectsItems0) validateIsDefault(formats strfmt.Registry) error {

	if err := validate.Required("is_default", "body", o.IsDefault); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DogmaEffectsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DogmaEffectsItems0) UnmarshalBinary(b []byte) error {
	var res DogmaEffectsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetUniverseTypesTypeIDNotFoundBody get_universe_types_type_id_not_found
//
// Not found
swagger:model GetUniverseTypesTypeIDNotFoundBody
*/

type GetUniverseTypesTypeIDNotFoundBody struct {

	// get_universe_types_type_id_404_not_found
	//
	// Not found message
	// Required: true
	Error *string `json:"error"`
}

/* polymorph GetUniverseTypesTypeIDNotFoundBody error false */

// Validate validates this get universe types type ID not found body
func (o *GetUniverseTypesTypeIDNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *GetUniverseTypesTypeIDNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseTypesTypeIdNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseTypesTypeIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseTypesTypeIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseTypesTypeIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetUniverseTypesTypeIDOKBody get_universe_types_type_id_ok
//
// 200 ok object
swagger:model GetUniverseTypesTypeIDOKBody
*/

type GetUniverseTypesTypeIDOKBody struct {

	// get_universe_types_type_id_capacity
	//
	// capacity number
	// Required: true
	Capacity *float32 `json:"capacity"`

	// get_universe_types_type_id_description
	//
	// description string
	// Required: true
	Description *string `json:"description"`

	// get_universe_types_type_id_dogma_attributes
	//
	// dogma_attributes array
	// Required: true
	// Max Items: 1000
	DogmaAttributes []*DogmaAttributesItems0 `json:"dogma_attributes"`

	// get_universe_types_type_id_dogma_effects
	//
	// dogma_effects array
	// Required: true
	// Max Items: 1000
	DogmaEffects []*DogmaEffectsItems0 `json:"dogma_effects"`

	// get_universe_types_type_id_graphic_id
	//
	// graphic_id integer
	// Required: true
	GraphicID *int32 `json:"graphic_id"`

	// get_universe_types_type_id_group_id
	//
	// group_id integer
	// Required: true
	GroupID *int32 `json:"group_id"`

	// get_universe_types_type_id_icon_id
	//
	// icon_id integer
	// Required: true
	IconID *int32 `json:"icon_id"`

	// get_universe_types_type_id_mass
	//
	// mass number
	// Required: true
	Mass *float32 `json:"mass"`

	// get_universe_types_type_id_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// get_universe_types_type_id_portion_size
	//
	// portion_size integer
	// Required: true
	PortionSize *int32 `json:"portion_size"`

	// get_universe_types_type_id_published
	//
	// published boolean
	// Required: true
	Published *bool `json:"published"`

	// get_universe_types_type_id_radius
	//
	// radius number
	// Required: true
	Radius *float32 `json:"radius"`

	// get_universe_types_type_id_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`

	// get_universe_types_type_id_volume
	//
	// volume number
	// Required: true
	Volume *float32 `json:"volume"`
}

/* polymorph GetUniverseTypesTypeIDOKBody capacity false */

/* polymorph GetUniverseTypesTypeIDOKBody description false */

/* polymorph GetUniverseTypesTypeIDOKBody dogma_attributes false */

/* polymorph GetUniverseTypesTypeIDOKBody dogma_effects false */

/* polymorph GetUniverseTypesTypeIDOKBody graphic_id false */

/* polymorph GetUniverseTypesTypeIDOKBody group_id false */

/* polymorph GetUniverseTypesTypeIDOKBody icon_id false */

/* polymorph GetUniverseTypesTypeIDOKBody mass false */

/* polymorph GetUniverseTypesTypeIDOKBody name false */

/* polymorph GetUniverseTypesTypeIDOKBody portion_size false */

/* polymorph GetUniverseTypesTypeIDOKBody published false */

/* polymorph GetUniverseTypesTypeIDOKBody radius false */

/* polymorph GetUniverseTypesTypeIDOKBody type_id false */

/* polymorph GetUniverseTypesTypeIDOKBody volume false */

// Validate validates this get universe types type ID o k body
func (o *GetUniverseTypesTypeIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCapacity(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateDogmaAttributes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateDogmaEffects(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateGraphicID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateGroupID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateIconID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateMass(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validatePortionSize(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validatePublished(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateRadius(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateTypeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateVolume(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseTypesTypeIDOKBody) validateCapacity(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseTypesTypeIdOK"+"."+"capacity", "body", o.Capacity); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseTypesTypeIDOKBody) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseTypesTypeIdOK"+"."+"description", "body", o.Description); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseTypesTypeIDOKBody) validateDogmaAttributes(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseTypesTypeIdOK"+"."+"dogma_attributes", "body", o.DogmaAttributes); err != nil {
		return err
	}

	iDogmaAttributesSize := int64(len(o.DogmaAttributes))

	if err := validate.MaxItems("getUniverseTypesTypeIdOK"+"."+"dogma_attributes", "body", iDogmaAttributesSize, 1000); err != nil {
		return err
	}

	for i := 0; i < len(o.DogmaAttributes); i++ {

		if swag.IsZero(o.DogmaAttributes[i]) { // not required
			continue
		}

		if o.DogmaAttributes[i] != nil {

			if err := o.DogmaAttributes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getUniverseTypesTypeIdOK" + "." + "dogma_attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetUniverseTypesTypeIDOKBody) validateDogmaEffects(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseTypesTypeIdOK"+"."+"dogma_effects", "body", o.DogmaEffects); err != nil {
		return err
	}

	iDogmaEffectsSize := int64(len(o.DogmaEffects))

	if err := validate.MaxItems("getUniverseTypesTypeIdOK"+"."+"dogma_effects", "body", iDogmaEffectsSize, 1000); err != nil {
		return err
	}

	for i := 0; i < len(o.DogmaEffects); i++ {

		if swag.IsZero(o.DogmaEffects[i]) { // not required
			continue
		}

		if o.DogmaEffects[i] != nil {

			if err := o.DogmaEffects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getUniverseTypesTypeIdOK" + "." + "dogma_effects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetUniverseTypesTypeIDOKBody) validateGraphicID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseTypesTypeIdOK"+"."+"graphic_id", "body", o.GraphicID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseTypesTypeIDOKBody) validateGroupID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseTypesTypeIdOK"+"."+"group_id", "body", o.GroupID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseTypesTypeIDOKBody) validateIconID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseTypesTypeIdOK"+"."+"icon_id", "body", o.IconID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseTypesTypeIDOKBody) validateMass(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseTypesTypeIdOK"+"."+"mass", "body", o.Mass); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseTypesTypeIDOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseTypesTypeIdOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseTypesTypeIDOKBody) validatePortionSize(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseTypesTypeIdOK"+"."+"portion_size", "body", o.PortionSize); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseTypesTypeIDOKBody) validatePublished(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseTypesTypeIdOK"+"."+"published", "body", o.Published); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseTypesTypeIDOKBody) validateRadius(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseTypesTypeIdOK"+"."+"radius", "body", o.Radius); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseTypesTypeIDOKBody) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseTypesTypeIdOK"+"."+"type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseTypesTypeIDOKBody) validateVolume(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseTypesTypeIdOK"+"."+"volume", "body", o.Volume); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseTypesTypeIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseTypesTypeIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseTypesTypeIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
