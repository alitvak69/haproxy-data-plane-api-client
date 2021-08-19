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

package configuration

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

// GetHAProxyConfigurationReader is a Reader for the GetHAProxyConfiguration structure.
type GetHAProxyConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHAProxyConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHAProxyConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetHAProxyConfigurationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHAProxyConfigurationOK creates a GetHAProxyConfigurationOK with default headers values
func NewGetHAProxyConfigurationOK() *GetHAProxyConfigurationOK {
	return &GetHAProxyConfigurationOK{}
}

/* GetHAProxyConfigurationOK describes a response with status code 200, with default header values.

Operation successful
*/
type GetHAProxyConfigurationOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetHAProxyConfigurationOKBody
}

func (o *GetHAProxyConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/raw][%d] getHAProxyConfigurationOK  %+v", 200, o.Payload)
}
func (o *GetHAProxyConfigurationOK) GetPayload() *GetHAProxyConfigurationOKBody {
	return o.Payload
}

func (o *GetHAProxyConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetHAProxyConfigurationOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHAProxyConfigurationDefault creates a GetHAProxyConfigurationDefault with default headers values
func NewGetHAProxyConfigurationDefault(code int) *GetHAProxyConfigurationDefault {
	return &GetHAProxyConfigurationDefault{
		_statusCode: code,
	}
}

/* GetHAProxyConfigurationDefault describes a response with status code -1, with default header values.

General Error
*/
type GetHAProxyConfigurationDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get h a proxy configuration default response
func (o *GetHAProxyConfigurationDefault) Code() int {
	return o._statusCode
}

func (o *GetHAProxyConfigurationDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/raw][%d] getHAProxyConfiguration default  %+v", o._statusCode, o.Payload)
}
func (o *GetHAProxyConfigurationDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHAProxyConfigurationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetHAProxyConfigurationOKBody get h a proxy configuration o k body
swagger:model GetHAProxyConfigurationOKBody
*/
type GetHAProxyConfigurationOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data *string `json:"data"`
}

// Validate validates this get h a proxy configuration o k body
func (o *GetHAProxyConfigurationOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHAProxyConfigurationOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getHAProxyConfigurationOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get h a proxy configuration o k body based on context it is used
func (o *GetHAProxyConfigurationOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetHAProxyConfigurationOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHAProxyConfigurationOKBody) UnmarshalBinary(b []byte) error {
	var res GetHAProxyConfigurationOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
