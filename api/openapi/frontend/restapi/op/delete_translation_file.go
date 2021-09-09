// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/freonservice/freon/internal/app"
)

// DeleteTranslationFileHandlerFunc turns a function with the right signature into a delete translation file handler
type DeleteTranslationFileHandlerFunc func(DeleteTranslationFileParams, *app.UserSession) DeleteTranslationFileResponder

// Handle executing the request and returning a response
func (fn DeleteTranslationFileHandlerFunc) Handle(params DeleteTranslationFileParams, principal *app.UserSession) DeleteTranslationFileResponder {
	return fn(params, principal)
}

// DeleteTranslationFileHandler interface for that can handle valid delete translation file params
type DeleteTranslationFileHandler interface {
	Handle(DeleteTranslationFileParams, *app.UserSession) DeleteTranslationFileResponder
}

// NewDeleteTranslationFile creates a new http.Handler for the delete translation file operation
func NewDeleteTranslationFile(ctx *middleware.Context, handler DeleteTranslationFileHandler) *DeleteTranslationFile {
	return &DeleteTranslationFile{Context: ctx, Handler: handler}
}

/*DeleteTranslationFile swagger:route DELETE /translation/file/{id} deleteTranslationFile

delete translation file by id

*/
type DeleteTranslationFile struct {
	Context *middleware.Context
	Handler DeleteTranslationFileHandler
}

func (o *DeleteTranslationFile) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteTranslationFileParams()

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
