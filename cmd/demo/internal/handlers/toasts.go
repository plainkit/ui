package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/button"
	"github.com/plainkit/ui/toast"
)

func RenderToastsContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Toasts")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Lightweight notifications with optional progress indicators and hover pause.")),
			),
			html.Div(
				html.AClass("space-y-6"),
				html.Div(
					html.AClass("flex flex-wrap gap-3"),
					toast.Trigger(
						toast.TriggerProps{
							Toast: toast.Props{
								Title:         "Settings saved",
								Description:   "We'll keep this workspace in sync and let you know if anything changes.",
								Variant:       toast.VariantDefault,
								Duration:      3000,
								Dismissible:   true,
								ShowIndicator: true,
								Icon:          true,
							},
						},
						button.Props{},
						html.Text("Show toast"),
					),
					toast.Trigger(
						toast.TriggerProps{
							Toast: toast.Props{
								Title:         "Success!",
								Description:   "Your changes have been saved successfully.",
								Variant:       toast.VariantSuccess,
								Duration:      4000,
								Dismissible:   true,
								ShowIndicator: true,
								Icon:          true,
							},
						},
						button.Props{Variant: button.VariantSecondary},
						html.Text("Success"),
					),
					toast.Trigger(
						toast.TriggerProps{
							Toast: toast.Props{
								Title:         "Error occurred",
								Description:   "Something went wrong. Please try again.",
								Variant:       toast.VariantError,
								Duration:      4000,
								Dismissible:   true,
								ShowIndicator: true,
								Icon:          true,
							},
						},
						button.Props{Variant: button.VariantDestructive},
						html.Text("Error"),
					),
					toast.Trigger(
						toast.TriggerProps{
							Toast: toast.Props{
								Title:         "Warning",
								Description:   "This action may have unintended consequences.",
								Variant:       toast.VariantWarning,
								Duration:      6000,
								Dismissible:   true,
								ShowIndicator: true,
								Icon:          true,
							},
						},
						button.Props{Variant: button.VariantOutline},
						html.Text("Warning"),
					),
				),
				html.Div(
					html.AClass("rounded-lg border bg-card p-6 text-sm text-muted-foreground"),
					html.P(html.Text("Toasts appear fixed near the chosen viewport corner. Hovering pauses the timer so users can finish reading.")),
				),
			),
		),
	)
}
