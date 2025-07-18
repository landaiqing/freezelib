# FreezeLib - Go Library for Beautiful Code Screenshots

**Language / 语言**: [English](README.md) | [中文](README_CN.md)

**Documentation / 文档**: [Usage Guide (English)](USAGE_EN.md) | [使用指南 (中文)](USAGE.md)

FreezeLib is a Go library for generating beautiful screenshots of code and terminal output. It's based on the popular [freeze](https://github.com/charmbracelet/freeze) CLI tool by Charm, but redesigned as a reusable library for Go applications.

## Features

- 🎨 **Syntax Highlighting**: Support for 100+ programming languages
- 🖼️ **Multiple Output Formats**: Generate SVG and PNG images
- 🎭 **Rich Themes**: Built-in themes including GitHub, Dracula, Monokai, and more
- 🪟 **Window Controls**: macOS-style window decorations
- 📏 **Line Numbers**: Optional line numbering
- 🌈 **ANSI Support**: Render colored terminal output
- ⚡ **Easy API**: Simple and chainable API design
- 🎯 **Presets**: Pre-configured styles for common use cases
- 🔧 **Highly Customizable**: Fine-tune every aspect of the output

## Installation

```bash
go get github.com/landaiqing/freezelib
```

## Quick Start

### Basic Usage

```go
package main

import (
    "os"
    "github.com/landaiqing/freezelib"
)

func main() {
    // Create a new freeze instance
    freeze := freezelib.New()

    // Go code to screenshot
    code := `package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}`

    // Generate SVG
    svgData, err := freeze.GenerateFromCode(code, "go")
    if err != nil {
        panic(err)
    }

    // Save to file
    os.WriteFile("hello.svg", svgData, 0644)
}
```

### QuickFreeze API

For a more fluent experience, use the QuickFreeze API:

```go
qf := freezelib.NewQuickFreeze()

svgData, err := qf.WithTheme("dracula").
    WithFont("Fira Code", 14).
    WithWindow().
    WithShadow().
    WithLineNumbers().
    CodeToSVG(code)
```

## API Reference

### Core Types

#### Freeze

The main interface for generating screenshots:

```go
freeze := freezelib.New()                    // Default config
freeze := freezelib.NewWithConfig(config)    // Custom config
freeze := freezelib.NewWithPreset("dark")    // Preset config
```

#### QuickFreeze

Simplified, chainable API:

```go
qf := freezelib.NewQuickFreeze()
qf := freezelib.NewQuickFreezeWithPreset("terminal")
```

### Generation Methods

#### From Code String
```go
svgData, err := freeze.GenerateFromCode(code, "python")
pngData, err := freeze.GeneratePNGFromCode(code, "python")
```

#### From File
```go
svgData, err := freeze.GenerateFromFile("main.go")
pngData, err := freeze.GeneratePNGFromFile("main.go")
```

#### From ANSI Terminal Output
```go
terminalOutput := "\033[32mSUCCESS\033[0m: Build completed"
svgData, err := freeze.GenerateFromANSI(terminalOutput)
pngData, err := freeze.GeneratePNGFromANSI(terminalOutput)
```

#### From Reader
```go
svgData, err := freeze.GenerateFromReader(reader, "javascript")
```

### Configuration

#### Basic Configuration
```go
config := freezelib.DefaultConfig()
config.SetTheme("github-dark")
config.SetFont("JetBrains Mono", 14)
config.SetBackground("#1e1e1e")
config.SetWindow(true)
config.SetLineNumbers(true)
```

#### Advanced Configuration
```go
config.SetPadding(20)           // All sides
config.SetPadding(20, 40)       // Vertical, horizontal
config.SetPadding(20, 40, 20, 40) // Top, right, bottom, left

config.SetShadow(20, 0, 10)     // Blur, X offset, Y offset
config.SetBorder(1, 8, "#333")  // Width, radius, color
config.SetDimensions(800, 600)  // Width, height
config.SetLines(10, 20)         // Line range (1-indexed)
```

### Presets

FreezeLib comes with several built-in presets:

```go
// Available presets
presets := []string{
    "base",         // Simple, clean
    "full",         // macOS-style with window controls
    "terminal",     // Optimized for terminal output
    "presentation", // High contrast for presentations
    "minimal",      // Minimal styling
    "dark",         // Dark theme
    "light",        // Light theme
    "retro",        // Retro terminal style
    "neon",         // Neon/cyberpunk style
    "compact",      // Compact for small snippets
}

freeze := freezelib.NewWithPreset("dark")
```

### Chainable Methods

Both `Freeze` and `QuickFreeze` support method chaining:

```go
freeze := freezelib.New().
    WithTheme("monokai").
    WithFont("Cascadia Code", 15).
    WithWindow(true).
    WithShadow(20, 0, 10).
    WithLineNumbers(true)

svgData, err := freeze.GenerateFromCode(code, "rust")
```

## Examples

### Terminal Output Screenshot

```go
freeze := freezelib.NewWithPreset("terminal")

ansiOutput := "\033[32m✓ Tests passed\033[0m\n" +
              "\033[31m✗ Build failed\033[0m\n" +
              "\033[33m⚠ Warning: deprecated API\033[0m"

svgData, err := freeze.GenerateFromANSI(ansiOutput)
```

### Custom Styling

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

### Batch Processing

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

## Supported Languages

FreezeLib supports syntax highlighting for 100+ programming languages including:

- Go, Rust, Python, JavaScript, TypeScript
- C, C++, C#, Java, Kotlin, Swift
- HTML, CSS, SCSS, JSON, YAML, XML
- Shell, PowerShell, Dockerfile
- SQL, GraphQL, Markdown
- And many more...

## Supported Themes

Popular themes include:
- `github` / `github-dark`
- `dracula`
- `monokai`
- `solarized-dark` / `solarized-light`
- `nord`
- `one-dark`
- `material`
- `vim`
- And many more...

## Error Handling

```go
svgData, err := freeze.GenerateFromCode(code, "go")
if err != nil {
    // Handle specific errors
    switch {
    case strings.Contains(err.Error(), "language"):
        // Language detection failed
    case strings.Contains(err.Error(), "config"):
        // Configuration error
    default:
        // Other errors
    }
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see the [LICENSE](./LICENSE) file for details.

## Acknowledgments

This library is based on the excellent [freeze](https://github.com/charmbracelet/freeze) CLI tool by [Charm](https://charm.sh). Special thanks to the Charm team for creating such a beautiful tool.
