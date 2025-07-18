package freezelib

import (
	"path/filepath"
	"strings"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/lexers"
)

// LanguageDetector provides enhanced language detection capabilities
type LanguageDetector struct {
	// EnableContentAnalysis enables content-based language detection
	EnableContentAnalysis bool
	// EnableFilenameAnalysis enables filename-based language detection
	EnableFilenameAnalysis bool
	// FallbackLanguage is used when detection fails
	FallbackLanguage string
	// CustomMappings allows custom file extension to language mappings
	CustomMappings map[string]string
}

// NewLanguageDetector creates a new language detector with default settings
func NewLanguageDetector() *LanguageDetector {
	return &LanguageDetector{
		EnableContentAnalysis:  true,
		EnableFilenameAnalysis: true,
		FallbackLanguage:       "text",
		CustomMappings:         make(map[string]string),
	}
}

// DetectLanguage detects the programming language from code content
func (ld *LanguageDetector) DetectLanguage(code string) string {
	if !ld.EnableContentAnalysis {
		return ld.FallbackLanguage
	}

	lexer := lexers.Analyse(code)
	if lexer != nil {
		config := lexer.Config()
		if config != nil && len(config.Aliases) > 0 {
			return config.Aliases[0]
		}
		if config != nil && config.Name != "" {
			return strings.ToLower(config.Name)
		}
	}

	return ld.FallbackLanguage
}

// DetectLanguageFromFilename detects the programming language from filename
func (ld *LanguageDetector) DetectLanguageFromFilename(filename string) string {
	if !ld.EnableFilenameAnalysis {
		return ld.FallbackLanguage
	}

	// Check custom mappings first
	ext := strings.ToLower(filepath.Ext(filename))
	if lang, exists := ld.CustomMappings[ext]; exists {
		return lang
	}

	// Use chroma's built-in filename detection
	lexer := lexers.Match(filename)
	if lexer != nil {
		config := lexer.Config()
		if config != nil && len(config.Aliases) > 0 {
			// Return the first alias which is usually the most common name
			return config.Aliases[0]
		}
		if config != nil && config.Name != "" {
			return strings.ToLower(config.Name)
		}
	}

	// Fallback to common extension mappings
	return ld.detectFromExtension(ext)
}

// DetectLanguageFromFile detects language from both filename and content
func (ld *LanguageDetector) DetectLanguageFromFile(filename, content string) string {
	// Try filename first
	if ld.EnableFilenameAnalysis {
		lang := ld.DetectLanguageFromFilename(filename)
		if lang != ld.FallbackLanguage {
			return lang
		}
	}

	// Try content analysis
	if ld.EnableContentAnalysis {
		lang := ld.DetectLanguage(content)
		if lang != ld.FallbackLanguage {
			return lang
		}
	}

	return ld.FallbackLanguage
}

// GetLexer returns a chroma lexer for the given language or content
func (ld *LanguageDetector) GetLexer(language, content string) chroma.Lexer {
	var lexer chroma.Lexer

	// Try to get lexer by language name
	if language != "" {
		lexer = lexers.Get(language)
		if lexer != nil {
			return lexer
		}
	}

	// Try content analysis if enabled
	if ld.EnableContentAnalysis && content != "" {
		lexer = lexers.Analyse(content)
		if lexer != nil {
			return lexer
		}
	}

	// Return fallback lexer
	return lexers.Get(ld.FallbackLanguage)
}

// GetLexerFromFile returns a chroma lexer for the given file
func (ld *LanguageDetector) GetLexerFromFile(filename, content string) chroma.Lexer {
	var lexer chroma.Lexer

	// Try filename detection first if enabled
	if ld.EnableFilenameAnalysis {
		lexer = lexers.Match(filename)
		if lexer != nil {
			return lexer
		}
	}

	// Try content analysis if enabled
	if ld.EnableContentAnalysis && content != "" {
		lexer = lexers.Analyse(content)
		if lexer != nil {
			return lexer
		}
	}

	// Return fallback lexer
	return lexers.Get(ld.FallbackLanguage)
}

// AddCustomMapping adds a custom file extension to language mapping
func (ld *LanguageDetector) AddCustomMapping(extension, language string) {
	if ld.CustomMappings == nil {
		ld.CustomMappings = make(map[string]string)
	}
	ld.CustomMappings[strings.ToLower(extension)] = language
}

// RemoveCustomMapping removes a custom file extension mapping
func (ld *LanguageDetector) RemoveCustomMapping(extension string) {
	if ld.CustomMappings != nil {
		delete(ld.CustomMappings, strings.ToLower(extension))
	}
}

// GetSupportedLanguages returns a list of all supported languages
func (ld *LanguageDetector) GetSupportedLanguages() []string {
	return lexers.Names(false) // false means don't include aliases
}

// IsLanguageSupported checks if a language is supported
func (ld *LanguageDetector) IsLanguageSupported(language string) bool {
	lexer := lexers.Get(language)
	return lexer != nil
}

// detectFromExtension provides fallback extension-based detection
func (ld *LanguageDetector) detectFromExtension(ext string) string {
	commonMappings := map[string]string{
		".go":         "go",
		".py":         "python",
		".js":         "javascript",
		".ts":         "typescript",
		".jsx":        "jsx",
		".tsx":        "tsx",
		".java":       "java",
		".c":          "c",
		".cpp":        "cpp",
		".cc":         "cpp",
		".cxx":        "cpp",
		".h":          "c",
		".hpp":        "cpp",
		".cs":         "csharp",
		".php":        "php",
		".rb":         "ruby",
		".rs":         "rust",
		".swift":      "swift",
		".kt":         "kotlin",
		".scala":      "scala",
		".clj":        "clojure",
		".hs":         "haskell",
		".ml":         "ocaml",
		".fs":         "fsharp",
		".vb":         "vbnet",
		".pl":         "perl",
		".r":          "r",
		".m":          "matlab",
		".lua":        "lua",
		".sh":         "bash",
		".bash":       "bash",
		".zsh":        "zsh",
		".fish":       "fish",
		".ps1":        "powershell",
		".bat":        "batch",
		".cmd":        "batch",
		".html":       "html",
		".htm":        "html",
		".xml":        "xml",
		".css":        "css",
		".scss":       "scss",
		".sass":       "sass",
		".less":       "less",
		".json":       "json",
		".yaml":       "yaml",
		".yml":        "yaml",
		".toml":       "toml",
		".ini":        "ini",
		".cfg":        "ini",
		".conf":       "ini",
		".sql":        "sql",
		".md":         "markdown",
		".markdown":   "markdown",
		".tex":        "latex",
		".dockerfile": "dockerfile",
		".makefile":   "makefile",
		".mk":         "makefile",
		".vim":        "vim",
		".proto":      "protobuf",
		".graphql":    "graphql",
		".gql":        "graphql",
		".dart":       "dart",
		".elm":        "elm",
		".ex":         "elixir",
		".exs":        "elixir",
		".erl":        "erlang",
		".hrl":        "erlang",
		".jl":         "julia",
		".nim":        "nim",
		".zig":        "zig",
		".v":          "v",
		".d":          "d",
		".pas":        "pascal",
		".pp":         "pascal",
		".ada":        "ada",
		".adb":        "ada",
		".ads":        "ada",
		".f":          "fortran",
		".f90":        "fortran",
		".f95":        "fortran",
		".f03":        "fortran",
		".f08":        "fortran",
		".cob":        "cobol",
		".cbl":        "cobol",
		".asm":        "nasm",
		".s":          "gas",
	}

	if lang, exists := commonMappings[ext]; exists {
		return lang
	}

	return ld.FallbackLanguage
}
