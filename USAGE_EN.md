# FreezeLib Usage Guide

**Language / 语言**: [English](USAGE_EN.md) | [中文](USAGE.md)

**Main Documentation / 主要文档**: [README (English)](README.md) | [README (中文)](README_CN.md)

FreezeLib is a Go library refactored from Charm's freeze CLI tool for generating beautiful code screenshots.

## 🚀 Quick Start

### Basic Usage

```go
package main

import (
	"github.com/landaiqing/freezelib"
	"os"
)

func main() {
	// Create freeze instance
	freeze := freezelib.New()

	// Code to screenshot
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

### Chainable API

```go
// Use QuickFreeze for method chaining
qf := freezelib.NewQuickFreeze()

svgData, err := qf.WithTheme("dracula").
    WithFont("Fira Code", 14).
    WithWindow().
    WithShadow().
    WithLineNumbers().
    WithLanguage("javascript").
    CodeToSVG(code)
```

## 📋 Main Features

### 1. Multiple Input Methods

```go
// Generate from code string
svgData, err := freeze.GenerateFromCode(code, "python")

// Generate from file
svgData, err := freeze.GenerateFromFile("main.go")

// Generate from ANSI terminal output
ansiOutput := "\033[32m✓ SUCCESS\033[0m: Build completed"
svgData, err := freeze.GenerateFromANSI(ansiOutput)

// Generate from Reader
svgData, err := freeze.GenerateFromReader(reader, "javascript")
```

### 2. Multiple Output Formats

```go
// Generate SVG
svgData, err := freeze.GenerateFromCode(code, "go")

// Generate PNG
pngData, err := freeze.GeneratePNGFromCode(code, "go")

// Save directly to file
err := freeze.SaveCodeToFile(code, "go", "output.svg")
err := freeze.SaveCodeToFile(code, "go", "output.png") // Auto-detect format
```

### 3. Preset Configurations

```go
// Use preset configurations
freeze := freezelib.NewWithPreset("dark")        // Dark theme
freeze := freezelib.NewWithPreset("terminal")    // Terminal style
freeze := freezelib.NewWithPreset("presentation") // Presentation style

// Available presets
presets := []string{
    "base",         // Basic style
    "full",         // macOS style
    "terminal",     // Terminal optimized
    "presentation", // Presentation optimized
    "minimal",      // Minimal style
    "dark",         // Dark theme
    "light",        // Light theme
    "retro",        // Retro style
    "neon",         // Neon style
    "compact",      // Compact style
}
```

### 4. Custom Configuration

```go
config := freezelib.DefaultConfig()

// Basic settings
config.SetTheme("github-dark")
config.SetFont("JetBrains Mono", 14)
config.SetBackground("#1e1e1e")
config.SetLanguage("python")

// Layout settings
config.SetPadding(20)           // All sides
config.SetPadding(20, 40)       // Vertical, horizontal
config.SetPadding(20, 40, 20, 40) // Top, right, bottom, left
config.SetMargin(15)
config.SetDimensions(800, 600)

// Decorative effects
config.SetWindow(true)          // Window controls
config.SetLineNumbers(true)     // Line numbers
config.SetShadow(20, 0, 10)     // Shadow: blur, X offset, Y offset
config.SetBorder(1, 8, "#333")  // Border: width, radius, color

// Line range (1-indexed)
config.SetLines(10, 20)         // Only capture lines 10-20

freeze := freezelib.NewWithConfig(config)
```

## 🎨 Supported Themes

- `github` / `github-dark`
- `dracula`
- `monokai`
- `solarized-dark` / `solarized-light`
- `nord`
- `one-dark`
- `material`
- `vim`
- And more...

## 💻 Supported Languages

Supports 100+ programming languages including:
- Go, Rust, Python, JavaScript, TypeScript
- C, C++, C#, Java, Kotlin, Swift
- HTML, CSS, SCSS, JSON, YAML, XML
- Shell, PowerShell, Dockerfile
- SQL, GraphQL, Markdown
- And more...

## 🔧 Advanced Usage

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

### Terminal Output Screenshots

```go
freeze := freezelib.NewWithPreset("terminal")

ansiOutput := "\033[32m✓ Tests passed\033[0m\n" +
              "\033[31m✗ Build failed\033[0m\n" +
              "\033[33m⚠ Warning: deprecated API\033[0m"

svgData, err := freeze.GenerateFromANSI(ansiOutput)
```

### Method Chaining

```go
freeze := freezelib.New().
    WithTheme("monokai").
    WithFont("Cascadia Code", 15).
    WithWindow(true).
    WithShadow(20, 0, 10).
    WithLineNumbers(true)

svgData, err := freeze.GenerateFromCode(code, "rust")
```

## 📊 Performance Optimization Tips

1. **Reuse instances**: Create one `Freeze` instance and reuse it
2. **Choose appropriate formats**: SVG for web, PNG for presentations
3. **Set specific dimensions**: Specifying dimensions improves performance
4. **Batch operations**: Process multiple files in a single session

## 🐛 Error Handling

```go
svgData, err := freeze.GenerateFromCode(code, "go")
if err != nil {
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

## 📁 Project Structure

```
freezelib/
├── freeze.go          # Main API interface
├── config.go          # Configuration structs
├── generator.go       # Core generation logic
├── quickfreeze.go     # Simplified API
├── presets.go         # Preset configurations
├── ansi.go           # ANSI processing
├── svg/              # SVG processing
├── font/             # Font processing
├── example/          # Usage examples
└── README.md         # Detailed documentation
```

## 🤝 Differences from Original freeze

| Feature | Original freeze | FreezeLib |
|---------|----------------|-----------|
| Usage | CLI tool | Go library |
| Integration | Command line calls | Direct import |
| Configuration | CLI args/config files | Go structs |
| Extensibility | Limited | Highly extensible |
| Performance | Process startup overhead | In-memory processing |

## 📝 Example Code

Check the complete examples in the `examples` directory:

This will generate multiple example SVG files showcasing various features of the library.
