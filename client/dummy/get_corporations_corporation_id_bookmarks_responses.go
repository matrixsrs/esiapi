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

// GetCorporationsCorporationIDBookmarksReader is a Reader for the GetCorporationsCorporationIDBookmarks structure.
type GetCorporationsCorporationIDBookmarksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDBookmarksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewGetCorporationsCorporationIDBookmarksNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetCorporationsCorporationIDBookmarksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDBookmarksNoContent creates a GetCorporationsCorporationIDBookmarksNoContent with default headers values
func NewGetCorporationsCorporationIDBookmarksNoContent() *GetCorporationsCorporationIDBookmarksNoContent {
	return &GetCorporationsCorporationIDBookmarksNoContent{}
}

/*GetCorporationsCorporationIDBookmarksNoContent handles this case with default header values.

Nice Dummy
*/
type GetCorporationsCorporationIDBookmarksNoContent struct {
}

func (o *GetCorporationsCorporationIDBookmarksNoContent) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/bookmarks/][%d] getCorporationsCorporationIdBookmarksNoContent ", 204)
}

func (o *GetCorporationsCorporationIDBookmarksNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCorporationsCorporationIDBookmarksInternalServerError creates a GetCorporationsCorporationIDBookmarksInternalServerError with default headers values
func NewGetCorporationsCorporationIDBookmarksInternalServerError() *GetCorporationsCorporationIDBookmarksInternalServerError {
	return &GetCorporationsCorporationIDBookmarksInternalServerError{}
}

/*GetCorporationsCorporationIDBookmarksInternalServerError handles this case with default header values.

Internal server error
*/
type GetCorporationsCorporationIDBookmarksInternalServerError struct {
	Payload *models.GetCorporationsCorporationIDBookmarksInternalServerError
}

func (o *GetCorporationsCorporationIDBookmarksInternalServerError) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/bookmarks/][%d] getCorporationsCorporationIdBookmarksInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationsCorporationIDBookmarksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCorporationsCorporationIDBookmarksInternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
