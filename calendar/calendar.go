package calendar

import (
	"crypto/rand"
	_ "embed"
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

//go:embed calendar.js
var calendarJS string
