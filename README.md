[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)

# null
null is a Go library providing an easy to use interface to handle `null` value or missing key. `Type[T any]` implements `json.Marshaler` and `json.Unmarshaler` . Because of generics it works for any type which can be unmarshalled from/marshalled into JSON.

Marshalling of undefined is an error though.

# Usage

```go
// define struct with a wrapper type
var s struct {
    Key Type[int] `json:"key"`
}
// unmarshal as usual
err := json.Unmarshal(data, &s)
// access value via Value() method, which returns error for null and undefined cases
value, err := s.Key.Value()
```

| JSON           | Value  | Error        |
|----------------|--------|--------------|
|`{"key": 42}`   | `42`   | `nil`        |
|`{"key": null}` | _zero_ | ErrNull      |
|`{}`            | _zero_ | ErrUndefined |

# Build and test

install gotip (until go 1.18 release)

```bash
go install golang.org/dl/gotip@latest
gotip download
gotip test -v
```
