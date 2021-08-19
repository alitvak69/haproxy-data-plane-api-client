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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/haproxy-data-plane-api-client/models"
)

// DeleteSpoeFileReader is a Reader for the DeleteSpoeFile structure.
type DeleteSpoeFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSpoeFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteSpoeFileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteSpoeFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteSpoeFileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSpoeFileNoContent creates a DeleteSpoeFileNoContent with default headers values
func NewDeleteSpoeFileNoContent() *DeleteSpoeFileNoContent {
	return &DeleteSpoeFileNoContent{}
}

/* DeleteSpoeFileNoContent describes a response with status code 204, with default header values.

SPOE file deleted
*/
type DeleteSpoeFileNoContent struct {
}

func (o *DeleteSpoeFileNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_files/{name}][%d] deleteSpoeFileNoContent ", 204)
}

func (o *DeleteSpoeFileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSpoeFileNotFound creates a DeleteSpoeFileNotFound with default headers values
func NewDeleteSpoeFileNotFound() *DeleteSpoeFileNotFound {
	return &DeleteSpoeFileNotFound{}
}

/* DeleteSpoeFileNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteSpoeFileNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteSpoeFileNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_files/{name}][%d] deleteSpoeFileNotFound  %+v", 404, o.Payload)
}
func (o *DeleteSpoeFileNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteSpoeFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteSpoeFileDefault creates a DeleteSpoeFileDefault with default headers values
func NewDeleteSpoeFileDefault(code int) *DeleteSpoeFileDefault {
	return &DeleteSpoeFileDefault{
		_statusCode: code,
	}
}

/* DeleteSpoeFileDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteSpoeFileDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete spoe file default response
func (o *DeleteSpoeFileDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSpoeFileDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_files/{name}][%d] deleteSpoeFile default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteSpoeFileDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteSpoeFileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
