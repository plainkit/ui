package checkbox

import (
	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/internal/classnames"
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
		className := classnames.Merge(
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
		"peer size-4 shrink-0 rounded-[4px] border border-input shadow-xs",
		"focus-visible:outline-none focus-visible:ring-[3px] focus-visible:ring-ring/50 focus-visible:border-ring",
		"disabled:cursor-not-allowed disabled:opacity-50",
		"checked:bg-primary checked:text-primary-foreground checked:border-primary",
		"appearance-none cursor-pointer transition-shadow",
		"relative",
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
		icon  html.Component
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

	// Add the icon overlay
	if icon == nil {
		icon = lucide.Check(html.AClass("size-3.5"))
	}

	divArgs = append(divArgs, html.Div(
		html.AClass("absolute left-0 top-0 h-4 w-4 pointer-events-none flex items-center justify-center text-primary-foreground opacity-0 peer-checked:opacity-100"),
		html.Child(icon),
	))

	return html.Div(divArgs...)
}
