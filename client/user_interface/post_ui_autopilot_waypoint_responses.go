// Code generated by go-swagger; DO NOT EDIT.

package user_interface

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/models"
)

// PostUIAutopilotWaypointReader is a Reader for the PostUIAutopilotWaypoint structure.
type PostUIAutopilotWaypointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUIAutopilotWaypointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPostUIAutopilotWaypointNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewPostUIAutopilotWaypointForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostUIAutopilotWaypointInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostUIAutopilotWaypointNoContent creates a PostUIAutopilotWaypointNoContent with default headers values
func NewPostUIAutopilotWaypointNoContent() *PostUIAutopilotWaypointNoContent {
	return &PostUIAutopilotWaypointNoContent{}
}

/*PostUIAutopilotWaypointNoContent handles this case with default header values.

Open window request received
*/
type PostUIAutopilotWaypointNoContent struct {
}

func (o *PostUIAutopilotWaypointNoContent) Error() string {
	return fmt.Sprintf("[POST /ui/autopilot/waypoint/][%d] postUiAutopilotWaypointNoContent ", 204)
}

func (o *PostUIAutopilotWaypointNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUIAutopilotWaypointForbidden creates a PostUIAutopilotWaypointForbidden with default headers values
func NewPostUIAutopilotWaypointForbidden() *PostUIAutopilotWaypointForbidden {
	return &PostUIAutopilotWaypointForbidden{}
}

/*PostUIAutopilotWaypointForbidden handles this case with default header values.

Forbidden
*/
type PostUIAutopilotWaypointForbidden struct {
	Payload *models.Forbidden
}

func (o *PostUIAutopilotWaypointForbidden) Error() string {
	return fmt.Sprintf("[POST /ui/autopilot/waypoint/][%d] postUiAutopilotWaypointForbidden  %+v", 403, o.Payload)
}

func (o *PostUIAutopilotWaypointForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIAutopilotWaypointInternalServerError creates a PostUIAutopilotWaypointInternalServerError with default headers values
func NewPostUIAutopilotWaypointInternalServerError() *PostUIAutopilotWaypointInternalServerError {
	return &PostUIAutopilotWaypointInternalServerError{}
}

/*PostUIAutopilotWaypointInternalServerError handles this case with default header values.

Internal server error
*/
type PostUIAutopilotWaypointInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *PostUIAutopilotWaypointInternalServerError) Error() string {
	return fmt.Sprintf("[POST /ui/autopilot/waypoint/][%d] postUiAutopilotWaypointInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUIAutopilotWaypointInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
