package dropdown

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/button"
	"github.com/plainkit/ui/internal/classnames"
	"github.com/plainkit/ui/popover"
)

type Placement = popover.Placement

const (
	PlacementTop         = popover.PlacementTop
	PlacementTopStart    = popover.PlacementTopStart
	PlacementTopEnd      = popover.PlacementTopEnd
	PlacementRight       = popover.PlacementRight
	PlacementRightStart  = popover.PlacementRightStart
	PlacementRightEnd    = popover.PlacementRightEnd
	PlacementBottom      = popover.PlacementBottom
	PlacementBottomStart = popover.PlacementBottomStart
	PlacementBottomEnd   = popover.PlacementBottomEnd
	PlacementLeft        = popover.PlacementLeft
	PlacementLeftStart   = popover.PlacementLeftStart
	PlacementLeftEnd     = popover.PlacementLeftEnd
)

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
	ID        string
	Class     string
	Attrs     []html.Global
	Width     string
	MaxHeight string
	Placement Placement
}

type GroupProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type LabelProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type ItemProps struct {
	ID           string
	Class        string
	Attrs        []html.Global
	Disabled     bool
	Href         string
	Target       string
	PreventClose bool
}

type SeparatorProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type ShortcutProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type SubProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type SubTriggerProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type SubContentProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

