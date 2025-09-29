package slider

import (
	"crypto/rand"
	_ "embed"
	"encoding/hex"
	"strconv"

	"github.com/plainkit/html"
	"github.com/plainkit/ui/internal/classnames"
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

// Slider is the root container around slider inputs/values.
func Slider(props Props, args ...html.DivArg) html.Node {
	divArgs := []html.DivArg{
		html.AClass(classnames.Merge("w-full", props.Class)),
		html.AData("pui-slider-wrapper", ""),
	}
	if props.ID != "" {
		divArgs = append(divArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}
	divArgs = append(divArgs, args...)

	node := html.Div(divArgs...)
	node = node.WithAssets("", sliderJS, "ui-slider")
	return node
}

// Input renders the range input element used within the slider wrapper.
func Input(props InputProps) html.Node {
	if props.ID == "" {
		props.ID = randomID("slider")
	}

	inputArgs := []html.InputArg{
		html.AId(props.ID),
		html.AType("range"),
		html.AClass(classnames.Merge(
			"w-full h-2 rounded-full bg-secondary appearance-none cursor-pointer",
			"focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2",
			"[&::-webkit-slider-thumb]:appearance-none [&::-webkit-slider-thumb]:w-4 [&::-webkit-slider-thumb]:h-4",
			"[&::-webkit-slider-thumb]:rounded-full [&::-webkit-slider-thumb]:bg-primary",
			"[&::-webkit-slider-thumb]:hover:bg-primary/90",
			"[&::-moz-range-thumb]:w-4 [&::-moz-range-thumb]:h-4 [&::-moz-range-thumb]:border-0",
			"[&::-moz-range-thumb]:rounded-full [&::-moz-range-thumb]:bg-primary",
			"[&::-moz-range-thumb]:hover:bg-primary/90",
			"disabled:opacity-50 disabled:cursor-not-allowed",
			props.Class,
		)),
		html.AData("pui-slider-input", ""),
	}
	if props.Name != "" {
		inputArgs = append(inputArgs, html.AName(props.Name))
	}
	if props.Value != 0 {
		inputArgs = append(inputArgs, html.AValue(strconv.Itoa(props.Value)))
	}
	if props.Min != 0 {
		inputArgs = append(inputArgs, html.AMin(strconv.Itoa(props.Min)))
	}
	if props.Max != 0 {
		inputArgs = append(inputArgs, html.AMax(strconv.Itoa(props.Max)))
	}
	if props.Step != 0 {
		inputArgs = append(inputArgs, html.AStep(strconv.Itoa(props.Step)))
	}
	if props.Disabled {
		inputArgs = append(inputArgs, html.ADisabled())
	}
	for _, attr := range props.Attrs {
		inputArgs = append(inputArgs, attr)
	}

	return html.Input(inputArgs...)
}

// Value renders a span that mirrors the slider value using data attributes.
func Value(props ValueProps, args ...html.SpanArg) html.Node {
	if props.For == "" {
		return html.Span(html.AClass("text-xs text-destructive"), html.Text("Slider value requires 'For'"))
	}

	spanArgs := []html.SpanArg{
		html.AClass(classnames.Merge("text-sm text-muted-foreground", props.Class)),
		html.AData("pui-slider-value", ""),
		html.AData("pui-slider-value-for", props.For),
	}
	if props.ID != "" {
		spanArgs = append(spanArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		spanArgs = append(spanArgs, attr)
	}
	spanArgs = append(spanArgs, args...)

	return html.Span(spanArgs...)
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
