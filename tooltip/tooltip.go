package tooltip

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/popover"
)

type Position string

const (
	PositionTop    Position = "top"
	PositionRight  Position = "right"
	PositionBottom Position = "bottom"
	PositionLeft   Position = "left"
)

// Map tooltip positions to popover positions
func mapTooltipPositionToPopover(position Position) popover.Placement {
	switch position {
	case PositionTop:
		return popover.PlacementTop
	case PositionRight:
		return popover.PlacementRight
	case PositionBottom:
		return popover.PlacementBottom
	case PositionLeft:
		return popover.PlacementLeft
	default:
		return popover.PlacementTop
	}
}

type Props struct {
	ID    string
	Class string
	Attrs []html.Global
}

type TriggerProps struct {
	ID    string
	Class string
	Attrs []html.Global
	For   string
}

type ContentProps struct {
	ID            string
	Class         string
	Attrs         []html.Global
	ShowArrow     bool
	Position      Position
	HoverDelay    int
	HoverOutDelay int
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
	for _, a := range divArgsFromProps("")(p) {
		a.ApplyDiv(attrs, children)
	}
}

// Tooltip renders a tooltip container (wrapper component)
func Tooltip(args ...html.DivArg) html.Node {
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

	return html.Div(append([]html.DivArg{props}, rest...)...)
}

func spanArgsFromProps(baseClass string, extra ...string) func(p TriggerProps) []html.SpanArg {
	return func(p TriggerProps) []html.SpanArg {
		args := []html.SpanArg{html.AClass(html.ClassMerge(append([]string{baseClass}, append(extra, p.Class)...)...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p TriggerProps) ApplySpan(attrs *html.SpanAttrs, children *[]html.Component) {
	for _, a := range spanArgsFromProps("")(p) {
		a.ApplySpan(attrs, children)
	}
}

// Trigger creates a tooltip trigger element
func Trigger(args ...html.SpanArg) html.Node {
	var (
		props TriggerProps
		rest  []html.SpanArg
	)

	for _, a := range args {
		if v, ok := a.(TriggerProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	triggerProps := popover.TriggerProps{
		ID:          props.ID,
		TriggerType: popover.TriggerTypeHover,
		For:         props.For,
		Class:       props.Class,
		Attrs:       props.Attrs,
	}

	return popover.Trigger(append([]html.SpanArg{triggerProps}, rest...)...)
}

func contentDivArgsFromProps(baseClass string, extra ...string) func(p ContentProps) []html.DivArg {
	return func(p ContentProps) []html.DivArg {
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

func (p ContentProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	for _, a := range contentDivArgsFromProps(
		"px-3 py-1.5 text-xs font-medium text-primary-foreground bg-primary border-primary rounded-md",
		"shadow-md max-w-xs",
	)(p) {
		a.ApplyDiv(attrs, children)
	}
}

// Content creates the tooltip content panel
func Content(args ...html.DivArg) html.Node {
	var (
		props ContentProps
		rest  []html.DivArg
	)

	for _, a := range args {
		if v, ok := a.(ContentProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	tooltipClass := html.ClassMerge(
		"px-3 py-1.5 text-xs font-medium text-primary-foreground bg-primary border-primary rounded-md",
		"shadow-md max-w-xs",
		props.Class,
	)

	contentProps := popover.ContentProps{
		ID:            props.ID,
		Class:         tooltipClass,
		Attrs:         props.Attrs,
		Placement:     mapTooltipPositionToPopover(props.Position),
		ShowArrow:     props.ShowArrow,
		HoverDelay:    props.HoverDelay,
		HoverOutDelay: props.HoverOutDelay,
	}

	return popover.Content(append([]html.DivArg{contentProps}, rest...)...)
}
