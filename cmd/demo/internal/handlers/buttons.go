package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/button"
)

func RenderButtonsContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Variants")),
				html.P(html.AClass("text-sm text-slate-400"), html.Text("Standard button appearances using design tokens.")),
			),
			html.Div(
				html.AClass("flex flex-wrap gap-3"),
				button.Button(button.Props{}, html.T("Default")),
				button.Button(button.Props{Variant: button.VariantSecondary}, html.T("Secondary")),
				button.Button(button.Props{Variant: button.VariantOutline}, html.T("Outline")),
				button.Button(button.Props{Variant: button.VariantDestructive}, html.T("Destructive")),
				button.Button(button.Props{Variant: button.VariantGhost}, html.T("Ghost")),
				button.Button(
					button.Props{Variant: button.VariantLink, Href: "https://plainkit.dev", Target: "_blank"},
					html.T("Link"),
				),
			),
		),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Sizes")),
				html.P(html.AClass("text-sm text-slate-400"), html.Text("Adjust sizing tokens for different contexts.")),
			),
			html.Div(
				html.AClass("flex flex-wrap items-end gap-3"),
				button.Button(button.Props{}, html.T("Default")),
				button.Button(button.Props{Size: button.SizeSm}, html.T("Small")),
				button.Button(button.Props{Size: button.SizeLg}, html.T("Large")),
				button.Button(button.Props{Size: button.SizeIcon}, html.T("+")),
			),
		),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("States")),
				html.P(html.AClass("text-sm text-slate-400"), html.Text("Examples of disabled and full-width buttons.")),
			),
			html.Div(
				html.AClass("flex flex-wrap gap-3"),
				button.Button(button.Props{Disabled: true}, html.T("Disabled")),
				button.Button(button.Props{FullWidth: true}, html.T("Full width")),
				button.Button(button.Props{Href: "https://plainkit.dev/docs"}, html.T("As link")),
			),
		),
	)
}
