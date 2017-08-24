// Code generated by go-swagger; DO NOT EDIT.

package universe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/models"
)

// GetUniverseStructuresReader is a Reader for the GetUniverseStructures structure.
type GetUniverseStructuresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseStructuresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUniverseStructuresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetUniverseStructuresInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUniverseStructuresOK creates a GetUniverseStructuresOK with default headers values
func NewGetUniverseStructuresOK() *GetUniverseStructuresOK {
	return &GetUniverseStructuresOK{}
}

/*GetUniverseStructuresOK handles this case with default header values.

List of public structure IDs
*/
type GetUniverseStructuresOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []*int64
}

func (o *GetUniverseStructuresOK) Error() string {
	return fmt.Sprintf("[GET /universe/structures/][%d] getUniverseStructuresOK  %+v", 200, o.Payload)
}

func (o *GetUniverseStructuresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseStructuresInternalServerError creates a GetUniverseStructuresInternalServerError with default headers values
func NewGetUniverseStructuresInternalServerError() *GetUniverseStructuresInternalServerError {
	return &GetUniverseStructuresInternalServerError{}
}

/*GetUniverseStructuresInternalServerError handles this case with default header values.

Internal server error
*/
type GetUniverseStructuresInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetUniverseStructuresInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/structures/][%d] getUniverseStructuresInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseStructuresInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
