// Code generated by go-swagger; DO NOT EDIT.

package corporation

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

// GetCorporationsCorporationIDIconsReader is a Reader for the GetCorporationsCorporationIDIcons structure.
type GetCorporationsCorporationIDIconsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDIconsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCorporationsCorporationIDIconsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetCorporationsCorporationIDIconsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCorporationsCorporationIDIconsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDIconsOK creates a GetCorporationsCorporationIDIconsOK with default headers values
func NewGetCorporationsCorporationIDIconsOK() *GetCorporationsCorporationIDIconsOK {
	return &GetCorporationsCorporationIDIconsOK{}
}

/*GetCorporationsCorporationIDIconsOK handles this case with default header values.

Urls for icons for the given corporation id and server
*/
type GetCorporationsCorporationIDIconsOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload GetCorporationsCorporationIDIconsOKBody
}

func (o *GetCorporationsCorporationIDIconsOK) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/icons/][%d] getCorporationsCorporationIdIconsOK  %+v", 200, o.Payload)
}

func (o *GetCorporationsCorporationIDIconsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDIconsNotFound creates a GetCorporationsCorporationIDIconsNotFound with default headers values
func NewGetCorporationsCorporationIDIconsNotFound() *GetCorporationsCorporationIDIconsNotFound {
	return &GetCorporationsCorporationIDIconsNotFound{}
}

/*GetCorporationsCorporationIDIconsNotFound handles this case with default header values.

No image server for this datasource
*/
type GetCorporationsCorporationIDIconsNotFound struct {
	Payload GetCorporationsCorporationIDIconsNotFoundBody
}

func (o *GetCorporationsCorporationIDIconsNotFound) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/icons/][%d] getCorporationsCorporationIdIconsNotFound  %+v", 404, o.Payload)
}

func (o *GetCorporationsCorporationIDIconsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDIconsInternalServerError creates a GetCorporationsCorporationIDIconsInternalServerError with default headers values
func NewGetCorporationsCorporationIDIconsInternalServerError() *GetCorporationsCorporationIDIconsInternalServerError {
	return &GetCorporationsCorporationIDIconsInternalServerError{}
}

/*GetCorporationsCorporationIDIconsInternalServerError handles this case with default header values.

Internal server error
*/
type GetCorporationsCorporationIDIconsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCorporationsCorporationIDIconsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/icons/][%d] getCorporationsCorporationIdIconsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationsCorporationIDIconsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCorporationsCorporationIDIconsNotFoundBody get_corporations_corporation_id_icons_not_found
//
// No image server for this datasource
swagger:model GetCorporationsCorporationIDIconsNotFoundBody
*/

type GetCorporationsCorporationIDIconsNotFoundBody struct {

	// get_corporations_corporation_id_icons_error
	//
	// error message
	// Required: true
	Error *string `json:"error"`
}

/* polymorph GetCorporationsCorporationIDIconsNotFoundBody error false */

// Validate validates this get corporations corporation ID icons not found body
func (o *GetCorporationsCorporationIDIconsNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *GetCorporationsCorporationIDIconsNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getCorporationsCorporationIdIconsNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCorporationsCorporationIDIconsNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCorporationsCorporationIDIconsNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDIconsNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetCorporationsCorporationIDIconsOKBody get_corporations_corporation_id_icons_ok
//
// 200 ok object
swagger:model GetCorporationsCorporationIDIconsOKBody
*/

type GetCorporationsCorporationIDIconsOKBody struct {

	// get_corporations_corporation_id_icons_px128x128
	//
	// px128x128 string
	// Required: true
	Px128x128 *string `json:"px128x128"`

	// get_corporations_corporation_id_icons_px256x256
	//
	// px256x256 string
	// Required: true
	Px256x256 *string `json:"px256x256"`

	// get_corporations_corporation_id_icons_px64x64
	//
	// px64x64 string
	// Required: true
	Px64x64 *string `json:"px64x64"`
}

/* polymorph GetCorporationsCorporationIDIconsOKBody px128x128 false */

/* polymorph GetCorporationsCorporationIDIconsOKBody px256x256 false */

/* polymorph GetCorporationsCorporationIDIconsOKBody px64x64 false */

// Validate validates this get corporations corporation ID icons o k body
func (o *GetCorporationsCorporationIDIconsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePx128x128(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validatePx256x256(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validatePx64x64(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCorporationsCorporationIDIconsOKBody) validatePx128x128(formats strfmt.Registry) error {

	if err := validate.Required("getCorporationsCorporationIdIconsOK"+"."+"px128x128", "body", o.Px128x128); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDIconsOKBody) validatePx256x256(formats strfmt.Registry) error {

	if err := validate.Required("getCorporationsCorporationIdIconsOK"+"."+"px256x256", "body", o.Px256x256); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDIconsOKBody) validatePx64x64(formats strfmt.Registry) error {

	if err := validate.Required("getCorporationsCorporationIdIconsOK"+"."+"px64x64", "body", o.Px64x64); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCorporationsCorporationIDIconsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCorporationsCorporationIDIconsOKBody) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDIconsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
