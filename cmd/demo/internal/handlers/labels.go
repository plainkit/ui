package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/label"
)

func RenderLabelsContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Input labelling")),
				html.P(html.AClass("text-sm text-slate-400"), html.Text("Link labels to form controls and surface validation errors.")),
			),
			html.Div(
				html.AClass("grid gap-6 md:grid-cols-2"),
				html.Div(
					html.AClass("space-y-2"),
					label.Label(label.Props{For: "email"}, html.Text("Email")),
					html.Input(
						html.AId("email"),
						html.AClass("w-full rounded-md border border-slate-800 bg-slate-900 px-3 py-2 text-sm text-slate-100"),
						html.APlaceholder("name@example.com"),
					),
				),
				html.Div(
					html.AClass("space-y-2"),
					label.Label(label.Props{For: "password", Error: "Password must be at least 8 characters"}, html.Text("Password")),
					html.Input(
						html.AId("password"),
						html.AClass("w-full rounded-md border border-destructive bg-slate-900 px-3 py-2 text-sm text-slate-100"),
						html.APlaceholder("Enter password"),
						html.AType("password"),
					),
					html.P(html.AClass("text-xs text-destructive"), html.Text("Password must be at least 8 characters.")),
				),
			),
		),
	)
}
