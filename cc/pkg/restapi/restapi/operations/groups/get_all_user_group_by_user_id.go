// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAllUserGroupByUserIDHandlerFunc turns a function with the right signature into a get all user group by user Id handler
type GetAllUserGroupByUserIDHandlerFunc func(GetAllUserGroupByUserIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAllUserGroupByUserIDHandlerFunc) Handle(params GetAllUserGroupByUserIDParams) middleware.Responder {
	return fn(params)
}

// GetAllUserGroupByUserIDHandler interface for that can handle valid get all user group by user Id params
type GetAllUserGroupByUserIDHandler interface {
	Handle(GetAllUserGroupByUserIDParams) middleware.Responder
}

// NewGetAllUserGroupByUserID creates a new http.Handler for the get all user group by user Id operation
func NewGetAllUserGroupByUserID(ctx *middleware.Context, handler GetAllUserGroupByUserIDHandler) *GetAllUserGroupByUserID {
	return &GetAllUserGroupByUserID{Context: ctx, Handler: handler}
}

/*GetAllUserGroupByUserID swagger:route GET /cc/v1/groups/userGroup/user/{userId} Groups getAllUserGroupByUserId

Returns a userGroup.

Optional extended description in Markdown.

*/
type GetAllUserGroupByUserID struct {
	Context *middleware.Context
	Handler GetAllUserGroupByUserIDHandler
}

func (o *GetAllUserGroupByUserID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAllUserGroupByUserIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
