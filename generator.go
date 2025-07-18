package freezelib

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/landaiqing/freezelib/font"
	"github.com/landaiqing/freezelib/svg"
	"os"
	"strings"

	"github.com/alecthomas/chroma/v2"
	formatter "github.com/alecthomas/chroma/v2/formatters/svg"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/beevik/etree"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
	"github.com/charmbracelet/x/cellbuf"
	"github.com/kanrichan/resvg-go"
)

const (
	defaultFontSize   = 14.0
	defaultLineHeight = 1.2
)

// Generator handles the core screenshot generation logic
type Generator struct {
	config           *Config
	languageDetector *LanguageDetector
}

// NewGenerator creates a new generator with the given configuration
func NewGenerator(config *Config) *Generator {
	if config == nil {
		config = DefaultConfig()
	}
	return &Generator{
		config:           config,
		languageDetector: NewLanguageDetector(),
	}
}

// GenerateFromCode generates an SVG from source code
func (g *Generator) GenerateFromCode(code, language string) ([]byte, error) {
	if err := g.config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	// Set language if provided
	if language != "" {
		g.config.Language = language
	}

	// Get lexer for the language using enhanced detection
	lexer := g.languageDetector.GetLexer(g.config.Language, code)
	if lexer == nil {
		return nil, errors.New("could not determine language for syntax highlighting")
	}

	return g.generateSVG(code, lexer, false)
}

// GenerateFromFile generates an SVG from a source code file
func (g *Generator) GenerateFromFile(filename string) ([]byte, error) {
	if err := g.config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	// Read file content
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	code := string(content)

	// Get lexer from filename and content using enhanced detection
	lexer := g.languageDetector.GetLexerFromFile(filename, code)
	if lexer == nil {
		return nil, errors.New("could not determine language for syntax highlighting")
	}

	return g.generateSVG(code, lexer, false)
}

// DetectLanguage detects the programming language from code content
func (g *Generator) DetectLanguage(code string) string {
	return g.languageDetector.DetectLanguage(code)
}

// DetectLanguageFromFilename detects the programming language from filename
func (g *Generator) DetectLanguageFromFilename(filename string) string {
	return g.languageDetector.DetectLanguageFromFilename(filename)
}

// DetectLanguageFromFile detects language from both filename and content
func (g *Generator) DetectLanguageFromFile(filename, content string) string {
	return g.languageDetector.DetectLanguageFromFile(filename, content)
}

// GetSupportedLanguages returns a list of all supported languages
func (g *Generator) GetSupportedLanguages() []string {
	return g.languageDetector.GetSupportedLanguages()
}

// IsLanguageSupported checks if a language is supported
func (g *Generator) IsLanguageSupported(language string) bool {
	return g.languageDetector.IsLanguageSupported(language)
}

// SetLanguageDetector sets a custom language detector
func (g *Generator) SetLanguageDetector(detector *LanguageDetector) {
	g.languageDetector = detector
}

// GetLanguageDetector returns the current language detector
func (g *Generator) GetLanguageDetector() *LanguageDetector {
	return g.languageDetector
}

// GenerateFromANSI generates an SVG from ANSI terminal output
func (g *Generator) GenerateFromANSI(ansiOutput string) ([]byte, error) {
	if err := g.config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	// For ANSI output, we use a text lexer but handle ANSI sequences specially
	strippedInput := ansi.Strip(ansiOutput)
	it := chroma.Literator(chroma.Token{Type: chroma.Text, Value: strippedInput})

	return g.generateSVGFromIterator(ansiOutput, it, true)
}

// generateSVG is the core SVG generation function
func (g *Generator) generateSVG(input string, lexer chroma.Lexer, isAnsi bool) ([]byte, error) {
	// Create token iterator
	var it chroma.Iterator
	var err error
	if isAnsi {
		strippedInput := ansi.Strip(input)
		it = chroma.Literator(chroma.Token{Type: chroma.Text, Value: strippedInput})
	} else {
		it, err = chroma.Coalesce(lexer).Tokenise(nil, input)
		if err != nil {
			return nil, fmt.Errorf("could not tokenize input: %w", err)
		}
	}

	return g.generateSVGFromIterator(input, it, isAnsi)
}

