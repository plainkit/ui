package toast

import (
	"crypto/rand"
	_ "embed"
	"encoding/hex"
	"strconv"

	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/button"
	"github.com/plainkit/ui/internal/classnames"
)

type Variant string
type Position string

type Props struct {
	ID            string
	Class         string
	Attrs         []html.Global
	Title         string
	Description   string
	Variant       Variant
	Position      Position
	Duration      int
	Dismissible   bool
	ShowIndicator bool
	Icon          bool
}

type TriggerProps struct {
	ID    string
	Class string
	Attrs []html.Global
	Toast Props // The toast configuration to spawn
}

const (
	VariantDefault Variant = "default"
	VariantSuccess Variant = "success"
	VariantError   Variant = "error"
	VariantWarning Variant = "warning"
	VariantInfo    Variant = "info"
)

const (
	PositionTopRight     Position = "top-right"
	PositionTopLeft      Position = "top-left"
	PositionTopCenter    Position = "top-center"
	PositionBottomRight  Position = "bottom-right"
	PositionBottomLeft   Position = "bottom-left"
	PositionBottomCenter Position = "bottom-center"
)

// Toast renders an interactive toast notification container with optional auto-dismiss logic.
func Toast(props Props, args ...html.DivArg) html.Node {
	variant := props.Variant
	if variant == "" {
		variant = VariantDefault
	}

	position := props.Position
	if position == "" {
		position = PositionBottomRight
	}

	duration := props.Duration
	if duration == 0 {
		duration = 3000
	}

	id := props.ID
	if id == "" {
		id = randomID("toast")
	}

	divArgs := []html.DivArg{
		html.AId(id),
		html.AClass(classnames.Merge(
			"z-50 fixed pointer-events-auto p-4 w-full md:max-w-[420px]",
			"animate-in fade-in slide-in-from-bottom-4 duration-300",
			"data-[position=top-right]:top-0 data-[position=top-right]:right-0",
			"data-[position=top-left]:top-0 data-[position=top-left]:left-0",
			"data-[position=top-center]:top-0 data-[position=top-center]:left-1/2 data-[position=top-center]:-translate-x-1/2",
			"data-[position=bottom-right]:bottom-0 data-[position=bottom-right]:right-0",
			"data-[position=bottom-left]:bottom-0 data-[position=bottom-left]:left-0",
			"data-[position=bottom-center]:bottom-0 data-[position=bottom-center]:left-1/2 data-[position=bottom-center]:-translate-x-1/2",
			props.Class,
		)),
		html.AData("pui-toast", ""),
		html.AData("pui-toast-duration", strconv.Itoa(duration)),
		html.AData("position", string(position)),
		html.AData("variant", string(variant)),
	}
	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}

	innerDivArgs := []html.DivArg{}

	if props.ShowIndicator && duration > 0 {
		progressBar := html.Div(
			html.AClass("absolute top-0 left-0 right-0 h-1 overflow-hidden"),
			html.Div(
				html.AClass(classnames.Merge(
					"toast-progress h-full origin-left transition-transform ease-linear",
					"data-[variant=default]:bg-gray-500",
					"data-[variant=success]:bg-green-500",
					"data-[variant=error]:bg-red-500",
					"data-[variant=warning]:bg-yellow-500",
					"data-[variant=info]:bg-blue-500",
				)),
				html.AData("variant", string(variant)),
				html.AData("duration", strconv.Itoa(duration)),
			),
		)
		innerDivArgs = append(innerDivArgs, progressBar)
	}

	contentDivArgs := []html.DivArg{html.AClass("w-full bg-popover text-popover-foreground rounded-lg shadow-xs border pt-5 pb-4 px-4 flex items-center justify-center relative overflow-hidden group gap-3")}

	if props.Icon && variant != VariantDefault {
		iconNode := variantIcon(variant)
		contentDivArgs = append(contentDivArgs, iconNode)
	}

	textSpanArgs := []html.SpanArg{html.AClass("flex-1 min-w-0")}
	if props.Title != "" {
		textSpanArgs = append(textSpanArgs, html.P(
			html.AClass("text-sm font-semibold truncate"),
			html.Text(props.Title),
		))
	}

	if props.Description != "" {
		textSpanArgs = append(textSpanArgs, html.P(
			html.AClass("text-sm opacity-90 mt-1"),
			html.Text(props.Description),
		))
	}

	textContainer := html.Span(textSpanArgs...)
	contentDivArgs = append(contentDivArgs, textContainer)

	if props.Dismissible {
		btn := button.Button(button.Props{
			Variant: button.VariantGhost,
			Size:    button.SizeIcon,
			Attrs: []html.Global{
				html.AAria("label", "Close"),
				html.AData("pui-toast-dismiss", ""),
			},
		}, lucide.X(
			html.AClass("size-4 opacity-75 hover:opacity-100"),
		))
		contentDivArgs = append(contentDivArgs, btn)
	}

	contentDiv := html.Div(contentDivArgs...)
	innerDivArgs = append(innerDivArgs, contentDiv)

	divArgs = append(divArgs, innerDivArgs...)
	divArgs = append(divArgs, args...)

	node := html.Div(divArgs...)

	return node.WithAssets("", toastJS, "ui-toast")
}

