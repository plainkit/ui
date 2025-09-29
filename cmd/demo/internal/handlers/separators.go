package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/separator"
)

func RenderSeparatorsContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Section dividers")),
				html.P(html.AClass("text-sm text-slate-400"), html.Text("Separate groups of content with optional labels.")),
			),
			html.Div(
				html.AClass("space-y-6"),
				separator.Separator(html.T("Team")),
				html.Div(
					html.AClass("flex h-32 items-center justify-center gap-6"),
					html.Text("Start"),
					separator.Separator(separator.Props{Orientation: separator.OrientationVertical, Decoration: separator.DecorationDotted}),
					html.Text("End"),
				),
			),
		),
	)
}