// generateSVGFromIterator generates SVG from a token iterator
func (g *Generator) generateSVGFromIterator(input string, it chroma.Iterator, isAnsi bool) ([]byte, error) {
	config := g.config

	// Calculate scale factor
	scale := 1.0
	autoHeight := config.Height == 0
	autoWidth := config.Width == 0

	// Expand padding and margin
	expandedMargin := config.expandMargin(scale)
	expandedPadding := config.expandPadding(scale)

	// Process input based on line selection
	processedInput := input
	if len(config.Lines) == 2 {
		processedInput = cutLines(input, config.Lines)
	}

	// Handle text wrapping
	if config.Wrap > 0 {
		processedInput = cellbuf.Wrap(processedInput, config.Wrap, "")
	}

	// Get style
	style, ok := styles.Registry[strings.ToLower(config.Theme)]
	if !ok || style == nil {
		style = styles.Get("github") // fallback to github style
	}

	// Add background color to style if not present
	if !style.Has(chroma.Background) {
		var err error
		style, err = style.Builder().Add(chroma.Background, "bg:"+config.Background).Build()
		if err != nil {
			return nil, fmt.Errorf("could not add background: %w", err)
		}
	}

	// Get font options
	fontOptions, err := font.FontOptions(config.Font.Family, config.Font.Size, config.Font.Ligatures, config.Font.File)
	if err != nil {
		return nil, fmt.Errorf("invalid font options: %w", err)
	}

	// Create SVG formatter
	f := formatter.New(fontOptions...)

	// Format to SVG
	buf := &bytes.Buffer{}
	err = f.Format(buf, style, it)
	if err != nil {
		return nil, fmt.Errorf("could not format to SVG: %w", err)
	}

	// Parse SVG document
	doc := etree.NewDocument()
	_, err = doc.ReadFrom(buf)
	if err != nil {
		return nil, fmt.Errorf("could not parse SVG: %w", err)
	}

	elements := doc.ChildElements()
	if len(elements) < 1 {
		return nil, errors.New("invalid SVG output")
	}

	image := elements[0]

	// Calculate dimensions
	w, h := svg.GetDimensions(image)
	imageWidth := float64(w) * scale
	imageHeight := float64(h) * scale

	// Adjust for font size and line height
	imageHeight *= config.Font.Size / defaultFontSize
	imageHeight *= config.LineHeight / defaultLineHeight

	terminalWidth := imageWidth
	terminalHeight := imageHeight

	hPadding := expandedPadding[left] + expandedPadding[right]
	hMargin := expandedMargin[left] + expandedMargin[right]
	vMargin := expandedMargin[top] + expandedMargin[bottom]
	vPadding := expandedPadding[top] + expandedPadding[bottom]

	// Calculate final dimensions
	if !autoWidth {
		imageWidth = config.Width
		terminalWidth = config.Width - hMargin
	} else {
		imageWidth += hMargin + hPadding
		terminalWidth += hPadding
	}

	if !autoHeight {
		imageHeight = config.Height
		terminalHeight = config.Height - vMargin
	} else {
		imageHeight += vMargin + vPadding
		terminalHeight += vPadding
	}

	// Get terminal background element
	terminal := image.SelectElement("rect")
	if terminal == nil {
		return nil, errors.New("could not find terminal background element")
	}

	// Add window controls if enabled
	if config.Window {
		windowControls := svg.NewWindowControls(5.5*scale, 19.0*scale, 12.0*scale)
		svg.Move(windowControls, expandedMargin[left], expandedMargin[top])
		image.AddChild(windowControls)
		expandedPadding[top] += 15 * scale
	}

	// Add corner radius
	if config.Border.Radius > 0 {
		svg.AddCornerRadius(terminal, config.Border.Radius*scale)
	}

	// Add shadow
	if config.Shadow.Blur > 0 || config.Shadow.X > 0 || config.Shadow.Y > 0 {
		id := "shadow"
		svg.AddShadow(image, id, config.Shadow.X*scale, config.Shadow.Y*scale, config.Shadow.Blur*scale)
		terminal.CreateAttr("filter", fmt.Sprintf("url(#%s)", id))
	}

	// Process text elements
	textGroup := image.SelectElement("g")
	if textGroup != nil {
		textGroup.CreateAttr("font-size", fmt.Sprintf("%.2fpx", config.Font.Size*scale))
		textGroup.CreateAttr("clip-path", "url(#terminalMask)")
		text := textGroup.SelectElements("text")

		offsetLine := 0
		if len(config.Lines) > 0 {
			offsetLine = config.Lines[0]
		}

		lineHeight := config.LineHeight * scale

		for i, line := range text {
			if isAnsi {
				line.SetText("")
			}

			// Add line numbers if enabled
			if config.ShowLineNumbers {
				ln := etree.NewElement("tspan")
				ln.CreateAttr("xml:space", "preserve")
				ln.CreateAttr("fill", style.Get(chroma.LineNumbers).Colour.String())
				ln.SetText(fmt.Sprintf("%3d  ", i+1+offsetLine))
				line.InsertChildAt(0, ln)
			}

			// Position the line
			x := expandedPadding[left] + expandedMargin[left]
			y := (float64(i+1))*(config.Font.Size*lineHeight) + expandedPadding[top] + expandedMargin[top]

			svg.Move(line, x, y)

			// Remove lines that are outside the visible area
			if y > imageHeight-expandedMargin[bottom]-expandedPadding[bottom] {
				textGroup.RemoveChild(line)
			}
		}

		// Process ANSI sequences if needed
		if isAnsi {
			processANSI(processedInput, text, textGroup, config, scale)
		}
	}

	// Calculate auto width based on content
	if autoWidth {
		tabWidth := 4
		if isAnsi {
			tabWidth = 6
		}
		strippedInput := ansi.Strip(processedInput)
		longestLine := lipgloss.Width(strings.ReplaceAll(strippedInput, "\t", strings.Repeat(" ", tabWidth)))
		terminalWidth = float64(longestLine+1) * (config.Font.Size / font.GetFontHeightToWidthRatio())
		terminalWidth *= scale
		terminalWidth += hPadding
		imageWidth = terminalWidth + hMargin
	}

	// Add border
	if config.Border.Width > 0 {
		svg.AddOutline(terminal, config.Border.Width, config.Border.Color)
		terminalHeight -= config.Border.Width * 2
		terminalWidth -= config.Border.Width * 2
	}

	// Adjust for line numbers
	if config.ShowLineNumbers {
		if autoWidth {
			terminalWidth += config.Font.Size * 3 * scale
			imageWidth += config.Font.Size * 3 * scale
		} else {
			terminalWidth -= config.Font.Size * 3
		}
	}

	// Add clipping path if needed
	if !autoHeight || !autoWidth {
		svg.AddClipPath(image, "terminalMask",
			expandedMargin[left], expandedMargin[top],
			terminalWidth, terminalHeight-expandedPadding[bottom])
	}

	// Set final positions and dimensions
	svg.Move(terminal, max(expandedMargin[left], config.Border.Width/2), max(expandedMargin[top], config.Border.Width/2))
	svg.SetDimensions(image, imageWidth, imageHeight)
	svg.SetDimensions(terminal, terminalWidth, terminalHeight)

	// Convert to bytes
	return doc.WriteToBytes()
}

