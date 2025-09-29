package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/card"
	"github.com/plainkit/ui/tagsinput"
)

func RenderTagsInputContent() html.Node {
	return card.Card(card.Props{},
		card.Header(card.HeaderProps{},
			card.Title(card.TitleProps{}, html.Text("Tags Input")),
			card.Description(card.DescriptionProps{}, html.Text("An input component for entering multiple tags.")),
		),
		card.Content(card.ContentProps{},
			html.Div(
				html.AClass("space-y-8"),

				// Basic Tags Input
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Basic Tags Input")),
					html.Div(
						html.AClass("max-w-md"),
						tagsinput.TagsInput(tagsinput.Props{
							ID:          "basic-tags",
							Name:        "tags",
							Placeholder: "Type and press Enter or comma to add tags",
						}),
					),
					html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Press Enter or comma to add tags. Backspace to remove the last tag when input is empty.")),
				),

				// Pre-filled Tags Input
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Pre-filled Tags")),
					html.Div(
						html.AClass("max-w-md"),
						tagsinput.TagsInput(tagsinput.Props{
							ID:          "prefilled-tags",
							Name:        "prefilled-tags",
							Value:       []string{"React", "Go", "JavaScript", "TypeScript"},
							Placeholder: "Add more tags...",
						}),
					),
				),

				// Disabled Tags Input
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Disabled Tags Input")),
					html.Div(
						html.AClass("max-w-md"),
						tagsinput.TagsInput(tagsinput.Props{
							ID:          "disabled-tags",
							Name:        "disabled-tags",
							Value:       []string{"Read-only", "Disabled"},
							Placeholder: "Cannot add tags",
							Disabled:    true,
						}),
					),
				),

				// Tags Input with Error State
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Error State")),
					html.Div(
						html.AClass("max-w-md"),
						tagsinput.TagsInput(tagsinput.Props{
							ID:          "error-tags",
							Name:        "error-tags",
							Placeholder: "This field has an error",
							HasError:    true,
						}),
					),
					html.P(html.AClass("text-sm text-destructive"), html.Text("Please add at least one tag.")),
				),
			),
		),
	)
}
