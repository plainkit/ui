package tagsinput

import (
	"crypto/rand"
	_ "embed"
	"encoding/hex"

	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/badge"
	"github.com/plainkit/ui/input"
	"github.com/plainkit/ui/internal/classnames"
)

type Props struct {
	ID          string
	Name        string
	Value       []string
	Form        string
	Placeholder string
	Class       string
	Attrs       []html.Global
	HasError    bool
	Disabled    bool
	Readonly    bool
}

func divArgsFromProps(baseClass string, extra ...string) func(p Props) []html.DivArg {
	return func(p Props) []html.DivArg {
		args := []html.DivArg{html.AClass(classnames.Merge(append([]string{baseClass}, append(extra, p.Class)...)...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID+"-container"))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p Props) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	id := p.ID
	if id == "" {
		id = randomID("tagsinput")
	}

	containerClass := classnames.Merge(
		// Base styles
		"flex items-center flex-wrap gap-2 p-2 rounded-md border border-input bg-transparent shadow-xs transition-[color,box-shadow] outline-none",
		// Dark mode background
		"dark:bg-input/30",
		// Focus styles
		"focus-within:border-ring focus-within:ring-ring/50 focus-within:ring-[3px]",
		// Disabled styles
		func() string {
			if p.Disabled {
				return "opacity-50 cursor-not-allowed"
			}

			return ""
		}(),
		// Width
		"w-full",
		// Error/Invalid styles
		func() string {
			if p.HasError {
				return "border-destructive ring-destructive/20 dark:ring-destructive/40"
			}

			return ""
		}(),
		p.Class,
	)

	args := divArgsFromProps(containerClass)(p)
	args = append([]html.DivArg{
		html.AId(id + "-container"),
		html.AData("pui-tagsinput", ""),
		html.AData("pui-tagsinput-name", p.Name),
		html.AData("pui-tagsinput-form", p.Form),
	}, args...)

	// Add existing tags as children
	tagChildren := make([]html.DivArg, 0, len(p.Value))
	for _, tag := range p.Value {
		tagBadge := badge.Badge(badge.Props{
			Attrs: []html.Global{
				html.AData("pui-tagsinput-chip", ""),
			},
		},
			html.Span(html.Text(tag)),
			html.Button(
				html.AType("button"),
				html.AClass("ml-1 text-current hover:text-destructive disabled:opacity-50 disabled:cursor-not-allowed cursor-pointer"),
				html.AData("pui-tagsinput-remove", ""),
				func() html.ButtonArg {
					if p.Disabled {
						return html.ADisabled()
					}

					return html.AAria("", "")
				}(),
				lucide.X(html.AClass("h-3 w-3 pointer-events-none")),
			),
		)
		tagChildren = append(tagChildren, tagBadge)
	}

	// Build the tags container
	tagsContainerArgs := []html.DivArg{
		html.AClass("flex items-center flex-wrap gap-2"),
		html.AData("pui-tagsinput-container", ""),
	}
	tagsContainerArgs = append(tagsContainerArgs, tagChildren...)
	tagsContainer := html.Div(tagsContainerArgs...)

	// Text input
	textInput := input.Input(input.Props{
		ID:          id,
		Class:       "border-0 shadow-none focus-visible:ring-0 h-auto py-0 px-0 bg-transparent rounded-none min-h-0 disabled:opacity-100 dark:bg-transparent",
		Type:        input.TypeText,
		Placeholder: p.Placeholder,
		Disabled:    p.Disabled,
		Readonly:    p.Readonly,
		Attrs: []html.Global{
			html.AData("pui-tagsinput-text-input", ""),
		},
	})

	// Add existing hidden inputs
	hiddenInputChildren := make([]html.DivArg, 0, len(p.Value))
	for _, tag := range p.Value {
		hiddenInput := html.Input(
			html.AType("hidden"),
			html.AName(p.Name),
			html.AValue(tag),
		)
		hiddenInputChildren = append(hiddenInputChildren, hiddenInput)
	}

	// Build the hidden inputs container
	hiddenContainerArgs := []html.DivArg{
		html.AData("pui-tagsinput-hidden-inputs", ""),
	}
	hiddenContainerArgs = append(hiddenContainerArgs, hiddenInputChildren...)
	hiddenInputsContainer := html.Div(hiddenContainerArgs...)

	*children = append(*children,
		tagsContainer,
		textInput,
		hiddenInputsContainer,
	)

	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

// TagsInput renders a tags input component for entering multiple tags
func TagsInput(args ...html.DivArg) html.Node {
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

	return html.Div(append([]html.DivArg{props}, rest...)...).WithAssets("", tagsinputJS, "ui-tagsinput")
}

func randomID(prefix string) string {
	buf := make([]byte, 4)
	if _, err := rand.Read(buf); err != nil {
		return prefix + "-id"
	}

	return prefix + "-" + hex.EncodeToString(buf)
}

//go:embed tagsinput.js
var tagsinputJS string
