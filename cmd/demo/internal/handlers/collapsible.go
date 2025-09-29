package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/collapsible"
)

func RenderCollapsibleContent() html.Node {
	triggerClasses := "flex items-center justify-between gap-3 rounded-md border border-border bg-card px-4 py-3 text-sm font-medium text-card-foreground cursor-pointer"
	contentClasses := "rounded-b-md border border-border border-t-0 bg-card px-4 py-3 text-sm text-muted-foreground"

	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Collapsible")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Progressively reveal supplemental details without leaving the current view.")),
			),
			html.Div(
				html.AClass("space-y-4"),
				collapsible.Collapsible(
					collapsible.Props{Class: "space-y-2"},
					collapsible.Trigger(
						collapsible.TriggerProps{Class: triggerClasses},
						html.Span(html.Text("What is Plain UI?")),
						lucide.ChevronDown(html.AClass("size-4 text-muted-foreground")),
					),
					collapsible.Content(
						html.Div(
							html.AClass(contentClasses),
							html.P(html.Text("Plain UI is a growing collection of composable components built on the plainkit html helpers.")),
						),
					),
				),
				collapsible.Collapsible(
					collapsible.Props{Class: "space-y-2", Open: true},
					collapsible.Trigger(
						collapsible.TriggerProps{Class: triggerClasses},
						html.Span(html.Text("Can I style it with Tailwind?")),
						lucide.ChevronDown(html.AClass("size-4 text-muted-foreground")),
					),
					collapsible.Content(
						html.Div(
							html.AClass(contentClasses),
							html.P(html.Text("Yes! Utility classes are merged automatically so you can override tokens per instance.")),
						),
					),
				),
			),
		),
	)
}
