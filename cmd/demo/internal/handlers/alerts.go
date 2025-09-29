package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/alert"
)

func RenderAlertsContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Contextual alerts")),
				html.P(html.AClass("text-sm text-slate-400"), html.Text("Communicate high priority feedback and destructive states.")),
			),
			html.Div(
				html.AClass("space-y-3"),
				alert.Alert(
					alert.Props{},
					alert.Title(alert.TitleProps{}, html.Text("Heads up!")),
					alert.Description(alert.DescriptionProps{}, html.P(html.Text("You can add components to your project from the left sidebar."))),
				),
				alert.Alert(
					alert.Props{Variant: alert.VariantDestructive},
					alert.Title(alert.TitleProps{}, html.Text("Action required")),
					alert.Description(alert.DescriptionProps{}, html.P(html.Text("This change is permanent and cannot be undone."))),
				),
			),
		),
	)
}
