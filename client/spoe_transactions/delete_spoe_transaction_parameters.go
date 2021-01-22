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

package spoe_transactions

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

// NewDeleteSpoeTransactionParams creates a new DeleteSpoeTransactionParams object
// with the default values initialized.
func NewDeleteSpoeTransactionParams() *DeleteSpoeTransactionParams {
	var ()
	return &DeleteSpoeTransactionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSpoeTransactionParamsWithTimeout creates a new DeleteSpoeTransactionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSpoeTransactionParamsWithTimeout(timeout time.Duration) *DeleteSpoeTransactionParams {
	var ()
	return &DeleteSpoeTransactionParams{

		timeout: timeout,
	}
}

// NewDeleteSpoeTransactionParamsWithContext creates a new DeleteSpoeTransactionParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSpoeTransactionParamsWithContext(ctx context.Context) *DeleteSpoeTransactionParams {
	var ()
	return &DeleteSpoeTransactionParams{

		Context: ctx,
	}
}

// NewDeleteSpoeTransactionParamsWithHTTPClient creates a new DeleteSpoeTransactionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSpoeTransactionParamsWithHTTPClient(client *http.Client) *DeleteSpoeTransactionParams {
	var ()
	return &DeleteSpoeTransactionParams{
		HTTPClient: client,
	}
}

/*DeleteSpoeTransactionParams contains all the parameters to send to the API endpoint
for the delete spoe transaction operation typically these are written to a http.Request
*/
type DeleteSpoeTransactionParams struct {

	/*ID
	  Transaction id

	*/
	ID string
	/*Spoe
	  Spoe file name

	*/
	Spoe string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete spoe transaction params
func (o *DeleteSpoeTransactionParams) WithTimeout(timeout time.Duration) *DeleteSpoeTransactionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete spoe transaction params
func (o *DeleteSpoeTransactionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete spoe transaction params
func (o *DeleteSpoeTransactionParams) WithContext(ctx context.Context) *DeleteSpoeTransactionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete spoe transaction params
func (o *DeleteSpoeTransactionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete spoe transaction params
func (o *DeleteSpoeTransactionParams) WithHTTPClient(client *http.Client) *DeleteSpoeTransactionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete spoe transaction params
func (o *DeleteSpoeTransactionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete spoe transaction params
func (o *DeleteSpoeTransactionParams) WithID(id string) *DeleteSpoeTransactionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete spoe transaction params
func (o *DeleteSpoeTransactionParams) SetID(id string) {
	o.ID = id
}

// WithSpoe adds the spoe to the delete spoe transaction params
func (o *DeleteSpoeTransactionParams) WithSpoe(spoe string) *DeleteSpoeTransactionParams {
	o.SetSpoe(spoe)
	return o
}

// SetSpoe adds the spoe to the delete spoe transaction params
func (o *DeleteSpoeTransactionParams) SetSpoe(spoe string) {
	o.Spoe = spoe
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSpoeTransactionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param spoe
	qrSpoe := o.Spoe
	qSpoe := qrSpoe
	if qSpoe != "" {
		if err := r.SetQueryParam("spoe", qSpoe); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
