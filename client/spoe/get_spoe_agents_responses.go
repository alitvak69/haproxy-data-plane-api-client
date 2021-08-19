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

package spoe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/form3tech-oss/haproxy-data-plane-api-client/models"
)

// GetSpoeAgentsReader is a Reader for the GetSpoeAgents structure.
type GetSpoeAgentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSpoeAgentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSpoeAgentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSpoeAgentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSpoeAgentsOK creates a GetSpoeAgentsOK with default headers values
func NewGetSpoeAgentsOK() *GetSpoeAgentsOK {
	return &GetSpoeAgentsOK{}
}

/* GetSpoeAgentsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetSpoeAgentsOK struct {

	/* Spoe configuration file version
	 */
	ConfigurationVersion string

	Payload *GetSpoeAgentsOKBody
}

func (o *GetSpoeAgentsOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_agents][%d] getSpoeAgentsOK  %+v", 200, o.Payload)
}
func (o *GetSpoeAgentsOK) GetPayload() *GetSpoeAgentsOKBody {
	return o.Payload
}

func (o *GetSpoeAgentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetSpoeAgentsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpoeAgentsDefault creates a GetSpoeAgentsDefault with default headers values
func NewGetSpoeAgentsDefault(code int) *GetSpoeAgentsDefault {
	return &GetSpoeAgentsDefault{
		_statusCode: code,
	}
}

/* GetSpoeAgentsDefault describes a response with status code -1, with default header values.

General Error
*/
type GetSpoeAgentsDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get spoe agents default response
func (o *GetSpoeAgentsDefault) Code() int {
	return o._statusCode
}

func (o *GetSpoeAgentsDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_agents][%d] getSpoeAgents default  %+v", o._statusCode, o.Payload)
}
func (o *GetSpoeAgentsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSpoeAgentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetSpoeAgentsOKBody get spoe agents o k body
swagger:model GetSpoeAgentsOKBody
*/
type GetSpoeAgentsOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.SpoeAgents `json:"data"`
}

// Validate validates this get spoe agents o k body
func (o *GetSpoeAgentsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSpoeAgentsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getSpoeAgentsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getSpoeAgentsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get spoe agents o k body based on the context it is used
func (o *GetSpoeAgentsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSpoeAgentsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Data.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getSpoeAgentsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSpoeAgentsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSpoeAgentsOKBody) UnmarshalBinary(b []byte) error {
	var res GetSpoeAgentsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
