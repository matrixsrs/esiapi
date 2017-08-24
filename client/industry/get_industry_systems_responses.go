// Code generated by go-swagger; DO NOT EDIT.

package industry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
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

// GetIndustrySystemsReader is a Reader for the GetIndustrySystems structure.
type GetIndustrySystemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIndustrySystemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIndustrySystemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetIndustrySystemsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetIndustrySystemsOK creates a GetIndustrySystemsOK with default headers values
func NewGetIndustrySystemsOK() *GetIndustrySystemsOK {
	return &GetIndustrySystemsOK{}
}

/*GetIndustrySystemsOK handles this case with default header values.

A list of cost indicies
*/
type GetIndustrySystemsOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []*GetIndustrySystemsOKBodyItems0
}

func (o *GetIndustrySystemsOK) Error() string {
	return fmt.Sprintf("[GET /industry/systems/][%d] getIndustrySystemsOK  %+v", 200, o.Payload)
}

func (o *GetIndustrySystemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIndustrySystemsInternalServerError creates a GetIndustrySystemsInternalServerError with default headers values
func NewGetIndustrySystemsInternalServerError() *GetIndustrySystemsInternalServerError {
	return &GetIndustrySystemsInternalServerError{}
}

/*GetIndustrySystemsInternalServerError handles this case with default header values.

Internal server error
*/
type GetIndustrySystemsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetIndustrySystemsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /industry/systems/][%d] getIndustrySystemsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIndustrySystemsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetIndustrySystemsOKBodyItems0 get_industry_systems_200_ok
//
// 200 ok object
swagger:model GetIndustrySystemsOKBodyItems0
*/

type GetIndustrySystemsOKBodyItems0 struct {

	// get_industry_systems_cost_indices
	//
	// cost_indices array
	// Required: true
	// Max Items: 10
	CostIndices []*GetIndustrySystemsOKBodyItems0CostIndicesItems0 `json:"cost_indices"`

	// get_industry_systems_solar_system_id
	//
	// solar_system_id integer
	// Required: true
	SolarSystemID *int32 `json:"solar_system_id"`
}

/* polymorph GetIndustrySystemsOKBodyItems0 cost_indices false */

/* polymorph GetIndustrySystemsOKBodyItems0 solar_system_id false */

// Validate validates this get industry systems o k body items0
func (o *GetIndustrySystemsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCostIndices(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateSolarSystemID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetIndustrySystemsOKBodyItems0) validateCostIndices(formats strfmt.Registry) error {

	if err := validate.Required("cost_indices", "body", o.CostIndices); err != nil {
		return err
	}

	iCostIndicesSize := int64(len(o.CostIndices))

	if err := validate.MaxItems("cost_indices", "body", iCostIndicesSize, 10); err != nil {
		return err
	}

	for i := 0; i < len(o.CostIndices); i++ {

		if swag.IsZero(o.CostIndices[i]) { // not required
			continue
		}

		if o.CostIndices[i] != nil {

			if err := o.CostIndices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cost_indices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetIndustrySystemsOKBodyItems0) validateSolarSystemID(formats strfmt.Registry) error {

	if err := validate.Required("solar_system_id", "body", o.SolarSystemID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetIndustrySystemsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetIndustrySystemsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetIndustrySystemsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetIndustrySystemsOKBodyItems0CostIndicesItems0 get_industry_systems_cost_indice
//
// cost_indice object
swagger:model GetIndustrySystemsOKBodyItems0CostIndicesItems0
*/

type GetIndustrySystemsOKBodyItems0CostIndicesItems0 struct {

	// get_industry_systems_activity
	//
	// activity string
	// Required: true
	Activity *string `json:"activity"`

	// get_industry_systems_cost_index
	//
	// cost_index number
	// Required: true
	CostIndex *float32 `json:"cost_index"`
}

/* polymorph GetIndustrySystemsOKBodyItems0CostIndicesItems0 activity false */

/* polymorph GetIndustrySystemsOKBodyItems0CostIndicesItems0 cost_index false */

// Validate validates this get industry systems o k body items0 cost indices items0
func (o *GetIndustrySystemsOKBodyItems0CostIndicesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateActivity(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateCostIndex(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getIndustrySystemsOKBodyItems0CostIndicesItems0TypeActivityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","manufacturing","researching_technology","researching_time_efficiency","researching_material_efficiency","copying","duplicating","invention","reverse_engineering"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getIndustrySystemsOKBodyItems0CostIndicesItems0TypeActivityPropEnum = append(getIndustrySystemsOKBodyItems0CostIndicesItems0TypeActivityPropEnum, v)
	}
}

const (
	// GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityNone captures enum value "none"
	GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityNone string = "none"
	// GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityManufacturing captures enum value "manufacturing"
	GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityManufacturing string = "manufacturing"
	// GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityResearchingTechnology captures enum value "researching_technology"
	GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityResearchingTechnology string = "researching_technology"
	// GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityResearchingTimeEfficiency captures enum value "researching_time_efficiency"
	GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityResearchingTimeEfficiency string = "researching_time_efficiency"
	// GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityResearchingMaterialEfficiency captures enum value "researching_material_efficiency"
	GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityResearchingMaterialEfficiency string = "researching_material_efficiency"
	// GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityCopying captures enum value "copying"
	GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityCopying string = "copying"
	// GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityDuplicating captures enum value "duplicating"
	GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityDuplicating string = "duplicating"
	// GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityInvention captures enum value "invention"
	GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityInvention string = "invention"
	// GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityReverseEngineering captures enum value "reverse_engineering"
	GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityReverseEngineering string = "reverse_engineering"
)

// prop value enum
func (o *GetIndustrySystemsOKBodyItems0CostIndicesItems0) validateActivityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getIndustrySystemsOKBodyItems0CostIndicesItems0TypeActivityPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetIndustrySystemsOKBodyItems0CostIndicesItems0) validateActivity(formats strfmt.Registry) error {

	if err := validate.Required("activity", "body", o.Activity); err != nil {
		return err
	}

	// value enum
	if err := o.validateActivityEnum("activity", "body", *o.Activity); err != nil {
		return err
	}

	return nil
}

func (o *GetIndustrySystemsOKBodyItems0CostIndicesItems0) validateCostIndex(formats strfmt.Registry) error {

	if err := validate.Required("cost_index", "body", o.CostIndex); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetIndustrySystemsOKBodyItems0CostIndicesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetIndustrySystemsOKBodyItems0CostIndicesItems0) UnmarshalBinary(b []byte) error {
	var res GetIndustrySystemsOKBodyItems0CostIndicesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
