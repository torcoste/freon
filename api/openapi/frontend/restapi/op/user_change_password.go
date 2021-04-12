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

// UserChangePasswordHandlerFunc turns a function with the right signature into a user change password handler
type UserChangePasswordHandlerFunc func(UserChangePasswordParams, *app.UserSession) UserChangePasswordResponder

// Handle executing the request and returning a response
func (fn UserChangePasswordHandlerFunc) Handle(params UserChangePasswordParams, principal *app.UserSession) UserChangePasswordResponder {
	return fn(params, principal)
}

// UserChangePasswordHandler interface for that can handle valid user change password params
type UserChangePasswordHandler interface {
	Handle(UserChangePasswordParams, *app.UserSession) UserChangePasswordResponder
}

// NewUserChangePassword creates a new http.Handler for the user change password operation
func NewUserChangePassword(ctx *middleware.Context, handler UserChangePasswordHandler) *UserChangePassword {
	return &UserChangePassword{Context: ctx, Handler: handler}
}

/*UserChangePassword swagger:route PUT /user/change-password userChangePassword

user change password

*/
type UserChangePassword struct {
	Context *middleware.Context
	Handler UserChangePasswordHandler
}

func (o *UserChangePassword) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUserChangePasswordParams()

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

// UserChangePasswordBody user change password body
//
// swagger:model UserChangePasswordBody
type UserChangePasswordBody struct {

	// new password
	// Required: true
	// Max Length: 100
	// Min Length: 8
	NewPassword *string `json:"new_password"`

	// old password
	// Required: true
	// Max Length: 100
	// Min Length: 8
	OldPassword *string `json:"old_password"`

	// repeat password
	// Required: true
	// Max Length: 100
	// Min Length: 8
	RepeatPassword *string `json:"repeat_password"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (o *UserChangePasswordBody) UnmarshalJSON(data []byte) error {
	var props struct {

		// new password
		// Required: true
		// Max Length: 100
		// Min Length: 8
		NewPassword *string `json:"new_password"`

		// old password
		// Required: true
		// Max Length: 100
		// Min Length: 8
		OldPassword *string `json:"old_password"`

		// repeat password
		// Required: true
		// Max Length: 100
		// Min Length: 8
		RepeatPassword *string `json:"repeat_password"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	o.NewPassword = props.NewPassword
	o.OldPassword = props.OldPassword
	o.RepeatPassword = props.RepeatPassword
	return nil
}

// Validate validates this user change password body
func (o *UserChangePasswordBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNewPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOldPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRepeatPassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UserChangePasswordBody) validateNewPassword(formats strfmt.Registry) error {

	if err := validate.Required("args"+"."+"new_password", "body", o.NewPassword); err != nil {
		return err
	}

	if err := validate.MinLength("args"+"."+"new_password", "body", string(*o.NewPassword), 8); err != nil {
		return err
	}

	if err := validate.MaxLength("args"+"."+"new_password", "body", string(*o.NewPassword), 100); err != nil {
		return err
	}

	return nil
}

func (o *UserChangePasswordBody) validateOldPassword(formats strfmt.Registry) error {

	if err := validate.Required("args"+"."+"old_password", "body", o.OldPassword); err != nil {
		return err
	}

	if err := validate.MinLength("args"+"."+"old_password", "body", string(*o.OldPassword), 8); err != nil {
		return err
	}

	if err := validate.MaxLength("args"+"."+"old_password", "body", string(*o.OldPassword), 100); err != nil {
		return err
	}

	return nil
}

func (o *UserChangePasswordBody) validateRepeatPassword(formats strfmt.Registry) error {

	if err := validate.Required("args"+"."+"repeat_password", "body", o.RepeatPassword); err != nil {
		return err
	}

	if err := validate.MinLength("args"+"."+"repeat_password", "body", string(*o.RepeatPassword), 8); err != nil {
		return err
	}

	if err := validate.MaxLength("args"+"."+"repeat_password", "body", string(*o.RepeatPassword), 100); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UserChangePasswordBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UserChangePasswordBody) UnmarshalBinary(b []byte) error {
	var res UserChangePasswordBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
