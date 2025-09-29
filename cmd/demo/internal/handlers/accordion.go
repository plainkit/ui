package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/accordion"
)

func RenderAccordionContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Accordion")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Stack collapsible sections to keep dense content organized.")),
			),
			accordion.Accordion(
				accordion.Props{},
				accordion.Item(
					accordion.ItemProps{},
					accordion.Trigger(accordion.TriggerProps{}, html.Text("Is Plain UI accessible?")),
					accordion.Content(
						accordion.ContentProps{},
						html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Yes. We build on semantic HTML, manage focus states, and respect reduced-motion preferences.")),
					),
				),
				accordion.Item(
					accordion.ItemProps{},
					accordion.Trigger(accordion.TriggerProps{}, html.Text("Can I use it with templ?")),
					accordion.Content(
						accordion.ContentProps{},
						html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Absolutely. The APIs are designed to feel templ-friendly and compose with html components.")),
					),
				),
				accordion.Item(
					accordion.ItemProps{},
					accordion.Trigger(accordion.TriggerProps{}, html.Text("Does it support nested content?")),
					accordion.Content(
						accordion.ContentProps{},
						html.Div(
							html.AClass("space-y-2 text-sm text-muted-foreground"),
							html.P(html.Text("Yes. You can place paragraphs, lists, and interactive elements inside each panel.")),
							html.Ul(
								html.AClass("list-disc pl-5"),
								html.Li(html.Text("Links and buttons")),
								html.Li(html.Text("Images or media")),
								html.Li(html.Text("Nested layout blocks")),
							),
						),
					),
				),
			),
		),
	)
}
