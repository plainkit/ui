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

// TagsInput renders a tags input component for entering multiple tags
func TagsInput(props Props, args ...html.DivArg) html.Node {
	id := props.ID
	if id == "" {
		id = randomID("tagsinput")
	}

	containerArgs := []html.DivArg{
		html.AId(id + "-container"),
		html.AClass(classnames.Merge(
			// Base styles
			"flex items-center flex-wrap gap-2 p-2 rounded-md border border-input bg-transparent shadow-xs transition-[color,box-shadow] outline-none",
			// Dark mode background
			"dark:bg-input/30",
			// Focus styles
			"focus-within:border-ring focus-within:ring-ring/50 focus-within:ring-[3px]",
			// Disabled styles
			func() string {
				if props.Disabled {
					return "opacity-50 cursor-not-allowed"
				}

				return ""
			}(),
			// Width
			"w-full",
			// Error/Invalid styles
			func() string {
				if props.HasError {
					return "border-destructive ring-destructive/20 dark:ring-destructive/40"
				}

				return ""
			}(),
			props.Class,
		)),
		html.AData("pui-tagsinput", ""),
		html.AData("pui-tagsinput-name", props.Name),
		html.AData("pui-tagsinput-form", props.Form),
	}

	for _, attr := range props.Attrs {
		containerArgs = append(containerArgs, attr)
	}

	// Add existing tags as children
	tagChildren := make([]html.DivArg, 0, len(props.Value))
	for _, tag := range props.Value {
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
					if props.Disabled {
						return html.ADisabled()
					}

					return html.AAria("", "")
				}(),
				lucide.X(html.AClass("h-3 w-3 pointer-events-none")),
			),
		)
		tagChildren = append(tagChildren, tagBadge)
	}

	// Build the full tagsContainer args
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
		Placeholder: props.Placeholder,
		Disabled:    props.Disabled,
		Readonly:    props.Readonly,
		Attrs: []html.Global{
			html.AData("pui-tagsinput-text-input", ""),
		},
	})

	// Add existing hidden inputs
	hiddenInputChildren := make([]html.DivArg, 0, len(props.Value))
	for _, tag := range props.Value {
		hiddenInput := html.Input(
			html.AType("hidden"),
			html.AName(props.Name),
			html.AValue(tag),
		)
		hiddenInputChildren = append(hiddenInputChildren, hiddenInput)
	}

	// Build the full hiddenInputsContainer args
	hiddenInputsArgs := []html.DivArg{
		html.AData("pui-tagsinput-hidden-inputs", ""),
	}
	hiddenInputsArgs = append(hiddenInputsArgs, hiddenInputChildren...)
	hiddenInputsContainer := html.Div(hiddenInputsArgs...)

	containerArgs = append(containerArgs,
		tagsContainer,
		textInput,
		hiddenInputsContainer,
	)
	containerArgs = append(containerArgs, args...)

	return html.Div(containerArgs...).WithAssets("", tagsinputJS, "ui-tagsinput")
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
