# rg - Royal Guard

[![Go Reference](https://pkg.go.dev/badge/github.com/yankeguo/rg.svg)](https://pkg.go.dev/github.com/yankeguo/rg)
[![Go](https://github.com/yankeguo/rg/actions/workflows/go.yml/badge.svg)](https://github.com/yankeguo/rg/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/yankeguo/rg/graph/badge.svg?token=TAJU25VDQJ)](https://codecov.io/gh/yankeguo/rg)

**A modern, generics-based error handling library for Go that brings throw-catch semantics to idiomatic Go code.**

[中文文档](README.zh.md)

## 🚀 Why rg?

Go's explicit error handling is powerful but can lead to verbose code. `rg` provides a clean, panic-based approach that:

- ✅ **Reduces boilerplate**: Eliminate repetitive `if err != nil` checks
- ✅ **Maintains safety**: Automatically converts panics back to errors
- ✅ **Type-safe**: Full generics support for any return type combination
- ✅ **Context-aware**: Built-in support for Go contexts
- ✅ **Hook-friendly**: Customizable error handling with callbacks
- ✅ **Zero dependencies**: Pure Go standard library

## 📦 Installation

```bash
go get github.com/yankeguo/rg
```

## 🎯 Quick Start

Transform verbose error handling:

```go
// Before: Traditional Go error handling
func processFile(filename string) ([]byte, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    processed, err := processData(data)
    if err != nil {
        return nil, err
    }

    return processed, nil
}

// After: With rg
func processFile(filename string) (result []byte, err error) {
    defer rg.Guard(&err)

    data := rg.Must(os.ReadFile(filename))
    result = rg.Must(processData(data))
    return
}
```

## 📖 Core Concepts

### Guard - The Safety Net

`rg.Guard(&err)` acts as a safety net that catches any panic and converts it to an error:

```go
func riskyOperation() (err error) {
    defer rg.Guard(&err)

    // Any panic here will be caught and converted to err
    rg.Must0(someFunctionThatMightFail())
    return nil // Success case
}
```

### Must Functions - The Throwers

The `Must` family of functions check for errors and panic if found:

- `rg.Must0(err)` - For functions returning only an error
- `rg.Must(value, err)` - For functions returning one value + error
- `rg.Must2(v1, v2, err)` - For functions returning two values + error
- ... up to `rg.Must7` for seven values + error

## 🔧 Advanced Features

### Context Support

Pass context information through the error handling chain:

```go
func processWithContext(ctx context.Context) (err error) {
    defer rg.Guard(&err, rg.WithContext(ctx))

    // Context is available in error callbacks
    result := rg.Must(someNetworkCall(ctx))
    return nil
}
```

### Error Hooks & Callbacks

Customize error handling with global hooks:

```go
func init() {
    // Global error hook (deprecated, use OnGuardWithContext)
    rg.OnGuard = func(r any) {
        log.Printf("Error caught: %v", r)
    }

    // Context-aware error hook
    rg.OnGuardWithContext = func(ctx context.Context, r any) {
        // Extract request ID, user info, etc. from context
        if reqID := ctx.Value("request_id"); reqID != nil {
            log.Printf("Error in request %v: %v", reqID, r)
        }
    }
}
```

## 💡 Real-World Examples

### File Processing Pipeline

```go
func convertJSONToYAML(inputFile string) (err error) {
    defer rg.Guard(&err)

    // Read and parse JSON
    jsonData := rg.Must(os.ReadFile(inputFile))
    var data map[string]interface{}
    rg.Must0(json.Unmarshal(jsonData, &data))

    // Convert to YAML and write
    yamlData := rg.Must(yaml.Marshal(data))
    rg.Must0(os.WriteFile(inputFile+".yaml", yamlData, 0644))

    return nil
}
```

### HTTP API Handler

```go
func handleUserCreation(w http.ResponseWriter, r *http.Request) {
    var err error
    defer rg.Guard(&err, rg.WithContext(r.Context()))
    defer func() {
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }()

    // Parse request
    var user User
    rg.Must0(json.NewDecoder(r.Body).Decode(&user))

    // Validate and save
    rg.Must0(user.Validate())
    savedUser := rg.Must(userService.Create(user))

    // Return response
    w.Header().Set("Content-Type", "application/json")
    rg.Must0(json.NewEncoder(w).Encode(savedUser))
}
```

### Database Transaction

```go
func transferMoney(from, to int64, amount decimal.Decimal) (err error) {
    defer rg.Guard(&err)

    tx := rg.Must(db.Begin())
    defer tx.Rollback() // Safe to call even after commit

    // Perform transfer operations
    rg.Must0(debitAccount(tx, from, amount))
    rg.Must0(creditAccount(tx, to, amount))
    rg.Must0(logTransfer(tx, from, to, amount))

    rg.Must0(tx.Commit())
    return nil
}
```

## 🎨 Best Practices

1. **Always use `defer rg.Guard(&err)`** at the beginning of functions that need error handling
2. **Keep guard simple**: Don't put complex logic in the defer statement
3. **Use context**: Pass context for better error tracking and debugging
4. **Combine with traditional error handling**: `rg` works well alongside standard Go error handling
5. **Test thoroughly**: Make sure your error paths are covered by tests

## 🤔 When to Use rg

**Great for:**

- Data processing pipelines
- API handlers with multiple validation steps
- File I/O operations
- Database transactions
- Any scenario with multiple sequential operations that can fail

**Consider alternatives for:**

- Simple functions with one or two error checks
- Performance-critical code (panic/recover has overhead)
- Libraries that need to expose traditional Go APIs

## 🛠 Comparison

| Feature        | Traditional Go            | rg                          |
| -------------- | ------------------------- | --------------------------- |
| Error handling | Explicit `if err != nil`  | Automatic with `Must`       |
| Code length    | Longer                    | Shorter                     |
| Performance    | Faster (no panic/recover) | Slightly slower             |
| Readability    | Good for simple cases     | Excellent for complex cases |
| Debugging      | Standard stack traces     | Enhanced with hooks         |

## 📚 API Reference

### Core Functions

- `Guard(err *error, opts ...Option)` - Recover from panic and set error
- `Must0(err error)` - Panic if error is not nil
- `Must[T](value T, err error) T` - Return value or panic on error
- `Must2` through `Must7` - Handle multiple return values

### Options

- `WithContext(ctx context.Context)` - Attach context to guard

### Hooks

- `OnGuard func(r any)` - Global panic hook (deprecated)
- `OnGuardWithContext func(ctx context.Context, r any)` - Context-aware panic hook

## 🤝 Contributing

We welcome contributions! Please feel free to submit issues, feature requests, or pull requests.

## 📄 License

MIT License - see LICENSE file for details.

## 👨‍💻 Author

**GUO YANKE** - [@yankeguo](https://github.com/yankeguo)

---

⭐ If you find this library helpful, please consider giving it a star!
