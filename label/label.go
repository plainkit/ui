package label

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/internal/styles"
)

type Props struct {
	ID    string
	Class string
	Attrs []html.Global
	For   string
	Error string
}

func (p Props) ApplyLabel(attrs *html.LabelAttrs, children *[]html.Component) {
	className := styles.Label(
		"inline-block leading-tight",
		conditionalClass(p.Error != "", "text-destructive"),
		p.Class,
	)

	args := []html.LabelArg{html.AClass(className), html.AData("pui-label-disabled-style", "opacity-50 cursor-not-allowed")}
	if p.ID != "" {
		args = append(args, html.AId(p.ID))
	}

	if p.For != "" {
		args = append(args, html.AFor(p.For))
	}

	for _, a := range p.Attrs {
		args = append(args, a)
	}

	for _, a := range args {
		a.ApplyLabel(attrs, children)
	}
}

func Label(args ...html.LabelArg) html.Node {
	var (
		props Props
		rest  []html.LabelArg
	)

	for _, a := range args {
		if v, ok := a.(Props); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Label(append([]html.LabelArg{props}, rest...)...)
}

func conditionalClass(cond bool, class string) string {
	if cond {
		return class
	}

	return ""
}
