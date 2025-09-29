package handlers

import (
	"github.com/plainkit/html"
	switchcomp "github.com/plainkit/ui/switch"
)

func RenderSwitchesContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Switches")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Switches express an on/off state for instant changes.")),
			),
			html.Div(
				html.AClass("space-y-3"),
				switchcomp.Switch(switchcomp.Props{ID: "switch-email", Name: "email"}, html.Child(html.Span(html.Text("Email notifications")))),
				switchcomp.Switch(switchcomp.Props{ID: "switch-push", Name: "push", Checked: true}, html.Child(html.Span(html.Text("Push notifications")))),
				switchcomp.Switch(switchcomp.Props{ID: "switch-disabled", Disabled: true}, html.Child(html.Span(html.Text("Public profile")))),
			),
		),
	)
}
