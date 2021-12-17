//go:build go1.18

// Copyright (c) 2020 gomoni contributors
// Under BSD 3-Clause license, see LICENSE file

package null

import (
	"bytes"
	"encoding/json"
)

// TODO: add a better type constaint
//	type marsheable interface {
//		~primiteTypes | json.Unmarshaller
//	}
//

type Type[T any] struct {
	value   T
	present bool
	isNull  bool
}

func (t Type[T]) Value() T {
	return t.value
}

func (t Type[T]) Present() bool {
	return t.present
}

func (t Type[T]) IsNull() bool {
	return t.isNull
}

func New[T any](x T) Type[T] {
	return Type[T]{
		value:   x,
		present: true,
		isNull:  false,
	}
}

func NewNull[T any](x T) Type[T] {
	return Type[T]{
		present: true,
		isNull:  true,
	}
}

func NewMissing[T any](x T) Type[T] {
	return Type[T]{
		present: false,
	}
}

func (t *Type[T]) UnmarshalJSON(data []byte) error {
	t.present = true
	if bytes.Equal(data, null()) {
		t.isNull = true
		return nil
	}
	t.isNull = false

	if err := json.Unmarshal(data, &t.value); err != nil {
		return err
	}

	return nil
}

func (t Type[T]) MarshalJSON() ([]byte, error) {
	if t.isNull {
		return null(), nil
	}
	return json.Marshal(t.value)
}

func null() []byte {
	return []byte(`null`)
}
