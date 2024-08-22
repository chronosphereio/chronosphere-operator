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

// RoutesNotifierList routes notifier list
//
// swagger:model RoutesNotifierList
type RoutesNotifierList struct {

	// group by
	GroupBy *NotificationPolicyRoutesGroupBy `json:"group_by,omitempty"`

	// Slugs of notifiers that will receive the alerts.
	NotifierSlugs []string `json:"notifier_slugs,omitempty"`

	// Frequency at which to resend alerts.
	RepeatIntervalSecs int32 `json:"repeat_interval_secs,omitempty"`
}

// Validate validates this routes notifier list
func (m *RoutesNotifierList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupBy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RoutesNotifierList) validateGroupBy(formats strfmt.Registry) error {
	if swag.IsZero(m.GroupBy) { // not required
		return nil
	}

	if m.GroupBy != nil {
		if err := m.GroupBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("group_by")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this routes notifier list based on the context it is used
func (m *RoutesNotifierList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGroupBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RoutesNotifierList) contextValidateGroupBy(ctx context.Context, formats strfmt.Registry) error {

	if m.GroupBy != nil {

		if swag.IsZero(m.GroupBy) { // not required
			return nil
		}

		if err := m.GroupBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("group_by")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RoutesNotifierList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoutesNotifierList) UnmarshalBinary(b []byte) error {
	var res RoutesNotifierList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
