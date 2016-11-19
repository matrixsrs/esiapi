package skills

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetCharactersCharacterIDSkillqueueParams creates a new GetCharactersCharacterIDSkillqueueParams object
// with the default values initialized.
func NewGetCharactersCharacterIDSkillqueueParams() *GetCharactersCharacterIDSkillqueueParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDSkillqueueParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDSkillqueueParamsWithTimeout creates a new GetCharactersCharacterIDSkillqueueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCharactersCharacterIDSkillqueueParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDSkillqueueParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDSkillqueueParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDSkillqueueParamsWithContext creates a new GetCharactersCharacterIDSkillqueueParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCharactersCharacterIDSkillqueueParamsWithContext(ctx context.Context) *GetCharactersCharacterIDSkillqueueParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDSkillqueueParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

/*GetCharactersCharacterIDSkillqueueParams contains all the parameters to send to the API endpoint
for the get characters character id skillqueue operation typically these are written to a http.Request
*/
type GetCharactersCharacterIDSkillqueueParams struct {

	/*CharacterID
	  Character id of the target character

	*/
	CharacterID int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get characters character id skillqueue params
func (o *GetCharactersCharacterIDSkillqueueParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDSkillqueueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id skillqueue params
func (o *GetCharactersCharacterIDSkillqueueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id skillqueue params
func (o *GetCharactersCharacterIDSkillqueueParams) WithContext(ctx context.Context) *GetCharactersCharacterIDSkillqueueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id skillqueue params
func (o *GetCharactersCharacterIDSkillqueueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithCharacterID adds the characterID to the get characters character id skillqueue params
func (o *GetCharactersCharacterIDSkillqueueParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDSkillqueueParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id skillqueue params
func (o *GetCharactersCharacterIDSkillqueueParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id skillqueue params
func (o *GetCharactersCharacterIDSkillqueueParams) WithDatasource(datasource *string) *GetCharactersCharacterIDSkillqueueParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id skillqueue params
func (o *GetCharactersCharacterIDSkillqueueParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDSkillqueueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param character_id
	if err := r.SetPathParam("character_id", swag.FormatInt32(o.CharacterID)); err != nil {
		return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}