package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/button"
	"github.com/plainkit/ui/card"
)

func RenderCardsContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Cards")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Surface key information inside a bordered container.")),
			),
			html.Div(
				html.AClass("grid gap-6 md:grid-cols-2"),
				card.Card(
					card.Props{},
					card.Header(
						card.HeaderProps{},
						html.Div(
							html.AClass("space-y-1"),
							card.Title(card.TitleProps{}, html.Text("Change subscription")),
							card.Description(card.DescriptionProps{}, html.Text("Upgrade or downgrade your current plan.")),
						),
					),
					card.Content(
						card.ContentProps{},
						html.Div(
							html.AClass("space-y-3"),
							html.P(html.AClass("text-sm text-muted-foreground"), html.Text("You're currently on the Team plan. Teams get collaborative features, SSO, and priority support.")),
							html.Ul(
								html.AClass("space-y-2 text-sm"),
								html.Li(html.Text("• Unlimited collaborators")),
								html.Li(html.Text("• Shared components")),
								html.Li(html.Text("• Priority support")),
							),
						),
					),
					card.Footer(
						card.FooterProps{},
						html.Div(
							html.AClass("ml-auto flex gap-3"),
							button.Button(button.Props{Variant: button.VariantOutline}, html.Text("Manage plan")),
							button.Button(html.Text("Upgrade")),
						),
					),
				),
			),
		),
	)
}
