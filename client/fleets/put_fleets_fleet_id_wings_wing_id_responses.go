// Code generated by go-swagger; DO NOT EDIT.

package fleets

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

// PutFleetsFleetIDWingsWingIDReader is a Reader for the PutFleetsFleetIDWingsWingID structure.
type PutFleetsFleetIDWingsWingIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutFleetsFleetIDWingsWingIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPutFleetsFleetIDWingsWingIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewPutFleetsFleetIDWingsWingIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutFleetsFleetIDWingsWingIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPutFleetsFleetIDWingsWingIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutFleetsFleetIDWingsWingIDNoContent creates a PutFleetsFleetIDWingsWingIDNoContent with default headers values
func NewPutFleetsFleetIDWingsWingIDNoContent() *PutFleetsFleetIDWingsWingIDNoContent {
	return &PutFleetsFleetIDWingsWingIDNoContent{}
}

/*PutFleetsFleetIDWingsWingIDNoContent handles this case with default header values.

Wing renamed
*/
type PutFleetsFleetIDWingsWingIDNoContent struct {
}

func (o *PutFleetsFleetIDWingsWingIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/wings/{wing_id}/][%d] putFleetsFleetIdWingsWingIdNoContent ", 204)
}

func (o *PutFleetsFleetIDWingsWingIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutFleetsFleetIDWingsWingIDForbidden creates a PutFleetsFleetIDWingsWingIDForbidden with default headers values
func NewPutFleetsFleetIDWingsWingIDForbidden() *PutFleetsFleetIDWingsWingIDForbidden {
	return &PutFleetsFleetIDWingsWingIDForbidden{}
}

/*PutFleetsFleetIDWingsWingIDForbidden handles this case with default header values.

Forbidden
*/
type PutFleetsFleetIDWingsWingIDForbidden struct {
	Payload *models.Forbidden
}

func (o *PutFleetsFleetIDWingsWingIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/wings/{wing_id}/][%d] putFleetsFleetIdWingsWingIdForbidden  %+v", 403, o.Payload)
}

func (o *PutFleetsFleetIDWingsWingIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDWingsWingIDNotFound creates a PutFleetsFleetIDWingsWingIDNotFound with default headers values
func NewPutFleetsFleetIDWingsWingIDNotFound() *PutFleetsFleetIDWingsWingIDNotFound {
	return &PutFleetsFleetIDWingsWingIDNotFound{}
}

/*PutFleetsFleetIDWingsWingIDNotFound handles this case with default header values.

The fleet does not exist or you don't have access to it
*/
type PutFleetsFleetIDWingsWingIDNotFound struct {
	Payload PutFleetsFleetIDWingsWingIDNotFoundBody
}

func (o *PutFleetsFleetIDWingsWingIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/wings/{wing_id}/][%d] putFleetsFleetIdWingsWingIdNotFound  %+v", 404, o.Payload)
}

func (o *PutFleetsFleetIDWingsWingIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDWingsWingIDInternalServerError creates a PutFleetsFleetIDWingsWingIDInternalServerError with default headers values
func NewPutFleetsFleetIDWingsWingIDInternalServerError() *PutFleetsFleetIDWingsWingIDInternalServerError {
	return &PutFleetsFleetIDWingsWingIDInternalServerError{}
}

/*PutFleetsFleetIDWingsWingIDInternalServerError handles this case with default header values.

Internal server error
*/
type PutFleetsFleetIDWingsWingIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *PutFleetsFleetIDWingsWingIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/wings/{wing_id}/][%d] putFleetsFleetIdWingsWingIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PutFleetsFleetIDWingsWingIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PutFleetsFleetIDWingsWingIDBody put_fleets_fleet_id_wings_wing_id_naming
//
// naming object
swagger:model PutFleetsFleetIDWingsWingIDBody
*/

type PutFleetsFleetIDWingsWingIDBody struct {

	// put_fleets_fleet_id_wings_wing_id_name
	//
	// name string
	// Required: true
	// Max Length: 10
	Name *string `json:"name"`
}

/* polymorph PutFleetsFleetIDWingsWingIDBody name false */

// MarshalBinary interface implementation
func (o *PutFleetsFleetIDWingsWingIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutFleetsFleetIDWingsWingIDBody) UnmarshalBinary(b []byte) error {
	var res PutFleetsFleetIDWingsWingIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PutFleetsFleetIDWingsWingIDNotFoundBody put_fleets_fleet_id_wings_wing_id_not_found
//
// Not found
swagger:model PutFleetsFleetIDWingsWingIDNotFoundBody
*/

type PutFleetsFleetIDWingsWingIDNotFoundBody struct {

	// put_fleets_fleet_id_wings_wing_id_404_not_found
	//
	// Not found message
	// Required: true
	Error *string `json:"error"`
}

/* polymorph PutFleetsFleetIDWingsWingIDNotFoundBody error false */

// Validate validates this put fleets fleet ID wings wing ID not found body
func (o *PutFleetsFleetIDWingsWingIDNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *PutFleetsFleetIDWingsWingIDNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("putFleetsFleetIdWingsWingIdNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutFleetsFleetIDWingsWingIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutFleetsFleetIDWingsWingIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PutFleetsFleetIDWingsWingIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
