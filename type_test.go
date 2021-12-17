//go:build go1.18

// Copyright (c) 2020 gomoni contributors
// Under BSD 3-Clause license, see LICENSE file

package null_test

import (
	"encoding/json"
	"testing"

	. "github.com/gomoni/null"
)

func TestJSON(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		name      string
		inp       string
		wantValue int
		wantJSON  string
	}{
		{
			name:      "string",
			inp:       `{"key": 42}`,
			wantValue: 42,
			wantJSON:  `{"key":42}`,
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

			got := s.Key.Value()
			if got != tt.wantValue {
				t.Fatalf("unexpected value: got %d, want: %d", got, tt.wantValue)
			}

			out, err := json.Marshal(s)
			if err != nil {
				t.Fatalf("unexpected err when marshaling: %s", err)
			}

			if string(out) != tt.wantJSON {
				t.Fatalf("unexpected json: got %s, want: %s", string(out), tt.wantJSON)
			}

		})
	}

}
