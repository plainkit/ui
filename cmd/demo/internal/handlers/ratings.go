package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/rating"
)

func RenderRatingsContent() html.Node {
	buildItems := func(style rating.Style) []html.DivArg {
		items := make([]html.DivArg, 0, 5)
		for i := 1; i <= 5; i++ {
			items = append(items, html.Child(rating.Item(rating.ItemProps{Value: i, Style: style})))
		}

		return items
	}

	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Ratings")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Support stars, hearts, or emoji reactions with keyboard and pointer interactivity.")),
			),
			html.Div(
				html.AClass("grid gap-6 md:grid-cols-2"),
				html.Div(
					html.AClass("space-y-4 rounded-lg border bg-card p-6 shadow-xs"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Half-star precision")),
					rating.Rating(
						rating.Props{ID: "rating-stars", Name: "product_rating", Precision: 0.5},
						rating.Group(rating.GroupProps{}, buildItems(rating.StyleStar)...),
					),
					html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Click or hover to preview partial fills before committing.")),
				),
				html.Div(
					html.AClass("space-y-4 rounded-lg border bg-card p-6 shadow-xs"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Heart feedback")),
					rating.Rating(
						rating.Props{ID: "rating-hearts", Precision: 1, OnlyInteger: true},
						rating.Group(rating.GroupProps{}, buildItems(rating.StyleHeart)...),
					),
					html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Perfect when you just need a simple like-meter.")),
				),
				html.Div(
					html.AClass("space-y-4 rounded-lg border bg-card p-6 shadow-xs md:col-span-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Emoji mood")),
					rating.Rating(
						rating.Props{ID: "rating-emoji", Value: 4.2, Precision: 0.5, ReadOnly: true},
						rating.Group(rating.GroupProps{}, buildItems(rating.StyleEmoji)...),
					),
					html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Read-only ratings show aggregate sentiment without interaction.")),
				),
			),
		),
	)
}
