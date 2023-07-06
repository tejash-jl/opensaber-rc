// Code generated by go-swagger; DO NOT EDIT.

package uploaded_files

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"bulk_issuance/swagger_gen/models"
)

// GetV1UploadsHandlerFunc turns a function with the right signature into a get v1 uploads handler
type GetV1UploadsHandlerFunc func(GetV1UploadsParams, *models.JWTClaimBody) middleware.Responder

// Handle executing the request and returning a response
func (fn GetV1UploadsHandlerFunc) Handle(params GetV1UploadsParams, principal *models.JWTClaimBody) middleware.Responder {
	return fn(params, principal)
}

// GetV1UploadsHandler interface for that can handle valid get v1 uploads params
type GetV1UploadsHandler interface {
	Handle(GetV1UploadsParams, *models.JWTClaimBody) middleware.Responder
}

// NewGetV1Uploads creates a new http.Handler for the get v1 uploads operation
func NewGetV1Uploads(ctx *middleware.Context, handler GetV1UploadsHandler) *GetV1Uploads {
	return &GetV1Uploads{Context: ctx, Handler: handler}
}

/*
	GetV1Uploads swagger:route GET /v1/uploads uploadedFiles getV1Uploads

get uploaded files
*/
type GetV1Uploads struct {
	Context *middleware.Context
	Handler GetV1UploadsHandler
}

func (o *GetV1Uploads) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetV1UploadsParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.JWTClaimBody
	if uprinc != nil {
		principal = uprinc.(*models.JWTClaimBody) // this is really a models.JWTClaimBody, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}