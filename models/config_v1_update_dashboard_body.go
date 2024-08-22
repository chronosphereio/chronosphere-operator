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

// ConfigV1UpdateDashboardBody config v1 update dashboard body
//
// swagger:model ConfigV1UpdateDashboardBody
type ConfigV1UpdateDashboardBody struct {

	// If true, the Dashboard will be created if it does not already exist, identified by slug. If false, an error will be returned if the Dashboard does not already exist.
	CreateIfMissing bool `json:"create_if_missing,omitempty"`

	// dashboard
	Dashboard *Configv1Dashboard `json:"dashboard,omitempty"`

	// If true, the Dashboard will not be created nor updated, and no response Dashboard will be returned. The response will return an error if the given Dashboard is invalid.
	DryRun bool `json:"dry_run,omitempty"`
}

// Validate validates this config v1 update dashboard body
func (m *ConfigV1UpdateDashboardBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDashboard(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigV1UpdateDashboardBody) validateDashboard(formats strfmt.Registry) error {
	if swag.IsZero(m.Dashboard) { // not required
		return nil
	}

	if m.Dashboard != nil {
		if err := m.Dashboard.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dashboard")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dashboard")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config v1 update dashboard body based on the context it is used
func (m *ConfigV1UpdateDashboardBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDashboard(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigV1UpdateDashboardBody) contextValidateDashboard(ctx context.Context, formats strfmt.Registry) error {

	if m.Dashboard != nil {

		if swag.IsZero(m.Dashboard) { // not required
			return nil
		}

		if err := m.Dashboard.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dashboard")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dashboard")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigV1UpdateDashboardBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigV1UpdateDashboardBody) UnmarshalBinary(b []byte) error {
	var res ConfigV1UpdateDashboardBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
