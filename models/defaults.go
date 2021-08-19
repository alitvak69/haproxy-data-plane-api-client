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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Defaults Defaults
//
// HAProxy defaults configuration
//
// swagger:model defaults
type Defaults struct {

	// error files
	ErrorFiles []*Errorfile `json:"error_files"`

	// abortonclose
	// Enum: [enabled disabled]
	Abortonclose string `json:"abortonclose,omitempty"`

	// adv check
	// Enum: [ssl-hello-chk smtpchk ldap-check mysql-check pgsql-check tcp-check redis-check httpchk]
	AdvCheck string `json:"adv_check,omitempty"`

	// allbackups
	// Enum: [enabled disabled]
	Allbackups string `json:"allbackups,omitempty"`

	// balance
	Balance *Balance `json:"balance,omitempty"`

	// bind process
	// Pattern: ^[^\s]+$
	BindProcess string `json:"bind_process,omitempty"`

	// check timeout
	CheckTimeout *int64 `json:"check_timeout,omitempty"`

	// clflog
	Clflog bool `json:"clflog,omitempty"`

	// client fin timeout
	ClientFinTimeout *int64 `json:"client_fin_timeout,omitempty"`

	// client timeout
	ClientTimeout *int64 `json:"client_timeout,omitempty"`

	// clitcpka
	// Enum: [enabled disabled]
	Clitcpka string `json:"clitcpka,omitempty"`

	// connect timeout
	ConnectTimeout *int64 `json:"connect_timeout,omitempty"`

	// contstats
	// Enum: [enabled]
	Contstats string `json:"contstats,omitempty"`

	// cookie
	Cookie *Cookie `json:"cookie,omitempty"`

	// default backend
	// Pattern: ^[A-Za-z0-9-_.:]+$
	DefaultBackend string `json:"default_backend,omitempty"`

	// default server
	DefaultServer *DefaultServer `json:"default_server,omitempty"`

	// dontlognull
	// Enum: [enabled disabled]
	Dontlognull string `json:"dontlognull,omitempty"`

	// external check
	// Enum: [enabled disabled]
	ExternalCheck string `json:"external_check,omitempty"`

	// external check command
	// Pattern: ^[^\s]+$
	ExternalCheckCommand string `json:"external_check_command,omitempty"`

	// external check path
	// Pattern: ^[^\s]+$
	ExternalCheckPath string `json:"external_check_path,omitempty"`

	// forwardfor
	Forwardfor *Forwardfor `json:"forwardfor,omitempty"`

	// http buffer request
	// Enum: [enabled disabled]
	HTTPBufferRequest string `json:"http-buffer-request,omitempty"`

	// http check
	HTTPCheck *HTTPCheck `json:"http-check,omitempty"`

	// http use htx
	// Enum: [enabled disabled]
	HTTPUseHtx string `json:"http-use-htx,omitempty"`

	// http connection mode
	// Enum: [httpclose http-server-close http-keep-alive]
	HTTPConnectionMode string `json:"http_connection_mode,omitempty"`

	// http keep alive timeout
	HTTPKeepAliveTimeout *int64 `json:"http_keep_alive_timeout,omitempty"`

	// http pretend keepalive
	// Enum: [enabled disabled]
	HTTPPretendKeepalive string `json:"http_pretend_keepalive,omitempty"`

	// http request timeout
	HTTPRequestTimeout *int64 `json:"http_request_timeout,omitempty"`

	// http reuse
	// Enum: [aggressive always never safe]
	HTTPReuse string `json:"http_reuse,omitempty"`

	// httpchk params
	HttpchkParams *HttpchkParams `json:"httpchk_params,omitempty"`

	// httplog
	Httplog bool `json:"httplog,omitempty"`

	// log format
	LogFormat string `json:"log_format,omitempty"`

	// log format sd
	LogFormatSd string `json:"log_format_sd,omitempty"`

	// log separate errors
	// Enum: [enabled disabled]
	LogSeparateErrors string `json:"log_separate_errors,omitempty"`

	// log tag
	// Pattern: ^[^\s]+$
	LogTag string `json:"log_tag,omitempty"`

	// logasap
	// Enum: [enabled disabled]
	Logasap string `json:"logasap,omitempty"`

	// maxconn
	Maxconn *int64 `json:"maxconn,omitempty"`

	// mode
	// Enum: [tcp http]
	Mode string `json:"mode,omitempty"`

	// monitor uri
	MonitorURI MonitorURI `json:"monitor_uri,omitempty"`

	// mysql check params
	MysqlCheckParams *MysqlCheckParams `json:"mysql_check_params,omitempty"`

	// pgsql check params
	PgsqlCheckParams *PgsqlCheckParams `json:"pgsql_check_params,omitempty"`

	// queue timeout
	QueueTimeout *int64 `json:"queue_timeout,omitempty"`

	// redispatch
	Redispatch *Redispatch `json:"redispatch,omitempty"`

	// retries
	Retries *int64 `json:"retries,omitempty"`

	// server fin timeout
	ServerFinTimeout *int64 `json:"server_fin_timeout,omitempty"`

	// server timeout
	ServerTimeout *int64 `json:"server_timeout,omitempty"`

	// smtpchk params
	SmtpchkParams *SmtpchkParams `json:"smtpchk_params,omitempty"`

	// stats options
	StatsOptions *StatsOptions `json:"stats_options,omitempty"`

	// tcplog
	Tcplog bool `json:"tcplog,omitempty"`

	// tunnel timeout
	TunnelTimeout *int64 `json:"tunnel_timeout,omitempty"`

	// unique id format
	UniqueIDFormat string `json:"unique_id_format,omitempty"`

	// unique id header
	UniqueIDHeader string `json:"unique_id_header,omitempty"`
}

