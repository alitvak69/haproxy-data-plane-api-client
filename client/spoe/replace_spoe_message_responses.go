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

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/models"
)

// ReplaceSpoeMessageReader is a Reader for the ReplaceSpoeMessage structure.
type ReplaceSpoeMessageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceSpoeMessageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceSpoeMessageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceSpoeMessageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceSpoeMessageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceSpoeMessageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceSpoeMessageOK creates a ReplaceSpoeMessageOK with default headers values
func NewReplaceSpoeMessageOK() *ReplaceSpoeMessageOK {
	return &ReplaceSpoeMessageOK{}
}

/*ReplaceSpoeMessageOK handles this case with default header values.

Spoe message replaced
*/
type ReplaceSpoeMessageOK struct {
	Payload *models.SpoeMessage
}

func (o *ReplaceSpoeMessageOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe/spoe_messages/{name}][%d] replaceSpoeMessageOK  %+v", 200, o.Payload)
}

func (o *ReplaceSpoeMessageOK) GetPayload() *models.SpoeMessage {
	return o.Payload
}

func (o *ReplaceSpoeMessageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SpoeMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceSpoeMessageBadRequest creates a ReplaceSpoeMessageBadRequest with default headers values
func NewReplaceSpoeMessageBadRequest() *ReplaceSpoeMessageBadRequest {
	return &ReplaceSpoeMessageBadRequest{
		ConfigurationVersion: 0,
	}
}

/*ReplaceSpoeMessageBadRequest handles this case with default header values.

Bad request
*/
type ReplaceSpoeMessageBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *ReplaceSpoeMessageBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe/spoe_messages/{name}][%d] replaceSpoeMessageBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceSpoeMessageBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceSpoeMessageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	configurationVersion, err := swag.ConvertInt64(response.GetHeader("Configuration-Version"))
	if err != nil {
		return errors.InvalidType("Configuration-Version", "header", "int64", response.GetHeader("Configuration-Version"))
	}
	o.ConfigurationVersion = configurationVersion

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceSpoeMessageNotFound creates a ReplaceSpoeMessageNotFound with default headers values
func NewReplaceSpoeMessageNotFound() *ReplaceSpoeMessageNotFound {
	return &ReplaceSpoeMessageNotFound{
		ConfigurationVersion: 0,
	}
}

/*ReplaceSpoeMessageNotFound handles this case with default header values.

The specified resource was not found
*/
type ReplaceSpoeMessageNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *ReplaceSpoeMessageNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe/spoe_messages/{name}][%d] replaceSpoeMessageNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceSpoeMessageNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceSpoeMessageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	configurationVersion, err := swag.ConvertInt64(response.GetHeader("Configuration-Version"))
	if err != nil {
		return errors.InvalidType("Configuration-Version", "header", "int64", response.GetHeader("Configuration-Version"))
	}
	o.ConfigurationVersion = configurationVersion

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceSpoeMessageDefault creates a ReplaceSpoeMessageDefault with default headers values
func NewReplaceSpoeMessageDefault(code int) *ReplaceSpoeMessageDefault {
	return &ReplaceSpoeMessageDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*ReplaceSpoeMessageDefault handles this case with default header values.

General Error
*/
type ReplaceSpoeMessageDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the replace spoe message default response
func (o *ReplaceSpoeMessageDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceSpoeMessageDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe/spoe_messages/{name}][%d] replaceSpoeMessage default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceSpoeMessageDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceSpoeMessageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	configurationVersion, err := swag.ConvertInt64(response.GetHeader("Configuration-Version"))
	if err != nil {
		return errors.InvalidType("Configuration-Version", "header", "int64", response.GetHeader("Configuration-Version"))
	}
	o.ConfigurationVersion = configurationVersion

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
