// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Status status
// swagger:model Status
type Status struct {

	// Кол-во разделов в базе данных.
	// Required: true
	Forum int32 `json:"forum"`

	// Кол-во сообщений в базе данных.
	// Required: true
	Post int64 `json:"post"`

	// Кол-во веток обсуждения в базе данных.
	// Required: true
	Thread int32 `json:"thread"`

	// Кол-во пользователей в базе данных.
	// Required: true
	User int32 `json:"user"`
}

// Validate validates this status
func (m *Status) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateForum(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePost(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateThread(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Status) validateForum(formats strfmt.Registry) error {

	if err := validate.Required("forum", "body", int32(m.Forum)); err != nil {
		return err
	}

	return nil
}

func (m *Status) validatePost(formats strfmt.Registry) error {

	if err := validate.Required("post", "body", int64(m.Post)); err != nil {
		return err
	}

	return nil
}

func (m *Status) validateThread(formats strfmt.Registry) error {

	if err := validate.Required("thread", "body", int32(m.Thread)); err != nil {
		return err
	}

	return nil
}

func (m *Status) validateUser(formats strfmt.Registry) error {

	if err := validate.Required("user", "body", int32(m.User)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Status) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Status) UnmarshalBinary(b []byte) error {
	var res Status
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
