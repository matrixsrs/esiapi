package live

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/models"
)

// GetCharactersCharacterIDSearchReader is a Reader for the GetCharactersCharacterIDSearch structure.
type GetCharactersCharacterIDSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCharactersCharacterIDSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDSearchOK creates a GetCharactersCharacterIDSearchOK with default headers values
func NewGetCharactersCharacterIDSearchOK() *GetCharactersCharacterIDSearchOK {
	return &GetCharactersCharacterIDSearchOK{}
}

/*GetCharactersCharacterIDSearchOK handles this case with default header values.

A list of search results
*/
type GetCharactersCharacterIDSearchOK struct {
	Payload *models.GetCharactersCharacterIDSearchOk
}

func (o *GetCharactersCharacterIDSearchOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/search/][%d] getCharactersCharacterIdSearchOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDSearchOk)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDSearchForbidden creates a GetCharactersCharacterIDSearchForbidden with default headers values
func NewGetCharactersCharacterIDSearchForbidden() *GetCharactersCharacterIDSearchForbidden {
	return &GetCharactersCharacterIDSearchForbidden{}
}

/*GetCharactersCharacterIDSearchForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDSearchForbidden struct {
	Payload *models.GetCharactersCharacterIDSearchForbidden
}

func (o *GetCharactersCharacterIDSearchForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/search/][%d] getCharactersCharacterIdSearchForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDSearchForbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDSearchInternalServerError creates a GetCharactersCharacterIDSearchInternalServerError with default headers values
func NewGetCharactersCharacterIDSearchInternalServerError() *GetCharactersCharacterIDSearchInternalServerError {
	return &GetCharactersCharacterIDSearchInternalServerError{}
}

/*GetCharactersCharacterIDSearchInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDSearchInternalServerError struct {
	Payload *models.GetCharactersCharacterIDSearchInternalServerError
}

func (o *GetCharactersCharacterIDSearchInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/search/][%d] getCharactersCharacterIdSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDSearchInternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
