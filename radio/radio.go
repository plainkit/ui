package radio

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/internal/classnames"
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

func Radio(props Props) html.Node {
	inputArgs := []html.InputArg{
		html.AType("radio"),
		html.AClass(classnames.Merge(
			"relative h-4 w-4",
			"appearance-none rounded-full",
			"border-2 border-primary",
			"before:absolute before:left-1/2 before:top-1/2",
			"before:h-1.5 before:w-1.5 before:-translate-x-1/2 before:-translate-y-1/2",
			"before:content[''] before:rounded-full before:bg-background",
			"checked:border-primary checked:bg-primary checked:before:visible",
			"focus-visible:outline-hidden focus-visible:ring-2 focus-visible:ring-ring",
			"focus-visible:ring-offset-2 focus-visible:ring-offset-background",
			"disabled:cursor-not-allowed",
			props.Class,
		)),
	}

	if props.ID != "" {
		inputArgs = append(inputArgs, html.AId(props.ID))
	}
	if props.Name != "" {
		inputArgs = append(inputArgs, html.AName(props.Name))
	}
	if props.Value != "" {
		inputArgs = append(inputArgs, html.AValue(props.Value))
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

	return html.Input(inputArgs...)
}
