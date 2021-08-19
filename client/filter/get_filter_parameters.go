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

package filter

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
	"github.com/go-openapi/swag"
)

// NewGetFilterParams creates a new GetFilterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFilterParams() *GetFilterParams {
	return &GetFilterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFilterParamsWithTimeout creates a new GetFilterParams object
// with the ability to set a timeout on a request.
func NewGetFilterParamsWithTimeout(timeout time.Duration) *GetFilterParams {
	return &GetFilterParams{
		timeout: timeout,
	}
}

// NewGetFilterParamsWithContext creates a new GetFilterParams object
// with the ability to set a context for a request.
func NewGetFilterParamsWithContext(ctx context.Context) *GetFilterParams {
	return &GetFilterParams{
		Context: ctx,
	}
}

// NewGetFilterParamsWithHTTPClient creates a new GetFilterParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFilterParamsWithHTTPClient(client *http.Client) *GetFilterParams {
	return &GetFilterParams{
		HTTPClient: client,
	}
}

/* GetFilterParams contains all the parameters to send to the API endpoint
   for the get filter operation.

   Typically these are written to a http.Request.
*/
type GetFilterParams struct {

	/* Index.

	   Filter Index
	*/
	Index int64

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

// WithDefaults hydrates default values in the get filter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFilterParams) WithDefaults() *GetFilterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get filter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFilterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get filter params
func (o *GetFilterParams) WithTimeout(timeout time.Duration) *GetFilterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get filter params
func (o *GetFilterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get filter params
func (o *GetFilterParams) WithContext(ctx context.Context) *GetFilterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get filter params
func (o *GetFilterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get filter params
func (o *GetFilterParams) WithHTTPClient(client *http.Client) *GetFilterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get filter params
func (o *GetFilterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIndex adds the index to the get filter params
func (o *GetFilterParams) WithIndex(index int64) *GetFilterParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the get filter params
func (o *GetFilterParams) SetIndex(index int64) {
	o.Index = index
}

// WithParentName adds the parentName to the get filter params
func (o *GetFilterParams) WithParentName(parentName string) *GetFilterParams {
	o.SetParentName(parentName)
	return o
}

// SetParentName adds the parentName to the get filter params
func (o *GetFilterParams) SetParentName(parentName string) {
	o.ParentName = parentName
}

// WithParentType adds the parentType to the get filter params
func (o *GetFilterParams) WithParentType(parentType string) *GetFilterParams {
	o.SetParentType(parentType)
	return o
}

// SetParentType adds the parentType to the get filter params
func (o *GetFilterParams) SetParentType(parentType string) {
	o.ParentType = parentType
}

// WithTransactionID adds the transactionID to the get filter params
func (o *GetFilterParams) WithTransactionID(transactionID *string) *GetFilterParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get filter params
func (o *GetFilterParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFilterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.Index)); err != nil {
		return err
	}

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
