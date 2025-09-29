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

func triggerSpanArgsFromProps(baseClass string, extra ...string) func(p TriggerProps) []html.SpanArg {
	return func(p TriggerProps) []html.SpanArg {
		triggerType := p.TriggerType
		if triggerType == "" {
			triggerType = TriggerTypeClick
		}

		args := []html.SpanArg{
			html.AClass(classnames.Merge(append([]string{baseClass}, append(extra, p.Class)...)...)),
			html.AData("pui-popover-open", "false"),
			html.AData("pui-popover-type", string(triggerType)),
		}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		if p.For != "" {
			args = append(args, html.AData("pui-popover-trigger", p.For))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p TriggerProps) ApplySpan(attrs *html.SpanAttrs, children *[]html.Component) {
	for _, a := range triggerSpanArgsFromProps("group cursor-pointer")(p) {
		a.ApplySpan(attrs, children)
	}
}

// Trigger renders the interactive element that toggles a popover.
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

	node := html.Span(append([]html.SpanArg{props}, rest...)...)

	return node.WithAssets("", popoverJS, "ui-popover")
}

func contentDivArgsFromProps(baseClass string, extra ...string) func(p ContentProps) []html.DivArg {
	return func(p ContentProps) []html.DivArg {
		placement := p.Placement
		if placement == "" {
			placement = PlacementBottom
		}

		offset := p.Offset
		if offset == 0 {
			if p.ShowArrow {
				offset = 8
			} else {
				offset = 4
			}
		}

		args := []html.DivArg{
			html.AClass(classnames.Merge(append([]string{baseClass}, append(extra, p.Class)...)...)),
			html.AData("pui-popover-id", p.ID),
			html.AData("pui-popover-open", "false"),
			html.AData("pui-popover-placement", string(placement)),
			html.AData("pui-popover-offset", strconv.Itoa(offset)),
			html.AData("pui-popover-disable-clickaway", strconv.FormatBool(p.DisableClickAway)),
			html.AData("pui-popover-disable-esc", strconv.FormatBool(p.DisableESC)),
			html.AData("pui-popover-show-arrow", strconv.FormatBool(p.ShowArrow)),
			html.AData("pui-popover-hover-delay", strconv.Itoa(p.HoverDelay)),
			html.AData("pui-popover-hover-out-delay", strconv.Itoa(p.HoverOutDelay)),
		}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		if p.MatchWidth {
			args = append(args, html.AData("pui-popover-match-width", "true"))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p ContentProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	for _, a := range contentDivArgsFromProps(
		"bg-popover rounded-lg border text-popover-foreground text-sm shadow-lg pointer-events-auto absolute z-[9999] hidden top-0 left-0",
	)(p) {
		a.ApplyDiv(attrs, children)
	}
}

// Content defines the floating panel positioning and styling.
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

	innerArgs := append([]html.DivArg{html.AClass("w-full overflow-hidden")}, rest...)

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

	node := html.Div(append([]html.DivArg{props}, contentInner...)...)

	return node.WithAssets("", popoverJS, "ui-popover")
}

//go:embed popover.js
var popoverJS string
