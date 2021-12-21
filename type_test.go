//go:build go1.18

// Copyright (c) 2020 gomoni contributors
// Under BSD 3-Clause license, see LICENSE file

package null_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"testing"

	. "github.com/gomoni/null"
)

func TestJSON(t *testing.T) {
	t.Parallel()

	var testCases = []struct {
		name      string
		inp       string
		wantValue Type[int]
		wantJSON  string
	}{
		{
			name:      "int value",
			inp:       `{"key": 42}`,
			wantValue: New(42),
			wantJSON:  `{"key":42}`,
		},
		{
			name:      "null value",
			inp:       `{"key": null}`,
			wantValue: NewNull[int](),
			wantJSON:  `{"key":null}`,
		},
		{
			name:      "undefined value",
			inp:       `{}`,
			wantValue: NewUndefined[int](),
			wantJSON:  `{}`,
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			var s struct {
				Key Type[int] `json:"key"`
			}
			err := json.Unmarshal([]byte(tt.inp), &s)
			if err != nil {
				t.Fatalf("unexpected err when unmarshaling: %s", err)
			}

			if !reflect.DeepEqual(s.Key, tt.wantValue) {
				t.Fatalf("unexpected value: got %+v, want: %+v", s.Key, tt.wantValue)
			}

			if _, err := s.Key.Value(); errors.Is(err, ErrUndefined) {
				_, err = json.Marshal(s)
				if !errors.Is(err, ErrUndefined) {
					t.Fatalf("expected err when marshaling undefined: got %+v, want ErrUndefined", err)
				}
			} else {
				out, err := json.Marshal(s)
				if err != nil {
					t.Fatalf("unexpected err when marshaling: %s", err)
				}

				if string(out) != tt.wantJSON {
					t.Fatalf("unexpected json: got %s, want: %s", string(out), tt.wantJSON)
				}
			}

		})
	}
}

func TestSprintf(t *testing.T) {
	type v struct {
		Int int
		Str string
	}
	type s struct {
		Key   Type[string]
		Value Type[v]
	}

	var testCases = []struct {
		name string
		inp  s
		want map[string]string
	}{
		{
			name: "undefined",
			inp:  s{Key: NewUndefined[string](), Value: NewUndefined[v]()},
			want: map[string]string{
				"%s":  "{undefined undefined}",
				"%q":  "{undefined undefined}",
				"%v":  "{undefined undefined}",
				"%+v": "{Key:undefined Value:undefined}",
				"%#v": "null_test.s{Key:NewUndefined[string](), Value:NewUndefined[null_test.v]()}",
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			for format, want := range tt.want {
				t.Run(format, func(t *testing.T) {
					got := fmt.Sprintf(format, tt.inp)
					if got != want {
						t.Errorf("format failed %q: got %q, want %q", format, got, want)
					}
				})
			}
		})
	}

	for _, format := range []string{"%s", "%q", "%v", "%+v", "%#v"} {
		t.Logf("`%s`: %s", format, fmt.Sprintf(format))
	}
}
