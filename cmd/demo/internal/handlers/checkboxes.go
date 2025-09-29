package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/checkbox"
)

func RenderCheckboxesContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Checkboxes")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Use checkboxes to represent binary options and multi-select choices.")),
			),
			html.Div(
				html.AClass("space-y-3"),
				html.Label(
					html.AClass("flex items-center gap-2"),
					html.Child(checkbox.Checkbox(checkbox.Props{ID: "cb-updates", Name: "updates"})),
					html.Text("Email me product updates"),
				),
				html.Label(
					html.AClass("flex items-center gap-2"),
					html.Child(checkbox.Checkbox(checkbox.Props{ID: "cb-terms", Name: "terms", Required: true})),
					html.Text("I agree to the terms"),
				),
				html.Label(
					html.AClass("flex items-center gap-2 text-muted-foreground"),
					html.Child(checkbox.Checkbox(checkbox.Props{ID: "cb-disabled", Disabled: true})),
					html.Text("Remember my choice (unavailable)"),
				),
			),
		),
	)
}