// Dropdown renders a dropdown container
func Dropdown(props Props, args ...html.DivArg) html.Node {
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

// Trigger creates a dropdown trigger using a button
func Trigger(triggerProps TriggerProps, buttonProps button.Props, args ...html.ButtonArg) html.Node {
	contentID := triggerProps.For
	if contentID == "" {
		contentID = randomID("dropdown")
	}

	return popover.Trigger(
		popover.TriggerProps{
			ID:          triggerProps.ID,
			For:         contentID,
			TriggerType: popover.TriggerTypeClick,
			Class:       triggerProps.Class,
			Attrs:       triggerProps.Attrs,
		},
		button.Button(append([]html.ButtonArg{buttonProps}, args...)...),
	)
}

// Content creates the dropdown content panel
func Content(props ContentProps, args ...html.DivArg) html.Node {
	contentID := props.ID
	if contentID == "" {
		contentID = randomID("dropdown-content")
	}

	placement := props.Placement
	if placement == "" {
		placement = PlacementBottomStart
	}

	maxHeight := "300px"
	if props.MaxHeight != "" {
		maxHeight = props.MaxHeight
	}

	contentClass := classnames.Merge(
		"z-50 rounded-md bg-popover p-1 shadow-md focus:outline-none overflow-auto",
		"border border-border",
		"min-w-[8rem]",
		"max-h-["+maxHeight+"]",
		props.Width,
		props.Class,
	)

	return popover.Content(
		popover.ContentProps{
			ID:        contentID,
			Placement: placement,
			Offset:    4,
			Class:     contentClass,
			Attrs:     props.Attrs,
		},
		args...,
	)
}

// Group creates a dropdown group container
func Group(props GroupProps, args ...html.DivArg) html.Node {
	divArgs := []html.DivArg{
		html.AClass(classnames.Merge("py-1", props.Class)),
		html.AAria("role", "group"),
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

// Label creates a dropdown label
func Label(props LabelProps, args ...html.DivArg) html.Node {
	divArgs := []html.DivArg{
		html.AClass(classnames.Merge("px-2 py-1.5 text-sm font-semibold", props.Class)),
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

// Item creates a dropdown item (button or link)
func Item(props ItemProps, args ...html.Node) html.Node {
	id := props.ID
	if id == "" {
		id = randomID("dropdown-item")
	}

	baseClasses := classnames.Merge(
		"flex text-left items-center justify-between px-2 py-1.5 text-sm rounded-sm",
		"focus:bg-accent focus:text-accent-foreground hover:bg-accent hover:text-accent-foreground cursor-default",
		props.Class,
	)

	if props.Disabled {
		baseClasses = classnames.Merge(baseClasses, "opacity-50 pointer-events-none")
	}

	attrs := []html.Global{
		html.AAria("role", "menuitem"),
		html.AData("pui-dropdown-item", ""),
	}
	if props.PreventClose {
		attrs = append(attrs, html.AData("pui-dropdown-prevent-close", "true"))
	}

	attrs = append(attrs, props.Attrs...)

	if props.Href != "" {
		// Create link
		aArgs := []html.AArg{
			html.AId(id),
			html.AHref(props.Href),
			html.AClass(baseClasses),
		}
		if props.Target != "" {
			aArgs = append(aArgs, html.ATarget(props.Target))
		}

		for _, attr := range attrs {
			aArgs = append(aArgs, attr)
		}

		for _, arg := range args {
			aArgs = append(aArgs, arg)
		}

		return html.A(aArgs...)
	}

	// Create button
	buttonArgs := []html.ButtonArg{
		html.AId(id),
		html.AClass(baseClasses),
		html.AType("button"),
	}
	if props.Disabled {
		buttonArgs = append(buttonArgs, html.ADisabled())
	}

	for _, attr := range attrs {
		buttonArgs = append(buttonArgs, attr)
	}

	for _, arg := range args {
		buttonArgs = append(buttonArgs, arg)
	}

	return html.Button(buttonArgs...)
}

// Separator creates a dropdown separator
func Separator(props SeparatorProps, args ...html.DivArg) html.Node {
	divArgs := []html.DivArg{
		html.AClass(classnames.Merge("h-px my-1 -mx-1 bg-muted", props.Class)),
		html.AAria("role", "separator"),
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

// Shortcut creates a dropdown shortcut indicator
func Shortcut(props ShortcutProps, args ...html.SpanArg) html.Node {
	spanArgs := []html.SpanArg{
		html.AClass(classnames.Merge("ml-auto text-xs tracking-widest opacity-60", props.Class)),
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

// Sub creates a dropdown submenu container
func Sub(props SubProps, args ...html.DivArg) html.Node {
	divArgs := []html.DivArg{
		html.AClass(classnames.Merge("relative", props.Class)),
		html.AData("pui-dropdown-submenu", ""),
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

// SubTrigger creates a submenu trigger
func SubTrigger(props SubTriggerProps, subContentID string, args ...html.Node) html.Node {
	if subContentID == "" {
		subContentID = randomID("submenu")
	}

	triggerContent := html.Button(
		html.AType("button"),
		html.AClass(classnames.Merge(
			"w-full text-left flex items-center justify-between px-2 py-1.5 text-sm rounded-sm",
			"focus:bg-accent focus:text-accent-foreground hover:bg-accent hover:text-accent-foreground cursor-default",
			props.Class,
		)),
		html.AData("pui-dropdown-submenu-trigger", ""),
		html.Span(
			func() []html.SpanArg {
				spanArgs := make([]html.SpanArg, 0, len(args))
				for _, arg := range args {
					spanArgs = append(spanArgs, arg)
				}

				return spanArgs
			}()...,
		),
		lucide.ChevronRight(html.AClass("h-4 w-4 ml-auto")),
	)

	return popover.Trigger(
		popover.TriggerProps{
			ID:          props.ID,
			For:         subContentID,
			TriggerType: popover.TriggerTypeHover,
			Attrs:       props.Attrs,
		},
		triggerContent,
	)
}

// SubContent creates submenu content
func SubContent(props SubContentProps, args ...html.DivArg) html.Node {
	subContentID := props.ID
	if subContentID == "" {
		subContentID = randomID("submenu-content")
	}

	return popover.Content(
		popover.ContentProps{
			ID:            subContentID,
			Placement:     PlacementRightStart,
			Offset:        -4,
			HoverDelay:    100,
			HoverOutDelay: 200,
			Class: classnames.Merge(
				"z-[9999] min-w-[8rem] rounded-md border bg-popover p-1 shadow-lg",
				props.Class,
			),
			Attrs: props.Attrs,
		},
		args...,
	)
}

func randomID(prefix string) string {
	buf := make([]byte, 4)
	if _, err := rand.Read(buf); err != nil {
		return prefix + "-id"
	}

	return prefix + "-" + hex.EncodeToString(buf)
}
