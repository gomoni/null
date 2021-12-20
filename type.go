//go:build go1.18

// Copyright (c) 2020 gomoni contributors
// Under BSD 3-Clause license, see LICENSE file

package null

import (
	"bytes"
	"encoding/json"
	"errors"
)

type state uint8

const (
	isUndefined state = 0
	isNull      state = 1
	isDefined   state = 2
)

var (
	ErrUndefined = errors.New("undefined")
	ErrNull      = errors.New("null")
	null         = []byte(`null`)
)

// Type is null and undefined aware type wrapper for parsing JSON. Unlike Go
// types it can distinguish cases "key": null and key is not present in JSON.
type Type[T any] struct {
	value T
	state state
}

// Value returns value from JSON, or ErrUndefined or ErrNull
func (t Type[T]) Value() (T, error) {
	switch t.state {
	case isUndefined:
		var zero T
		return zero, ErrUndefined
	case isNull:
		var zero T
		return zero, ErrNull
	default:
		return t.value, nil
	}
}

// New creates new wrapped value
func New[T any](x T) Type[T] {
	return Type[T]{
		value: x,
		state: isDefined,
	}
}

// NewNull creates new null value.
func NewNull[T any]() Type[T] {
	return Type[T]{
		state: isNull,
	}
}

// NewNull creates new undefined value.
func NewUndefined[T any]() Type[T] {
	return Type[T]{
		state: isUndefined,
	}
}

// UnmarshalJSON implements json.Unmarshaler
func (t *Type[T]) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, null) {
		t.state = isNull
		return nil
	}

	t.state = isDefined
	if err := json.Unmarshal(data, &t.value); err != nil {
		return err
	}

	return nil
}

// MarshalJSON implements json.Marshaler. Marshaling undefined values results in error.
func (t Type[T]) MarshalJSON() ([]byte, error) {
	switch t.state {
	case isUndefined:
		return nil, ErrUndefined
	case isNull:
		return null, nil
	default:
		return json.Marshal(t.value)
	}
}

// TODO
//	add fmt.Printf support(?)
//		equal method?
