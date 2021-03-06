// Code generated by go-swagger; DO NOT EDIT.

package auths

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "mlss-controlcenter-go/pkg/models"
)

// CheckNamespaceOKCode is the HTTP code returned for type CheckNamespaceOK
const CheckNamespaceOKCode int = 200

/*CheckNamespaceOK auth by namespace.

swagger:response checkNamespaceOK
*/
type CheckNamespaceOK struct {

	/*
	  In: Body
	*/
	Payload *models.Result `json:"body,omitempty"`
}

// NewCheckNamespaceOK creates CheckNamespaceOK with default headers values
func NewCheckNamespaceOK() *CheckNamespaceOK {

	return &CheckNamespaceOK{}
}

// WithPayload adds the payload to the check namespace o k response
func (o *CheckNamespaceOK) WithPayload(payload *models.Result) *CheckNamespaceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check namespace o k response
func (o *CheckNamespaceOK) SetPayload(payload *models.Result) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckNamespaceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckNamespaceUnauthorizedCode is the HTTP code returned for type CheckNamespaceUnauthorized
const CheckNamespaceUnauthorizedCode int = 401

/*CheckNamespaceUnauthorized Unauthorized

swagger:response checkNamespaceUnauthorized
*/
type CheckNamespaceUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCheckNamespaceUnauthorized creates CheckNamespaceUnauthorized with default headers values
func NewCheckNamespaceUnauthorized() *CheckNamespaceUnauthorized {

	return &CheckNamespaceUnauthorized{}
}

// WithPayload adds the payload to the check namespace unauthorized response
func (o *CheckNamespaceUnauthorized) WithPayload(payload *models.Error) *CheckNamespaceUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check namespace unauthorized response
func (o *CheckNamespaceUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckNamespaceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckNamespaceNotFoundCode is the HTTP code returned for type CheckNamespaceNotFound
const CheckNamespaceNotFoundCode int = 404

/*CheckNamespaceNotFound url to add namespace not found.

swagger:response checkNamespaceNotFound
*/
type CheckNamespaceNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCheckNamespaceNotFound creates CheckNamespaceNotFound with default headers values
func NewCheckNamespaceNotFound() *CheckNamespaceNotFound {

	return &CheckNamespaceNotFound{}
}

// WithPayload adds the payload to the check namespace not found response
func (o *CheckNamespaceNotFound) WithPayload(payload *models.Error) *CheckNamespaceNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check namespace not found response
func (o *CheckNamespaceNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckNamespaceNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
