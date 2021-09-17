// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Translation translation
//
// swagger:model Translation
type Translation struct {

	// created at
	// Required: true
	CreatedAt *int64 `json:"created_at"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// identifier
	// Required: true
	Identifier *Identifier `json:"identifier"`

	// localization
	// Required: true
	Localization *Localization `json:"localization"`

	// plural
	Plural string `json:"plural,omitempty"`

	// singular
	// Required: true
	Singular *string `json:"singular"`

	// status
	Status string `json:"status,omitempty"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *Translation) UnmarshalJSON(data []byte) error {
	var props struct {

		// created at
		// Required: true
		CreatedAt *int64 `json:"created_at"`

		// id
		// Required: true
		ID *int64 `json:"id"`

		// identifier
		// Required: true
		Identifier *Identifier `json:"identifier"`

		// localization
		// Required: true
		Localization *Localization `json:"localization"`

		// plural
		Plural string `json:"plural,omitempty"`

		// singular
		// Required: true
		Singular *string `json:"singular"`

		// status
		Status string `json:"status,omitempty"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.CreatedAt = props.CreatedAt
	m.ID = props.ID
	m.Identifier = props.Identifier
	m.Localization = props.Localization
	m.Plural = props.Plural
	m.Singular = props.Singular
	m.Status = props.Status
	return nil
}

// Validate validates this translation
func (m *Translation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSingular(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Translation) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *Translation) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Translation) validateIdentifier(formats strfmt.Registry) error {

	if err := validate.Required("identifier", "body", m.Identifier); err != nil {
		return err
	}

	if m.Identifier != nil {
		if err := m.Identifier.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identifier")
			}
			return err
		}
	}

	return nil
}

func (m *Translation) validateLocalization(formats strfmt.Registry) error {

	if err := validate.Required("localization", "body", m.Localization); err != nil {
		return err
	}

	if m.Localization != nil {
		if err := m.Localization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("localization")
			}
			return err
		}
	}

	return nil
}

func (m *Translation) validateSingular(formats strfmt.Registry) error {

	if err := validate.Required("singular", "body", m.Singular); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this translation based on the context it is used
func (m *Translation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIdentifier(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocalization(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Translation) contextValidateIdentifier(ctx context.Context, formats strfmt.Registry) error {

	if m.Identifier != nil {
		if err := m.Identifier.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identifier")
			}
			return err
		}
	}

	return nil
}

func (m *Translation) contextValidateLocalization(ctx context.Context, formats strfmt.Registry) error {

	if m.Localization != nil {
		if err := m.Localization.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("localization")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Translation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Translation) UnmarshalBinary(b []byte) error {
	var res Translation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
