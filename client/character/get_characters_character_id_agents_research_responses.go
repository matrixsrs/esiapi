// Code generated by go-swagger; DO NOT EDIT.

package character

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

// GetCharactersCharacterIDAgentsResearchReader is a Reader for the GetCharactersCharacterIDAgentsResearch structure.
type GetCharactersCharacterIDAgentsResearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDAgentsResearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDAgentsResearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCharactersCharacterIDAgentsResearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDAgentsResearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDAgentsResearchOK creates a GetCharactersCharacterIDAgentsResearchOK with default headers values
func NewGetCharactersCharacterIDAgentsResearchOK() *GetCharactersCharacterIDAgentsResearchOK {
	return &GetCharactersCharacterIDAgentsResearchOK{}
}

/*GetCharactersCharacterIDAgentsResearchOK handles this case with default header values.

A list of agents research information
*/
type GetCharactersCharacterIDAgentsResearchOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []*GetCharactersCharacterIDAgentsResearchOKBodyItems0
}

func (o *GetCharactersCharacterIDAgentsResearchOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/agents_research/][%d] getCharactersCharacterIdAgentsResearchOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDAgentsResearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDAgentsResearchForbidden creates a GetCharactersCharacterIDAgentsResearchForbidden with default headers values
func NewGetCharactersCharacterIDAgentsResearchForbidden() *GetCharactersCharacterIDAgentsResearchForbidden {
	return &GetCharactersCharacterIDAgentsResearchForbidden{}
}

/*GetCharactersCharacterIDAgentsResearchForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDAgentsResearchForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDAgentsResearchForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/agents_research/][%d] getCharactersCharacterIdAgentsResearchForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDAgentsResearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDAgentsResearchInternalServerError creates a GetCharactersCharacterIDAgentsResearchInternalServerError with default headers values
func NewGetCharactersCharacterIDAgentsResearchInternalServerError() *GetCharactersCharacterIDAgentsResearchInternalServerError {
	return &GetCharactersCharacterIDAgentsResearchInternalServerError{}
}

/*GetCharactersCharacterIDAgentsResearchInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDAgentsResearchInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDAgentsResearchInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/agents_research/][%d] getCharactersCharacterIdAgentsResearchInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDAgentsResearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDAgentsResearchOKBodyItems0 get_characters_character_id_agents_research_200_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDAgentsResearchOKBodyItems0
*/

type GetCharactersCharacterIDAgentsResearchOKBodyItems0 struct {

	// get_characters_character_id_agents_research_agent_id
	//
	// agent_id integer
	// Required: true
	AgentID *int32 `json:"agent_id"`

	// get_characters_character_id_agents_research_points_per_day
	//
	// points_per_day number
	// Required: true
	PointsPerDay *float32 `json:"points_per_day"`

	// get_characters_character_id_agents_research_remainder_points
	//
	// remainder_points number
	// Required: true
	RemainderPoints *float32 `json:"remainder_points"`

	// get_characters_character_id_agents_research_skill_type_id
	//
	// skill_type_id integer
	// Required: true
	SkillTypeID *int32 `json:"skill_type_id"`

	// get_characters_character_id_agents_research_started_at
	//
	// started_at string
	// Required: true
	StartedAt *strfmt.DateTime `json:"started_at"`
}

/* polymorph GetCharactersCharacterIDAgentsResearchOKBodyItems0 agent_id false */

/* polymorph GetCharactersCharacterIDAgentsResearchOKBodyItems0 points_per_day false */

/* polymorph GetCharactersCharacterIDAgentsResearchOKBodyItems0 remainder_points false */

/* polymorph GetCharactersCharacterIDAgentsResearchOKBodyItems0 skill_type_id false */

/* polymorph GetCharactersCharacterIDAgentsResearchOKBodyItems0 started_at false */

// Validate validates this get characters character ID agents research o k body items0
func (o *GetCharactersCharacterIDAgentsResearchOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAgentID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validatePointsPerDay(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateRemainderPoints(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateSkillTypeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateStartedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCharactersCharacterIDAgentsResearchOKBodyItems0) validateAgentID(formats strfmt.Registry) error {

	if err := validate.Required("agent_id", "body", o.AgentID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDAgentsResearchOKBodyItems0) validatePointsPerDay(formats strfmt.Registry) error {

	if err := validate.Required("points_per_day", "body", o.PointsPerDay); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDAgentsResearchOKBodyItems0) validateRemainderPoints(formats strfmt.Registry) error {

	if err := validate.Required("remainder_points", "body", o.RemainderPoints); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDAgentsResearchOKBodyItems0) validateSkillTypeID(formats strfmt.Registry) error {

	if err := validate.Required("skill_type_id", "body", o.SkillTypeID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDAgentsResearchOKBodyItems0) validateStartedAt(formats strfmt.Registry) error {

	if err := validate.Required("started_at", "body", o.StartedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("started_at", "body", "date-time", o.StartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDAgentsResearchOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDAgentsResearchOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDAgentsResearchOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
