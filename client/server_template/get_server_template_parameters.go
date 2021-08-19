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

package server_template

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

// NewGetServerTemplateParams creates a new GetServerTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetServerTemplateParams() *GetServerTemplateParams {
	return &GetServerTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetServerTemplateParamsWithTimeout creates a new GetServerTemplateParams object
// with the ability to set a timeout on a request.
func NewGetServerTemplateParamsWithTimeout(timeout time.Duration) *GetServerTemplateParams {
	return &GetServerTemplateParams{
		timeout: timeout,
	}
}

// NewGetServerTemplateParamsWithContext creates a new GetServerTemplateParams object
// with the ability to set a context for a request.
func NewGetServerTemplateParamsWithContext(ctx context.Context) *GetServerTemplateParams {
	return &GetServerTemplateParams{
		Context: ctx,
	}
}

// NewGetServerTemplateParamsWithHTTPClient creates a new GetServerTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetServerTemplateParamsWithHTTPClient(client *http.Client) *GetServerTemplateParams {
	return &GetServerTemplateParams{
		HTTPClient: client,
	}
}

/* GetServerTemplateParams contains all the parameters to send to the API endpoint
   for the get server template operation.

   Typically these are written to a http.Request.
*/
type GetServerTemplateParams struct {

	/* Backend.

	   Parent backend name
	*/
	Backend string

	/* Prefix.

	   Server template prefix
	*/
	Prefix string

	/* TransactionID.

	   ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	*/
	TransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get server template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServerTemplateParams) WithDefaults() *GetServerTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get server template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServerTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get server template params
func (o *GetServerTemplateParams) WithTimeout(timeout time.Duration) *GetServerTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get server template params
func (o *GetServerTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get server template params
func (o *GetServerTemplateParams) WithContext(ctx context.Context) *GetServerTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get server template params
func (o *GetServerTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get server template params
func (o *GetServerTemplateParams) WithHTTPClient(client *http.Client) *GetServerTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get server template params
func (o *GetServerTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackend adds the backend to the get server template params
func (o *GetServerTemplateParams) WithBackend(backend string) *GetServerTemplateParams {
	o.SetBackend(backend)
	return o
}

// SetBackend adds the backend to the get server template params
func (o *GetServerTemplateParams) SetBackend(backend string) {
	o.Backend = backend
}

// WithPrefix adds the prefix to the get server template params
func (o *GetServerTemplateParams) WithPrefix(prefix string) *GetServerTemplateParams {
	o.SetPrefix(prefix)
	return o
}

// SetPrefix adds the prefix to the get server template params
func (o *GetServerTemplateParams) SetPrefix(prefix string) {
	o.Prefix = prefix
}

// WithTransactionID adds the transactionID to the get server template params
func (o *GetServerTemplateParams) WithTransactionID(transactionID *string) *GetServerTemplateParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get server template params
func (o *GetServerTemplateParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetServerTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param backend
	qrBackend := o.Backend
	qBackend := qrBackend
	if qBackend != "" {

		if err := r.SetQueryParam("backend", qBackend); err != nil {
			return err
		}
	}

	// path param prefix
	if err := r.SetPathParam("prefix", o.Prefix); err != nil {
		return err
	}

	if o.TransactionID != nil {

		// query param transaction_id
		var qrTransactionID string

		if o.TransactionID != nil {
			qrTransactionID = *o.TransactionID
		}
		qTransactionID := qrTransactionID
		if qTransactionID != "" {

			if err := r.SetQueryParam("transaction_id", qTransactionID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
