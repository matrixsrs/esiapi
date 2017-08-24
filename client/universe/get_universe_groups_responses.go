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

// GetUniverseGroupsReader is a Reader for the GetUniverseGroups structure.
type GetUniverseGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUniverseGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetUniverseGroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUniverseGroupsOK creates a GetUniverseGroupsOK with default headers values
func NewGetUniverseGroupsOK() *GetUniverseGroupsOK {
	return &GetUniverseGroupsOK{}
}

/*GetUniverseGroupsOK handles this case with default header values.

A list of item group ids
*/
type GetUniverseGroupsOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []int32
}

func (o *GetUniverseGroupsOK) Error() string {
	return fmt.Sprintf("[GET /universe/groups/][%d] getUniverseGroupsOK  %+v", 200, o.Payload)
}

func (o *GetUniverseGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseGroupsInternalServerError creates a GetUniverseGroupsInternalServerError with default headers values
func NewGetUniverseGroupsInternalServerError() *GetUniverseGroupsInternalServerError {
	return &GetUniverseGroupsInternalServerError{}
}

/*GetUniverseGroupsInternalServerError handles this case with default header values.

Internal server error
*/
type GetUniverseGroupsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetUniverseGroupsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/groups/][%d] getUniverseGroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseGroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
