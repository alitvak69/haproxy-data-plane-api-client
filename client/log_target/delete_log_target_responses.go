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

package log_target

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

// DeleteLogTargetReader is a Reader for the DeleteLogTarget structure.
type DeleteLogTargetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLogTargetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteLogTargetAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteLogTargetNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteLogTargetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteLogTargetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteLogTargetAccepted creates a DeleteLogTargetAccepted with default headers values
func NewDeleteLogTargetAccepted() *DeleteLogTargetAccepted {
	return &DeleteLogTargetAccepted{}
}

/*DeleteLogTargetAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type DeleteLogTargetAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string
}

func (o *DeleteLogTargetAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/log_targets/{index}][%d] deleteLogTargetAccepted ", 202)
}

func (o *DeleteLogTargetAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	return nil
}

// NewDeleteLogTargetNoContent creates a DeleteLogTargetNoContent with default headers values
func NewDeleteLogTargetNoContent() *DeleteLogTargetNoContent {
	return &DeleteLogTargetNoContent{}
}

/*DeleteLogTargetNoContent handles this case with default header values.

Log Target deleted
*/
type DeleteLogTargetNoContent struct {
}

func (o *DeleteLogTargetNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/log_targets/{index}][%d] deleteLogTargetNoContent ", 204)
}

func (o *DeleteLogTargetNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLogTargetNotFound creates a DeleteLogTargetNotFound with default headers values
func NewDeleteLogTargetNotFound() *DeleteLogTargetNotFound {
	return &DeleteLogTargetNotFound{
		ConfigurationVersion: 0,
	}
}

/*DeleteLogTargetNotFound handles this case with default header values.

The specified resource was not found
*/
type DeleteLogTargetNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *DeleteLogTargetNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/log_targets/{index}][%d] deleteLogTargetNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLogTargetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteLogTargetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteLogTargetDefault creates a DeleteLogTargetDefault with default headers values
func NewDeleteLogTargetDefault(code int) *DeleteLogTargetDefault {
	return &DeleteLogTargetDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*DeleteLogTargetDefault handles this case with default header values.

General Error
*/
type DeleteLogTargetDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the delete log target default response
func (o *DeleteLogTargetDefault) Code() int {
	return o._statusCode
}

func (o *DeleteLogTargetDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/log_targets/{index}][%d] deleteLogTarget default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteLogTargetDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteLogTargetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
