# FreezeLib Usage Guide

**Language / 语言**: [English](USAGE_EN.md) | [中文](USAGE.md)

FreezeLib is a Go library for generating beautiful code screenshots.

## Installation

```bash
go get github.com/landaiqing/freezelib
```

## Basic Usage

### 1. Simplest Example

```go
package main

import (
	"github.com/landaiqing/freezelib"
	"os"
)

func main() {
	// Create instance
	freeze := freezelib.New()

	// Code content
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

	// Save file
	os.WriteFile("hello.svg", svgData, 0644)
}
```

### 2. Generate from File

```go
// Generate directly from file
svgData, err := freeze.GenerateFromFile("main.go")
if err != nil {
    panic(err)
}
os.WriteFile("main.svg", svgData, 0644)
```

### 3. Generate PNG Format

```go
// Generate PNG instead of SVG
pngData, err := freeze.GeneratePNGFromCode(code, "go")
if err != nil {
    panic(err)
}
os.WriteFile("hello.png", pngData, 0644)
```

### 4. Save Directly to File

```go
// One step: generate and save
err := freeze.SaveCodeToFile(code, "go", "output.svg")
if err != nil {
    panic(err)
}

// Auto-detect format (by file extension)
err = freeze.SaveCodeToFile(code, "go", "output.png")
```

## Using Preset Styles

```go
// Use preset configurations for quick start
freeze := freezelib.NewWithPreset("dark")        // Dark theme
freeze := freezelib.NewWithPreset("terminal")    // Terminal style
freeze := freezelib.NewWithPreset("presentation") // Presentation style

// View all available presets
presets := freeze.GetAvailablePresets()
// Returns: ["base", "compact", "dark", "full", "light", "minimal", "neon", "presentation", "retro", "terminal"]
```

## Custom Styling

### Basic Settings

```go
freeze := freezelib.New().
    WithTheme("github-dark").           // Set theme
    WithFont("JetBrains Mono", 14).     // Set font and size
    WithBackground("#1e1e1e").          // Set background color
    WithPadding(20)                     // Set padding

svgData, err := freeze.GenerateFromCode(code, "go")
```

### Add Decorative Effects

```go
freeze := freezelib.New().
    WithTheme("dracula").
    WithWindow(true).                   // Add window controls
    WithLineNumbers(true).              // Show line numbers
    WithShadow(20, 0, 10)               // Add shadow

svgData, err := freeze.GenerateFromCode(code, "python")
```

## View Supported Options

```go
freeze := freezelib.New()

// View all supported languages (270+)
languages := freeze.GetSupportedLanguages()
fmt.Printf("Supports %d languages\n", len(languages))

// View all supported themes (67+)
themes := freeze.GetSupportedThemes()
fmt.Printf("Supports %d themes\n", len(themes))

// View all available presets (10)
presets := freeze.GetAvailablePresets()
fmt.Printf("Available presets: %v\n", presets)
```

## Terminal Output Screenshots

```go
// Screenshot terminal output (supports ANSI colors)
freeze := freezelib.NewWithPreset("terminal")

ansiOutput := "\033[32m✓ Tests passed\033[0m\n" +
              "\033[31m✗ Build failed\033[0m\n" +
              "\033[33m⚠ Warning: deprecated API\033[0m"

svgData, err := freeze.GenerateFromANSI(ansiOutput)
if err != nil {
    panic(err)
}
os.WriteFile("terminal.svg", svgData, 0644)
```

## Batch Processing Files

```go
freeze := freezelib.NewWithPreset("dark")
files := []string{"main.go", "config.go", "utils.go"}

for _, file := range files {
    err := freeze.SaveFileToFile(file, file+".svg")
    if err != nil {
        fmt.Printf("Failed to process %s: %v\n", file, err)
        continue
    }
    fmt.Printf("Generated: %s.svg\n", file)
}
```

## Auto Language Detection

```go
freeze := freezelib.New()

// Don't specify language, auto-detect
code := `function hello() {
    console.log("Hello, World!");
}`

// Auto-detected as JavaScript
svgData, err := freeze.GenerateFromCodeAuto(code)
if err != nil {
    panic(err)
}
os.WriteFile("auto.svg", svgData, 0644)
```

## Error Handling

```go
svgData, err := freeze.GenerateFromCode(code, "go")
if err != nil {
    fmt.Printf("Generation failed: %v\n", err)
    return
}

// Save file
err = os.WriteFile("output.svg", svgData, 0644)
if err != nil {
    fmt.Printf("Save failed: %v\n", err)
    return
}

fmt.Println("Generated successfully!")
```

## Common Themes

- `github` - GitHub light theme
- `github-dark` - GitHub dark theme
- `dracula` - Dracula theme
- `monokai` - Monokai theme
- `nord` - Nord theme

## Common Languages

- `go` - Go language
- `python` - Python
- `javascript` - JavaScript
- `typescript` - TypeScript
- `java` - Java
- `rust` - Rust
- `c` - C language
- `cpp` - C++
- `html` - HTML
- `css` - CSS
- `json` - JSON
- `yaml` - YAML
- `markdown` - Markdown
- `bash` - Shell script