func variantIcon(variant Variant) html.Node {
	switch variant {
	case VariantSuccess:
		return lucide.CircleCheck(html.AClass("size-5 text-green-500 flex-shrink-0"))
	case VariantError:
		return lucide.CircleX(html.AClass("size-5 text-destructive flex-shrink-0"))
	case VariantWarning:
		return lucide.TriangleAlert(html.AClass("size-5 text-yellow-500 flex-shrink-0"))
	case VariantInfo:
		return lucide.Info(html.AClass("size-5 text-blue-500 flex-shrink-0"))
	default:
		// Return an empty span as a placeholder
		return html.Span()
	}
}

func randomID(prefix string) string {
	buf := make([]byte, 4)
	if _, err := rand.Read(buf); err != nil {
		return prefix + "-id"
	}

	return prefix + "-" + hex.EncodeToString(buf)
}

// ToastTrigger creates a button that spawns a toast when clicked
func ToastTrigger(props TriggerProps, buttonProps button.Props, args ...html.ButtonArg) html.Node {
	id := props.ID
	if id == "" {
		id = randomID("toast-trigger")
	}

	// Encode toast configuration in data attributes
	attrs := buttonProps.Attrs
	if attrs == nil {
		attrs = []html.Global{}
	}

	attrs = append(attrs,
		html.AData("pui-toast-trigger", ""),
		html.AData("toast-title", props.Toast.Title),
		html.AData("toast-description", props.Toast.Description),
		html.AData("toast-variant", string(props.Toast.Variant)),
		html.AData("toast-position", string(props.Toast.Position)),
		html.AData("toast-duration", strconv.Itoa(props.Toast.Duration)),
		html.AData("toast-dismissible", strconv.FormatBool(props.Toast.Dismissible)),
		html.AData("toast-show-indicator", strconv.FormatBool(props.Toast.ShowIndicator)),
		html.AData("toast-icon", strconv.FormatBool(props.Toast.Icon)),
	)

	attrs = append(attrs, props.Attrs...)

	buttonProps.Attrs = attrs

	buttonProps.ID = id
	if props.Class != "" {
		buttonProps.Class = classnames.Merge(buttonProps.Class, props.Class)
	}

	return button.Button(append([]html.ButtonArg{buttonProps}, args...)...).WithAssets("", toastJS, "ui-toast")
}

// ToastContainer creates a container for toasts to be spawned into
func ToastContainer(position Position) html.Node {
	if position == "" {
		position = PositionBottomRight
	}

	positionClasses := map[Position]string{
		PositionTopRight:     "top-0 right-0",
		PositionTopLeft:      "top-0 left-0",
		PositionTopCenter:    "top-0 left-1/2 -translate-x-1/2",
		PositionBottomRight:  "bottom-0 right-0",
		PositionBottomLeft:   "bottom-0 left-0",
		PositionBottomCenter: "bottom-0 left-1/2 -translate-x-1/2",
	}

	return html.Div(
		html.AId("toast-container-"+string(position)),
		html.AClass(classnames.Merge(
			"fixed z-50 pointer-events-none p-4",
			positionClasses[position],
		)),
		html.AData("pui-toast-container", string(position)),
	).WithAssets("", toastJS, "ui-toast")
}

//go:embed toast.js
var toastJS string
