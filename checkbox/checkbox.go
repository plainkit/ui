package checkbox

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
	Disabled bool
	Required bool
	Checked  bool
	Form     string
}

func inputArgsFromProps(baseClass string, extra ...string) func(p Props) []html.InputArg {
	return func(p Props) []html.InputArg {
		className := html.ClassMerge(
			append([]string{baseClass},
				append(extra, p.Class)...)...)

		args := []html.InputArg{
			html.AType("checkbox"),
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
			"rounded-lg shadow-sm",
			"checked:bg-gradient-to-br checked:from-primary checked:via-primary/90 checked:to-primary/80",
			"checked:border-transparent",
			"appearance-none",
		),
		"checked:text-primary-foreground",
		"transition-shadow duration-200",
	)(p)

	if p.Name != "" {
		args = append(args, html.AName(p.Name))
	}

	if p.Value != "" {
		args = append(args, html.AValue(p.Value))
	} else {
		args = append(args, html.AValue("on"))
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

// Checkbox renders a styled checkbox input using the composable pattern.
// Accepts variadic html.InputArg arguments, with Props as an optional first argument.
func Checkbox(args ...html.InputArg) html.Node {
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

	// Create the wrapper div
	divArgs := []html.DivArg{html.AClass("relative inline-flex items-center")}

	// Add the input with all arguments applied
	divArgs = append(divArgs, html.Input(append([]html.InputArg{props}, rest...)...))

	divArgs = append(divArgs, html.Div(
		html.AClass("absolute left-0 top-0 flex h-4 w-4 items-center justify-center text-primary-foreground opacity-0 transition-opacity duration-150 peer-checked:opacity-100"),
	))

	return html.Div(divArgs...)
}
