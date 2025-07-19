# FreezeLib - 美观代码截图的 Go 库

**Language / 语言**: [English](README.md) | [中文](README_CN.md)

**Documentation / 文档**: [Usage Guide (English)](USAGE_EN.md) | [使用指南 (中文)](USAGE.md)

FreezeLib 是一个用于生成美观代码和终端输出截图的 Go 库。它基于 Charm 团队广受欢迎的 [freeze](https://github.com/charmbracelet/freeze) CLI 工具，但重新设计为可在 Go 应用程序中重复使用的库。

## 特性

- 🎨 **语法高亮**: 支持 270+ 种编程语言
- 🔍 **自动语言检测**: 智能检测代码内容和文件名的语言
- 📋 **简单列表**: 轻松访问所有可用的语言、主题和预设
- 🖼️ **多种输出格式**: 生成 SVG 和 PNG 图像
- 🎭 **丰富主题**: 67+ 内置主题，包括 GitHub、Dracula、Monokai 等
- 🪟 **窗口控件**: macOS 风格的窗口装饰
- 📏 **行号**: 可选的行号显示
- 🌈 **ANSI 支持**: 渲染彩色终端输出
- ⚡ **简易 API**: 简单且可链式调用的 API 设计
- 🎯 **预设配置**: 10 种预配置样式，适用于常见用例
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

### 自动语言检测

FreezeLib 可以自动检测编程语言：

```go
freeze := freezelib.New()

// 从代码内容自动检测语言
svgData, err := freeze.GenerateFromCodeAuto(code)

// 手动检测语言
language := freeze.DetectLanguage(code)
fmt.Printf("检测到的语言: %s", language)

// 从文件名检测语言
language = freeze.DetectLanguageFromFilename("script.py")

// 组合检测（文件名 + 内容）
language = freeze.DetectLanguageFromFile("script.py", code)
```

### 可用选项

获取所有可用选项的列表：

```go
freeze := freezelib.New()

// 获取所有可用选项（排序列表）
languages := freeze.GetSupportedLanguages()  // 270+ 种语言
themes := freeze.GetSupportedThemes()        // 67+ 种主题
presets := freeze.GetAvailablePresets()      // 10 种预设

// 测试支持
isSupported := freeze.IsLanguageSupported("go")
isSupported = freeze.IsThemeSupported("github")
isValid := freezelib.IsValidPreset("dark")

// 全局函数也可用
languages = freezelib.GetSupportedLanguages()
themes = freezelib.GetSupportedThemes()
presets = freezelib.GetAvailablePresets()
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
    CodeToSVGAuto(code) // 自动检测语言
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

// 使用自动语言检测
svgData, err := freeze.GenerateFromCodeAuto(code)
pngData, err := freeze.GeneratePNGFromCodeAuto(code)
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
// 可用预设（10 种）
presets := []string{
    "base",         // 基础配置
    "full",         // 完整功能配置
    "terminal",     // 终端风格
    "dark",         // 深色主题
    "light",        // 浅色主题
    "minimal",      // 极简风格
    "professional", // 专业风格
    "vibrant",      // 鲜艳配色
    "retro",        // 复古风格
    "neon",         // 霓虹风格
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

查看 `examples/` 目录获取完整示例：

- `01-basic/` - 基础用法示例
- `02-formats/` - 输出格式示例
- `03-themes/` - 主题展示
- `04-languages/` - 语言支持示例
- `05-terminal/` - 终端输出示例
- `06-advanced/` - 高级配置
- `07-batch/` - 批量处理
- `08-auto-language-detection/` - 自动语言检测
- `09-supported-options/` - 支持选项列表

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

FreezeLib 支持 270+ 种编程语言的语法高亮：

### 热门语言
- **系统级**: Go, Rust, C, C++, Zig, D, Nim, V
- **Web**: JavaScript, TypeScript, HTML, CSS, SCSS, PHP
- **企业级**: Java, C#, Kotlin, Scala, Swift
- **脚本**: Python, Ruby, Perl, Lua, Bash, PowerShell
- **数据**: JSON, YAML, TOML, XML, SQL, GraphQL

### 分类
- **热门** (30): 最常用的语言
- **Web** (15): 前端和后端 Web 技术
- **系统** (13): 底层和系统编程
- **脚本** (12): 自动化和脚本语言
- **数据** (11): 配置和数据格式
- **更多**: 总共支持 270+ 种语言

## 支持的主题

FreezeLib 包含 67+ 种语法高亮主题：

### 热门主题
- **GitHub**: `github`, `github-dark`
- **现代**: `dracula`, `monokai`, `nord`, `one-dark`
- **经典**: `solarized-dark`, `solarized-light`, `material`, `vim`
- **多彩**: `colorful`, `friendly`, `fruity`, `rainbow_dash`

### 分类
- **热门** (30): 最常用的主题
- **深色** (10): 适合低光环境的深色配色方案
- **浅色** (14): 适合明亮环境的浅色配色方案
- **更多**: 总共 67+ 种主题可用

### 简单高效
所有列表都按字母顺序排序，便于浏览和选择。

## 语言检测功能

### 自动检测方法
```go
freeze := freezelib.New()

// 从代码内容检测语言
language := freeze.DetectLanguage(code)

// 从文件名检测
language = freeze.DetectLanguageFromFilename("script.py")

// 组合检测（文件名 + 内容）
language = freeze.DetectLanguageFromFile("script.py", code)

// 检查语言支持
supported := freeze.IsLanguageSupported("go")

// 获取所有支持的语言
languages := freeze.GetSupportedLanguages()
```

### 自定义语言检测器
```go
detector := freeze.GetLanguageDetector()

// 添加自定义文件扩展名映射
detector.AddCustomMapping(".myext", "python")
detector.AddCustomMapping(".config", "json")

// 配置检测策略
detector.EnableContentAnalysis = true
detector.EnableFilenameAnalysis = true
detector.FallbackLanguage = "text"
```

## 支持选项

### 获取所有可用选项
```go
freeze := freezelib.New()

// 获取所有可用选项（排序列表）
languages := freeze.GetSupportedLanguages()  // 270+ 种语言
themes := freeze.GetSupportedThemes()        // 67+ 种主题
presets := freeze.GetAvailablePresets()      // 10 种预设

// 测试支持
isSupported := freeze.IsLanguageSupported("go")
isSupported = freeze.IsThemeSupported("github")
isValid := freezelib.IsValidPreset("dark")

// 全局函数也可用
languages = freezelib.GetSupportedLanguages()
themes = freezelib.GetSupportedThemes()
presets = freezelib.GetAvailablePresets()
```

### 验证支持
```go
// 检查支持
if freeze.IsLanguageSupported("go") {
    // 生成 Go 代码截图
}

if freeze.IsThemeSupported("dracula") {
    // 使用 Dracula 主题
}

if freezelib.IsValidPreset("dark") {
    // 使用深色预设
}
```

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
