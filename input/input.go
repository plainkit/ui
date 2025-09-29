package input

import (
	"crypto/rand"
	_ "embed"
	"encoding/hex"

	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/button"
)

type Type string

const (
	TypeText     Type = "text"
	TypePassword Type = "password"
	TypeEmail    Type = "email"
	TypeNumber   Type = "number"
	TypeTel      Type = "tel"
	TypeURL      Type = "url"
	TypeSearch   Type = "search"
	TypeDate     Type = "date"
	TypeDateTime Type = "datetime-local"
	TypeTime     Type = "time"
	TypeFile     Type = "file"
	TypeColor    Type = "color"
	TypeWeek     Type = "week"
	TypeMonth    Type = "month"
)

type Props struct {
	ID                 string
	Class              string
	Attrs              []html.Global
	Name               string
	Type               Type
	Form               string
	Placeholder        string
	Value              string
	Disabled           bool
	Readonly           bool
	Required           bool
	FileAccept         string
	HasError           bool
	ShowPasswordToggle bool
}

func inputArgsFromProps(baseClass string, extra ...string) func(p Props) []html.InputArg {
	return func(p Props) []html.InputArg {
		inputType := p.Type
		if inputType == "" {
			inputType = TypeText
		}

		id := p.ID
		if id == "" {
			id = randomID()
		}

		extraPadding := ""
		if p.Type == TypePassword && p.ShowPasswordToggle {
			extraPadding = "pr-8"
		}

		errorClass := ""
		if p.HasError {
			errorClass = "border-destructive ring-destructive/20 dark:ring-destructive/40"
		}

		className := html.ClassMerge(
			append([]string{baseClass},
				append(extra,
					errorClass,
					extraPadding,
					p.Class,
				)...)...)

		args := []html.InputArg{
			html.AId(id),
			html.AType(string(inputType)),
			html.AClass(className),
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

// ApplyInput implements the html.InputArg interface for Props
func (p Props) ApplyInput(attrs *html.InputAttrs, children *[]html.Component) {
	args := inputArgsFromProps(
		"flex h-9 w-full min-w-0 rounded-md border border-input bg-transparent px-3 py-1 text-base shadow-xs transition-[color,box-shadow] outline-none md:text-sm",
		"dark:bg-input/30",
		"selection:bg-primary selection:text-primary-foreground",
		"placeholder:text-muted-foreground",
		"file:inline-flex file:h-7 file:border-0 file:bg-transparent file:text-sm file:font-medium file:text-foreground",
		"focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px]",
		"disabled:pointer-events-none disabled:cursor-not-allowed disabled:opacity-50",
		"aria-invalid:ring-destructive/20 aria-invalid:border-destructive dark:aria-invalid:ring-destructive/40",
	)(p)

	if p.Name != "" {
		args = append(args, html.AName(p.Name))
	}

	if p.Placeholder != "" {
		args = append(args, html.APlaceholder(p.Placeholder))
	}

	if p.Value != "" {
		args = append(args, html.AValue(p.Value))
	}

	if p.Type == TypeFile && p.FileAccept != "" {
		args = append(args, html.AAccept(p.FileAccept))
	}

	if p.Form != "" {
		args = append(args, html.AForm(p.Form))
	}

	if p.Disabled {
		args = append(args, html.ADisabled())
	}

	if p.Readonly {
		args = append(args, html.AReadonly())
	}

	if p.Required {
		args = append(args, html.ARequired())
	}

	if p.HasError {
		args = append(args, html.AAria("invalid", "true"))
	}

	for _, a := range args {
		a.ApplyInput(attrs, children)
	}
}

// Input renders a styled input control using the composable pattern.
// Accepts variadic html.InputArg arguments, with Props as an optional first argument.
func Input(args ...html.InputArg) html.Node {
	var (
		props Props
		rest  []html.InputArg
	)

	// Separate Props from other arguments
	for _, a := range args {
		if v, ok := a.(Props); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	if props.Type == "" {
		props.Type = TypeText
	}

	if props.ID == "" {
		props.ID = randomID()
	}

	children := []html.Component{html.Input(append([]html.InputArg{props}, rest...)...)}

	if props.Type == TypePassword && props.ShowPasswordToggle {
		children = append(children, passwordToggleButton(props.ID))
	}

	divArgs := []html.DivArg{html.AClass("relative w-full")}
	for _, child := range children {
		divArgs = append(divArgs, html.Child(child))
	}

	node := html.Div(divArgs...)
	if props.Type == TypePassword && props.ShowPasswordToggle {
		node = node.WithAssets("", passwordToggleJS, "ui-input-toggle")
	}

	return node
}

func passwordToggleButton(inputID string) html.Node {
	openIcon := html.Span(
		html.AClass("icon-open block"),
		lucide.Eye(html.AClass("size-4")),
	)
	closedIcon := html.Span(
		html.AClass("icon-closed hidden"),
		lucide.EyeOff(html.AClass("size-4")),
	)

	return button.Button(
		button.Props{
			Size:    button.SizeIcon,
			Variant: button.VariantGhost,
			Class:   "absolute right-0 top-1/2 -translate-y-1/2 opacity-50 cursor-pointer",
			Attrs: []html.Global{
				html.AData("pui-input-toggle-password", inputID),
			},
		},
		openIcon,
		closedIcon,
	)
}

func randomID() string {
	bytes := make([]byte, 6)
	if _, err := rand.Read(bytes); err != nil {
		return "input-id"
	}

	return "input-" + hex.EncodeToString(bytes)
}

//go:embed input.js
var passwordToggleJS string
