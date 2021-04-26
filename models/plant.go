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

// Plant plant
//
// swagger:model Plant
type Plant struct {

	// id
	ID int64 `json:"id,omitempty"`

	// name
	// Example: rose
	// Required: true
	Name *string `json:"name"`

	// photo urls
	// Required: true
	PhotoUrls []string `json:"photoUrls" xml:"photoUrl"`
}

// Validate validates this plant
func (m *Plant) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhotoUrls(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Plant) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Plant) validatePhotoUrls(formats strfmt.Registry) error {

	if err := validate.Required("photoUrls", "body", m.PhotoUrls); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this plant based on context it is used
func (m *Plant) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Plant) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Plant) UnmarshalBinary(b []byte) error {
	var res Plant
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
