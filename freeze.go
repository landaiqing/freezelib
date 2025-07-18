// Package freezelib provides a Go library for generating beautiful code screenshots
// from source code and terminal output.
//
// This library is based on the freeze CLI tool by Charm and provides a programmatic
// interface for creating code screenshots with syntax highlighting, themes, and
// various styling options.
package freezelib

import (
	"fmt"
	"io"
	"os"
)

// Freeze is the main interface for generating code screenshots
type Freeze struct {
	generator *Generator
	config    *Config
}

// New creates a new Freeze instance with default configuration
func New() *Freeze {
	config := DefaultConfig()
	return &Freeze{
		generator: NewGenerator(config),
		config:    config,
	}
}

// NewWithConfig creates a new Freeze instance with the provided configuration
func NewWithConfig(config *Config) *Freeze {
	if config == nil {
		config = DefaultConfig()
	}
	return &Freeze{
		generator: NewGenerator(config),
		config:    config,
	}
}

// NewWithPreset creates a new Freeze instance with a preset configuration
func NewWithPreset(presetName string) *Freeze {
	config := GetPreset(presetName)
	return &Freeze{
		generator: NewGenerator(config),
		config:    config,
	}
}

// Config returns the current configuration
func (f *Freeze) Config() *Config {
	return f.config
}

// SetConfig updates the configuration and recreates the generator
func (f *Freeze) SetConfig(config *Config) *Freeze {
	f.config = config
	f.generator = NewGenerator(config)
	return f
}

// UpdateConfig allows modifying the current configuration
func (f *Freeze) UpdateConfig(fn func(*Config)) *Freeze {
	fn(f.config)
	f.generator = NewGenerator(f.config)
	return f
}

// GenerateFromCode generates an SVG screenshot from source code
func (f *Freeze) GenerateFromCode(code, language string) ([]byte, error) {
	return f.generator.GenerateFromCode(code, language)
}

// GenerateFromFile generates an SVG screenshot from a source code file
func (f *Freeze) GenerateFromFile(filename string) ([]byte, error) {
	return f.generator.GenerateFromFile(filename)
}

// GenerateFromReader generates an SVG screenshot from a reader containing source code
func (f *Freeze) GenerateFromReader(reader io.Reader, language string) ([]byte, error) {
	content, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read from reader: %w", err)
	}
	return f.generator.GenerateFromCode(string(content), language)
}

// GenerateFromANSI generates an SVG screenshot from ANSI terminal output
func (f *Freeze) GenerateFromANSI(ansiOutput string) ([]byte, error) {
	return f.generator.GenerateFromANSI(ansiOutput)
}

// GeneratePNGFromCode generates a PNG screenshot from source code
func (f *Freeze) GeneratePNGFromCode(code, language string) ([]byte, error) {
	svgData, err := f.generator.GenerateFromCode(code, language)
	if err != nil {
		return nil, err
	}

	// Calculate dimensions for PNG (use 4x scale for better quality)
	width := f.config.Width
	height := f.config.Height
	if width == 0 || height == 0 {
		// Use default dimensions with 4x scale
		width = 800 * 4
		height = 600 * 4
	} else {
		width *= 4
		height *= 4
	}

	return f.generator.ConvertToPNG(svgData, width, height)
}

// GeneratePNGFromFile generates a PNG screenshot from a source code file
func (f *Freeze) GeneratePNGFromFile(filename string) ([]byte, error) {
	svgData, err := f.generator.GenerateFromFile(filename)
	if err != nil {
		return nil, err
	}

	// Calculate dimensions for PNG
	width := f.config.Width
	height := f.config.Height
	if width == 0 || height == 0 {
		width = 800 * 4
		height = 600 * 4
	} else {
		width *= 4
		height *= 4
	}

	return f.generator.ConvertToPNG(svgData, width, height)
}

// GeneratePNGFromANSI generates a PNG screenshot from ANSI terminal output
func (f *Freeze) GeneratePNGFromANSI(ansiOutput string) ([]byte, error) {
	svgData, err := f.generator.GenerateFromANSI(ansiOutput)
	if err != nil {
		return nil, err
	}

	// Calculate dimensions for PNG
	width := f.config.Width
	height := f.config.Height
	if width == 0 || height == 0 {
		width = 800 * 4
		height = 600 * 4
	} else {
		width *= 4
		height *= 4
	}

	return f.generator.ConvertToPNG(svgData, width, height)
}

// SaveToFile saves the generated SVG to a file
func (f *Freeze) SaveToFile(data []byte, filename string) error {
	return os.WriteFile(filename, data, 0644)
}

