// Code generated by go-swagger; DO NOT EDIT.

package sovereignty

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

// GetSovereigntyMapReader is a Reader for the GetSovereigntyMap structure.
type GetSovereigntyMapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSovereigntyMapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSovereigntyMapOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetSovereigntyMapInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSovereigntyMapOK creates a GetSovereigntyMapOK with default headers values
func NewGetSovereigntyMapOK() *GetSovereigntyMapOK {
	return &GetSovereigntyMapOK{}
}

/*GetSovereigntyMapOK handles this case with default header values.

A list of sovereignty information for solar systems in New Eden
*/
type GetSovereigntyMapOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []*GetSovereigntyMapOKBodyItems0
}

func (o *GetSovereigntyMapOK) Error() string {
	return fmt.Sprintf("[GET /sovereignty/map/][%d] getSovereigntyMapOK  %+v", 200, o.Payload)
}

func (o *GetSovereigntyMapOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSovereigntyMapInternalServerError creates a GetSovereigntyMapInternalServerError with default headers values
func NewGetSovereigntyMapInternalServerError() *GetSovereigntyMapInternalServerError {
	return &GetSovereigntyMapInternalServerError{}
}

/*GetSovereigntyMapInternalServerError handles this case with default header values.

Internal server error
*/
type GetSovereigntyMapInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetSovereigntyMapInternalServerError) Error() string {
	return fmt.Sprintf("[GET /sovereignty/map/][%d] getSovereigntyMapInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSovereigntyMapInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetSovereigntyMapOKBodyItems0 get_sovereignty_map_200_ok
//
// 200 ok object
swagger:model GetSovereigntyMapOKBodyItems0
*/

type GetSovereigntyMapOKBodyItems0 struct {

	// get_sovereignty_map_alliance_id
	//
	// alliance_id integer
	AllianceID int32 `json:"alliance_id,omitempty"`

	// get_sovereignty_map_corporation_id
	//
	// corporation_id integer
	CorporationID int32 `json:"corporation_id,omitempty"`

	// get_sovereignty_map_faction_id
	//
	// faction_id integer
	FactionID int32 `json:"faction_id,omitempty"`

	// get_sovereignty_map_system_id
	//
	// system_id integer
	// Required: true
	SystemID *int32 `json:"system_id"`
}

/* polymorph GetSovereigntyMapOKBodyItems0 alliance_id false */

/* polymorph GetSovereigntyMapOKBodyItems0 corporation_id false */

/* polymorph GetSovereigntyMapOKBodyItems0 faction_id false */

/* polymorph GetSovereigntyMapOKBodyItems0 system_id false */

// Validate validates this get sovereignty map o k body items0
func (o *GetSovereigntyMapOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSystemID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSovereigntyMapOKBodyItems0) validateSystemID(formats strfmt.Registry) error {

	if err := validate.Required("system_id", "body", o.SystemID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSovereigntyMapOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSovereigntyMapOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetSovereigntyMapOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
