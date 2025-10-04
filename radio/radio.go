package radio

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/internal/styles"
)

type Props struct {
	ID       string
	Class    string
	Attrs    []html.Global
	Name     string
	Value    string
	Form     string
	Disabled bool
	Required bool
	Checked  bool
}

func inputArgsFromProps(baseClass string, extra ...string) func(p Props) []html.InputArg {
	return func(p Props) []html.InputArg {
		className := html.ClassMerge(
			append([]string{baseClass},
				append(extra, p.Class)...)...)

		args := []html.InputArg{
			html.AType("radio"),
			html.AClass(className),
		}

		if p.ID != "" {
			args = append(args, html.AId(p.ID))
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
		styles.Control(
			"size-4 rounded-full border-2",
			"before:absolute before:left-1/2 before:top-1/2",
			"before:h-2 before:w-2 before:-translate-x-1/2 before:-translate-y-1/2",
			"before:rounded-full before:bg-primary/80 before:opacity-0",
			"checked:border-primary checked:bg-primary/10 checked:before:opacity-100",
			"transition-all duration-200 before:transition-opacity",
		),
	)(p)

	if p.Name != "" {
		args = append(args, html.AName(p.Name))
	}

	if p.Value != "" {
		args = append(args, html.AValue(p.Value))
	}

	if p.Form != "" {
		args = append(args, html.AForm(p.Form))
	}

	if p.Checked {
		args = append(args, html.AChecked())
	}

	if p.Disabled {
		args = append(args, html.ADisabled())
	}

	if p.Required {
		args = append(args, html.ARequired())
	}

	for _, a := range args {
		a.ApplyInput(attrs, children)
	}
}

// Radio renders a styled radio input using the composable pattern.
// Accepts variadic html.InputArg arguments, with Props as an optional first argument.
func Radio(args ...html.InputArg) html.Node {
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

	return html.Input(append([]html.InputArg{props}, rest...)...)
}
