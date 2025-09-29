package aspectratio

import (
	"github.com/plainkit/html"
)

type Ratio string

const (
	RatioAuto     Ratio = "auto"
	RatioSquare   Ratio = "square"
	RatioVideo    Ratio = "video"
	RatioPortrait Ratio = "portrait"
	RatioWide     Ratio = "wide"
)

type Props struct {
	ID    string
	Class string
	Attrs []html.Global
	Ratio Ratio
}

func divArgsFromProps(baseClass string, extra ...string) func(p Props) []html.DivArg {
	return func(p Props) []html.DivArg {
		classNames := append([]string{baseClass}, extra...)
		classNames = append(classNames, ratioClass(p.Ratio), p.Class)

		args := []html.DivArg{html.AClass(html.ClassMerge(classNames...))}
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
	for _, a := range divArgsFromProps("relative w-full")(p) {
		a.ApplyDiv(attrs, children)
	}
}

func AspectRatio(args ...html.DivArg) html.Node {
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

	innerArgs := []html.DivArg{html.AClass("absolute inset-0")}
	innerArgs = append(innerArgs, rest...)

	return html.Div(append([]html.DivArg{props}, html.Div(innerArgs...))...)
}

func ratioClass(r Ratio) string {
	switch r {
	case RatioSquare:
		return "aspect-square"
	case RatioVideo:
		return "aspect-video"
	case RatioPortrait:
		return "aspect-[3/4]"
	case RatioWide:
		return "aspect-[2/1]"
	case RatioAuto:
		fallthrough
	default:
		return "aspect-auto"
	}
}
