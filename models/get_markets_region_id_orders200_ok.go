package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// GetMarketsRegionIDOrders200Ok 200 ok object
// swagger:model get_markets_region_id_orders_200_ok
type GetMarketsRegionIDOrders200Ok struct {

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
	MinVolume *int64 `json:"min_volume"`

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
	TypeID *int64 `json:"type_id"`

	// get_markets_region_id_orders_volume_remain
	//
	// volume_remain integer
	// Required: true
	VolumeRemain *int64 `json:"volume_remain"`

	// get_markets_region_id_orders_volume_total
	//
	// volume_total integer
	// Required: true
	VolumeTotal *int64 `json:"volume_total"`
}

// Validate validates this get markets region id orders 200 ok
func (m *GetMarketsRegionIDOrders200Ok) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDuration(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIsBuyOrder(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIssued(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLocationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMinVolume(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrderID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePrice(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRange(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTypeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVolumeRemain(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVolumeTotal(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetMarketsRegionIDOrders200Ok) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *GetMarketsRegionIDOrders200Ok) validateIsBuyOrder(formats strfmt.Registry) error {

	if err := validate.Required("is_buy_order", "body", m.IsBuyOrder); err != nil {
		return err
	}

	return nil
}

func (m *GetMarketsRegionIDOrders200Ok) validateIssued(formats strfmt.Registry) error {

	if err := validate.Required("issued", "body", m.Issued); err != nil {
		return err
	}

	return nil
}

func (m *GetMarketsRegionIDOrders200Ok) validateLocationID(formats strfmt.Registry) error {

	if err := validate.Required("location_id", "body", m.LocationID); err != nil {
		return err
	}

	return nil
}

func (m *GetMarketsRegionIDOrders200Ok) validateMinVolume(formats strfmt.Registry) error {

	if err := validate.Required("min_volume", "body", m.MinVolume); err != nil {
		return err
	}

	return nil
}

func (m *GetMarketsRegionIDOrders200Ok) validateOrderID(formats strfmt.Registry) error {

	if err := validate.Required("order_id", "body", m.OrderID); err != nil {
		return err
	}

	return nil
}

func (m *GetMarketsRegionIDOrders200Ok) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", m.Price); err != nil {
		return err
	}

	return nil
}

var getMarketsRegionIdOrders200OkTypeRangePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["station","region","solarsystem","1","2","3","4","5","10","20","30","40"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getMarketsRegionIdOrders200OkTypeRangePropEnum = append(getMarketsRegionIdOrders200OkTypeRangePropEnum, v)
	}
}

const (
	GetMarketsRegionIDOrders200OkRangeStation     string = "station"
	GetMarketsRegionIDOrders200OkRangeRegion      string = "region"
	GetMarketsRegionIDOrders200OkRangeSolarsystem string = "solarsystem"
	GetMarketsRegionIDOrders200OkRangeNr1         string = "1"
	GetMarketsRegionIDOrders200OkRangeNr2         string = "2"
	GetMarketsRegionIDOrders200OkRangeNr3         string = "3"
	GetMarketsRegionIDOrders200OkRangeNr4         string = "4"
	GetMarketsRegionIDOrders200OkRangeNr5         string = "5"
	GetMarketsRegionIDOrders200OkRangeNr10        string = "10"
	GetMarketsRegionIDOrders200OkRangeNr20        string = "20"
	GetMarketsRegionIDOrders200OkRangeNr30        string = "30"
	GetMarketsRegionIDOrders200OkRangeNr40        string = "40"
)

// prop value enum
func (m *GetMarketsRegionIDOrders200Ok) validateRangeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getMarketsRegionIdOrders200OkTypeRangePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetMarketsRegionIDOrders200Ok) validateRange(formats strfmt.Registry) error {

	if err := validate.Required("range", "body", m.Range); err != nil {
		return err
	}

	// value enum
	if err := m.validateRangeEnum("range", "body", *m.Range); err != nil {
		return err
	}

	return nil
}

func (m *GetMarketsRegionIDOrders200Ok) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", m.TypeID); err != nil {
		return err
	}

	return nil
}

func (m *GetMarketsRegionIDOrders200Ok) validateVolumeRemain(formats strfmt.Registry) error {

	if err := validate.Required("volume_remain", "body", m.VolumeRemain); err != nil {
		return err
	}

	return nil
}

func (m *GetMarketsRegionIDOrders200Ok) validateVolumeTotal(formats strfmt.Registry) error {

	if err := validate.Required("volume_total", "body", m.VolumeTotal); err != nil {
		return err
	}

	return nil
}
