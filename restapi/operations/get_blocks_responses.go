// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/faddat/sawmint/models"
)

// GetBlocksOKCode is the HTTP code returned for type GetBlocksOK
const GetBlocksOKCode int = 200

/*GetBlocksOK Successfully retrieved blocks

swagger:response getBlocksOK
*/
type GetBlocksOK struct {

	/*
	  In: Body
	*/
	Payload *GetBlocksOKBody `json:"body,omitempty"`
}

// NewGetBlocksOK creates GetBlocksOK with default headers values
func NewGetBlocksOK() *GetBlocksOK {

	return &GetBlocksOK{}
}

// WithPayload adds the payload to the get blocks o k response
func (o *GetBlocksOK) WithPayload(payload *GetBlocksOKBody) *GetBlocksOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get blocks o k response
func (o *GetBlocksOK) SetPayload(payload *GetBlocksOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBlocksOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBlocksBadRequestCode is the HTTP code returned for type GetBlocksBadRequest
const GetBlocksBadRequestCode int = 400

/*GetBlocksBadRequest Request was malformed

swagger:response getBlocksBadRequest
*/
type GetBlocksBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetBlocksBadRequest creates GetBlocksBadRequest with default headers values
func NewGetBlocksBadRequest() *GetBlocksBadRequest {

	return &GetBlocksBadRequest{}
}

// WithPayload adds the payload to the get blocks bad request response
func (o *GetBlocksBadRequest) WithPayload(payload *models.Error) *GetBlocksBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get blocks bad request response
func (o *GetBlocksBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBlocksBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBlocksInternalServerErrorCode is the HTTP code returned for type GetBlocksInternalServerError
const GetBlocksInternalServerErrorCode int = 500

/*GetBlocksInternalServerError Something went wrong within the validator

swagger:response getBlocksInternalServerError
*/
type GetBlocksInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetBlocksInternalServerError creates GetBlocksInternalServerError with default headers values
func NewGetBlocksInternalServerError() *GetBlocksInternalServerError {

	return &GetBlocksInternalServerError{}
}

// WithPayload adds the payload to the get blocks internal server error response
func (o *GetBlocksInternalServerError) WithPayload(payload *models.Error) *GetBlocksInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get blocks internal server error response
func (o *GetBlocksInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBlocksInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBlocksServiceUnavailableCode is the HTTP code returned for type GetBlocksServiceUnavailable
const GetBlocksServiceUnavailableCode int = 503

/*GetBlocksServiceUnavailable API is unable to reach the validator

swagger:response getBlocksServiceUnavailable
*/
type GetBlocksServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetBlocksServiceUnavailable creates GetBlocksServiceUnavailable with default headers values
func NewGetBlocksServiceUnavailable() *GetBlocksServiceUnavailable {

	return &GetBlocksServiceUnavailable{}
}

// WithPayload adds the payload to the get blocks service unavailable response
func (o *GetBlocksServiceUnavailable) WithPayload(payload *models.Error) *GetBlocksServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get blocks service unavailable response
func (o *GetBlocksServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBlocksServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}