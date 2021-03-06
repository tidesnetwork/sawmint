// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	models "github.com/faddat/sawmint/models"
)

// GetTransactionsTransactionIDHandlerFunc turns a function with the right signature into a get transactions transaction ID handler
type GetTransactionsTransactionIDHandlerFunc func(GetTransactionsTransactionIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTransactionsTransactionIDHandlerFunc) Handle(params GetTransactionsTransactionIDParams) middleware.Responder {
	return fn(params)
}

// GetTransactionsTransactionIDHandler interface for that can handle valid get transactions transaction ID params
type GetTransactionsTransactionIDHandler interface {
	Handle(GetTransactionsTransactionIDParams) middleware.Responder
}

// NewGetTransactionsTransactionID creates a new http.Handler for the get transactions transaction ID operation
func NewGetTransactionsTransactionID(ctx *middleware.Context, handler GetTransactionsTransactionIDHandler) *GetTransactionsTransactionID {
	return &GetTransactionsTransactionID{Context: ctx, Handler: handler}
}

/*GetTransactionsTransactionID swagger:route GET /transactions/{transaction_id} getTransactionsTransactionId

Fetches a particular transaction

*/
type GetTransactionsTransactionID struct {
	Context *middleware.Context
	Handler GetTransactionsTransactionIDHandler
}

func (o *GetTransactionsTransactionID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetTransactionsTransactionIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetTransactionsTransactionIDOKBody get transactions transaction ID o k body
// swagger:model GetTransactionsTransactionIDOKBody
type GetTransactionsTransactionIDOKBody struct {

	// data
	Data *models.Block `json:"data,omitempty"`

	// head
	Head models.Head `json:"head,omitempty"`

	// link
	Link models.Link `json:"link,omitempty"`
}

// Validate validates this get transactions transaction ID o k body
func (o *GetTransactionsTransactionIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateHead(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLink(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTransactionsTransactionIDOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getTransactionsTransactionIdOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

func (o *GetTransactionsTransactionIDOKBody) validateHead(formats strfmt.Registry) error {

	if swag.IsZero(o.Head) { // not required
		return nil
	}

	if err := o.Head.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getTransactionsTransactionIdOK" + "." + "head")
		}
		return err
	}

	return nil
}

func (o *GetTransactionsTransactionIDOKBody) validateLink(formats strfmt.Registry) error {

	if swag.IsZero(o.Link) { // not required
		return nil
	}

	if err := o.Link.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getTransactionsTransactionIdOK" + "." + "link")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionsTransactionIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionsTransactionIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetTransactionsTransactionIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
