// Code generated by go-swagger; DO NOT EDIT.

package market

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/models"
)

// GetMarketsRegionIDOrdersReader is a Reader for the GetMarketsRegionIDOrders structure.
type GetMarketsRegionIDOrdersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMarketsRegionIDOrdersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMarketsRegionIDOrdersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 422:
		result := NewGetMarketsRegionIDOrdersUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetMarketsRegionIDOrdersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMarketsRegionIDOrdersOK creates a GetMarketsRegionIDOrdersOK with default headers values
func NewGetMarketsRegionIDOrdersOK() *GetMarketsRegionIDOrdersOK {
	return &GetMarketsRegionIDOrdersOK{
		XPages: 1,
	}
}

/*GetMarketsRegionIDOrdersOK handles this case with default header values.

A list of orders
*/
type GetMarketsRegionIDOrdersOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string
	/*Maximum page number
	 */
	XPages int32

	Payload []*GetMarketsRegionIDOrdersOKBodyItems0
}

func (o *GetMarketsRegionIDOrdersOK) Error() string {
	return fmt.Sprintf("[GET /markets/{region_id}/orders/][%d] getMarketsRegionIdOrdersOK  %+v", 200, o.Payload)
}

func (o *GetMarketsRegionIDOrdersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	// response header X-Pages
	xPages, err := swag.ConvertInt32(response.GetHeader("X-Pages"))
	if err != nil {
		return errors.InvalidType("X-Pages", "header", "int32", response.GetHeader("X-Pages"))
	}
	o.XPages = xPages

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarketsRegionIDOrdersUnprocessableEntity creates a GetMarketsRegionIDOrdersUnprocessableEntity with default headers values
func NewGetMarketsRegionIDOrdersUnprocessableEntity() *GetMarketsRegionIDOrdersUnprocessableEntity {
	return &GetMarketsRegionIDOrdersUnprocessableEntity{}
}

/*GetMarketsRegionIDOrdersUnprocessableEntity handles this case with default header values.

Not found
*/
type GetMarketsRegionIDOrdersUnprocessableEntity struct {
	Payload GetMarketsRegionIDOrdersUnprocessableEntityBody
}

func (o *GetMarketsRegionIDOrdersUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /markets/{region_id}/orders/][%d] getMarketsRegionIdOrdersUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetMarketsRegionIDOrdersUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarketsRegionIDOrdersInternalServerError creates a GetMarketsRegionIDOrdersInternalServerError with default headers values
func NewGetMarketsRegionIDOrdersInternalServerError() *GetMarketsRegionIDOrdersInternalServerError {
	return &GetMarketsRegionIDOrdersInternalServerError{}
}

/*GetMarketsRegionIDOrdersInternalServerError handles this case with default header values.

Internal server error
*/
type GetMarketsRegionIDOrdersInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetMarketsRegionIDOrdersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /markets/{region_id}/orders/][%d] getMarketsRegionIdOrdersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetMarketsRegionIDOrdersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetMarketsRegionIDOrdersOKBodyItems0 get_markets_region_id_orders_200_ok
//
// 200 ok object
swagger:model GetMarketsRegionIDOrdersOKBodyItems0
*/

type GetMarketsRegionIDOrdersOKBodyItems0 struct {

	// get_markets_region_id_orders_duration
	//
	// duration integer
	// Required: true
	Duration *int32 `json:"duration"`

	// get_markets_region_id_orders_is_buy_order
	//
	// is_buy_order boolean
	// Required: true
	IsBuyOrder *bool `json:"is_buy_order"`

	// get_markets_region_id_orders_issued
	//
	// issued string
	// Required: true
	Issued *strfmt.DateTime `json:"issued"`

	// get_markets_region_id_orders_location_id
	//
	// location_id integer
	// Required: true
	LocationID *int64 `json:"location_id"`

	// get_markets_region_id_orders_min_volume
	//
	// min_volume integer
	// Required: true
	MinVolume *int32 `json:"min_volume"`

	// get_markets_region_id_orders_order_id
	//
	// order_id integer
	// Required: true
	OrderID *int64 `json:"order_id"`

	// get_markets_region_id_orders_price
	//
	// price number
	// Required: true
	Price *float32 `json:"price"`

	// get_markets_region_id_orders_range
	//
	// range string
	// Required: true
	Range *string `json:"range"`

	// get_markets_region_id_orders_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`

	// get_markets_region_id_orders_volume_remain
	//
	// volume_remain integer
	// Required: true
	VolumeRemain *int32 `json:"volume_remain"`

	// get_markets_region_id_orders_volume_total
	//
	// volume_total integer
	// Required: true
	VolumeTotal *int32 `json:"volume_total"`
}