// Validate validates this defaults
func (m *Defaults) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrorFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAbortonclose(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdvCheck(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAllbackups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBindProcess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClitcpka(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContstats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCookie(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultBackend(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDontlognull(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalCheck(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalCheckCommand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalCheckPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateForwardfor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPBufferRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPCheck(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPUseHtx(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPConnectionMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPPretendKeepalive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPReuse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHttpchkParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogSeparateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogTag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogasap(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonitorURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMysqlCheckParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePgsqlCheckParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedispatch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSmtpchkParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatsOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Defaults) validateErrorFiles(formats strfmt.Registry) error {
	if swag.IsZero(m.ErrorFiles) { // not required
		return nil
	}

	for i := 0; i < len(m.ErrorFiles); i++ {
		if swag.IsZero(m.ErrorFiles[i]) { // not required
			continue
		}

		if m.ErrorFiles[i] != nil {
			if err := m.ErrorFiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("error_files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var defaultsTypeAbortonclosePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		defaultsTypeAbortonclosePropEnum = append(defaultsTypeAbortonclosePropEnum, v)
	}
}

const (

	// DefaultsAbortoncloseEnabled captures enum value "enabled"
	DefaultsAbortoncloseEnabled string = "enabled"

	// DefaultsAbortoncloseDisabled captures enum value "disabled"
	DefaultsAbortoncloseDisabled string = "disabled"
)

// prop value enum
func (m *Defaults) validateAbortoncloseEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, defaultsTypeAbortonclosePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Defaults) validateAbortonclose(formats strfmt.Registry) error {
	if swag.IsZero(m.Abortonclose) { // not required
		return nil
	}

	// value enum
	if err := m.validateAbortoncloseEnum("abortonclose", "body", m.Abortonclose); err != nil {
		return err
	}

	return nil
}

var defaultsTypeAdvCheckPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ssl-hello-chk","smtpchk","ldap-check","mysql-check","pgsql-check","tcp-check","redis-check","httpchk"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		defaultsTypeAdvCheckPropEnum = append(defaultsTypeAdvCheckPropEnum, v)
	}
}

const (

	// DefaultsAdvCheckSslDashHelloDashChk captures enum value "ssl-hello-chk"
	DefaultsAdvCheckSslDashHelloDashChk string = "ssl-hello-chk"

	// DefaultsAdvCheckSmtpchk captures enum value "smtpchk"
	DefaultsAdvCheckSmtpchk string = "smtpchk"

	// DefaultsAdvCheckLdapDashCheck captures enum value "ldap-check"
	DefaultsAdvCheckLdapDashCheck string = "ldap-check"

	// DefaultsAdvCheckMysqlDashCheck captures enum value "mysql-check"
	DefaultsAdvCheckMysqlDashCheck string = "mysql-check"

	// DefaultsAdvCheckPgsqlDashCheck captures enum value "pgsql-check"
	DefaultsAdvCheckPgsqlDashCheck string = "pgsql-check"

	// DefaultsAdvCheckTCPDashCheck captures enum value "tcp-check"
	DefaultsAdvCheckTCPDashCheck string = "tcp-check"

	// DefaultsAdvCheckRedisDashCheck captures enum value "redis-check"
	DefaultsAdvCheckRedisDashCheck string = "redis-check"

	// DefaultsAdvCheckHttpchk captures enum value "httpchk"
	DefaultsAdvCheckHttpchk string = "httpchk"
)

// prop value enum
func (m *Defaults) validateAdvCheckEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, defaultsTypeAdvCheckPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Defaults) validateAdvCheck(formats strfmt.Registry) error {
	if swag.IsZero(m.AdvCheck) { // not required
		return nil
	}

	// value enum
	if err := m.validateAdvCheckEnum("adv_check", "body", m.AdvCheck); err != nil {
		return err
	}

	return nil
}

var defaultsTypeAllbackupsPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		defaultsTypeAllbackupsPropEnum = append(defaultsTypeAllbackupsPropEnum, v)
	}
}

const (

	// DefaultsAllbackupsEnabled captures enum value "enabled"
	DefaultsAllbackupsEnabled string = "enabled"

	// DefaultsAllbackupsDisabled captures enum value "disabled"
	DefaultsAllbackupsDisabled string = "disabled"
)

// prop value enum
func (m *Defaults) validateAllbackupsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, defaultsTypeAllbackupsPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Defaults) validateAllbackups(formats strfmt.Registry) error {
	if swag.IsZero(m.Allbackups) { // not required
		return nil
	}

	// value enum
	if err := m.validateAllbackupsEnum("allbackups", "body", m.Allbackups); err != nil {
		return err
	}

	return nil
}

func (m *Defaults) validateBalance(formats strfmt.Registry) error {
	if swag.IsZero(m.Balance) { // not required
		return nil
	}

	if m.Balance != nil {
		if err := m.Balance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("balance")
			}
			return err
		}
	}

	return nil
}

func (m *Defaults) validateBindProcess(formats strfmt.Registry) error {
	if swag.IsZero(m.BindProcess) { // not required
		return nil
	}

	if err := validate.Pattern("bind_process", "body", m.BindProcess, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var defaultsTypeClitcpkaPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		defaultsTypeClitcpkaPropEnum = append(defaultsTypeClitcpkaPropEnum, v)
	}
}

const (

	// DefaultsClitcpkaEnabled captures enum value "enabled"
	DefaultsClitcpkaEnabled string = "enabled"

	// DefaultsClitcpkaDisabled captures enum value "disabled"
	DefaultsClitcpkaDisabled string = "disabled"
)

// prop value enum
func (m *Defaults) validateClitcpkaEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, defaultsTypeClitcpkaPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Defaults) validateClitcpka(formats strfmt.Registry) error {
	if swag.IsZero(m.Clitcpka) { // not required
		return nil
	}

	// value enum
	if err := m.validateClitcpkaEnum("clitcpka", "body", m.Clitcpka); err != nil {
		return err
	}

	return nil
}

var defaultsTypeContstatsPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		defaultsTypeContstatsPropEnum = append(defaultsTypeContstatsPropEnum, v)
	}
}

