// Code generated by go-swagger; DO NOT EDIT.

package mail

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

// PutCharactersCharacterIDMailMailIDReader is a Reader for the PutCharactersCharacterIDMailMailID structure.
type PutCharactersCharacterIDMailMailIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCharactersCharacterIDMailMailIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPutCharactersCharacterIDMailMailIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutCharactersCharacterIDMailMailIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPutCharactersCharacterIDMailMailIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPutCharactersCharacterIDMailMailIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutCharactersCharacterIDMailMailIDNoContent creates a PutCharactersCharacterIDMailMailIDNoContent with default headers values
func NewPutCharactersCharacterIDMailMailIDNoContent() *PutCharactersCharacterIDMailMailIDNoContent {
	return &PutCharactersCharacterIDMailMailIDNoContent{}
}

/*PutCharactersCharacterIDMailMailIDNoContent handles this case with default header values.

Mail updated
*/
type PutCharactersCharacterIDMailMailIDNoContent struct {
}

func (o *PutCharactersCharacterIDMailMailIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /characters/{character_id}/mail/{mail_id}/][%d] putCharactersCharacterIdMailMailIdNoContent ", 204)
}

func (o *PutCharactersCharacterIDMailMailIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCharactersCharacterIDMailMailIDBadRequest creates a PutCharactersCharacterIDMailMailIDBadRequest with default headers values
func NewPutCharactersCharacterIDMailMailIDBadRequest() *PutCharactersCharacterIDMailMailIDBadRequest {
	return &PutCharactersCharacterIDMailMailIDBadRequest{}
}

/*PutCharactersCharacterIDMailMailIDBadRequest handles this case with default header values.

Invalid label ID; or No parameters in body -- nothing to do
*/
type PutCharactersCharacterIDMailMailIDBadRequest struct {
	Payload PutCharactersCharacterIDMailMailIDBadRequestBody
}

func (o *PutCharactersCharacterIDMailMailIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /characters/{character_id}/mail/{mail_id}/][%d] putCharactersCharacterIdMailMailIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutCharactersCharacterIDMailMailIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCharactersCharacterIDMailMailIDForbidden creates a PutCharactersCharacterIDMailMailIDForbidden with default headers values
func NewPutCharactersCharacterIDMailMailIDForbidden() *PutCharactersCharacterIDMailMailIDForbidden {
	return &PutCharactersCharacterIDMailMailIDForbidden{}
}

/*PutCharactersCharacterIDMailMailIDForbidden handles this case with default header values.

Forbidden
*/
type PutCharactersCharacterIDMailMailIDForbidden struct {
	Payload *models.Forbidden
}

func (o *PutCharactersCharacterIDMailMailIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /characters/{character_id}/mail/{mail_id}/][%d] putCharactersCharacterIdMailMailIdForbidden  %+v", 403, o.Payload)
}

func (o *PutCharactersCharacterIDMailMailIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCharactersCharacterIDMailMailIDInternalServerError creates a PutCharactersCharacterIDMailMailIDInternalServerError with default headers values
func NewPutCharactersCharacterIDMailMailIDInternalServerError() *PutCharactersCharacterIDMailMailIDInternalServerError {
	return &PutCharactersCharacterIDMailMailIDInternalServerError{}
}

/*PutCharactersCharacterIDMailMailIDInternalServerError handles this case with default header values.

Internal server error
*/
type PutCharactersCharacterIDMailMailIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *PutCharactersCharacterIDMailMailIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /characters/{character_id}/mail/{mail_id}/][%d] putCharactersCharacterIdMailMailIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PutCharactersCharacterIDMailMailIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PutCharactersCharacterIDMailMailIDBadRequestBody put_characters_character_id_mail_mail_id_bad_request
//
// Bad request
swagger:model PutCharactersCharacterIDMailMailIDBadRequestBody
*/

type PutCharactersCharacterIDMailMailIDBadRequestBody struct {

	// put_characters_character_id_mail_mail_id_400_bad_request
	//
	// Bad request message
	// Required: true
	Error *string `json:"error"`
}

/* polymorph PutCharactersCharacterIDMailMailIDBadRequestBody error false */

// Validate validates this put characters character ID mail mail ID bad request body
func (o *PutCharactersCharacterIDMailMailIDBadRequestBody) Validate(formats strfmt.Registry) error {
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

func (o *PutCharactersCharacterIDMailMailIDBadRequestBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("putCharactersCharacterIdMailMailIdBadRequest"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutCharactersCharacterIDMailMailIDBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutCharactersCharacterIDMailMailIDBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PutCharactersCharacterIDMailMailIDBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PutCharactersCharacterIDMailMailIDBody put_characters_character_id_mail_mail_id_contents
//
// contents object
swagger:model PutCharactersCharacterIDMailMailIDBody
*/

type PutCharactersCharacterIDMailMailIDBody struct {

	// put_characters_character_id_mail_mail_id_labels
	//
	// Labels to assign to the mail. Pre-existing labels are unassigned.
	// Required: true
	// Max Items: 25
	Labels []*int64 `json:"labels"`

	// put_characters_character_id_mail_mail_id_read
	//
	// Whether the mail is flagged as read
	// Required: true
	Read *bool `json:"read"`
}

/* polymorph PutCharactersCharacterIDMailMailIDBody labels false */

/* polymorph PutCharactersCharacterIDMailMailIDBody read false */

// MarshalBinary interface implementation
func (o *PutCharactersCharacterIDMailMailIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutCharactersCharacterIDMailMailIDBody) UnmarshalBinary(b []byte) error {
	var res PutCharactersCharacterIDMailMailIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
