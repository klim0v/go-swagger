package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/casualjim/go-swagger/examples/generated/models"
	"github.com/casualjim/go-swagger/httpkit/middleware"
)

// UpdatePetHandlerFunc turns a function with the right signature into a update pet handler
type UpdatePetHandlerFunc func(UpdatePetParams, *models.User) error

func (fn UpdatePetHandlerFunc) Handle(params UpdatePetParams, principal *models.User) error {
	return fn(params, principal)
}

// UpdatePetHandler interface for that can handle valid update pet params
type UpdatePetHandler interface {
	Handle(UpdatePetParams, *models.User) error
}

// NewUpdatePet creates a new http.Handler for the update pet operation
func NewUpdatePet(ctx *middleware.Context, handler UpdatePetHandler) *UpdatePet {
	return &UpdatePet{Context: ctx, Handler: handler}
}

// UpdatePet
type UpdatePet struct {
	Context *middleware.Context
	Params  UpdatePetParams
	Handler UpdatePetHandler
}

func (o *UpdatePet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal *models.User
	if uprinc != nil {
		principal = uprinc.(*models.User) // it's ok this is really a models.User
	}

	if err := o.Context.BindValidRequest(r, route, &o.Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	err = o.Handler.Handle(o.Params, principal) // actually handle the request
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	o.Context.Respond(rw, r, route.Produces, route, nil)

}
