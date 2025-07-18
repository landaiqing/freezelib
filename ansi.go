package freezelib

import (
	"fmt"
	"strings"

	"github.com/beevik/etree"
	"github.com/charmbracelet/x/ansi"
	"github.com/mattn/go-runewidth"
)

// dispatcher handles ANSI escape sequences and converts them to SVG
type dispatcher struct {
	lines   []*etree.Element
	svg     *etree.Element
	config  *Config
	scale   float64
	row     int
	col     int
	bg      *etree.Element
	bgWidth int
}

// newDispatcher creates a new ANSI dispatcher
func newDispatcher(lines []*etree.Element, svg *etree.Element, config *Config, scale float64) *dispatcher {
	return &dispatcher{
		lines:  lines,
		svg:    svg,
		config: config,
		scale:  scale,
		row:    0,
		col:    0,
	}
}

// Print handles printable characters
func (p *dispatcher) Print(r rune) {
	p.row = clamp(p.row, 0, len(p.lines)-1)
	// insert the rune in the last tspan
	children := p.lines[p.row].ChildElements()
	var lastChild *etree.Element
	isFirstChild := len(children) == 0
	if isFirstChild {
		lastChild = etree.NewElement("tspan")
		lastChild.CreateAttr("xml:space", "preserve")
		p.lines[p.row].AddChild(lastChild)
	} else {
		lastChild = children[len(children)-1]
	}

	if runewidth.RuneWidth(r) > 1 {
		newChild := lastChild.Copy()
		newChild.SetText(string(r))
		newChild.CreateAttr("dx", fmt.Sprintf("%.2fpx", (p.config.Font.Size/5)*p.scale))
		p.lines[p.row].AddChild(newChild)
	} else {
		lastChild.SetText(lastChild.Text() + string(r))
	}

	p.col += runewidth.RuneWidth(r)
	if p.bg != nil {
		p.bgWidth += runewidth.RuneWidth(r)
	}
}

// Execute handles control characters
func (p *dispatcher) Execute(code byte) {
	if code == '\t' {
		for p.col%16 != 0 {
			p.Print(' ')
		}
	}
	if code == '\n' {
		p.endBackground()
		p.row++
		p.col = 0
	}
}

// endBackground ends the current background span
func (p *dispatcher) endBackground() {
	if p.bg == nil {
		return
	}
	p.bg.CreateAttr("width", fmt.Sprintf("%.2fpx", float64(p.bgWidth)*(p.config.Font.Size/fontHeightToWidthRatio)*p.scale))
	p.bg = nil
	p.bgWidth = 0
}

