// Code generated by go-swagger; DO NOT EDIT.

package alliance

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

// GetAlliancesAllianceIDReader is a Reader for the GetAlliancesAllianceID structure.
type GetAlliancesAllianceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlliancesAllianceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAlliancesAllianceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetAlliancesAllianceIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAlliancesAllianceIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAlliancesAllianceIDOK creates a GetAlliancesAllianceIDOK with default headers values
func NewGetAlliancesAllianceIDOK() *GetAlliancesAllianceIDOK {
	return &GetAlliancesAllianceIDOK{}
}

/*GetAlliancesAllianceIDOK handles this case with default header values.

Public data about an alliance
*/
type GetAlliancesAllianceIDOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload GetAlliancesAllianceIDOKBody
}

func (o *GetAlliancesAllianceIDOK) Error() string {
	return fmt.Sprintf("[GET /alliances/{alliance_id}/][%d] getAlliancesAllianceIdOK  %+v", 200, o.Payload)
}

func (o *GetAlliancesAllianceIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAlliancesAllianceIDNotFound creates a GetAlliancesAllianceIDNotFound with default headers values
func NewGetAlliancesAllianceIDNotFound() *GetAlliancesAllianceIDNotFound {
	return &GetAlliancesAllianceIDNotFound{}
}

/*GetAlliancesAllianceIDNotFound handles this case with default header values.

Alliance not found
*/
type GetAlliancesAllianceIDNotFound struct {
	Payload GetAlliancesAllianceIDNotFoundBody
}

func (o *GetAlliancesAllianceIDNotFound) Error() string {
	return fmt.Sprintf("[GET /alliances/{alliance_id}/][%d] getAlliancesAllianceIdNotFound  %+v", 404, o.Payload)
}

func (o *GetAlliancesAllianceIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesAllianceIDInternalServerError creates a GetAlliancesAllianceIDInternalServerError with default headers values
func NewGetAlliancesAllianceIDInternalServerError() *GetAlliancesAllianceIDInternalServerError {
	return &GetAlliancesAllianceIDInternalServerError{}
}

/*GetAlliancesAllianceIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetAlliancesAllianceIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetAlliancesAllianceIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /alliances/{alliance_id}/][%d] getAlliancesAllianceIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlliancesAllianceIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetAlliancesAllianceIDNotFoundBody get_alliances_alliance_id_not_found
//
// Alliance not found
swagger:model GetAlliancesAllianceIDNotFoundBody
*/

type GetAlliancesAllianceIDNotFoundBody struct {

	// get_alliances_alliance_id_error
	//
	// error message
	// Required: true
	Error *string `json:"error"`
}

/* polymorph GetAlliancesAllianceIDNotFoundBody error false */

// Validate validates this get alliances alliance ID not found body
func (o *GetAlliancesAllianceIDNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *GetAlliancesAllianceIDNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getAlliancesAllianceIdNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAlliancesAllianceIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAlliancesAllianceIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetAlliancesAllianceIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAlliancesAllianceIDOKBody get_alliances_alliance_id_ok
//
// 200 ok object
swagger:model GetAlliancesAllianceIDOKBody
*/

type GetAlliancesAllianceIDOKBody struct {

	// get_alliances_alliance_id_alliance_name
	//
	// the full name of the alliance
	// Required: true
	AllianceName *string `json:"alliance_name"`

	// get_alliances_alliance_id_date_founded
	//
	// date_founded string
	// Required: true
	DateFounded *strfmt.DateTime `json:"date_founded"`

	// get_alliances_alliance_id_executor_corp
	//
	// the executor corporation ID, if this alliance is not closed
	// Required: true
	ExecutorCorp *int32 `json:"executor_corp"`

	// get_alliances_alliance_id_ticker
	//
	// the short name of the alliance
	// Required: true
	Ticker *string `json:"ticker"`
}

/* polymorph GetAlliancesAllianceIDOKBody alliance_name false */

/* polymorph GetAlliancesAllianceIDOKBody date_founded false */

/* polymorph GetAlliancesAllianceIDOKBody executor_corp false */

/* polymorph GetAlliancesAllianceIDOKBody ticker false */

// Validate validates this get alliances alliance ID o k body
func (o *GetAlliancesAllianceIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAllianceName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateDateFounded(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateExecutorCorp(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateTicker(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAlliancesAllianceIDOKBody) validateAllianceName(formats strfmt.Registry) error {

	if err := validate.Required("getAlliancesAllianceIdOK"+"."+"alliance_name", "body", o.AllianceName); err != nil {
		return err
	}

	return nil
}

func (o *GetAlliancesAllianceIDOKBody) validateDateFounded(formats strfmt.Registry) error {

	if err := validate.Required("getAlliancesAllianceIdOK"+"."+"date_founded", "body", o.DateFounded); err != nil {
		return err
	}

	if err := validate.FormatOf("getAlliancesAllianceIdOK"+"."+"date_founded", "body", "date-time", o.DateFounded.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetAlliancesAllianceIDOKBody) validateExecutorCorp(formats strfmt.Registry) error {

	if err := validate.Required("getAlliancesAllianceIdOK"+"."+"executor_corp", "body", o.ExecutorCorp); err != nil {
		return err
	}

	return nil
}

func (o *GetAlliancesAllianceIDOKBody) validateTicker(formats strfmt.Registry) error {

	if err := validate.Required("getAlliancesAllianceIdOK"+"."+"ticker", "body", o.Ticker); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAlliancesAllianceIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAlliancesAllianceIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetAlliancesAllianceIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
