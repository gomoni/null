//go:build go1.18

// Copyright (c) 2020 gomoni contributors
// Under BSD 3-Clause license, see LICENSE file

package null

import (
	"bytes"
	"encoding/json"
	"io"
)

// TODO: add a better type constaint
//	type marsheable interface {
//		~primiteTypes | json.Unmarshaller
//	}
//
//	add fmt.Printf support(?)
//		equal method?
//  how to marshal undefined value? - if err, then use better one than io.EOF

type Type[T any] struct {
	value   T
	defined bool
	isNull  bool
}

func (t Type[T]) Value() T {
	return t.value
}

func (t Type[T]) IsUndefined() bool {
	return !t.defined
}

func (t Type[T]) IsNull() bool {
	return t.isNull
}

func New[T any](x T) Type[T] {
	return Type[T]{
		value:   x,
		defined: true,
		isNull:  false,
	}
}

func NewNull[T any]() Type[T] {
	return Type[T]{
		defined: true,
		isNull:  true,
	}
}

func NewUndefined[T any]() Type[T] {
	return Type[T]{
		defined: false,
	}
}

func (t *Type[T]) UnmarshalJSON(data []byte) error {
	t.defined = true
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
	// FIXME: marshalling for t.defined = false must err
	// in theory it can make a sense with omitempty, but there's no way of knowing it
	if t.IsUndefined() {
		return nil, io.EOF		// FIXME: better error
	}
	if t.isNull {
		return null(), nil
	}
	return json.Marshal(t.value)
}

func null() []byte {
	return []byte(`null`)
}
