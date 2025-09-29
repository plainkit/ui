package selectbox

import (
	"crypto/rand"
	"encoding/hex"
	"strconv"

	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/button"
	"github.com/plainkit/ui/input"
	"github.com/plainkit/ui/internal/classnames"
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
		args := []html.DivArg{html.AClass(classnames.Merge(append([]string{baseClass}, append(extra, p.Class)...)...))}
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
		args := []html.SpanArg{html.AClass(classnames.Merge(append([]string{baseClass}, append(extra, p.Class)...)...))}
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

	args := divArgsFromProps("select-container w-full relative")(p)
	args = append([]html.DivArg{html.AId(wrapperID)}, args...)

	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

func (p ValueProps) ApplySpan(attrs *html.SpanAttrs, children *[]html.Component) {
	args := spanArgsFromProps("block truncate select-value text-muted-foreground")(p)

	if p.Placeholder != "" {
		args = append(args, html.AData("pui-selectbox-placeholder", p.Placeholder))
	}

	for _, a := range args {
		a.ApplySpan(attrs, children)
	}
}

func (p GroupProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	args := []html.DivArg{
		html.AClass(classnames.Merge("p-1", p.Class)),
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
	args := []html.SpanArg{html.AClass(classnames.Merge("px-2 py-1.5 text-sm font-medium", p.Class))}

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
			html.AClass("pointer-events-none ml-1"),
			lucide.ChevronDown(html.AClass("size-4 text-muted-foreground")),
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
				Class: classnames.Merge(
					// Required class for JavaScript
					"select-trigger",
					// Base styles matching input
					"w-full h-9 px-3 py-1 text-base md:text-sm",
					"flex items-center justify-between",
					"rounded-md border border-input bg-transparent shadow-xs transition-[color,box-shadow] outline-none",
					// Dark mode background
					"dark:bg-input/30",
					// Selection styles
					"selection:bg-primary selection:text-primary-foreground",
					// Focus styles
					"focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px]",
					// Error/Invalid styles
					"aria-invalid:ring-destructive/20 aria-invalid:border-destructive dark:aria-invalid:ring-destructive/40",
					func() string {
						if props.HasError {
							return "border-destructive ring-destructive/20 dark:ring-destructive/40"
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
			html.AClass("sticky top-0 bg-popover p-1"),
			html.Div(
				html.AClass("relative"),
				html.Span(
					html.AClass("absolute left-3 top-1/2 -translate-y-1/2 text-muted-foreground z-10 pointer-events-none"),
					lucide.Search(html.AClass("size-4")),
				),
				input.Input(input.Props{
					Type:        input.TypeSearch,
					Class:       "pl-8",
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
		Class: classnames.Merge(
			"p-1 select-content z-50 overflow-hidden rounded-md border bg-popover text-popover-foreground shadow-md",
			"min-w-[var(--popover-trigger-width)] w-[var(--popover-trigger-width)]",
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
		html.AClass(classnames.Merge("p-1", props.Class)),
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
		html.AClass(classnames.Merge(
			"select-item relative flex w-full cursor-default select-none items-center rounded-sm py-1.5 px-2 text-sm font-light outline-none",
			"hover:bg-accent hover:text-accent-foreground",
			"focus:bg-accent focus:text-accent-foreground",
			func() string {
				if props.Selected {
					return "bg-accent text-accent-foreground"
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
	spanArgs := []html.SpanArg{html.AClass("truncate select-item-text")}
	spanArgs = append(spanArgs, args...)

	divArgs = append(divArgs,
		html.Span(spanArgs...),
		html.Span(
			html.AClass(classnames.Merge(
				"select-check absolute right-2 flex h-3.5 w-3.5 items-center justify-center",
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
