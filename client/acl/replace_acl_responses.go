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

package acl

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/haproxytech/models"
)

// ReplaceACLReader is a Reader for the ReplaceACL structure.
type ReplaceACLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceACLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceACLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceACLAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceACLBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceACLNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceACLDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceACLOK creates a ReplaceACLOK with default headers values
func NewReplaceACLOK() *ReplaceACLOK {
	return &ReplaceACLOK{}
}

/*ReplaceACLOK handles this case with default header values.

ACL line replaced
*/
type ReplaceACLOK struct {
	Payload *models.ACL
}

func (o *ReplaceACLOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/acls/{id}][%d] replaceAclOK  %+v", 200, o.Payload)
}

func (o *ReplaceACLOK) GetPayload() *models.ACL {
	return o.Payload
}

func (o *ReplaceACLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ACL)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceACLAccepted creates a ReplaceACLAccepted with default headers values
func NewReplaceACLAccepted() *ReplaceACLAccepted {
	return &ReplaceACLAccepted{}
}

/*ReplaceACLAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type ReplaceACLAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string

	Payload *models.ACL
}

func (o *ReplaceACLAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/acls/{id}][%d] replaceAclAccepted  %+v", 202, o.Payload)
}

func (o *ReplaceACLAccepted) GetPayload() *models.ACL {
	return o.Payload
}

func (o *ReplaceACLAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	o.Payload = new(models.ACL)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceACLBadRequest creates a ReplaceACLBadRequest with default headers values
func NewReplaceACLBadRequest() *ReplaceACLBadRequest {
	return &ReplaceACLBadRequest{}
}

/*ReplaceACLBadRequest handles this case with default header values.

Bad request
*/
type ReplaceACLBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *ReplaceACLBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/acls/{id}][%d] replaceAclBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceACLBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceACLBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceACLNotFound creates a ReplaceACLNotFound with default headers values
func NewReplaceACLNotFound() *ReplaceACLNotFound {
	return &ReplaceACLNotFound{}
}

/*ReplaceACLNotFound handles this case with default header values.

The specified resource was not found
*/
type ReplaceACLNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *ReplaceACLNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/acls/{id}][%d] replaceAclNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceACLNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceACLNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceACLDefault creates a ReplaceACLDefault with default headers values
func NewReplaceACLDefault(code int) *ReplaceACLDefault {
	return &ReplaceACLDefault{
		_statusCode: code,
	}
}

/*ReplaceACLDefault handles this case with default header values.

General Error
*/
type ReplaceACLDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the replace Acl default response
func (o *ReplaceACLDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceACLDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/acls/{id}][%d] replaceAcl default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceACLDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceACLDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
