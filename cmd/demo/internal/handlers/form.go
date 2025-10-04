package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/button"
	"github.com/plainkit/ui/form"
	"github.com/plainkit/ui/input"
	switchcomp "github.com/plainkit/ui/switch"
)

func RenderFormContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Form Building Blocks")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Compose labels, descriptions, and validation messages around existing inputs.")),
			),
			html.Div(
				html.AClass("grid gap-6 lg:grid-cols-2"),
				html.Form(
					html.AClass("space-y-6 rounded-lg border bg-card p-6 shadow-xs"),
					form.Item(
						form.ItemProps{},
						form.Label(form.LabelProps{For: "full-name"}, html.Text("Full name")),
						input.Input(input.Props{
							ID:          "full-name",
							Name:        "full_name",
							Placeholder: "Ada Lovelace",
							Required:    true,
						}),
						form.Description(form.DescriptionProps{}, html.Text("We use your full name to personalize communications.")),
					),
					form.Item(
						form.ItemProps{},
						form.Label(form.LabelProps{For: "email-address"}, html.Text("Email address")),
						input.Input(input.Props{
							ID:          "email-address",
							Name:        "email",
							Type:        input.TypeEmail,
							Placeholder: "ada@example.com",
							Required:    true,
						}, html.AAutocomplete("email")),
						form.Description(form.DescriptionProps{}, html.Text("We'll only send important updates.")),
					),
					form.Item(
						form.ItemProps{},
						form.Label(form.LabelProps{For: "account-password"}, html.Text("Password")),
						input.Input(input.Props{
							ID:                 "account-password",
							Name:               "password",
							Type:               input.TypePassword,
							Placeholder:        "Create a strong password",
							Required:           true,
							ShowPasswordToggle: true,
							HasError:           true,
						}),
						form.Message(form.MessageProps{Variant: form.MessageVariantError}, html.Text("Use at least 8 characters with a number and symbol.")),
					),
					html.Div(
						html.AClass("flex items-center justify-end gap-3"),
						button.Button(button.Props{Variant: button.VariantOutline, Type: button.TypeButton}, html.Text("Cancel")),
						button.Button(button.Props{Type: button.TypeSubmit}, html.Text("Save changes")),
					),
				),
				html.Div(
					html.AClass("space-y-4 rounded-lg border bg-card p-6 shadow-xs"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Inline controls")),
					html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Use ItemFlex to align toggles and short helpers on a single row.")),
					form.ItemFlex(
						form.ItemProps{},
						switchcomp.Switch(switchcomp.Props{
							ID:      "alerts-switch",
							Name:    "alerts",
							Checked: true,
						}, html.Text("Email notifications")),
						html.Div(
							html.AClass("text-sm text-muted-foreground"),
							html.Text("Send me important product and billing updates."),
						),
					),
					form.ItemFlex(
						form.ItemProps{},
						html.Child(switchcomp.Switch(switchcomp.Props{
							ID:   "beta-switch",
							Name: "beta_program",
						}, html.Text("Join beta program"))),
						html.Div(
							html.AClass("text-sm text-muted-foreground"),
							html.Text("Access experimental features before they ship."),
						),
					),
				),
			),
		),
	)
}
