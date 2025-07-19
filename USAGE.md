# FreezeLib 使用指南

**Language / 语言**: [English](USAGE_EN.md) | [中文](USAGE.md)

FreezeLib 是一个 Go 库，用于生成美观的代码截图。

## 安装

```bash
go get github.com/landaiqing/freezelib
```

## 基本使用

### 1. 最简单的例子

```go
package main

import (
	"github.com/landaiqing/freezelib"
	"os"
)

func main() {
	// 创建实例
	freeze := freezelib.New()

	// 代码内容
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

	// 保存文件
	os.WriteFile("hello.svg", svgData, 0644)
}
```

### 2. 从文件生成

```go
// 直接从文件生成
svgData, err := freeze.GenerateFromFile("main.go")
if err != nil {
    panic(err)
}
os.WriteFile("main.svg", svgData, 0644)
```

### 3. 生成 PNG 格式

```go
// 生成 PNG 而不是 SVG
pngData, err := freeze.GeneratePNGFromCode(code, "go")
if err != nil {
    panic(err)
}
os.WriteFile("hello.png", pngData, 0644)
```

### 4. 直接保存到文件

```go
// 一步完成：生成并保存
err := freeze.SaveCodeToFile(code, "go", "output.svg")
if err != nil {
    panic(err)
}

// 自动检测格式（根据文件扩展名）
err = freeze.SaveCodeToFile(code, "go", "output.png")
```

## 使用预设样式

```go
// 使用预设配置，快速开始
freeze := freezelib.NewWithPreset("dark")        // 深色主题
freeze := freezelib.NewWithPreset("terminal")    // 终端风格
freeze := freezelib.NewWithPreset("presentation") // 演示风格

// 查看所有可用预设
presets := freeze.GetAvailablePresets()
// 返回: ["base", "compact", "dark", "full", "light", "minimal", "neon", "presentation", "retro", "terminal"]
```

## 自定义样式

### 基本设置

```go
freeze := freezelib.New().
    WithTheme("github-dark").           // 设置主题
    WithFont("JetBrains Mono", 14).     // 设置字体和大小
    WithBackground("#1e1e1e").          // 设置背景色
    WithPadding(20)                     // 设置内边距

svgData, err := freeze.GenerateFromCode(code, "go")
```

### 添加装饰效果

```go
freeze := freezelib.New().
    WithTheme("dracula").
    WithWindow(true).                   // 添加窗口控件
    WithLineNumbers(true).              // 显示行号
    WithShadow(20, 0, 10)               // 添加阴影

svgData, err := freeze.GenerateFromCode(code, "python")
```

## 查看支持的选项

```go
freeze := freezelib.New()

// 查看所有支持的语言（270+ 种）
languages := freeze.GetSupportedLanguages()
fmt.Printf("支持 %d 种语言\n", len(languages))

// 查看所有支持的主题（67+ 种）
themes := freeze.GetSupportedThemes()
fmt.Printf("支持 %d 种主题\n", len(themes))

// 查看所有可用预设（10 种）
presets := freeze.GetAvailablePresets()
fmt.Printf("可用预设: %v\n", presets)
```

## 终端输出截图

```go
// 截图终端输出（支持 ANSI 颜色）
freeze := freezelib.NewWithPreset("terminal")

ansiOutput := "\033[32m✓ 测试通过\033[0m\n" +
              "\033[31m✗ 构建失败\033[0m\n" +
              "\033[33m⚠ 警告: API 已废弃\033[0m"

svgData, err := freeze.GenerateFromANSI(ansiOutput)
if err != nil {
    panic(err)
}
os.WriteFile("terminal.svg", svgData, 0644)
```

## 批量处理文件

```go
freeze := freezelib.NewWithPreset("dark")
files := []string{"main.go", "config.go", "utils.go"}

for _, file := range files {
    err := freeze.SaveFileToFile(file, file+".svg")
    if err != nil {
        fmt.Printf("处理 %s 失败: %v\n", file, err)
        continue
    }
    fmt.Printf("已生成: %s.svg\n", file)
}
```

## 自动语言检测

```go
freeze := freezelib.New()

// 不指定语言，自动检测
code := `function hello() {
    console.log("Hello, World!");
}`

// 自动检测为 JavaScript
svgData, err := freeze.GenerateFromCodeAuto(code)
if err != nil {
    panic(err)
}
os.WriteFile("auto.svg", svgData, 0644)
```

## 错误处理

```go
svgData, err := freeze.GenerateFromCode(code, "go")
if err != nil {
    fmt.Printf("生成失败: %v\n", err)
    return
}

// 保存文件
err = os.WriteFile("output.svg", svgData, 0644)
if err != nil {
    fmt.Printf("保存失败: %v\n", err)
    return
}

fmt.Println("生成成功!")
```

## 常用主题

- `github` - GitHub 浅色主题
- `github-dark` - GitHub 深色主题
- `dracula` - Dracula 主题
- `monokai` - Monokai 主题
- `nord` - Nord 主题

## 常用语言

- `go` - Go 语言
- `python` - Python
- `javascript` - JavaScript
- `typescript` - TypeScript
- `java` - Java
- `rust` - Rust
- `c` - C 语言
- `cpp` - C++
- `html` - HTML
- `css` - CSS
- `json` - JSON
- `yaml` - YAML
- `markdown` - Markdown
- `bash` - Shell 脚本
