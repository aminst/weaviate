/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */ // Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateSchemaActionsDeleteOKCode is the HTTP code returned for type WeaviateSchemaActionsDeleteOK
const WeaviateSchemaActionsDeleteOKCode int = 200

/*WeaviateSchemaActionsDeleteOK Removed the Action class from the ontology.

swagger:response weaviateSchemaActionsDeleteOK
*/
type WeaviateSchemaActionsDeleteOK struct {
}

// NewWeaviateSchemaActionsDeleteOK creates WeaviateSchemaActionsDeleteOK with default headers values
func NewWeaviateSchemaActionsDeleteOK() *WeaviateSchemaActionsDeleteOK {

	return &WeaviateSchemaActionsDeleteOK{}
}

// WriteResponse to the client
func (o *WeaviateSchemaActionsDeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// WeaviateSchemaActionsDeleteBadRequestCode is the HTTP code returned for type WeaviateSchemaActionsDeleteBadRequest
const WeaviateSchemaActionsDeleteBadRequestCode int = 400

/*WeaviateSchemaActionsDeleteBadRequest Could not delete the Action class.

swagger:response weaviateSchemaActionsDeleteBadRequest
*/
type WeaviateSchemaActionsDeleteBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateSchemaActionsDeleteBadRequest creates WeaviateSchemaActionsDeleteBadRequest with default headers values
func NewWeaviateSchemaActionsDeleteBadRequest() *WeaviateSchemaActionsDeleteBadRequest {

	return &WeaviateSchemaActionsDeleteBadRequest{}
}

// WithPayload adds the payload to the weaviate schema actions delete bad request response
func (o *WeaviateSchemaActionsDeleteBadRequest) WithPayload(payload *models.ErrorResponse) *WeaviateSchemaActionsDeleteBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate schema actions delete bad request response
func (o *WeaviateSchemaActionsDeleteBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateSchemaActionsDeleteBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateSchemaActionsDeleteUnauthorizedCode is the HTTP code returned for type WeaviateSchemaActionsDeleteUnauthorized
const WeaviateSchemaActionsDeleteUnauthorizedCode int = 401

/*WeaviateSchemaActionsDeleteUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateSchemaActionsDeleteUnauthorized
*/
type WeaviateSchemaActionsDeleteUnauthorized struct {
}

// NewWeaviateSchemaActionsDeleteUnauthorized creates WeaviateSchemaActionsDeleteUnauthorized with default headers values
func NewWeaviateSchemaActionsDeleteUnauthorized() *WeaviateSchemaActionsDeleteUnauthorized {

	return &WeaviateSchemaActionsDeleteUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateSchemaActionsDeleteUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
