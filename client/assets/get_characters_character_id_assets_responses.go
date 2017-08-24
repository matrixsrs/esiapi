// Code generated by go-swagger; DO NOT EDIT.

package assets

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

// GetCharactersCharacterIDAssetsReader is a Reader for the GetCharactersCharacterIDAssets structure.
type GetCharactersCharacterIDAssetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDAssetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDAssetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCharactersCharacterIDAssetsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDAssetsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDAssetsOK creates a GetCharactersCharacterIDAssetsOK with default headers values
func NewGetCharactersCharacterIDAssetsOK() *GetCharactersCharacterIDAssetsOK {
	return &GetCharactersCharacterIDAssetsOK{}
}

/*GetCharactersCharacterIDAssetsOK handles this case with default header values.

A flat list of the users assets
*/
type GetCharactersCharacterIDAssetsOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []*GetCharactersCharacterIDAssetsOKBodyItems0
}

func (o *GetCharactersCharacterIDAssetsOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/assets/][%d] getCharactersCharacterIdAssetsOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDAssetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDAssetsForbidden creates a GetCharactersCharacterIDAssetsForbidden with default headers values
func NewGetCharactersCharacterIDAssetsForbidden() *GetCharactersCharacterIDAssetsForbidden {
	return &GetCharactersCharacterIDAssetsForbidden{}
}

/*GetCharactersCharacterIDAssetsForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDAssetsForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDAssetsForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/assets/][%d] getCharactersCharacterIdAssetsForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDAssetsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDAssetsInternalServerError creates a GetCharactersCharacterIDAssetsInternalServerError with default headers values
func NewGetCharactersCharacterIDAssetsInternalServerError() *GetCharactersCharacterIDAssetsInternalServerError {
	return &GetCharactersCharacterIDAssetsInternalServerError{}
}

/*GetCharactersCharacterIDAssetsInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDAssetsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDAssetsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/assets/][%d] getCharactersCharacterIdAssetsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDAssetsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDAssetsOKBodyItems0 get_characters_character_id_assets_200_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDAssetsOKBodyItems0
*/

type GetCharactersCharacterIDAssetsOKBodyItems0 struct {

	// get_characters_character_id_assets_is_singleton
	//
	// is_singleton boolean
	// Required: true
	IsSingleton *bool `json:"is_singleton"`

	// get_characters_character_id_assets_item_id
	//
	// item_id integer
	// Required: true
	ItemID *int64 `json:"item_id"`

	// get_characters_character_id_assets_location_flag
	//
	// location_flag string
	// Required: true
	LocationFlag *string `json:"location_flag"`

	// get_characters_character_id_assets_location_id
	//
	// location_id integer
	// Required: true
	LocationID *int64 `json:"location_id"`

	// get_characters_character_id_assets_location_type
	//
	// location_type string
	// Required: true
	LocationType *string `json:"location_type"`

	// get_characters_character_id_assets_quantity
	//
	// quantity integer
	Quantity int32 `json:"quantity,omitempty"`

	// get_characters_character_id_assets_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`
}

/* polymorph GetCharactersCharacterIDAssetsOKBodyItems0 is_singleton false */

/* polymorph GetCharactersCharacterIDAssetsOKBodyItems0 item_id false */

/* polymorph GetCharactersCharacterIDAssetsOKBodyItems0 location_flag false */

/* polymorph GetCharactersCharacterIDAssetsOKBodyItems0 location_id false */

/* polymorph GetCharactersCharacterIDAssetsOKBodyItems0 location_type false */

/* polymorph GetCharactersCharacterIDAssetsOKBodyItems0 quantity false */

/* polymorph GetCharactersCharacterIDAssetsOKBodyItems0 type_id false */

