package font

import (
	"embed"
	"fmt"

	formatter "github.com/alecthomas/chroma/v2/formatters/svg"
)

//go:embed *.ttf
var fonts embed.FS

// JetBrainsMonoTTF contains the JetBrains Mono font data
var JetBrainsMonoTTF []byte

// JetBrainsMonoNLTTF contains the JetBrains Mono NL font data
var JetBrainsMonoNLTTF []byte

func init() {
	var err error
	JetBrainsMonoTTF, err = fonts.ReadFile("JetBrainsMono-Regular.ttf")
	if err != nil {
		// If embedded font is not available, use empty slice
		JetBrainsMonoTTF = []byte{}
	}

	JetBrainsMonoNLTTF, err = fonts.ReadFile("JetBrainsMonoNL-Regular.ttf")
	if err != nil {
		// If embedded font is not available, use empty slice
		JetBrainsMonoNLTTF = []byte{}
	}
}

// FontOptions creates formatter options for the given font configuration
func FontOptions(family string, size float64, ligatures bool, fontFile string) ([]formatter.Option, error) {
	var options []formatter.Option

	// Set font family
	if family != "" {
		options = append(options, formatter.FontFamily(family))
	}

	// Embed font file if specified
	if fontFile != "" {
		option, err := formatter.EmbedFontFile(family, fontFile)
		if err != nil {
			return nil, fmt.Errorf("failed to embed font file: %w", err)
		}
		options = append(options, option)
	}

	return options, nil
}

// GetDefaultFontFamily returns the default font family
func GetDefaultFontFamily() string {
	return "JetBrains Mono"
}

// GetDefaultFontSize returns the default font size
func GetDefaultFontSize() float64 {
	return 14.0
}

// IsMonospaceFont checks if a font family is monospace
func IsMonospaceFont(family string) bool {
	monospaceFonts := map[string]bool{
		"JetBrains Mono":   true,
		"Fira Code":        true,
		"Source Code Pro":  true,
		"Monaco":           true,
		"Menlo":            true,
		"Consolas":         true,
		"Courier New":      true,
		"monospace":        true,
		"SF Mono":          true,
		"Cascadia Code":    true,
		"Ubuntu Mono":      true,
		"DejaVu Sans Mono": true,
		"Liberation Mono":  true,
		"Inconsolata":      true,
		"Roboto Mono":      true,
	}
	return monospaceFonts[family]
}

// ValidateFontFamily validates if a font family name is valid
func ValidateFontFamily(family string) error {
	if family == "" {
		return fmt.Errorf("font family cannot be empty")
	}
	return nil
}

// ValidateFontSize validates if a font size is valid
func ValidateFontSize(size float64) error {
	if size <= 0 {
		return fmt.Errorf("font size must be positive, got %.2f", size)
	}
	if size > 100 {
		return fmt.Errorf("font size too large, got %.2f", size)
	}
	return nil
}

// GetFontHeightToWidthRatio returns the typical height to width ratio for monospace fonts
func GetFontHeightToWidthRatio() float64 {
	return 1.68
}

// CalculateTextWidth estimates the width of text in pixels
func CalculateTextWidth(text string, fontSize float64) float64 {
	return float64(len(text)) * (fontSize / GetFontHeightToWidthRatio())
}

// CalculateLineHeight calculates the line height in pixels
func CalculateLineHeight(fontSize, lineHeightRatio float64) float64 {
	return fontSize * lineHeightRatio
}

// GetEmbeddedFontData returns embedded font data if available
func GetEmbeddedFontData(fontName string) []byte {
	switch fontName {
	case "JetBrains Mono", "JetBrainsMono":
		return JetBrainsMonoTTF
	case "JetBrains Mono NL", "JetBrainsMonoNL":
		return JetBrainsMonoNLTTF
	default:
		return nil
	}
}

// FontConfig represents font configuration
type FontConfig struct {
	Family    string
	Size      float64
	Ligatures bool
	File      string
}

// NewFontConfig creates a new font configuration with defaults
func NewFontConfig() *FontConfig {
	return &FontConfig{
		Family:    GetDefaultFontFamily(),
		Size:      GetDefaultFontSize(),
		Ligatures: true,
		File:      "",
	}
}

// Validate validates the font configuration
func (fc *FontConfig) Validate() error {
	if err := ValidateFontFamily(fc.Family); err != nil {
		return err
	}
	if err := ValidateFontSize(fc.Size); err != nil {
		return err
	}
	return nil
}

// ToFormatterOptions converts font config to formatter options
func (fc *FontConfig) ToFormatterOptions() ([]formatter.Option, error) {
	return FontOptions(fc.Family, fc.Size, fc.Ligatures, fc.File)
}
