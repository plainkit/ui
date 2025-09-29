package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/selectbox"
)

func RenderSelectBoxesContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Select Boxes")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Dropdown selectors with search and multi-selection capabilities.")),
			),
			html.Div(
				html.AClass("grid gap-6 md:grid-cols-2"),
				html.Div(
					html.AClass("space-y-4"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Single Select")),
					selectbox.SelectBox(
						selectbox.Props{},
						selectbox.Trigger(
							selectbox.TriggerProps{Name: "framework"},
							"single-select-content",
							selectbox.Value(selectbox.ValueProps{Placeholder: "Select framework..."}),
						),
						selectbox.Content(
							selectbox.ContentProps{ID: "single-select-content"},
							selectbox.Group(
								selectbox.GroupProps{},
								selectbox.Label(selectbox.LabelProps{}, html.Text("Frontend")),
								selectbox.Item(selectbox.ItemProps{Value: "react"}, html.Text("React")),
								selectbox.Item(selectbox.ItemProps{Value: "vue"}, html.Text("Vue.js")),
								selectbox.Item(selectbox.ItemProps{Value: "angular"}, html.Text("Angular")),
							),
							selectbox.Group(
								selectbox.GroupProps{},
								selectbox.Label(selectbox.LabelProps{}, html.Text("Backend")),
								selectbox.Item(selectbox.ItemProps{Value: "go"}, html.Text("Go")),
								selectbox.Item(selectbox.ItemProps{Value: "node"}, html.Text("Node.js")),
								selectbox.Item(selectbox.ItemProps{Value: "python"}, html.Text("Python")),
							),
						),
					),
				),
				html.Div(
					html.AClass("space-y-4"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("With Search")),
					selectbox.SelectBox(
						selectbox.Props{},
						selectbox.Trigger(
							selectbox.TriggerProps{Name: "country"},
							"searchable-select-content",
							selectbox.Value(selectbox.ValueProps{Placeholder: "Select country..."}),
						),
						selectbox.Content(
							selectbox.ContentProps{
								ID:                "searchable-select-content",
								SearchPlaceholder: "Search countries...",
							},
							selectbox.Item(selectbox.ItemProps{Value: "us"}, html.Text("United States")),
							selectbox.Item(selectbox.ItemProps{Value: "uk"}, html.Text("United Kingdom")),
							selectbox.Item(selectbox.ItemProps{Value: "de"}, html.Text("Germany")),
							selectbox.Item(selectbox.ItemProps{Value: "fr"}, html.Text("France")),
							selectbox.Item(selectbox.ItemProps{Value: "jp"}, html.Text("Japan")),
						),
					),
				),
			),
		),
	)
}
