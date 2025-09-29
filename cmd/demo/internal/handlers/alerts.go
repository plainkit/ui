package handlers

import (
	. "github.com/plainkit/html"
	"github.com/plainkit/ui/alert"
)

func RenderAlertsContent() Node {
	return Div(
		AClass("space-y-10"),
		Section(
			AClass("space-y-4"),
			Div(
				AClass("space-y-1"),
				H2(AClass("text-2xl font-semibold"), T("Contextual alerts")),
				P(AClass("text-sm text-slate-400"), T("Communicate high priority feedback and destructive states.")),
			),
			Div(
				AClass("space-y-3"),
				alert.Alert(
					alert.Title(T("Heads up!")),
					alert.Description(P(T("You can add components to your project from the left sidebar."))),
				),
				alert.Alert(
					alert.Props{Variant: alert.VariantDestructive},
					alert.Title(T("Action required")),
					alert.Description(P(T("This change is permanent and cannot be undone."))),
				),
			),
		),
	)
}
