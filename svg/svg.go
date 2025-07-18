package svg

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/beevik/etree"
)

// AddShadow adds a definition of a shadow to the <defs> with the given id.
func AddShadow(element *etree.Element, id string, x, y, blur float64) {
	f := etree.NewElement("filter")
	f.CreateAttr("id", id)
	f.CreateAttr("filterUnits", "userSpaceOnUse")

	b := etree.NewElement("feGaussianBlur")
	b.CreateAttr("in", "SourceAlpha")
	b.CreateAttr("stdDeviation", fmt.Sprintf("%.2f", blur))

	o := etree.NewElement("feOffset")
	o.CreateAttr("result", "offsetblur")
	o.CreateAttr("dx", fmt.Sprintf("%.2f", x))
	o.CreateAttr("dy", fmt.Sprintf("%.2f", y))

	m := etree.NewElement("feMerge")
	mn1 := etree.NewElement("feMergeNode")
	mn2 := etree.NewElement("feMergeNode")
	mn2.CreateAttr("in", "SourceGraphic")
	m.AddChild(mn1)
	m.AddChild(mn2)

	f.AddChild(b)
	f.AddChild(o)
	f.AddChild(m)

	defs := etree.NewElement("defs")
	defs.AddChild(f)
	element.AddChild(defs)
}

// AddClipPath adds a definition of a clip path to the <defs> with the given id.
func AddClipPath(element *etree.Element, id string, x, y, w, h float64) {
	p := etree.NewElement("clipPath")
	p.CreateAttr("id", id)

	rect := etree.NewElement("rect")
	rect.CreateAttr("x", fmt.Sprintf("%.2f", x))
	rect.CreateAttr("y", fmt.Sprintf("%.2f", y))
	rect.CreateAttr("width", fmt.Sprintf("%.2f", w))
	rect.CreateAttr("height", fmt.Sprintf("%.2f", h))

	p.AddChild(rect)

	defs := etree.NewElement("defs")
	defs.AddChild(p)
	element.AddChild(defs)
}

// AddCornerRadius adds corner radius to an element.
func AddCornerRadius(e *etree.Element, radius float64) {
	e.CreateAttr("rx", fmt.Sprintf("%.2f", radius))
	e.CreateAttr("ry", fmt.Sprintf("%.2f", radius))
}

// Move moves the given element to the (x, y) position.
func Move(e *etree.Element, x, y float64) {
	e.CreateAttr("x", fmt.Sprintf("%.2fpx", x))
	e.CreateAttr("y", fmt.Sprintf("%.2fpx", y))
}

// AddOutline adds an outline to the given element.
func AddOutline(e *etree.Element, width float64, color string) {
	e.CreateAttr("stroke", color)
	e.CreateAttr("stroke-width", fmt.Sprintf("%.2f", width))
}

const (
	red    string = "#FF5A54"
	yellow string = "#E6BF29"
	green  string = "#52C12B"
)

// NewWindowControls returns a colorful window bar element.
func NewWindowControls(r float64, x, y float64) *etree.Element {
	bar := etree.NewElement("svg")
	for i, color := range []string{red, yellow, green} {
		circle := etree.NewElement("circle")
		circle.CreateAttr("cx", fmt.Sprintf("%.2f", float64(i+1)*float64(x)-float64(r)))
		circle.CreateAttr("cy", fmt.Sprintf("%.2f", y))
		circle.CreateAttr("r", fmt.Sprintf("%.2f", r))
		circle.CreateAttr("fill", color)
		bar.AddChild(circle)
	}
	return bar
}

// SetDimensions sets the width and height of the given element.
func SetDimensions(element *etree.Element, width, height float64) {
	widthAttr := element.SelectAttr("width")
	heightAttr := element.SelectAttr("height")
	if heightAttr != nil {
		heightAttr.Value = fmt.Sprintf("%.2f", height)
	}
	if widthAttr != nil {
		widthAttr.Value = fmt.Sprintf("%.2f", width)
	}
}

// GetDimensions returns the width and height of the element.
func GetDimensions(element *etree.Element) (int, int) {
	widthValue := element.SelectAttrValue("width", "0px")
	heightValue := element.SelectAttrValue("height", "0px")
	width := dimensionToInt(widthValue)
	height := dimensionToInt(heightValue)
	return width, height
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

// CreateSVGElement creates a new SVG root element with basic attributes
func CreateSVGElement(width, height float64) *etree.Element {
	svg := etree.NewElement("svg")
	svg.CreateAttr("xmlns", "http://www.w3.org/2000/svg")
	svg.CreateAttr("width", fmt.Sprintf("%.2f", width))
	svg.CreateAttr("height", fmt.Sprintf("%.2f", height))
	svg.CreateAttr("viewBox", fmt.Sprintf("0 0 %.2f %.2f", width, height))
	return svg
}

// CreateRect creates a rectangle element
func CreateRect(x, y, width, height float64, fill string) *etree.Element {
	rect := etree.NewElement("rect")
	rect.CreateAttr("x", fmt.Sprintf("%.2f", x))
	rect.CreateAttr("y", fmt.Sprintf("%.2f", y))
	rect.CreateAttr("width", fmt.Sprintf("%.2f", width))
	rect.CreateAttr("height", fmt.Sprintf("%.2f", height))
	if fill != "" {
		rect.CreateAttr("fill", fill)
	}
	return rect
}

// CreateText creates a text element
func CreateText(x, y float64, content string) *etree.Element {
	text := etree.NewElement("text")
	text.CreateAttr("x", fmt.Sprintf("%.2f", x))
	text.CreateAttr("y", fmt.Sprintf("%.2f", y))
	text.SetText(content)
	return text
}

// CreateGroup creates a group element
func CreateGroup() *etree.Element {
	return etree.NewElement("g")
}

// SetFontAttributes sets font-related attributes on an element
func SetFontAttributes(element *etree.Element, family string, size float64) {
	if family != "" {
		element.CreateAttr("font-family", family)
	}
	if size > 0 {
		element.CreateAttr("font-size", fmt.Sprintf("%.2fpx", size))
	}
}

// SetTextAttributes sets text-related attributes
func SetTextAttributes(element *etree.Element, fill, textAnchor string) {
	if fill != "" {
		element.CreateAttr("fill", fill)
	}
	if textAnchor != "" {
		element.CreateAttr("text-anchor", textAnchor)
	}
}

// AddStyle adds a style attribute to an element
func AddStyle(element *etree.Element, style string) {
	existing := element.SelectAttrValue("style", "")
	if existing != "" {
		style = existing + "; " + style
	}
	element.CreateAttr("style", style)
}

// Max returns the maximum of two float64 values
func Max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// Min returns the minimum of two float64 values
func Min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// Clamp constrains a value between min and max
func Clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
