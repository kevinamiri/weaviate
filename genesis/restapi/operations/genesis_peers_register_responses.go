//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/semi-technologies/weaviate/genesis/models"
)

// GenesisPeersRegisterOKCode is the HTTP code returned for type GenesisPeersRegisterOK
const GenesisPeersRegisterOKCode int = 200

/*
GenesisPeersRegisterOK Successfully registred the peer to the network.

swagger:response genesisPeersRegisterOK
*/
type GenesisPeersRegisterOK struct {
	/*
	  In: Body
	*/
	Payload *models.PeerRegistrationResponse `json:"body,omitempty"`
}

// NewGenesisPeersRegisterOK creates GenesisPeersRegisterOK with default headers values
func NewGenesisPeersRegisterOK() *GenesisPeersRegisterOK {
	return &GenesisPeersRegisterOK{}
}

// WithPayload adds the payload to the genesis peers register o k response
func (o *GenesisPeersRegisterOK) WithPayload(payload *models.PeerRegistrationResponse) *GenesisPeersRegisterOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the genesis peers register o k response
func (o *GenesisPeersRegisterOK) SetPayload(payload *models.PeerRegistrationResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GenesisPeersRegisterOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GenesisPeersRegisterBadRequestCode is the HTTP code returned for type GenesisPeersRegisterBadRequest
const GenesisPeersRegisterBadRequestCode int = 400

/*
GenesisPeersRegisterBadRequest The weaviate peer is not reachable from the Gensis service.

swagger:response genesisPeersRegisterBadRequest
*/
type GenesisPeersRegisterBadRequest struct{}

// NewGenesisPeersRegisterBadRequest creates GenesisPeersRegisterBadRequest with default headers values
func NewGenesisPeersRegisterBadRequest() *GenesisPeersRegisterBadRequest {
	return &GenesisPeersRegisterBadRequest{}
}

// WriteResponse to the client
func (o *GenesisPeersRegisterBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GenesisPeersRegisterForbiddenCode is the HTTP code returned for type GenesisPeersRegisterForbidden
const GenesisPeersRegisterForbiddenCode int = 403

/*
GenesisPeersRegisterForbidden You are not allowed on the network.

swagger:response genesisPeersRegisterForbidden
*/
type GenesisPeersRegisterForbidden struct{}

// NewGenesisPeersRegisterForbidden creates GenesisPeersRegisterForbidden with default headers values
func NewGenesisPeersRegisterForbidden() *GenesisPeersRegisterForbidden {
	return &GenesisPeersRegisterForbidden{}
}

// WriteResponse to the client
func (o *GenesisPeersRegisterForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