// ConvertToPNG converts SVG data to PNG format
func (g *Generator) ConvertToPNG(svgData []byte, width, height float64) ([]byte, error) {
	// Parse SVG document
	doc := etree.NewDocument()
	err := doc.ReadFromBytes(svgData)
	if err != nil {
		return nil, fmt.Errorf("could not parse SVG: %w", err)
	}

	// Use resvg for conversion
	worker, err := resvg.NewDefaultWorker(context.Background())
	if err != nil {
		return nil, fmt.Errorf("could not create resvg worker: %w", err)
	}
	defer worker.Close()

	fontdb, err := worker.NewFontDBDefault()
	if err != nil {
		return nil, fmt.Errorf("could not create font database: %w", err)
	}
	defer fontdb.Close()

	// Load embedded fonts
	if len(font.JetBrainsMonoTTF) > 0 {
		err = fontdb.LoadFontData(font.JetBrainsMonoTTF)
		if err != nil {
			return nil, fmt.Errorf("could not load JetBrains Mono font: %w", err)
		}
	}

	pixmap, err := worker.NewPixmap(uint32(width), uint32(height))
	if err != nil {
		return nil, fmt.Errorf("could not create pixmap: %w", err)
	}
	defer pixmap.Close()

	tree, err := worker.NewTreeFromData(svgData, &resvg.Options{
		Dpi:                192,
		ShapeRenderingMode: resvg.ShapeRenderingModeGeometricPrecision,
		TextRenderingMode:  resvg.TextRenderingModeOptimizeLegibility,
		ImageRenderingMode: resvg.ImageRenderingModeOptimizeQuality,
		DefaultSizeWidth:   float32(width),
		DefaultSizeHeight:  float32(height),
	})
	if err != nil {
		return nil, fmt.Errorf("could not create SVG tree: %w", err)
	}
	defer tree.Close()

	err = tree.ConvertText(fontdb)
	if err != nil {
		return nil, fmt.Errorf("could not convert text: %w", err)
	}

	err = tree.Render(resvg.TransformIdentity(), pixmap)
	if err != nil {
		return nil, fmt.Errorf("could not render SVG: %w", err)
	}

	pngData, err := pixmap.EncodePNG()
	if err != nil {
		return nil, fmt.Errorf("could not encode PNG: %w", err)
	}

	return pngData, nil
}

// cutLines cuts the input to the specified line range
func cutLines(input string, lines []int) string {
	if len(lines) != 2 {
		return input
	}

	inputLines := strings.Split(input, "\n")
	start := lines[0]
	end := lines[1]

	if start < 0 {
		start = 0
	}
	if end >= len(inputLines) || end < 0 {
		end = len(inputLines) - 1
	}
	if start > end {
		return ""
	}

	return strings.Join(inputLines[start:end+1], "\n")
}

// max returns the maximum of two float64 values
func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
