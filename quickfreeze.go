package freezelib

import (
	"fmt"
	"strings"
)

// QuickFreeze provides a simplified, chainable API for quick code screenshots
type QuickFreeze struct {
	config *Config
}

// NewQuickFreeze creates a new QuickFreeze instance with default configuration
func NewQuickFreeze() *QuickFreeze {
	return &QuickFreeze{
		config: DefaultConfig(),
	}
}

// NewQuickFreezeWithPreset creates a new QuickFreeze instance with a preset
func NewQuickFreezeWithPreset(presetName string) *QuickFreeze {
	return &QuickFreeze{
		config: GetPreset(presetName),
	}
}

// WithTheme sets the syntax highlighting theme
func (qf *QuickFreeze) WithTheme(theme string) *QuickFreeze {
	qf.config.SetTheme(theme)
	return qf
}

// WithFont sets the font family and size
func (qf *QuickFreeze) WithFont(family string, size float64) *QuickFreeze {
	qf.config.SetFont(family, size)
	return qf
}

// WithBackground sets the background color
func (qf *QuickFreeze) WithBackground(color string) *QuickFreeze {
	qf.config.SetBackground(color)
	return qf
}

// WithWindow enables window controls
func (qf *QuickFreeze) WithWindow() *QuickFreeze {
	qf.config.SetWindow(true)
	return qf
}

// WithoutWindow disables window controls
func (qf *QuickFreeze) WithoutWindow() *QuickFreeze {
	qf.config.SetWindow(false)
	return qf
}

// WithLineNumbers enables line numbers
func (qf *QuickFreeze) WithLineNumbers() *QuickFreeze {
	qf.config.SetLineNumbers(true)
	return qf
}

// WithoutLineNumbers disables line numbers
func (qf *QuickFreeze) WithoutLineNumbers() *QuickFreeze {
	qf.config.SetLineNumbers(false)
	return qf
}

// WithShadow adds a shadow effect
func (qf *QuickFreeze) WithShadow() *QuickFreeze {
	qf.config.SetShadow(20, 0, 10)
	return qf
}

// WithCustomShadow adds a custom shadow effect
func (qf *QuickFreeze) WithCustomShadow(blur, x, y float64) *QuickFreeze {
	qf.config.SetShadow(blur, x, y)
	return qf
}

// WithoutShadow removes shadow effect
func (qf *QuickFreeze) WithoutShadow() *QuickFreeze {
	qf.config.SetShadow(0, 0, 0)
	return qf
}

// WithBorder adds a border
func (qf *QuickFreeze) WithBorder() *QuickFreeze {
	qf.config.SetBorder(1, 8, "#515151")
	return qf
}

// WithCustomBorder adds a custom border
func (qf *QuickFreeze) WithCustomBorder(width, radius float64, color string) *QuickFreeze {
	qf.config.SetBorder(width, radius, color)
	return qf
}

// WithoutBorder removes border
func (qf *QuickFreeze) WithoutBorder() *QuickFreeze {
	qf.config.SetBorder(0, 0, "")
	return qf
}

// WithPadding sets padding (1, 2, or 4 values like CSS)
func (qf *QuickFreeze) WithPadding(values ...float64) *QuickFreeze {
	qf.config.SetPadding(values...)
	return qf
}

// WithMargin sets margin (1, 2, or 4 values like CSS)
func (qf *QuickFreeze) WithMargin(values ...float64) *QuickFreeze {
	qf.config.SetMargin(values...)
	return qf
}

// WithDimensions sets specific width and height
func (qf *QuickFreeze) WithDimensions(width, height float64) *QuickFreeze {
	qf.config.SetDimensions(width, height)
	return qf
}

// WithWidth sets specific width (height auto)
func (qf *QuickFreeze) WithWidth(width float64) *QuickFreeze {
	qf.config.Width = width
	return qf
}

// WithHeight sets specific height (width auto)
func (qf *QuickFreeze) WithHeight(height float64) *QuickFreeze {
	qf.config.Height = height
	return qf
}

// WithLines sets the line range to capture (1-indexed)
func (qf *QuickFreeze) WithLines(start, end int) *QuickFreeze {
	qf.config.SetLines(start, end)
	return qf
}

// WithLanguage sets the programming language for syntax highlighting
func (qf *QuickFreeze) WithLanguage(language string) *QuickFreeze {
	qf.config.SetLanguage(language)
	return qf
}

// WithLineHeight sets the line height ratio
func (qf *QuickFreeze) WithLineHeight(ratio float64) *QuickFreeze {
	qf.config.LineHeight = ratio
	return qf
}

// WithWrap sets text wrapping at specified column
func (qf *QuickFreeze) WithWrap(columns int) *QuickFreeze {
	qf.config.Wrap = columns
	return qf
}

// CodeToSVG generates SVG from source code
func (qf *QuickFreeze) CodeToSVG(code string) ([]byte, error) {
	generator := NewGenerator(qf.config)
	return generator.GenerateFromCode(code, qf.config.Language)
}

