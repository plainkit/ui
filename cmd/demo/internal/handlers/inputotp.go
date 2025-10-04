package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/form"
	"github.com/plainkit/ui/inputotp"
)

func RenderInputOTPContent() html.Node {
	buildSlots := func(start int, hasError bool) []html.DivArg {
		items := make([]html.DivArg, 0, 3)
		for i := 0; i < 3; i++ {
			items = append(items, inputotp.Slot(inputotp.SlotProps{
				Index:       start + i,
				Placeholder: "0",
				HasError:    hasError,
			}))
		}

		return items
	}

	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Input OTP")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Capture short verification codes with smart focus management and paste handling.")),
			),
			html.Div(
				html.AClass("space-y-8"),
				html.Div(
					html.AClass("space-y-3 rounded-lg border bg-card p-6 shadow-xs"),
					html.Label(html.AClass("text-sm font-medium"), html.Text("Two-factor code")),
					inputotp.InputOTP(
						inputotp.Props{ID: "otp-primary", Name: "security_code", Required: true, HasError: true},
						inputotp.Group(buildSlots(0, true)...),
						inputotp.Separator(html.Text("-")),
						inputotp.Group(buildSlots(3, true)...),
					),
					form.Message(form.MessageProps{Variant: form.MessageVariantError}, html.Text("The verification code is incorrect.")),
				),
				html.Div(
					html.AClass("space-y-3 rounded-lg border bg-card p-6 shadow-xs"),
					html.Label(html.AClass("text-sm font-medium"), html.Text("SMS code")),
					inputotp.InputOTP(
						inputotp.Props{ID: "otp-secondary", Name: "sms_code", Value: "123456", Autofocus: true, Class: "flex-wrap gap-y-3"},
						inputotp.Group(buildSlots(0, false)...),
						inputotp.Separator(html.Text("-")),
						inputotp.Group(buildSlots(3, false)...),
					),
					form.Description(form.DescriptionProps{}, html.Text("Autofocus jumps to the next slot as the user types.")),
				),
			),
		),
	)
}
