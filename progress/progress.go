package progress

import (
	"crypto/rand"
	"encoding/hex"
	"strconv"

	"github.com/plainkit/html"
	"github.com/plainkit/ui/internal/classnames"
)

type Size string
type Variant string

type Props struct {
	ID        string
	Class     string
	Attrs     []html.Global
	Max       int
	Value     int
	Label     string
	ShowValue bool
	Size      Size
	Variant   Variant
	BarClass  string
}

const (
	SizeSm Size = "sm"
	SizeLg Size = "lg"
)

const (
	VariantDefault Variant = "default"
	VariantSuccess Variant = "success"
	VariantDanger  Variant = "danger"
	VariantWarning Variant = "warning"
)

func Progress(props Props) html.Node {
	propsMax := maxValue(props.Max)
	if props.ID == "" {
		props.ID = randomID()
	}

	outerClass := classnames.Merge("w-full", props.Class)

	outerArgs := []html.DivArg{
		html.AId(props.ID),
		html.AClass(outerClass),
		html.ACustom("role", "progressbar"),
		html.AAria("valuemin", "0"),
		html.AAria("valuemax", strconv.Itoa(propsMax)),
		html.AAria("valuenow", strconv.Itoa(clamp(props.Value, 0, propsMax))),
	}
	for _, attr := range props.Attrs {
		outerArgs = append(outerArgs, attr)
	}

	children := make([]html.Component, 0, 2)
	if props.Label != "" || props.ShowValue {
		labelArgs := []html.DivArg{html.AClass("flex justify-between items-center mb-1")}
		if props.Label != "" {
			labelArgs = append(labelArgs, html.Span(html.AClass("text-sm font-medium"), html.Text(props.Label)))
		}
		if props.ShowValue {
			pct := percentage(props.Value, propsMax)
			labelArgs = append(labelArgs, html.Span(html.AClass("text-sm font-medium"), html.Text(strconv.Itoa(pct)+"%")))
		}
		children = append(children, html.Div(labelArgs...))
	}

	bar := html.Div(
		html.AData("pui-progress-indicator", ""),
		html.AClass(classnames.Merge(
			"h-full rounded-full transition-all",
			sizeClass(props.Size),
			variantClass(props.Variant),
			props.BarClass,
		)),
		html.AStyle("width: "+strconv.Itoa(percentage(props.Value, propsMax))+"%;"),
	)
	barWrapper := html.Div(
		html.AClass("w-full overflow-hidden rounded-full bg-secondary"),
		bar,
	)
	children = append(children, barWrapper)

	for _, child := range children {
		outerArgs = append(outerArgs, html.Child(child))
	}
	return html.Div(outerArgs...)
}

func sizeClass(size Size) string {
	switch size {
	case SizeSm:
		return "h-1"
	case SizeLg:
		return "h-4"
	default:
		return "h-2.5"
	}
}

func variantClass(variant Variant) string {
	switch variant {
	case VariantSuccess:
		return "bg-green-500"
	case VariantDanger:
		return "bg-destructive"
	case VariantWarning:
		return "bg-yellow-500"
	default:
		return "bg-primary"
	}
}

func maxValue(max int) int {
	if max <= 0 {
		return 100
	}
	return max
}

func percentage(value, max int) int {
	clamped := clamp(value, 0, max)
	if max == 0 {
		return 0
	}
	return (clamped * 100) / max
}

func clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func randomID() string {
	bytes := make([]byte, 6)
	if _, err := rand.Read(bytes); err != nil {
		return "progress-id"
	}
	return "progress-" + hex.EncodeToString(bytes)
}
