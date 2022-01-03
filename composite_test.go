package null_test

import (
	"encoding/json"
	"reflect"
	"testing"

	. "github.com/gomoni/null"
)

func TestNull(t *testing.T) {

	var testCases = []struct {
		name      string
		inp       string
		wantValue Null[int]
	}{
		{
			name:      "int value",
			inp:       `{"key": 42}`,
			wantValue: NewNull(42),
		},
		{
			name:      "null value",
			inp:       `{"key": null}`,
			wantValue: NewNull[int](),
		},
		{
			name:      "empty value",
			inp:       `{}`,
			wantValue: NewNull[int](0),
		},
	}
	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			var s struct {
				Key Null[int] `json:"key"`
			}
			err := json.Unmarshal([]byte(tt.inp), &s)
			if err != nil {
				t.Fatalf("unexpected err when unmarshaling: %s", err)
			}

			if !reflect.DeepEqual(s.Key, tt.wantValue) {
				t.Fatalf("unexpected value: got %+v, want: %+v", s.Key, tt.wantValue)
			}
		})
	}
}

func TestComposite(t *testing.T) {

	var testCases = []struct {
		name      string
		inp       string
		wantValue Option[Null[int]]
	}{
		{
			name:      "int value",
			inp:       `{"key": 42}`,
			wantValue: NewOption(NewNull(42)),
		},
		{
			name:      "null value",
			inp:       `{"key": null}`,
			wantValue: NewOption(NewNull[int]()),
		},
		{
			name:      "empty value",
			inp:       `{}`,
			wantValue: NewOption[Null[int]](),
		},
	}
	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			var s struct {
				Key Option[Null[int]] `json:"key"`
			}
			err := json.Unmarshal([]byte(tt.inp), &s)
			if err != nil {
				t.Fatalf("unexpected err when unmarshaling: %s", err)
			}

			if !reflect.DeepEqual(s.Key, tt.wantValue) {
				t.Fatalf("unexpected value: got %+v, want: %+v", s.Key, tt.wantValue)
			}
		})
	}
}
