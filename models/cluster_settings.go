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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterSettings Cluster Settings
//
// Settings related to a cluster.
//
// swagger:model cluster_settings
type ClusterSettings struct {

	// bootstrap key
	BootstrapKey string `json:"bootstrap_key,omitempty"`

	// cluster
	Cluster *ClusterSettingsCluster `json:"cluster,omitempty"`

	// mode
	// Enum: [single cluster]
	Mode string `json:"mode,omitempty"`

	// status
	// Read Only: true
	// Enum: [active unreachable waiting_approval]
	Status string `json:"status,omitempty"`
}

// Validate validates this cluster settings
func (m *ClusterSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSettings) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

var clusterSettingsTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["single","cluster"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterSettingsTypeModePropEnum = append(clusterSettingsTypeModePropEnum, v)
	}
}

const (

	// ClusterSettingsModeSingle captures enum value "single"
	ClusterSettingsModeSingle string = "single"

	// ClusterSettingsModeCluster captures enum value "cluster"
	ClusterSettingsModeCluster string = "cluster"
)

// prop value enum
func (m *ClusterSettings) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clusterSettingsTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ClusterSettings) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

var clusterSettingsTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","unreachable","waiting_approval"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterSettingsTypeStatusPropEnum = append(clusterSettingsTypeStatusPropEnum, v)
	}
}

const (

	// ClusterSettingsStatusActive captures enum value "active"
	ClusterSettingsStatusActive string = "active"

	// ClusterSettingsStatusUnreachable captures enum value "unreachable"
	ClusterSettingsStatusUnreachable string = "unreachable"

	// ClusterSettingsStatusWaitingApproval captures enum value "waiting_approval"
	ClusterSettingsStatusWaitingApproval string = "waiting_approval"
)

// prop value enum
func (m *ClusterSettings) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clusterSettingsTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ClusterSettings) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cluster settings based on the context it is used
func (m *ClusterSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSettings) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSettings) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSettings) UnmarshalBinary(b []byte) error {
	var res ClusterSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterSettingsCluster Cluster controller information
//
// swagger:model ClusterSettingsCluster
type ClusterSettingsCluster struct {

	// address
	// Read Only: true
	// Pattern: ^[^\s]+$
	Address string `json:"address,omitempty"`

	// api base path
	// Read Only: true
	APIBasePath string `json:"api_base_path,omitempty"`

	// description
	// Read Only: true
	Description string `json:"description,omitempty"`

	// name
	// Read Only: true
	Name string `json:"name,omitempty"`

	// port
	// Read Only: true
	// Maximum: 65535
	// Minimum: 1
	Port *int64 `json:"port,omitempty"`
}

// Validate validates this cluster settings cluster
func (m *ClusterSettingsCluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSettingsCluster) validateAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if err := validate.Pattern("cluster"+"."+"address", "body", m.Address, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *ClusterSettingsCluster) validatePort(formats strfmt.Registry) error {
	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if err := validate.MinimumInt("cluster"+"."+"port", "body", *m.Port, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("cluster"+"."+"port", "body", *m.Port, 65535, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cluster settings cluster based on the context it is used
func (m *ClusterSettingsCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAPIBasePath(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSettingsCluster) contextValidateAddress(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cluster"+"."+"address", "body", string(m.Address)); err != nil {
		return err
	}

	return nil
}

func (m *ClusterSettingsCluster) contextValidateAPIBasePath(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cluster"+"."+"api_base_path", "body", string(m.APIBasePath)); err != nil {
		return err
	}

	return nil
}

func (m *ClusterSettingsCluster) contextValidateDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cluster"+"."+"description", "body", string(m.Description)); err != nil {
		return err
	}

	return nil
}

func (m *ClusterSettingsCluster) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cluster"+"."+"name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *ClusterSettingsCluster) contextValidatePort(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cluster"+"."+"port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterSettingsCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSettingsCluster) UnmarshalBinary(b []byte) error {
	var res ClusterSettingsCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
