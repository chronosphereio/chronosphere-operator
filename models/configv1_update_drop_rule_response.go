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

// Configv1UpdateDropRuleResponse configv1 update drop rule response
//
// swagger:model configv1UpdateDropRuleResponse
type Configv1UpdateDropRuleResponse struct {

	// drop rule
	DropRule *Configv1DropRule `json:"drop_rule,omitempty"`
}

// Validate validates this configv1 update drop rule response
func (m *Configv1UpdateDropRuleResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDropRule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1UpdateDropRuleResponse) validateDropRule(formats strfmt.Registry) error {
	if swag.IsZero(m.DropRule) { // not required
		return nil
	}

	if m.DropRule != nil {
		if err := m.DropRule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("drop_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("drop_rule")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configv1 update drop rule response based on the context it is used
func (m *Configv1UpdateDropRuleResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDropRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1UpdateDropRuleResponse) contextValidateDropRule(ctx context.Context, formats strfmt.Registry) error {

	if m.DropRule != nil {

		if swag.IsZero(m.DropRule) { // not required
			return nil
		}

		if err := m.DropRule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("drop_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("drop_rule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Configv1UpdateDropRuleResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configv1UpdateDropRuleResponse) UnmarshalBinary(b []byte) error {
	var res Configv1UpdateDropRuleResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
