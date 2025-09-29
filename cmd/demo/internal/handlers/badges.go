package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/badge"
)

func RenderBadgesContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Variants")),
				html.P(html.AClass("text-sm text-slate-400"), html.Text("Status indicators for different contexts.")),
			),
			html.Div(
				html.AClass("flex flex-wrap gap-3"),
				badge.Badge(html.T("Default")),
				badge.Badge(badge.Props{Variant: badge.VariantSecondary}, html.T("Secondary")),
				badge.Badge(badge.Props{Variant: badge.VariantOutline}, html.T("Outline")),
				badge.Badge(badge.Props{Variant: badge.VariantDestructive}, html.T("Destructive")),
			),
		),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Custom content")),
				html.P(html.AClass("text-sm text-slate-400"), html.Text("Badges can contain arbitrary content.")),
			),
			html.Div(
				html.AClass("flex flex-wrap gap-3"),
				badge.Badge(badge.Props{Class: "gap-1.5"}, html.T("v2.0")),
				badge.Badge(
					badge.Props{Variant: badge.VariantSecondary},
					html.Span(
						html.AClass("inline-flex items-center gap-1"),
						html.Text("Active"),
					),
				),
			),
		),
	)
}
