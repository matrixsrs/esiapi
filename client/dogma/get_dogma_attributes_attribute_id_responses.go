// Code generated by go-swagger; DO NOT EDIT.

package dogma

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

// GetDogmaAttributesAttributeIDReader is a Reader for the GetDogmaAttributesAttributeID structure.
type GetDogmaAttributesAttributeIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDogmaAttributesAttributeIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDogmaAttributesAttributeIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetDogmaAttributesAttributeIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetDogmaAttributesAttributeIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDogmaAttributesAttributeIDOK creates a GetDogmaAttributesAttributeIDOK with default headers values
func NewGetDogmaAttributesAttributeIDOK() *GetDogmaAttributesAttributeIDOK {
	return &GetDogmaAttributesAttributeIDOK{}
}

/*GetDogmaAttributesAttributeIDOK handles this case with default header values.

Information about a dogma attribute
*/
type GetDogmaAttributesAttributeIDOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload GetDogmaAttributesAttributeIDOKBody
}

func (o *GetDogmaAttributesAttributeIDOK) Error() string {
	return fmt.Sprintf("[GET /dogma/attributes/{attribute_id}/][%d] getDogmaAttributesAttributeIdOK  %+v", 200, o.Payload)
}

func (o *GetDogmaAttributesAttributeIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDogmaAttributesAttributeIDNotFound creates a GetDogmaAttributesAttributeIDNotFound with default headers values
func NewGetDogmaAttributesAttributeIDNotFound() *GetDogmaAttributesAttributeIDNotFound {
	return &GetDogmaAttributesAttributeIDNotFound{}
}

/*GetDogmaAttributesAttributeIDNotFound handles this case with default header values.

Dogma attribute not found
*/
type GetDogmaAttributesAttributeIDNotFound struct {
	Payload GetDogmaAttributesAttributeIDNotFoundBody
}

func (o *GetDogmaAttributesAttributeIDNotFound) Error() string {
	return fmt.Sprintf("[GET /dogma/attributes/{attribute_id}/][%d] getDogmaAttributesAttributeIdNotFound  %+v", 404, o.Payload)
}

func (o *GetDogmaAttributesAttributeIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDogmaAttributesAttributeIDInternalServerError creates a GetDogmaAttributesAttributeIDInternalServerError with default headers values
func NewGetDogmaAttributesAttributeIDInternalServerError() *GetDogmaAttributesAttributeIDInternalServerError {
	return &GetDogmaAttributesAttributeIDInternalServerError{}
}

/*GetDogmaAttributesAttributeIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetDogmaAttributesAttributeIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetDogmaAttributesAttributeIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dogma/attributes/{attribute_id}/][%d] getDogmaAttributesAttributeIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDogmaAttributesAttributeIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetDogmaAttributesAttributeIDNotFoundBody get_dogma_attributes_attribute_id_not_found
//
// Not found
swagger:model GetDogmaAttributesAttributeIDNotFoundBody
*/

type GetDogmaAttributesAttributeIDNotFoundBody struct {

	// get_dogma_attributes_attribute_id_404_not_found
	//
	// Not found message
	// Required: true
	Error *string `json:"error"`
}

/* polymorph GetDogmaAttributesAttributeIDNotFoundBody error false */

// Validate validates this get dogma attributes attribute ID not found body
func (o *GetDogmaAttributesAttributeIDNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *GetDogmaAttributesAttributeIDNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getDogmaAttributesAttributeIdNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDogmaAttributesAttributeIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDogmaAttributesAttributeIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetDogmaAttributesAttributeIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetDogmaAttributesAttributeIDOKBody get_dogma_attributes_attribute_id_ok
//
// 200 ok object
swagger:model GetDogmaAttributesAttributeIDOKBody
*/

type GetDogmaAttributesAttributeIDOKBody struct {

	// get_dogma_attributes_attribute_id_attribute_id
	//
	// attribute_id integer
	// Required: true
	AttributeID *int32 `json:"attribute_id"`

	// get_dogma_attributes_attribute_id_default_value
	//
	// default_value number
	// Required: true
	DefaultValue *float32 `json:"default_value"`

	// get_dogma_attributes_attribute_id_description
	//
	// description string
	// Required: true
	Description *string `json:"description"`

	// get_dogma_attributes_attribute_id_display_name
	//
	// display_name string
	// Required: true
	DisplayName *string `json:"display_name"`

	// get_dogma_attributes_attribute_id_high_is_good
	//
	// high_is_good boolean
	// Required: true
	HighIsGood *bool `json:"high_is_good"`

	// get_dogma_attributes_attribute_id_icon_id
	//
	// icon_id integer
	// Required: true
	IconID *int32 `json:"icon_id"`

	// get_dogma_attributes_attribute_id_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// get_dogma_attributes_attribute_id_published
	//
	// published boolean
	// Required: true
	Published *bool `json:"published"`

	// get_dogma_attributes_attribute_id_stackable
	//
	// stackable boolean
	// Required: true
	Stackable *bool `json:"stackable"`

	// get_dogma_attributes_attribute_id_unit_id
	//
	// unit_id integer
	// Required: true
	UnitID *int32 `json:"unit_id"`
}

/* polymorph GetDogmaAttributesAttributeIDOKBody attribute_id false */

/* polymorph GetDogmaAttributesAttributeIDOKBody default_value false */

/* polymorph GetDogmaAttributesAttributeIDOKBody description false */

/* polymorph GetDogmaAttributesAttributeIDOKBody display_name false */

/* polymorph GetDogmaAttributesAttributeIDOKBody high_is_good false */

/* polymorph GetDogmaAttributesAttributeIDOKBody icon_id false */

/* polymorph GetDogmaAttributesAttributeIDOKBody name false */

/* polymorph GetDogmaAttributesAttributeIDOKBody published false */

/* polymorph GetDogmaAttributesAttributeIDOKBody stackable false */

/* polymorph GetDogmaAttributesAttributeIDOKBody unit_id false */

// Validate validates this get dogma attributes attribute ID o k body
func (o *GetDogmaAttributesAttributeIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAttributeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateDefaultValue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateDisplayName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateHighIsGood(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateIconID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validatePublished(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateStackable(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateUnitID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDogmaAttributesAttributeIDOKBody) validateAttributeID(formats strfmt.Registry) error {

	if err := validate.Required("getDogmaAttributesAttributeIdOK"+"."+"attribute_id", "body", o.AttributeID); err != nil {
		return err
	}

	return nil
}

func (o *GetDogmaAttributesAttributeIDOKBody) validateDefaultValue(formats strfmt.Registry) error {

	if err := validate.Required("getDogmaAttributesAttributeIdOK"+"."+"default_value", "body", o.DefaultValue); err != nil {
		return err
	}

	return nil
}

func (o *GetDogmaAttributesAttributeIDOKBody) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("getDogmaAttributesAttributeIdOK"+"."+"description", "body", o.Description); err != nil {
		return err
	}

	return nil
}

func (o *GetDogmaAttributesAttributeIDOKBody) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("getDogmaAttributesAttributeIdOK"+"."+"display_name", "body", o.DisplayName); err != nil {
		return err
	}

	return nil
}

func (o *GetDogmaAttributesAttributeIDOKBody) validateHighIsGood(formats strfmt.Registry) error {

	if err := validate.Required("getDogmaAttributesAttributeIdOK"+"."+"high_is_good", "body", o.HighIsGood); err != nil {
		return err
	}

	return nil
}

func (o *GetDogmaAttributesAttributeIDOKBody) validateIconID(formats strfmt.Registry) error {

	if err := validate.Required("getDogmaAttributesAttributeIdOK"+"."+"icon_id", "body", o.IconID); err != nil {
		return err
	}

	return nil
}

func (o *GetDogmaAttributesAttributeIDOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("getDogmaAttributesAttributeIdOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetDogmaAttributesAttributeIDOKBody) validatePublished(formats strfmt.Registry) error {

	if err := validate.Required("getDogmaAttributesAttributeIdOK"+"."+"published", "body", o.Published); err != nil {
		return err
	}

	return nil
}

func (o *GetDogmaAttributesAttributeIDOKBody) validateStackable(formats strfmt.Registry) error {

	if err := validate.Required("getDogmaAttributesAttributeIdOK"+"."+"stackable", "body", o.Stackable); err != nil {
		return err
	}

	return nil
}

func (o *GetDogmaAttributesAttributeIDOKBody) validateUnitID(formats strfmt.Registry) error {

	if err := validate.Required("getDogmaAttributesAttributeIdOK"+"."+"unit_id", "body", o.UnitID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDogmaAttributesAttributeIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDogmaAttributesAttributeIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetDogmaAttributesAttributeIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
