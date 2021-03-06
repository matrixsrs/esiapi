// Code generated by go-swagger; DO NOT EDIT.

package wars

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/models"
)

// GetWarsWarIDReader is a Reader for the GetWarsWarID structure.
type GetWarsWarIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWarsWarIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWarsWarIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 422:
		result := NewGetWarsWarIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetWarsWarIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetWarsWarIDOK creates a GetWarsWarIDOK with default headers values
func NewGetWarsWarIDOK() *GetWarsWarIDOK {
	return &GetWarsWarIDOK{}
}

/*GetWarsWarIDOK handles this case with default header values.

Details about a war
*/
type GetWarsWarIDOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload GetWarsWarIDOKBody
}

func (o *GetWarsWarIDOK) Error() string {
	return fmt.Sprintf("[GET /wars/{war_id}/][%d] getWarsWarIdOK  %+v", 200, o.Payload)
}

func (o *GetWarsWarIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetWarsWarIDUnprocessableEntity creates a GetWarsWarIDUnprocessableEntity with default headers values
func NewGetWarsWarIDUnprocessableEntity() *GetWarsWarIDUnprocessableEntity {
	return &GetWarsWarIDUnprocessableEntity{}
}

/*GetWarsWarIDUnprocessableEntity handles this case with default header values.

War not found
*/
type GetWarsWarIDUnprocessableEntity struct {
	Payload GetWarsWarIDUnprocessableEntityBody
}

func (o *GetWarsWarIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /wars/{war_id}/][%d] getWarsWarIdUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetWarsWarIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWarsWarIDInternalServerError creates a GetWarsWarIDInternalServerError with default headers values
func NewGetWarsWarIDInternalServerError() *GetWarsWarIDInternalServerError {
	return &GetWarsWarIDInternalServerError{}
}

/*GetWarsWarIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetWarsWarIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetWarsWarIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /wars/{war_id}/][%d] getWarsWarIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWarsWarIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AlliesItems0 get_wars_war_id_ally
//
// ally object
swagger:model AlliesItems0
*/

type AlliesItems0 struct {

	// get_wars_war_id_alliance_id
	//
	// Alliance ID if and only if this ally is an alliance
	AllianceID int32 `json:"alliance_id,omitempty"`

	// get_wars_war_id_corporation_id
	//
	// Corporation ID if and only if this ally is a corporation
	CorporationID int32 `json:"corporation_id,omitempty"`
}

/* polymorph AlliesItems0 alliance_id false */

/* polymorph AlliesItems0 corporation_id false */

// Validate validates this allies items0
func (o *AlliesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *AlliesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AlliesItems0) UnmarshalBinary(b []byte) error {
	var res AlliesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWarsWarIDOKBody get_wars_war_id_ok
//
// 200 ok object
swagger:model GetWarsWarIDOKBody
*/

type GetWarsWarIDOKBody struct {

	// aggressor
	// Required: true
	Aggressor *GetWarsWarIDOKBodyAggressor `json:"aggressor"`

	// get_wars_war_id_allies
	//
	// allied corporations or alliances, each object contains either corporation_id or alliance_id
	// Required: true
	// Max Items: 10000
	Allies []*AlliesItems0 `json:"allies"`

	// get_wars_war_id_declared
	//
	// Time that the war was declared
	// Required: true
	Declared *strfmt.DateTime `json:"declared"`

	// defender
	// Required: true
	Defender *GetWarsWarIDOKBodyDefender `json:"defender"`

	// get_wars_war_id_finished
	//
	// Time the war ended and shooting was no longer allowed
	// Required: true
	Finished *strfmt.DateTime `json:"finished"`

	// get_wars_war_id_id
	//
	// ID of the specified war
	// Required: true
	ID *int32 `json:"id"`

	// get_wars_war_id_mutual
	//
	// Was the war declared mutual by both parties
	// Required: true
	Mutual *bool `json:"mutual"`

	// get_wars_war_id_open_for_allies
	//
	// Is the war currently open for allies or not
	// Required: true
	OpenForAllies *bool `json:"open_for_allies"`

	// get_wars_war_id_retracted
	//
	// Time the war was retracted but both sides could still shoot each other
	// Required: true
	Retracted *strfmt.DateTime `json:"retracted"`

	// get_wars_war_id_started
	//
	// Time when the war started and both sides could shoot each other
	// Required: true
	Started *strfmt.DateTime `json:"started"`
}

/* polymorph GetWarsWarIDOKBody aggressor false */

/* polymorph GetWarsWarIDOKBody allies false */

/* polymorph GetWarsWarIDOKBody declared false */

/* polymorph GetWarsWarIDOKBody defender false */

/* polymorph GetWarsWarIDOKBody finished false */

/* polymorph GetWarsWarIDOKBody id false */

/* polymorph GetWarsWarIDOKBody mutual false */

/* polymorph GetWarsWarIDOKBody open_for_allies false */

/* polymorph GetWarsWarIDOKBody retracted false */

/* polymorph GetWarsWarIDOKBody started false */

// Validate validates this get wars war ID o k body
func (o *GetWarsWarIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAggressor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateAllies(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateDeclared(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateDefender(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateFinished(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateMutual(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateOpenForAllies(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateRetracted(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateStarted(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWarsWarIDOKBody) validateAggressor(formats strfmt.Registry) error {

	if err := validate.Required("getWarsWarIdOK"+"."+"aggressor", "body", o.Aggressor); err != nil {
		return err
	}

	if o.Aggressor != nil {

		if err := o.Aggressor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWarsWarIdOK" + "." + "aggressor")
			}
			return err
		}
	}

	return nil
}

func (o *GetWarsWarIDOKBody) validateAllies(formats strfmt.Registry) error {

	if err := validate.Required("getWarsWarIdOK"+"."+"allies", "body", o.Allies); err != nil {
		return err
	}

	iAlliesSize := int64(len(o.Allies))

	if err := validate.MaxItems("getWarsWarIdOK"+"."+"allies", "body", iAlliesSize, 10000); err != nil {
		return err
	}

	for i := 0; i < len(o.Allies); i++ {

		if swag.IsZero(o.Allies[i]) { // not required
			continue
		}

		if o.Allies[i] != nil {

			if err := o.Allies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getWarsWarIdOK" + "." + "allies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetWarsWarIDOKBody) validateDeclared(formats strfmt.Registry) error {

	if err := validate.Required("getWarsWarIdOK"+"."+"declared", "body", o.Declared); err != nil {
		return err
	}

	if err := validate.FormatOf("getWarsWarIdOK"+"."+"declared", "body", "date-time", o.Declared.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetWarsWarIDOKBody) validateDefender(formats strfmt.Registry) error {

	if err := validate.Required("getWarsWarIdOK"+"."+"defender", "body", o.Defender); err != nil {
		return err
	}

	if o.Defender != nil {

		if err := o.Defender.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWarsWarIdOK" + "." + "defender")
			}
			return err
		}
	}

	return nil
}

func (o *GetWarsWarIDOKBody) validateFinished(formats strfmt.Registry) error {

	if err := validate.Required("getWarsWarIdOK"+"."+"finished", "body", o.Finished); err != nil {
		return err
	}

	if err := validate.FormatOf("getWarsWarIdOK"+"."+"finished", "body", "date-time", o.Finished.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetWarsWarIDOKBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("getWarsWarIdOK"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *GetWarsWarIDOKBody) validateMutual(formats strfmt.Registry) error {

	if err := validate.Required("getWarsWarIdOK"+"."+"mutual", "body", o.Mutual); err != nil {
		return err
	}

	return nil
}

func (o *GetWarsWarIDOKBody) validateOpenForAllies(formats strfmt.Registry) error {

	if err := validate.Required("getWarsWarIdOK"+"."+"open_for_allies", "body", o.OpenForAllies); err != nil {
		return err
	}

	return nil
}

func (o *GetWarsWarIDOKBody) validateRetracted(formats strfmt.Registry) error {

	if err := validate.Required("getWarsWarIdOK"+"."+"retracted", "body", o.Retracted); err != nil {
		return err
	}

	if err := validate.FormatOf("getWarsWarIdOK"+"."+"retracted", "body", "date-time", o.Retracted.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetWarsWarIDOKBody) validateStarted(formats strfmt.Registry) error {

	if err := validate.Required("getWarsWarIdOK"+"."+"started", "body", o.Started); err != nil {
		return err
	}

	if err := validate.FormatOf("getWarsWarIdOK"+"."+"started", "body", "date-time", o.Started.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWarsWarIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWarsWarIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetWarsWarIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWarsWarIDOKBodyAggressor get_wars_war_id_aggressor
//
// The aggressor corporation or alliance that declared this war, only contains either corporation_id or alliance_id
swagger:model GetWarsWarIDOKBodyAggressor
*/

type GetWarsWarIDOKBodyAggressor struct {

	// get_wars_war_id_alliance_id
	//
	// Alliance ID if and only if the aggressor is an alliance
	AllianceID int32 `json:"alliance_id,omitempty"`

	// get_wars_war_id_corporation_id
	//
	// Corporation ID if and only if the aggressor is a corporation
	CorporationID int32 `json:"corporation_id,omitempty"`

	// get_wars_war_id_isk_destroyed
	//
	// ISK value of ships the aggressor has destroyed
	// Required: true
	IskDestroyed *float32 `json:"isk_destroyed"`

	// get_wars_war_id_ships_killed
	//
	// The number of ships the aggressor has killed
	// Required: true
	ShipsKilled *int32 `json:"ships_killed"`
}

/* polymorph GetWarsWarIDOKBodyAggressor alliance_id false */

/* polymorph GetWarsWarIDOKBodyAggressor corporation_id false */

/* polymorph GetWarsWarIDOKBodyAggressor isk_destroyed false */

/* polymorph GetWarsWarIDOKBodyAggressor ships_killed false */

// Validate validates this get wars war ID o k body aggressor
func (o *GetWarsWarIDOKBodyAggressor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIskDestroyed(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateShipsKilled(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWarsWarIDOKBodyAggressor) validateIskDestroyed(formats strfmt.Registry) error {

	if err := validate.Required("getWarsWarIdOK"+"."+"aggressor"+"."+"isk_destroyed", "body", o.IskDestroyed); err != nil {
		return err
	}

	return nil
}

func (o *GetWarsWarIDOKBodyAggressor) validateShipsKilled(formats strfmt.Registry) error {

	if err := validate.Required("getWarsWarIdOK"+"."+"aggressor"+"."+"ships_killed", "body", o.ShipsKilled); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWarsWarIDOKBodyAggressor) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWarsWarIDOKBodyAggressor) UnmarshalBinary(b []byte) error {
	var res GetWarsWarIDOKBodyAggressor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWarsWarIDOKBodyDefender get_wars_war_id_defender
//
// The defending corporation or alliance that declared this war, only contains either corporation_id or alliance_id
swagger:model GetWarsWarIDOKBodyDefender
*/

type GetWarsWarIDOKBodyDefender struct {

	// get_wars_war_id_alliance_id
	//
	// Alliance ID if and only if the defender is an alliance
	AllianceID int32 `json:"alliance_id,omitempty"`

	// get_wars_war_id_corporation_id
	//
	// Corporation ID if and only if the defender is a corporation
	CorporationID int32 `json:"corporation_id,omitempty"`

	// get_wars_war_id_isk_destroyed
	//
	// ISK value of ships the defender has killed
	// Required: true
	IskDestroyed *float32 `json:"isk_destroyed"`

	// get_wars_war_id_ships_killed
	//
	// The number of ships the defender has killed
	// Required: true
	ShipsKilled *int32 `json:"ships_killed"`
}

/* polymorph GetWarsWarIDOKBodyDefender alliance_id false */

/* polymorph GetWarsWarIDOKBodyDefender corporation_id false */

/* polymorph GetWarsWarIDOKBodyDefender isk_destroyed false */

/* polymorph GetWarsWarIDOKBodyDefender ships_killed false */

// Validate validates this get wars war ID o k body defender
func (o *GetWarsWarIDOKBodyDefender) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIskDestroyed(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateShipsKilled(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWarsWarIDOKBodyDefender) validateIskDestroyed(formats strfmt.Registry) error {

	if err := validate.Required("getWarsWarIdOK"+"."+"defender"+"."+"isk_destroyed", "body", o.IskDestroyed); err != nil {
		return err
	}

	return nil
}

func (o *GetWarsWarIDOKBodyDefender) validateShipsKilled(formats strfmt.Registry) error {

	if err := validate.Required("getWarsWarIdOK"+"."+"defender"+"."+"ships_killed", "body", o.ShipsKilled); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWarsWarIDOKBodyDefender) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWarsWarIDOKBodyDefender) UnmarshalBinary(b []byte) error {
	var res GetWarsWarIDOKBodyDefender
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWarsWarIDUnprocessableEntityBody get_wars_war_id_unprocessable_entity
//
// Unprocessable entity
swagger:model GetWarsWarIDUnprocessableEntityBody
*/

type GetWarsWarIDUnprocessableEntityBody struct {

	// get_wars_war_id_422_unprocessable_entity
	//
	// Unprocessable entity message
	// Required: true
	Error *string `json:"error"`
}

/* polymorph GetWarsWarIDUnprocessableEntityBody error false */

// Validate validates this get wars war ID unprocessable entity body
func (o *GetWarsWarIDUnprocessableEntityBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWarsWarIDUnprocessableEntityBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getWarsWarIdUnprocessableEntity"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWarsWarIDUnprocessableEntityBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWarsWarIDUnprocessableEntityBody) UnmarshalBinary(b []byte) error {
	var res GetWarsWarIDUnprocessableEntityBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
