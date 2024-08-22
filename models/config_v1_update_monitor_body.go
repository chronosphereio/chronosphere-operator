// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConfigV1UpdateMonitorBody config v1 update monitor body
//
// swagger:model ConfigV1UpdateMonitorBody
type ConfigV1UpdateMonitorBody struct {

	// If true, the Monitor will be created if it does not already exist, identified by slug. If false, an error will be returned if the Monitor does not already exist.
	CreateIfMissing bool `json:"create_if_missing,omitempty"`

	// If true, the Monitor will not be created nor updated, and no response Monitor will be returned. The response will return an error if the given Monitor is invalid.
	DryRun bool `json:"dry_run,omitempty"`

	// monitor
	Monitor *Configv1Monitor `json:"monitor,omitempty"`
}

// Validate validates this config v1 update monitor body
func (m *ConfigV1UpdateMonitorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMonitor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigV1UpdateMonitorBody) validateMonitor(formats strfmt.Registry) error {
	if swag.IsZero(m.Monitor) { // not required
		return nil
	}

	if m.Monitor != nil {
		if err := m.Monitor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("monitor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("monitor")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config v1 update monitor body based on the context it is used
func (m *ConfigV1UpdateMonitorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMonitor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigV1UpdateMonitorBody) contextValidateMonitor(ctx context.Context, formats strfmt.Registry) error {

	if m.Monitor != nil {

		if swag.IsZero(m.Monitor) { // not required
			return nil
		}

		if err := m.Monitor.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("monitor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("monitor")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigV1UpdateMonitorBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigV1UpdateMonitorBody) UnmarshalBinary(b []byte) error {
	var res ConfigV1UpdateMonitorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
