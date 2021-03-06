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

// GetUniverseStationsStationIDReader is a Reader for the GetUniverseStationsStationID structure.
type GetUniverseStationsStationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseStationsStationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUniverseStationsStationIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetUniverseStationsStationIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUniverseStationsStationIDOK creates a GetUniverseStationsStationIDOK with default headers values
func NewGetUniverseStationsStationIDOK() *GetUniverseStationsStationIDOK {
	return &GetUniverseStationsStationIDOK{}
}

/*GetUniverseStationsStationIDOK handles this case with default header values.

Public data about a station
*/
type GetUniverseStationsStationIDOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload *models.GetUniverseStationsStationIDOk
}

func (o *GetUniverseStationsStationIDOK) Error() string {
	return fmt.Sprintf("[GET /universe/stations/{station_id}/][%d] getUniverseStationsStationIdOK  %+v", 200, o.Payload)
}

func (o *GetUniverseStationsStationIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	o.Payload = new(models.GetUniverseStationsStationIDOk)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStationsStationIDInternalServerError creates a GetUniverseStationsStationIDInternalServerError with default headers values
func NewGetUniverseStationsStationIDInternalServerError() *GetUniverseStationsStationIDInternalServerError {
	return &GetUniverseStationsStationIDInternalServerError{}
}

/*GetUniverseStationsStationIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetUniverseStationsStationIDInternalServerError struct {
	Payload *models.GetUniverseStationsStationIDInternalServerError
}

func (o *GetUniverseStationsStationIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/stations/{station_id}/][%d] getUniverseStationsStationIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseStationsStationIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetUniverseStationsStationIDInternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
