package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/card"
	"github.com/plainkit/ui/timepicker"
)

func RenderTimePickerContent() html.Node {
	return card.Card(card.Props{},
		card.Header(card.HeaderProps{},
			card.Title(card.TitleProps{}, html.Text("Time Picker")),
			card.Description(card.DescriptionProps{}, html.Text("A time picker component for selecting hours and minutes.")),
		),
		card.Content(card.ContentProps{},
			html.Div(
				html.AClass("space-y-8"),

				// Basic Time Picker (24-hour)
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("24-Hour Format")),
					html.Div(
						html.AClass("max-w-xs"),
						timepicker.TimePicker(timepicker.Props{
							ID:          "time-24h",
							Name:        "time-24h",
							Use12Hours:  false,
							Placeholder: "Select time (24h)",
						}),
					),
				),

				// 12-Hour Time Picker
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("12-Hour Format")),
					html.Div(
						html.AClass("max-w-xs"),
						timepicker.TimePicker(timepicker.Props{
							ID:          "time-12h",
							Name:        "time-12h",
							Use12Hours:  true,
							Placeholder: "Select time (12h)",
						}),
					),
				),

				// Time Picker with 15-minute Steps
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("15-Minute Steps")),
					html.Div(
						html.AClass("max-w-xs"),
						timepicker.TimePicker(timepicker.Props{
							ID:          "time-15min",
							Name:        "time-15min",
							Use12Hours:  true,
							Step:        15,
							Placeholder: "Select time (15min steps)",
						}),
					),
					html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Minutes are shown in 15-minute increments.")),
				),

				// Disabled Time Picker
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Disabled")),
					html.Div(
						html.AClass("max-w-xs"),
						timepicker.TimePicker(timepicker.Props{
							ID:          "time-disabled",
							Name:        "time-disabled",
							Placeholder: "Disabled time picker",
							Disabled:    true,
						}),
					),
				),

				// Time Picker with Error State
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Error State")),
					html.Div(
						html.AClass("max-w-xs"),
						timepicker.TimePicker(timepicker.Props{
							ID:          "time-error",
							Name:        "time-error",
							Placeholder: "Time with error",
							HasError:    true,
							Required:    true,
						}),
					),
					html.P(html.AClass("text-sm text-destructive"), html.Text("Please select a time.")),
				),
			),
		),
	)
}
