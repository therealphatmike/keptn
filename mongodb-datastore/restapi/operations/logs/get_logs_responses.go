// Code generated by go-swagger; DO NOT EDIT.

package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/keptn/keptn/mongodb-datastore/models"
)

// GetLogsOKCode is the HTTP code returned for type GetLogsOK
const GetLogsOKCode int = 200

/*GetLogsOK ok

swagger:response getLogsOK
*/
type GetLogsOK struct {

	/*
	  In: Body
	*/
	Payload *GetLogsOKBody `json:"body,omitempty"`
}

// NewGetLogsOK creates GetLogsOK with default headers values
func NewGetLogsOK() *GetLogsOK {

	return &GetLogsOK{}
}

// WithPayload adds the payload to the get logs o k response
func (o *GetLogsOK) WithPayload(payload *GetLogsOKBody) *GetLogsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get logs o k response
func (o *GetLogsOK) SetPayload(payload *GetLogsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLogsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetLogsDefault error

swagger:response getLogsDefault
*/
type GetLogsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetLogsDefault creates GetLogsDefault with default headers values
func NewGetLogsDefault(code int) *GetLogsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetLogsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get logs default response
func (o *GetLogsDefault) WithStatusCode(code int) *GetLogsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get logs default response
func (o *GetLogsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get logs default response
func (o *GetLogsDefault) WithPayload(payload *models.Error) *GetLogsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get logs default response
func (o *GetLogsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLogsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
