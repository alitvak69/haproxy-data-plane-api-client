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

package cluster

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

// NewInitiateCertificateRefreshParams creates a new InitiateCertificateRefreshParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInitiateCertificateRefreshParams() *InitiateCertificateRefreshParams {
	return &InitiateCertificateRefreshParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInitiateCertificateRefreshParamsWithTimeout creates a new InitiateCertificateRefreshParams object
// with the ability to set a timeout on a request.
func NewInitiateCertificateRefreshParamsWithTimeout(timeout time.Duration) *InitiateCertificateRefreshParams {
	return &InitiateCertificateRefreshParams{
		timeout: timeout,
	}
}

// NewInitiateCertificateRefreshParamsWithContext creates a new InitiateCertificateRefreshParams object
// with the ability to set a context for a request.
func NewInitiateCertificateRefreshParamsWithContext(ctx context.Context) *InitiateCertificateRefreshParams {
	return &InitiateCertificateRefreshParams{
		Context: ctx,
	}
}

// NewInitiateCertificateRefreshParamsWithHTTPClient creates a new InitiateCertificateRefreshParams object
// with the ability to set a custom HTTPClient for a request.
func NewInitiateCertificateRefreshParamsWithHTTPClient(client *http.Client) *InitiateCertificateRefreshParams {
	return &InitiateCertificateRefreshParams{
		HTTPClient: client,
	}
}

/* InitiateCertificateRefreshParams contains all the parameters to send to the API endpoint
   for the initiate certificate refresh operation.

   Typically these are written to a http.Request.
*/
type InitiateCertificateRefreshParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the initiate certificate refresh params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitiateCertificateRefreshParams) WithDefaults() *InitiateCertificateRefreshParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the initiate certificate refresh params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitiateCertificateRefreshParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the initiate certificate refresh params
func (o *InitiateCertificateRefreshParams) WithTimeout(timeout time.Duration) *InitiateCertificateRefreshParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the initiate certificate refresh params
func (o *InitiateCertificateRefreshParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the initiate certificate refresh params
func (o *InitiateCertificateRefreshParams) WithContext(ctx context.Context) *InitiateCertificateRefreshParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the initiate certificate refresh params
func (o *InitiateCertificateRefreshParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the initiate certificate refresh params
func (o *InitiateCertificateRefreshParams) WithHTTPClient(client *http.Client) *InitiateCertificateRefreshParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the initiate certificate refresh params
func (o *InitiateCertificateRefreshParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *InitiateCertificateRefreshParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
