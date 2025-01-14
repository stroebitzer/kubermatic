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

// GPU g p u
//
// swagger:model GPU
type GPU struct {

	// device name
	DeviceName string `json:"deviceName,omitempty"`

	// Name of the GPU device as exposed by a device plugin
	Name string `json:"name,omitempty"`

	// virtual g p u options
	VirtualGPUOptions *VGPUOptions `json:"virtualGPUOptions,omitempty"`
}

// Validate validates this g p u
func (m *GPU) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVirtualGPUOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GPU) validateVirtualGPUOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.VirtualGPUOptions) { // not required
		return nil
	}

	if m.VirtualGPUOptions != nil {
		if err := m.VirtualGPUOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtualGPUOptions")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this g p u based on the context it is used
func (m *GPU) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVirtualGPUOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GPU) contextValidateVirtualGPUOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.VirtualGPUOptions != nil {
		if err := m.VirtualGPUOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtualGPUOptions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GPU) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GPU) UnmarshalBinary(b []byte) error {
	var res GPU
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
