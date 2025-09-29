package skeleton

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/internal/classnames"
)

type Props struct {
	ID    string
	Class string
	Attrs []html.Global
}

func (p Props) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	args := []html.DivArg{html.AClass(classnames.Merge("animate-pulse rounded bg-muted", p.Class))}
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

func Skeleton(args ...html.DivArg) html.Node {
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

	return html.Div(append([]html.DivArg{props}, rest...)...)
}
