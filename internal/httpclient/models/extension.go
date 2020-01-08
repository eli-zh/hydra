// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Extension Extension Extension Extension Extension Extension represents the ASN.1 structure of the same name. See RFC
// 5280, section 4.2.
// swagger:model Extension
type Extension struct {

	// critical
	Critical bool `json:"Critical,omitempty"`

	// Id
	ID ObjectIdentifier `json:"Id,omitempty"`

	// value
	Value []uint8 `json:"Value"`
}

// Validate validates this extension
func (m *Extension) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Extension) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Id")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Extension) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Extension) UnmarshalBinary(b []byte) error {
	var res Extension
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
