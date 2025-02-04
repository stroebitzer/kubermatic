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

// WatchdogDevice Hardware watchdog device.
//
// Exactly one of its members must be set.
//
// swagger:model WatchdogDevice
type WatchdogDevice struct {

	// i6300esb
	I6300esb *I6300ESBWatchdog `json:"i6300esb,omitempty"`
}

// Validate validates this watchdog device
func (m *WatchdogDevice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateI6300esb(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WatchdogDevice) validateI6300esb(formats strfmt.Registry) error {
	if swag.IsZero(m.I6300esb) { // not required
		return nil
	}

	if m.I6300esb != nil {
		if err := m.I6300esb.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("i6300esb")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this watchdog device based on the context it is used
func (m *WatchdogDevice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateI6300esb(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WatchdogDevice) contextValidateI6300esb(ctx context.Context, formats strfmt.Registry) error {

	if m.I6300esb != nil {
		if err := m.I6300esb.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("i6300esb")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WatchdogDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WatchdogDevice) UnmarshalBinary(b []byte) error {
	var res WatchdogDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
