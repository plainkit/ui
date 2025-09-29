package separator

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/internal/classnames"
)

type Orientation string
type Decoration string

type Props struct {
	ID          string
	Class       string
	Attrs       []html.Global
	Orientation Orientation
	Decoration  Decoration
}

const (
	OrientationHorizontal Orientation = "horizontal"
	OrientationVertical   Orientation = "vertical"
)

const (
	DecorationDashed Decoration = "dashed"
	DecorationDotted Decoration = "dotted"
)

// Separator renders a decorative divider. Label args populate the centered span.
func Separator(props Props, labelArgs ...html.SpanArg) html.Node {
	if props.Orientation == "" {
		props.Orientation = OrientationHorizontal
	}

	if props.Orientation == OrientationVertical {
		return verticalSeparator(props, labelArgs)
	}

	return horizontalSeparator(props, labelArgs)
}

func horizontalSeparator(props Props, labelArgs []html.SpanArg) html.Node {
	outerArgs := []html.DivArg{
		html.AClass(classnames.Merge("shrink-0 w-full", props.Class)),
		html.ACustom("role", "separator"),
		html.AAria("orientation", "horizontal"),
	}
	if props.ID != "" {
		outerArgs = append(outerArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		outerArgs = append(outerArgs, attr)
	}

	inner := html.Div(
		html.AClass("relative flex items-center w-full"),
		html.Span(
			html.AClass(classnames.Merge("absolute w-full border-t h-[1px]", decorationClass(props.Decoration))),
			html.ACustom("aria-hidden", "true"),
		),
		html.Span(append([]html.SpanArg{html.AClass("relative mx-auto bg-background px-2 text-xs text-muted-foreground")}, labelArgs...)...),
	)

	outerArgs = append(outerArgs, inner)

	return html.Div(outerArgs...)
}

func verticalSeparator(props Props, labelArgs []html.SpanArg) html.Node {
	outerArgs := []html.DivArg{
		html.AClass(classnames.Merge("shrink-0 h-full", props.Class)),
		html.ACustom("role", "separator"),
		html.AAria("orientation", "vertical"),
	}
	if props.ID != "" {
		outerArgs = append(outerArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		outerArgs = append(outerArgs, attr)
	}

	inner := html.Div(
		html.AClass("relative flex flex-col items-center h-full"),
		html.Span(
			html.AClass(classnames.Merge("absolute h-full border-l w-[1px]", decorationClass(props.Decoration))),
			html.ACustom("aria-hidden", "true"),
		),
		html.Span(append([]html.SpanArg{html.AClass("relative my-auto bg-background py-2 text-xs text-muted-foreground")}, labelArgs...)...),
	)

	outerArgs = append(outerArgs, inner)

	return html.Div(outerArgs...)
}

func decorationClass(dec Decoration) string {
	switch dec {
	case DecorationDashed:
		return "border-dashed"
	case DecorationDotted:
		return "border-dotted"
	default:
		return ""
	}
}
