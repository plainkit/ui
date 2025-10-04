package calendar

import (
	"crypto/rand"
	_ "embed"
	"encoding/hex"
	"strconv"
	"time"

	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/internal/styles"
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

func divArgsFromProps(baseClass string, extra ...string) func(p Props) []html.DivArg {
	return func(p Props) []html.DivArg {
		args := []html.DivArg{html.AClass(html.ClassMerge(append([]string{baseClass}, append(extra, p.Class)...)...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID+"-wrapper"))
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
		id = randomID("calendar")
	}

	name := p.Name
	if name == "" {
		name = id + "-value"
	}

	localeTag := p.LocaleTag
	if localeTag == "" {
		localeTag = LocaleDefaultTag
	}

	initialStartOfWeek := Monday
	if p.StartOfWeek != nil {
		initialStartOfWeek = *p.StartOfWeek
	}

	initialView := time.Now()
	if p.Value != nil {
		initialView = *p.Value
	}

	initialMonth := p.InitialMonth
	initialYear := p.InitialYear

	// Use year from initialView if InitialYear prop is invalid/unset (<= 0)
	if initialYear <= 0 {
		initialYear = initialView.Year()
	}

	// Use month from initialView if InitialMonth prop is invalid OR
	// if InitialMonth is default 0 AND InitialYear was also defaulted (meaning neither was likely set explicitly)
	if (initialMonth < 0 || initialMonth > 11) || (initialMonth == 0 && p.InitialYear <= 0) {
		initialMonth = int(initialView.Month()) - 1 // time.Month is 1-12
	}

	initialSelectedISO := ""
	if p.Value != nil {
		initialSelectedISO = p.Value.Format("2006-01-02")
	}

	args := divArgsFromProps(styles.Surface("inline-flex flex-col gap-4 p-6"))(p)
	args = append(args, html.AData("pui-calendar-wrapper", "true"))

	// Build wrapper content
	wrapperContent := make([]html.Component, 0)

	// Add hidden input if requested (default: true)
	if p.RenderHiddenInput {
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
			html.AClass("mb-4 flex items-center justify-between gap-3 rounded-2xl bg-muted/50 px-4 py-3"),
			html.Span(
				html.AData("pui-calendar-month-display", ""),
				html.AClass(styles.DisplayHeading("text-lg")),
			),
			html.Div(
				html.AClass("flex items-center gap-2"),
				html.Button(
					html.AType("button"),
					html.AData("pui-calendar-prev", ""),
					html.AClass(styles.InteractiveGhost("size-8 rounded-full text-muted-foreground hover:text-foreground")),
					lucide.ChevronLeft(html.AClass("h-4 w-4")),
				),
				html.Button(
					html.AType("button"),
					html.AData("pui-calendar-next", ""),
					html.AClass(styles.InteractiveGhost("size-8 rounded-full text-muted-foreground hover:text-foreground")),
					lucide.ChevronRight(html.AClass("h-4 w-4")),
				),
			),
		),

		// Weekday Headers
		html.Div(
			html.AData("pui-calendar-weekdays", ""),
			html.AClass("mb-1 grid grid-cols-7 place-items-center gap-2 text-[0.7rem] font-semibold uppercase tracking-[0.3em] text-muted-foreground/70"),
		),

		// Calendar Day Grid
		html.Div(
			html.AData("pui-calendar-days", ""),
			html.AClass("grid grid-cols-7 gap-2"),
		),
	)

	wrapperContent = append(wrapperContent, calendarContainer)
	*children = append(*children, wrapperContent...)

	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

// Calendar renders a calendar component for date selection
func Calendar(args ...html.DivArg) html.Node {
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

	return html.Div(append([]html.DivArg{props}, rest...)...).WithAssets("", calendarJS, "ui-calendar")
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
