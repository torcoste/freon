// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/freonservice/freon/internal/app"
)

// CreateLocalizationHandlerFunc turns a function with the right signature into a create localization handler
type CreateLocalizationHandlerFunc func(CreateLocalizationParams, *app.UserSession) CreateLocalizationResponder

// Handle executing the request and returning a response
func (fn CreateLocalizationHandlerFunc) Handle(params CreateLocalizationParams, principal *app.UserSession) CreateLocalizationResponder {
	return fn(params, principal)
}

// CreateLocalizationHandler interface for that can handle valid create localization params
type CreateLocalizationHandler interface {
	Handle(CreateLocalizationParams, *app.UserSession) CreateLocalizationResponder
}

// NewCreateLocalization creates a new http.Handler for the create localization operation
func NewCreateLocalization(ctx *middleware.Context, handler CreateLocalizationHandler) *CreateLocalization {
	return &CreateLocalization{Context: ctx, Handler: handler}
}

/*CreateLocalization swagger:route POST /localization createLocalization

create new localization type

*/
type CreateLocalization struct {
	Context *middleware.Context
	Handler CreateLocalizationHandler
}

func (o *CreateLocalization) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateLocalizationParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *app.UserSession
	if uprinc != nil {
		principal = uprinc.(*app.UserSession) // this is really a app.UserSession, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// CreateLocalizationBody create localization body
//
// swagger:model CreateLocalizationBody
type CreateLocalizationBody struct {

	// icon
	Icon string `json:"icon,omitempty"`

	// lang name
	// Required: true
	// Max Length: 255
	// Min Length: 1
	LangName *string `json:"lang_name"`

	// locale
	// Required: true
	// Max Length: 10
	// Min Length: 2
	Locale *string `json:"locale"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (o *CreateLocalizationBody) UnmarshalJSON(data []byte) error {
	var props struct {

		// icon
		Icon string `json:"icon,omitempty"`

		// lang name
		// Required: true
		// Max Length: 255
		// Min Length: 1
		LangName *string `json:"lang_name"`

		// locale
		// Required: true
		// Max Length: 10
		// Min Length: 2
		Locale *string `json:"locale"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	o.Icon = props.Icon
	o.LangName = props.LangName
	o.Locale = props.Locale
	return nil
}

// Validate validates this create localization body
func (o *CreateLocalizationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLangName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLocale(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateLocalizationBody) validateLangName(formats strfmt.Registry) error {

	if err := validate.Required("args"+"."+"lang_name", "body", o.LangName); err != nil {
		return err
	}

	if err := validate.MinLength("args"+"."+"lang_name", "body", string(*o.LangName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("args"+"."+"lang_name", "body", string(*o.LangName), 255); err != nil {
		return err
	}

	return nil
}

func (o *CreateLocalizationBody) validateLocale(formats strfmt.Registry) error {

	if err := validate.Required("args"+"."+"locale", "body", o.Locale); err != nil {
		return err
	}

	if err := validate.MinLength("args"+"."+"locale", "body", string(*o.Locale), 2); err != nil {
		return err
	}

	if err := validate.MaxLength("args"+"."+"locale", "body", string(*o.Locale), 10); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateLocalizationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateLocalizationBody) UnmarshalBinary(b []byte) error {
	var res CreateLocalizationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
