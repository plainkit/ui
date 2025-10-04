package dropdown

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/button"
	"github.com/plainkit/ui/internal/styles"
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

// Dropdown renders a dropdown container
func Dropdown(args ...html.DivArg) html.Node {
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

func triggerSpanArgsFromProps(baseClass string, extra ...string) func(p TriggerProps) []html.SpanArg {
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
	for _, a := range triggerSpanArgsFromProps("")(p) {
		a.ApplySpan(attrs, children)
	}
}

// Trigger creates a dropdown trigger using a button
func Trigger(args ...interface{}) html.Node {
	var (
		triggerProps TriggerProps
		buttonProps  button.Props
		rest         []html.ButtonArg
	)

	// Parse arguments
	for _, a := range args {
		if v, ok := a.(TriggerProps); ok {
			triggerProps = v
		} else if v, ok := a.(button.Props); ok {
			buttonProps = v
		} else if buttonArg, ok := a.(html.ButtonArg); ok {
			rest = append(rest, buttonArg)
		}
	}

	contentID := triggerProps.For
	if contentID == "" {
		contentID = randomID("dropdown")
	}

	if buttonProps.Variant == "" {
		buttonProps.Variant = button.VariantOutline
	}

	buttonProps.Class = html.ClassMerge(
		styles.Interactive(
			"dropdown-trigger inline-flex w-full items-center justify-between gap-2 px-4 py-2 text-sm",
			"text-left font-medium",
		),
		"select-none",
		buttonProps.Class,
	)

	return popover.Trigger(
		popover.TriggerProps{
			ID:          triggerProps.ID,
			For:         contentID,
			TriggerType: popover.TriggerTypeClick,
			Class:       triggerProps.Class,
			Attrs:       triggerProps.Attrs,
		},
		button.Button(append([]html.ButtonArg{buttonProps}, rest...)...),
	)
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
	for _, a := range contentDivArgsFromProps("")(p) {
		a.ApplyDiv(attrs, children)
	}
}

// Content creates the dropdown content panel
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

	contentClass := html.ClassMerge(
		styles.Panel("dropdown-content z-50 max-h-["+maxHeight+"] overflow-auto p-2 shadow-xl"),
		"min-w-[8rem]",
		props.Width,
		props.Class,
	)

	contentProps := popover.ContentProps{
		ID:        contentID,
		Placement: placement,
		Offset:    4,
		Class:     contentClass,
		Attrs:     props.Attrs,
	}

	return popover.Content(append([]html.DivArg{contentProps}, rest...)...)
}

