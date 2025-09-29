package calendar

import (
	"crypto/rand"
	"encoding/hex"
	"strconv"
	"time"

	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/internal/classnames"
)

type LocaleTag string

const (
	LocaleDefaultTag    LocaleTag = "en-US"
	LocaleTagChinese    LocaleTag = "zh-CN"
	LocaleTagFrench     LocaleTag = "fr-FR"
	LocaleTagGerman     LocaleTag = "de-DE"
	LocaleTagItalian    LocaleTag = "it-IT"
	LocaleTagJapanese   LocaleTag = "ja-JP"
	LocaleTagPortuguese LocaleTag = "pt-PT"
	LocaleTagSpanish    LocaleTag = "es-ES"
)

type Day int

const (
	Sunday    Day = 0
	Monday    Day = 1
	Tuesday   Day = 2
	Wednesday Day = 3
	Thursday  Day = 4
	Friday    Day = 5
	Saturday  Day = 6
)

type Props struct {
	ID                string
	Class             string
	Attrs             []html.Global
	LocaleTag         LocaleTag
	Value             *time.Time
	Name              string
	InitialMonth      int  // Optional: 0-11 (Default: current or from Value). Controls the initially displayed month view.
	InitialYear       int  // Optional: (Default: current or from Value). Controls the initially displayed year view.
	StartOfWeek       *Day // Optional: 0-6 [Sun-Sat] (Default: 1).
	RenderHiddenInput bool // Optional: Whether to render the hidden input (Default: true). Set to false when used inside DatePicker.
}

// Calendar renders a calendar component for date selection
func Calendar(props Props, args ...html.DivArg) html.Node {
	id := props.ID
	if id == "" {
		id = randomID("calendar")
	}

	name := props.Name
	if name == "" {
		name = id + "-value"
	}

	localeTag := props.LocaleTag
	if localeTag == "" {
		localeTag = LocaleDefaultTag
	}

	initialStartOfWeek := Monday
	if props.StartOfWeek != nil {
		initialStartOfWeek = *props.StartOfWeek
	}

	initialView := time.Now()
	if props.Value != nil {
		initialView = *props.Value
	}

	initialMonth := props.InitialMonth
	initialYear := props.InitialYear

	// Use year from initialView if InitialYear prop is invalid/unset (<= 0)
	if initialYear <= 0 {
		initialYear = initialView.Year()
	}

	// Use month from initialView if InitialMonth prop is invalid OR
	// if InitialMonth is default 0 AND InitialYear was also defaulted (meaning neither was likely set explicitly)
	if (initialMonth < 0 || initialMonth > 11) || (initialMonth == 0 && props.InitialYear <= 0) {
		initialMonth = int(initialView.Month()) - 1 // time.Month is 1-12
	}

	initialSelectedISO := ""
	if props.Value != nil {
		initialSelectedISO = props.Value.Format("2006-01-02")
	}

	wrapperArgs := []html.DivArg{
		html.AClass(classnames.Merge("", props.Class)),
		html.AId(id + "-wrapper"),
		html.AData("pui-calendar-wrapper", "true"),
	}

	for _, attr := range props.Attrs {
		wrapperArgs = append(wrapperArgs, attr)
	}

	// Build wrapper content
	wrapperContent := make([]html.DivArg, 0)

	// Add hidden input if requested (default: true)
	if props.RenderHiddenInput {
		hiddenInput := html.Input(
			html.AType("hidden"),
			html.AName(name),
			html.AValue(initialSelectedISO),
			html.AId(id+"-hidden"),
			html.AData("pui-calendar-hidden-input", ""),
		)
		wrapperContent = append(wrapperContent, hiddenInput)
	}

	// Calendar container
	calendarContainer := html.Div(
		html.AId(id),
		html.AData("pui-calendar-container", "true"),
		html.AData("pui-calendar-locale-tag", string(localeTag)),
		html.AData("pui-calendar-initial-month", strconv.Itoa(initialMonth)),
		html.AData("pui-calendar-initial-year", strconv.Itoa(initialYear)),
		html.AData("pui-calendar-selected-date", initialSelectedISO),
		html.AData("pui-calendar-start-of-week", strconv.Itoa(int(initialStartOfWeek))),

		// Calendar Header
		html.Div(
			html.AClass("flex items-center justify-between mb-4"),
			html.Span(
				html.AData("pui-calendar-month-display", ""),
				html.AClass("text-sm font-medium"),
			),
			html.Div(
				html.AClass("flex gap-1"),
				html.Button(
					html.AType("button"),
					html.AData("pui-calendar-prev", ""),
					html.AClass("inline-flex items-center justify-center rounded-md text-sm font-medium h-7 w-7 hover:bg-accent hover:text-accent-foreground focus:outline-none disabled:opacity-50"),
					lucide.ChevronLeft(html.AClass("h-4 w-4")),
				),
				html.Button(
					html.AType("button"),
					html.AData("pui-calendar-next", ""),
					html.AClass("inline-flex items-center justify-center rounded-md text-sm font-medium h-7 w-7 hover:bg-accent hover:text-accent-foreground focus:outline-none disabled:opacity-50"),
					lucide.ChevronRight(html.AClass("h-4 w-4")),
				),
			),
		),

		// Weekday Headers
		html.Div(
			html.AData("pui-calendar-weekdays", ""),
			html.AClass("grid grid-cols-7 gap-1 mb-1 place-items-center"),
		),

		// Calendar Day Grid
		html.Div(
			html.AData("pui-calendar-days", ""),
			html.AClass("grid grid-cols-7 gap-1 place-items-center"),
		),
	)

	wrapperContent = append(wrapperContent, calendarContainer)
	wrapperContent = append(wrapperContent, args...)
	wrapperArgs = append(wrapperArgs, wrapperContent...)

	return html.Div(wrapperArgs...).WithAssets("", calendarJS, "ui-calendar")
}

