package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/card"
	"github.com/plainkit/ui/carousel"
)

func RenderCarouselContent() html.Node {
	return card.Card(card.Props{},
		card.Header(card.HeaderProps{},
			card.Title(card.TitleProps{}, html.Text("Carousel")),
			card.Description(card.DescriptionProps{}, html.Text("A carousel component for displaying sliding content.")),
		),
		card.Content(card.ContentProps{},
			html.Div(
				html.AClass("space-y-8"),

				// Basic Carousel
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Basic Carousel")),
					html.Div(
						html.AClass("relative"),
						carousel.Carousel(carousel.Props{
							ID: "basic-carousel",
						},
							carousel.Content(carousel.ContentProps{},
								carousel.Item(carousel.ItemProps{},
									html.Div(
										html.AClass("bg-gradient-to-r from-blue-500 to-purple-600 h-64 flex items-center justify-center text-white text-xl font-bold rounded-lg"),
										html.Text("Slide 1"),
									),
								),
								carousel.Item(carousel.ItemProps{},
									html.Div(
										html.AClass("bg-gradient-to-r from-green-500 to-blue-600 h-64 flex items-center justify-center text-white text-xl font-bold rounded-lg"),
										html.Text("Slide 2"),
									),
								),
								carousel.Item(carousel.ItemProps{},
									html.Div(
										html.AClass("bg-gradient-to-r from-red-500 to-pink-600 h-64 flex items-center justify-center text-white text-xl font-bold rounded-lg"),
										html.Text("Slide 3"),
									),
								),
							),
							carousel.Previous(carousel.PreviousProps{}),
							carousel.Next(carousel.NextProps{}),
							carousel.Indicators(carousel.IndicatorsProps{Count: 3}),
						),
					),
				),

				// Autoplay Carousel with Loop
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Autoplay Carousel")),
					html.Div(
						html.AClass("relative"),
						carousel.Carousel(carousel.Props{
							ID:       "autoplay-carousel",
							Autoplay: true,
							Interval: 3000,
							Loop:     true,
						},
							carousel.Content(carousel.ContentProps{},
								carousel.Item(carousel.ItemProps{},
									html.Div(
										html.AClass("bg-gradient-to-r from-yellow-400 to-orange-500 h-48 flex items-center justify-center text-white text-lg font-semibold rounded-lg"),
										html.Text("Auto Slide 1 - 3s intervals"),
									),
								),
								carousel.Item(carousel.ItemProps{},
									html.Div(
										html.AClass("bg-gradient-to-r from-purple-500 to-indigo-600 h-48 flex items-center justify-center text-white text-lg font-semibold rounded-lg"),
										html.Text("Auto Slide 2 - Loops infinitely"),
									),
								),
								carousel.Item(carousel.ItemProps{},
									html.Div(
										html.AClass("bg-gradient-to-r from-teal-500 to-cyan-600 h-48 flex items-center justify-center text-white text-lg font-semibold rounded-lg"),
										html.Text("Auto Slide 3 - Hover to pause"),
									),
								),
							),
							carousel.Previous(carousel.PreviousProps{}),
							carousel.Next(carousel.NextProps{}),
							carousel.Indicators(carousel.IndicatorsProps{Count: 3}),
						),
					),
				),
			),
		),
	)
}
