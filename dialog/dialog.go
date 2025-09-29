package dialog

import (
	"crypto/rand"
	_ "embed"
	"encoding/hex"

	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/button"
	"github.com/plainkit/ui/internal/classnames"
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

// Dialog renders a dialog container component
func Dialog(props Props, args ...html.DivArg) html.Node {
	instanceID := props.ID
	if instanceID == "" {
		instanceID = randomID("dialog")
	}

	divArgs := []html.DivArg{
		html.AData("pui-dialog", ""),
		html.AData("dialog-instance", instanceID),
		html.AClass(classnames.Merge("", props.Class)),
	}

	if props.DisableClickAway {
		divArgs = append(divArgs, html.AData("pui-dialog-disable-click-away", "true"))
	}
	if props.DisableESC {
		divArgs = append(divArgs, html.AData("pui-dialog-disable-esc", "true"))
	}
	if props.ID != "" {
		divArgs = append(divArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}
	divArgs = append(divArgs, args...)

	return html.Div(divArgs...).WithAssets("", dialogJS, "ui-dialog")
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
		buttonProps.Class = classnames.Merge(buttonProps.Class, triggerProps.Class)
	}
	buttonProps.Attrs = attrs
	if triggerProps.ID != "" {
		buttonProps.ID = triggerProps.ID
	}

	return button.Button(buttonProps, args...).WithAssets("", dialogJS, "ui-dialog")
}

// Content creates the dialog content panel
func Content(props ContentProps, args ...html.DivArg) html.Node {
	instanceID := props.ID
	if instanceID == "" {
		instanceID = randomID("dialog-content")
	}

	// Overlay/backdrop
	overlayClasses := classnames.Merge(
		"fixed inset-0 z-50 bg-black/50",
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
	contentClasses := classnames.Merge(
		// Base positioning
		"fixed z-50 left-[50%] top-[50%] translate-x-[-50%] translate-y-[-50%]",
		// Style
		"bg-background rounded-lg border shadow-lg",
		// Layout
		"grid gap-4 p-6",
		// Size
		"w-full max-w-[calc(100%-2rem)] sm:max-w-lg",
		// Transitions
		"transition-all duration-200",
		// Scale animation
		"data-[pui-dialog-open=false]:scale-95",
		"data-[pui-dialog-open=true]:scale-100",
		// Opacity
		"data-[pui-dialog-open=false]:opacity-0",
		"data-[pui-dialog-open=true]:opacity-100",
		// Pointer events
		"data-[pui-dialog-open=false]:pointer-events-none",
		"data-[pui-dialog-open=true]:pointer-events-auto",
		// Hidden state
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
			html.AClass(classnames.Merge(
				// Positioning
				"absolute top-4 right-4",
				// Style
				"rounded-xs opacity-70",
				// Interactions
				"transition-opacity hover:opacity-100",
				// Focus states
				"focus:outline-none focus:ring-2",
				"focus:ring-ring focus:ring-offset-2",
				"ring-offset-background",
				// Hover/Data states
				"data-[pui-dialog-open=true]:bg-accent",
				"data-[pui-dialog-open=true]:text-muted-foreground",
				// Disabled state
				"disabled:pointer-events-none",
				// Icon styles
				"[&_svg]:pointer-events-none",
				"[&_svg]:shrink-0",
				"[&_svg:not([class*='size-'])]:size-4",
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

// Close creates a dialog close trigger
func Close(props CloseProps, args ...html.SpanArg) html.Node {
	spanArgs := []html.SpanArg{
		html.AClass(classnames.Merge("contents cursor-pointer", props.Class)),
	}

	if props.ID != "" {
		spanArgs = append(spanArgs, html.AId(props.ID))
	}
	if props.For != "" {
		spanArgs = append(spanArgs, html.AData("pui-dialog-close", props.For))
	} else {
		spanArgs = append(spanArgs, html.AData("pui-dialog-close", ""))
	}

	for _, attr := range props.Attrs {
		spanArgs = append(spanArgs, attr)
	}
	spanArgs = append(spanArgs, args...)

	return html.Span(spanArgs...)
}

// Header creates a dialog header container
func Header(props HeaderProps, args ...html.DivArg) html.Node {
	divArgs := []html.DivArg{
		html.AClass(classnames.Merge("flex flex-col gap-2 text-center sm:text-left", props.Class)),
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

// Footer creates a dialog footer container
func Footer(props FooterProps, args ...html.DivArg) html.Node {
	divArgs := []html.DivArg{
		html.AClass(classnames.Merge("flex flex-col-reverse gap-2 sm:flex-row sm:justify-end", props.Class)),
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

// Title creates a dialog title
func Title(props TitleProps, args ...html.H2Arg) html.Node {
	h2Args := []html.H2Arg{
		html.AClass(classnames.Merge("text-lg leading-none font-semibold", props.Class)),
	}

	if props.ID != "" {
		h2Args = append(h2Args, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		h2Args = append(h2Args, attr)
	}
	h2Args = append(h2Args, args...)

	return html.H2(h2Args...)
}

// Description creates a dialog description
func Description(props DescriptionProps, args ...html.PArg) html.Node {
	pArgs := []html.PArg{
		html.AClass(classnames.Merge("text-muted-foreground text-sm", props.Class)),
	}

	if props.ID != "" {
		pArgs = append(pArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		pArgs = append(pArgs, attr)
	}
	pArgs = append(pArgs, args...)

	return html.P(pArgs...)
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
