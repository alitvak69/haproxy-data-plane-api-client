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

package stick_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/haproxy-data-plane-api-client/models"
)

// CreateStickRuleReader is a Reader for the CreateStickRule structure.
type CreateStickRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateStickRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateStickRuleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateStickRuleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateStickRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateStickRuleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateStickRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateStickRuleCreated creates a CreateStickRuleCreated with default headers values
func NewCreateStickRuleCreated() *CreateStickRuleCreated {
	return &CreateStickRuleCreated{}
}

/* CreateStickRuleCreated describes a response with status code 201, with default header values.

Stick Rule created
*/
type CreateStickRuleCreated struct {
	Payload *models.StickRule
}

func (o *CreateStickRuleCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/stick_rules][%d] createStickRuleCreated  %+v", 201, o.Payload)
}
func (o *CreateStickRuleCreated) GetPayload() *models.StickRule {
	return o.Payload
}

func (o *CreateStickRuleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StickRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStickRuleAccepted creates a CreateStickRuleAccepted with default headers values
func NewCreateStickRuleAccepted() *CreateStickRuleAccepted {
	return &CreateStickRuleAccepted{}
}

/* CreateStickRuleAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateStickRuleAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.StickRule
}

func (o *CreateStickRuleAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/stick_rules][%d] createStickRuleAccepted  %+v", 202, o.Payload)
}
func (o *CreateStickRuleAccepted) GetPayload() *models.StickRule {
	return o.Payload
}

func (o *CreateStickRuleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.StickRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStickRuleBadRequest creates a CreateStickRuleBadRequest with default headers values
func NewCreateStickRuleBadRequest() *CreateStickRuleBadRequest {
	return &CreateStickRuleBadRequest{}
}

/* CreateStickRuleBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateStickRuleBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateStickRuleBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/stick_rules][%d] createStickRuleBadRequest  %+v", 400, o.Payload)
}
func (o *CreateStickRuleBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateStickRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateStickRuleConflict creates a CreateStickRuleConflict with default headers values
func NewCreateStickRuleConflict() *CreateStickRuleConflict {
	return &CreateStickRuleConflict{}
}

/* CreateStickRuleConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateStickRuleConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateStickRuleConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/stick_rules][%d] createStickRuleConflict  %+v", 409, o.Payload)
}
func (o *CreateStickRuleConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateStickRuleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateStickRuleDefault creates a CreateStickRuleDefault with default headers values
func NewCreateStickRuleDefault(code int) *CreateStickRuleDefault {
	return &CreateStickRuleDefault{
		_statusCode: code,
	}
}

/* CreateStickRuleDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateStickRuleDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create stick rule default response
func (o *CreateStickRuleDefault) Code() int {
	return o._statusCode
}

func (o *CreateStickRuleDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/stick_rules][%d] createStickRule default  %+v", o._statusCode, o.Payload)
}
func (o *CreateStickRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateStickRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
