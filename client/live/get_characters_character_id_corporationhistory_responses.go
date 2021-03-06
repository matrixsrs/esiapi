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

// GetCharactersCharacterIDCorporationhistoryReader is a Reader for the GetCharactersCharacterIDCorporationhistory structure.
type GetCharactersCharacterIDCorporationhistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDCorporationhistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDCorporationhistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetCharactersCharacterIDCorporationhistoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDCorporationhistoryOK creates a GetCharactersCharacterIDCorporationhistoryOK with default headers values
func NewGetCharactersCharacterIDCorporationhistoryOK() *GetCharactersCharacterIDCorporationhistoryOK {
	return &GetCharactersCharacterIDCorporationhistoryOK{}
}

/*GetCharactersCharacterIDCorporationhistoryOK handles this case with default header values.

Corporation history for the given character
*/
type GetCharactersCharacterIDCorporationhistoryOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []*models.GetCharactersCharacterIDCorporationhistory200Ok
}

func (o *GetCharactersCharacterIDCorporationhistoryOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/corporationhistory/][%d] getCharactersCharacterIdCorporationhistoryOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDCorporationhistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDCorporationhistoryInternalServerError creates a GetCharactersCharacterIDCorporationhistoryInternalServerError with default headers values
func NewGetCharactersCharacterIDCorporationhistoryInternalServerError() *GetCharactersCharacterIDCorporationhistoryInternalServerError {
	return &GetCharactersCharacterIDCorporationhistoryInternalServerError{}
}

/*GetCharactersCharacterIDCorporationhistoryInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDCorporationhistoryInternalServerError struct {
	Payload *models.GetCharactersCharacterIDCorporationhistoryInternalServerError
}

func (o *GetCharactersCharacterIDCorporationhistoryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/corporationhistory/][%d] getCharactersCharacterIdCorporationhistoryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDCorporationhistoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDCorporationhistoryInternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
