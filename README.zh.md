# rg

皇家守卫

[![Go Reference](https://pkg.go.dev/badge/github.com/yankeguo/rg.svg)](https://pkg.go.dev/github.com/yankeguo/rg)
[![Go](https://github.com/yankeguo/rg/actions/workflows/go.yml/badge.svg)](https://github.com/yankeguo/rg/actions/workflows/go.yml)

在 Go 中使用泛型实现 Throw-Catch 异常处理

## 用法

任何一个末尾返回值为 `error` 类型的函数，均可以使用 `rg.Must` 包裹起来（或者 `rg.Must2`, `rg.Must3` ...），实现异常抛出

使用 `defer rg.Guard(&err)` 实现异常捕获

## 距离

```go
package demo

import (
	"context"
	"encoding/json"
	"gopkg.in/yaml.v3"
	"github.com/yankeguo/rg"
	"os"
)

// 每次都处理 err, 何其丑陋
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

// 使用 rg，优雅的处理异常
func jsonFileToYAML(filename string) (err error) {
	defer rg.Guard(&err)
	buf := rg.Must(os.ReadFile(filename))
	var m map[string]interface{}
	rg.Must0(json.Unmarshal(buf, &m))
	buf = rg.Must(yaml.Marshal(m))
	rg.Must0(os.WriteFile(filename+".yaml", buf, 0640))
	return
}

// 可以设置全局的异常观测回调，支持 context.Context
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

## 许可证

GUO YANKE, MIT License