func groupDivArgsFromProps(baseClass string, extra ...string) func(p GroupProps) []html.DivArg {
	return func(p GroupProps) []html.DivArg {
		args := []html.DivArg{
			html.AClass(html.ClassMerge(append([]string{baseClass}, append(extra, p.Class)...)...)),
			html.AAria("role", "group"),
		}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p GroupProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	for _, a := range groupDivArgsFromProps(styles.SurfaceMuted("space-y-1 rounded-xl bg-transparent py-1.5"))(p) {
		a.ApplyDiv(attrs, children)
	}
}

// Group creates a dropdown group container
func Group(args ...html.DivArg) html.Node {
	var (
		props GroupProps
		rest  []html.DivArg
	)

	for _, a := range args {
		if v, ok := a.(GroupProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Div(append([]html.DivArg{props}, rest...)...)
}

func labelDivArgsFromProps(baseClass string, extra ...string) func(p LabelProps) []html.DivArg {
	return func(p LabelProps) []html.DivArg {
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

func (p LabelProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	for _, a := range labelDivArgsFromProps(styles.SubHeading("px-3 py-2 text-xs uppercase tracking-wide text-muted-foreground/60"))(p) {
		a.ApplyDiv(attrs, children)
	}
}

// Label creates a dropdown label
func Label(args ...html.DivArg) html.Node {
	var (
		props LabelProps
		rest  []html.DivArg
	)

	for _, a := range args {
		if v, ok := a.(LabelProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Div(append([]html.DivArg{props}, rest...)...)
}

// Item creates a dropdown item (button or link)
func Item(props ItemProps, args ...html.Node) html.Node {
	id := props.ID
	if id == "" {
		id = randomID("dropdown-item")
	}

	itemClass := html.ClassMerge(
		styles.InteractiveGhost(
			"dropdown-item w-full items-center justify-between gap-3 px-3 py-2 text-sm",
			"text-left",
		),
		props.Class,
	)

	if props.Disabled {
		itemClass = html.ClassMerge(itemClass, "pointer-events-none opacity-50")
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
			html.AClass(itemClass),
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
		html.AClass(itemClass),
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

func separatorDivArgsFromProps(baseClass string, extra ...string) func(p SeparatorProps) []html.DivArg {
	return func(p SeparatorProps) []html.DivArg {
		args := []html.DivArg{
			html.AClass(html.ClassMerge(append([]string{baseClass}, append(extra, p.Class)...)...)),
			html.AAria("role", "separator"),
		}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p SeparatorProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	for _, a := range separatorDivArgsFromProps("my-2 h-px -mx-2 bg-gradient-to-r from-transparent via-border/60 to-transparent")(p) {
		a.ApplyDiv(attrs, children)
	}
}

// Separator creates a dropdown separator
func Separator(args ...html.DivArg) html.Node {
	var (
		props SeparatorProps
		rest  []html.DivArg
	)

	for _, a := range args {
		if v, ok := a.(SeparatorProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Div(append([]html.DivArg{props}, rest...)...)
}

func shortcutSpanArgsFromProps(baseClass string, extra ...string) func(p ShortcutProps) []html.SpanArg {
	return func(p ShortcutProps) []html.SpanArg {
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

func (p ShortcutProps) ApplySpan(attrs *html.SpanAttrs, children *[]html.Component) {
	for _, a := range shortcutSpanArgsFromProps(styles.SubtleText("ml-auto text-[11px] tracking-[0.25em] uppercase"))(p) {
		a.ApplySpan(attrs, children)
	}
}

// Shortcut creates a dropdown shortcut indicator
func Shortcut(args ...html.SpanArg) html.Node {
	var (
		props ShortcutProps
		rest  []html.SpanArg
	)

	for _, a := range args {
		if v, ok := a.(ShortcutProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Span(append([]html.SpanArg{props}, rest...)...)
}

func subDivArgsFromProps(baseClass string, extra ...string) func(p SubProps) []html.DivArg {
	return func(p SubProps) []html.DivArg {
		args := []html.DivArg{
			html.AClass(html.ClassMerge(append([]string{baseClass}, append(extra, p.Class)...)...)),
			html.AData("pui-dropdown-submenu", ""),
		}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p SubProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	for _, a := range subDivArgsFromProps("relative")(p) {
		a.ApplyDiv(attrs, children)
	}
}

// Sub creates a dropdown submenu container
func Sub(args ...html.DivArg) html.Node {
	var (
		props SubProps
		rest  []html.DivArg
	)

	for _, a := range args {
		if v, ok := a.(SubProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Div(append([]html.DivArg{props}, rest...)...)
}

// SubTrigger creates a submenu trigger
func SubTrigger(props SubTriggerProps, subContentID string, args ...html.Node) html.Node {
	if subContentID == "" {
		subContentID = randomID("submenu")
	}

	triggerContent := html.Button(
		html.AType("button"),
		html.AClass(html.ClassMerge(
			styles.InteractiveGhost(
				"dropdown-subtrigger flex w-full items-center justify-between gap-2 px-3 py-2 text-sm text-left",
			),
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
		lucide.ChevronRight(html.AClass("ml-auto size-4 text-muted-foreground")),
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

	contentProps := popover.ContentProps{
		ID:            subContentID,
		Placement:     PlacementRightStart,
		Offset:        -4,
		HoverDelay:    100,
		HoverOutDelay: 200,
		Class: html.ClassMerge(
			styles.Panel("z-[9999] min-w-[8rem] p-2 shadow-xl"),
			props.Class,
		),
		Attrs: props.Attrs,
	}

	return popover.Content(append([]html.DivArg{contentProps}, args...)...)
}

func randomID(prefix string) string {
	buf := make([]byte, 4)
	if _, err := rand.Read(buf); err != nil {
		return prefix + "-id"
	}

	return prefix + "-" + hex.EncodeToString(buf)
}
