package tooltip

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/internal/classnames"
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

// Tooltip renders a tooltip container (wrapper component)
func Tooltip(props Props, args ...html.DivArg) html.Node {
	divArgs := []html.DivArg{
		html.AClass(classnames.Merge("", props.Class)),
	}

	if props.ID != "" {
		divArgs = append(divArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}
	divArgs = append(divArgs, args...)

	return html.Div(divArgs...)
}

// Trigger creates a tooltip trigger element
func Trigger(props TriggerProps, args ...html.SpanArg) html.Node {
	return popover.Trigger(
		popover.TriggerProps{
			ID:          props.ID,
			TriggerType: popover.TriggerTypeHover,
			For:         props.For,
			Class:       props.Class,
			Attrs:       props.Attrs,
		},
		args...,
	)
}

// Content creates the tooltip content panel
func Content(props ContentProps, args ...html.DivArg) html.Node {
	tooltipClass := classnames.Merge(
		"px-3 py-1.5 text-xs font-medium text-primary-foreground bg-primary border-primary rounded-md",
		"shadow-md max-w-xs",
		props.Class,
	)

	return popover.Content(
		popover.ContentProps{
			ID:            props.ID,
			Class:         tooltipClass,
			Attrs:         props.Attrs,
			Placement:     mapTooltipPositionToPopover(props.Position),
			ShowArrow:     props.ShowArrow,
			HoverDelay:    props.HoverDelay,
			HoverOutDelay: props.HoverOutDelay,
		},
		args...,
	)
}
