package dialog

import (
	"crypto/rand"
	_ "embed"
	"encoding/hex"

	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/button"
	"github.com/plainkit/ui/internal/styles"
)

type Props struct {
	ID               string
	Class            string
	Attrs            []html.Global
	DisableClickAway bool
	DisableESC       bool
	Open             bool
}

type TriggerProps struct {
	ID       string
	Class    string
	Attrs    []html.Global
	Disabled bool
	For      string // Reference to a specific dialog ID (for external triggers)
}

type ContentProps struct {
	ID              string
	Class           string
	Attrs           []html.Global
	HideCloseButton bool
	Open            bool // Initial open state for standalone usage
}

type CloseProps struct {
	ID    string
	Class string
	Attrs []html.Global
	For   string // ID of the dialog to close (optional, defaults to closest dialog)
}

type HeaderProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type FooterProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type TitleProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type DescriptionProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

func dialogDivArgsFromProps(baseClass string, extra ...string) func(p Props) []html.DivArg {
	return func(p Props) []html.DivArg {
		instanceID := p.ID
		if instanceID == "" {
			instanceID = randomID("dialog")
		}

		args := []html.DivArg{
			html.AData("pui-dialog", ""),
			html.AData("dialog-instance", instanceID),
			html.AClass(html.ClassMerge(append([]string{baseClass}, append(extra, p.Class)...)...)),
		}

		if p.DisableClickAway {
			args = append(args, html.AData("pui-dialog-disable-click-away", "true"))
		}

		if p.DisableESC {
			args = append(args, html.AData("pui-dialog-disable-esc", "true"))
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

func (p Props) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	for _, a := range dialogDivArgsFromProps("relative z-50")(p) {
		a.ApplyDiv(attrs, children)
	}
}

// Dialog renders a dialog container component
func Dialog(args ...html.DivArg) html.Node {
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

	return html.Div(append([]html.DivArg{props}, rest...)...).WithAssets("", dialogJS, "ui-dialog")
}

// Trigger creates a dialog trigger element
func Trigger(triggerProps TriggerProps, buttonProps button.Props, args ...html.ButtonArg) html.Node {
	instanceID := triggerProps.For
	if instanceID == "" && triggerProps.ID != "" {
		instanceID = triggerProps.ID
	}

	// Merge button attributes with trigger-specific data attributes
	attrs := buttonProps.Attrs
	if attrs == nil {
		attrs = []html.Global{}
	}

	attrs = append(attrs,
		html.AData("pui-dialog-trigger", instanceID),
		html.AData("dialog-instance", instanceID),
		html.AData("pui-dialog-trigger-open", "false"),
	)
	attrs = append(attrs, triggerProps.Attrs...)

	if triggerProps.Disabled {
		buttonProps.Disabled = true
	}

	if triggerProps.Class != "" {
		buttonProps.Class = html.ClassMerge(buttonProps.Class, triggerProps.Class)
	}

	buttonProps.Attrs = attrs
	if triggerProps.ID != "" {
		buttonProps.ID = triggerProps.ID
	}

	return button.Button(append([]html.ButtonArg{buttonProps}, args...)...).WithAssets("", dialogJS, "ui-dialog")
}

// Content creates the dialog content panel
func Content(props ContentProps, args ...html.DivArg) html.Node {
	instanceID := props.ID
	if instanceID == "" {
		instanceID = randomID("dialog-content")
	}

	// Overlay/backdrop
	overlayClasses := html.ClassMerge(
		"fixed inset-0 z-40 bg-black/50 backdrop-blur-sm",
		"transition-opacity duration-300",
		"data-[pui-dialog-open=false]:opacity-0",
		"data-[pui-dialog-open=true]:opacity-100",
		"data-[pui-dialog-open=false]:pointer-events-none",
		"data-[pui-dialog-open=true]:pointer-events-auto",
		"data-[pui-dialog-hidden=true]:!hidden",
	)

	overlayArgs := []html.DivArg{
		html.AClass(overlayClasses),
		html.AData("pui-dialog-backdrop", ""),
		html.AData("dialog-instance", instanceID),
	}

	if props.Open {
		overlayArgs = append(overlayArgs, html.AData("pui-dialog-open", "true"))
	} else {
		overlayArgs = append(overlayArgs,
			html.AData("pui-dialog-open", "false"),
			html.AData("pui-dialog-hidden", "true"),
		)
	}

	// Content panel
	contentClasses := html.ClassMerge(
		"fixed left-1/2 top-1/2 z-50 w-full max-w-[min(90vw,620px)] -translate-x-1/2 -translate-y-1/2",
		styles.Panel("relative grid gap-6 p-8"),
		"transition-all duration-200",
		"data-[pui-dialog-open=false]:scale-95 data-[pui-dialog-open=false]:opacity-0",
		"data-[pui-dialog-open=true]:scale-100 data-[pui-dialog-open=true]:opacity-100",
		"data-[pui-dialog-open=false]:pointer-events-none data-[pui-dialog-open=true]:pointer-events-auto",
		"data-[pui-dialog-hidden=true]:!hidden",
		props.Class,
	)

	contentArgs := []html.DivArg{
		html.AClass(contentClasses),
		html.AData("pui-dialog-content", ""),
		html.AData("dialog-instance", instanceID),
	}

	if props.Open {
		contentArgs = append(contentArgs, html.AData("pui-dialog-open", "true"))
	} else {
		contentArgs = append(contentArgs,
			html.AData("pui-dialog-open", "false"),
			html.AData("pui-dialog-hidden", "true"),
		)
	}

	for _, attr := range props.Attrs {
		contentArgs = append(contentArgs, attr)
	}

	contentArgs = append(contentArgs, args...)

	// Add close button if not hidden
	if !props.HideCloseButton {
		closeButton := html.Button(
			html.AClass(html.ClassMerge(
				"absolute right-4 top-4",
				styles.InteractiveGhost(
					"size-8 rounded-full",
					"bg-muted/40 text-muted-foreground/80",
					"hover:text-foreground",
				),
				"transition-opacity hover:opacity-100",
				"data-[pui-dialog-open=false]:opacity-0",
				"data-[pui-dialog-open=true]:opacity-80",
				"[&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4",
			)),
			html.AData("pui-dialog-close", instanceID),
			html.AAria("label", "Close"),
			html.AType("button"),
			lucide.X(),
			html.Span(html.AClass("sr-only"), html.Text("Close")),
		)
		contentArgs = append(contentArgs, closeButton)
	}

	return html.Div(
		html.Div(overlayArgs...),
		html.Div(contentArgs...),
	).WithAssets("", dialogJS, "ui-dialog")
}

func closeSpanArgsFromProps(baseClass string, extra ...string) func(p CloseProps) []html.SpanArg {
	return func(p CloseProps) []html.SpanArg {
		args := []html.SpanArg{
			html.AClass(html.ClassMerge(append([]string{baseClass}, append(extra, p.Class)...)...)),
		}

		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		if p.For != "" {
			args = append(args, html.AData("pui-dialog-close", p.For))
		} else {
			args = append(args, html.AData("pui-dialog-close", ""))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p CloseProps) ApplySpan(attrs *html.SpanAttrs, children *[]html.Component) {
	for _, a := range closeSpanArgsFromProps("contents cursor-pointer")(p) {
		a.ApplySpan(attrs, children)
	}
}

// Close creates a dialog close trigger
func Close(args ...html.SpanArg) html.Node {
	var (
		props CloseProps
		rest  []html.SpanArg
	)

	for _, a := range args {
		if v, ok := a.(CloseProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Span(append([]html.SpanArg{props}, rest...)...)
}

func headerDivArgsFromProps(baseClass string, extra ...string) func(p HeaderProps) []html.DivArg {
	return func(p HeaderProps) []html.DivArg {
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

func (p HeaderProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	for _, a := range headerDivArgsFromProps("flex flex-col gap-3 text-center sm:text-left")(p) {
		a.ApplyDiv(attrs, children)
	}
}

// Header creates a dialog header container
func Header(args ...html.DivArg) html.Node {
	var (
		props HeaderProps
		rest  []html.DivArg
	)

	for _, a := range args {
		if v, ok := a.(HeaderProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Div(append([]html.DivArg{props}, rest...)...)
}

func footerDivArgsFromProps(baseClass string, extra ...string) func(p FooterProps) []html.DivArg {
	return func(p FooterProps) []html.DivArg {
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

func (p FooterProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	for _, a := range footerDivArgsFromProps("flex flex-col-reverse gap-3 sm:flex-row sm:items-center sm:justify-end")(p) {
		a.ApplyDiv(attrs, children)
	}
}

// Footer creates a dialog footer container
func Footer(args ...html.DivArg) html.Node {
	var (
		props FooterProps
		rest  []html.DivArg
	)

	for _, a := range args {
		if v, ok := a.(FooterProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Div(append([]html.DivArg{props}, rest...)...)
}

func titleH2ArgsFromProps(baseClass string, extra ...string) func(p TitleProps) []html.H2Arg {
	return func(p TitleProps) []html.H2Arg {
		args := []html.H2Arg{html.AClass(html.ClassMerge(append([]string{baseClass}, append(extra, p.Class)...)...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p TitleProps) ApplyH2(attrs *html.H2Attrs, children *[]html.Component) {
	for _, a := range titleH2ArgsFromProps(styles.DisplayHeading("text-pretty text-balance"))(p) {
		a.ApplyH2(attrs, children)
	}
}

// Title creates a dialog title
func Title(args ...html.H2Arg) html.Node {
	var (
		props TitleProps
		rest  []html.H2Arg
	)

	for _, a := range args {
		if v, ok := a.(TitleProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.H2(append([]html.H2Arg{props}, rest...)...)
}

func descriptionPArgsFromProps(baseClass string, extra ...string) func(p DescriptionProps) []html.PArg {
	return func(p DescriptionProps) []html.PArg {
		args := []html.PArg{html.AClass(html.ClassMerge(append([]string{baseClass}, append(extra, p.Class)...)...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p DescriptionProps) ApplyP(attrs *html.PAttrs, children *[]html.Component) {
	for _, a := range descriptionPArgsFromProps(styles.SubtleText("leading-relaxed"))(p) {
		a.ApplyP(attrs, children)
	}
}

// Description creates a dialog description
func Description(args ...html.PArg) html.Node {
	var (
		props DescriptionProps
		rest  []html.PArg
	)

	for _, a := range args {
		if v, ok := a.(DescriptionProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.P(append([]html.PArg{props}, rest...)...)
}

func randomID(prefix string) string {
	buf := make([]byte, 4)
	if _, err := rand.Read(buf); err != nil {
		return prefix + "-id"
	}

	return prefix + "-" + hex.EncodeToString(buf)
}

//go:embed dialog.js
var dialogJS string
