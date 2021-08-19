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

package defaults

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new defaults API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for defaults API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetDefaults(params *GetDefaultsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDefaultsOK, error)

	ReplaceDefaults(params *ReplaceDefaultsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceDefaultsOK, *ReplaceDefaultsAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetDefaults returns defaults part of configuration

  Returns defaults part of configuration.
*/
func (a *Client) GetDefaults(params *GetDefaultsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDefaultsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDefaultsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDefaults",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/defaults",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDefaultsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDefaultsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDefaultsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ReplaceDefaults replaces defaults

  Replace defaults part of config
*/
func (a *Client) ReplaceDefaults(params *ReplaceDefaultsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceDefaultsOK, *ReplaceDefaultsAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceDefaultsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "replaceDefaults",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/configuration/defaults",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceDefaultsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReplaceDefaultsOK:
		return value, nil, nil
	case *ReplaceDefaultsAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReplaceDefaultsDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
