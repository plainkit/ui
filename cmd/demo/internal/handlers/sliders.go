package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/slider"
)

func RenderSlidersContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Sliders")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Adjust numeric values within a range.")),
			),
			html.Div(
				html.AClass("space-y-6"),
				slider.Slider(
					slider.Input(slider.InputProps{ID: "volume", Name: "volume", Min: 0, Max: 100, Value: 50}),
					html.Div(
						html.AClass("flex items-center justify-between text-sm"),
						html.Span(html.Text("Volume")),
						slider.Value(slider.ValueProps{For: "volume"}),
					),
				),
				slider.Slider(
					slider.Input(slider.InputProps{ID: "brightness", Name: "brightness", Min: 20, Max: 120, Value: 90, Step: 5}),
					html.Div(
						html.AClass("flex items-center justify-between text-sm"),
						html.Span(html.Text("Brightness")),
						slider.Value(slider.ValueProps{For: "brightness"}),
					),
				),
			),
		),
	)
}
