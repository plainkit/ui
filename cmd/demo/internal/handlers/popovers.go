package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/button"
	"github.com/plainkit/ui/popover"
)

func RenderPopoversContent() html.Node {
	const (
		infoPopoverID  = "profile-popover"
		hoverPopoverID = "shortcut-popover"
	)

	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Popovers")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Anchor floating panels to any trigger and choose between click or hover interactions.")),
			),
			html.Div(
				html.AClass("grid gap-6 md:grid-cols-2"),
				html.Div(
					html.AClass("space-y-4 rounded-lg border bg-card p-6 shadow-xs"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Click trigger")),
					html.Div(
						html.AClass("flex items-center gap-3"),
						popover.Trigger(
							popover.TriggerProps{ID: "profile-trigger", For: infoPopoverID, TriggerType: popover.TriggerTypeClick},
							button.Button(button.Props{Variant: button.VariantOutline}, html.Text("Open profile")),
						),
						popover.Content(
							popover.ContentProps{ID: infoPopoverID, Class: "w-60 space-y-3 p-4", ShowArrow: true},
							html.Div(
								html.AClass("flex items-center gap-3"),
								html.Div(
									html.AClass("h-10 w-10 rounded-full bg-muted"),
								),
								html.Div(
									html.AClass("space-y-0.5"),
									html.P(html.AClass("text-sm font-medium"), html.Text("Ada Lovelace")),
									html.P(html.AClass("text-xs text-muted-foreground"), html.Text("ada@example.com")),
								),
							),
							html.Div(
								html.AClass("flex gap-2"),
								button.Button(button.Props{Variant: button.VariantOutline, Size: button.SizeSm}, html.Text("View profile")),
								button.Button(button.Props{Size: button.SizeSm}, html.Text("Message")),
							),
						),
					),
				),
				html.Div(
					html.AClass("space-y-4 rounded-lg border bg-card p-6 shadow-xs"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Hover trigger")),
					html.Div(
						html.AClass("flex items-center gap-3"),
						popover.Trigger(
							popover.TriggerProps{ID: "shortcut-trigger", For: hoverPopoverID, TriggerType: popover.TriggerTypeHover},
							html.Span(html.AClass("inline-flex items-center gap-2 rounded-md border border-dashed border-border px-3 py-1 text-sm text-muted-foreground"),
								html.Text("Hover for shortcuts"),
								html.Kbd(html.AClass("rounded bg-muted px-1.5 py-0.5 text-xs"), html.Text("⌘")),
								html.Kbd(html.AClass("rounded bg-muted px-1.5 py-0.5 text-xs"), html.Text("K")),
							),
						),
						popover.Content(
							popover.ContentProps{ID: hoverPopoverID, Class: "w-48 space-y-2 p-3 text-sm", Placement: popover.PlacementTop, ShowArrow: true, MatchWidth: true, HoverDelay: 80, HoverOutDelay: 120},
							html.P(html.AClass("text-xs uppercase tracking-wide text-muted-foreground"), html.Text("Quick actions")),
							html.Ul(
								html.AClass("space-y-1"),
								html.Li(html.Text("Create new doc")),
								html.Li(html.Text("Invite teammate")),
								html.Li(html.Text("Open command palette")),
							),
						),
					),
				),
			),
		),
	)
}
