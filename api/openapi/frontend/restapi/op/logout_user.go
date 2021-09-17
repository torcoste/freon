// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/freonservice/freon/internal/app"
)

// LogoutUserHandlerFunc turns a function with the right signature into a logout user handler
type LogoutUserHandlerFunc func(LogoutUserParams, *app.UserSession) LogoutUserResponder

// Handle executing the request and returning a response
func (fn LogoutUserHandlerFunc) Handle(params LogoutUserParams, principal *app.UserSession) LogoutUserResponder {
	return fn(params, principal)
}

// LogoutUserHandler interface for that can handle valid logout user params
type LogoutUserHandler interface {
	Handle(LogoutUserParams, *app.UserSession) LogoutUserResponder
}

// NewLogoutUser creates a new http.Handler for the logout user operation
func NewLogoutUser(ctx *middleware.Context, handler LogoutUserHandler) *LogoutUser {
	return &LogoutUser{Context: ctx, Handler: handler}
}

/* LogoutUser swagger:route POST /logout logoutUser

logout user, remove jwt session

*/
type LogoutUser struct {
	Context *middleware.Context
	Handler LogoutUserHandler
}

func (o *LogoutUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewLogoutUserParams()
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
