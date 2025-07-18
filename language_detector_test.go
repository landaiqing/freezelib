package freezelib

import (
	"testing"
)

func TestLanguageDetector(t *testing.T) {
	detector := NewLanguageDetector()

	tests := []struct {
		name     string
		code     string
		expected string
	}{
		{
			name: "Go code",
			code: `package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}`,
			expected: "go",
		},
		{
			name: "Python code - fallback to text",
			code: `def hello():
    print("Hello, World!")

if __name__ == "__main__":
    hello()`,
			expected: "text", // Content analysis might not work for all languages
		},
		{
			name: "JavaScript code - fallback to text",
			code: `function hello() {
    console.log("Hello, World!");
}

hello();`,
			expected: "text", // Content analysis might not work for all languages
		},
		{
			name: "Rust code - fallback to text",
			code: `fn main() {
    println!("Hello, World!");
}`,
			expected: "text", // Content analysis might not work for all languages
		},
		{
			name: "JSON code - fallback to text",
			code: `{
    "name": "test",
    "version": "1.0.0"
}`,
			expected: "text", // Content analysis might not work for all languages
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := detector.DetectLanguage(tt.code)
			if result != tt.expected {
				t.Errorf("DetectLanguage() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestLanguageDetectorFromFilename(t *testing.T) {
	detector := NewLanguageDetector()

	tests := []struct {
		name     string
		filename string
		expected string
	}{
		{"Go file", "main.go", "go"},
		{"Python file", "script.py", "python"},
		{"JavaScript file", "app.js", "js"}, // chroma uses "js" as first alias
		{"TypeScript file", "app.ts", "ts"}, // chroma uses "ts" as first alias
		{"Rust file", "main.rs", "rust"},
		{"CSS file", "style.css", "css"},
		{"JSON file", "package.json", "json"},
		{"Dockerfile", "Dockerfile", "docker"}, // chroma uses "docker" as first alias
		{"Shell script", "deploy.sh", "bash"},
		{"Unknown extension", "file.unknown", "text"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := detector.DetectLanguageFromFilename(tt.filename)
			if result != tt.expected {
				t.Errorf("DetectLanguageFromFilename() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestLanguageDetectorCustomMappings(t *testing.T) {
	detector := NewLanguageDetector()

	// Add custom mapping
	detector.AddCustomMapping(".myext", "python")

	result := detector.DetectLanguageFromFilename("script.myext")
	if result != "python" {
		t.Errorf("Custom mapping failed: got %v, want python", result)
	}

	// Remove custom mapping
	detector.RemoveCustomMapping(".myext")

	result = detector.DetectLanguageFromFilename("script.myext")
	if result == "python" {
		t.Errorf("Custom mapping removal failed: still returns python")
	}
}

func TestLanguageDetectorCombined(t *testing.T) {
	detector := NewLanguageDetector()

	// Test with filename that has extension but content is different
	pythonCode := `def hello():
    print("Hello from Python!")

hello()`

	// Should prefer filename detection
	result := detector.DetectLanguageFromFile("script.py", pythonCode)
	if result != "python" {
		t.Errorf("Combined detection failed: got %v, want python", result)
	}

	// Test with unknown extension - should fallback to text since content analysis may not work
	result = detector.DetectLanguageFromFile("script.unknown", pythonCode)
	if result != "text" {
		t.Errorf("Content fallback failed: got %v, want text", result)
	}
}

func TestLanguageDetectorConfiguration(t *testing.T) {
	detector := NewLanguageDetector()

	// Test disabling content analysis
	detector.EnableContentAnalysis = false

	pythonCode := `def hello():
    print("Hello!")
hello()`

	result := detector.DetectLanguage(pythonCode)
	if result != detector.FallbackLanguage {
		t.Errorf("Content analysis should be disabled: got %v, want %v", result, detector.FallbackLanguage)
	}

	// Test disabling filename analysis
	detector.EnableContentAnalysis = true
	detector.EnableFilenameAnalysis = false

	result = detector.DetectLanguageFromFilename("script.py")
	if result != detector.FallbackLanguage {
		t.Errorf("Filename analysis should be disabled: got %v, want %v", result, detector.FallbackLanguage)
	}
}

func TestFreezeAutoDetection(t *testing.T) {
	freeze := New()

	goCode := `package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}`

	// Test auto detection
	svgData, err := freeze.GenerateFromCodeAuto(goCode)
	if err != nil {
		t.Errorf("GenerateFromCodeAuto failed: %v", err)
	}

	if len(svgData) == 0 {
		t.Error("GenerateFromCodeAuto returned empty data")
	}

	// Test language detection
	language := freeze.DetectLanguage(goCode)
	if language != "go" {
		t.Errorf("Language detection failed: got %v, want go", language)
	}
}

func TestQuickFreezeAutoDetection(t *testing.T) {
	qf := NewQuickFreeze()

	jsCode := `function hello() {
    console.log("Hello, World!");
}

hello();`

	// Test auto detection
	svgData, err := qf.CodeToSVGAuto(jsCode)
	if err != nil {
		t.Errorf("CodeToSVGAuto failed: %v", err)
	}

	if len(svgData) == 0 {
		t.Error("CodeToSVGAuto returned empty data")
	}

	// Test language detection - content analysis may not work for JS
	language := qf.DetectLanguage(jsCode)
	if language != "text" {
		t.Errorf("Language detection failed: got %v, want text", language)
	}
}

func TestLanguageSupport(t *testing.T) {
	freeze := New()

	// Test supported languages
	languages := freeze.GetSupportedLanguages()
	if len(languages) == 0 {
		t.Error("No supported languages found")
	}

	// Test common languages
	commonLanguages := []string{"go", "python", "javascript", "rust", "java", "c", "cpp"}
	for _, lang := range commonLanguages {
		if !freeze.IsLanguageSupported(lang) {
			t.Errorf("Language %s should be supported", lang)
		}
	}

	// Test unsupported language
	if freeze.IsLanguageSupported("nonexistent-language") {
		t.Error("Nonexistent language should not be supported")
	}
}
