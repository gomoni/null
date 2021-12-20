# null
Null and undefined aware type wrapper for Go types. Provides richer semantics
for parsing json when needed.

# Usage

```go
var s struct {
    Key Type[int] `json:"key"`
}
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
