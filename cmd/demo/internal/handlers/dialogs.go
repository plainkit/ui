package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/button"
	"github.com/plainkit/ui/dialog"
)

func RenderDialogsContent() html.Node {
	const dialogID = "demo-dialog"

	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Dialogs")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Modal dialogs for confirmations, forms, and important communications.")),
			),
			html.Div(
				html.AClass("space-y-6"),
				html.Div(
					html.AClass("flex flex-wrap gap-3"),
					dialog.Trigger(
						dialog.TriggerProps{For: dialogID},
						button.Props{},
						html.Text("Open Dialog"),
					),
				),
				dialog.Content(
					dialog.ContentProps{ID: dialogID},
					dialog.Header(
						dialog.HeaderProps{},
						dialog.Title(dialog.TitleProps{}, html.Text("Confirm Action")),
						dialog.Description(dialog.DescriptionProps{}, html.Text("Are you sure you want to delete this item? This action cannot be undone.")),
					),
					dialog.Footer(
						dialog.FooterProps{},
						dialog.Close(dialog.CloseProps{For: dialogID},
							button.Button(button.Props{Variant: button.VariantOutline}, html.Text("Cancel"))),
						button.Button(button.Props{Variant: button.VariantDestructive}, html.Text("Delete")),
					),
				),
			),
		),
	)
}
