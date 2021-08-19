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

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/haproxy-data-plane-api-client/models"
)

// ReplaceSiteReader is a Reader for the ReplaceSite structure.
type ReplaceSiteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceSiteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceSiteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceSiteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceSiteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceSiteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceSiteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceSiteOK creates a ReplaceSiteOK with default headers values
func NewReplaceSiteOK() *ReplaceSiteOK {
	return &ReplaceSiteOK{}
}

/* ReplaceSiteOK describes a response with status code 200, with default header values.

Site replaced
*/
type ReplaceSiteOK struct {
	Payload *models.Site
}

func (o *ReplaceSiteOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/sites/{name}][%d] replaceSiteOK  %+v", 200, o.Payload)
}
func (o *ReplaceSiteOK) GetPayload() *models.Site {
	return o.Payload
}

func (o *ReplaceSiteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Site)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceSiteAccepted creates a ReplaceSiteAccepted with default headers values
func NewReplaceSiteAccepted() *ReplaceSiteAccepted {
	return &ReplaceSiteAccepted{}
}

/* ReplaceSiteAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type ReplaceSiteAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.Site
}

func (o *ReplaceSiteAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/sites/{name}][%d] replaceSiteAccepted  %+v", 202, o.Payload)
}
func (o *ReplaceSiteAccepted) GetPayload() *models.Site {
	return o.Payload
}

func (o *ReplaceSiteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.Site)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceSiteBadRequest creates a ReplaceSiteBadRequest with default headers values
func NewReplaceSiteBadRequest() *ReplaceSiteBadRequest {
	return &ReplaceSiteBadRequest{}
}

/* ReplaceSiteBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceSiteBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceSiteBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/sites/{name}][%d] replaceSiteBadRequest  %+v", 400, o.Payload)
}
func (o *ReplaceSiteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceSiteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceSiteNotFound creates a ReplaceSiteNotFound with default headers values
func NewReplaceSiteNotFound() *ReplaceSiteNotFound {
	return &ReplaceSiteNotFound{}
}

/* ReplaceSiteNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type ReplaceSiteNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceSiteNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/sites/{name}][%d] replaceSiteNotFound  %+v", 404, o.Payload)
}
func (o *ReplaceSiteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceSiteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceSiteDefault creates a ReplaceSiteDefault with default headers values
func NewReplaceSiteDefault(code int) *ReplaceSiteDefault {
	return &ReplaceSiteDefault{
		_statusCode: code,
	}
}

/* ReplaceSiteDefault describes a response with status code -1, with default header values.

General Error
*/
type ReplaceSiteDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the replace site default response
func (o *ReplaceSiteDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceSiteDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/sites/{name}][%d] replaceSite default  %+v", o._statusCode, o.Payload)
}
func (o *ReplaceSiteDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceSiteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
