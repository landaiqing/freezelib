# FreezeLib Examples

**Language / 语言**: [English](README.md) | [中文](README_CN.md)

This directory contains comprehensive examples demonstrating various features of FreezeLib.

## 📁 Example Categories

### [01-basic/](01-basic/) - Basic Usage
- Simple code screenshot generation
- Basic configuration
- Getting started examples

### [02-formats/](02-formats/) - Output Formats
- SVG output examples
- PNG output examples
- Format comparison
- Quality settings

### [03-themes/](03-themes/) - Theme Showcase
- Popular themes demonstration
- Theme comparison
- Custom theme creation

### [04-languages/](04-languages/) - Programming Languages
- Syntax highlighting for different languages
- Language-specific optimizations
- Multi-language projects

### [05-terminal/](05-terminal/) - Terminal Output
- ANSI color support
- Terminal styling
- Command output screenshots

### [06-advanced/](06-advanced/) - Advanced Configuration
- Complex styling options
- Performance optimization
- Custom fonts and layouts

### [07-batch/](07-batch/) - Batch Processing
- Multiple file processing
- Automated workflows
- Bulk operations

### [08-auto-language-detection/](08-auto-language-detection/) - Auto Language Detection
- Content-based language detection
- Filename-based detection
- Custom detector configuration
- Detection API usage

### [09-supported-options/](09-supported-options/) - Supported Options
- Get all supported languages list
- Get all supported themes list
- Get all available presets list
- Validate language, theme, preset support
- Simple and efficient API

## 🚀 Quick Start

To run all examples:

```bash
cd examples
go run run_all_examples.go
```

To run specific category:

```bash
# Basic usage
cd examples/01-basic
go run main.go

# Auto language detection
cd examples/08-auto-language-detection
go run main.go

# Supported options
cd examples/09-supported-options
go run main.go
```

## 🤝 Contributing Examples

To add new examples:

1. Choose appropriate category or create new one
2. Follow the naming convention: `example_name.go`
3. Include both code and generated output
4. Add documentation in README.md
5. Test with `go run main.go`