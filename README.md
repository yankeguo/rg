# rg

[![Go Reference](https://pkg.go.dev/badge/github.com/yankeguo/rg.svg)](https://pkg.go.dev/github.com/yankeguo/rg)
[![Go](https://github.com/yankeguo/rg/actions/workflows/go.yml/badge.svg)](https://github.com/yankeguo/rg/actions/workflows/go.yml)

`rg (Royal Guard)` is a generics based throw-catch approach in Go

## Usage

Any function with the latest return value of type `error` can be wrapped by `rg.Must` (or `rg.Must2`, `rg.Must3` ...)

## Example

```go
package demo

import (
	"context"
	"encoding/json"
	"gopkg.in/yaml.v3"
	"github.com/yankeguo/rg"
	"os"
)

// jsonFileToYAMLUgly this is a demo function WITHOUT rg
func jsonFileToYAMLUgly(filename string) (err error) {
	var buf []byte
	if buf, err = os.ReadFile(filename); err != nil {
		return
	}
	var m map[string]interface{}
	if err = json.Unmarshal(buf, &m); err != nil {
		return
	}
	if buf, err = yaml.Marshal(m); err != nil {
		return
	}
	buf = rg.Must(yaml.Marshal(m))
	if err = os.WriteFile(filename+".yaml", buf, 0640); err != nil {
		return
	}
	return
}

// jsonFileToYAML this is a demo function WITH rg
func jsonFileToYAML(filename string) (err error) {
	defer rg.Guard(&err)
	buf := rg.Must(os.ReadFile(filename))
	var m map[string]interface{}
	rg.Must0(json.Unmarshal(buf, &m))
	buf = rg.Must(yaml.Marshal(m))
	rg.Must0(os.WriteFile(filename+".yaml", buf, 0640))
	return
}

func GuardCallbackOnErrWithContext(ctx context.Context)(err error) {
	rg.OnGuardWithContext = func(ctx context.Context, r any) {
        // do something like logging with ctx on guarded
	}
	// for recovery panic
	defer rg.Guard(&err, rg.WithContext(ctx))
	// if err is not nil, it will panic
	file:=rg.Must(os.ReadFile("something.txt"))
	rg.Must0(os.WriteFile("something.txt", file, 0640))
	return
}
```

## Credits

GUO YANKE, MIT License
