// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IPSecPolicyOptions IP sec policy options
//
// swagger:model IPSecPolicyOptions
type IPSecPolicyOptions struct {

	// authentications
	// Required: true
	Authentications IPSECPolicyAuthentications `json:"authentications"`

	// dh groups
	// Required: true
	DhGroups IPSECPolicyDhGroups `json:"dhGroups"`

	// encryptions
	// Required: true
	Encryptions IPSECPolicyEncryptions `json:"encryptions"`
}

// Validate validates this IP sec policy options
func (m *IPSecPolicyOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthentications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDhGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncryptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPSecPolicyOptions) validateAuthentications(formats strfmt.Registry) error {

	if err := validate.Required("authentications", "body", m.Authentications); err != nil {
		return err
	}

	if err := m.Authentications.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("authentications")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("authentications")
		}
		return err
	}

	return nil
}

func (m *IPSecPolicyOptions) validateDhGroups(formats strfmt.Registry) error {

	if err := validate.Required("dhGroups", "body", m.DhGroups); err != nil {
		return err
	}

	if err := m.DhGroups.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("dhGroups")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("dhGroups")
		}
		return err
	}

	return nil
}

func (m *IPSecPolicyOptions) validateEncryptions(formats strfmt.Registry) error {

	if err := validate.Required("encryptions", "body", m.Encryptions); err != nil {
		return err
	}

	if err := m.Encryptions.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("encryptions")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("encryptions")
		}
		return err
	}

	return nil
}

// ContextValidate validate this IP sec policy options based on the context it is used
func (m *IPSecPolicyOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthentications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDhGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEncryptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPSecPolicyOptions) contextValidateAuthentications(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Authentications.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("authentications")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("authentications")
		}
		return err
	}

	return nil
}

func (m *IPSecPolicyOptions) contextValidateDhGroups(ctx context.Context, formats strfmt.Registry) error {

	if err := m.DhGroups.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("dhGroups")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("dhGroups")
		}
		return err
	}

	return nil
}

func (m *IPSecPolicyOptions) contextValidateEncryptions(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Encryptions.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("encryptions")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("encryptions")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPSecPolicyOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPSecPolicyOptions) UnmarshalBinary(b []byte) error {
	var res IPSecPolicyOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
