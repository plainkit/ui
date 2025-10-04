package tagsinput

import (
	"crypto/rand"
	_ "embed"
	"encoding/hex"

	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/badge"
	"github.com/plainkit/ui/input"
	"github.com/plainkit/ui/internal/styles"
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
		args := []html.DivArg{html.AClass(html.ClassMerge(append([]string{baseClass}, append(extra, p.Class)...)...))}
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

	containerClass := html.ClassMerge(
		styles.Input("flex w-full flex-wrap items-center gap-2 rounded-2xl border border-input/60 bg-background/60 py-2 pl-3 pr-10 text-sm transition-[border,box-shadow]"),
		"min-h-[2.75rem] cursor-text",
		func() string {
			if p.Disabled {
				return "cursor-not-allowed opacity-60"
			}

			return ""
		}(),
		func() string {
			if p.HasError {
				return "border-destructive ring-destructive/30"
			}

			return ""
		}(),
	)

	args := divArgsFromProps(html.ClassMerge(containerClass, p.Class))(p)
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
				html.AClass(styles.InteractiveGhost("ml-1 size-6 justify-center rounded-full text-muted-foreground/70", "hover:text-destructive")),
				html.AData("pui-tagsinput-remove", ""),
				func() html.ButtonArg {
					if p.Disabled {
						return html.ADisabled()
					}

					return html.AAria("", "")
				}(),
				lucide.X(html.AClass("pointer-events-none size-3")),
			),
		)
		tagChildren = append(tagChildren, tagBadge)
	}

	// Build the tags container
	tagsContainerArgs := []html.DivArg{
		html.AClass("flex flex-wrap items-center gap-2"),
		html.AData("pui-tagsinput-container", ""),
	}
	tagsContainerArgs = append(tagsContainerArgs, tagChildren...)
	tagsContainer := html.Div(tagsContainerArgs...)

	// Text input
	textInput := input.Input(input.Props{
		ID:          id,
		Class:       "min-h-0 grow border-0 bg-transparent px-0 py-1 text-sm shadow-none focus-visible:ring-0",
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
