# FreezeLib - 美观代码截图的 Go 库

**Language / 语言**: [English](README.md) | [中文](README_CN.md)

**Documentation / 文档**: [Usage Guide (English)](USAGE_EN.md) | [使用指南 (中文)](USAGE.md)

FreezeLib 是一个用于生成美观代码和终端输出截图的 Go 库。它基于 Charm 团队广受欢迎的 [freeze](https://github.com/charmbracelet/freeze) CLI 工具，但重新设计为可在 Go 应用程序中重复使用的库。

## 特性

- 🎨 **语法高亮**: 支持 100+ 种编程语言
- 🖼️ **多种输出格式**: 生成 SVG 和 PNG 图像
- 🎭 **丰富主题**: 内置主题包括 GitHub、Dracula、Monokai 等
- 🪟 **窗口控件**: macOS 风格的窗口装饰
- 📏 **行号**: 可选的行号显示
- 🌈 **ANSI 支持**: 渲染彩色终端输出
- ⚡ **简易 API**: 简单且可链式调用的 API 设计
- 🎯 **预设配置**: 常见用例的预配置样式
- 🔧 **高度可定制**: 精细调整输出的每个方面

## 安装

```bash
go get github.com/landaiqing/freezelib
```

## 快速开始

### 基本用法

```go
package main

import (
    "os"
    "github.com/landaiqing/freezelib"
)

func main() {
    // 创建新的 freeze 实例
    freeze := freezelib.New()

    // 要截图的 Go 代码
    code := `package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}`

    // 生成 SVG
    svgData, err := freeze.GenerateFromCode(code, "go")
    if err != nil {
        panic(err)
    }

    // 保存到文件
    os.WriteFile("hello.svg", svgData, 0644)
}
```

### QuickFreeze API

为了更流畅的体验，使用 QuickFreeze API：

```go
qf := freezelib.NewQuickFreeze()

svgData, err := qf.WithTheme("dracula").
    WithFont("Fira Code", 14).
    WithWindow().
    WithShadow().
    WithLineNumbers().
    CodeToSVG(code)
```

## API 参考

### 核心类型

#### Freeze

生成截图的主要接口：

```go
freeze := freezelib.New()                    // 默认配置
freeze := freezelib.NewWithConfig(config)    // 自定义配置
freeze := freezelib.NewWithPreset("dark")    // 预设配置
```

#### QuickFreeze

简化的链式 API：

```go
qf := freezelib.NewQuickFreeze()
qf := freezelib.NewQuickFreezeWithPreset("terminal")
```

### 生成方法

#### 从代码字符串
```go
svgData, err := freeze.GenerateFromCode(code, "python")
pngData, err := freeze.GeneratePNGFromCode(code, "python")
```

#### 从文件
```go
svgData, err := freeze.GenerateFromFile("main.go")
pngData, err := freeze.GeneratePNGFromFile("main.go")
```

#### 从 ANSI 终端输出
```go
terminalOutput := "\033[32mSUCCESS\033[0m: Build completed"
svgData, err := freeze.GenerateFromANSI(terminalOutput)
pngData, err := freeze.GeneratePNGFromANSI(terminalOutput)
```

#### 从 Reader
```go
svgData, err := freeze.GenerateFromReader(reader, "javascript")
```

### 配置

#### 基本配置
```go
config := freezelib.DefaultConfig()
config.SetTheme("github-dark")
config.SetFont("JetBrains Mono", 14)
config.SetBackground("#1e1e1e")
config.SetWindow(true)
config.SetLineNumbers(true)
```

#### 高级配置
```go
config.SetPadding(20)           // 所有边
config.SetPadding(20, 40)       // 垂直，水平
config.SetPadding(20, 40, 20, 40) // 上，右，下，左

config.SetShadow(20, 0, 10)     // 模糊，X 偏移，Y 偏移
config.SetBorder(1, 8, "#333")  // 宽度，圆角，颜色
config.SetDimensions(800, 600)  // 宽度，高度
config.SetLines(10, 20)         // 行范围（1-indexed）
```

### 预设

FreezeLib 提供了几个内置预设：

```go
// 可用预设
presets := []string{
    "base",         // 简洁干净
    "full",         // macOS 风格窗口控件
    "terminal",     // 终端输出优化
    "presentation", // 演示高对比度
    "minimal",      // 极简样式
    "dark",         // 深色主题
    "light",        // 浅色主题
    "retro",        // 复古终端风格
    "neon",         // 霓虹/赛博朋克风格
    "compact",      // 小代码片段紧凑型
}

freeze := freezelib.NewWithPreset("dark")
```

### 链式方法

`Freeze` 和 `QuickFreeze` 都支持方法链：

```go
freeze := freezelib.New().
    WithTheme("monokai").
    WithFont("Cascadia Code", 15).
    WithWindow(true).
    WithShadow(20, 0, 10).
    WithLineNumbers(true)

svgData, err := freeze.GenerateFromCode(code, "rust")
```

## 示例

### 终端输出截图

```go
freeze := freezelib.NewWithPreset("terminal")

ansiOutput := "\033[32m✓ Tests passed\033[0m\n" +
              "\033[31m✗ Build failed\033[0m\n" +
              "\033[33m⚠ Warning: deprecated API\033[0m"

svgData, err := freeze.GenerateFromANSI(ansiOutput)
```

### 自定义样式

```go
config := freezelib.DefaultConfig()
config.Theme = "github"
config.Background = "#f6f8fa"
config.Font.Family = "SF Mono"
config.Font.Size = 16
config.SetPadding(30)
config.SetMargin(20)
config.Window = true
config.ShowLineNumbers = true
config.Border.Radius = 12
config.Shadow.Blur = 25

freeze := freezelib.NewWithConfig(config)
```

### 批量处理

```go
files := []string{"main.go", "config.go", "utils.go"}

for _, file := range files {
    svgData, err := freeze.GenerateFromFile(file)
    if err != nil {
        continue
    }
    
    outputFile := strings.TrimSuffix(file, ".go") + ".svg"
    os.WriteFile(outputFile, svgData, 0644)
}
```

## 支持的语言

FreezeLib 支持 100+ 种编程语言的语法高亮，包括：

- Go, Rust, Python, JavaScript, TypeScript
- C, C++, C#, Java, Kotlin, Swift
- HTML, CSS, SCSS, JSON, YAML, XML
- Shell, PowerShell, Dockerfile
- SQL, GraphQL, Markdown
- 等等...

## 支持的主题

流行主题包括：
- `github` / `github-dark`
- `dracula`
- `monokai`
- `solarized-dark` / `solarized-light`
- `nord`
- `one-dark`
- `material`
- `vim`
- 等等...

## 错误处理

```go
svgData, err := freeze.GenerateFromCode(code, "go")
if err != nil {
    // 处理特定错误
    switch {
    case strings.Contains(err.Error(), "language"):
        // 语言检测失败
    case strings.Contains(err.Error(), "config"):
        // 配置错误
    default:
        // 其他错误
    }
}
```

## 贡献

欢迎贡献！请随时提交 Pull Request。

## 许可证

MIT 许可证 - 详见 [LICENSE](./LICENSE) 文件。

## 致谢

本库基于 [Charm](https://charm.sh) 团队出色的 [freeze](https://github.com/charmbracelet/freeze) CLI 工具。特别感谢 Charm 团队创造了如此美观的工具。
