// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RollupRuleGraphiteLabelPolicy rollup rule graphite label policy
//
// swagger:model RollupRuleGraphiteLabelPolicy
type RollupRuleGraphiteLabelPolicy struct {

	// Required list of labels to replace. Useful for discarding
	// high-cardinality values while still preserving the original positions of
	// the Graphite metric.
	Replace []*GraphiteLabelPolicyReplace `json:"replace,omitempty"`
}

// Validate validates this rollup rule graphite label policy
func (m *RollupRuleGraphiteLabelPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReplace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RollupRuleGraphiteLabelPolicy) validateReplace(formats strfmt.Registry) error {
	if swag.IsZero(m.Replace) { // not required
		return nil
	}

	for i := 0; i < len(m.Replace); i++ {
		if swag.IsZero(m.Replace[i]) { // not required
			continue
		}

		if m.Replace[i] != nil {
			if err := m.Replace[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("replace" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("replace" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this rollup rule graphite label policy based on the context it is used
func (m *RollupRuleGraphiteLabelPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReplace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RollupRuleGraphiteLabelPolicy) contextValidateReplace(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Replace); i++ {

		if m.Replace[i] != nil {

			if swag.IsZero(m.Replace[i]) { // not required
				return nil
			}

			if err := m.Replace[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("replace" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("replace" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RollupRuleGraphiteLabelPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RollupRuleGraphiteLabelPolicy) UnmarshalBinary(b []byte) error {
	var res RollupRuleGraphiteLabelPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
