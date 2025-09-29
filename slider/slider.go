package slider

import (
	"crypto/rand"
	_ "embed"
	"encoding/hex"
	"strconv"

	"github.com/plainkit/html"
)

type Props struct {
	ID    string
	Class string
	Attrs []html.Global
}

type InputProps struct {
	ID       string
	Class    string
	Attrs    []html.Global
	Name     string
	Min      int
	Max      int
	Step     int
	Value    int
	Disabled bool
}

type ValueProps struct {
	ID    string
	Class string
	Attrs []html.Global
	For   string
}

func divArgsFromProps(baseClass string, extra ...string) func(p Props) []html.DivArg {
	return func(p Props) []html.DivArg {
		args := []html.DivArg{html.AClass(html.ClassMerge(append([]string{baseClass}, append(extra, p.Class)...)...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p Props) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	args := divArgsFromProps("w-full")(p)
	args = append(args, html.AData("pui-slider-wrapper", ""))

	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

func (p InputProps) ApplyInput(attrs *html.InputAttrs, children *[]html.Component) {
	id := p.ID
	if id == "" {
		id = randomID("slider")
	}

	args := []html.InputArg{
		html.AId(id),
		html.AType("range"),
		html.AClass(html.ClassMerge(
			"w-full h-2 rounded-full bg-secondary appearance-none cursor-pointer",
			"focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2",
			"[&::-webkit-slider-thumb]:appearance-none [&::-webkit-slider-thumb]:w-4 [&::-webkit-slider-thumb]:h-4",
			"[&::-webkit-slider-thumb]:rounded-full [&::-webkit-slider-thumb]:bg-primary",
			"[&::-webkit-slider-thumb]:hover:bg-primary/90",
			"[&::-moz-range-thumb]:w-4 [&::-moz-range-thumb]:h-4 [&::-moz-range-thumb]:border-0",
			"[&::-moz-range-thumb]:rounded-full [&::-moz-range-thumb]:bg-primary",
			"[&::-moz-range-thumb]:hover:bg-primary/90",
			"disabled:opacity-50 disabled:cursor-not-allowed",
			p.Class,
		)),
		html.AData("pui-slider-input", ""),
	}
	if p.Name != "" {
		args = append(args, html.AName(p.Name))
	}

	if p.Value != 0 {
		args = append(args, html.AValue(strconv.Itoa(p.Value)))
	}

	if p.Min != 0 {
		args = append(args, html.AMin(strconv.Itoa(p.Min)))
	}

	if p.Max != 0 {
		args = append(args, html.AMax(strconv.Itoa(p.Max)))
	}

	if p.Step != 0 {
		args = append(args, html.AStep(strconv.Itoa(p.Step)))
	}

	if p.Disabled {
		args = append(args, html.ADisabled())
	}

	for _, a := range p.Attrs {
		args = append(args, a)
	}

	for _, a := range args {
		a.ApplyInput(attrs, children)
	}
}

func (p ValueProps) ApplySpan(attrs *html.SpanAttrs, children *[]html.Component) {
	if p.For == "" {
		// Apply error styling
		args := []html.SpanArg{html.AClass("text-xs text-destructive")}
		for _, a := range args {
			a.ApplySpan(attrs, children)
		}

		return
	}

	args := []html.SpanArg{
		html.AClass(html.ClassMerge("text-sm text-muted-foreground", p.Class)),
		html.AData("pui-slider-value", ""),
		html.AData("pui-slider-value-for", p.For),
	}
	if p.ID != "" {
		args = append(args, html.AId(p.ID))
	}

	for _, a := range p.Attrs {
		args = append(args, a)
	}

	for _, a := range args {
		a.ApplySpan(attrs, children)
	}
}

// Slider is the root container around slider inputs/values.
func Slider(args ...html.DivArg) html.Node {
	var (
		props Props
		rest  []html.DivArg
	)

	for _, a := range args {
		if v, ok := a.(Props); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Div(append([]html.DivArg{props}, rest...)...).WithAssets("", sliderJS, "ui-slider")
}

// Input renders the range input element used within the slider wrapper.
func Input(args ...html.InputArg) html.Node {
	var (
		props InputProps
		rest  []html.InputArg
	)

	for _, a := range args {
		if v, ok := a.(InputProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Input(append([]html.InputArg{props}, rest...)...)
}

// Value renders a span that mirrors the slider value using data attributes.
func Value(args ...html.SpanArg) html.Node {
	var (
		props ValueProps
		rest  []html.SpanArg
	)

	for _, a := range args {
		if v, ok := a.(ValueProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Span(append([]html.SpanArg{props}, rest...)...)
}

func randomID(prefix string) string {
	buf := make([]byte, 4)
	if _, err := rand.Read(buf); err != nil {
		return prefix + "-id"
	}

	return prefix + "-" + hex.EncodeToString(buf)
}

//go:embed slider.js
var sliderJS string
