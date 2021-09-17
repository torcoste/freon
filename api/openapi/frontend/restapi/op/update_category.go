// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/freonservice/freon/internal/app"
)

// UpdateCategoryHandlerFunc turns a function with the right signature into a update category handler
type UpdateCategoryHandlerFunc func(UpdateCategoryParams, *app.UserSession) UpdateCategoryResponder

// Handle executing the request and returning a response
func (fn UpdateCategoryHandlerFunc) Handle(params UpdateCategoryParams, principal *app.UserSession) UpdateCategoryResponder {
	return fn(params, principal)
}

// UpdateCategoryHandler interface for that can handle valid update category params
type UpdateCategoryHandler interface {
	Handle(UpdateCategoryParams, *app.UserSession) UpdateCategoryResponder
}

// NewUpdateCategory creates a new http.Handler for the update category operation
func NewUpdateCategory(ctx *middleware.Context, handler UpdateCategoryHandler) *UpdateCategory {
	return &UpdateCategory{Context: ctx, Handler: handler}
}

/* UpdateCategory swagger:route PUT /category/{id} updateCategory

update category

*/
type UpdateCategory struct {
	Context *middleware.Context
	Handler UpdateCategoryHandler
}

func (o *UpdateCategory) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateCategoryParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
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

// UpdateCategoryBody update category body
//
// swagger:model UpdateCategoryBody
type UpdateCategoryBody struct {

	// name
	// Required: true
	// Max Length: 255
	// Min Length: 1
	Name *string `json:"name"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (o *UpdateCategoryBody) UnmarshalJSON(data []byte) error {
	var props struct {

		// name
		// Required: true
		// Max Length: 255
		// Min Length: 1
		Name *string `json:"name"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	o.Name = props.Name
	return nil
}

// Validate validates this update category body
func (o *UpdateCategoryBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateCategoryBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("args"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	if err := validate.MinLength("args"+"."+"name", "body", *o.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("args"+"."+"name", "body", *o.Name, 255); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update category body based on context it is used
func (o *UpdateCategoryBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateCategoryBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateCategoryBody) UnmarshalBinary(b []byte) error {
	var res UpdateCategoryBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
