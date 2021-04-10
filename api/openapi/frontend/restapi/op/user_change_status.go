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

// UserChangeStatusHandlerFunc turns a function with the right signature into a user change status handler
type UserChangeStatusHandlerFunc func(UserChangeStatusParams, *app.UserSession) UserChangeStatusResponder

// Handle executing the request and returning a response
func (fn UserChangeStatusHandlerFunc) Handle(params UserChangeStatusParams, principal *app.UserSession) UserChangeStatusResponder {
	return fn(params, principal)
}

// UserChangeStatusHandler interface for that can handle valid user change status params
type UserChangeStatusHandler interface {
	Handle(UserChangeStatusParams, *app.UserSession) UserChangeStatusResponder
}

// NewUserChangeStatus creates a new http.Handler for the user change status operation
func NewUserChangeStatus(ctx *middleware.Context, handler UserChangeStatusHandler) *UserChangeStatus {
	return &UserChangeStatus{Context: ctx, Handler: handler}
}

/*UserChangeStatus swagger:route PUT /user/change-status userChangeStatus

user change status

*/
type UserChangeStatus struct {
	Context *middleware.Context
	Handler UserChangeStatusHandler
}

func (o *UserChangeStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUserChangeStatusParams()

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

// UserChangeStatusBody user change status body
//
// swagger:model UserChangeStatusBody
type UserChangeStatusBody struct {

	// status
	// Required: true
	Status *int64 `json:"status"`

	// user id
	// Required: true
	UserID *int64 `json:"user_id"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (o *UserChangeStatusBody) UnmarshalJSON(data []byte) error {
	var props struct {

		// status
		// Required: true
		Status *int64 `json:"status"`

		// user id
		// Required: true
		UserID *int64 `json:"user_id"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	o.Status = props.Status
	o.UserID = props.UserID
	return nil
}

// Validate validates this user change status body
func (o *UserChangeStatusBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UserChangeStatusBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("args"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

func (o *UserChangeStatusBody) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("args"+"."+"user_id", "body", o.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UserChangeStatusBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UserChangeStatusBody) UnmarshalBinary(b []byte) error {
	var res UserChangeStatusBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
