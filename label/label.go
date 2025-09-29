package label

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/internal/classnames"
)

type Props struct {
	ID    string
	Class string
	Attrs []html.Global
	For   string
	Error string
}

func Label(props Props, args ...html.LabelArg) html.Node {
	className := classnames.Merge(
		"text-sm font-medium leading-none inline-block",
		conditionalClass(props.Error != "", "text-destructive"),
		props.Class,
	)

	labelArgs := []html.LabelArg{html.AClass(className), html.AData("pui-label-disabled-style", "opacity-50 cursor-not-allowed")}
	if props.ID != "" {
		labelArgs = append(labelArgs, html.AId(props.ID))
	}

	if props.For != "" {
		labelArgs = append(labelArgs, html.AFor(props.For))
	}

	for _, attr := range props.Attrs {
		labelArgs = append(labelArgs, attr)
	}

	labelArgs = append(labelArgs, args...)

	return html.Label(labelArgs...)
}

func conditionalClass(cond bool, class string) string {
	if cond {
		return class
	}

	return ""
}
