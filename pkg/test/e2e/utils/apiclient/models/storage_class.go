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
	"github.com/go-openapi/validate"
)

// StorageClass StorageClass represents a Kubernetes StorageClass
//
// swagger:model StorageClass
type StorageClass struct {

	// AllowVolumeExpansion shows whether the storage class allow volume expand
	// +optional
	AllowVolumeExpansion bool `json:"allowVolumeExpansion,omitempty"`

	// Restrict the node topologies where volumes can be dynamically provisioned.
	// Each volume plugin defines its own supported topology specifications.
	// An empty TopologySelectorTerm list means there is no topology restriction.
	// This field is only honored by servers that enable the VolumeScheduling feature.
	// +optional
	// +listType=atomic
	AllowedTopologies []*TopologySelectorTerm `json:"allowedTopologies"`

	// Annotations that can be added to the resource
	Annotations map[string]string `json:"annotations,omitempty"`

	// CreationTimestamp is a timestamp representing the server time when this object was created.
	// Format: date-time
	CreationTimestamp strfmt.DateTime `json:"creationTimestamp,omitempty"`

	// DeletionTimestamp is a timestamp representing the server time when this object was deleted.
	// Format: date-time
	DeletionTimestamp strfmt.DateTime `json:"deletionTimestamp,omitempty"`

	// ID unique value that identifies the resource generated by the server. Read-Only.
	ID string `json:"id,omitempty"`

	// Dynamically provisioned PersistentVolumes of this storage class are
	// created with these mountOptions, e.g. ["ro", "soft"]. Not validated -
	// mount of the PVs will simply fail if one is invalid.
	// +optional
	MountOptions []string `json:"mountOptions"`

	// Name represents human readable name for the resource
	Name string `json:"name,omitempty"`

	// Parameters holds the parameters for the provisioner that should
	// create volumes of this storage class.
	// +optional
	Parameters map[string]string `json:"parameters,omitempty"`

	// Provisioner indicates the type of the provisioner.
	Provisioner string `json:"provisioner,omitempty"`

	// reclaim policy
	ReclaimPolicy PersistentVolumeReclaimPolicy `json:"reclaimPolicy,omitempty"`

	// volume binding mode
	VolumeBindingMode VolumeBindingMode `json:"volumeBindingMode,omitempty"`
}

// Validate validates this storage class
func (m *StorageClass) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowedTopologies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeletionTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReclaimPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeBindingMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageClass) validateAllowedTopologies(formats strfmt.Registry) error {
	if swag.IsZero(m.AllowedTopologies) { // not required
		return nil
	}

	for i := 0; i < len(m.AllowedTopologies); i++ {
		if swag.IsZero(m.AllowedTopologies[i]) { // not required
			continue
		}

		if m.AllowedTopologies[i] != nil {
			if err := m.AllowedTopologies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allowedTopologies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StorageClass) validateCreationTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("creationTimestamp", "body", "date-time", m.CreationTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StorageClass) validateDeletionTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.DeletionTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("deletionTimestamp", "body", "date-time", m.DeletionTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StorageClass) validateReclaimPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.ReclaimPolicy) { // not required
		return nil
	}

	if err := m.ReclaimPolicy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("reclaimPolicy")
		}
		return err
	}

	return nil
}

func (m *StorageClass) validateVolumeBindingMode(formats strfmt.Registry) error {
	if swag.IsZero(m.VolumeBindingMode) { // not required
		return nil
	}

	if err := m.VolumeBindingMode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("volumeBindingMode")
		}
		return err
	}

	return nil
}

// ContextValidate validate this storage class based on the context it is used
func (m *StorageClass) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllowedTopologies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReclaimPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolumeBindingMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageClass) contextValidateAllowedTopologies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AllowedTopologies); i++ {

		if m.AllowedTopologies[i] != nil {
			if err := m.AllowedTopologies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allowedTopologies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StorageClass) contextValidateReclaimPolicy(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ReclaimPolicy.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("reclaimPolicy")
		}
		return err
	}

	return nil
}

func (m *StorageClass) contextValidateVolumeBindingMode(ctx context.Context, formats strfmt.Registry) error {

	if err := m.VolumeBindingMode.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("volumeBindingMode")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageClass) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageClass) UnmarshalBinary(b []byte) error {
	var res StorageClass
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
