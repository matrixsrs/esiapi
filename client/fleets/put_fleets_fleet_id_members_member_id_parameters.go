// Code generated by go-swagger; DO NOT EDIT.

package fleets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPutFleetsFleetIDMembersMemberIDParams creates a new PutFleetsFleetIDMembersMemberIDParams object
// with the default values initialized.
func NewPutFleetsFleetIDMembersMemberIDParams() *PutFleetsFleetIDMembersMemberIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PutFleetsFleetIDMembersMemberIDParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPutFleetsFleetIDMembersMemberIDParamsWithTimeout creates a new PutFleetsFleetIDMembersMemberIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutFleetsFleetIDMembersMemberIDParamsWithTimeout(timeout time.Duration) *PutFleetsFleetIDMembersMemberIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PutFleetsFleetIDMembersMemberIDParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewPutFleetsFleetIDMembersMemberIDParamsWithContext creates a new PutFleetsFleetIDMembersMemberIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutFleetsFleetIDMembersMemberIDParamsWithContext(ctx context.Context) *PutFleetsFleetIDMembersMemberIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PutFleetsFleetIDMembersMemberIDParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewPutFleetsFleetIDMembersMemberIDParamsWithHTTPClient creates a new PutFleetsFleetIDMembersMemberIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutFleetsFleetIDMembersMemberIDParamsWithHTTPClient(client *http.Client) *PutFleetsFleetIDMembersMemberIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PutFleetsFleetIDMembersMemberIDParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*PutFleetsFleetIDMembersMemberIDParams contains all the parameters to send to the API endpoint
for the put fleets fleet id members member id operation typically these are written to a http.Request
*/
type PutFleetsFleetIDMembersMemberIDParams struct {

	/*XUserAgent
	  Client identifier, takes precedence over User-Agent

	*/
	XUserAgent *string
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*FleetID
	  ID for a fleet

	*/
	FleetID int64
	/*MemberID
	  The character ID of a member in this fleet

	*/
	MemberID int32
	/*Movement
	  Details of the invitation

	*/
	Movement PutFleetsFleetIDMembersMemberIDBody
	/*Token
	  Access token to use if unable to set a header

	*/
	Token *string
	/*UserAgent
	  Client identifier, takes precedence over headers

	*/
	UserAgent *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put fleets fleet id members member id params
func (o *PutFleetsFleetIDMembersMemberIDParams) WithTimeout(timeout time.Duration) *PutFleetsFleetIDMembersMemberIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put fleets fleet id members member id params
func (o *PutFleetsFleetIDMembersMemberIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put fleets fleet id members member id params
func (o *PutFleetsFleetIDMembersMemberIDParams) WithContext(ctx context.Context) *PutFleetsFleetIDMembersMemberIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put fleets fleet id members member id params
func (o *PutFleetsFleetIDMembersMemberIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put fleets fleet id members member id params
func (o *PutFleetsFleetIDMembersMemberIDParams) WithHTTPClient(client *http.Client) *PutFleetsFleetIDMembersMemberIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put fleets fleet id members member id params
func (o *PutFleetsFleetIDMembersMemberIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the put fleets fleet id members member id params
func (o *PutFleetsFleetIDMembersMemberIDParams) WithXUserAgent(xUserAgent *string) *PutFleetsFleetIDMembersMemberIDParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the put fleets fleet id members member id params
func (o *PutFleetsFleetIDMembersMemberIDParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithDatasource adds the datasource to the put fleets fleet id members member id params
func (o *PutFleetsFleetIDMembersMemberIDParams) WithDatasource(datasource *string) *PutFleetsFleetIDMembersMemberIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the put fleets fleet id members member id params
func (o *PutFleetsFleetIDMembersMemberIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithFleetID adds the fleetID to the put fleets fleet id members member id params
func (o *PutFleetsFleetIDMembersMemberIDParams) WithFleetID(fleetID int64) *PutFleetsFleetIDMembersMemberIDParams {
	o.SetFleetID(fleetID)
	return o
}

// SetFleetID adds the fleetId to the put fleets fleet id members member id params
func (o *PutFleetsFleetIDMembersMemberIDParams) SetFleetID(fleetID int64) {
	o.FleetID = fleetID
}

// WithMemberID adds the memberID to the put fleets fleet id members member id params
func (o *PutFleetsFleetIDMembersMemberIDParams) WithMemberID(memberID int32) *PutFleetsFleetIDMembersMemberIDParams {
	o.SetMemberID(memberID)
	return o
}

// SetMemberID adds the memberId to the put fleets fleet id members member id params
func (o *PutFleetsFleetIDMembersMemberIDParams) SetMemberID(memberID int32) {
	o.MemberID = memberID
}

// WithMovement adds the movement to the put fleets fleet id members member id params
func (o *PutFleetsFleetIDMembersMemberIDParams) WithMovement(movement PutFleetsFleetIDMembersMemberIDBody) *PutFleetsFleetIDMembersMemberIDParams {
	o.SetMovement(movement)
	return o
}

// SetMovement adds the movement to the put fleets fleet id members member id params
func (o *PutFleetsFleetIDMembersMemberIDParams) SetMovement(movement PutFleetsFleetIDMembersMemberIDBody) {
	o.Movement = movement
}

// WithToken adds the token to the put fleets fleet id members member id params
func (o *PutFleetsFleetIDMembersMemberIDParams) WithToken(token *string) *PutFleetsFleetIDMembersMemberIDParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the put fleets fleet id members member id params
func (o *PutFleetsFleetIDMembersMemberIDParams) SetToken(token *string) {
	o.Token = token
}

// WithUserAgent adds the userAgent to the put fleets fleet id members member id params
func (o *PutFleetsFleetIDMembersMemberIDParams) WithUserAgent(userAgent *string) *PutFleetsFleetIDMembersMemberIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the put fleets fleet id members member id params
func (o *PutFleetsFleetIDMembersMemberIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *PutFleetsFleetIDMembersMemberIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XUserAgent != nil {

		// header param X-User-Agent
		if err := r.SetHeaderParam("X-User-Agent", *o.XUserAgent); err != nil {
			return err
		}

	}

	if o.Datasource != nil {

		// query param datasource
		var qrDatasource string
		if o.Datasource != nil {
			qrDatasource = *o.Datasource
		}
		qDatasource := qrDatasource
		if qDatasource != "" {
			if err := r.SetQueryParam("datasource", qDatasource); err != nil {
				return err
			}
		}

	}

	// path param fleet_id
	if err := r.SetPathParam("fleet_id", swag.FormatInt64(o.FleetID)); err != nil {
		return err
	}

	// path param member_id
	if err := r.SetPathParam("member_id", swag.FormatInt32(o.MemberID)); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.Movement); err != nil {
		return err
	}

	if o.Token != nil {

		// query param token
		var qrToken string
		if o.Token != nil {
			qrToken = *o.Token
		}
		qToken := qrToken
		if qToken != "" {
			if err := r.SetQueryParam("token", qToken); err != nil {
				return err
			}
		}

	}

	if o.UserAgent != nil {

		// query param user_agent
		var qrUserAgent string
		if o.UserAgent != nil {
			qrUserAgent = *o.UserAgent
		}
		qUserAgent := qrUserAgent
		if qUserAgent != "" {
			if err := r.SetQueryParam("user_agent", qUserAgent); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
