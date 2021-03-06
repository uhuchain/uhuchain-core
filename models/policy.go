// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Policy policy
// swagger:model Policy

type Policy struct {

	// claims
	Claims []*Claim `json:"claims"`

	// end date
	EndDate string `json:"endDate,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// insurance Id
	InsuranceID int64 `json:"insuranceId,omitempty"`

	// insurance name
	InsuranceName string `json:"insuranceName,omitempty"`

	// start date
	StartDate string `json:"startDate,omitempty"`
}

/* polymorph Policy claims false */

/* polymorph Policy endDate false */

/* polymorph Policy id false */

/* polymorph Policy insuranceId false */

/* polymorph Policy insuranceName false */

/* polymorph Policy startDate false */

// Validate validates this policy
func (m *Policy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClaims(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Policy) validateClaims(formats strfmt.Registry) error {

	if swag.IsZero(m.Claims) { // not required
		return nil
	}

	for i := 0; i < len(m.Claims); i++ {

		if swag.IsZero(m.Claims[i]) { // not required
			continue
		}

		if m.Claims[i] != nil {

			if err := m.Claims[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("claims" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Policy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Policy) UnmarshalBinary(b []byte) error {
	var res Policy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
