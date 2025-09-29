package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/tabs"
)

func RenderTabsContent() html.Node {
	tabsID := "plans-tabs"

	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Tabs")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Organize content horizontally with triggers and panels.")),
			),
			tabs.Tabs(
				tabs.Props{ID: tabsID},
				html.Child(
					tabs.List(tabs.ListProps{TabsID: tabsID},
						tabs.Trigger(tabs.TriggerProps{TabsID: tabsID, Value: "overview", IsActive: true}, html.T("Overview")),
						tabs.Trigger(tabs.TriggerProps{TabsID: tabsID, Value: "billing"}, html.T("Billing")),
						tabs.Trigger(tabs.TriggerProps{TabsID: tabsID, Value: "usage"}, html.T("Usage")),
					),
				),
				html.Child(
					tabs.Content(tabs.ContentProps{TabsID: tabsID, Value: "overview", IsActive: true},
						html.Div(
							html.AClass("space-y-2"),
							html.H3(html.AClass("text-lg font-semibold"), html.Text("Workspace overview")),
							html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Monitor member activity, invites, and billing status.")),
						),
					),
				),
				html.Child(
					tabs.Content(tabs.ContentProps{TabsID: tabsID, Value: "billing"},
						html.Div(
							html.AClass("space-y-2"),
							html.H3(html.AClass("text-lg font-semibold"), html.Text("Billing")),
							html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Manage invoices, payment methods, and receipts.")),
						),
					),
				),
				html.Child(
					tabs.Content(tabs.ContentProps{TabsID: tabsID, Value: "usage"},
						html.Div(
							html.AClass("space-y-2"),
							html.H3(html.AClass("text-lg font-semibold"), html.Text("Usage")),
							html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Track API calls and seat consumption by team.")),
						),
					),
				),
			),
		),
	)
}
