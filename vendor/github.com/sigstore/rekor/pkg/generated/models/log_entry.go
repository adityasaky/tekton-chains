// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright 2021 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

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

// LogEntry log entry
//
// swagger:model LogEntry
type LogEntry map[string]LogEntryAnon

// Validate validates this log entry
func (m LogEntry) Validate(formats strfmt.Registry) error {
	var res []error

	for k := range m {

		if swag.IsZero(m[k]) { // not required
			continue
		}
		if val, ok := m[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this log entry based on the context it is used
func (m LogEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	for k := range m {

		if val, ok := m[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// LogEntryAnon log entry anon
//
// swagger:model LogEntryAnon
type LogEntryAnon struct {

	// body
	// Required: true
	Body interface{} `json:"body"`

	// integrated time
	IntegratedTime int64 `json:"integratedTime,omitempty"`

	// log index
	// Required: true
	// Minimum: 0
	LogIndex *int64 `json:"logIndex"`

	// verification
	Verification *LogEntryAnonVerification `json:"verification,omitempty"`
}

// Validate validates this log entry anon
func (m *LogEntryAnon) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerification(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogEntryAnon) validateBody(formats strfmt.Registry) error {

	if m.Body == nil {
		return errors.Required("body", "body", nil)
	}

	return nil
}

func (m *LogEntryAnon) validateLogIndex(formats strfmt.Registry) error {

	if err := validate.Required("logIndex", "body", m.LogIndex); err != nil {
		return err
	}

	if err := validate.MinimumInt("logIndex", "body", *m.LogIndex, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *LogEntryAnon) validateVerification(formats strfmt.Registry) error {
	if swag.IsZero(m.Verification) { // not required
		return nil
	}

	if m.Verification != nil {
		if err := m.Verification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("verification")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this log entry anon based on the context it is used
func (m *LogEntryAnon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVerification(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogEntryAnon) contextValidateVerification(ctx context.Context, formats strfmt.Registry) error {

	if m.Verification != nil {
		if err := m.Verification.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("verification")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogEntryAnon) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogEntryAnon) UnmarshalBinary(b []byte) error {
	var res LogEntryAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LogEntryAnonVerification log entry anon verification
//
// swagger:model LogEntryAnonVerification
type LogEntryAnonVerification struct {

	// inclusion proof
	InclusionProof *InclusionProof `json:"inclusionProof,omitempty"`

	// Signature over the logIndex, body and integratedTime.
	// Format: byte
	SignedEntryTimestamp strfmt.Base64 `json:"signedEntryTimestamp,omitempty"`
}

// Validate validates this log entry anon verification
func (m *LogEntryAnonVerification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInclusionProof(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogEntryAnonVerification) validateInclusionProof(formats strfmt.Registry) error {
	if swag.IsZero(m.InclusionProof) { // not required
		return nil
	}

	if m.InclusionProof != nil {
		if err := m.InclusionProof.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("verification" + "." + "inclusionProof")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this log entry anon verification based on the context it is used
func (m *LogEntryAnonVerification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInclusionProof(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogEntryAnonVerification) contextValidateInclusionProof(ctx context.Context, formats strfmt.Registry) error {

	if m.InclusionProof != nil {
		if err := m.InclusionProof.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("verification" + "." + "inclusionProof")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogEntryAnonVerification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogEntryAnonVerification) UnmarshalBinary(b []byte) error {
	var res LogEntryAnonVerification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}