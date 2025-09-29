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

// Checkbox renders a styled checkbox input with an optional custom icon overlay.
func Checkbox(props Props, icon html.Component) html.Node {
	divArgs := []html.DivArg{html.AClass("relative inline-flex items-center")}

	inputArgs := []html.InputArg{
		html.AType("checkbox"),
		html.AClass(checkboxClass(props)),
	}
	if props.ID != "" {
		inputArgs = append(inputArgs, html.AId(props.ID))
	}
	if props.Name != "" {
		inputArgs = append(inputArgs, html.AName(props.Name))
	}
	if props.Value != "" {
		inputArgs = append(inputArgs, html.AValue(props.Value))
	} else {
		inputArgs = append(inputArgs, html.AValue("on"))
	}
	if props.Form != "" {
		inputArgs = append(inputArgs, html.AForm(props.Form))
	}
	if props.Checked {
		inputArgs = append(inputArgs, html.AChecked())
	}
	if props.Disabled {
		inputArgs = append(inputArgs, html.ADisabled())
	}
	if props.Required {
		inputArgs = append(inputArgs, html.ARequired())
	}
	for _, attr := range props.Attrs {
		inputArgs = append(inputArgs, attr)
	}

	divArgs = append(divArgs, html.Input(inputArgs...))

	iconNode := icon
	if iconNode == nil {
		iconNode = lucide.Check(html.AClass("size-3.5"))
	}

	divArgs = append(divArgs, html.Div(
		html.AClass("absolute left-0 top-0 h-4 w-4 pointer-events-none flex items-center justify-center text-primary-foreground opacity-0 peer-checked:opacity-100"),
		html.Child(iconNode),
	))

	return html.Div(divArgs...)
}

func checkboxClass(props Props) string {
	return classnames.Merge(
		"peer size-4 shrink-0 rounded-[4px] border border-input shadow-xs",
		"focus-visible:outline-none focus-visible:ring-[3px] focus-visible:ring-ring/50 focus-visible:border-ring",
		"disabled:cursor-not-allowed disabled:opacity-50",
		"checked:bg-primary checked:text-primary-foreground checked:border-primary",
		"appearance-none cursor-pointer transition-shadow",
		"relative",
		props.Class,
	)
}
