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

package discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetAPIEndpointsParams creates a new GetAPIEndpointsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIEndpointsParams() *GetAPIEndpointsParams {
	return &GetAPIEndpointsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIEndpointsParamsWithTimeout creates a new GetAPIEndpointsParams object
// with the ability to set a timeout on a request.
func NewGetAPIEndpointsParamsWithTimeout(timeout time.Duration) *GetAPIEndpointsParams {
	return &GetAPIEndpointsParams{
		timeout: timeout,
	}
}

// NewGetAPIEndpointsParamsWithContext creates a new GetAPIEndpointsParams object
// with the ability to set a context for a request.
func NewGetAPIEndpointsParamsWithContext(ctx context.Context) *GetAPIEndpointsParams {
	return &GetAPIEndpointsParams{
		Context: ctx,
	}
}

// NewGetAPIEndpointsParamsWithHTTPClient creates a new GetAPIEndpointsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIEndpointsParamsWithHTTPClient(client *http.Client) *GetAPIEndpointsParams {
	return &GetAPIEndpointsParams{
		HTTPClient: client,
	}
}

/* GetAPIEndpointsParams contains all the parameters to send to the API endpoint
   for the get API endpoints operation.

   Typically these are written to a http.Request.
*/
type GetAPIEndpointsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API endpoints params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIEndpointsParams) WithDefaults() *GetAPIEndpointsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API endpoints params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIEndpointsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API endpoints params
func (o *GetAPIEndpointsParams) WithTimeout(timeout time.Duration) *GetAPIEndpointsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API endpoints params
func (o *GetAPIEndpointsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API endpoints params
func (o *GetAPIEndpointsParams) WithContext(ctx context.Context) *GetAPIEndpointsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API endpoints params
func (o *GetAPIEndpointsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API endpoints params
func (o *GetAPIEndpointsParams) WithHTTPClient(client *http.Client) *GetAPIEndpointsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API endpoints params
func (o *GetAPIEndpointsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIEndpointsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
