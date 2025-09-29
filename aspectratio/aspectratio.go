package aspectratio

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/internal/classnames"
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

func AspectRatio(props Props, args ...html.DivArg) html.Node {
	containerArgs := []html.DivArg{
		html.AClass(classnames.Merge("relative w-full", ratioClass(props.Ratio), props.Class)),
	}
	if props.ID != "" {
		containerArgs = append(containerArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		containerArgs = append(containerArgs, attr)
	}

	innerArgs := []html.DivArg{html.AClass("absolute inset-0")}
	innerArgs = append(innerArgs, args...)

	containerArgs = append(containerArgs, html.Div(innerArgs...))

	return html.Div(containerArgs...)
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
