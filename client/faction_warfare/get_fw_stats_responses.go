// Code generated by go-swagger; DO NOT EDIT.

package faction_warfare

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

// GetFwStatsReader is a Reader for the GetFwStats structure.
type GetFwStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFwStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFwStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetFwStatsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFwStatsOK creates a GetFwStatsOK with default headers values
func NewGetFwStatsOK() *GetFwStatsOK {
	return &GetFwStatsOK{}
}

/*GetFwStatsOK handles this case with default header values.

Per faction breakdown of faction warfare statistics
*/
type GetFwStatsOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []*GetFwStatsOKBodyItems0
}

func (o *GetFwStatsOK) Error() string {
	return fmt.Sprintf("[GET /fw/stats/][%d] getFwStatsOK  %+v", 200, o.Payload)
}

func (o *GetFwStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFwStatsInternalServerError creates a GetFwStatsInternalServerError with default headers values
func NewGetFwStatsInternalServerError() *GetFwStatsInternalServerError {
	return &GetFwStatsInternalServerError{}
}

/*GetFwStatsInternalServerError handles this case with default header values.

Internal server error
*/
type GetFwStatsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetFwStatsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /fw/stats/][%d] getFwStatsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFwStatsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetFwStatsOKBodyItems0 get_fw_stats_200_ok
//
// 200 ok object
swagger:model GetFwStatsOKBodyItems0
*/

type GetFwStatsOKBodyItems0 struct {

	// get_fw_stats_faction_id
	//
	// faction_id integer
	// Required: true
	FactionID *int32 `json:"faction_id"`

	// kills
	// Required: true
	Kills *GetFwStatsOKBodyItems0Kills `json:"kills"`

	// get_fw_stats_pilots
	//
	// How many pilots fight for the given faction
	// Required: true
	Pilots *int32 `json:"pilots"`

	// get_fw_stats_systems_controlled
	//
	// The number of solar systems controlled by the given faction
	// Required: true
	SystemsControlled *int32 `json:"systems_controlled"`

	// victory points
	// Required: true
	VictoryPoints *GetFwStatsOKBodyItems0VictoryPoints `json:"victory_points"`
}

/* polymorph GetFwStatsOKBodyItems0 faction_id false */

/* polymorph GetFwStatsOKBodyItems0 kills false */

/* polymorph GetFwStatsOKBodyItems0 pilots false */

/* polymorph GetFwStatsOKBodyItems0 systems_controlled false */

/* polymorph GetFwStatsOKBodyItems0 victory_points false */

