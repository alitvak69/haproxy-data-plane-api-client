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

package frontend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new frontend API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for frontend API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateFrontend(params *CreateFrontendParams, authInfo runtime.ClientAuthInfoWriter) (*CreateFrontendCreated, *CreateFrontendAccepted, error)

	DeleteFrontend(params *DeleteFrontendParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteFrontendAccepted, *DeleteFrontendNoContent, error)

	GetFrontend(params *GetFrontendParams, authInfo runtime.ClientAuthInfoWriter) (*GetFrontendOK, error)

	GetFrontends(params *GetFrontendsParams, authInfo runtime.ClientAuthInfoWriter) (*GetFrontendsOK, error)

	ReplaceFrontend(params *ReplaceFrontendParams, authInfo runtime.ClientAuthInfoWriter) (*ReplaceFrontendOK, *ReplaceFrontendAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateFrontend adds a frontend

  Adds a new frontend to the configuration file.
*/
func (a *Client) CreateFrontend(params *CreateFrontendParams, authInfo runtime.ClientAuthInfoWriter) (*CreateFrontendCreated, *CreateFrontendAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateFrontendParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createFrontend",
		Method:             "POST",
		PathPattern:        "/services/haproxy/configuration/frontends",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateFrontendReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateFrontendCreated:
		return value, nil, nil
	case *CreateFrontendAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateFrontendDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteFrontend deletes a frontend

  Deletes a frontend from the configuration by it's name.
*/
func (a *Client) DeleteFrontend(params *DeleteFrontendParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteFrontendAccepted, *DeleteFrontendNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFrontendParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteFrontend",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/configuration/frontends/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteFrontendReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteFrontendAccepted:
		return value, nil, nil
	case *DeleteFrontendNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteFrontendDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetFrontend returns a frontend

  Returns one frontend configuration by it's name.
*/
func (a *Client) GetFrontend(params *GetFrontendParams, authInfo runtime.ClientAuthInfoWriter) (*GetFrontendOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFrontendParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFrontend",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/frontends/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetFrontendReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFrontendOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetFrontendDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetFrontends returns an array of frontends

  Returns an array of all configured frontends.
*/
func (a *Client) GetFrontends(params *GetFrontendsParams, authInfo runtime.ClientAuthInfoWriter) (*GetFrontendsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFrontendsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFrontends",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/frontends",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetFrontendsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFrontendsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetFrontendsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ReplaceFrontend replaces a frontend

  Replaces a frontend configuration by it's name.
*/
func (a *Client) ReplaceFrontend(params *ReplaceFrontendParams, authInfo runtime.ClientAuthInfoWriter) (*ReplaceFrontendOK, *ReplaceFrontendAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceFrontendParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceFrontend",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/configuration/frontends/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceFrontendReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReplaceFrontendOK:
		return value, nil, nil
	case *ReplaceFrontendAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReplaceFrontendDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