// Validate validates this get characters character ID assets o k body items0
func (o *GetCharactersCharacterIDAssetsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIsSingleton(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateItemID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateLocationFlag(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateLocationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateLocationType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateTypeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCharactersCharacterIDAssetsOKBodyItems0) validateIsSingleton(formats strfmt.Registry) error {

	if err := validate.Required("is_singleton", "body", o.IsSingleton); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDAssetsOKBodyItems0) validateItemID(formats strfmt.Registry) error {

	if err := validate.Required("item_id", "body", o.ItemID); err != nil {
		return err
	}

	return nil
}

var getCharactersCharacterIdAssetsOKBodyItems0TypeLocationFlagPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AutoFit","Cargo","CorpseBay","DroneBay","FleetHangar","Deliveries","HiddenModifiers","Hangar","HangarAll","LoSlot0","LoSlot1","LoSlot2","LoSlot3","LoSlot4","LoSlot5","LoSlot6","LoSlot7","MedSlot0","MedSlot1","MedSlot2","MedSlot3","MedSlot4","MedSlot5","MedSlot6","MedSlot7","HiSlot0","HiSlot1","HiSlot2","HiSlot3","HiSlot4","HiSlot5","HiSlot6","HiSlot7","AssetSafety","Locked","Unlocked","Implant","QuafeBay","RigSlot0","RigSlot1","RigSlot2","RigSlot3","RigSlot4","RigSlot5","RigSlot6","RigSlot7","ShipHangar","SpecializedFuelBay","SpecializedOreHold","SpecializedGasHold","SpecializedMineralHold","SpecializedSalvageHold","SpecializedShipHold","SpecializedSmallShipHold","SpecializedMediumShipHold","SpecializedLargeShipHold","SpecializedIndustrialShipHold","SpecializedAmmoHold","SpecializedCommandCenterHold","SpecializedPlanetaryCommoditiesHold","SpecializedMaterialBay","SubSystemBay","SubSystemSlot0","SubSystemSlot1","SubSystemSlot2","SubSystemSlot3","SubSystemSlot4","SubSystemSlot5","SubSystemSlot6","SubSystemSlot7","FighterBay","FighterTube0","FighterTube1","FighterTube2","FighterTube3","FighterTube4","Module","Wardrobe"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdAssetsOKBodyItems0TypeLocationFlagPropEnum = append(getCharactersCharacterIdAssetsOKBodyItems0TypeLocationFlagPropEnum, v)
	}
}

