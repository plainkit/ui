package timepicker

import (
	"crypto/rand"
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

const timepickerJS = `(function () {
  'use strict';

  // Utility functions
  function parseTime(str) {
    const match = str && str.match(/^(\d{1,2}):(\d{2})$/);
    if (!match) return null;
    const hour = Number(match[1]);
    const minute = Number(match[2]);
    return (hour >= 0 && hour <= 23 && minute >= 0 && minute <= 59) ? { hour, minute } : null;
  }

  function formatTime(hour, minute, use12Hours) {
    if (hour === null || minute === null) return null;
    const pad = n => n.toString().padStart(2, '0');

    if (use12Hours) {
      const h = hour === 0 ? 12 : hour > 12 ? hour - 12 : hour;
      return pad(h) + ':' + pad(minute) + ' ' + (hour >= 12 ? 'PM' : 'AM');
    }
    return pad(hour) + ':' + pad(minute);
  }

  function isValidTime(hour, minute, minTime, maxTime) {
    if (!minTime && !maxTime) return true;
    const timeInMinutes = hour * 60 + minute;

    if (minTime) {
      const minInMinutes = minTime.hour * 60 + minTime.minute;
      if (timeInMinutes < minInMinutes) return false;
    }

    if (maxTime) {
      const maxInMinutes = maxTime.hour * 60 + maxTime.minute;
      if (timeInMinutes > maxInMinutes) return false;
    }

    return true;
  }

  // DOM helpers
  function findTrigger(element) {
    const popup = element.closest('[data-pui-timepicker-popup]');
    if (!popup) return null;

    const popupId = popup.closest('[id]') && popup.closest('[id]').id;
    if (!popupId) return null;

    return document.getElementById(popupId.replace('-content', ''));
  }

  function getElements(trigger) {
    const popupId = trigger.id + '-content';
    const popupContainer = document.getElementById(popupId);
    const popup = popupContainer && popupContainer.querySelector('[data-pui-timepicker-popup]');
    if (!popup) return null;

    return {
      popup,
      hourList: popup.querySelector('[data-pui-timepicker-hour-list]'),
      minuteList: popup.querySelector('[data-pui-timepicker-minute-list]'),
      hiddenInput: document.getElementById(trigger.id + '-hidden') ||
                   (trigger.parentElement && trigger.parentElement.querySelector('[data-pui-timepicker-hidden-input]'))
    };
  }

  // State management
  function getState(trigger) {
    return {
      hour: trigger.dataset.tuiTimepickerCurrentHour ? parseInt(trigger.dataset.tuiTimepickerCurrentHour) : null,
      minute: trigger.dataset.tuiTimepickerCurrentMinute ? parseInt(trigger.dataset.tuiTimepickerCurrentMinute) : null,
      use12Hours: trigger.getAttribute('data-pui-timepicker-use12hours') === 'true',
      step: parseInt(trigger.getAttribute('data-pui-timepicker-step') || '1'),
      minTime: parseTime(trigger.getAttribute('data-pui-timepicker-min-time')),
      maxTime: parseTime(trigger.getAttribute('data-pui-timepicker-max-time')),
      placeholder: trigger.getAttribute('data-pui-timepicker-placeholder') || 'Select time'
    };
  }

  function setState(trigger, hour, minute) {
    if (hour !== null) {
      trigger.dataset.tuiTimepickerCurrentHour = hour;
    } else {
      delete trigger.dataset.tuiTimepickerCurrentHour;
    }

    if (minute !== null) {
      trigger.dataset.tuiTimepickerCurrentMinute = minute;
    } else {
      delete trigger.dataset.tuiTimepickerCurrentMinute;
    }

    updateDisplay(trigger);
  }

  // Display updates
  function updateDisplay(trigger) {
    const state = getState(trigger);
    const elements = getElements(trigger);

    // Update trigger display
    const display = trigger.querySelector('[data-pui-timepicker-display]');
    if (display) {
      const formatted = formatTime(state.hour, state.minute, state.use12Hours);
      display.textContent = formatted || state.placeholder;
      display.classList.toggle('text-muted-foreground', !formatted);
    }

    // Update hidden input
    if (elements && elements.hiddenInput) {
      elements.hiddenInput.value = (state.hour !== null && state.minute !== null) ?
        formatTime(state.hour, state.minute, false) : '';
    }

    // Update selections if popup is visible
    if (elements && elements.hourList && elements.minuteList) {
      updateSelections(elements, state);
    }
  }

  function updateSelections(elements, state) {
    // Update hour buttons
    elements.hourList.querySelectorAll('[data-pui-timepicker-hour]').forEach(btn => {
      const hour = parseInt(btn.getAttribute('data-pui-timepicker-hour'));
      let isSelected = false;

      if (state.hour !== null) {
        if (state.use12Hours) {
          isSelected = (hour === state.hour) ||
                      (hour === 0 && state.hour === 12) ||
                      (hour === state.hour - 12 && state.hour > 12);
        } else {
          isSelected = hour === state.hour;
        }
      }

      btn.setAttribute('data-pui-timepicker-selected', isSelected);

      // Check validity
      let valid = false;
      for (let m = 0; m < 60; m++) {
        if (isValidTime(hour, m, state.minTime, state.maxTime)) {
          valid = true;
          break;
        }
      }

      btn.disabled = !valid;
      btn.classList.toggle('opacity-50', !valid);
      btn.classList.toggle('cursor-not-allowed', !valid);
    });

    // Update minute buttons
    elements.minuteList.querySelectorAll('[data-pui-timepicker-minute]').forEach(btn => {
      const minute = parseInt(btn.getAttribute('data-pui-timepicker-minute'));
      const isSelected = minute === state.minute;
      const valid = state.hour === null || isValidTime(state.hour, minute, state.minTime, state.maxTime);

      btn.setAttribute('data-pui-timepicker-selected', isSelected);
      btn.disabled = !valid;
      btn.classList.toggle('opacity-50', !valid);
      btn.classList.toggle('cursor-not-allowed', !valid);
    });

    // Update AM/PM buttons
    const amBtn = elements.popup.querySelector('[data-pui-timepicker-period="AM"]');
    const pmBtn = elements.popup.querySelector('[data-pui-timepicker-period="PM"]');

    if (amBtn && pmBtn) {
      const isAM = state.hour === null || state.hour < 12;
      amBtn.setAttribute('data-pui-timepicker-active', isAM);
      pmBtn.setAttribute('data-pui-timepicker-active', !isAM);
    }
  }

  // Event handlers
  document.addEventListener('click', (e) => {
    const target = e.target;

    // Hour selection
    if (target.matches('[data-pui-timepicker-hour]') && !target.disabled) {
      const trigger = findTrigger(target);
      if (!trigger) return;

      const state = getState(trigger);
      let hour = parseInt(target.getAttribute('data-pui-timepicker-hour'));

      if (state.use12Hours) {
        const isPM = state.hour !== null && state.hour >= 12;
        hour = hour === 0 ? (isPM ? 12 : 0) : (isPM ? hour + 12 : hour);
      }

      if (!isValidTime(hour, state.minute, state.minTime, state.maxTime)) {
        // Find first valid minute
        for (let m = 0; m < 60; m += state.step) {
          if (isValidTime(hour, m, state.minTime, state.maxTime)) {
            setState(trigger, hour, m);
            return;
          }
        }
      } else {
        setState(trigger, hour, state.minute);
      }
      return;
    }

    // Minute selection
    if (target.matches('[data-pui-timepicker-minute]') && !target.disabled) {
      const trigger = findTrigger(target);
      if (!trigger) return;

      const state = getState(trigger);
      const minute = parseInt(target.getAttribute('data-pui-timepicker-minute'));

      if (state.hour === null || isValidTime(state.hour, minute, state.minTime, state.maxTime)) {
        setState(trigger, state.hour, minute);
      }
      return;
    }

    // AM/PM selection
    if (target.matches('[data-pui-timepicker-period]')) {
      const trigger = findTrigger(target);
      if (!trigger) return;

      const state = getState(trigger);
      if (state.hour === null) return;

      const period = target.getAttribute('data-pui-timepicker-period');
      let newHour = state.hour;

      if (period === 'AM' && state.hour >= 12) {
        newHour = state.hour === 12 ? 0 : state.hour - 12;
      } else if (period === 'PM' && state.hour < 12) {
        newHour = state.hour === 0 ? 12 : state.hour + 12;
      }

      if (newHour !== state.hour) {
        if (!isValidTime(newHour, state.minute, state.minTime, state.maxTime)) {
          // Find first valid minute
          for (let m = 0; m < 60; m += state.step) {
            if (isValidTime(newHour, m, state.minTime, state.maxTime)) {
              setState(trigger, newHour, m);
              return;
            }
          }
        } else {
          setState(trigger, newHour, state.minute);
        }
      }
      return;
    }

    // Done button
    if (target.matches('[data-pui-timepicker-done]')) {
      const trigger = findTrigger(target);
      if (trigger && window.closePopover) {
        window.closePopover(trigger.id + '-content');
      }
      return;
    }
  });

  // Form reset
  document.addEventListener('reset', (e) => {
    if (!e.target.matches('form')) return;

    e.target.querySelectorAll('[data-pui-timepicker="true"]').forEach(trigger => {
      setState(trigger, null, null);
      const elements = getElements(trigger);
      if (elements && elements.hiddenInput) {
        elements.hiddenInput.value = '';
      }
    });
  });

  // MutationObserver for initial rendering
  const observer = new MutationObserver(() => {
    document.querySelectorAll('[data-pui-timepicker="true"]:not([data-rendered])').forEach(trigger => {
      trigger.setAttribute('data-rendered', 'true');

      // Read initial value from hidden input
      const elements = getElements(trigger);
      const initialValue = (elements && elements.hiddenInput && elements.hiddenInput.value) ||
                          (elements && elements.popup && elements.popup.getAttribute('data-pui-timepicker-value'));

      if (initialValue) {
        const parsed = parseTime(initialValue);
        if (parsed) {
          setState(trigger, parsed.hour, parsed.minute);
        }
      }

      updateDisplay(trigger);
    });
  });

  observer.observe(document.body, { childList: true, subtree: true });
})();`
