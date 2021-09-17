// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/freonservice/freon/internal/app"
)

// ListTranslationFilesHandlerFunc turns a function with the right signature into a list translation files handler
type ListTranslationFilesHandlerFunc func(ListTranslationFilesParams, *app.UserSession) ListTranslationFilesResponder

// Handle executing the request and returning a response
func (fn ListTranslationFilesHandlerFunc) Handle(params ListTranslationFilesParams, principal *app.UserSession) ListTranslationFilesResponder {
	return fn(params, principal)
}

// ListTranslationFilesHandler interface for that can handle valid list translation files params
type ListTranslationFilesHandler interface {
	Handle(ListTranslationFilesParams, *app.UserSession) ListTranslationFilesResponder
}

// NewListTranslationFiles creates a new http.Handler for the list translation files operation
func NewListTranslationFiles(ctx *middleware.Context, handler ListTranslationFilesHandler) *ListTranslationFiles {
	return &ListTranslationFiles{Context: ctx, Handler: handler}
}

/* ListTranslationFiles swagger:route GET /translation/files listTranslationFiles

get full list of available translation files

*/
type ListTranslationFiles struct {
	Context *middleware.Context
	Handler ListTranslationFilesHandler
}

func (o *ListTranslationFiles) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListTranslationFilesParams()
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
