# rg - 皇家守卫

[![Go Reference](https://pkg.go.dev/badge/github.com/yankeguo/rg.svg)](https://pkg.go.dev/github.com/yankeguo/rg)
[![Go](https://github.com/yankeguo/rg/actions/workflows/go.yml/badge.svg)](https://github.com/yankeguo/rg/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/yankeguo/rg/graph/badge.svg?token=TAJU25VDQJ)](https://codecov.io/gh/yankeguo/rg)

**一个现代化的、基于泛型的 Go 错误处理库，为惯用的 Go 代码带来抛出-捕获语义。**

[English Documentation](README.md)

## 🚀 为什么选择 rg？

Go 语言的显式错误处理很强大，但可能导致冗长的代码。`rg` 提供了一种简洁的、基于 panic 的方法：

- ✅ **减少样板代码**：消除重复的 `if err != nil` 检查
- ✅ **保持安全性**：自动将 panic 转换回 error
- ✅ **类型安全**：完整支持泛型，适用于任何返回值类型组合
- ✅ **上下文感知**：内置支持 Go 上下文
- ✅ **钩子友好**：可自定义的错误处理回调
- ✅ **零依赖**：纯 Go 标准库实现

## 📦 安装

```bash
go get github.com/yankeguo/rg
```

## 🎯 快速开始

转换冗长的错误处理：

```go
// 之前：传统的 Go 错误处理
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

// 之后：使用 rg
func processFile(filename string) (result []byte, err error) {
    defer rg.Guard(&err)

    data := rg.Must(os.ReadFile(filename))
    result = rg.Must(processData(data))
    return
}
```

## 📖 核心概念

### Guard - 安全网

`rg.Guard(&err)` 充当安全网，捕获任何 panic 并将其转换为错误：

```go
func riskyOperation() (err error) {
    defer rg.Guard(&err)

    // 这里的任何 panic 都会被捕获并转换为 err
    rg.Must0(someFunctionThatMightFail())
    return nil // 成功情况
}
```

### Must 函数系列 - 抛出器

`Must` 函数系列检查错误，如果发现错误就会 panic：

- `rg.Must0(err)` - 用于只返回错误的函数
- `rg.Must(value, err)` - 用于返回一个值 + 错误的函数
- `rg.Must2(v1, v2, err)` - 用于返回两个值 + 错误的函数
- ... 最多支持到 `rg.Must7`，用于七个值 + 错误

## 🔧 高级特性

### 上下文支持

通过错误处理链传递上下文信息：

```go
func processWithContext(ctx context.Context) (err error) {
    defer rg.Guard(&err, rg.WithContext(ctx))

    // 上下文在错误回调中可用
    result := rg.Must(someNetworkCall(ctx))
    return nil
}
```

### 错误钩子和回调

使用全局钩子自定义错误处理：

```go
func init() {
    // 全局错误钩子（已弃用，请使用 OnGuardWithContext）
    rg.OnGuard = func(r any) {
        log.Printf("捕获到错误: %v", r)
    }

    // 上下文感知的错误钩子
    rg.OnGuardWithContext = func(ctx context.Context, r any) {
        // 从上下文中提取请求ID、用户信息等
        if reqID := ctx.Value("request_id"); reqID != nil {
            log.Printf("请求 %v 中的错误: %v", reqID, r)
        }
    }
}
```

## 💡 实际应用示例

### 文件处理管道

```go
func convertJSONToYAML(inputFile string) (err error) {
    defer rg.Guard(&err)

    // 读取并解析 JSON
    jsonData := rg.Must(os.ReadFile(inputFile))
    var data map[string]interface{}
    rg.Must0(json.Unmarshal(jsonData, &data))

    // 转换为 YAML 并写入
    yamlData := rg.Must(yaml.Marshal(data))
    rg.Must0(os.WriteFile(inputFile+".yaml", yamlData, 0644))

    return nil
}
```

### HTTP API 处理器

```go
func handleUserCreation(w http.ResponseWriter, r *http.Request) {
    var err error
    defer rg.Guard(&err, rg.WithContext(r.Context()))
    defer func() {
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }()

    // 解析请求
    var user User
    rg.Must0(json.NewDecoder(r.Body).Decode(&user))

    // 验证并保存
    rg.Must0(user.Validate())
    savedUser := rg.Must(userService.Create(user))

    // 返回响应
    w.Header().Set("Content-Type", "application/json")
    rg.Must0(json.NewEncoder(w).Encode(savedUser))
}
```

### 数据库事务

```go
func transferMoney(from, to int64, amount decimal.Decimal) (err error) {
    defer rg.Guard(&err)

    tx := rg.Must(db.Begin())
    defer tx.Rollback() // 即使在提交后调用也是安全的

    // 执行转账操作
    rg.Must0(debitAccount(tx, from, amount))
    rg.Must0(creditAccount(tx, to, amount))
    rg.Must0(logTransfer(tx, from, to, amount))

    rg.Must0(tx.Commit())
    return nil
}
```

## 🎨 最佳实践

1. **总是使用 `defer rg.Guard(&err)`** 在需要错误处理的函数开头
2. **保持守卫简单**：不要在 defer 语句中放置复杂逻辑
3. **使用上下文**：传递上下文以获得更好的错误跟踪和调试
4. **与传统错误处理结合**：`rg` 与标准 Go 错误处理能很好地协同工作
5. **彻底测试**：确保错误路径被测试覆盖

## 🤔 何时使用 rg

**适用场景：**

- 数据处理管道
- 有多个验证步骤的 API 处理器
- 文件 I/O 操作
- 数据库事务
- 任何有多个可能失败的顺序操作的场景

**考虑替代方案的场景：**

- 只有一两个错误检查的简单函数
- 性能关键代码（panic/recover 有开销）
- 需要暴露传统 Go API 的库

## 🛠 对比

| 特性     | 传统 Go                  | rg                   |
| -------- | ------------------------ | -------------------- |
| 错误处理 | 显式 `if err != nil`     | 使用 `Must` 自动处理 |
| 代码长度 | 较长                     | 较短                 |
| 性能     | 更快（无 panic/recover） | 略慢                 |
| 可读性   | 简单情况下良好           | 复杂情况下优秀       |
| 调试     | 标准堆栈跟踪             | 通过钩子增强         |

## 📚 API 参考

### 核心函数

- `Guard(err *error, opts ...Option)` - 从 panic 中恢复并设置错误
- `Must0(err error)` - 如果错误不为空就 panic
- `Must[T](value T, err error) T` - 返回值或在错误时 panic
- `Must2` 到 `Must7` - 处理多个返回值

### 选项

- `WithContext(ctx context.Context)` - 为守卫附加上下文

### 钩子

- `OnGuard func(r any)` - 全局 panic 钩子（已弃用）
- `OnGuardWithContext func(ctx context.Context, r any)` - 上下文感知的 panic 钩子

## 🤝 贡献

我们欢迎贡献！请随时提交问题、功能请求或拉取请求。

## 📄 许可证

MIT 许可证 - 详见 LICENSE 文件。

## 👨‍💻 作者

**GUO YANKE** - [@yankeguo](https://github.com/yankeguo)

---

⭐ 如果你觉得这个库有帮助，请考虑给它一个星标！
