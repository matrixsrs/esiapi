// Code generated by go-swagger; DO NOT EDIT.

package killmails

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

// GetCharactersCharacterIDKillmailsRecentReader is a Reader for the GetCharactersCharacterIDKillmailsRecent structure.
type GetCharactersCharacterIDKillmailsRecentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDKillmailsRecentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDKillmailsRecentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCharactersCharacterIDKillmailsRecentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDKillmailsRecentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDKillmailsRecentOK creates a GetCharactersCharacterIDKillmailsRecentOK with default headers values
func NewGetCharactersCharacterIDKillmailsRecentOK() *GetCharactersCharacterIDKillmailsRecentOK {
	return &GetCharactersCharacterIDKillmailsRecentOK{}
}

/*GetCharactersCharacterIDKillmailsRecentOK handles this case with default header values.

A list of killmail IDs and hashes
*/
type GetCharactersCharacterIDKillmailsRecentOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []*GetCharactersCharacterIDKillmailsRecentOKBodyItems0
}

func (o *GetCharactersCharacterIDKillmailsRecentOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/killmails/recent/][%d] getCharactersCharacterIdKillmailsRecentOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDKillmailsRecentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDKillmailsRecentForbidden creates a GetCharactersCharacterIDKillmailsRecentForbidden with default headers values
func NewGetCharactersCharacterIDKillmailsRecentForbidden() *GetCharactersCharacterIDKillmailsRecentForbidden {
	return &GetCharactersCharacterIDKillmailsRecentForbidden{}
}

/*GetCharactersCharacterIDKillmailsRecentForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDKillmailsRecentForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDKillmailsRecentForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/killmails/recent/][%d] getCharactersCharacterIdKillmailsRecentForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDKillmailsRecentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDKillmailsRecentInternalServerError creates a GetCharactersCharacterIDKillmailsRecentInternalServerError with default headers values
func NewGetCharactersCharacterIDKillmailsRecentInternalServerError() *GetCharactersCharacterIDKillmailsRecentInternalServerError {
	return &GetCharactersCharacterIDKillmailsRecentInternalServerError{}
}

/*GetCharactersCharacterIDKillmailsRecentInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDKillmailsRecentInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDKillmailsRecentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/killmails/recent/][%d] getCharactersCharacterIdKillmailsRecentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDKillmailsRecentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDKillmailsRecentOKBodyItems0 get_characters_character_id_killmails_recent_200_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDKillmailsRecentOKBodyItems0
*/

type GetCharactersCharacterIDKillmailsRecentOKBodyItems0 struct {

	// get_characters_character_id_killmails_recent_killmail_hash
	//
	// A hash of this killmail
	// Required: true
	KillmailHash *string `json:"killmail_hash"`

	// get_characters_character_id_killmails_recent_killmail_id
	//
	// ID of this killmail
	// Required: true
	KillmailID *int32 `json:"killmail_id"`
}

/* polymorph GetCharactersCharacterIDKillmailsRecentOKBodyItems0 killmail_hash false */

/* polymorph GetCharactersCharacterIDKillmailsRecentOKBodyItems0 killmail_id false */

// Validate validates this get characters character ID killmails recent o k body items0
func (o *GetCharactersCharacterIDKillmailsRecentOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateKillmailHash(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateKillmailID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCharactersCharacterIDKillmailsRecentOKBodyItems0) validateKillmailHash(formats strfmt.Registry) error {

	if err := validate.Required("killmail_hash", "body", o.KillmailHash); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDKillmailsRecentOKBodyItems0) validateKillmailID(formats strfmt.Registry) error {

	if err := validate.Required("killmail_id", "body", o.KillmailID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDKillmailsRecentOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDKillmailsRecentOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDKillmailsRecentOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
