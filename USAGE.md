# FreezeLib 使用指南

**Language / 语言**: [English](USAGE_EN.md) | [中文](USAGE.md)

**Main Documentation / 主要文档**: [README (English)](README.md) | [README (中文)](README_CN.md)

FreezeLib 是一个基于 Charm 的 freeze CLI 工具重构的 Go 公共库，用于生成美观的代码截图。

## 🚀 快速开始

### 基本用法

```go
package main

import (
	"github.com/landaiqing/freezelib"
	"os"
)

func main() {
	// 创建 freeze 实例
	freeze := freezelib.New()

	// 要截图的代码
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

### 链式调用 API

```go
// 使用 QuickFreeze 进行链式调用
qf := freezelib.NewQuickFreeze()

svgData, err := qf.WithTheme("dracula").
    WithFont("Fira Code", 14).
    WithWindow().
    WithShadow().
    WithLineNumbers().
    WithLanguage("javascript").
    CodeToSVG(code)
```

## 📋 主要功能

### 1. 多种输入方式

```go
// 从代码字符串生成
svgData, err := freeze.GenerateFromCode(code, "python")

// 从文件生成
svgData, err := freeze.GenerateFromFile("main.go")

// 从 ANSI 终端输出生成
ansiOutput := "\033[32m✓ SUCCESS\033[0m: Build completed"
svgData, err := freeze.GenerateFromANSI(ansiOutput)

// 从 Reader 生成
svgData, err := freeze.GenerateFromReader(reader, "javascript")
```

### 2. 多种输出格式

```go
// 生成 SVG
svgData, err := freeze.GenerateFromCode(code, "go")

// 生成 PNG
pngData, err := freeze.GeneratePNGFromCode(code, "go")

// 直接保存到文件
err := freeze.SaveCodeToFile(code, "go", "output.svg")
err := freeze.SaveCodeToFile(code, "go", "output.png") // 自动检测格式
```

### 3. 预设配置

```go
// 使用预设配置
freeze := freezelib.NewWithPreset("dark")        // 深色主题
freeze := freezelib.NewWithPreset("terminal")    // 终端风格
freeze := freezelib.NewWithPreset("presentation") // 演示风格

// 可用预设
presets := []string{
    "base",         // 基础样式
    "full",         // macOS 风格
    "terminal",     // 终端优化
    "presentation", // 演示优化
    "minimal",      // 极简风格
    "dark",         // 深色主题
    "light",        // 浅色主题
    "retro",        // 复古风格
    "neon",         // 霓虹风格
    "compact",      // 紧凑风格
}
```

### 4. 自定义配置

```go
config := freezelib.DefaultConfig()

// 基本设置
config.SetTheme("github-dark")
config.SetFont("JetBrains Mono", 14)
config.SetBackground("#1e1e1e")
config.SetLanguage("python")

// 布局设置
config.SetPadding(20)           // 所有边
config.SetPadding(20, 40)       // 垂直，水平
config.SetPadding(20, 40, 20, 40) // 上，右，下，左
config.SetMargin(15)
config.SetDimensions(800, 600)

// 装饰效果
config.SetWindow(true)          // 窗口控件
config.SetLineNumbers(true)     // 行号
config.SetShadow(20, 0, 10)     // 阴影：模糊，X偏移，Y偏移
config.SetBorder(1, 8, "#333")  // 边框：宽度，圆角，颜色

// 行范围（1-indexed）
config.SetLines(10, 20)         // 只截取第10-20行

freeze := freezelib.NewWithConfig(config)
```

## 🎨 支持的主题

- `github` / `github-dark`
- `dracula`
- `monokai`
- `solarized-dark` / `solarized-light`
- `nord`
- `one-dark`
- `material`
- `vim`
- 等等...

## 💻 支持的语言

支持 100+ 种编程语言，包括：
- Go, Rust, Python, JavaScript, TypeScript
- C, C++, C#, Java, Kotlin, Swift
- HTML, CSS, SCSS, JSON, YAML, XML
- Shell, PowerShell, Dockerfile
- SQL, GraphQL, Markdown
- 等等...

## 🔧 高级用法

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

### 终端输出截图

```go
freeze := freezelib.NewWithPreset("terminal")

ansiOutput := "\033[32m✓ Tests passed\033[0m\n" +
              "\033[31m✗ Build failed\033[0m\n" +
              "\033[33m⚠ Warning: deprecated API\033[0m"

svgData, err := freeze.GenerateFromANSI(ansiOutput)
```

### 链式方法

```go
freeze := freezelib.New().
    WithTheme("monokai").
    WithFont("Cascadia Code", 15).
    WithWindow(true).
    WithShadow(20, 0, 10).
    WithLineNumbers(true)

svgData, err := freeze.GenerateFromCode(code, "rust")
```

## 📊 性能优化建议

1. **重用实例**：创建一个 `Freeze` 实例并重复使用
2. **选择合适格式**：网页用 SVG，演示用 PNG
3. **设置具体尺寸**：指定尺寸可提高性能
4. **批量操作**：在单个会话中处理多个文件

## 🐛 错误处理

```go
svgData, err := freeze.GenerateFromCode(code, "go")
if err != nil {
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

## 📁 项目结构

```
freezelib/
├── freeze.go          # 主要 API 接口
├── config.go          # 配置结构体
├── generator.go       # 核心生成逻辑
├── quickfreeze.go     # 简化 API
├── presets.go         # 预设配置
├── ansi.go           # ANSI 处理
├── svg/              # SVG 处理
├── font/             # 字体处理
├── example/          # 使用示例
└── README.md         # 详细文档
```

## 🤝 与原版 freeze 的区别

| 特性 | 原版 freeze | FreezeLib |
|------|-------------|-----------|
| 使用方式 | CLI 工具 | Go 库 |
| 集成方式 | 命令行调用 | 直接导入 |
| 配置方式 | 命令行参数/配置文件 | Go 结构体 |
| 扩展性 | 有限 | 高度可扩展 |
| 性能 | 进程启动开销 | 内存中处理 |

## 📝 示例代码

查看 `examples` 目录中的完整示例：

这将生成多个示例 SVG 文件，展示库的各种功能。