// Validate validates this get fw stats o k body items0
func (o *GetFwStatsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFactionID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateKills(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validatePilots(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateSystemsControlled(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateVictoryPoints(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetFwStatsOKBodyItems0) validateFactionID(formats strfmt.Registry) error {

	if err := validate.Required("faction_id", "body", o.FactionID); err != nil {
		return err
	}

	return nil
}

func (o *GetFwStatsOKBodyItems0) validateKills(formats strfmt.Registry) error {

	if err := validate.Required("kills", "body", o.Kills); err != nil {
		return err
	}

	if o.Kills != nil {

		if err := o.Kills.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kills")
			}
			return err
		}
	}

	return nil
}

func (o *GetFwStatsOKBodyItems0) validatePilots(formats strfmt.Registry) error {

	if err := validate.Required("pilots", "body", o.Pilots); err != nil {
		return err
	}

	return nil
}

func (o *GetFwStatsOKBodyItems0) validateSystemsControlled(formats strfmt.Registry) error {

	if err := validate.Required("systems_controlled", "body", o.SystemsControlled); err != nil {
		return err
	}

	return nil
}

func (o *GetFwStatsOKBodyItems0) validateVictoryPoints(formats strfmt.Registry) error {

	if err := validate.Required("victory_points", "body", o.VictoryPoints); err != nil {
		return err
	}

	if o.VictoryPoints != nil {

		if err := o.VictoryPoints.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("victory_points")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetFwStatsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFwStatsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetFwStatsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetFwStatsOKBodyItems0Kills get_fw_stats_kills
//
// Summary of kills against an enemy faction for the given faction
swagger:model GetFwStatsOKBodyItems0Kills
*/

type GetFwStatsOKBodyItems0Kills struct {

	// get_fw_stats_last_week
	//
	// Last week's total number of kills against enemy factions
	// Required: true
	LastWeek *int32 `json:"last_week"`

	// get_fw_stats_total
	//
	// Total number of kills against enemy factions since faction warfare began
	// Required: true
	Total *int32 `json:"total"`

	// get_fw_stats_yesterday
	//
	// Yesterday's total number of kills against enemy factions
	// Required: true
	Yesterday *int32 `json:"yesterday"`
}

/* polymorph GetFwStatsOKBodyItems0Kills last_week false */

/* polymorph GetFwStatsOKBodyItems0Kills total false */

/* polymorph GetFwStatsOKBodyItems0Kills yesterday false */

// Validate validates this get fw stats o k body items0 kills
func (o *GetFwStatsOKBodyItems0Kills) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLastWeek(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateTotal(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateYesterday(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetFwStatsOKBodyItems0Kills) validateLastWeek(formats strfmt.Registry) error {

	if err := validate.Required("kills"+"."+"last_week", "body", o.LastWeek); err != nil {
		return err
	}

	return nil
}

func (o *GetFwStatsOKBodyItems0Kills) validateTotal(formats strfmt.Registry) error {

	if err := validate.Required("kills"+"."+"total", "body", o.Total); err != nil {
		return err
	}

	return nil
}

func (o *GetFwStatsOKBodyItems0Kills) validateYesterday(formats strfmt.Registry) error {

	if err := validate.Required("kills"+"."+"yesterday", "body", o.Yesterday); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetFwStatsOKBodyItems0Kills) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFwStatsOKBodyItems0Kills) UnmarshalBinary(b []byte) error {
	var res GetFwStatsOKBodyItems0Kills
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetFwStatsOKBodyItems0VictoryPoints get_fw_stats_victory_points
//
// Summary of victory points gained for the given faction
swagger:model GetFwStatsOKBodyItems0VictoryPoints
*/

type GetFwStatsOKBodyItems0VictoryPoints struct {

	// get_fw_stats_last_week
	//
	// Last week's victory points gained
	// Required: true
	LastWeek *int32 `json:"last_week"`

	// get_fw_stats_total
	//
	// Total victory points gained since faction warfare began
	// Required: true
	Total *int32 `json:"total"`

	// get_fw_stats_yesterday
	//
	// Yesterday's victory points gained
	// Required: true
	Yesterday *int32 `json:"yesterday"`
}

/* polymorph GetFwStatsOKBodyItems0VictoryPoints last_week false */

/* polymorph GetFwStatsOKBodyItems0VictoryPoints total false */

/* polymorph GetFwStatsOKBodyItems0VictoryPoints yesterday false */

// Validate validates this get fw stats o k body items0 victory points
func (o *GetFwStatsOKBodyItems0VictoryPoints) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLastWeek(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateTotal(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateYesterday(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetFwStatsOKBodyItems0VictoryPoints) validateLastWeek(formats strfmt.Registry) error {

	if err := validate.Required("victory_points"+"."+"last_week", "body", o.LastWeek); err != nil {
		return err
	}

	return nil
}

func (o *GetFwStatsOKBodyItems0VictoryPoints) validateTotal(formats strfmt.Registry) error {

	if err := validate.Required("victory_points"+"."+"total", "body", o.Total); err != nil {
		return err
	}

	return nil
}

func (o *GetFwStatsOKBodyItems0VictoryPoints) validateYesterday(formats strfmt.Registry) error {

	if err := validate.Required("victory_points"+"."+"yesterday", "body", o.Yesterday); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetFwStatsOKBodyItems0VictoryPoints) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFwStatsOKBodyItems0VictoryPoints) UnmarshalBinary(b []byte) error {
	var res GetFwStatsOKBodyItems0VictoryPoints
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
