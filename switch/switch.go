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

// Switch renders an accessible checkbox-based toggle control.
func Switch(props Props, labelArgs ...html.LabelArg) html.Node {
	if props.ID == "" {
		props.ID = randomID()
	}

	labelClasses := classnames.Merge("inline-flex cursor-pointer items-center gap-2", conditional(props.Disabled, "cursor-not-allowed"))

	labelArgs = append([]html.LabelArg{html.AFor(props.ID), html.AClass(labelClasses)}, labelArgs...)

	inputArgs := []html.InputArg{
		html.AId(props.ID),
		html.AType("checkbox"),
		html.AClass("peer hidden"),
		html.ACustom("role", "switch"),
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

	for _, attr := range props.Attrs {
		inputArgs = append(inputArgs, attr)
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
			props.Class,
		)),
		html.AAria("hidden", "true"),
	)

	labelArgs = append(labelArgs,
		html.Input(inputArgs...),
		visual,
	)

	return html.Label(labelArgs...)
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
