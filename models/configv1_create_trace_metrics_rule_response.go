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

// Configv1CreateTraceMetricsRuleResponse configv1 create trace metrics rule response
//
// swagger:model configv1CreateTraceMetricsRuleResponse
type Configv1CreateTraceMetricsRuleResponse struct {

	// trace metrics rule
	TraceMetricsRule *Configv1TraceMetricsRule `json:"trace_metrics_rule,omitempty"`
}

// Validate validates this configv1 create trace metrics rule response
func (m *Configv1CreateTraceMetricsRuleResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTraceMetricsRule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1CreateTraceMetricsRuleResponse) validateTraceMetricsRule(formats strfmt.Registry) error {
	if swag.IsZero(m.TraceMetricsRule) { // not required
		return nil
	}

	if m.TraceMetricsRule != nil {
		if err := m.TraceMetricsRule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trace_metrics_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trace_metrics_rule")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configv1 create trace metrics rule response based on the context it is used
func (m *Configv1CreateTraceMetricsRuleResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTraceMetricsRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1CreateTraceMetricsRuleResponse) contextValidateTraceMetricsRule(ctx context.Context, formats strfmt.Registry) error {

	if m.TraceMetricsRule != nil {

		if swag.IsZero(m.TraceMetricsRule) { // not required
			return nil
		}

		if err := m.TraceMetricsRule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trace_metrics_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trace_metrics_rule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Configv1CreateTraceMetricsRuleResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configv1CreateTraceMetricsRuleResponse) UnmarshalBinary(b []byte) error {
	var res Configv1CreateTraceMetricsRuleResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