const (

	// DefaultsContstatsEnabled captures enum value "enabled"
	DefaultsContstatsEnabled string = "enabled"
)

// prop value enum
func (m *Defaults) validateContstatsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, defaultsTypeContstatsPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Defaults) validateContstats(formats strfmt.Registry) error {
	if swag.IsZero(m.Contstats) { // not required
		return nil
	}

	// value enum
	if err := m.validateContstatsEnum("contstats", "body", m.Contstats); err != nil {
		return err
	}

	return nil
}

func (m *Defaults) validateCookie(formats strfmt.Registry) error {
	if swag.IsZero(m.Cookie) { // not required
		return nil
	}

	if m.Cookie != nil {
		if err := m.Cookie.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cookie")
			}
			return err
		}
	}

	return nil
}

func (m *Defaults) validateDefaultBackend(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultBackend) { // not required
		return nil
	}

	if err := validate.Pattern("default_backend", "body", m.DefaultBackend, `^[A-Za-z0-9-_.:]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Defaults) validateDefaultServer(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultServer) { // not required
		return nil
	}

	if m.DefaultServer != nil {
		if err := m.DefaultServer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_server")
			}
			return err
		}
	}

	return nil
}

var defaultsTypeDontlognullPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		defaultsTypeDontlognullPropEnum = append(defaultsTypeDontlognullPropEnum, v)
	}
}

const (

	// DefaultsDontlognullEnabled captures enum value "enabled"
	DefaultsDontlognullEnabled string = "enabled"

	// DefaultsDontlognullDisabled captures enum value "disabled"
	DefaultsDontlognullDisabled string = "disabled"
)

// prop value enum
func (m *Defaults) validateDontlognullEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, defaultsTypeDontlognullPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Defaults) validateDontlognull(formats strfmt.Registry) error {
	if swag.IsZero(m.Dontlognull) { // not required
		return nil
	}

	// value enum
	if err := m.validateDontlognullEnum("dontlognull", "body", m.Dontlognull); err != nil {
		return err
	}

	return nil
}

var defaultsTypeExternalCheckPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		defaultsTypeExternalCheckPropEnum = append(defaultsTypeExternalCheckPropEnum, v)
	}
}

const (

	// DefaultsExternalCheckEnabled captures enum value "enabled"
	DefaultsExternalCheckEnabled string = "enabled"

	// DefaultsExternalCheckDisabled captures enum value "disabled"
	DefaultsExternalCheckDisabled string = "disabled"
)

// prop value enum
func (m *Defaults) validateExternalCheckEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, defaultsTypeExternalCheckPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Defaults) validateExternalCheck(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalCheck) { // not required
		return nil
	}

	// value enum
	if err := m.validateExternalCheckEnum("external_check", "body", m.ExternalCheck); err != nil {
		return err
	}

	return nil
}

func (m *Defaults) validateExternalCheckCommand(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalCheckCommand) { // not required
		return nil
	}

	if err := validate.Pattern("external_check_command", "body", m.ExternalCheckCommand, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Defaults) validateExternalCheckPath(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalCheckPath) { // not required
		return nil
	}

	if err := validate.Pattern("external_check_path", "body", m.ExternalCheckPath, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Defaults) validateForwardfor(formats strfmt.Registry) error {
	if swag.IsZero(m.Forwardfor) { // not required
		return nil
	}

	if m.Forwardfor != nil {
		if err := m.Forwardfor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("forwardfor")
			}
			return err
		}
	}

	return nil
}

var defaultsTypeHTTPBufferRequestPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		defaultsTypeHTTPBufferRequestPropEnum = append(defaultsTypeHTTPBufferRequestPropEnum, v)
	}
}

const (

	// DefaultsHTTPBufferRequestEnabled captures enum value "enabled"
	DefaultsHTTPBufferRequestEnabled string = "enabled"

	// DefaultsHTTPBufferRequestDisabled captures enum value "disabled"
	DefaultsHTTPBufferRequestDisabled string = "disabled"
)

// prop value enum
func (m *Defaults) validateHTTPBufferRequestEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, defaultsTypeHTTPBufferRequestPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Defaults) validateHTTPBufferRequest(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPBufferRequest) { // not required
		return nil
	}

	// value enum
	if err := m.validateHTTPBufferRequestEnum("http-buffer-request", "body", m.HTTPBufferRequest); err != nil {
		return err
	}

	return nil
}

func (m *Defaults) validateHTTPCheck(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPCheck) { // not required
		return nil
	}

	if m.HTTPCheck != nil {
		if err := m.HTTPCheck.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http-check")
			}
			return err
		}
	}

	return nil
}

var defaultsTypeHTTPUseHtxPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		defaultsTypeHTTPUseHtxPropEnum = append(defaultsTypeHTTPUseHtxPropEnum, v)
	}
}

const (

	// DefaultsHTTPUseHtxEnabled captures enum value "enabled"
	DefaultsHTTPUseHtxEnabled string = "enabled"

	// DefaultsHTTPUseHtxDisabled captures enum value "disabled"
	DefaultsHTTPUseHtxDisabled string = "disabled"
)

// prop value enum
func (m *Defaults) validateHTTPUseHtxEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, defaultsTypeHTTPUseHtxPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Defaults) validateHTTPUseHtx(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPUseHtx) { // not required
		return nil
	}

	// value enum
	if err := m.validateHTTPUseHtxEnum("http-use-htx", "body", m.HTTPUseHtx); err != nil {
		return err
	}

	return nil
}

var defaultsTypeHTTPConnectionModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["httpclose","http-server-close","http-keep-alive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		defaultsTypeHTTPConnectionModePropEnum = append(defaultsTypeHTTPConnectionModePropEnum, v)
	}
}

const (

	// DefaultsHTTPConnectionModeHttpclose captures enum value "httpclose"
	DefaultsHTTPConnectionModeHttpclose string = "httpclose"

	// DefaultsHTTPConnectionModeHTTPDashServerDashClose captures enum value "http-server-close"
	DefaultsHTTPConnectionModeHTTPDashServerDashClose string = "http-server-close"

	// DefaultsHTTPConnectionModeHTTPDashKeepDashAlive captures enum value "http-keep-alive"
	DefaultsHTTPConnectionModeHTTPDashKeepDashAlive string = "http-keep-alive"
)

// prop value enum
func (m *Defaults) validateHTTPConnectionModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, defaultsTypeHTTPConnectionModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Defaults) validateHTTPConnectionMode(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPConnectionMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateHTTPConnectionModeEnum("http_connection_mode", "body", m.HTTPConnectionMode); err != nil {
		return err
	}

	return nil
}

var defaultsTypeHTTPPretendKeepalivePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		defaultsTypeHTTPPretendKeepalivePropEnum = append(defaultsTypeHTTPPretendKeepalivePropEnum, v)
	}
}

const (

	// DefaultsHTTPPretendKeepaliveEnabled captures enum value "enabled"
	DefaultsHTTPPretendKeepaliveEnabled string = "enabled"

	// DefaultsHTTPPretendKeepaliveDisabled captures enum value "disabled"
	DefaultsHTTPPretendKeepaliveDisabled string = "disabled"
)

// prop value enum
func (m *Defaults) validateHTTPPretendKeepaliveEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, defaultsTypeHTTPPretendKeepalivePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Defaults) validateHTTPPretendKeepalive(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPPretendKeepalive) { // not required
		return nil
	}

	// value enum
	if err := m.validateHTTPPretendKeepaliveEnum("http_pretend_keepalive", "body", m.HTTPPretendKeepalive); err != nil {
		return err
	}

	return nil
}

var defaultsTypeHTTPReusePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["aggressive","always","never","safe"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		defaultsTypeHTTPReusePropEnum = append(defaultsTypeHTTPReusePropEnum, v)
	}
}

const (

	// DefaultsHTTPReuseAggressive captures enum value "aggressive"
	DefaultsHTTPReuseAggressive string = "aggressive"

	// DefaultsHTTPReuseAlways captures enum value "always"
	DefaultsHTTPReuseAlways string = "always"

	// DefaultsHTTPReuseNever captures enum value "never"
	DefaultsHTTPReuseNever string = "never"

	// DefaultsHTTPReuseSafe captures enum value "safe"
	DefaultsHTTPReuseSafe string = "safe"
)

// prop value enum
func (m *Defaults) validateHTTPReuseEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, defaultsTypeHTTPReusePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Defaults) validateHTTPReuse(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPReuse) { // not required
		return nil
	}

	// value enum
	if err := m.validateHTTPReuseEnum("http_reuse", "body", m.HTTPReuse); err != nil {
		return err
	}

	return nil
}

func (m *Defaults) validateHttpchkParams(formats strfmt.Registry) error {
	if swag.IsZero(m.HttpchkParams) { // not required
		return nil
	}

	if m.HttpchkParams != nil {
		if err := m.HttpchkParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("httpchk_params")
			}
			return err
		}
	}

	return nil
}

var defaultsTypeLogSeparateErrorsPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		defaultsTypeLogSeparateErrorsPropEnum = append(defaultsTypeLogSeparateErrorsPropEnum, v)
	}
}

const (

	// DefaultsLogSeparateErrorsEnabled captures enum value "enabled"
	DefaultsLogSeparateErrorsEnabled string = "enabled"

	// DefaultsLogSeparateErrorsDisabled captures enum value "disabled"
	DefaultsLogSeparateErrorsDisabled string = "disabled"
)

// prop value enum
func (m *Defaults) validateLogSeparateErrorsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, defaultsTypeLogSeparateErrorsPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Defaults) validateLogSeparateErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.LogSeparateErrors) { // not required
		return nil
	}

	// value enum
	if err := m.validateLogSeparateErrorsEnum("log_separate_errors", "body", m.LogSeparateErrors); err != nil {
		return err
	}

	return nil
}

func (m *Defaults) validateLogTag(formats strfmt.Registry) error {
	if swag.IsZero(m.LogTag) { // not required
		return nil
	}

	if err := validate.Pattern("log_tag", "body", m.LogTag, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var defaultsTypeLogasapPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		defaultsTypeLogasapPropEnum = append(defaultsTypeLogasapPropEnum, v)
	}
}

const (

	// DefaultsLogasapEnabled captures enum value "enabled"
	DefaultsLogasapEnabled string = "enabled"

	// DefaultsLogasapDisabled captures enum value "disabled"
	DefaultsLogasapDisabled string = "disabled"
)

// prop value enum
func (m *Defaults) validateLogasapEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, defaultsTypeLogasapPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Defaults) validateLogasap(formats strfmt.Registry) error {
	if swag.IsZero(m.Logasap) { // not required
		return nil
	}

	// value enum
	if err := m.validateLogasapEnum("logasap", "body", m.Logasap); err != nil {
		return err
	}

	return nil
}

var defaultsTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["tcp","http"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		defaultsTypeModePropEnum = append(defaultsTypeModePropEnum, v)
	}
}

const (

	// DefaultsModeTCP captures enum value "tcp"
	DefaultsModeTCP string = "tcp"

	// DefaultsModeHTTP captures enum value "http"
	DefaultsModeHTTP string = "http"
)

// prop value enum
func (m *Defaults) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, defaultsTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Defaults) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *Defaults) validateMonitorURI(formats strfmt.Registry) error {
	if swag.IsZero(m.MonitorURI) { // not required
		return nil
	}

	if err := m.MonitorURI.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("monitor_uri")
		}
		return err
	}

	return nil
}

func (m *Defaults) validateMysqlCheckParams(formats strfmt.Registry) error {
	if swag.IsZero(m.MysqlCheckParams) { // not required
		return nil
	}

	if m.MysqlCheckParams != nil {
		if err := m.MysqlCheckParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mysql_check_params")
			}
			return err
		}
	}

	return nil
}

func (m *Defaults) validatePgsqlCheckParams(formats strfmt.Registry) error {
	if swag.IsZero(m.PgsqlCheckParams) { // not required
		return nil
	}

	if m.PgsqlCheckParams != nil {
		if err := m.PgsqlCheckParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pgsql_check_params")
			}
			return err
		}
	}

	return nil
}

func (m *Defaults) validateRedispatch(formats strfmt.Registry) error {
	if swag.IsZero(m.Redispatch) { // not required
		return nil
	}

	if m.Redispatch != nil {
		if err := m.Redispatch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("redispatch")
			}
			return err
		}
	}

	return nil
}

func (m *Defaults) validateSmtpchkParams(formats strfmt.Registry) error {
	if swag.IsZero(m.SmtpchkParams) { // not required
		return nil
	}

	if m.SmtpchkParams != nil {
		if err := m.SmtpchkParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("smtpchk_params")
			}
			return err
		}
	}

	return nil
}

func (m *Defaults) validateStatsOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.StatsOptions) { // not required
		return nil
	}

	if m.StatsOptions != nil {
		if err := m.StatsOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats_options")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this defaults based on the context it is used
func (m *Defaults) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrorFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBalance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCookie(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefaultServer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateForwardfor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHTTPCheck(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHttpchkParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMonitorURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMysqlCheckParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePgsqlCheckParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRedispatch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSmtpchkParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatsOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Defaults) contextValidateErrorFiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ErrorFiles); i++ {

		if m.ErrorFiles[i] != nil {
			if err := m.ErrorFiles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("error_files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Defaults) contextValidateBalance(ctx context.Context, formats strfmt.Registry) error {

	if m.Balance != nil {
		if err := m.Balance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("balance")
			}
			return err
		}
	}

	return nil
}

func (m *Defaults) contextValidateCookie(ctx context.Context, formats strfmt.Registry) error {

	if m.Cookie != nil {
		if err := m.Cookie.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cookie")
			}
			return err
		}
	}

	return nil
}

func (m *Defaults) contextValidateDefaultServer(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultServer != nil {
		if err := m.DefaultServer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_server")
			}
			return err
		}
	}

	return nil
}

func (m *Defaults) contextValidateForwardfor(ctx context.Context, formats strfmt.Registry) error {

	if m.Forwardfor != nil {
		if err := m.Forwardfor.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("forwardfor")
			}
			return err
		}
	}

	return nil
}

func (m *Defaults) contextValidateHTTPCheck(ctx context.Context, formats strfmt.Registry) error {

	if m.HTTPCheck != nil {
		if err := m.HTTPCheck.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http-check")
			}
			return err
		}
	}

	return nil
}

func (m *Defaults) contextValidateHttpchkParams(ctx context.Context, formats strfmt.Registry) error {

	if m.HttpchkParams != nil {
		if err := m.HttpchkParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("httpchk_params")
			}
			return err
		}
	}

	return nil
}

func (m *Defaults) contextValidateMonitorURI(ctx context.Context, formats strfmt.Registry) error {

	if err := m.MonitorURI.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("monitor_uri")
		}
		return err
	}

	return nil
}

func (m *Defaults) contextValidateMysqlCheckParams(ctx context.Context, formats strfmt.Registry) error {

	if m.MysqlCheckParams != nil {
		if err := m.MysqlCheckParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mysql_check_params")
			}
			return err
		}
	}

	return nil
}

func (m *Defaults) contextValidatePgsqlCheckParams(ctx context.Context, formats strfmt.Registry) error {

	if m.PgsqlCheckParams != nil {
		if err := m.PgsqlCheckParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pgsql_check_params")
			}
			return err
		}
	}

	return nil
}

func (m *Defaults) contextValidateRedispatch(ctx context.Context, formats strfmt.Registry) error {

	if m.Redispatch != nil {
		if err := m.Redispatch.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("redispatch")
			}
			return err
		}
	}

	return nil
}

func (m *Defaults) contextValidateSmtpchkParams(ctx context.Context, formats strfmt.Registry) error {

	if m.SmtpchkParams != nil {
		if err := m.SmtpchkParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("smtpchk_params")
			}
			return err
		}
	}

	return nil
}

func (m *Defaults) contextValidateStatsOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.StatsOptions != nil {
		if err := m.StatsOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats_options")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Defaults) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Defaults) UnmarshalBinary(b []byte) error {
	var res Defaults
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
