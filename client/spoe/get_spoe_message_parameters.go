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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetSpoeMessageParams creates a new GetSpoeMessageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSpoeMessageParams() *GetSpoeMessageParams {
	return &GetSpoeMessageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSpoeMessageParamsWithTimeout creates a new GetSpoeMessageParams object
// with the ability to set a timeout on a request.
func NewGetSpoeMessageParamsWithTimeout(timeout time.Duration) *GetSpoeMessageParams {
	return &GetSpoeMessageParams{
		timeout: timeout,
	}
}

// NewGetSpoeMessageParamsWithContext creates a new GetSpoeMessageParams object
// with the ability to set a context for a request.
func NewGetSpoeMessageParamsWithContext(ctx context.Context) *GetSpoeMessageParams {
	return &GetSpoeMessageParams{
		Context: ctx,
	}
}

// NewGetSpoeMessageParamsWithHTTPClient creates a new GetSpoeMessageParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSpoeMessageParamsWithHTTPClient(client *http.Client) *GetSpoeMessageParams {
	return &GetSpoeMessageParams{
		HTTPClient: client,
	}
}

/* GetSpoeMessageParams contains all the parameters to send to the API endpoint
   for the get spoe message operation.

   Typically these are written to a http.Request.
*/
type GetSpoeMessageParams struct {

	/* Name.

	   Spoe message name
	*/
	Name string

	/* Scope.

	   Spoe scope
	*/
	Scope string

	/* Spoe.

	   Spoe file name
	*/
	Spoe string

	/* TransactionID.

	   ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	*/
	TransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get spoe message params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSpoeMessageParams) WithDefaults() *GetSpoeMessageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get spoe message params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSpoeMessageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get spoe message params
func (o *GetSpoeMessageParams) WithTimeout(timeout time.Duration) *GetSpoeMessageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get spoe message params
func (o *GetSpoeMessageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get spoe message params
func (o *GetSpoeMessageParams) WithContext(ctx context.Context) *GetSpoeMessageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get spoe message params
func (o *GetSpoeMessageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get spoe message params
func (o *GetSpoeMessageParams) WithHTTPClient(client *http.Client) *GetSpoeMessageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get spoe message params
func (o *GetSpoeMessageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get spoe message params
func (o *GetSpoeMessageParams) WithName(name string) *GetSpoeMessageParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get spoe message params
func (o *GetSpoeMessageParams) SetName(name string) {
	o.Name = name
}

// WithScope adds the scope to the get spoe message params
func (o *GetSpoeMessageParams) WithScope(scope string) *GetSpoeMessageParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the get spoe message params
func (o *GetSpoeMessageParams) SetScope(scope string) {
	o.Scope = scope
}

// WithSpoe adds the spoe to the get spoe message params
func (o *GetSpoeMessageParams) WithSpoe(spoe string) *GetSpoeMessageParams {
	o.SetSpoe(spoe)
	return o
}

// SetSpoe adds the spoe to the get spoe message params
func (o *GetSpoeMessageParams) SetSpoe(spoe string) {
	o.Spoe = spoe
}

// WithTransactionID adds the transactionID to the get spoe message params
func (o *GetSpoeMessageParams) WithTransactionID(transactionID *string) *GetSpoeMessageParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get spoe message params
func (o *GetSpoeMessageParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSpoeMessageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// query param scope
	qrScope := o.Scope
	qScope := qrScope
	if qScope != "" {

		if err := r.SetQueryParam("scope", qScope); err != nil {
			return err
		}
	}

	// query param spoe
	qrSpoe := o.Spoe
	qSpoe := qrSpoe
	if qSpoe != "" {

		if err := r.SetQueryParam("spoe", qSpoe); err != nil {
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
