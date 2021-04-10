// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/freonservice/freon/internal/app"
)

// ListCategoriesHandlerFunc turns a function with the right signature into a list categories handler
type ListCategoriesHandlerFunc func(ListCategoriesParams, *app.UserSession) ListCategoriesResponder

// Handle executing the request and returning a response
func (fn ListCategoriesHandlerFunc) Handle(params ListCategoriesParams, principal *app.UserSession) ListCategoriesResponder {
	return fn(params, principal)
}

// ListCategoriesHandler interface for that can handle valid list categories params
type ListCategoriesHandler interface {
	Handle(ListCategoriesParams, *app.UserSession) ListCategoriesResponder
}

// NewListCategories creates a new http.Handler for the list categories operation
func NewListCategories(ctx *middleware.Context, handler ListCategoriesHandler) *ListCategories {
	return &ListCategories{Context: ctx, Handler: handler}
}

/*ListCategories swagger:route GET /categories listCategories

get full list of available categories

*/
type ListCategories struct {
	Context *middleware.Context
	Handler ListCategoriesHandler
}

func (o *ListCategories) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListCategoriesParams()

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
