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

// GetCorporationsCorporationIDAlliancehistoryReader is a Reader for the GetCorporationsCorporationIDAlliancehistory structure.
type GetCorporationsCorporationIDAlliancehistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDAlliancehistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCorporationsCorporationIDAlliancehistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetCorporationsCorporationIDAlliancehistoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDAlliancehistoryOK creates a GetCorporationsCorporationIDAlliancehistoryOK with default headers values
func NewGetCorporationsCorporationIDAlliancehistoryOK() *GetCorporationsCorporationIDAlliancehistoryOK {
	return &GetCorporationsCorporationIDAlliancehistoryOK{}
}

/*GetCorporationsCorporationIDAlliancehistoryOK handles this case with default header values.

Alliance history for the given corporation
*/
type GetCorporationsCorporationIDAlliancehistoryOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []*models.GetCorporationsCorporationIDAlliancehistory200Ok
}

func (o *GetCorporationsCorporationIDAlliancehistoryOK) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/alliancehistory/][%d] getCorporationsCorporationIdAlliancehistoryOK  %+v", 200, o.Payload)
}

func (o *GetCorporationsCorporationIDAlliancehistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDAlliancehistoryInternalServerError creates a GetCorporationsCorporationIDAlliancehistoryInternalServerError with default headers values
func NewGetCorporationsCorporationIDAlliancehistoryInternalServerError() *GetCorporationsCorporationIDAlliancehistoryInternalServerError {
	return &GetCorporationsCorporationIDAlliancehistoryInternalServerError{}
}

/*GetCorporationsCorporationIDAlliancehistoryInternalServerError handles this case with default header values.

Internal server error
*/
type GetCorporationsCorporationIDAlliancehistoryInternalServerError struct {
	Payload *models.GetCorporationsCorporationIDAlliancehistoryInternalServerError
}

func (o *GetCorporationsCorporationIDAlliancehistoryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/alliancehistory/][%d] getCorporationsCorporationIdAlliancehistoryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationsCorporationIDAlliancehistoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCorporationsCorporationIDAlliancehistoryInternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
