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

package nameserver

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

// NewGetNameserverParams creates a new GetNameserverParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNameserverParams() *GetNameserverParams {
	return &GetNameserverParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNameserverParamsWithTimeout creates a new GetNameserverParams object
// with the ability to set a timeout on a request.
func NewGetNameserverParamsWithTimeout(timeout time.Duration) *GetNameserverParams {
	return &GetNameserverParams{
		timeout: timeout,
	}
}

// NewGetNameserverParamsWithContext creates a new GetNameserverParams object
// with the ability to set a context for a request.
func NewGetNameserverParamsWithContext(ctx context.Context) *GetNameserverParams {
	return &GetNameserverParams{
		Context: ctx,
	}
}

// NewGetNameserverParamsWithHTTPClient creates a new GetNameserverParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNameserverParamsWithHTTPClient(client *http.Client) *GetNameserverParams {
	return &GetNameserverParams{
		HTTPClient: client,
	}
}

/* GetNameserverParams contains all the parameters to send to the API endpoint
   for the get nameserver operation.

   Typically these are written to a http.Request.
*/
type GetNameserverParams struct {

	/* Name.

	   Nameserver name
	*/
	Name string

	/* Resolver.

	   Parent resolver name
	*/
	Resolver string

	/* TransactionID.

	   ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	*/
	TransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get nameserver params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNameserverParams) WithDefaults() *GetNameserverParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get nameserver params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNameserverParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get nameserver params
func (o *GetNameserverParams) WithTimeout(timeout time.Duration) *GetNameserverParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nameserver params
func (o *GetNameserverParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nameserver params
func (o *GetNameserverParams) WithContext(ctx context.Context) *GetNameserverParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nameserver params
func (o *GetNameserverParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nameserver params
func (o *GetNameserverParams) WithHTTPClient(client *http.Client) *GetNameserverParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nameserver params
func (o *GetNameserverParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get nameserver params
func (o *GetNameserverParams) WithName(name string) *GetNameserverParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get nameserver params
func (o *GetNameserverParams) SetName(name string) {
	o.Name = name
}

// WithResolver adds the resolver to the get nameserver params
func (o *GetNameserverParams) WithResolver(resolver string) *GetNameserverParams {
	o.SetResolver(resolver)
	return o
}

// SetResolver adds the resolver to the get nameserver params
func (o *GetNameserverParams) SetResolver(resolver string) {
	o.Resolver = resolver
}

// WithTransactionID adds the transactionID to the get nameserver params
func (o *GetNameserverParams) WithTransactionID(transactionID *string) *GetNameserverParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get nameserver params
func (o *GetNameserverParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNameserverParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// query param resolver
	qrResolver := o.Resolver
	qResolver := qrResolver
	if qResolver != "" {

		if err := r.SetQueryParam("resolver", qResolver); err != nil {
			return err
		}
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
