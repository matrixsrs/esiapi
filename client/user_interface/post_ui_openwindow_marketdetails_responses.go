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

// PostUIOpenwindowMarketdetailsReader is a Reader for the PostUIOpenwindowMarketdetails structure.
type PostUIOpenwindowMarketdetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUIOpenwindowMarketdetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPostUIOpenwindowMarketdetailsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewPostUIOpenwindowMarketdetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostUIOpenwindowMarketdetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostUIOpenwindowMarketdetailsNoContent creates a PostUIOpenwindowMarketdetailsNoContent with default headers values
func NewPostUIOpenwindowMarketdetailsNoContent() *PostUIOpenwindowMarketdetailsNoContent {
	return &PostUIOpenwindowMarketdetailsNoContent{}
}

/*PostUIOpenwindowMarketdetailsNoContent handles this case with default header values.

Open window request received
*/
type PostUIOpenwindowMarketdetailsNoContent struct {
}

func (o *PostUIOpenwindowMarketdetailsNoContent) Error() string {
	return fmt.Sprintf("[POST /ui/openwindow/marketdetails/][%d] postUiOpenwindowMarketdetailsNoContent ", 204)
}

func (o *PostUIOpenwindowMarketdetailsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUIOpenwindowMarketdetailsForbidden creates a PostUIOpenwindowMarketdetailsForbidden with default headers values
func NewPostUIOpenwindowMarketdetailsForbidden() *PostUIOpenwindowMarketdetailsForbidden {
	return &PostUIOpenwindowMarketdetailsForbidden{}
}

/*PostUIOpenwindowMarketdetailsForbidden handles this case with default header values.

Forbidden
*/
type PostUIOpenwindowMarketdetailsForbidden struct {
	Payload *models.Forbidden
}

func (o *PostUIOpenwindowMarketdetailsForbidden) Error() string {
	return fmt.Sprintf("[POST /ui/openwindow/marketdetails/][%d] postUiOpenwindowMarketdetailsForbidden  %+v", 403, o.Payload)
}

func (o *PostUIOpenwindowMarketdetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIOpenwindowMarketdetailsInternalServerError creates a PostUIOpenwindowMarketdetailsInternalServerError with default headers values
func NewPostUIOpenwindowMarketdetailsInternalServerError() *PostUIOpenwindowMarketdetailsInternalServerError {
	return &PostUIOpenwindowMarketdetailsInternalServerError{}
}

/*PostUIOpenwindowMarketdetailsInternalServerError handles this case with default header values.

Internal server error
*/
type PostUIOpenwindowMarketdetailsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *PostUIOpenwindowMarketdetailsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /ui/openwindow/marketdetails/][%d] postUiOpenwindowMarketdetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUIOpenwindowMarketdetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
