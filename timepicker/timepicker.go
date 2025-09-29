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
	"github.com/plainkit/ui/internal/classnames"
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

// TimePicker renders a time picker component with hour/minute selection
func TimePicker(props Props, args ...html.DivArg) html.Node {
	id := props.ID
	if id == "" {
		id = randomID("timepicker")
	}

	name := props.Name
	if name == "" {
		name = id
	}

	placeholder := props.Placeholder
	if placeholder == "" {
		placeholder = "Select time"
	}

	amLabel := props.AMLabel
	if amLabel == "" {
		amLabel = "AM"
	}

	pmLabel := props.PMLabel
	if pmLabel == "" {
		pmLabel = "PM"
	}

	step := props.Step
	if step <= 0 {
		step = 1
	}

	contentID := id + "-content"

	var valueString string
	if !props.Value.IsZero() {
		valueString = props.Value.Format("15:04")
	}

	var minTimeString string
	if !props.MinTime.IsZero() {
		minTimeString = props.MinTime.Format("15:04")
	}

	var maxTimeString string
	if !props.MaxTime.IsZero() {
		maxTimeString = props.MaxTime.Format("15:04")
	}

	divArgs := []html.DivArg{
		html.AClass(classnames.Merge("relative inline-block w-full", "")),
	}

	// Hidden input
	hiddenInput := html.Input(
		html.AType("hidden"),
		html.AName(name),
		html.AValue(valueString),
		html.AId(id+"-hidden"),
		html.AData("pui-timepicker-hidden-input", "true"),
		func() html.InputArg {
			if props.Form != "" {
				return html.AForm(props.Form)
			}
			return html.AAria("", "")
		}(),
		func() html.InputArg {
			if props.Required {
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
			Class: classnames.Merge(
				// Base styles matching input
				"w-full h-9 px-3 py-1 text-base md:text-sm",
				"flex items-center justify-between",
				"rounded-md border border-input bg-transparent shadow-xs transition-[color,box-shadow] outline-none",
				// Dark mode background
				"dark:bg-input/30",
				// Selection styles
				"selection:bg-primary selection:text-primary-foreground",
				// Focus styles
				"focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px]",
				// Error/Invalid styles
				"aria-invalid:ring-destructive/20 aria-invalid:border-destructive dark:aria-invalid:ring-destructive/40",
				func() string {
					if props.HasError {
						return "border-destructive ring-destructive/20 dark:ring-destructive/40"
					}
					return ""
				}(),
				props.Class,
			),
			Disabled: props.Disabled,
			Attrs: []html.Global{
				html.AData("pui-timepicker", "true"),
				html.AData("pui-timepicker-use12hours", fmt.Sprintf("%t", props.Use12Hours)),
				html.AData("pui-timepicker-am-label", amLabel),
				html.AData("pui-timepicker-pm-label", pmLabel),
				html.AData("pui-timepicker-placeholder", placeholder),
				html.AData("pui-timepicker-step", fmt.Sprintf("%d", step)),
				html.AData("pui-timepicker-min-time", minTimeString),
				html.AData("pui-timepicker-max-time", maxTimeString),
				func() html.Global {
					if props.HasError {
						return html.AAria("invalid", "true")
					}
					return html.AAria("", "")
				}(),
			},
		},
			html.Span(
				html.AData("pui-timepicker-display", ""),
				html.AClass("text-left grow text-muted-foreground"),
				html.Text(placeholder),
			),
			html.Span(
				html.AClass("text-muted-foreground flex items-center ml-2"),
				lucide.Clock(html.AClass("h-4 w-4")),
			),
		),
	)

	// Popup content
	popupContent := popover.Content(
		popover.ContentProps{
			ID:        contentID,
			Placement: popover.PlacementBottomStart,
			Class:     "p-0 w-80",
		},
		card.Card(card.Props{
			Class: "border-0 shadow-none",
		},
			card.Content(card.ContentProps{
				Class: "p-4",
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
						html.AClass("grid grid-cols-2 gap-3 mb-4"),

						// Hour selection
						html.Div(
							html.AClass("space-y-2"),
							html.Label(
								html.AClass("text-sm font-medium"),
								html.Text("Hour"),
							),
							html.Div(
								html.AClass("max-h-32 overflow-y-auto border rounded-md bg-background"),
								createHourList(props.Use12Hours),
							),
						),

						// Minute selection
						html.Div(
							html.AClass("space-y-2"),
							html.Label(
								html.AClass("text-sm font-medium"),
								html.Text("Minute"),
							),
							html.Div(
								html.AClass("max-h-32 overflow-y-auto border rounded-md bg-background"),
								createMinuteList(step),
							),
						),
					),

					// AM/PM selector and action buttons
					html.Div(
						html.AClass("flex justify-between items-center"),

						// AM/PM selector (conditionally rendered)
						func() html.Node {
							if props.Use12Hours {
								return html.Div(
									html.AClass("flex gap-1"),
									html.Button(
										html.AType("button"),
										html.AData("pui-timepicker-period", "AM"),
										html.AData("pui-timepicker-active", "false"),
										html.AClass("px-3 py-1 text-sm rounded-md border transition-colors hover:bg-accent hover:text-accent-foreground data-[pui-timepicker-active=true]:bg-primary data-[pui-timepicker-active=true]:text-primary-foreground data-[pui-timepicker-active=true]:hover:bg-primary/90"),
										html.Text(amLabel),
									),
									html.Button(
										html.AType("button"),
										html.AData("pui-timepicker-period", "PM"),
										html.AData("pui-timepicker-active", "false"),
										html.AClass("px-3 py-1 text-sm rounded-md border transition-colors hover:bg-accent hover:text-accent-foreground data-[pui-timepicker-active=true]:bg-primary data-[pui-timepicker-active=true]:text-primary-foreground data-[pui-timepicker-active=true]:hover:bg-primary/90"),
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

	divArgs = append(divArgs,
		hiddenInput,
		triggerButton,
		popupContent,
	)
	divArgs = append(divArgs, args...)

	return html.Div(divArgs...).WithAssets("", timepickerJS, "ui-timepicker")
}

func createHourList(use12Hours bool) html.Node {
	divArgs := []html.DivArg{
		html.AData("pui-timepicker-hour-list", "true"),
		html.AClass("p-1 space-y-0.5"),
	}

	buttonClass := "w-full px-2 py-1 text-sm rounded transition-colors text-left hover:bg-accent hover:text-accent-foreground data-[pui-timepicker-selected=true]:bg-primary data-[pui-timepicker-selected=true]:text-primary-foreground data-[pui-timepicker-selected=true]:hover:bg-primary/90"

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
		html.AClass("p-1 space-y-0.5"),
	}

	buttonClass := "w-full px-2 py-1 text-sm rounded transition-colors text-left hover:bg-accent hover:text-accent-foreground data-[pui-timepicker-selected=true]:bg-primary data-[pui-timepicker-selected=true]:text-primary-foreground data-[pui-timepicker-selected=true]:hover:bg-primary/90"

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