/* polymorph GetMarketsRegionIDOrdersOKBodyItems0 duration false */

/* polymorph GetMarketsRegionIDOrdersOKBodyItems0 is_buy_order false */

/* polymorph GetMarketsRegionIDOrdersOKBodyItems0 issued false */

/* polymorph GetMarketsRegionIDOrdersOKBodyItems0 location_id false */

/* polymorph GetMarketsRegionIDOrdersOKBodyItems0 min_volume false */

/* polymorph GetMarketsRegionIDOrdersOKBodyItems0 order_id false */

/* polymorph GetMarketsRegionIDOrdersOKBodyItems0 price false */

/* polymorph GetMarketsRegionIDOrdersOKBodyItems0 range false */

/* polymorph GetMarketsRegionIDOrdersOKBodyItems0 type_id false */

/* polymorph GetMarketsRegionIDOrdersOKBodyItems0 volume_remain false */

/* polymorph GetMarketsRegionIDOrdersOKBodyItems0 volume_total false */

// Validate validates this get markets region ID orders o k body items0
func (o *GetMarketsRegionIDOrdersOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDuration(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateIsBuyOrder(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateIssued(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateLocationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateMinVolume(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateOrderID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validatePrice(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateRange(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateTypeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateVolumeRemain(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateVolumeTotal(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetMarketsRegionIDOrdersOKBodyItems0) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("duration", "body", o.Duration); err != nil {
		return err
	}

	return nil
}

func (o *GetMarketsRegionIDOrdersOKBodyItems0) validateIsBuyOrder(formats strfmt.Registry) error {

	if err := validate.Required("is_buy_order", "body", o.IsBuyOrder); err != nil {
		return err
	}

	return nil
}

func (o *GetMarketsRegionIDOrdersOKBodyItems0) validateIssued(formats strfmt.Registry) error {

	if err := validate.Required("issued", "body", o.Issued); err != nil {
		return err
	}

	if err := validate.FormatOf("issued", "body", "date-time", o.Issued.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetMarketsRegionIDOrdersOKBodyItems0) validateLocationID(formats strfmt.Registry) error {

	if err := validate.Required("location_id", "body", o.LocationID); err != nil {
		return err
	}

	return nil
}

func (o *GetMarketsRegionIDOrdersOKBodyItems0) validateMinVolume(formats strfmt.Registry) error {

	if err := validate.Required("min_volume", "body", o.MinVolume); err != nil {
		return err
	}

	return nil
}

func (o *GetMarketsRegionIDOrdersOKBodyItems0) validateOrderID(formats strfmt.Registry) error {

	if err := validate.Required("order_id", "body", o.OrderID); err != nil {
		return err
	}

	return nil
}

func (o *GetMarketsRegionIDOrdersOKBodyItems0) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", o.Price); err != nil {
		return err
	}

	return nil
}

var getMarketsRegionIdOrdersOKBodyItems0TypeRangePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["station","region","solarsystem","1","2","3","4","5","10","20","30","40"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getMarketsRegionIdOrdersOKBodyItems0TypeRangePropEnum = append(getMarketsRegionIdOrdersOKBodyItems0TypeRangePropEnum, v)
	}
}

const (
	// GetMarketsRegionIDOrdersOKBodyItems0RangeStation captures enum value "station"
	GetMarketsRegionIDOrdersOKBodyItems0RangeStation string = "station"
	// GetMarketsRegionIDOrdersOKBodyItems0RangeRegion captures enum value "region"
	GetMarketsRegionIDOrdersOKBodyItems0RangeRegion string = "region"
	// GetMarketsRegionIDOrdersOKBodyItems0RangeSolarsystem captures enum value "solarsystem"
	GetMarketsRegionIDOrdersOKBodyItems0RangeSolarsystem string = "solarsystem"
	// GetMarketsRegionIDOrdersOKBodyItems0RangeNr1 captures enum value "1"
	GetMarketsRegionIDOrdersOKBodyItems0RangeNr1 string = "1"
	// GetMarketsRegionIDOrdersOKBodyItems0RangeNr2 captures enum value "2"
	GetMarketsRegionIDOrdersOKBodyItems0RangeNr2 string = "2"
	// GetMarketsRegionIDOrdersOKBodyItems0RangeNr3 captures enum value "3"
	GetMarketsRegionIDOrdersOKBodyItems0RangeNr3 string = "3"
	// GetMarketsRegionIDOrdersOKBodyItems0RangeNr4 captures enum value "4"
	GetMarketsRegionIDOrdersOKBodyItems0RangeNr4 string = "4"
	// GetMarketsRegionIDOrdersOKBodyItems0RangeNr5 captures enum value "5"
	GetMarketsRegionIDOrdersOKBodyItems0RangeNr5 string = "5"
	// GetMarketsRegionIDOrdersOKBodyItems0RangeNr10 captures enum value "10"
	GetMarketsRegionIDOrdersOKBodyItems0RangeNr10 string = "10"
	// GetMarketsRegionIDOrdersOKBodyItems0RangeNr20 captures enum value "20"
	GetMarketsRegionIDOrdersOKBodyItems0RangeNr20 string = "20"
	// GetMarketsRegionIDOrdersOKBodyItems0RangeNr30 captures enum value "30"
	GetMarketsRegionIDOrdersOKBodyItems0RangeNr30 string = "30"
	// GetMarketsRegionIDOrdersOKBodyItems0RangeNr40 captures enum value "40"
	GetMarketsRegionIDOrdersOKBodyItems0RangeNr40 string = "40"
)

// prop value enum
func (o *GetMarketsRegionIDOrdersOKBodyItems0) validateRangeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getMarketsRegionIdOrdersOKBodyItems0TypeRangePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetMarketsRegionIDOrdersOKBodyItems0) validateRange(formats strfmt.Registry) error {

	if err := validate.Required("range", "body", o.Range); err != nil {
		return err
	}

	// value enum
	if err := o.validateRangeEnum("range", "body", *o.Range); err != nil {
		return err
	}

	return nil
}

func (o *GetMarketsRegionIDOrdersOKBodyItems0) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

func (o *GetMarketsRegionIDOrdersOKBodyItems0) validateVolumeRemain(formats strfmt.Registry) error {

	if err := validate.Required("volume_remain", "body", o.VolumeRemain); err != nil {
		return err
	}

	return nil
}

func (o *GetMarketsRegionIDOrdersOKBodyItems0) validateVolumeTotal(formats strfmt.Registry) error {

	if err := validate.Required("volume_total", "body", o.VolumeTotal); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetMarketsRegionIDOrdersOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetMarketsRegionIDOrdersOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetMarketsRegionIDOrdersOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetMarketsRegionIDOrdersUnprocessableEntityBody get_markets_region_id_orders_unprocessable_entity
//
// Unprocessable entity
swagger:model GetMarketsRegionIDOrdersUnprocessableEntityBody
*/

type GetMarketsRegionIDOrdersUnprocessableEntityBody struct {

	// get_markets_region_id_orders_422_unprocessable_entity
	//
	// Unprocessable entity message
	// Required: true
	Error *string `json:"error"`
}

/* polymorph GetMarketsRegionIDOrdersUnprocessableEntityBody error false */

// Validate validates this get markets region ID orders unprocessable entity body
func (o *GetMarketsRegionIDOrdersUnprocessableEntityBody) Validate(formats strfmt.Registry) error {
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

func (o *GetMarketsRegionIDOrdersUnprocessableEntityBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getMarketsRegionIdOrdersUnprocessableEntity"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetMarketsRegionIDOrdersUnprocessableEntityBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetMarketsRegionIDOrdersUnprocessableEntityBody) UnmarshalBinary(b []byte) error {
	var res GetMarketsRegionIDOrdersUnprocessableEntityBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
