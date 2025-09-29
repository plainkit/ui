package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/skeleton"
)

func RenderSkeletonsContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Skeletons")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Display lightweight placeholders while content loads.")),
			),
			html.Div(
				html.AClass("grid gap-6 md:grid-cols-2"),
				html.Div(
					html.AClass("space-y-4 rounded-lg border p-6"),
					skeleton.Skeleton(skeleton.Props{Class: "h-8 w-1/2"}),
					html.Div(
						html.AClass("space-y-2"),
						skeleton.Skeleton(skeleton.Props{Class: "h-4 w-full"}),
						skeleton.Skeleton(skeleton.Props{Class: "h-4 w-2/3"}),
						skeleton.Skeleton(skeleton.Props{Class: "h-4 w-3/4"}),
					),
					html.Div(
						html.AClass("flex gap-3"),
						skeleton.Skeleton(skeleton.Props{Class: "h-9 w-24"}),
						skeleton.Skeleton(skeleton.Props{Class: "h-9 w-32"}),
					),
				),
				html.Div(
					html.AClass("space-y-4 rounded-lg border p-6"),
					skeleton.Skeleton(skeleton.Props{Class: "h-48 w-full rounded-md"}),
					html.Div(
						html.AClass("space-y-2"),
						skeleton.Skeleton(skeleton.Props{Class: "h-4 w-3/5"}),
						skeleton.Skeleton(skeleton.Props{Class: "h-4 w-4/5"}),
					),
				),
			),
		),
	)
}
