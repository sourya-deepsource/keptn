// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/keptn/keptn/shipyard-controller/models"
)

// GetTriggeredEventsOKCode is the HTTP code returned for type GetTriggeredEventsOK
const GetTriggeredEventsOKCode int = 200

/*GetTriggeredEventsOK Success

swagger:response getTriggeredEventsOK
*/
type GetTriggeredEventsOK struct {

	/*
	  In: Body
	*/
	Payload *models.Events `json:"body,omitempty"`
}

// NewGetTriggeredEventsOK creates GetTriggeredEventsOK with default headers values
func NewGetTriggeredEventsOK() *GetTriggeredEventsOK {

	return &GetTriggeredEventsOK{}
}

// WithPayload adds the payload to the get triggered events o k response
func (o *GetTriggeredEventsOK) WithPayload(payload *models.Events) *GetTriggeredEventsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get triggered events o k response
func (o *GetTriggeredEventsOK) SetPayload(payload *models.Events) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTriggeredEventsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetTriggeredEventsDefault Error

swagger:response getTriggeredEventsDefault
*/
type GetTriggeredEventsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTriggeredEventsDefault creates GetTriggeredEventsDefault with default headers values
func NewGetTriggeredEventsDefault(code int) *GetTriggeredEventsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetTriggeredEventsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get triggered events default response
func (o *GetTriggeredEventsDefault) WithStatusCode(code int) *GetTriggeredEventsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get triggered events default response
func (o *GetTriggeredEventsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get triggered events default response
func (o *GetTriggeredEventsDefault) WithPayload(payload *models.Error) *GetTriggeredEventsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get triggered events default response
func (o *GetTriggeredEventsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTriggeredEventsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
