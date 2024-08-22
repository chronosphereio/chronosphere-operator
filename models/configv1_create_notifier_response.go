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

// Configv1CreateNotifierResponse configv1 create notifier response
//
// swagger:model configv1CreateNotifierResponse
type Configv1CreateNotifierResponse struct {

	// notifier
	Notifier *Configv1Notifier `json:"notifier,omitempty"`
}

// Validate validates this configv1 create notifier response
func (m *Configv1CreateNotifierResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNotifier(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1CreateNotifierResponse) validateNotifier(formats strfmt.Registry) error {
	if swag.IsZero(m.Notifier) { // not required
		return nil
	}

	if m.Notifier != nil {
		if err := m.Notifier.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notifier")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("notifier")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configv1 create notifier response based on the context it is used
func (m *Configv1CreateNotifierResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNotifier(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1CreateNotifierResponse) contextValidateNotifier(ctx context.Context, formats strfmt.Registry) error {

	if m.Notifier != nil {

		if swag.IsZero(m.Notifier) { // not required
			return nil
		}

		if err := m.Notifier.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notifier")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("notifier")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Configv1CreateNotifierResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configv1CreateNotifierResponse) UnmarshalBinary(b []byte) error {
	var res Configv1CreateNotifierResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
