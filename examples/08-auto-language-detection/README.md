# Auto Language Detection Examples

This example demonstrates FreezeLib's enhanced automatic language detection capabilities.

## Features

### 🎯 Automatic Language Detection
- **Content-based detection**: Analyzes code content to identify the programming language
- **Filename-based detection**: Uses file extensions and names to determine language
- **Combined detection**: Intelligently combines both methods for best results
- **Fallback support**: Gracefully handles unknown languages

### 🔧 Enhanced API
- `GenerateFromCodeAuto()` - Generate screenshots without specifying language
- `DetectLanguage()` - Detect language from code content
- `DetectLanguageFromFilename()` - Detect language from filename
- `DetectLanguageFromFile()` - Combined detection from both filename and content
- `GetSupportedLanguages()` - List all supported languages
- `IsLanguageSupported()` - Check if a language is supported

### ⚙️ Customizable Detection
- Custom file extension mappings
- Configurable detection strategies
- Fallback language settings

## Usage Examples

### Basic Auto Detection

```go
freeze := freezelib.New()

code := `
def hello_world():
    print("Hello, World!")

if __name__ == "__main__":
    hello_world()
`

// Automatically detect language and generate screenshot
svgData, err := freeze.GenerateFromCodeAuto(code)
```

### Language Detection API

```go
freeze := freezelib.New()

// Detect language from content
language := freeze.DetectLanguage(code)
fmt.Printf("Detected language: %s\n", language)

// Detect from filename
language = freeze.DetectLanguageFromFilename("script.py")
fmt.Printf("Language from filename: %s\n", language)

// Combined detection
language = freeze.DetectLanguageFromFile("script.py", code)
fmt.Printf("Combined detection: %s\n", language)
```

### Custom Language Detector

```go
freeze := freezelib.New()

// Get and customize the language detector
detector := freeze.GetLanguageDetector()

// Add custom file extension mappings
detector.AddCustomMapping(".myext", "python")
detector.AddCustomMapping(".config", "json")

// Use custom mappings
language := freeze.DetectLanguageFromFilename("app.config")
// Returns "json" due to custom mapping
```

### QuickFreeze Auto Detection

```go
qf := freezelib.NewQuickFreeze()

svgData, err := qf.WithTheme("dracula").
    WithFont("Fira Code", 14).
    WithWindow().
    CodeToSVGAuto(code) // Auto-detect language
```

## Supported Languages

FreezeLib supports 100+ programming languages including:

### Popular Languages
- **Go** - `.go`
- **Python** - `.py`, `.pyw`
- **JavaScript** - `.js`, `.mjs`
- **TypeScript** - `.ts`, `.tsx`
- **Rust** - `.rs`
- **Java** - `.java`
- **C/C++** - `.c`, `.cpp`, `.cc`, `.cxx`, `.h`, `.hpp`
- **C#** - `.cs`
- **PHP** - `.php`
- **Ruby** - `.rb`

### Web Technologies
- **HTML** - `.html`, `.htm`
- **CSS** - `.css`
- **SCSS/Sass** - `.scss`, `.sass`
- **JSON** - `.json`
- **XML** - `.xml`

### Shell & Scripts
- **Bash** - `.sh`, `.bash`
- **PowerShell** - `.ps1`
- **Batch** - `.bat`, `.cmd`
- **Fish** - `.fish`
- **Zsh** - `.zsh`

### Configuration & Data
- **YAML** - `.yaml`, `.yml`
- **TOML** - `.toml`
- **INI** - `.ini`, `.cfg`, `.conf`
- **SQL** - `.sql`
- **Dockerfile** - `Dockerfile`, `.dockerfile`

### And Many More...
- Kotlin, Swift, Scala, Clojure, Haskell, OCaml, F#, Erlang, Elixir
- Julia, Nim, Zig, V, D, Pascal, Ada, Fortran, COBOL
- Assembly (NASM, GAS), MATLAB, R, Lua, Dart, Elm
- GraphQL, Protocol Buffers, Markdown, LaTeX, Vim script

## Detection Strategies

### 1. Content Analysis
Uses Chroma's built-in lexer analysis to examine code patterns, keywords, and syntax.

```go
// Analyzes code structure and syntax
language := freeze.DetectLanguage(`
package main
import "fmt"
func main() { fmt.Println("Hello") }
`)
// Returns: "go"
```

### 2. Filename Analysis
Examines file extensions and special filenames.

```go
// Uses file extension mapping
language := freeze.DetectLanguageFromFilename("script.py")
// Returns: "python"

// Recognizes special files
language := freeze.DetectLanguageFromFilename("Dockerfile")
// Returns: "dockerfile"
```

### 3. Combined Analysis
Intelligently combines both methods for best accuracy.

```go
// Tries filename first, then content analysis
language := freeze.DetectLanguageFromFile("script.unknown", pythonCode)
// Returns: "python" (from content analysis)
```

## Configuration Options

### Language Detector Settings

```go
detector := freeze.GetLanguageDetector()

// Enable/disable detection methods
detector.EnableContentAnalysis = true
detector.EnableFilenameAnalysis = true

// Set fallback language
detector.FallbackLanguage = "text"

// Add custom mappings
detector.AddCustomMapping(".myext", "python")
```

## Error Handling

```go
svgData, err := freeze.GenerateFromCodeAuto(code)
if err != nil {
    if strings.Contains(err.Error(), "could not determine language") {
        // Language detection failed
        // Try with explicit language or check supported languages
        fmt.Println("Supported languages:", freeze.GetSupportedLanguages())
    }
}
```

## Running the Examples

```bash
cd examples/08-auto-language-detection
go run main.go
```

This will generate various screenshots demonstrating:
- Basic auto detection with different languages
- Language detection API usage
- Custom language detector configuration
- Batch processing with auto detection
- Language analysis and support information

## Output Files

The examples generate several SVG files in the `output/` directory:
- `auto_*.svg` - Basic auto detection examples
- `detection_*.svg` - Language detection API examples  
- `custom_*.svg` - Custom detector examples
- `batch_*.svg` - Batch processing examples
- `language_summary.svg` - Language support summary
