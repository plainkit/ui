package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/input"
	"github.com/plainkit/ui/label"
)

func RenderInputsContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Form inputs")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Text inputs adapt to validation, states, and password toggles.")),
			),
			html.Div(
				html.AClass("grid gap-6 md:grid-cols-2"),
				html.Div(
					html.AClass("space-y-2"),
					label.Label(label.Props{For: "name"}, html.Text("Full name")),
					input.Input(input.Props{ID: "name", Name: "name", Placeholder: "Ada Lovelace"}),
				),
				html.Div(
					html.AClass("space-y-2"),
					label.Label(label.Props{For: "email"}, html.Text("Email address")),
					input.Input(input.Props{ID: "email", Name: "email", Type: input.TypeEmail, Placeholder: "name@example.com"}),
				),
				html.Div(
					html.AClass("space-y-2"),
					label.Label(label.Props{For: "password"}, html.Text("Password")),
					input.Input(input.Props{ID: "password", Name: "password", Type: input.TypePassword, Placeholder: "••••••••", ShowPasswordToggle: true}),
				),
				html.Div(
					html.AClass("space-y-2"),
					label.Label(label.Props{For: "search", Error: "No results"}, html.Text("Search")),
					input.Input(input.Props{ID: "search", Name: "search", Placeholder: "Search docs", HasError: true}),
				),
			),
		),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("File & disabled inputs")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Provide metadata for file uploads and communicate disabled states.")),
			),
			html.Div(
				html.AClass("grid gap-6 md:grid-cols-2"),
				html.Div(
					html.AClass("space-y-2"),
					label.Label(label.Props{For: "upload"}, html.Text("Project assets")),
					input.Input(input.Props{ID: "upload", Name: "assets", Type: input.TypeFile, FileAccept: "image/*"}),
				),
				html.Div(
					html.AClass("space-y-2"),
					label.Label(label.Props{For: "company"}, html.Text("Company")),
					input.Input(input.Props{ID: "company", Name: "company", Placeholder: "Acme Inc.", Disabled: true}),
				),
			),
		),
	)
}
