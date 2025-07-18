package freezelib

import (
	"fmt"
	"strconv"
	"strings"
)

// Config represents the configuration for generating code screenshots
type Config struct {
	// Window settings
	Background string    `json:"background"`
	Margin     []float64 `json:"margin"`
	Padding    []float64 `json:"padding"`
	Window     bool      `json:"window"`
	Width      float64   `json:"width"`
	Height     float64   `json:"height"`

	// Language and theme
	Language string `json:"language"`
	Theme    string `json:"theme"`
	Wrap     int    `json:"wrap"`

	// Decoration
	Border Border `json:"border"`
	Shadow Shadow `json:"shadow"`

	// Font
	Font Font `json:"font"`

	// Line settings
	LineHeight      float64 `json:"line_height"`
	Lines           []int   `json:"lines"`
	ShowLineNumbers bool    `json:"show_line_numbers"`
}

// Shadow configuration for drop shadow effects
type Shadow struct {
	Blur float64 `json:"blur"`
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
}

// Border configuration for window borders
type Border struct {
	Radius float64 `json:"radius"`
	Width  float64 `json:"width"`
	Color  string  `json:"color"`
}

// Font configuration
type Font struct {
	Family    string  `json:"family"`
	File      string  `json:"file"`
	Size      float64 `json:"size"`
	Ligatures bool    `json:"ligatures"`
}

// DefaultConfig returns a default configuration
func DefaultConfig() *Config {
	return &Config{
		Background:      "#171717",
		Margin:          []float64{0},
		Padding:         []float64{20},
		Window:          false,
		Width:           0,
		Height:          0,
		Language:        "",
		Theme:           "charm",
		Wrap:            0,
		Border:          Border{Radius: 0, Width: 0, Color: "#515151"},
		Shadow:          Shadow{Blur: 0, X: 0, Y: 0},
		Font:            Font{Family: "JetBrains Mono", Size: 14, Ligatures: true},
		LineHeight:      1.2,
		Lines:           []int{},
		ShowLineNumbers: false,
	}
}

// SetPadding sets padding for all sides or specific sides
// Accepts 1, 2, or 4 values like CSS padding
func (c *Config) SetPadding(values ...float64) *Config {
	c.Padding = values
	return c
}

// SetMargin sets margin for all sides or specific sides
// Accepts 1, 2, or 4 values like CSS margin
func (c *Config) SetMargin(values ...float64) *Config {
	c.Margin = values
	return c
}

// SetFont sets font family and size
func (c *Config) SetFont(family string, size float64) *Config {
	c.Font.Family = family
	c.Font.Size = size
	return c
}

// SetTheme sets the syntax highlighting theme
func (c *Config) SetTheme(theme string) *Config {
	c.Theme = theme
	return c
}

// SetLanguage sets the programming language for syntax highlighting
func (c *Config) SetLanguage(language string) *Config {
	c.Language = language
	return c
}

// SetBackground sets the background color
func (c *Config) SetBackground(color string) *Config {
	c.Background = color
	return c
}

// SetWindow enables or disables window controls
func (c *Config) SetWindow(enabled bool) *Config {
	c.Window = enabled
	return c
}

// SetLineNumbers enables or disables line numbers
func (c *Config) SetLineNumbers(enabled bool) *Config {
	c.ShowLineNumbers = enabled
	return c
}

// SetShadow sets shadow properties
func (c *Config) SetShadow(blur, x, y float64) *Config {
	c.Shadow = Shadow{Blur: blur, X: x, Y: y}
	return c
}

// SetBorder sets border properties
func (c *Config) SetBorder(width, radius float64, color string) *Config {
	c.Border = Border{Width: width, Radius: radius, Color: color}
	return c
}

// SetDimensions sets the output dimensions
func (c *Config) SetDimensions(width, height float64) *Config {
	c.Width = width
	c.Height = height
	return c
}

// SetLines sets the line range to capture (1-indexed)
func (c *Config) SetLines(start, end int) *Config {
	if start > 0 && end > 0 && start <= end {
		c.Lines = []int{start - 1, end - 1} // Convert to 0-indexed
	}
	return c
}

// expandPadding expands padding values according to CSS rules
func (c *Config) expandPadding(scale float64) []float64 {
	p := c.Padding
	switch len(p) {
	case 1:
		return []float64{p[0] * scale, p[0] * scale, p[0] * scale, p[0] * scale}
	case 2:
		return []float64{p[0] * scale, p[1] * scale, p[0] * scale, p[1] * scale}
	case 4:
		return []float64{p[0] * scale, p[1] * scale, p[2] * scale, p[3] * scale}
	default:
		return []float64{0, 0, 0, 0}
	}
}

// expandMargin expands margin values according to CSS rules
func (c *Config) expandMargin(scale float64) []float64 {
	m := c.Margin
	switch len(m) {
	case 1:
		return []float64{m[0] * scale, m[0] * scale, m[0] * scale, m[0] * scale}
	case 2:
		return []float64{m[0] * scale, m[1] * scale, m[0] * scale, m[1] * scale}
	case 4:
		return []float64{m[0] * scale, m[1] * scale, m[2] * scale, m[3] * scale}
	default:
		return []float64{0, 0, 0, 0}
	}
}

// Clone creates a deep copy of the configuration
func (c *Config) Clone() *Config {
	clone := *c
	clone.Margin = make([]float64, len(c.Margin))
	copy(clone.Margin, c.Margin)
	clone.Padding = make([]float64, len(c.Padding))
	copy(clone.Padding, c.Padding)
	clone.Lines = make([]int, len(c.Lines))
	copy(clone.Lines, c.Lines)
	return &clone
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.Font.Size <= 0 {
		return fmt.Errorf("font size must be positive")
	}
	if c.LineHeight <= 0 {
		return fmt.Errorf("line height must be positive")
	}
	if len(c.Lines) == 2 && c.Lines[0] > c.Lines[1] {
		return fmt.Errorf("start line must be less than or equal to end line")
	}
	return nil
}

// parseColor validates and normalizes color values
func parseColor(color string) string {
	color = strings.TrimSpace(color)
	if color == "" {
		return "#000000"
	}
	if !strings.HasPrefix(color, "#") {
		color = "#" + color
	}
	return color
}

// dimensionToInt converts dimension strings to integers
func dimensionToInt(dimension string) int {
	dimension = strings.TrimSuffix(dimension, "px")
	val, err := strconv.Atoi(dimension)
	if err != nil {
		return 0
	}
	return val
}

// side constants for padding/margin indexing
const (
	top    = 0
	right  = 1
	bottom = 2
	left   = 3
)
