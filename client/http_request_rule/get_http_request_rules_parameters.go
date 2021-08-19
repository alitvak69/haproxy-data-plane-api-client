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

package http_request_rule

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

// NewGetHTTPRequestRulesParams creates a new GetHTTPRequestRulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetHTTPRequestRulesParams() *GetHTTPRequestRulesParams {
	return &GetHTTPRequestRulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetHTTPRequestRulesParamsWithTimeout creates a new GetHTTPRequestRulesParams object
// with the ability to set a timeout on a request.
func NewGetHTTPRequestRulesParamsWithTimeout(timeout time.Duration) *GetHTTPRequestRulesParams {
	return &GetHTTPRequestRulesParams{
		timeout: timeout,
	}
}

// NewGetHTTPRequestRulesParamsWithContext creates a new GetHTTPRequestRulesParams object
// with the ability to set a context for a request.
func NewGetHTTPRequestRulesParamsWithContext(ctx context.Context) *GetHTTPRequestRulesParams {
	return &GetHTTPRequestRulesParams{
		Context: ctx,
	}
}

// NewGetHTTPRequestRulesParamsWithHTTPClient creates a new GetHTTPRequestRulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetHTTPRequestRulesParamsWithHTTPClient(client *http.Client) *GetHTTPRequestRulesParams {
	return &GetHTTPRequestRulesParams{
		HTTPClient: client,
	}
}

/* GetHTTPRequestRulesParams contains all the parameters to send to the API endpoint
   for the get HTTP request rules operation.

   Typically these are written to a http.Request.
*/
type GetHTTPRequestRulesParams struct {

	/* ParentName.

	   Parent name
	*/
	ParentName string

	/* ParentType.

	   Parent type
	*/
	ParentType string

	/* TransactionID.

	   ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	*/
	TransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get HTTP request rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHTTPRequestRulesParams) WithDefaults() *GetHTTPRequestRulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get HTTP request rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHTTPRequestRulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get HTTP request rules params
func (o *GetHTTPRequestRulesParams) WithTimeout(timeout time.Duration) *GetHTTPRequestRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get HTTP request rules params
func (o *GetHTTPRequestRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get HTTP request rules params
func (o *GetHTTPRequestRulesParams) WithContext(ctx context.Context) *GetHTTPRequestRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get HTTP request rules params
func (o *GetHTTPRequestRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get HTTP request rules params
func (o *GetHTTPRequestRulesParams) WithHTTPClient(client *http.Client) *GetHTTPRequestRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get HTTP request rules params
func (o *GetHTTPRequestRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithParentName adds the parentName to the get HTTP request rules params
func (o *GetHTTPRequestRulesParams) WithParentName(parentName string) *GetHTTPRequestRulesParams {
	o.SetParentName(parentName)
	return o
}

// SetParentName adds the parentName to the get HTTP request rules params
func (o *GetHTTPRequestRulesParams) SetParentName(parentName string) {
	o.ParentName = parentName
}

// WithParentType adds the parentType to the get HTTP request rules params
func (o *GetHTTPRequestRulesParams) WithParentType(parentType string) *GetHTTPRequestRulesParams {
	o.SetParentType(parentType)
	return o
}

// SetParentType adds the parentType to the get HTTP request rules params
func (o *GetHTTPRequestRulesParams) SetParentType(parentType string) {
	o.ParentType = parentType
}

// WithTransactionID adds the transactionID to the get HTTP request rules params
func (o *GetHTTPRequestRulesParams) WithTransactionID(transactionID *string) *GetHTTPRequestRulesParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get HTTP request rules params
func (o *GetHTTPRequestRulesParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetHTTPRequestRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param parent_name
	qrParentName := o.ParentName
	qParentName := qrParentName
	if qParentName != "" {

		if err := r.SetQueryParam("parent_name", qParentName); err != nil {
			return err
		}
	}

	// query param parent_type
	qrParentType := o.ParentType
	qParentType := qrParentType
	if qParentType != "" {

		if err := r.SetQueryParam("parent_type", qParentType); err != nil {
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
