package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/progress"
)

func RenderProgressContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Upload progress")),
				html.P(html.AClass("text-sm text-slate-400"), html.Text("Display progress with optional labels and value output.")),
			),
			html.Div(
				html.AClass("space-y-4"),
				progress.Progress(progress.Props{Label: "Preparing files", Value: 30, ShowValue: true}),
				progress.Progress(progress.Props{Label: "Uploading", Value: 65, ShowValue: true, Variant: progress.VariantSuccess}),
				progress.Progress(progress.Props{Label: "Processing", Value: 45, ShowValue: true, Variant: progress.VariantWarning, Size: progress.SizeLg}),
			),
		),
	)
}
