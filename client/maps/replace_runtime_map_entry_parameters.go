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

package maps

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

// NewReplaceRuntimeMapEntryParams creates a new ReplaceRuntimeMapEntryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplaceRuntimeMapEntryParams() *ReplaceRuntimeMapEntryParams {
	return &ReplaceRuntimeMapEntryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceRuntimeMapEntryParamsWithTimeout creates a new ReplaceRuntimeMapEntryParams object
// with the ability to set a timeout on a request.
func NewReplaceRuntimeMapEntryParamsWithTimeout(timeout time.Duration) *ReplaceRuntimeMapEntryParams {
	return &ReplaceRuntimeMapEntryParams{
		timeout: timeout,
	}
}

// NewReplaceRuntimeMapEntryParamsWithContext creates a new ReplaceRuntimeMapEntryParams object
// with the ability to set a context for a request.
func NewReplaceRuntimeMapEntryParamsWithContext(ctx context.Context) *ReplaceRuntimeMapEntryParams {
	return &ReplaceRuntimeMapEntryParams{
		Context: ctx,
	}
}

// NewReplaceRuntimeMapEntryParamsWithHTTPClient creates a new ReplaceRuntimeMapEntryParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplaceRuntimeMapEntryParamsWithHTTPClient(client *http.Client) *ReplaceRuntimeMapEntryParams {
	return &ReplaceRuntimeMapEntryParams{
		HTTPClient: client,
	}
}

/* ReplaceRuntimeMapEntryParams contains all the parameters to send to the API endpoint
   for the replace runtime map entry operation.

   Typically these are written to a http.Request.
*/
type ReplaceRuntimeMapEntryParams struct {

	// Data.
	Data ReplaceRuntimeMapEntryBody

	/* ForceSync.

	   If true, immediately syncs changes to disk
	*/
	ForceSync *bool

	/* ID.

	   Map id
	*/
	ID string

	/* Map.

	   Mapfile attribute storage_name
	*/
	Map string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the replace runtime map entry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceRuntimeMapEntryParams) WithDefaults() *ReplaceRuntimeMapEntryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replace runtime map entry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceRuntimeMapEntryParams) SetDefaults() {
	var (
		forceSyncDefault = bool(false)
	)

	val := ReplaceRuntimeMapEntryParams{
		ForceSync: &forceSyncDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the replace runtime map entry params
func (o *ReplaceRuntimeMapEntryParams) WithTimeout(timeout time.Duration) *ReplaceRuntimeMapEntryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace runtime map entry params
func (o *ReplaceRuntimeMapEntryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace runtime map entry params
func (o *ReplaceRuntimeMapEntryParams) WithContext(ctx context.Context) *ReplaceRuntimeMapEntryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace runtime map entry params
func (o *ReplaceRuntimeMapEntryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace runtime map entry params
func (o *ReplaceRuntimeMapEntryParams) WithHTTPClient(client *http.Client) *ReplaceRuntimeMapEntryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace runtime map entry params
func (o *ReplaceRuntimeMapEntryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the replace runtime map entry params
func (o *ReplaceRuntimeMapEntryParams) WithData(data ReplaceRuntimeMapEntryBody) *ReplaceRuntimeMapEntryParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the replace runtime map entry params
func (o *ReplaceRuntimeMapEntryParams) SetData(data ReplaceRuntimeMapEntryBody) {
	o.Data = data
}

// WithForceSync adds the forceSync to the replace runtime map entry params
func (o *ReplaceRuntimeMapEntryParams) WithForceSync(forceSync *bool) *ReplaceRuntimeMapEntryParams {
	o.SetForceSync(forceSync)
	return o
}

// SetForceSync adds the forceSync to the replace runtime map entry params
func (o *ReplaceRuntimeMapEntryParams) SetForceSync(forceSync *bool) {
	o.ForceSync = forceSync
}

// WithID adds the id to the replace runtime map entry params
func (o *ReplaceRuntimeMapEntryParams) WithID(id string) *ReplaceRuntimeMapEntryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the replace runtime map entry params
func (o *ReplaceRuntimeMapEntryParams) SetID(id string) {
	o.ID = id
}

// WithMap adds the mapVar to the replace runtime map entry params
func (o *ReplaceRuntimeMapEntryParams) WithMap(mapVar string) *ReplaceRuntimeMapEntryParams {
	o.SetMap(mapVar)
	return o
}

// SetMap adds the map to the replace runtime map entry params
func (o *ReplaceRuntimeMapEntryParams) SetMap(mapVar string) {
	o.Map = mapVar
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceRuntimeMapEntryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Data); err != nil {
		return err
	}

	if o.ForceSync != nil {

		// query param force_sync
		var qrForceSync bool

		if o.ForceSync != nil {
			qrForceSync = *o.ForceSync
		}
		qForceSync := swag.FormatBool(qrForceSync)
		if qForceSync != "" {

			if err := r.SetQueryParam("force_sync", qForceSync); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param map
	qrMap := o.Map
	qMap := qrMap
	if qMap != "" {

		if err := r.SetQueryParam("map", qMap); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