const (
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagAutoFit captures enum value "AutoFit"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagAutoFit string = "AutoFit"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagCargo captures enum value "Cargo"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagCargo string = "Cargo"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagCorpseBay captures enum value "CorpseBay"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagCorpseBay string = "CorpseBay"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagDroneBay captures enum value "DroneBay"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagDroneBay string = "DroneBay"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagFleetHangar captures enum value "FleetHangar"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagFleetHangar string = "FleetHangar"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagDeliveries captures enum value "Deliveries"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagDeliveries string = "Deliveries"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagHiddenModifiers captures enum value "HiddenModifiers"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagHiddenModifiers string = "HiddenModifiers"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagHangar captures enum value "Hangar"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagHangar string = "Hangar"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagHangarAll captures enum value "HangarAll"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagHangarAll string = "HangarAll"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagLoSlot0 captures enum value "LoSlot0"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagLoSlot0 string = "LoSlot0"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagLoSlot1 captures enum value "LoSlot1"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagLoSlot1 string = "LoSlot1"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagLoSlot2 captures enum value "LoSlot2"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagLoSlot2 string = "LoSlot2"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagLoSlot3 captures enum value "LoSlot3"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagLoSlot3 string = "LoSlot3"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagLoSlot4 captures enum value "LoSlot4"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagLoSlot4 string = "LoSlot4"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagLoSlot5 captures enum value "LoSlot5"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagLoSlot5 string = "LoSlot5"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagLoSlot6 captures enum value "LoSlot6"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagLoSlot6 string = "LoSlot6"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagLoSlot7 captures enum value "LoSlot7"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagLoSlot7 string = "LoSlot7"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagMedSlot0 captures enum value "MedSlot0"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagMedSlot0 string = "MedSlot0"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagMedSlot1 captures enum value "MedSlot1"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagMedSlot1 string = "MedSlot1"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagMedSlot2 captures enum value "MedSlot2"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagMedSlot2 string = "MedSlot2"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagMedSlot3 captures enum value "MedSlot3"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagMedSlot3 string = "MedSlot3"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagMedSlot4 captures enum value "MedSlot4"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagMedSlot4 string = "MedSlot4"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagMedSlot5 captures enum value "MedSlot5"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagMedSlot5 string = "MedSlot5"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagMedSlot6 captures enum value "MedSlot6"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagMedSlot6 string = "MedSlot6"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagMedSlot7 captures enum value "MedSlot7"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagMedSlot7 string = "MedSlot7"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagHiSlot0 captures enum value "HiSlot0"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagHiSlot0 string = "HiSlot0"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagHiSlot1 captures enum value "HiSlot1"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagHiSlot1 string = "HiSlot1"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagHiSlot2 captures enum value "HiSlot2"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagHiSlot2 string = "HiSlot2"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagHiSlot3 captures enum value "HiSlot3"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagHiSlot3 string = "HiSlot3"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagHiSlot4 captures enum value "HiSlot4"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagHiSlot4 string = "HiSlot4"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagHiSlot5 captures enum value "HiSlot5"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagHiSlot5 string = "HiSlot5"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagHiSlot6 captures enum value "HiSlot6"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagHiSlot6 string = "HiSlot6"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagHiSlot7 captures enum value "HiSlot7"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagHiSlot7 string = "HiSlot7"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagAssetSafety captures enum value "AssetSafety"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagAssetSafety string = "AssetSafety"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagLocked captures enum value "Locked"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagLocked string = "Locked"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagUnlocked captures enum value "Unlocked"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagUnlocked string = "Unlocked"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagImplant captures enum value "Implant"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagImplant string = "Implant"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagQuafeBay captures enum value "QuafeBay"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagQuafeBay string = "QuafeBay"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagRigSlot0 captures enum value "RigSlot0"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagRigSlot0 string = "RigSlot0"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagRigSlot1 captures enum value "RigSlot1"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagRigSlot1 string = "RigSlot1"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagRigSlot2 captures enum value "RigSlot2"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagRigSlot2 string = "RigSlot2"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagRigSlot3 captures enum value "RigSlot3"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagRigSlot3 string = "RigSlot3"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagRigSlot4 captures enum value "RigSlot4"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagRigSlot4 string = "RigSlot4"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagRigSlot5 captures enum value "RigSlot5"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagRigSlot5 string = "RigSlot5"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagRigSlot6 captures enum value "RigSlot6"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagRigSlot6 string = "RigSlot6"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagRigSlot7 captures enum value "RigSlot7"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagRigSlot7 string = "RigSlot7"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagShipHangar captures enum value "ShipHangar"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagShipHangar string = "ShipHangar"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSpecializedFuelBay captures enum value "SpecializedFuelBay"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSpecializedFuelBay string = "SpecializedFuelBay"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSpecializedOreHold captures enum value "SpecializedOreHold"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSpecializedOreHold string = "SpecializedOreHold"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSpecializedGasHold captures enum value "SpecializedGasHold"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSpecializedGasHold string = "SpecializedGasHold"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSpecializedMineralHold captures enum value "SpecializedMineralHold"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSpecializedMineralHold string = "SpecializedMineralHold"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSpecializedSalvageHold captures enum value "SpecializedSalvageHold"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSpecializedSalvageHold string = "SpecializedSalvageHold"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSpecializedShipHold captures enum value "SpecializedShipHold"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSpecializedShipHold string = "SpecializedShipHold"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSpecializedSmallShipHold captures enum value "SpecializedSmallShipHold"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSpecializedSmallShipHold string = "SpecializedSmallShipHold"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSpecializedMediumShipHold captures enum value "SpecializedMediumShipHold"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSpecializedMediumShipHold string = "SpecializedMediumShipHold"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSpecializedLargeShipHold captures enum value "SpecializedLargeShipHold"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSpecializedLargeShipHold string = "SpecializedLargeShipHold"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSpecializedIndustrialShipHold captures enum value "SpecializedIndustrialShipHold"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSpecializedIndustrialShipHold string = "SpecializedIndustrialShipHold"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSpecializedAmmoHold captures enum value "SpecializedAmmoHold"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSpecializedAmmoHold string = "SpecializedAmmoHold"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSpecializedCommandCenterHold captures enum value "SpecializedCommandCenterHold"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSpecializedCommandCenterHold string = "SpecializedCommandCenterHold"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSpecializedPlanetaryCommoditiesHold captures enum value "SpecializedPlanetaryCommoditiesHold"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSpecializedPlanetaryCommoditiesHold string = "SpecializedPlanetaryCommoditiesHold"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSpecializedMaterialBay captures enum value "SpecializedMaterialBay"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSpecializedMaterialBay string = "SpecializedMaterialBay"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSubSystemBay captures enum value "SubSystemBay"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSubSystemBay string = "SubSystemBay"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSubSystemSlot0 captures enum value "SubSystemSlot0"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSubSystemSlot0 string = "SubSystemSlot0"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSubSystemSlot1 captures enum value "SubSystemSlot1"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSubSystemSlot1 string = "SubSystemSlot1"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSubSystemSlot2 captures enum value "SubSystemSlot2"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSubSystemSlot2 string = "SubSystemSlot2"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSubSystemSlot3 captures enum value "SubSystemSlot3"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSubSystemSlot3 string = "SubSystemSlot3"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSubSystemSlot4 captures enum value "SubSystemSlot4"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSubSystemSlot4 string = "SubSystemSlot4"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSubSystemSlot5 captures enum value "SubSystemSlot5"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSubSystemSlot5 string = "SubSystemSlot5"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSubSystemSlot6 captures enum value "SubSystemSlot6"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSubSystemSlot6 string = "SubSystemSlot6"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSubSystemSlot7 captures enum value "SubSystemSlot7"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagSubSystemSlot7 string = "SubSystemSlot7"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagFighterBay captures enum value "FighterBay"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagFighterBay string = "FighterBay"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagFighterTube0 captures enum value "FighterTube0"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagFighterTube0 string = "FighterTube0"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagFighterTube1 captures enum value "FighterTube1"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagFighterTube1 string = "FighterTube1"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagFighterTube2 captures enum value "FighterTube2"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagFighterTube2 string = "FighterTube2"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagFighterTube3 captures enum value "FighterTube3"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagFighterTube3 string = "FighterTube3"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagFighterTube4 captures enum value "FighterTube4"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagFighterTube4 string = "FighterTube4"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagModule captures enum value "Module"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagModule string = "Module"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagWardrobe captures enum value "Wardrobe"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationFlagWardrobe string = "Wardrobe"
)