// CsiDispatch handles CSI (Control Sequence Introducer) sequences
func (p *dispatcher) CsiDispatch(cmd ansi.Cmd, params ansi.Params) {
	if cmd != 'm' {
		// ignore incomplete or non Style (SGR) sequences
		return
	}

	span := etree.NewElement("tspan")
	span.CreateAttr("xml:space", "preserve")
	reset := func() {
		// reset ANSI, this is done by creating a new empty tspan,
		// which would reset all the styles such that when text is appended to the last
		// child of this line there is no styling applied.
		if p.row < len(p.lines) {
			p.lines[p.row].AddChild(span)
		}
		p.endBackground()
	}

	if len(params) == 0 {
		// zero params means reset
		reset()
		return
	}

	var i int
	for i < len(params) {
		v := params[i].Param(0)
		switch v {
		case 0:
			reset()
		case 1:
			// Bold - not implemented in SVG for now
			p.lines[p.row].AddChild(span)
		case 9:
			span.CreateAttr("text-decoration", "line-through")
			p.lines[p.row].AddChild(span)
		case 3:
			span.CreateAttr("font-style", "italic")
			p.lines[p.row].AddChild(span)
		case 4:
			span.CreateAttr("text-decoration", "underline")
			p.lines[p.row].AddChild(span)
		case 30, 31, 32, 33, 34, 35, 36, 37, 90, 91, 92, 93, 94, 95, 96, 97:
			span.CreateAttr("fill", ansiPalette[v])
			p.lines[p.row].AddChild(span)
		case 38:
			i++
			if i < len(params) {
				switch params[i].Param(0) {
				case 5:
					if i+1 < len(params) {
						n := params[i+1].Param(0)
						i++
						fill := palette[n]
						span.CreateAttr("fill", fill)
						p.lines[p.row].AddChild(span)
					}
				case 2:
					if i+3 < len(params) {
						r := params[i+1].Param(0)
						g := params[i+2].Param(0)
						b := params[i+3].Param(0)
						i += 3
						fill := fmt.Sprintf("rgb(%d,%d,%d)", r, g, b)
						span.CreateAttr("fill", fill)
						p.lines[p.row].AddChild(span)
					}
				}
			}
		case 40, 41, 42, 43, 44, 45, 46, 47, 100, 101, 102, 103, 104, 105, 106, 107:
			// Background colors
			p.endBackground()
			p.bg = etree.NewElement("rect")
			p.bg.CreateAttr("fill", ansiPalette[v-10])
			p.bg.CreateAttr("height", fmt.Sprintf("%.2fpx", p.config.Font.Size*p.config.LineHeight))
			p.bg.CreateAttr("x", fmt.Sprintf("%.2fpx", float64(p.col)*(p.config.Font.Size/fontHeightToWidthRatio)*p.scale))
			p.bg.CreateAttr("y", fmt.Sprintf("%.2fpx", float64(p.row)*p.config.Font.Size*p.config.LineHeight))
			p.svg.InsertChildAt(0, p.bg)
		case 48:
			i++
			if i < len(params) {
				switch params[i].Param(0) {
				case 5:
					if i+1 < len(params) {
						n := params[i+1].Param(0)
						i++
						p.endBackground()
						p.bg = etree.NewElement("rect")
						p.bg.CreateAttr("fill", palette[n])
						p.bg.CreateAttr("height", fmt.Sprintf("%.2fpx", p.config.Font.Size*p.config.LineHeight))
						p.bg.CreateAttr("x", fmt.Sprintf("%.2fpx", float64(p.col)*(p.config.Font.Size/fontHeightToWidthRatio)*p.scale))
						p.bg.CreateAttr("y", fmt.Sprintf("%.2fpx", float64(p.row)*p.config.Font.Size*p.config.LineHeight))
						p.svg.InsertChildAt(0, p.bg)
					}
				case 2:
					if i+3 < len(params) {
						r := params[i+1].Param(0)
						g := params[i+2].Param(0)
						b := params[i+3].Param(0)
						i += 3
						p.endBackground()
						p.bg = etree.NewElement("rect")
						p.bg.CreateAttr("fill", fmt.Sprintf("rgb(%d,%d,%d)", r, g, b))
						p.bg.CreateAttr("height", fmt.Sprintf("%.2fpx", p.config.Font.Size*p.config.LineHeight))
						p.bg.CreateAttr("x", fmt.Sprintf("%.2fpx", float64(p.col)*(p.config.Font.Size/fontHeightToWidthRatio)*p.scale))
						p.bg.CreateAttr("y", fmt.Sprintf("%.2fpx", float64(p.row)*p.config.Font.Size*p.config.LineHeight))
						p.svg.InsertChildAt(0, p.bg)
					}
				}
			}
		}
		i++
	}
}

// processANSI processes ANSI escape sequences in the input text
func processANSI(input string, lines []*etree.Element, svg *etree.Element, config *Config, scale float64) {
	d := newDispatcher(lines, svg, config, scale)
	parser := ansi.NewParser()
	parser.SetHandler(ansi.Handler{
		Print:     d.Print,
		HandleCsi: d.CsiDispatch,
		Execute:   d.Execute,
	})

	for _, line := range strings.Split(input, "\n") {
		parser.Parse([]byte(line))
		d.Execute(ansi.LF) // simulate a newline
	}
}

// stripANSI removes ANSI escape sequences from text
func stripANSI(input string) string {
	return ansi.Strip(input)
}

// isANSI checks if the input contains ANSI escape sequences
func isANSI(input string) bool {
	return stripANSI(input) != input
}

// clamp constrains a value between min and max
func clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

const fontHeightToWidthRatio = 1.68

// ANSI color palette
var ansiPalette = map[int]string{
	30: "#000000", // black
	31: "#FF0000", // red
	32: "#00FF00", // green
	33: "#FFFF00", // yellow
	34: "#0000FF", // blue
	35: "#FF00FF", // magenta
	36: "#00FFFF", // cyan
	37: "#FFFFFF", // white
	90: "#808080", // bright black (gray)
	91: "#FF8080", // bright red
	92: "#80FF80", // bright green
	93: "#FFFF80", // bright yellow
	94: "#8080FF", // bright blue
	95: "#FF80FF", // bright magenta
	96: "#80FFFF", // bright cyan
	97: "#FFFFFF", // bright white
}

