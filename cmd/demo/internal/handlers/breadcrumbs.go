package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/breadcrumb"
)

func RenderBreadcrumbsContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Breadcrumbs")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Provide location context and quick navigation back to parent pages.")),
			),
			breadcrumb.Breadcrumb(
				breadcrumb.Props{Class: "rounded-lg border bg-card px-4 py-3"},
				breadcrumb.List(
					breadcrumb.ListProps{},
					breadcrumb.Item(
						breadcrumb.ItemProps{},
						breadcrumb.Link(breadcrumb.LinkProps{Href: "#"}, html.Text("Dashboard")),
						breadcrumb.Separator(breadcrumb.SeparatorProps{}),
					),
					breadcrumb.Item(
						breadcrumb.ItemProps{},
						breadcrumb.Link(breadcrumb.LinkProps{Href: "#"}, html.Text("Projects")),
						breadcrumb.Separator(breadcrumb.SeparatorProps{}),
					),
					breadcrumb.Item(
						breadcrumb.ItemProps{},
						breadcrumb.Page(breadcrumb.ItemProps{}, html.Text("Plain UI")),
					),
				),
			),
		),
	)
}