// CodeToPNG generates PNG from source code
func (qf *QuickFreeze) CodeToPNG(code string) ([]byte, error) {
	generator := NewGenerator(qf.config)
	svgData, err := generator.GenerateFromCode(code, qf.config.Language)
	if err != nil {
		return nil, err
	}

	width := qf.config.Width
	height := qf.config.Height
	if width == 0 || height == 0 {
		width = 800 * 4
		height = 600 * 4
	} else {
		width *= 4
		height *= 4
	}

	return generator.ConvertToPNG(svgData, width, height)
}

// FileToSVG generates SVG from a source code file
func (qf *QuickFreeze) FileToSVG(filename string) ([]byte, error) {
	generator := NewGenerator(qf.config)
	return generator.GenerateFromFile(filename)
}

// FileToPNG generates PNG from a source code file
func (qf *QuickFreeze) FileToPNG(filename string) ([]byte, error) {
	generator := NewGenerator(qf.config)
	svgData, err := generator.GenerateFromFile(filename)
	if err != nil {
		return nil, err
	}

	width := qf.config.Width
	height := qf.config.Height
	if width == 0 || height == 0 {
		width = 800 * 4
		height = 600 * 4
	} else {
		width *= 4
		height *= 4
	}

	return generator.ConvertToPNG(svgData, width, height)
}

// ANSIToSVG generates SVG from ANSI terminal output
func (qf *QuickFreeze) ANSIToSVG(ansiOutput string) ([]byte, error) {
	generator := NewGenerator(qf.config)
	return generator.GenerateFromANSI(ansiOutput)
}

// ANSIToPNG generates PNG from ANSI terminal output
func (qf *QuickFreeze) ANSIToPNG(ansiOutput string) ([]byte, error) {
	generator := NewGenerator(qf.config)
	svgData, err := generator.GenerateFromANSI(ansiOutput)
	if err != nil {
		return nil, err
	}

	width := qf.config.Width
	height := qf.config.Height
	if width == 0 || height == 0 {
		width = 800 * 4
		height = 600 * 4
	} else {
		width *= 4
		height *= 4
	}

	return generator.ConvertToPNG(svgData, width, height)
}

// SaveCodeToFile generates and saves code screenshot to file
func (qf *QuickFreeze) SaveCodeToFile(code, filename string) error {
	var data []byte
	var err error

	if isPNGFile(filename) {
		data, err = qf.CodeToPNG(code)
	} else {
		data, err = qf.CodeToSVG(code)
	}

	if err != nil {
		return err
	}

	return saveToFile(data, filename)
}

// SaveFileToFile generates and saves file screenshot to file
func (qf *QuickFreeze) SaveFileToFile(inputFile, outputFile string) error {
	var data []byte
	var err error

	if isPNGFile(outputFile) {
		data, err = qf.FileToPNG(inputFile)
	} else {
		data, err = qf.FileToSVG(inputFile)
	}

	if err != nil {
		return err
	}

	return saveToFile(data, outputFile)
}

// SaveANSIToFile generates and saves ANSI screenshot to file
func (qf *QuickFreeze) SaveANSIToFile(ansiOutput, filename string) error {
	var data []byte
	var err error

	if isPNGFile(filename) {
		data, err = qf.ANSIToPNG(ansiOutput)
	} else {
		data, err = qf.ANSIToSVG(ansiOutput)
	}

	if err != nil {
		return err
	}

	return saveToFile(data, filename)
}

// Config returns the current configuration
func (qf *QuickFreeze) Config() *Config {
	return qf.config
}

// Clone creates a copy of the QuickFreeze instance
func (qf *QuickFreeze) Clone() *QuickFreeze {
	return &QuickFreeze{
		config: qf.config.Clone(),
	}
}

// Reset resets the configuration to defaults
func (qf *QuickFreeze) Reset() *QuickFreeze {
	qf.config = DefaultConfig()
	return qf
}

// ResetToPreset resets the configuration to a specific preset
func (qf *QuickFreeze) ResetToPreset(presetName string) *QuickFreeze {
	qf.config = GetPreset(presetName)
	return qf
}

// String returns a string representation of the current configuration
func (qf *QuickFreeze) String() string {
	var parts []string

	parts = append(parts, fmt.Sprintf("Theme: %s", qf.config.Theme))
	parts = append(parts, fmt.Sprintf("Font: %s %.1fpx", qf.config.Font.Family, qf.config.Font.Size))
	parts = append(parts, fmt.Sprintf("Background: %s", qf.config.Background))

	if qf.config.Window {
		parts = append(parts, "Window: enabled")
	}

	if qf.config.ShowLineNumbers {
		parts = append(parts, "Line numbers: enabled")
	}

	if qf.config.Shadow.Blur > 0 {
		parts = append(parts, fmt.Sprintf("Shadow: blur=%.1f", qf.config.Shadow.Blur))
	}

	if qf.config.Border.Width > 0 {
		parts = append(parts, fmt.Sprintf("Border: width=%.1f", qf.config.Border.Width))
	}

	return "QuickFreeze{" + strings.Join(parts, ", ") + "}"
}

// saveToFile is a helper function to save data to file
func saveToFile(data []byte, filename string) error {
	return NewWithConfig(DefaultConfig()).SaveToFile(data, filename)
}