// SaveCodeToFile generates and saves a code screenshot to a file
func (f *Freeze) SaveCodeToFile(code, language, filename string) error {
	var data []byte
	var err error

	if isPNGFile(filename) {
		data, err = f.GeneratePNGFromCode(code, language)
	} else {
		data, err = f.GenerateFromCode(code, language)
	}

	if err != nil {
		return err
	}

	return f.SaveToFile(data, filename)
}

// SaveFileToFile generates and saves a file screenshot to a file
func (f *Freeze) SaveFileToFile(inputFile, outputFile string) error {
	var data []byte
	var err error

	if isPNGFile(outputFile) {
		data, err = f.GeneratePNGFromFile(inputFile)
	} else {
		data, err = f.GenerateFromFile(inputFile)
	}

	if err != nil {
		return err
	}

	return f.SaveToFile(data, outputFile)
}

// SaveANSIToFile generates and saves an ANSI screenshot to a file
func (f *Freeze) SaveANSIToFile(ansiOutput, filename string) error {
	var data []byte
	var err error

	if isPNGFile(filename) {
		data, err = f.GeneratePNGFromANSI(ansiOutput)
	} else {
		data, err = f.GenerateFromANSI(ansiOutput)
	}

	if err != nil {
		return err
	}

	return f.SaveToFile(data, filename)
}

// Clone creates a copy of the Freeze instance with the same configuration
func (f *Freeze) Clone() *Freeze {
	return NewWithConfig(f.config.Clone())
}

// WithTheme creates a new Freeze instance with the specified theme
func (f *Freeze) WithTheme(theme string) *Freeze {
	clone := f.Clone()
	clone.config.SetTheme(theme)
	clone.generator = NewGenerator(clone.config)
	return clone
}

// WithFont creates a new Freeze instance with the specified font
func (f *Freeze) WithFont(family string, size float64) *Freeze {
	clone := f.Clone()
	clone.config.SetFont(family, size)
	clone.generator = NewGenerator(clone.config)
	return clone
}

// WithBackground creates a new Freeze instance with the specified background color
func (f *Freeze) WithBackground(color string) *Freeze {
	clone := f.Clone()
	clone.config.SetBackground(color)
	clone.generator = NewGenerator(clone.config)
	return clone
}

// WithWindow creates a new Freeze instance with window controls enabled/disabled
func (f *Freeze) WithWindow(enabled bool) *Freeze {
	clone := f.Clone()
	clone.config.SetWindow(enabled)
	clone.generator = NewGenerator(clone.config)
	return clone
}

// WithLineNumbers creates a new Freeze instance with line numbers enabled/disabled
func (f *Freeze) WithLineNumbers(enabled bool) *Freeze {
	clone := f.Clone()
	clone.config.SetLineNumbers(enabled)
	clone.generator = NewGenerator(clone.config)
	return clone
}

// WithShadow creates a new Freeze instance with shadow settings
func (f *Freeze) WithShadow(blur, x, y float64) *Freeze {
	clone := f.Clone()
	clone.config.SetShadow(blur, x, y)
	clone.generator = NewGenerator(clone.config)
	return clone
}

// WithBorder creates a new Freeze instance with border settings
func (f *Freeze) WithBorder(width, radius float64, color string) *Freeze {
	clone := f.Clone()
	clone.config.SetBorder(width, radius, color)
	clone.generator = NewGenerator(clone.config)
	return clone
}

// WithPadding creates a new Freeze instance with padding settings
func (f *Freeze) WithPadding(values ...float64) *Freeze {
	clone := f.Clone()
	clone.config.SetPadding(values...)
	clone.generator = NewGenerator(clone.config)
	return clone
}

// WithMargin creates a new Freeze instance with margin settings
func (f *Freeze) WithMargin(values ...float64) *Freeze {
	clone := f.Clone()
	clone.config.SetMargin(values...)
	clone.generator = NewGenerator(clone.config)
	return clone
}

// WithDimensions creates a new Freeze instance with specific dimensions
func (f *Freeze) WithDimensions(width, height float64) *Freeze {
	clone := f.Clone()
	clone.config.SetDimensions(width, height)
	clone.generator = NewGenerator(clone.config)
	return clone
}

// isPNGFile checks if the filename has a PNG extension
func isPNGFile(filename string) bool {
	return len(filename) > 4 && filename[len(filename)-4:] == ".png"
}

// Version information
const (
	Version = "1.0.0"
	Author  = "Charm"
)
