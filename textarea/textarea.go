package textarea

import (
	"crypto/rand"
	_ "embed"
	"encoding/hex"
	"strconv"

	"github.com/plainkit/html"
)

type Props struct {
	ID          string
	Class       string
	Attrs       []html.Global
	Name        string
	Value       string
	Form        string
	Placeholder string
	Rows        int
	AutoResize  bool
	Disabled    bool
	Required    bool
	Readonly    bool
	HasError    bool
}

func textareaArgsFromProps(baseClass string, extra ...string) func(p Props) []html.TextareaArg {
	return func(p Props) []html.TextareaArg {
		id := p.ID
		if id == "" {
			id = randomID()
		}

		autoResizeExtra := ""
		if p.AutoResize {
			autoResizeExtra = "overflow-hidden resize-none"
		}

		errorClass := ""
		if p.HasError {
			errorClass = "border-destructive ring-destructive/20 dark:ring-destructive/40"
		}

		className := html.ClassMerge(
			append([]string{baseClass},
				append(extra,
					errorClass,
					autoResizeExtra,
					p.Class,
				)...)...)

		args := []html.TextareaArg{
			html.AId(id),
			html.AClass(className),
			html.AData("pui-textarea", ""),
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

// ApplyTextarea implements the html.TextareaArg interface for Props
func (p Props) ApplyTextarea(attrs *html.TextareaAttrs, children *[]html.Component) {
	args := textareaArgsFromProps(
		"flex w-full min-w-0 rounded-md border border-input bg-transparent px-3 py-1 text-base shadow-xs transition-[color,box-shadow] outline-none md:text-sm",
		"min-h-[80px]",
		"dark:bg-input/30",
		"selection:bg-primary selection:text-primary-foreground",
		"placeholder:text-muted-foreground",
		"focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px]",
		"disabled:pointer-events-none disabled:cursor-not-allowed disabled:opacity-50",
		"aria-invalid:ring-destructive/20 aria-invalid:border-destructive dark:aria-invalid:ring-destructive/40",
	)(p)

	if p.Name != "" {
		args = append(args, html.AName(p.Name))
	}

	if p.Form != "" {
		args = append(args, html.AForm(p.Form))
	}

	if p.Placeholder != "" {
		args = append(args, html.APlaceholder(p.Placeholder))
	}

	if p.Rows > 0 {
		args = append(args, html.ARows(strconv.Itoa(p.Rows)))
	}

	if p.Disabled {
		args = append(args, html.ADisabled())
	}

	if p.Required {
		args = append(args, html.ARequired())
	}

	if p.Readonly {
		args = append(args, html.AReadonly())
	}

	if p.HasError {
		args = append(args, html.AAria("invalid", "true"))
	}

	if p.AutoResize {
		args = append(args, html.AData("pui-textarea-auto-resize", "true"))
	}

	for _, a := range args {
		a.ApplyTextarea(attrs, children)
	}
}

// Textarea renders a styled multi-line input using the composable pattern.
// Accepts variadic html.TextareaArg arguments, with Props as an optional first argument.
func Textarea(args ...html.TextareaArg) html.Node {
	var (
		props Props
		rest  []html.TextareaArg
	)

	// Separate Props from other arguments
	for _, a := range args {
		if v, ok := a.(Props); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	if props.ID == "" {
		props.ID = randomID()
	}

	// Add the value as text content if provided
	if props.Value != "" {
		rest = append(rest, html.Text(props.Value))
	}

	node := html.Textarea(append([]html.TextareaArg{props}, rest...)...)
	if props.AutoResize {
		node = node.WithAssets("", textareaResizeJS, "ui-textarea-autoresize")
	}

	return node
}

func randomID() string {
	buf := make([]byte, 4)
	if _, err := rand.Read(buf); err != nil {
		return "textarea-id"
	}

	return "textarea-" + hex.EncodeToString(buf)
}

//go:embed textarea.js
var textareaResizeJS string
