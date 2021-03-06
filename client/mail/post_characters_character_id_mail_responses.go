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

// PostCharactersCharacterIDMailReader is a Reader for the PostCharactersCharacterIDMail structure.
type PostCharactersCharacterIDMailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCharactersCharacterIDMailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostCharactersCharacterIDMailCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostCharactersCharacterIDMailBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostCharactersCharacterIDMailForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostCharactersCharacterIDMailInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostCharactersCharacterIDMailCreated creates a PostCharactersCharacterIDMailCreated with default headers values
func NewPostCharactersCharacterIDMailCreated() *PostCharactersCharacterIDMailCreated {
	return &PostCharactersCharacterIDMailCreated{}
}

/*PostCharactersCharacterIDMailCreated handles this case with default header values.

Mail created
*/
type PostCharactersCharacterIDMailCreated struct {
	Payload int32
}

func (o *PostCharactersCharacterIDMailCreated) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/mail/][%d] postCharactersCharacterIdMailCreated  %+v", 201, o.Payload)
}

func (o *PostCharactersCharacterIDMailCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDMailBadRequest creates a PostCharactersCharacterIDMailBadRequest with default headers values
func NewPostCharactersCharacterIDMailBadRequest() *PostCharactersCharacterIDMailBadRequest {
	return &PostCharactersCharacterIDMailBadRequest{}
}

/*PostCharactersCharacterIDMailBadRequest handles this case with default header values.

Only one corporation, alliance, or mailing list can be the recipient of a mail
*/
type PostCharactersCharacterIDMailBadRequest struct {
	Payload PostCharactersCharacterIDMailBadRequestBody
}

func (o *PostCharactersCharacterIDMailBadRequest) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/mail/][%d] postCharactersCharacterIdMailBadRequest  %+v", 400, o.Payload)
}

func (o *PostCharactersCharacterIDMailBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDMailForbidden creates a PostCharactersCharacterIDMailForbidden with default headers values
func NewPostCharactersCharacterIDMailForbidden() *PostCharactersCharacterIDMailForbidden {
	return &PostCharactersCharacterIDMailForbidden{}
}

/*PostCharactersCharacterIDMailForbidden handles this case with default header values.

Forbidden
*/
type PostCharactersCharacterIDMailForbidden struct {
	Payload *models.Forbidden
}

func (o *PostCharactersCharacterIDMailForbidden) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/mail/][%d] postCharactersCharacterIdMailForbidden  %+v", 403, o.Payload)
}

func (o *PostCharactersCharacterIDMailForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDMailInternalServerError creates a PostCharactersCharacterIDMailInternalServerError with default headers values
func NewPostCharactersCharacterIDMailInternalServerError() *PostCharactersCharacterIDMailInternalServerError {
	return &PostCharactersCharacterIDMailInternalServerError{}
}

/*PostCharactersCharacterIDMailInternalServerError handles this case with default header values.

Internal server error
*/
type PostCharactersCharacterIDMailInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *PostCharactersCharacterIDMailInternalServerError) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/mail/][%d] postCharactersCharacterIdMailInternalServerError  %+v", 500, o.Payload)
}

func (o *PostCharactersCharacterIDMailInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostCharactersCharacterIDMailBadRequestBody post_characters_character_id_mail_bad_request
//
// Bad request
swagger:model PostCharactersCharacterIDMailBadRequestBody
*/

type PostCharactersCharacterIDMailBadRequestBody struct {

	// post_characters_character_id_mail_400_bad_request
	//
	// Bad request message
	// Required: true
	Error *string `json:"error"`
}

/* polymorph PostCharactersCharacterIDMailBadRequestBody error false */

// Validate validates this post characters character ID mail bad request body
func (o *PostCharactersCharacterIDMailBadRequestBody) Validate(formats strfmt.Registry) error {
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

func (o *PostCharactersCharacterIDMailBadRequestBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("postCharactersCharacterIdMailBadRequest"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostCharactersCharacterIDMailBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCharactersCharacterIDMailBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PostCharactersCharacterIDMailBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostCharactersCharacterIDMailBody post_characters_character_id_mail_mail
//
// mail schema
swagger:model PostCharactersCharacterIDMailBody
*/

type PostCharactersCharacterIDMailBody struct {

	// post_characters_character_id_mail_approved_cost
	//
	// approved_cost integer
	// Required: true
	ApprovedCost *int64 `json:"approved_cost"`

	// post_characters_character_id_mail_body
	//
	// body string
	// Required: true
	// Max Length: 10000
	Body *string `json:"body"`

	// post_characters_character_id_mail_recipients
	//
	// recipients array
	// Required: true
	// Max Items: 50
	// Min Items: 1
	Recipients []*RecipientsItems0 `json:"recipients"`

	// post_characters_character_id_mail_subject
	//
	// subject string
	// Required: true
	// Max Length: 1000
	Subject *string `json:"subject"`
}

/* polymorph PostCharactersCharacterIDMailBody approved_cost false */

/* polymorph PostCharactersCharacterIDMailBody body false */

/* polymorph PostCharactersCharacterIDMailBody recipients false */

/* polymorph PostCharactersCharacterIDMailBody subject false */

// MarshalBinary interface implementation
func (o *PostCharactersCharacterIDMailBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCharactersCharacterIDMailBody) UnmarshalBinary(b []byte) error {
	var res PostCharactersCharacterIDMailBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
