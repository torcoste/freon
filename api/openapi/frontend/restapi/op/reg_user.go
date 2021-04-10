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

// RegUserHandlerFunc turns a function with the right signature into a reg user handler
type RegUserHandlerFunc func(RegUserParams, *app.UserSession) RegUserResponder

// Handle executing the request and returning a response
func (fn RegUserHandlerFunc) Handle(params RegUserParams, principal *app.UserSession) RegUserResponder {
	return fn(params, principal)
}

// RegUserHandler interface for that can handle valid reg user params
type RegUserHandler interface {
	Handle(RegUserParams, *app.UserSession) RegUserResponder
}

// NewRegUser creates a new http.Handler for the reg user operation
func NewRegUser(ctx *middleware.Context, handler RegUserHandler) *RegUser {
	return &RegUser{Context: ctx, Handler: handler}
}

/*RegUser swagger:route POST /user/register regUser

registration new user

*/
type RegUser struct {
	Context *middleware.Context
	Handler RegUserHandler
}

func (o *RegUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRegUserParams()

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

// RegUserBody reg user body
//
// swagger:model RegUserBody
type RegUserBody struct {

	// email
	// Required: true
	// Max Length: 255
	// Min Length: 1
	// Pattern: ^[\x21-\x7F]{1,64}@[\x21-\x3F\x41-\x7F]+$
	Email *string `json:"email"`

	// first name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	FirstName *string `json:"first_name"`

	// password
	// Required: true
	// Max Length: 100
	// Min Length: 8
	// Format: password
	Password *strfmt.Password `json:"password"`

	// repeat password
	// Required: true
	// Max Length: 100
	// Min Length: 8
	// Format: password
	RepeatPassword *strfmt.Password `json:"repeat_password"`

	// role
	// Required: true
	Role *string `json:"role"`

	// second name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	SecondName *string `json:"second_name"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (o *RegUserBody) UnmarshalJSON(data []byte) error {
	var props struct {

		// email
		// Required: true
		// Max Length: 255
		// Min Length: 1
		// Pattern: ^[\x21-\x7F]{1,64}@[\x21-\x3F\x41-\x7F]+$
		Email *string `json:"email"`

		// first name
		// Required: true
		// Max Length: 100
		// Min Length: 1
		FirstName *string `json:"first_name"`

		// password
		// Required: true
		// Max Length: 100
		// Min Length: 8
		// Format: password
		Password *strfmt.Password `json:"password"`

		// repeat password
		// Required: true
		// Max Length: 100
		// Min Length: 8
		// Format: password
		RepeatPassword *strfmt.Password `json:"repeat_password"`

		// role
		// Required: true
		Role *string `json:"role"`

		// second name
		// Required: true
		// Max Length: 100
		// Min Length: 1
		SecondName *string `json:"second_name"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	o.Email = props.Email
	o.FirstName = props.FirstName
	o.Password = props.Password
	o.RepeatPassword = props.RepeatPassword
	o.Role = props.Role
	o.SecondName = props.SecondName
	return nil
}

// Validate validates this reg user body
func (o *RegUserBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFirstName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRepeatPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSecondName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RegUserBody) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("args"+"."+"email", "body", o.Email); err != nil {
		return err
	}

	if err := validate.MinLength("args"+"."+"email", "body", string(*o.Email), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("args"+"."+"email", "body", string(*o.Email), 255); err != nil {
		return err
	}

	if err := validate.Pattern("args"+"."+"email", "body", string(*o.Email), `^[\x21-\x7F]{1,64}@[\x21-\x3F\x41-\x7F]+$`); err != nil {
		return err
	}

	return nil
}

func (o *RegUserBody) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("args"+"."+"first_name", "body", o.FirstName); err != nil {
		return err
	}

	if err := validate.MinLength("args"+"."+"first_name", "body", string(*o.FirstName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("args"+"."+"first_name", "body", string(*o.FirstName), 100); err != nil {
		return err
	}

	return nil
}

func (o *RegUserBody) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("args"+"."+"password", "body", o.Password); err != nil {
		return err
	}

	if err := validate.MinLength("args"+"."+"password", "body", string(*o.Password), 8); err != nil {
		return err
	}

	if err := validate.MaxLength("args"+"."+"password", "body", string(*o.Password), 100); err != nil {
		return err
	}

	if err := validate.FormatOf("args"+"."+"password", "body", "password", o.Password.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *RegUserBody) validateRepeatPassword(formats strfmt.Registry) error {

	if err := validate.Required("args"+"."+"repeat_password", "body", o.RepeatPassword); err != nil {
		return err
	}

	if err := validate.MinLength("args"+"."+"repeat_password", "body", string(*o.RepeatPassword), 8); err != nil {
		return err
	}

	if err := validate.MaxLength("args"+"."+"repeat_password", "body", string(*o.RepeatPassword), 100); err != nil {
		return err
	}

	if err := validate.FormatOf("args"+"."+"repeat_password", "body", "password", o.RepeatPassword.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *RegUserBody) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("args"+"."+"role", "body", o.Role); err != nil {
		return err
	}

	return nil
}

func (o *RegUserBody) validateSecondName(formats strfmt.Registry) error {

	if err := validate.Required("args"+"."+"second_name", "body", o.SecondName); err != nil {
		return err
	}

	if err := validate.MinLength("args"+"."+"second_name", "body", string(*o.SecondName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("args"+"."+"second_name", "body", string(*o.SecondName), 100); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RegUserBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RegUserBody) UnmarshalBinary(b []byte) error {
	var res RegUserBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
