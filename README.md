# null
Null and undefined aware type wrapper for Go types. Provides richer semantics
for parsing json when needed.

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
