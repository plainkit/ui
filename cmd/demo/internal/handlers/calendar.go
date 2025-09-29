package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/calendar"
	"github.com/plainkit/ui/card"
)

func RenderCalendarContent() html.Node {
	return card.Card(card.Props{},
		card.Header(card.HeaderProps{},
			card.Title(card.TitleProps{}, html.Text("Calendar")),
			card.Description(card.DescriptionProps{}, html.Text("A calendar component for date selection.")),
		),
		card.Content(card.ContentProps{},
			html.Div(
				html.AClass("space-y-8"),

				// Basic Calendar
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Basic Calendar")),
					html.Div(
						html.AClass("max-w-md"),
						calendar.Calendar(calendar.Props{
							ID:   "basic-calendar",
							Name: "selected-date",
						}),
					),
				),

				// Calendar with Initial Value
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Calendar with Initial Value")),
					html.Div(
						html.AClass("max-w-md"),
						calendar.Calendar(calendar.Props{
							ID:   "calendar-with-value",
							Name: "preset-date",
							// Value: time.Now(), // Would set current date
						}),
					),
				),

				// Calendar with Custom Locale
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Calendar with French Locale")),
					html.Div(
						html.AClass("max-w-md"),
						calendar.Calendar(calendar.Props{
							ID:        "calendar-french",
							Name:      "french-date",
							LocaleTag: calendar.LocaleTagFrench,
						}),
					),
				),

				// Calendar Starting on Sunday
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Calendar Starting on Sunday")),
					html.Div(
						html.AClass("max-w-md"),
						calendar.Calendar(calendar.Props{
							ID:          "calendar-sunday",
							Name:        "sunday-date",
							StartOfWeek: &[]calendar.Day{calendar.Sunday}[0],
						}),
					),
				),
			),
		),
	)
}
