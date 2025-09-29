package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/label"
	"github.com/plainkit/ui/textarea"
)

func RenderTextareasContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Textareas")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Collect multi-line input with optional auto-resize.")),
			),
			html.Div(
				html.AClass("grid gap-6 md:grid-cols-2"),
				html.Div(
					html.AClass("space-y-2"),
					label.Label(label.Props{For: "bio"}, html.Text("Short bio")),
					textarea.Textarea(textarea.Props{ID: "bio", Name: "bio", Placeholder: "Tell us about yourself", Rows: 4}),
				),
				html.Div(
					html.AClass("space-y-2"),
					label.Label(label.Props{For: "notes"}, html.Text("Meeting notes")),
					textarea.Textarea(textarea.Props{ID: "notes", Name: "notes", AutoResize: true, Placeholder: "Notes will auto-resize as you type."}),
				),
			),
		),
	)
}
