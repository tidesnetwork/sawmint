// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	"strconv"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	models "github.com/faddat/sawmint/models"
)

// GetBatchesHandlerFunc turns a function with the right signature into a get batches handler
type GetBatchesHandlerFunc func(GetBatchesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBatchesHandlerFunc) Handle(params GetBatchesParams) middleware.Responder {
	return fn(params)
}

// GetBatchesHandler interface for that can handle valid get batches params
type GetBatchesHandler interface {
	Handle(GetBatchesParams) middleware.Responder
}

// NewGetBatches creates a new http.Handler for the get batches operation
func NewGetBatches(ctx *middleware.Context, handler GetBatchesHandler) *GetBatches {
	return &GetBatches{Context: ctx, Handler: handler}
}

/*GetBatches swagger:route GET /batches getBatches

Fetches a list of batches

Fetches a paginated list of batches from the validator.


*/
type GetBatches struct {
	Context *middleware.Context
	Handler GetBatchesHandler
}

func (o *GetBatches) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetBatchesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetBatchesOKBody get batches o k body
// swagger:model GetBatchesOKBody
type GetBatchesOKBody struct {

	// data
	Data []*models.Batch `json:"data"`

	// head
	Head models.Head `json:"head,omitempty"`

	// link
	Link models.Link `json:"link,omitempty"`

	// paging
	Paging *models.Paging `json:"paging,omitempty"`
}

// Validate validates this get batches o k body
func (o *GetBatchesOKBody) Validate(formats strfmt.Registry) error {
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

	if err := o.validatePaging(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBatchesOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getBatchesOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetBatchesOKBody) validateHead(formats strfmt.Registry) error {

	if swag.IsZero(o.Head) { // not required
		return nil
	}

	if err := o.Head.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getBatchesOK" + "." + "head")
		}
		return err
	}

	return nil
}

func (o *GetBatchesOKBody) validateLink(formats strfmt.Registry) error {

	if swag.IsZero(o.Link) { // not required
		return nil
	}

	if err := o.Link.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getBatchesOK" + "." + "link")
		}
		return err
	}

	return nil
}

func (o *GetBatchesOKBody) validatePaging(formats strfmt.Registry) error {

	if swag.IsZero(o.Paging) { // not required
		return nil
	}

	if o.Paging != nil {
		if err := o.Paging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getBatchesOK" + "." + "paging")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetBatchesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetBatchesOKBody) UnmarshalBinary(b []byte) error {
	var res GetBatchesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