// prop value enum
func (o *GetCharactersCharacterIDAssetsOKBodyItems0) validateLocationFlagEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCharactersCharacterIdAssetsOKBodyItems0TypeLocationFlagPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetCharactersCharacterIDAssetsOKBodyItems0) validateLocationFlag(formats strfmt.Registry) error {

	if err := validate.Required("location_flag", "body", o.LocationFlag); err != nil {
		return err
	}

	// value enum
	if err := o.validateLocationFlagEnum("location_flag", "body", *o.LocationFlag); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDAssetsOKBodyItems0) validateLocationID(formats strfmt.Registry) error {

	if err := validate.Required("location_id", "body", o.LocationID); err != nil {
		return err
	}

	return nil
}

var getCharactersCharacterIdAssetsOKBodyItems0TypeLocationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["station","solar_system","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdAssetsOKBodyItems0TypeLocationTypePropEnum = append(getCharactersCharacterIdAssetsOKBodyItems0TypeLocationTypePropEnum, v)
	}
}

const (
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationTypeStation captures enum value "station"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationTypeStation string = "station"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationTypeSolarSystem captures enum value "solar_system"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationTypeSolarSystem string = "solar_system"
	// GetCharactersCharacterIDAssetsOKBodyItems0LocationTypeOther captures enum value "other"
	GetCharactersCharacterIDAssetsOKBodyItems0LocationTypeOther string = "other"
)

// prop value enum
func (o *GetCharactersCharacterIDAssetsOKBodyItems0) validateLocationTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCharactersCharacterIdAssetsOKBodyItems0TypeLocationTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetCharactersCharacterIDAssetsOKBodyItems0) validateLocationType(formats strfmt.Registry) error {

	if err := validate.Required("location_type", "body", o.LocationType); err != nil {
		return err
	}

	// value enum
	if err := o.validateLocationTypeEnum("location_type", "body", *o.LocationType); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDAssetsOKBodyItems0) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDAssetsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDAssetsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDAssetsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
