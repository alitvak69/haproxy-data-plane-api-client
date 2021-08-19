// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package service_discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/haproxy-data-plane-api-client/models"
)

// CreateConsulReader is a Reader for the CreateConsul structure.
type CreateConsulReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateConsulReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateConsulCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateConsulBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateConsulConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateConsulDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateConsulCreated creates a CreateConsulCreated with default headers values
func NewCreateConsulCreated() *CreateConsulCreated {
	return &CreateConsulCreated{}
}

/* CreateConsulCreated describes a response with status code 201, with default header values.

Consul created
*/
type CreateConsulCreated struct {
	Payload *models.Consul
}

func (o *CreateConsulCreated) Error() string {
	return fmt.Sprintf("[POST /service_discovery/consul][%d] createConsulCreated  %+v", 201, o.Payload)
}
func (o *CreateConsulCreated) GetPayload() *models.Consul {
	return o.Payload
}

func (o *CreateConsulCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Consul)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConsulBadRequest creates a CreateConsulBadRequest with default headers values
func NewCreateConsulBadRequest() *CreateConsulBadRequest {
	return &CreateConsulBadRequest{}
}

/* CreateConsulBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateConsulBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateConsulBadRequest) Error() string {
	return fmt.Sprintf("[POST /service_discovery/consul][%d] createConsulBadRequest  %+v", 400, o.Payload)
}
func (o *CreateConsulBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateConsulBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConsulConflict creates a CreateConsulConflict with default headers values
func NewCreateConsulConflict() *CreateConsulConflict {
	return &CreateConsulConflict{}
}

/* CreateConsulConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateConsulConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateConsulConflict) Error() string {
	return fmt.Sprintf("[POST /service_discovery/consul][%d] createConsulConflict  %+v", 409, o.Payload)
}
func (o *CreateConsulConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateConsulConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConsulDefault creates a CreateConsulDefault with default headers values
func NewCreateConsulDefault(code int) *CreateConsulDefault {
	return &CreateConsulDefault{
		_statusCode: code,
	}
}

/* CreateConsulDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateConsulDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create consul default response
func (o *CreateConsulDefault) Code() int {
	return o._statusCode
}

func (o *CreateConsulDefault) Error() string {
	return fmt.Sprintf("[POST /service_discovery/consul][%d] createConsul default  %+v", o._statusCode, o.Payload)
}
func (o *CreateConsulDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateConsulDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
