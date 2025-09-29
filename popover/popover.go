package popover

import (
	_ "embed"
	"strconv"

	"github.com/plainkit/html"
	"github.com/plainkit/ui/internal/classnames"
)

type Placement string

type TriggerType string

type TriggerProps struct {
	ID          string
	Class       string
	Attrs       []html.Global
	For         string
	TriggerType TriggerType
}

type ContentProps struct {
	ID               string
	Class            string
	Attrs            []html.Global
	Placement        Placement
	Offset           int
	DisableClickAway bool
	DisableESC       bool
	ShowArrow        bool
	HoverDelay       int
	HoverOutDelay    int
	MatchWidth       bool
}

const (
	PlacementTop         Placement = "top"
	PlacementTopStart    Placement = "top-start"
	PlacementTopEnd      Placement = "top-end"
	PlacementRight       Placement = "right"
	PlacementRightStart  Placement = "right-start"
	PlacementRightEnd    Placement = "right-end"
	PlacementBottom      Placement = "bottom"
	PlacementBottomStart Placement = "bottom-start"
	PlacementBottomEnd   Placement = "bottom-end"
	PlacementLeft        Placement = "left"
	PlacementLeftStart   Placement = "left-start"
	PlacementLeftEnd     Placement = "left-end"
)

const (
	TriggerTypeHover TriggerType = "hover"
	TriggerTypeClick TriggerType = "click"
)

// Trigger renders the interactive element that toggles a popover.
func Trigger(props TriggerProps, args ...html.SpanArg) html.Node {
	triggerType := props.TriggerType
	if triggerType == "" {
		triggerType = TriggerTypeClick
	}

	spanArgs := []html.SpanArg{
		html.AClass(classnames.Merge("group cursor-pointer", props.Class)),
		html.AData("pui-popover-open", "false"),
		html.AData("pui-popover-type", string(triggerType)),
	}
	if props.ID != "" {
		spanArgs = append(spanArgs, html.AId(props.ID))
	}

	if props.For != "" {
		spanArgs = append(spanArgs, html.AData("pui-popover-trigger", props.For))
	}

	for _, attr := range props.Attrs {
		spanArgs = append(spanArgs, attr)
	}

	spanArgs = append(spanArgs, args...)

	node := html.Span(spanArgs...)

	return node.WithAssets("", popoverJS, "ui-popover")
}

// Content defines the floating panel positioning and styling.
func Content(props ContentProps, args ...html.DivArg) html.Node {
	placement := props.Placement
	if placement == "" {
		placement = PlacementBottom
	}

	offset := props.Offset
	if offset == 0 {
		if props.ShowArrow {
			offset = 8
		} else {
			offset = 4
		}
	}

	divArgs := []html.DivArg{
		html.AClass(classnames.Merge(
			"bg-popover rounded-lg border text-popover-foreground text-sm shadow-lg pointer-events-auto absolute z-[9999] hidden top-0 left-0",
			props.Class,
		)),
		html.AData("pui-popover-id", props.ID),
		html.AData("pui-popover-open", "false"),
		html.AData("pui-popover-placement", string(placement)),
		html.AData("pui-popover-offset", strconv.Itoa(offset)),
		html.AData("pui-popover-disable-clickaway", strconv.FormatBool(props.DisableClickAway)),
		html.AData("pui-popover-disable-esc", strconv.FormatBool(props.DisableESC)),
		html.AData("pui-popover-show-arrow", strconv.FormatBool(props.ShowArrow)),
		html.AData("pui-popover-hover-delay", strconv.Itoa(props.HoverDelay)),
		html.AData("pui-popover-hover-out-delay", strconv.Itoa(props.HoverOutDelay)),
	}
	if props.ID != "" {
		divArgs = append(divArgs, html.AId(props.ID))
	}

	if props.MatchWidth {
		divArgs = append(divArgs, html.AData("pui-popover-match-width", "true"))
	}

	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}

	innerArgs := append([]html.DivArg{html.AClass("w-full overflow-hidden")}, args...)

	contentInner := []html.DivArg{html.Div(innerArgs...)}
	if props.ShowArrow {
		contentInner = append(contentInner, html.Div(
			html.AData("pui-popover-arrow", ""),
			html.AClass("absolute h-2.5 w-2.5 rotate-45 bg-popover border border-border "+
				"data-[pui-popover-placement^=top]:-bottom-[5px] data-[pui-popover-placement^=top]:border-t-transparent data-[pui-popover-placement^=top]:border-l-transparent "+
				"data-[pui-popover-placement^=bottom]:-top-[5px] data-[pui-popover-placement^=bottom]:border-b-transparent data-[pui-popover-placement^=bottom]:border-r-transparent "+
				"data-[pui-popover-placement^=left]:-right-[5px] data-[pui-popover-placement^=left]:border-l-transparent data-[pui-popover-placement^=left]:border-b-transparent "+
				"data-[pui-popover-placement^=right]:-left-[5px] data-[pui-popover-placement^=right]:border-r-transparent data-[pui-popover-placement^=right]:border-t-transparent"),
		))
	}

	node := html.Div(append(divArgs, contentInner...)...)

	return node.WithAssets("", popoverJS, "ui-popover")
}

//go:embed popover.js
var popoverJS string
