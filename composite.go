package null

import (
	"bytes"
	"encoding/json"
)

// Optional[Nullable[int]]

type Null[T any] struct {
	value T
	isNull bool
}

// Value returns value, false if value is not null, or $zero, true otherwise
func (n Null[T]) Value() (T, bool) {
	var zero T
	if n.isNull {
		return zero, true
	}
	return n.value, false
}

// NewNull creates new null value.
func NewNull[T any](vals ...T) Null[T] {
	switch len(vals) {
	case 0:
		return Null[T]{
			isNull: true,
		}
	case 1:
		return Null[T]{
			value: vals[0],
			isNull: false,
		}
	default:
		panic("Can't pass more than one value")
	}
}

// UnmarshalJSON implements json.Unmarshaler
func (t *Null[T]) UnmarshalJSON(data []byte) error {
	var zero T
	t.value = zero
	if bytes.Equal(data, null) {
		t.isNull = true
		return nil
	}

	if err := json.Unmarshal(data, &t.value); err != nil {
		return err
	}
	t.isNull = false

	return nil
}

type Option[T any] struct {
	value T
	isValid bool
}

// Value returns value, false if value is present, or $zero, true otherwise
func (n Option[T]) Value() (T, bool) {
	var zero T
	if !n.isValid {
		return zero, true
	}
	return n.value, false
}

// NewOption creates new optional value, calling as NewOption[int]() creates missing value
// providing more than one argument panics
func NewOption[T any](vals ...T) Option[T] {
	switch len(vals) {
	case 0:
		return Option[T]{
			isValid: false,
		}
	case 1:
		return Option[T]{
			value: vals[0],
			isValid: true,
		}
	default:
		panic("Can't pass more than one value")
	}
}

// UnmarshalJSON implements json.Unmarshaler
func (t *Option[T]) UnmarshalJSON(data []byte) error {
	var zero T
	t.value = zero
	t.isValid = false

	if err := json.Unmarshal(data, &t.value); err != nil {
		return err
	}
	t.isValid = true

	return nil
}
