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

func Skeleton(props Props, args ...html.DivArg) html.Node {
	divArgs := []html.DivArg{html.AClass(classnames.Merge("animate-pulse rounded bg-muted", props.Class))}
	if props.ID != "" {
		divArgs = append(divArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}

	divArgs = append(divArgs, args...)

	return html.Div(divArgs...)
}
