package alliance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

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

	Payload *models.GetAlliancesAllianceIDOk
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

	o.Payload = new(models.GetAlliancesAllianceIDOk)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
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
	Payload *models.GetAlliancesAllianceIDNotFound
}

func (o *GetAlliancesAllianceIDNotFound) Error() string {
	return fmt.Sprintf("[GET /alliances/{alliance_id}/][%d] getAlliancesAllianceIdNotFound  %+v", 404, o.Payload)
}

func (o *GetAlliancesAllianceIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetAlliancesAllianceIDNotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
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
	Payload *models.GetAlliancesAllianceIDInternalServerError
}

func (o *GetAlliancesAllianceIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /alliances/{alliance_id}/][%d] getAlliancesAllianceIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlliancesAllianceIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetAlliancesAllianceIDInternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}