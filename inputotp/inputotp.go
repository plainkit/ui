package inputotp

import (
	_ "embed"
	"strconv"

	"github.com/plainkit/html"
)

type Props struct {
	ID        string
	Class     string
	Attrs     []html.Global
	Value     string
	Required  bool
	Name      string
	Form      string
	HasError  bool
	Autofocus bool
}

type GroupProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type SlotProps struct {
	ID          string
	Class       string
	Attrs       []html.Global
	Index       int
	Type        string
	Placeholder string
	Disabled    bool
	HasError    bool
}

type SeparatorProps struct {
	ID    string
	Class string
	Attrs []html.Global
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
	args := divArgsFromProps("flex flex-row items-center gap-2 w-fit")(p)
	args = append(args, html.AData("pui-inputotp", ""))

	if p.Value != "" {
		args = append(args, html.AData("pui-inputotp-value", p.Value))
	}

	if p.Autofocus {
		args = append(args, html.AAutofocus())
	}

	// Create hidden input
	hiddenArgs := []html.InputArg{
		html.AType("hidden"),
		html.AData("pui-inputotp-value-target", ""),
	}
	if p.ID != "" {
		hiddenArgs = append(hiddenArgs, html.AId(p.ID))
	}

	if p.Name != "" {
		hiddenArgs = append(hiddenArgs, html.AName(p.Name))
	}

	if p.Form != "" {
		hiddenArgs = append(hiddenArgs, html.AForm(p.Form))
	}

	if p.HasError {
		hiddenArgs = append(hiddenArgs, html.AAria("invalid", "true"))
	}

	if p.Required {
		hiddenArgs = append(hiddenArgs, html.ARequired())
	}

	hidden := html.Input(hiddenArgs...)
	*children = append(*children, hidden)

	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

func (p GroupProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	args := []html.DivArg{html.AClass(html.ClassMerge("flex gap-2", p.Class))}
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

func (p SlotProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	inputType := p.Type
	if inputType == "" {
		inputType = "text"
	}

	args := []html.DivArg{html.AClass("relative")}
	if p.ID != "" {
		args = append(args, html.AId(p.ID))
	}

	for _, a := range p.Attrs {
		args = append(args, a)
	}

	classes := []string{
		"w-10 h-12 text-center rounded-md border border-input bg-transparent text-base shadow-xs transition-[color,box-shadow] outline-none md:text-sm",
		"dark:bg-input/30",
		"selection:bg-primary selection:text-primary-foreground",
		"placeholder:text-muted-foreground",
		"focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px]",
		"disabled:pointer-events-none disabled:cursor-not-allowed disabled:opacity-50",
		"aria-invalid:ring-destructive/20 aria-invalid:border-destructive dark:aria-invalid:ring-destructive/40",
	}
	if p.HasError {
		classes = append(classes, "border-destructive ring-destructive/20 dark:ring-destructive/40")
	}

	classes = append(classes, p.Class)

	inputArgs := []html.InputArg{
		html.AType(inputType),
		html.AInputmode("numeric"),
		html.AMaxlength("1"),
		html.AClass(html.ClassMerge(classes...)),
		html.AData("pui-inputotp-index", strconv.Itoa(p.Index)),
		html.AData("pui-inputotp-slot", ""),
	}
	if p.Placeholder != "" {
		inputArgs = append(inputArgs, html.APlaceholder(p.Placeholder))
	}

	if p.Disabled {
		inputArgs = append(inputArgs, html.ADisabled())
	}

	if p.HasError {
		inputArgs = append(inputArgs, html.AAria("invalid", "true"))
	}

	input := html.Input(inputArgs...)
	*children = append(*children, input)

	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

func (p SeparatorProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	args := []html.DivArg{html.AClass(html.ClassMerge("flex items-center text-muted-foreground text-xl", p.Class))}
	if p.ID != "" {
		args = append(args, html.AId(p.ID))
	}

	for _, a := range p.Attrs {
		args = append(args, a)
	}

	// Add default separator if no children
	if len(*children) == 0 {
		*children = append(*children, html.Span(html.Text("-")))
	}

	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

func InputOTP(args ...html.DivArg) html.Node {
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

	return html.Div(append([]html.DivArg{props}, rest...)...).WithAssets("", inputOTPJS, "ui-inputotp")
}

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

	return html.Div(append([]html.DivArg{props}, rest...)...)
}

func Slot(args ...html.DivArg) html.Node {
	var (
		props SlotProps
		rest  []html.DivArg
	)

	for _, a := range args {
		if v, ok := a.(SlotProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Div(append([]html.DivArg{props}, rest...)...)
}

func Separator(args ...html.DivArg) html.Node {
	var (
		props SeparatorProps
		rest  []html.DivArg
	)

	for _, a := range args {
		if v, ok := a.(SeparatorProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	// Add default separator if no other content
	if len(rest) == 0 {
		rest = append(rest, html.Span(html.Text("-")))
	}

	return html.Div(append([]html.DivArg{props}, rest...)...)
}

//go:embed inputotp.js
var inputOTPJS string
