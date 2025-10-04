package selectbox

import (
	"crypto/rand"
	"encoding/hex"
	"strconv"

	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/button"
	"github.com/plainkit/ui/input"
	"github.com/plainkit/ui/internal/styles"
	"github.com/plainkit/ui/popover"
)

type Props struct {
	ID       string
	Class    string
	Attrs    []html.Global
	Multiple bool
}

type TriggerProps struct {
	ID                string
	Class             string
	Attrs             []html.Global
	Name              string
	Form              string
	Required          bool
	Disabled          bool
	HasError          bool
	Multiple          bool
	ShowPills         bool
	SelectedCountText string
}

type ValueProps struct {
	ID          string
	Class       string
	Attrs       []html.Global
	Placeholder string
	Multiple    bool
}

type ContentProps struct {
	ID                string
	Class             string
	Attrs             []html.Global
	NoSearch          bool
	SearchPlaceholder string
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
	ID       string
	Class    string
	Attrs    []html.Global
	Value    string
	Selected bool
	Disabled bool
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

func spanArgsFromProps(baseClass string, extra ...string) func(p ValueProps) []html.SpanArg {
	return func(p ValueProps) []html.SpanArg {
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

func (p Props) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	wrapperID := p.ID
	if wrapperID == "" {
		wrapperID = randomID("selectbox")
	}

	args := divArgsFromProps("select-container relative w-full space-y-2")(p)
	args = append([]html.DivArg{html.AId(wrapperID)}, args...)

	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

func (p ValueProps) ApplySpan(attrs *html.SpanAttrs, children *[]html.Component) {
	args := spanArgsFromProps(styles.SubtleText("block truncate select-value text-left"))(p)

	if p.Placeholder != "" {
		args = append(args, html.AData("pui-selectbox-placeholder", p.Placeholder))
	}

	for _, a := range args {
		a.ApplySpan(attrs, children)
	}
}

func (p GroupProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	args := []html.DivArg{
		html.AClass(html.ClassMerge("p-1", p.Class)),
		html.AAria("role", "group"),
	}

	if p.ID != "" {
		args = append(args, html.AId(p.ID))
	}

	for _, a := range p.Attrs {
		args = append(args, a)
	}

	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

func (p LabelProps) ApplySpan(attrs *html.SpanAttrs, children *[]html.Component) {
	args := []html.SpanArg{html.AClass(html.ClassMerge(styles.SubHeading("px-3 py-2 text-xs uppercase tracking-wide text-muted-foreground/70"), p.Class))}

	if p.ID != "" {
		args = append(args, html.AId(p.ID))
	}

	for _, a := range p.Attrs {
		args = append(args, a)
	}

	for _, a := range args {
		a.ApplySpan(attrs, children)
	}
}

// SelectBox renders a select box container
func SelectBox(args ...html.DivArg) html.Node {
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

// Trigger creates a select box trigger button
func Trigger(props TriggerProps, contentID string, args ...html.Node) html.Node {
	if contentID == "" {
		contentID = randomID("selectbox-content")
	}

	if props.ShowPills {
		props.Multiple = true
	}

	// Hidden input for form submission
	inputArgs := []html.InputArg{html.AType("hidden")}
	if props.Name != "" {
		inputArgs = append(inputArgs, html.AName(props.Name))
	}

	if props.Form != "" {
		inputArgs = append(inputArgs, html.AForm(props.Form))
	}

	if props.Required {
		inputArgs = append(inputArgs, html.ARequired())
	}

	for _, attr := range props.Attrs {
		inputArgs = append(inputArgs, attr)
	}

	hiddenInput := html.Input(inputArgs...)

	// Button content with children and chevron
	buttonContent := make([]html.ButtonArg, 0, len(args)+2)

	buttonContent = append(buttonContent, hiddenInput)
	for _, arg := range args {
		buttonContent = append(buttonContent, arg)
	}

	buttonContent = append(buttonContent,
		html.Span(
			html.AClass("pointer-events-none ml-auto pl-2 text-muted-foreground/70"),
			lucide.ChevronDown(html.AClass("size-4")),
		),
	)

	return popover.Trigger(
		popover.TriggerProps{
			For:         contentID,
			TriggerType: popover.TriggerTypeClick,
		},
		button.Button(append([]html.ButtonArg{
			button.Props{
				ID:      props.ID,
				Type:    button.TypeButton,
				Variant: button.VariantOutline,
				Class: html.ClassMerge(
					styles.Input(
						"select-trigger flex h-11 w-full items-center justify-between gap-3",
						"cursor-pointer text-left text-sm md:text-base",
						"pr-10",
						"transition-transform hover:-translate-y-0.5",
					),
					"aria-invalid:border-destructive aria-invalid:ring-destructive/30",
					func() string {
						if props.HasError {
							return "border-destructive ring-destructive/30"
						}

						return ""
					}(),
					props.Class,
				),
				Disabled: props.Disabled,
				Attrs: []html.Global{
					html.AData("pui-selectbox-content-id", contentID),
					html.AData("pui-selectbox-multiple", strconv.FormatBool(props.Multiple)),
					html.AData("pui-selectbox-show-pills", strconv.FormatBool(props.ShowPills)),
					html.AData("pui-selectbox-selected-count-text", props.SelectedCountText),
					html.ATabindex(0),
					func() html.Global {
						if props.Required {
							return html.AAria("required", "true")
						}

						return html.AAria("", "")
					}(),
					func() html.Global {
						if props.HasError {
							return html.AAria("invalid", "true")
						}

						return html.AAria("", "")
					}(),
				},
			},
		}, buttonContent...)...),
	)
}

// Value creates a select box value display
func Value(args ...html.SpanArg) html.Node {
	var (
		props ValueProps
		rest  []html.SpanArg
	)

	for _, a := range args {
		if v, ok := a.(ValueProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	if props.Placeholder != "" && len(rest) == 0 {
		rest = append(rest, html.Text(props.Placeholder))
	}

	return html.Span(append([]html.SpanArg{props}, rest...)...)
}

// Content creates the select box dropdown content
func Content(props ContentProps, args ...html.DivArg) html.Node {
	contentID := props.ID
	if contentID == "" {
		contentID = randomID("selectbox-content")
	}

	contentArgs := []html.DivArg{
		html.AClass("max-h-[300px] overflow-y-auto"),
	}
	contentArgs = append(contentArgs, args...)

	var popoverContent []html.DivArg

	if !props.NoSearch {
		searchPlaceholder := "Search..."
		if props.SearchPlaceholder != "" {
			searchPlaceholder = props.SearchPlaceholder
		}

		searchDiv := html.Div(
			html.AClass("sticky top-0 z-10 -mx-2 -mt-2 bg-popover/95 p-2 backdrop-blur supports-[backdrop-filter]:bg-popover/80"),
			html.Div(
				html.AClass("relative"),
				html.Span(
					html.AClass("pointer-events-none absolute left-3 top-1/2 z-10 -translate-y-1/2 text-muted-foreground/70"),
					lucide.Search(html.AClass("size-4")),
				),
				input.Input(input.Props{
					Type:        input.TypeSearch,
					Class:       "pl-10",
					Placeholder: searchPlaceholder,
					Attrs: []html.Global{
						html.AData("pui-selectbox-search", ""),
					},
				}),
			),
		)
		popoverContent = append(popoverContent, searchDiv)
	}

	popoverContent = append(popoverContent, html.Div(contentArgs...))

	contentProps := popover.ContentProps{
		ID:         contentID,
		Placement:  popover.PlacementBottomStart,
		Offset:     4,
		MatchWidth: true,
		Class: html.ClassMerge(
			styles.Panel(
				"select-content z-50 overflow-hidden p-2",
				"min-w-[var(--popover-trigger-width)] w-[var(--popover-trigger-width)]",
			),
			props.Class,
		),
		Attrs: []html.Global{
			html.AAria("role", "listbox"),
			html.ATabindex(-1),
		},
	}

	return popover.Content(append([]html.DivArg{contentProps}, popoverContent...)...)
}

// Group creates a select box option group
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

	divArgs := []html.DivArg{
		html.AClass(html.ClassMerge("p-1", props.Class)),
		html.AAria("role", "group"),
	}

	if props.ID != "" {
		divArgs = append(divArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}

	divArgs = append(divArgs, rest...)

	return html.Div(divArgs...)
}

// Label creates a select box group label
func Label(args ...html.SpanArg) html.Node {
	var (
		props LabelProps
		rest  []html.SpanArg
	)

	for _, a := range args {
		if v, ok := a.(LabelProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Span(append([]html.SpanArg{props}, rest...)...)
}

// Item creates a select box option item
func Item(props ItemProps, args ...html.SpanArg) html.Node {
	divArgs := []html.DivArg{
		html.AClass(html.ClassMerge(
			styles.InteractiveGhost(
				"select-item relative flex w-full cursor-pointer select-none items-center gap-3",
				"rounded-lg px-3 py-2 text-sm",
				"justify-between",
			),
			"focus-visible:ring-0",
			func() string {
				if props.Selected {
					return "bg-accent/90 text-accent-foreground"
				}

				return ""
			}(),
			func() string {
				if props.Disabled {
					return "pointer-events-none opacity-50"
				}

				return ""
			}(),
			props.Class,
		)),
		html.AAria("role", "option"),
		html.AData("pui-selectbox-value", props.Value),
		html.AData("pui-selectbox-selected", strconv.FormatBool(props.Selected)),
		html.AData("pui-selectbox-disabled", strconv.FormatBool(props.Disabled)),
		html.ATabindex(0),
	}

	if props.ID != "" {
		divArgs = append(divArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}

	// Add text content and check icon
	spanArgs := []html.SpanArg{html.AClass(styles.SubtleText("truncate select-item-text text-sm"))}
	spanArgs = append(spanArgs, args...)

	divArgs = append(divArgs,
		html.Span(spanArgs...),
		html.Span(
			html.AClass(html.ClassMerge(
				"select-check absolute right-3 flex h-4 w-4 items-center justify-center text-primary",
				"transition-opacity duration-150",
				func() string {
					if props.Selected {
						return "opacity-100"
					}

					return "opacity-0"
				}(),
			)),
			lucide.Check(html.AClass("size-4")),
		),
	)

	return html.Div(divArgs...)
}

func randomID(prefix string) string {
	buf := make([]byte, 4)
	if _, err := rand.Read(buf); err != nil {
		return prefix + "-id"
	}

	return prefix + "-" + hex.EncodeToString(buf)
}
