package textarea

import (
	"crypto/rand"
	_ "embed"
	"encoding/hex"
	"strconv"

	"github.com/plainkit/html"
	"github.com/plainkit/ui/internal/classnames"
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

// Textarea renders a styled multi-line input.
func Textarea(props Props, extra ...html.TextareaArg) html.Node {
	if props.ID == "" {
		props.ID = randomID()
	}

	args := []html.TextareaArg{
		html.AId(props.ID),
		html.AClass(textareaClass(props)),
		html.AData("pui-textarea", ""),
	}
	if props.Name != "" {
		args = append(args, html.AName(props.Name))
	}
	if props.Form != "" {
		args = append(args, html.AForm(props.Form))
	}
	if props.Placeholder != "" {
		args = append(args, html.APlaceholder(props.Placeholder))
	}
	if props.Rows > 0 {
		args = append(args, html.ARows(strconv.Itoa(props.Rows)))
	}
	if props.Disabled {
		args = append(args, html.ADisabled())
	}
	if props.Required {
		args = append(args, html.ARequired())
	}
	if props.Readonly {
		args = append(args, html.AReadonly())
	}
	if props.HasError {
		args = append(args, html.AAria("invalid", "true"))
	}
	if props.AutoResize {
		args = append(args, html.AData("pui-textarea-auto-resize", "true"))
	}
	for _, attr := range props.Attrs {
		args = append(args, attr)
	}
	args = append(args, extra...)

	if props.Value != "" {
		args = append(args, html.Text(props.Value))
	}

	node := html.Textarea(args...)
	if props.AutoResize {
		node = node.WithAssets("", textareaResizeJS, "ui-textarea-autoresize")
	}
	return node
}

func textareaClass(props Props) string {
	extra := ""
	if props.AutoResize {
		extra = "overflow-hidden resize-none"
	}
	error := ""
	if props.HasError {
		error = "border-destructive ring-destructive/20 dark:ring-destructive/40"
	}

	return classnames.Merge(
		"flex w-full min-w-0 rounded-md border border-input bg-transparent px-3 py-1 text-base shadow-xs transition-[color,box-shadow] outline-none md:text-sm",
		"min-h-[80px]",
		"dark:bg-input/30",
		"selection:bg-primary selection:text-primary-foreground",
		"placeholder:text-muted-foreground",
		"focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px]",
		"disabled:pointer-events-none disabled:cursor-not-allowed disabled:opacity-50",
		"aria-invalid:ring-destructive/20 aria-invalid:border-destructive dark:aria-invalid:ring-destructive/40",
		error,
		extra,
		props.Class,
	)
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
