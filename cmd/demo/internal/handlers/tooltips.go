package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/button"
	"github.com/plainkit/ui/tooltip"
)

func RenderTooltipsContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Tooltips")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Contextual information on hover with flexible positioning.")),
			),
			html.Div(
				html.AClass("grid gap-6 md:grid-cols-2"),
				html.Div(
					html.AClass("space-y-4"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Basic Tooltips")),
					html.Div(
						html.AClass("flex flex-wrap gap-4"),
						tooltip.Trigger(
							tooltip.TriggerProps{For: "tooltip-top"},
							button.Button(button.Props{Variant: button.VariantOutline}, html.Text("Top")),
						),
						tooltip.Content(
							tooltip.ContentProps{
								ID:        "tooltip-top",
								Position:  tooltip.PositionTop,
								ShowArrow: true,
							},
							html.Text("This tooltip appears on top"),
						),
						tooltip.Trigger(
							tooltip.TriggerProps{For: "tooltip-right"},
							button.Button(button.Props{Variant: button.VariantOutline}, html.Text("Right")),
						),
						tooltip.Content(
							tooltip.ContentProps{
								ID:        "tooltip-right",
								Position:  tooltip.PositionRight,
								ShowArrow: true,
							},
							html.Text("This tooltip appears on the right"),
						),
					),
				),
				html.Div(
					html.AClass("space-y-4"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("With Delays")),
					html.Div(
						html.AClass("flex flex-wrap gap-4"),
						tooltip.Trigger(
							tooltip.TriggerProps{For: "tooltip-delayed"},
							button.Button(button.Props{Variant: button.VariantOutline}, html.Text("Hover me")),
						),
						tooltip.Content(
							tooltip.ContentProps{
								ID:            "tooltip-delayed",
								Position:      tooltip.PositionBottom,
								ShowArrow:     true,
								HoverDelay:    500,
								HoverOutDelay: 200,
							},
							html.Text("This tooltip has custom delays"),
						),
					),
				),
			),
		),
	)
}
