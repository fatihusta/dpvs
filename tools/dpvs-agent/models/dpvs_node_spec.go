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

// DpvsNodeSpec dpvs node spec
//
// swagger:model DpvsNodeSpec
type DpvsNodeSpec struct {

	// announce port
	AnnouncePort *VsAnnouncePort `json:"AnnouncePort,omitempty"`

	// laddrs
	Laddrs *LocalAddressExpandList `json:"Laddrs,omitempty"`
}

// Validate validates this dpvs node spec
func (m *DpvsNodeSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnnouncePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLaddrs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DpvsNodeSpec) validateAnnouncePort(formats strfmt.Registry) error {
	if swag.IsZero(m.AnnouncePort) { // not required
		return nil
	}

	if m.AnnouncePort != nil {
		if err := m.AnnouncePort.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AnnouncePort")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("AnnouncePort")
			}
			return err
		}
	}

	return nil
}

func (m *DpvsNodeSpec) validateLaddrs(formats strfmt.Registry) error {
	if swag.IsZero(m.Laddrs) { // not required
		return nil
	}

	if m.Laddrs != nil {
		if err := m.Laddrs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Laddrs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Laddrs")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this dpvs node spec based on the context it is used
func (m *DpvsNodeSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAnnouncePort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLaddrs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DpvsNodeSpec) contextValidateAnnouncePort(ctx context.Context, formats strfmt.Registry) error {

	if m.AnnouncePort != nil {
		if err := m.AnnouncePort.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AnnouncePort")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("AnnouncePort")
			}
			return err
		}
	}

	return nil
}

func (m *DpvsNodeSpec) contextValidateLaddrs(ctx context.Context, formats strfmt.Registry) error {

	if m.Laddrs != nil {
		if err := m.Laddrs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Laddrs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Laddrs")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DpvsNodeSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DpvsNodeSpec) UnmarshalBinary(b []byte) error {
	var res DpvsNodeSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
