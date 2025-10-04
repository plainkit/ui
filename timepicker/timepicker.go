package timepicker

import (
	"crypto/rand"
	_ "embed"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/button"
	"github.com/plainkit/ui/card"
	"github.com/plainkit/ui/internal/styles"
	"github.com/plainkit/ui/popover"
)

type Props struct {
	ID          string
	Class       string
	Attrs       []html.Global
	Name        string
	Form        string
	Value       time.Time
	MinTime     time.Time
	MaxTime     time.Time
	Step        int
	Use12Hours  bool
	AMLabel     string
	PMLabel     string
	Placeholder string
	Required    bool
	Disabled    bool
	HasError    bool
}

func divArgsFromProps(baseClass string, extra ...string) func(p Props) []html.DivArg {
	return func(p Props) []html.DivArg {
		args := []html.DivArg{html.AClass(html.ClassMerge(append([]string{baseClass}, append(extra, p.Class)...)...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p Props) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	id := p.ID
	if id == "" {
		id = randomID("timepicker")
	}

	name := p.Name
	if name == "" {
		name = id
	}

	placeholder := p.Placeholder
	if placeholder == "" {
		placeholder = "Select time"
	}

	amLabel := p.AMLabel
	if amLabel == "" {
		amLabel = "AM"
	}

	pmLabel := p.PMLabel
	if pmLabel == "" {
		pmLabel = "PM"
	}

	step := p.Step
	if step <= 0 {
		step = 1
	}

	contentID := id + "-content"

	var valueString string
	if !p.Value.IsZero() {
		valueString = p.Value.Format("15:04")
	}

	var minTimeString string
	if !p.MinTime.IsZero() {
		minTimeString = p.MinTime.Format("15:04")
	}

	var maxTimeString string
	if !p.MaxTime.IsZero() {
		maxTimeString = p.MaxTime.Format("15:04")
	}

	args := divArgsFromProps(styles.SurfaceMuted("relative inline-flex w-full flex-col gap-3 p-4"))(p)

	// Hidden input
	hiddenInput := html.Input(
		html.AType("hidden"),
		html.AName(name),
		html.AValue(valueString),
		html.AId(id+"-hidden"),
		html.AData("pui-timepicker-hidden-input", "true"),
		func() html.InputArg {
			if p.Form != "" {
				return html.AForm(p.Form)
			}

			return html.AAria("", "")
		}(),
		func() html.InputArg {
			if p.Required {
				return html.ARequired()
			}

			return html.AAria("", "")
		}(),
	)

	// Trigger button
	triggerButton := popover.Trigger(
		popover.TriggerProps{
			For:         contentID,
			TriggerType: popover.TriggerTypeClick,
		},
		button.Button(button.Props{
			ID:      id,
			Variant: button.VariantOutline,
			Class: html.ClassMerge(
				styles.Input("timepicker-trigger flex h-11 w-full items-center justify-between gap-3 pr-3 text-sm", "aria-invalid:border-destructive aria-invalid:ring-destructive/30"),
				func() string {
					if p.HasError {
						return "border-destructive ring-destructive/30"
					}

					return ""
				}(),
				p.Class,
			),
			Disabled: p.Disabled,
			Attrs: []html.Global{
				html.AData("pui-timepicker", "true"),
				html.AData("pui-timepicker-use12hours", fmt.Sprintf("%t", p.Use12Hours)),
				html.AData("pui-timepicker-am-label", amLabel),
				html.AData("pui-timepicker-pm-label", pmLabel),
				html.AData("pui-timepicker-placeholder", placeholder),
				html.AData("pui-timepicker-step", fmt.Sprintf("%d", step)),
				html.AData("pui-timepicker-min-time", minTimeString),
				html.AData("pui-timepicker-max-time", maxTimeString),
				func() html.Global {
					if p.HasError {
						return html.AAria("invalid", "true")
					}

					return html.AAria("", "")
				}(),
			},
		},
			html.Span(
				html.AData("pui-timepicker-display", ""),
				html.AClass(styles.SubtleText("grow text-left text-sm")),
				html.Text(placeholder),
			),
			html.Span(
				html.AClass("ml-3 flex items-center text-muted-foreground/70"),
				lucide.Clock(html.AClass("h-4 w-4")),
			),
		),
	)

	// Popup content
	popupContent := popover.Content(
		popover.ContentProps{
			ID:        contentID,
			Placement: popover.PlacementBottomStart,
			Class:     styles.Panel("w-80 p-0"),
		},
		card.Card(card.Props{
			Class: "border-none bg-transparent p-0 shadow-none",
		},
			card.Content(card.ContentProps{
				Class: "space-y-4 p-4",
			},
				html.Div(
					html.AData("pui-timepicker-popup", "true"),
					html.AData("pui-timepicker-input-name", name),
					html.AData("pui-timepicker-parent-id", id),
					func() html.Global {
						if valueString != "" {
							return html.AData("pui-timepicker-value", valueString)
						}

						return html.AAria("", "")
					}(),

					// Time selection grid
					html.Div(
						html.AClass("grid grid-cols-2 gap-4"),

						// Hour selection
						html.Div(
							html.AClass("flex flex-col gap-2"),
							html.Label(
								html.AClass(styles.SubHeading("text-xs uppercase tracking-wide text-muted-foreground/70")),
								html.Text("Hour"),
							),
							html.Div(
								html.AClass(styles.SurfaceMuted("max-h-48 overflow-y-auto rounded-xl p-1")),
								createHourList(p.Use12Hours),
							),
						),

						// Minute selection
						html.Div(
							html.AClass("flex flex-col gap-2"),
							html.Label(
								html.AClass(styles.SubHeading("text-xs uppercase tracking-wide text-muted-foreground/70")),
								html.Text("Minute"),
							),
							html.Div(
								html.AClass(styles.SurfaceMuted("max-h-48 overflow-y-auto rounded-xl p-1")),
								createMinuteList(step),
							),
						),
					),

					// AM/PM selector and action buttons
					html.Div(
						html.AClass("flex items-center justify-between"),

						// AM/PM selector (conditionally rendered)
						func() html.Node {
							if p.Use12Hours {
								return html.Div(
									html.AClass("flex gap-2"),
									html.Button(
										html.AType("button"),
										html.AData("pui-timepicker-period", "AM"),
										html.AData("pui-timepicker-active", "false"),
										html.AClass(styles.InteractiveGhost("px-4 py-1.5 text-sm", "data-[pui-timepicker-active=true]:bg-primary data-[pui-timepicker-active=true]:text-primary-foreground data-[pui-timepicker-active=true]:shadow")),
										html.Text(amLabel),
									),
									html.Button(
										html.AType("button"),
										html.AData("pui-timepicker-period", "PM"),
										html.AData("pui-timepicker-active", "false"),
										html.AClass(styles.InteractiveGhost("px-4 py-1.5 text-sm", "data-[pui-timepicker-active=true]:bg-primary data-[pui-timepicker-active=true]:text-primary-foreground data-[pui-timepicker-active=true]:shadow")),
										html.Text(pmLabel),
									),
								)
							}

							return html.Div()
						}()),

					// Done button
					button.Button(button.Props{
						Type:    button.TypeButton,
						Variant: button.VariantSecondary,
						Size:    button.SizeSm,
						Attrs: []html.Global{
							html.AData("pui-timepicker-done", "true"),
						},
					}, html.Text("Done")),
				),
			),
		),
	)

	*children = append(*children,
		hiddenInput,
		triggerButton,
		popupContent,
	)

	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

// TimePicker renders a time picker component with hour/minute selection
func TimePicker(args ...html.DivArg) html.Node {
	var (
		props Props
		rest  []html.DivArg
	)

	for _, a := range args {
		if v, ok := a.(Props); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Div(append([]html.DivArg{props}, rest...)...).WithAssets("", timepickerJS, "ui-timepicker")
}

func createHourList(use12Hours bool) html.Node {
	divArgs := []html.DivArg{
		html.AData("pui-timepicker-hour-list", "true"),
		html.AClass("space-y-1"),
	}

	buttonClass := styles.InteractiveGhost(
		"w-full justify-start rounded-lg px-3 py-2 text-sm",
		"data-[pui-timepicker-selected=true]:bg-primary data-[pui-timepicker-selected=true]:text-primary-foreground data-[pui-timepicker-selected=true]:shadow",
	)

	if use12Hours {
		// 12-hour format: 12, 01-11
		divArgs = append(divArgs, html.Button(
			html.AType("button"),
			html.AData("pui-timepicker-hour", "0"),
			html.AData("pui-timepicker-selected", "false"),
			html.AClass(buttonClass),
			html.Text("12"),
		))

		for hour := 1; hour <= 11; hour++ {
			divArgs = append(divArgs, html.Button(
				html.AType("button"),
				html.AData("pui-timepicker-hour", strconv.Itoa(hour)),
				html.AData("pui-timepicker-selected", "false"),
				html.AClass(buttonClass),
				html.Text(fmt.Sprintf("%02d", hour)),
			))
		}
	} else {
		// 24-hour format: 00-23
		for hour := 0; hour < 24; hour++ {
			divArgs = append(divArgs, html.Button(
				html.AType("button"),
				html.AData("pui-timepicker-hour", strconv.Itoa(hour)),
				html.AData("pui-timepicker-selected", "false"),
				html.AClass(buttonClass),
				html.Text(fmt.Sprintf("%02d", hour)),
			))
		}
	}

	return html.Div(divArgs...)
}

func createMinuteList(step int) html.Node {
	divArgs := []html.DivArg{
		html.AData("pui-timepicker-minute-list", "true"),
		html.AClass("space-y-1"),
	}

	buttonClass := styles.InteractiveGhost(
		"w-full justify-start rounded-lg px-3 py-2 text-sm",
		"data-[pui-timepicker-selected=true]:bg-primary data-[pui-timepicker-selected=true]:text-primary-foreground data-[pui-timepicker-selected=true]:shadow",
	)

	for minute := 0; minute < 60; minute += step {
		divArgs = append(divArgs, html.Button(
			html.AType("button"),
			html.AData("pui-timepicker-minute", strconv.Itoa(minute)),
			html.AData("pui-timepicker-selected", "false"),
			html.AClass(buttonClass),
			html.Text(fmt.Sprintf("%02d", minute)),
		))
	}

	return html.Div(divArgs...)
}

func randomID(prefix string) string {
	buf := make([]byte, 4)
	if _, err := rand.Read(buf); err != nil {
		return prefix + "-id"
	}

	return prefix + "-" + hex.EncodeToString(buf)
}

//go:embed timepicker.js
var timepickerJS string
