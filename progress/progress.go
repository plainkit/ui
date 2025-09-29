package progress

import (
	"crypto/rand"
	"encoding/hex"
	"strconv"

	"github.com/plainkit/html"
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

func divArgsFromProps(baseClass string, extra ...string) func(p Props) []html.DivArg {
	return func(p Props) []html.DivArg {
		id := p.ID
		if id == "" {
			id = randomID()
		}

		className := html.ClassMerge(
			append([]string{baseClass},
				append(extra, p.Class)...)...)

		args := []html.DivArg{
			html.AId(id),
			html.AClass(className),
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

// ApplyDiv implements the html.DivArg interface for Props
func (p Props) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	propsMax := maxValue(p.Max)

	args := divArgsFromProps("w-full")(p)
	args = append([]html.DivArg{
		html.ACustom("role", "progressbar"),
		html.AAria("valuemin", "0"),
		html.AAria("valuemax", strconv.Itoa(propsMax)),
		html.AAria("valuenow", strconv.Itoa(clamp(p.Value, 0, propsMax))),
	}, args...)

	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

// Progress renders a progress bar using the composable pattern.
// Accepts variadic html.DivArg arguments, with Props as an optional first argument.
func Progress(args ...html.DivArg) html.Node {
	var (
		props Props
		rest  []html.DivArg
	)

	// Separate Props from other arguments
	for _, a := range args {
		if v, ok := a.(Props); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	propsMax := maxValue(props.Max)
	if props.ID == "" {
		props.ID = randomID()
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
		html.AClass(html.ClassMerge(
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

	divArgs := append([]html.DivArg{props}, rest...)
	for _, child := range children {
		divArgs = append(divArgs, html.Child(child))
	}

	return html.Div(divArgs...)
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
