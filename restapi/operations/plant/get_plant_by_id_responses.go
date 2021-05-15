// Code generated by go-swagger; DO NOT EDIT.

package plant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kaatinga/plantbook/models"
)

// GetPlantByIDOKCode is the HTTP code returned for type GetPlantByIDOK
const GetPlantByIDOKCode int = 200

/*GetPlantByIDOK successful operation

swagger:response getPlantByIdOK
*/
type GetPlantByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Plant `json:"body,omitempty"`
}

// NewGetPlantByIDOK creates GetPlantByIDOK with default headers values
func NewGetPlantByIDOK() *GetPlantByIDOK {

	return &GetPlantByIDOK{}
}

// WithPayload adds the payload to the get plant by Id o k response
func (o *GetPlantByIDOK) WithPayload(payload *models.Plant) *GetPlantByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get plant by Id o k response
func (o *GetPlantByIDOK) SetPayload(payload *models.Plant) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPlantByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPlantByIDBadRequestCode is the HTTP code returned for type GetPlantByIDBadRequest
const GetPlantByIDBadRequestCode int = 400

/*GetPlantByIDBadRequest Invalid ID supplied

swagger:response getPlantByIdBadRequest
*/
type GetPlantByIDBadRequest struct {
}

// NewGetPlantByIDBadRequest creates GetPlantByIDBadRequest with default headers values
func NewGetPlantByIDBadRequest() *GetPlantByIDBadRequest {

	return &GetPlantByIDBadRequest{}
}

// WriteResponse to the client
func (o *GetPlantByIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetPlantByIDNotFoundCode is the HTTP code returned for type GetPlantByIDNotFound
const GetPlantByIDNotFoundCode int = 404

/*GetPlantByIDNotFound Plant not found

swagger:response getPlantByIdNotFound
*/
type GetPlantByIDNotFound struct {
}

// NewGetPlantByIDNotFound creates GetPlantByIDNotFound with default headers values
func NewGetPlantByIDNotFound() *GetPlantByIDNotFound {

	return &GetPlantByIDNotFound{}
}

// WriteResponse to the client
func (o *GetPlantByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
