package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/button"
	"github.com/plainkit/ui/dropdown"
)

func RenderDropdownsContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Dropdowns")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Floating menus with items, groups, and keyboard navigation.")),
			),
			html.Div(
				html.AClass("flex flex-wrap gap-6"),
				html.Div(
					html.AClass("space-y-4"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Basic Menu")),
					dropdown.Dropdown(
						dropdown.Props{},
						dropdown.Trigger(
							dropdown.TriggerProps{For: "basic-dropdown"},
							button.Props{Variant: button.VariantOutline},
							html.Text("Options"),
							lucide.ChevronDown(html.AClass("ml-2 size-4")),
						),
						dropdown.Content(
							dropdown.ContentProps{ID: "basic-dropdown"},
							dropdown.Item(dropdown.ItemProps{}, html.Span(html.Text("Profile"))),
							dropdown.Item(dropdown.ItemProps{}, html.Span(html.Text("Settings"))),
							dropdown.Separator(dropdown.SeparatorProps{}),
							dropdown.Item(dropdown.ItemProps{}, html.Span(html.Text("Sign out"))),
						),
					),
				),
				html.Div(
					html.AClass("space-y-4"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("With Groups")),
					dropdown.Dropdown(
						dropdown.Props{},
						dropdown.Trigger(
							dropdown.TriggerProps{For: "grouped-dropdown"},
							button.Props{Variant: button.VariantOutline},
							html.Text("Actions"),
							lucide.ChevronDown(html.AClass("ml-2 size-4")),
						),
						dropdown.Content(
							dropdown.ContentProps{ID: "grouped-dropdown"},
							dropdown.Group(
								dropdown.GroupProps{},
								dropdown.Label(dropdown.LabelProps{}, html.Text("Account")),
								dropdown.Item(dropdown.ItemProps{}, html.Span(html.Text("Profile"))),
								dropdown.Item(dropdown.ItemProps{}, html.Span(html.Text("Billing"))),
							),
							dropdown.Separator(dropdown.SeparatorProps{}),
							dropdown.Group(
								dropdown.GroupProps{},
								dropdown.Label(dropdown.LabelProps{}, html.Text("Actions")),
								dropdown.Item(dropdown.ItemProps{}, html.Span(html.Text("New file")), dropdown.Shortcut(dropdown.ShortcutProps{}, html.Text("⌘N"))),
								dropdown.Item(dropdown.ItemProps{}, html.Span(html.Text("Save")), dropdown.Shortcut(dropdown.ShortcutProps{}, html.Text("⌘S"))),
							),
						),
					),
				),
			),
		),
	)
}
