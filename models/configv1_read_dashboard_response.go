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

// Configv1ReadDashboardResponse configv1 read dashboard response
//
// swagger:model configv1ReadDashboardResponse
type Configv1ReadDashboardResponse struct {

	// dashboard
	Dashboard *Configv1Dashboard `json:"dashboard,omitempty"`
}

// Validate validates this configv1 read dashboard response
func (m *Configv1ReadDashboardResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDashboard(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1ReadDashboardResponse) validateDashboard(formats strfmt.Registry) error {
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

// ContextValidate validate this configv1 read dashboard response based on the context it is used
func (m *Configv1ReadDashboardResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDashboard(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1ReadDashboardResponse) contextValidateDashboard(ctx context.Context, formats strfmt.Registry) error {

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
func (m *Configv1ReadDashboardResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configv1ReadDashboardResponse) UnmarshalBinary(b []byte) error {
	var res Configv1ReadDashboardResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
