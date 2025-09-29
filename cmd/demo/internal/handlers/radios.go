package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/radio"
)

func RenderRadiosContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Radios")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Allow users to make a single choice from a concise list.")),
			),
			html.Fieldset(
				html.AClass("space-y-4 rounded-lg border p-6"),
				html.Legend(html.AClass("text-sm font-semibold uppercase tracking-wide text-muted-foreground"), html.Text("Delivery speed")),
				html.Div(
					html.AClass("space-y-3"),
					html.Label(
						html.AClass("flex items-center gap-3 text-sm"),
						radio.Radio(radio.Props{ID: "delivery-standard", Name: "delivery", Value: "standard", Checked: true}),
						html.Div(
							html.AClass("flex flex-col"),
							html.Span(html.AClass("font-medium"), html.Text("Standard")),
							html.Span(html.AClass("text-xs text-muted-foreground"), html.Text("3-5 business days")),
						),
					),
					html.Label(
						html.AClass("flex items-center gap-3 text-sm"),
						radio.Radio(radio.Props{ID: "delivery-express", Name: "delivery", Value: "express"}),
						html.Div(
							html.AClass("flex flex-col"),
							html.Span(html.AClass("font-medium"), html.Text("Express")),
							html.Span(html.AClass("text-xs text-muted-foreground"), html.Text("Arrives tomorrow")),
						),
					),
					html.Label(
						html.AClass("flex items-center gap-3 text-sm"),
						radio.Radio(radio.Props{ID: "delivery-same-day", Name: "delivery", Value: "same-day"}),
						html.Div(
							html.AClass("flex flex-col"),
							html.Span(html.AClass("font-medium"), html.Text("Same day")),
							html.Span(html.AClass("text-xs text-muted-foreground"), html.Text("Available in select cities")),
						),
					),
				),
			),
		),
	)
}