func randomID(prefix string) string {
	buf := make([]byte, 4)
	if _, err := rand.Read(buf); err != nil {
		return prefix + "-id"
	}
	return prefix + "-" + hex.EncodeToString(buf)
}

const calendarJS = `(function () {
  "use strict";

  // Utility functions
  function parseISODate(isoStr) {
    if (!isoStr) return null;
    const parts = isoStr.split("-");
    if (parts.length !== 3) return null;

    const year = parseInt(parts[0], 10);
    const month = parseInt(parts[1], 10) - 1;
    const day = parseInt(parts[2], 10);
    const date = new Date(Date.UTC(year, month, day));

    if (
      date.getUTCFullYear() === year &&
      date.getUTCMonth() === month &&
      date.getUTCDate() === day
    ) {
      return date;
    }
    return null;
  }

  function getMonthNames(locale) {
    try {
      return Array.from({ length: 12 }, (_, i) =>
        new Intl.DateTimeFormat(locale, {
          month: "long",
          timeZone: "UTC",
        }).format(new Date(Date.UTC(2000, i, 1))),
      );
    } catch {
      return [
        "January",
        "February",
        "March",
        "April",
        "May",
        "June",
        "July",
        "August",
        "September",
        "October",
        "November",
        "December",
      ];
    }
  }

  function getDayNames(locale, startOfWeek) {
    try {
      return Array.from({ length: 7 }, (_, i) =>
        new Intl.DateTimeFormat(locale, {
          weekday: "short",
          timeZone: "UTC",
        }).format(new Date(Date.UTC(2000, 0, i + 2 + startOfWeek))),
      );
    } catch {
      return ["Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"];
    }
  }

  function findHiddenInput(container) {
    // Check wrapper first
    const wrapper = container.closest("[data-pui-calendar-wrapper]");
    let hiddenInput = wrapper?.querySelector(
      "[data-pui-calendar-hidden-input]",
    );

    // For datepicker integration
    if (!hiddenInput && container.id) {
      const parentId = container.id.replace("-calendar-instance", "");
      hiddenInput = document.getElementById(parentId + "-hidden");
    }

    return hiddenInput;
  }

  function renderCalendar(container) {
    const monthDisplay = container.querySelector(
      "[data-pui-calendar-month-display]",
    );
    const weekdaysContainer = container.querySelector(
      "[data-pui-calendar-weekdays]",
    );
    const daysContainer = container.querySelector("[data-pui-calendar-days]");

    if (!monthDisplay || !weekdaysContainer || !daysContainer) return;

    // Get current viewing month/year (or use initial/defaults)
    let currentMonth = parseInt(container.dataset.tuiCalendarCurrentMonth);
    let currentYear = parseInt(container.dataset.tuiCalendarCurrentYear);

    // If not set, use initial values or current date
    if (isNaN(currentMonth) || isNaN(currentYear)) {
      const initialMonth = parseInt(
        container.getAttribute("data-pui-calendar-initial-month"),
      );
      const initialYear = parseInt(
        container.getAttribute("data-pui-calendar-initial-year"),
      );
      const selectedDate = container.getAttribute(
        "data-pui-calendar-selected-date",
      );

      if (selectedDate) {
        const parsed = parseISODate(selectedDate);
        if (parsed) {
          currentMonth = parsed.getUTCMonth();
          currentYear = parsed.getUTCFullYear();
        }
      }

      if (isNaN(currentMonth)) {
        currentMonth = !isNaN(initialMonth)
          ? initialMonth
          : new Date().getMonth();
      }
      if (isNaN(currentYear)) {
        currentYear =
          !isNaN(initialYear) && initialYear > 0
            ? initialYear
            : new Date().getFullYear();
      }

      // Store for navigation
      container.dataset.tuiCalendarCurrentMonth = currentMonth;
      container.dataset.tuiCalendarCurrentYear = currentYear;
    }

    // Get other settings
    const locale =
      container.getAttribute("data-pui-calendar-locale-tag") || "en-US";
    const startOfWeek =
      parseInt(container.getAttribute("data-pui-calendar-start-of-week")) || 1;
    const selectedDateStr = container.getAttribute(
      "data-pui-calendar-selected-date",
    );
    const selectedDate = selectedDateStr ? parseISODate(selectedDateStr) : null;

    // Update month display
    const monthNames = getMonthNames(locale);
    monthDisplay.textContent = monthNames[currentMonth] + " " + currentYear;

    // Render weekdays if empty
    if (!weekdaysContainer.children.length) {
      const dayNames = getDayNames(locale, startOfWeek);
      weekdaysContainer.innerHTML = dayNames
        .map(
          (day) =>
            '<div class="text-center text-xs text-muted-foreground font-medium">' + day + '</div>',
        )
        .join("");
    }

    // Render days
    daysContainer.innerHTML = "";

    const firstDay = new Date(Date.UTC(currentYear, currentMonth, 1));
    const startOffset = (((firstDay.getUTCDay() - startOfWeek) % 7) + 7) % 7;
    const daysInMonth = new Date(
      Date.UTC(currentYear, currentMonth + 1, 0),
    ).getUTCDate();

    const today = new Date();
    const todayUTC = new Date(
      Date.UTC(today.getFullYear(), today.getMonth(), today.getDate()),
    );

    // Add empty cells for offset
    for (let i = 0; i < startOffset; i++) {
      daysContainer.innerHTML += '<div class="h-8 w-8"></div>';
    }

    // Add day buttons
    for (let day = 1; day <= daysInMonth; day++) {
      const currentDate = new Date(Date.UTC(currentYear, currentMonth, day));
      const isSelected =
        selectedDate && currentDate.getTime() === selectedDate.getTime();
      const isToday = currentDate.getTime() === todayUTC.getTime();

      let classes =
        "inline-flex h-8 w-8 items-center justify-center rounded-md text-sm font-medium focus:outline-none focus:ring-1 focus:ring-ring";

      if (isSelected) {
        classes += " bg-primary text-primary-foreground hover:bg-primary/90";
      } else if (isToday) {
        classes += " bg-accent text-accent-foreground";
      } else {
        classes += " hover:bg-accent hover:text-accent-foreground";
      }

      daysContainer.innerHTML += '<button type="button" class="' + classes + '" data-pui-calendar-day="' + day + '">' + day + '</button>';
    }
  }

  // Event delegation for calendar navigation and selection
  document.addEventListener("click", (e) => {
    // Previous month
    const prevBtn = e.target.closest("[data-pui-calendar-prev]");
    if (prevBtn) {
      const container = prevBtn.closest("[data-pui-calendar-container]");
      if (!container) return;

      let month = parseInt(container.dataset.tuiCalendarCurrentMonth, 10);
      let year = parseInt(container.dataset.tuiCalendarCurrentYear, 10);

      // Only use fallback if truly not initialized (should not happen after init)
      if (isNaN(month)) month = new Date().getMonth();
      if (isNaN(year)) year = new Date().getFullYear();

      month--;
      if (month < 0) {
        month = 11;
        year--;
      }

      container.dataset.tuiCalendarCurrentMonth = month;
      container.dataset.tuiCalendarCurrentYear = year;
      renderCalendar(container);
      return;
    }

    // Next month
    const nextBtn = e.target.closest("[data-pui-calendar-next]");
    if (nextBtn) {
      const container = nextBtn.closest("[data-pui-calendar-container]");
      if (!container) return;

      let month = parseInt(container.dataset.tuiCalendarCurrentMonth, 10);
      let year = parseInt(container.dataset.tuiCalendarCurrentYear, 10);

      // Only use fallback if truly not initialized (should not happen after init)
      if (isNaN(month)) month = new Date().getMonth();
      if (isNaN(year)) year = new Date().getFullYear();

      month++;
      if (month > 11) {
        month = 0;
        year++;
      }

      container.dataset.tuiCalendarCurrentMonth = month;
      container.dataset.tuiCalendarCurrentYear = year;
      renderCalendar(container);
      return;
    }

    // Day selection
    if (e.target.matches("[data-pui-calendar-day]")) {
      const container = e.target.closest("[data-pui-calendar-container]");
      if (!container) return;

      const day = parseInt(e.target.dataset.tuiCalendarDay);
      let month = parseInt(container.dataset.tuiCalendarCurrentMonth, 10);
      let year = parseInt(container.dataset.tuiCalendarCurrentYear, 10);

      // Only use fallback if truly not initialized (should not happen after init)
      if (isNaN(month)) month = new Date().getMonth();
      if (isNaN(year)) year = new Date().getFullYear();
      const selectedDate = new Date(Date.UTC(year, month, day));

      // Update selected date attribute
      container.setAttribute(
        "data-pui-calendar-selected-date",
        selectedDate.toISOString().split("T")[0],
      );

      // Update hidden input
      const hiddenInput = findHiddenInput(container);
      if (hiddenInput) {
        hiddenInput.value = selectedDate.toISOString().split("T")[0];
        hiddenInput.dispatchEvent(new Event("change", { bubbles: true }));
      }

      // Dispatch custom event
      container.dispatchEvent(
        new CustomEvent("calendar-date-selected", {
          bubbles: true,
          detail: { date: selectedDate },
        }),
      );

      renderCalendar(container);
    }
  });

  // Form reset handling
  document.addEventListener("reset", (e) => {
    if (!e.target.matches("form")) return;

    e.target
      .querySelectorAll("[data-pui-calendar-container]")
      .forEach((container) => {
        const hiddenInput = findHiddenInput(container);
        if (hiddenInput) {
          hiddenInput.value = "";
        }

        // Clear selected date and reset to current month
        container.removeAttribute("data-pui-calendar-selected-date");
        const today = new Date();
        container.dataset.tuiCalendarCurrentMonth = today.getMonth();
        container.dataset.tuiCalendarCurrentYear = today.getFullYear();
        renderCalendar(container);
      });
  });

  // MutationObserver for dynamic content (framework-agnostic)
  const observer = new MutationObserver(() => {
    document
      .querySelectorAll("[data-pui-calendar-container]")
      .forEach((container) => {
        const daysContainer = container.querySelector(
          "[data-pui-calendar-days]",
        );
        // Only render if not already rendered
        if (daysContainer && !daysContainer.children.length) {
          renderCalendar(container);
        }
      });
  });

  observer.observe(document.body, { childList: true, subtree: true });

  // Initialize calendars on page load
  function initCalendars() {
    document
      .querySelectorAll("[data-pui-calendar-container]")
      .forEach(renderCalendar);
  }

  if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", initCalendars);
  } else {
    initCalendars();
  }
})();`