// 256-color palette
var palette = []string{
	"#000000", "#800000", "#008000", "#808000", "#000080", "#800080", "#008080", "#c0c0c0",
	"#808080", "#ff0000", "#00ff00", "#ffff00", "#0000ff", "#ff00ff", "#00ffff", "#ffffff",
	"#000000", "#00005f", "#000087", "#0000af", "#0000d7", "#0000ff", "#005f00", "#005f5f",
	"#005f87", "#005faf", "#005fd7", "#005fff", "#008700", "#00875f", "#008787", "#0087af",
	"#0087d7", "#0087ff", "#00af00", "#00af5f", "#00af87", "#00afaf", "#00afd7", "#00afff",
	"#00d700", "#00d75f", "#00d787", "#00d7af", "#00d7d7", "#00d7ff", "#00ff00", "#00ff5f",
	"#00ff87", "#00ffaf", "#00ffd7", "#00ffff", "#5f0000", "#5f005f", "#5f0087", "#5f00af",
	"#5f00d7", "#5f00ff", "#5f5f00", "#5f5f5f", "#5f5f87", "#5f5faf", "#5f5fd7", "#5f5fff",
	"#5f8700", "#5f875f", "#5f8787", "#5f87af", "#5f87d7", "#5f87ff", "#5faf00", "#5faf5f",
	"#5faf87", "#5fafaf", "#5fafd7", "#5fafff", "#5fd700", "#5fd75f", "#5fd787", "#5fd7af",
	"#5fd7d7", "#5fd7ff", "#5fff00", "#5fff5f", "#5fff87", "#5fffaf", "#5fffd7", "#5fffff",
	"#870000", "#87005f", "#870087", "#8700af", "#8700d7", "#8700ff", "#875f00", "#875f5f",
	"#875f87", "#875faf", "#875fd7", "#875fff", "#878700", "#87875f", "#878787", "#8787af",
	"#8787d7", "#8787ff", "#87af00", "#87af5f", "#87af87", "#87afaf", "#87afd7", "#87afff",
	"#87d700", "#87d75f", "#87d787", "#87d7af", "#87d7d7", "#87d7ff", "#87ff00", "#87ff5f",
	"#87ff87", "#87ffaf", "#87ffd7", "#87ffff", "#af0000", "#af005f", "#af0087", "#af00af",
	"#af00d7", "#af00ff", "#af5f00", "#af5f5f", "#af5f87", "#af5faf", "#af5fd7", "#af5fff",
	"#af8700", "#af875f", "#af8787", "#af87af", "#af87d7", "#af87ff", "#afaf00", "#afaf5f",
	"#afaf87", "#afafaf", "#afafd7", "#afafff", "#afd700", "#afd75f", "#afd787", "#afd7af",
	"#afd7d7", "#afd7ff", "#afff00", "#afff5f", "#afff87", "#afffaf", "#afffd7", "#afffff",
	"#d70000", "#d7005f", "#d70087", "#d700af", "#d700d7", "#d700ff", "#d75f00", "#d75f5f",
	"#d75f87", "#d75faf", "#d75fd7", "#d75fff", "#d78700", "#d7875f", "#d78787", "#d787af",
	"#d787d7", "#d787ff", "#d7af00", "#d7af5f", "#d7af87", "#d7afaf", "#d7afd7", "#d7afff",
	"#d7d700", "#d7d75f", "#d7d787", "#d7d7af", "#d7d7d7", "#d7d7ff", "#d7ff00", "#d7ff5f",
	"#d7ff87", "#d7ffaf", "#d7ffd7", "#d7ffff", "#ff0000", "#ff005f", "#ff0087", "#ff00af",
	"#ff00d7", "#ff00ff", "#ff5f00", "#ff5f5f", "#ff5f87", "#ff5faf", "#ff5fd7", "#ff5fff",
	"#ff8700", "#ff875f", "#ff8787", "#ff87af", "#ff87d7", "#ff87ff", "#ffaf00", "#ffaf5f",
	"#ffaf87", "#ffafaf", "#ffafd7", "#ffafff", "#ffd700", "#ffd75f", "#ffd787", "#ffd7af",
	"#ffd7d7", "#ffd7ff", "#ffff00", "#ffff5f", "#ffff87", "#ffffaf", "#ffffd7", "#ffffff",
	"#080808", "#121212", "#1c1c1c", "#262626", "#303030", "#3a3a3a", "#444444", "#4e4e4e",
	"#585858", "#626262", "#6c6c6c", "#767676", "#808080", "#8a8a8a", "#949494", "#9e9e9e",
	"#a8a8a8", "#b2b2b2", "#bcbcbc", "#c6c6c6", "#d0d0d0", "#dadada", "#e4e4e4", "#eeeeee",
}
