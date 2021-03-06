package dummy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/models"
)

// GetCharactersCharacterIDWalletsJournalReader is a Reader for the GetCharactersCharacterIDWalletsJournal structure.
type GetCharactersCharacterIDWalletsJournalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDWalletsJournalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDWalletsJournalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCharactersCharacterIDWalletsJournalForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDWalletsJournalInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDWalletsJournalOK creates a GetCharactersCharacterIDWalletsJournalOK with default headers values
func NewGetCharactersCharacterIDWalletsJournalOK() *GetCharactersCharacterIDWalletsJournalOK {
	return &GetCharactersCharacterIDWalletsJournalOK{}
}

/*GetCharactersCharacterIDWalletsJournalOK handles this case with default header values.

Journal entries
*/
type GetCharactersCharacterIDWalletsJournalOK struct {
	Payload []*models.GetCharactersCharacterIDWalletsJournal200Ok
}

func (o *GetCharactersCharacterIDWalletsJournalOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/wallets/journal/][%d] getCharactersCharacterIdWalletsJournalOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDWalletsJournalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDWalletsJournalForbidden creates a GetCharactersCharacterIDWalletsJournalForbidden with default headers values
func NewGetCharactersCharacterIDWalletsJournalForbidden() *GetCharactersCharacterIDWalletsJournalForbidden {
	return &GetCharactersCharacterIDWalletsJournalForbidden{}
}

/*GetCharactersCharacterIDWalletsJournalForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDWalletsJournalForbidden struct {
	Payload *models.GetCharactersCharacterIDWalletsJournalForbidden
}

func (o *GetCharactersCharacterIDWalletsJournalForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/wallets/journal/][%d] getCharactersCharacterIdWalletsJournalForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDWalletsJournalForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDWalletsJournalForbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDWalletsJournalInternalServerError creates a GetCharactersCharacterIDWalletsJournalInternalServerError with default headers values
func NewGetCharactersCharacterIDWalletsJournalInternalServerError() *GetCharactersCharacterIDWalletsJournalInternalServerError {
	return &GetCharactersCharacterIDWalletsJournalInternalServerError{}
}

/*GetCharactersCharacterIDWalletsJournalInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDWalletsJournalInternalServerError struct {
	Payload *models.GetCharactersCharacterIDWalletsJournalInternalServerError
}

func (o *GetCharactersCharacterIDWalletsJournalInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/wallets/journal/][%d] getCharactersCharacterIdWalletsJournalInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDWalletsJournalInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDWalletsJournalInternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
