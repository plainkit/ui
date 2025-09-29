package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/aspectratio"
)

func RenderAspectRatiosContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Aspect Ratios")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Maintain consistent media sizing across breakpoints.")),
			),
			html.Div(
				html.AClass("grid gap-6 md:grid-cols-2"),
				aspectratio.AspectRatio(
					aspectratio.Props{Ratio: aspectratio.RatioVideo, Class: "overflow-hidden rounded-lg border"},
					html.Img(
						html.ASrc("https://images.unsplash.com/photo-1522199755839-a2bacb67c546?auto=format&fit=crop&w=800&q=80"),
						html.AAlt("Team working together"),
						html.AClass("h-full w-full object-cover"),
					),
				),
				aspectratio.AspectRatio(
					aspectratio.Props{Ratio: aspectratio.RatioSquare, Class: "overflow-hidden rounded-lg border bg-muted"},
					html.Div(
						html.AClass("flex h-full w-full items-center justify-center text-center"),
						html.Div(
							html.AClass("space-y-2"),
							html.H3(html.AClass("text-lg font-semibold"), html.Text("Square preview")),
							html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Use any content inside the ratio wrapper.")),
						),
					),
				),
			),
		),
	)
}
