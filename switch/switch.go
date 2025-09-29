package switchcomp

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/plainkit/html"
	"github.com/plainkit/ui/internal/classnames"
)

type Props struct {
	ID       string
	Class    string
	Attrs    []html.Global
	Name     string
	Value    string
	Disabled bool
	Checked  bool
	Form     string
}

func labelArgsFromProps(baseClass string, extra ...string) func(p Props) []html.LabelArg {
	return func(p Props) []html.LabelArg {
		id := p.ID
		if id == "" {
			id = randomID()
		}

		className := classnames.Merge(
			append([]string{baseClass},
				append(extra, conditional(p.Disabled, "cursor-not-allowed"))...)...)

		args := []html.LabelArg{
			html.AFor(id),
			html.AClass(className),
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

// ApplyLabel implements the html.LabelArg interface for Props
func (p Props) ApplyLabel(attrs *html.LabelAttrs, children *[]html.Component) {
	id := p.ID
	if id == "" {
		id = randomID()
	}

	args := labelArgsFromProps("inline-flex cursor-pointer items-center gap-2")(p)

	inputArgs := []html.InputArg{
		html.AId(id),
		html.AType("checkbox"),
		html.AClass("peer hidden"),
		html.ACustom("role", "switch"),
	}
	if p.Name != "" {
		inputArgs = append(inputArgs, html.AName(p.Name))
	}

	if p.Value != "" {
		inputArgs = append(inputArgs, html.AValue(p.Value))
	} else {
		inputArgs = append(inputArgs, html.AValue("on"))
	}

	if p.Form != "" {
		inputArgs = append(inputArgs, html.AForm(p.Form))
	}

	if p.Checked {
		inputArgs = append(inputArgs, html.AChecked())
	}

	if p.Disabled {
		inputArgs = append(inputArgs, html.ADisabled())
	}

	visual := html.Div(
		html.AClass(classnames.Merge(
			"relative inline-flex h-5 w-9 shrink-0 cursor-pointer items-center",
			"rounded-full border-2 border-transparent",
			"transition-colors",
			"bg-input",
			"peer-checked:bg-primary",
			"peer-focus-visible:outline-none peer-focus-visible:ring-2",
			"peer-focus-visible:ring-ring peer-focus-visible:ring-offset-2",
			"peer-focus-visible:ring-offset-background",
			"peer-disabled:cursor-not-allowed peer-disabled:opacity-50",
			"after:pointer-events-none after:block",
			"after:h-4 after:w-4",
			"after:rounded-full after:bg-background",
			"after:shadow-lg after:ring-0",
			"after:transition-transform",
			"after:content-['']",
			"peer-checked:after:translate-x-4",
			p.Class,
		)),
		html.AAria("hidden", "true"),
	)

	// Apply base label args first
	for _, a := range args {
		a.ApplyLabel(attrs, children)
	}

	// Add the input and visual elements as children
	*children = append(*children, html.Input(inputArgs...), visual)
}

// Switch renders an accessible checkbox-based toggle control using the composable pattern.
// Accepts variadic html.LabelArg arguments, with Props as an optional first argument.
func Switch(args ...html.LabelArg) html.Node {
	var (
		props Props
		rest  []html.LabelArg
	)

	// Separate Props from other arguments
	for _, a := range args {
		if v, ok := a.(Props); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Label(append([]html.LabelArg{props}, rest...)...)
}

func randomID() string {
	bytes := make([]byte, 6)
	if _, err := rand.Read(bytes); err != nil {
		return "switch-id"
	}

	return "switch-" + hex.EncodeToString(bytes)
}

func conditional(cond bool, class string) string {
	if cond {
		return class
	}

	return ""
}
