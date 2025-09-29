package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/accordion"
	"github.com/plainkit/ui/alert"
	"github.com/plainkit/ui/aspectratio"
	"github.com/plainkit/ui/avatar"
	"github.com/plainkit/ui/badge"
	"github.com/plainkit/ui/breadcrumb"
	"github.com/plainkit/ui/button"
	"github.com/plainkit/ui/calendar"
	"github.com/plainkit/ui/card"
	"github.com/plainkit/ui/carousel"
	"github.com/plainkit/ui/checkbox"
	democss "github.com/plainkit/ui/cmd/demo/internal/css"
	"github.com/plainkit/ui/code"
	"github.com/plainkit/ui/collapsible"
	"github.com/plainkit/ui/dialog"
	"github.com/plainkit/ui/dropdown"
	"github.com/plainkit/ui/form"
	"github.com/plainkit/ui/input"
	"github.com/plainkit/ui/inputotp"
	"github.com/plainkit/ui/label"
	"github.com/plainkit/ui/pagination"
	"github.com/plainkit/ui/popover"
	"github.com/plainkit/ui/progress"
	"github.com/plainkit/ui/radio"
	"github.com/plainkit/ui/rating"
	"github.com/plainkit/ui/selectbox"
	"github.com/plainkit/ui/separator"
	"github.com/plainkit/ui/skeleton"
	"github.com/plainkit/ui/slider"
	switchcomp "github.com/plainkit/ui/switch"
	"github.com/plainkit/ui/table"
	"github.com/plainkit/ui/tabs"
	"github.com/plainkit/ui/tagsinput"
	"github.com/plainkit/ui/textarea"
	"github.com/plainkit/ui/timepicker"
	"github.com/plainkit/ui/toast"
	"github.com/plainkit/ui/tooltip"
)

type page struct {
	Path    string
	Label   string
	Content func() html.Node
}

var pages = []page{
	{Path: "/accordion", Label: "Accordion", Content: renderAccordionContent},
	{Path: "/alerts", Label: "Alerts", Content: renderAlertsContent},
	{Path: "/aspect-ratios", Label: "Aspect Ratios", Content: renderAspectRatiosContent},
	{Path: "/avatars", Label: "Avatars", Content: renderAvatarsContent},
	{Path: "/badges", Label: "Badges", Content: renderBadgesContent},
	{Path: "/breadcrumbs", Label: "Breadcrumbs", Content: renderBreadcrumbsContent},
	{Path: "/buttons", Label: "Buttons", Content: renderButtonsContent},
	{Path: "/calendar", Label: "Calendar", Content: renderCalendarContent},
	{Path: "/cards", Label: "Cards", Content: renderCardsContent},
	{Path: "/carousel", Label: "Carousel", Content: renderCarouselContent},
	{Path: "/checkboxes", Label: "Checkboxes", Content: renderCheckboxesContent},
	{Path: "/code", Label: "Code", Content: renderCodeContent},
	{Path: "/collapsible", Label: "Collapsible", Content: renderCollapsibleContent},
	{Path: "/dialogs", Label: "Dialogs", Content: renderDialogsContent},
	{Path: "/dropdowns", Label: "Dropdowns", Content: renderDropdownsContent},
	{Path: "/forms", Label: "Form Helpers", Content: renderFormContent},
	{Path: "/input-otp", Label: "Input OTP", Content: renderInputOTPContent},
	{Path: "/inputs", Label: "Inputs", Content: renderInputsContent},
	{Path: "/labels", Label: "Labels", Content: renderLabelsContent},
	{Path: "/pagination", Label: "Pagination", Content: renderPaginationContent},
	{Path: "/popovers", Label: "Popovers", Content: renderPopoversContent},
	{Path: "/progress", Label: "Progress", Content: renderProgressContent},
	{Path: "/radios", Label: "Radios", Content: renderRadiosContent},
	{Path: "/ratings", Label: "Ratings", Content: renderRatingsContent},
	{Path: "/selectboxes", Label: "Select Boxes", Content: renderSelectBoxesContent},
	{Path: "/separators", Label: "Separators", Content: renderSeparatorsContent},
	{Path: "/skeletons", Label: "Skeletons", Content: renderSkeletonsContent},
	{Path: "/sliders", Label: "Sliders", Content: renderSlidersContent},
	{Path: "/switches", Label: "Switches", Content: renderSwitchesContent},
	{Path: "/tables", Label: "Tables", Content: renderTablesContent},
	{Path: "/tabs", Label: "Tabs", Content: renderTabsContent},
	{Path: "/tags-input", Label: "Tags Input", Content: renderTagsInputContent},
	{Path: "/textareas", Label: "Textareas", Content: renderTextareasContent},
	{Path: "/timepicker", Label: "Time Picker", Content: renderTimePickerContent},
	{Path: "/toasts", Label: "Toasts", Content: renderToastsContent},
	{Path: "/tooltips", Label: "Tooltips", Content: renderTooltipsContent},
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/assets/styles.css", cssHandler)

	for _, pg := range pages {
		p := pg
		mux.HandleFunc(p.Path, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			body := p.Content()
			if _, err := w.Write([]byte(renderPage(p.Path, body))); err != nil {
				log.Printf("write response: %v", err)
			}
		})
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, pages[0].Path, http.StatusFound)
	})

	addr := ":8080"
	log.Printf("UI components demo available at http://localhost%v", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

func cssHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=31536000")
	if _, err := w.Write([]byte(democss.TailwindCSS)); err != nil {
		log.Printf("write css: %v", err)
	}
}

func renderPage(activePath string, body html.Node) string {
	assets := html.NewAssets()
	assets.Collect(body)

	headChildren := []html.HeadArg{
		html.Child(html.Meta(html.ACharset("utf-8"))),
		html.Child(html.Meta(html.AName("viewport"), html.AContent("width=device-width, initial-scale=1"))),
		html.Child(html.Title(html.Text("Plain UI Components"))),
		html.Child(html.Link(html.ARel("stylesheet"), html.AHref("/assets/styles.css"))),
	}

	bodyChildren := []html.Component{
		html.Div(
			html.AClass("flex min-h-screen"),
			renderSidebar(activePath),
			html.Main(
				html.AClass("flex-1 overflow-y-auto"),
				html.Div(
					html.AClass("mx-auto w-full max-w-4xl px-8 py-12 space-y-12"),
					body,
				),
			),
		),
	}

	if cssSnippets := assets.CSS(); len(cssSnippets) > 0 {
		headChildren = append(headChildren, html.Child(html.Style(html.UnsafeText(strings.Join(cssSnippets, "\n")))))
	}

	if jsSnippets := assets.JS(); len(jsSnippets) > 0 {
		bodyChildren = append([]html.Component{html.Script(html.UnsafeText(strings.Join(jsSnippets, "\n")))}, bodyChildren...)
	}

	bodyArgs := []html.BodyArg{html.AClass("min-h-screen bg-background text-foreground")}
	for _, child := range bodyChildren {
		bodyArgs = append(bodyArgs, html.Child(child))
	}

	page := html.Html(
		html.Head(headChildren...),
		html.Body(bodyArgs...),
	)

	return "<!DOCTYPE html>\n" + html.Render(page)
}

func renderSidebar(activePath string) html.Node {
	links := make([]html.UlArg, 0, len(pages))
	for _, pg := range pages {
		isActive := pg.Path == activePath
		className := "block rounded-md px-3 py-2 text-sm font-medium transition-colors"
		if isActive {
			className += " bg-sidebar-accent text-sidebar-accent-foreground"
		} else {
			className += " text-sidebar-foreground/70 hover:bg-sidebar-accent hover:text-sidebar-accent-foreground"
		}

		links = append(links,
			html.Li(
				html.A(
					html.AHref(pg.Path),
					html.AClass(className),
					html.Text(pg.Label),
				),
			),
		)
	}

	return html.Aside(
		html.AClass("w-64 border-r border-sidebar-border bg-sidebar"),
		html.Div(
			html.AClass("sticky top-0 flex h-screen flex-col gap-6 p-6"),
			html.Div(
				html.AClass("flex items-center gap-2 py-2"),
				lucide.Layers(html.AClass("size-6 text-primary")),
				html.H1(
					html.AClass("text-xl font-semibold text-sidebar-foreground"),
					html.Text("Plain UI"),
				),
			),
			html.Nav(
				html.AClass("flex-1 overflow-y-auto scrollbar-thin scrollbar-thumb-sidebar-border/30 hover:scrollbar-thumb-sidebar-border/50"),
				html.AStyle("scrollbar-width: thin; scrollbar-color: rgba(0,0,0,0.1) transparent;"),
				html.Ul(append([]html.UlArg{html.AClass("space-y-1 pb-6 pr-2")}, links...)...),
			),
		),
	)
}

func renderButtonsContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Variants")),
				html.P(html.AClass("text-sm text-slate-400"), html.Text("Standard button appearances using design tokens.")),
			),
			html.Div(
				html.AClass("flex flex-wrap gap-3"),
				button.Button(button.Props{}, html.T("Default")),
				button.Button(button.Props{Variant: button.VariantSecondary}, html.T("Secondary")),
				button.Button(button.Props{Variant: button.VariantOutline}, html.T("Outline")),
				button.Button(button.Props{Variant: button.VariantDestructive}, html.T("Destructive")),
				button.Button(button.Props{Variant: button.VariantGhost}, html.T("Ghost")),
				button.Button(
					button.Props{Variant: button.VariantLink, Href: "https://plainkit.dev", Target: "_blank"},
					html.T("Link"),
				),
			),
		),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Sizes")),
				html.P(html.AClass("text-sm text-slate-400"), html.Text("Adjust sizing tokens for different contexts.")),
			),
			html.Div(
				html.AClass("flex flex-wrap items-end gap-3"),
				button.Button(button.Props{}, html.T("Default")),
				button.Button(button.Props{Size: button.SizeSm}, html.T("Small")),
				button.Button(button.Props{Size: button.SizeLg}, html.T("Large")),
				button.Button(button.Props{Size: button.SizeIcon}, html.T("+")),
			),
		),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("States")),
				html.P(html.AClass("text-sm text-slate-400"), html.Text("Examples of disabled and full-width buttons.")),
			),
			html.Div(
				html.AClass("flex flex-wrap gap-3"),
				button.Button(button.Props{Disabled: true}, html.T("Disabled")),
				button.Button(button.Props{FullWidth: true}, html.T("Full width")),
				button.Button(button.Props{Href: "https://plainkit.dev/docs"}, html.T("As link")),
			),
		),
	)
}

func renderBadgesContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Variants")),
				html.P(html.AClass("text-sm text-slate-400"), html.Text("Status indicators for different contexts.")),
			),
			html.Div(
				html.AClass("flex flex-wrap gap-3"),
				badge.Badge(badge.Props{}, html.T("Default")),
				badge.Badge(badge.Props{Variant: badge.VariantSecondary}, html.T("Secondary")),
				badge.Badge(badge.Props{Variant: badge.VariantOutline}, html.T("Outline")),
				badge.Badge(badge.Props{Variant: badge.VariantDestructive}, html.T("Destructive")),
			),
		),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Custom content")),
				html.P(html.AClass("text-sm text-slate-400"), html.Text("Badges can contain arbitrary content.")),
			),
			html.Div(
				html.AClass("flex flex-wrap gap-3"),
				badge.Badge(badge.Props{Class: "gap-1.5"}, html.T("v2.0")),
				badge.Badge(
					badge.Props{Variant: badge.VariantSecondary},
					html.Span(
						html.AClass("inline-flex items-center gap-1"),
						html.Text("Active"),
					),
				),
			),
		),
	)
}

func renderCheckboxesContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Checkboxes")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Use checkboxes to represent binary options and multi-select choices.")),
			),
			html.Div(
				html.AClass("space-y-3"),
				html.Label(
					html.AClass("flex items-center gap-2"),
					html.Child(checkbox.Checkbox(checkbox.Props{ID: "cb-updates", Name: "updates"}, nil)),
					html.Text("Email me product updates"),
				),
				html.Label(
					html.AClass("flex items-center gap-2"),
					html.Child(checkbox.Checkbox(checkbox.Props{ID: "cb-terms", Name: "terms", Required: true}, nil)),
					html.Text("I agree to the terms"),
				),
				html.Label(
					html.AClass("flex items-center gap-2 text-muted-foreground"),
					html.Child(checkbox.Checkbox(checkbox.Props{ID: "cb-disabled", Disabled: true}, nil)),
					html.Text("Remember my choice (unavailable)"),
				),
			),
		),
	)
}

func renderSwitchesContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Switches")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Switches express an on/off state for instant changes.")),
			),
			html.Div(
				html.AClass("space-y-3"),
				switchcomp.Switch(switchcomp.Props{ID: "switch-email", Name: "email"}, html.Child(html.Span(html.Text("Email notifications")))),
				switchcomp.Switch(switchcomp.Props{ID: "switch-push", Name: "push", Checked: true}, html.Child(html.Span(html.Text("Push notifications")))),
				switchcomp.Switch(switchcomp.Props{ID: "switch-disabled", Disabled: true}, html.Child(html.Span(html.Text("Public profile")))),
			),
		),
	)
}

func renderSlidersContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Sliders")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Adjust numeric values within a range.")),
			),
			html.Div(
				html.AClass("space-y-6"),
				slider.Slider(
					slider.Props{},
					slider.Input(slider.InputProps{ID: "volume", Name: "volume", Min: 0, Max: 100, Value: 50}),
					html.Div(
						html.AClass("flex items-center justify-between text-sm"),
						html.Span(html.Text("Volume")),
						slider.Value(slider.ValueProps{For: "volume"}),
					),
				),
				slider.Slider(
					slider.Props{},
					slider.Input(slider.InputProps{ID: "brightness", Name: "brightness", Min: 20, Max: 120, Value: 90, Step: 5}),
					html.Div(
						html.AClass("flex items-center justify-between text-sm"),
						html.Span(html.Text("Brightness")),
						slider.Value(slider.ValueProps{For: "brightness"}),
					),
				),
			),
		),
	)
}

func renderTextareasContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Textareas")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Collect multi-line input with optional auto-resize.")),
			),
			html.Div(
				html.AClass("grid gap-6 md:grid-cols-2"),
				html.Div(
					html.AClass("space-y-2"),
					label.Label(label.Props{For: "bio"}, html.Text("Short bio")),
					textarea.Textarea(textarea.Props{ID: "bio", Name: "bio", Placeholder: "Tell us about yourself", Rows: 4}),
				),
				html.Div(
					html.AClass("space-y-2"),
					label.Label(label.Props{For: "notes"}, html.Text("Meeting notes")),
					textarea.Textarea(textarea.Props{ID: "notes", Name: "notes", AutoResize: true, Placeholder: "Notes will auto-resize as you type."}),
				),
			),
		),
	)
}

func renderTabsContent() html.Node {
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

func renderInputsContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Form inputs")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Text inputs adapt to validation, states, and password toggles.")),
			),
			html.Div(
				html.AClass("grid gap-6 md:grid-cols-2"),
				html.Div(
					html.AClass("space-y-2"),
					label.Label(label.Props{For: "name"}, html.Text("Full name")),
					input.Input(input.Props{ID: "name", Name: "name", Placeholder: "Ada Lovelace"}),
				),
				html.Div(
					html.AClass("space-y-2"),
					label.Label(label.Props{For: "email"}, html.Text("Email address")),
					input.Input(input.Props{ID: "email", Name: "email", Type: input.TypeEmail, Placeholder: "name@example.com"}),
				),
				html.Div(
					html.AClass("space-y-2"),
					label.Label(label.Props{For: "password"}, html.Text("Password")),
					input.Input(input.Props{ID: "password", Name: "password", Type: input.TypePassword, Placeholder: "••••••••", ShowPasswordToggle: true}),
				),
				html.Div(
					html.AClass("space-y-2"),
					label.Label(label.Props{For: "search", Error: "No results"}, html.Text("Search")),
					input.Input(input.Props{ID: "search", Name: "search", Placeholder: "Search docs", HasError: true}),
				),
			),
		),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("File & disabled inputs")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Provide metadata for file uploads and communicate disabled states.")),
			),
			html.Div(
				html.AClass("grid gap-6 md:grid-cols-2"),
				html.Div(
					html.AClass("space-y-2"),
					label.Label(label.Props{For: "upload"}, html.Text("Project assets")),
					input.Input(input.Props{ID: "upload", Name: "assets", Type: input.TypeFile, FileAccept: "image/*"}),
				),
				html.Div(
					html.AClass("space-y-2"),
					label.Label(label.Props{For: "company"}, html.Text("Company")),
					input.Input(input.Props{ID: "company", Name: "company", Placeholder: "Acme Inc.", Disabled: true}),
				),
			),
		),
	)
}

func renderAlertsContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Contextual alerts")),
				html.P(html.AClass("text-sm text-slate-400"), html.Text("Communicate high priority feedback and destructive states.")),
			),
			html.Div(
				html.AClass("space-y-3"),
				alert.Alert(
					alert.Props{},
					alert.Title(alert.TitleProps{}, html.Text("Heads up!")),
					alert.Description(alert.DescriptionProps{}, html.P(html.Text("You can add components to your project from the left sidebar."))),
				),
				alert.Alert(
					alert.Props{Variant: alert.VariantDestructive},
					alert.Title(alert.TitleProps{}, html.Text("Action required")),
					alert.Description(alert.DescriptionProps{}, html.P(html.Text("This change is permanent and cannot be undone."))),
				),
			),
		),
	)
}

func renderProgressContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Upload progress")),
				html.P(html.AClass("text-sm text-slate-400"), html.Text("Display progress with optional labels and value output.")),
			),
			html.Div(
				html.AClass("space-y-4"),
				progress.Progress(progress.Props{Label: "Preparing files", Value: 30, ShowValue: true}),
				progress.Progress(progress.Props{Label: "Uploading", Value: 65, ShowValue: true, Variant: progress.VariantSuccess}),
				progress.Progress(progress.Props{Label: "Processing", Value: 45, ShowValue: true, Variant: progress.VariantWarning, Size: progress.SizeLg}),
			),
		),
	)
}

func renderSeparatorsContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Section dividers")),
				html.P(html.AClass("text-sm text-slate-400"), html.Text("Separate groups of content with optional labels.")),
			),
			html.Div(
				html.AClass("space-y-6"),
				separator.Separator(separator.Props{}, html.T("Team")),
				html.Div(
					html.AClass("flex h-32 items-center justify-center gap-6"),
					html.Text("Start"),
					separator.Separator(separator.Props{Orientation: separator.OrientationVertical, Decoration: separator.DecorationDotted}),
					html.Text("End"),
				),
			),
		),
	)
}

func renderAvatarsContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("User avatars")),
				html.P(html.AClass("text-sm text-slate-400"), html.Text("Fallbacks automatically show initials when no image is available.")),
			),
			html.Div(
				html.AClass("flex flex-wrap items-center gap-6"),
				avatar.Avatar(
					avatar.Props{},
					avatar.Image(avatar.ImageProps{Src: "https://i.pravatar.cc/64?img=12", Alt: "Ryan"}),
					avatar.Fallback(avatar.FallbackProps{}, html.Text("RY")),
				),
				avatar.Avatar(
					avatar.Props{Size: avatar.SizeSm},
					avatar.Fallback(avatar.FallbackProps{}, html.Text("DC")),
				),
			),
		),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Avatar group")),
				html.P(html.AClass("text-sm text-slate-400"), html.Text("Stack avatars with spacing and overflow indicator.")),
			),
			html.Div(
				html.AClass("flex items-center gap-8"),
				avatar.Group(
					avatar.GroupProps{Spacing: avatar.GroupSpacingMd},
					avatar.Avatar(avatar.Props{InGroup: true}, avatar.Image(avatar.ImageProps{Src: "https://i.pravatar.cc/64?img=32", Alt: "Yara"})),
					avatar.Avatar(avatar.Props{InGroup: true}, avatar.Image(avatar.ImageProps{Src: "https://i.pravatar.cc/64?img=19", Alt: "Lee"})),
					avatar.Avatar(avatar.Props{InGroup: true}, avatar.Fallback(avatar.FallbackProps{}, html.Text("MW"))),
					avatar.GroupOverflow(4, avatar.Props{InGroup: true, Size: avatar.SizeSm}),
				),
			),
		),
	)
}

func renderLabelsContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Input labelling")),
				html.P(html.AClass("text-sm text-slate-400"), html.Text("Link labels to form controls and surface validation errors.")),
			),
			html.Div(
				html.AClass("grid gap-6 md:grid-cols-2"),
				html.Div(
					html.AClass("space-y-2"),
					label.Label(label.Props{For: "email"}, html.Text("Email")),
					html.Input(
						html.AId("email"),
						html.AClass("w-full rounded-md border border-slate-800 bg-slate-900 px-3 py-2 text-sm text-slate-100"),
						html.APlaceholder("name@example.com"),
					),
				),
				html.Div(
					html.AClass("space-y-2"),
					label.Label(label.Props{For: "password", Error: "Password must be at least 8 characters"}, html.Text("Password")),
					html.Input(
						html.AId("password"),
						html.AClass("w-full rounded-md border border-destructive bg-slate-900 px-3 py-2 text-sm text-slate-100"),
						html.APlaceholder("Enter password"),
						html.AType("password"),
					),
					html.P(html.AClass("text-xs text-destructive"), html.Text("Password must be at least 8 characters.")),
				),
			),
		),
	)
}

func renderCardsContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Cards")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Surface key information inside a bordered container.")),
			),
			html.Div(
				html.AClass("grid gap-6 md:grid-cols-2"),
				card.Card(
					card.Props{},
					card.Header(
						card.HeaderProps{},
						html.Div(
							html.AClass("space-y-1"),
							card.Title(card.TitleProps{}, html.Text("Change subscription")),
							card.Description(card.DescriptionProps{}, html.Text("Upgrade or downgrade your current plan.")),
						),
					),
					card.Content(
						card.ContentProps{},
						html.Div(
							html.AClass("space-y-3"),
							html.P(html.AClass("text-sm text-muted-foreground"), html.Text("You're currently on the Team plan. Teams get collaborative features, SSO, and priority support.")),
							html.Ul(
								html.AClass("space-y-2 text-sm"),
								html.Li(html.Text("• Unlimited collaborators")),
								html.Li(html.Text("• Shared components")),
								html.Li(html.Text("• Priority support")),
							),
						),
					),
					card.Footer(
						card.FooterProps{},
						html.Div(
							html.AClass("ml-auto flex gap-3"),
							button.Button(button.Props{Variant: button.VariantOutline}, html.Text("Manage plan")),
							button.Button(button.Props{}, html.Text("Upgrade")),
						),
					),
				),
			),
		),
	)
}

func renderAccordionContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Accordion")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Stack collapsible sections to keep dense content organized.")),
			),
			accordion.Accordion(
				accordion.Props{},
				accordion.Item(
					accordion.ItemProps{},
					accordion.Trigger(accordion.TriggerProps{}, html.Text("Is Plain UI accessible?")),
					accordion.Content(
						accordion.ContentProps{},
						html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Yes. We build on semantic HTML, manage focus states, and respect reduced-motion preferences.")),
					),
				),
				accordion.Item(
					accordion.ItemProps{},
					accordion.Trigger(accordion.TriggerProps{}, html.Text("Can I use it with templ?")),
					accordion.Content(
						accordion.ContentProps{},
						html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Absolutely. The APIs are designed to feel templ-friendly and compose with html components.")),
					),
				),
				accordion.Item(
					accordion.ItemProps{},
					accordion.Trigger(accordion.TriggerProps{}, html.Text("Does it support nested content?")),
					accordion.Content(
						accordion.ContentProps{},
						html.Div(
							html.AClass("space-y-2 text-sm text-muted-foreground"),
							html.P(html.Text("Yes. You can place paragraphs, lists, and interactive elements inside each panel.")),
							html.Ul(
								html.AClass("list-disc pl-5"),
								html.Li(html.Text("Links and buttons")),
								html.Li(html.Text("Images or media")),
								html.Li(html.Text("Nested layout blocks")),
							),
						),
					),
				),
			),
		),
	)
}

func renderRadiosContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Radios")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Allow users to make a single choice from a concise list.")),
			),
			html.Fieldset(
				html.AClass("space-y-4 rounded-lg border p-6"),
				html.Legend(html.AClass("text-sm font-semibold uppercase tracking-wide text-muted-foreground"), html.Text("Delivery speed")),
				html.Div(
					html.AClass("space-y-3"),
					html.Label(
						html.AClass("flex items-center gap-3 text-sm"),
						radio.Radio(radio.Props{ID: "delivery-standard", Name: "delivery", Value: "standard", Checked: true}),
						html.Div(
							html.AClass("flex flex-col"),
							html.Span(html.AClass("font-medium"), html.Text("Standard")),
							html.Span(html.AClass("text-xs text-muted-foreground"), html.Text("3-5 business days")),
						),
					),
					html.Label(
						html.AClass("flex items-center gap-3 text-sm"),
						radio.Radio(radio.Props{ID: "delivery-express", Name: "delivery", Value: "express"}),
						html.Div(
							html.AClass("flex flex-col"),
							html.Span(html.AClass("font-medium"), html.Text("Express")),
							html.Span(html.AClass("text-xs text-muted-foreground"), html.Text("Arrives tomorrow")),
						),
					),
					html.Label(
						html.AClass("flex items-center gap-3 text-sm"),
						radio.Radio(radio.Props{ID: "delivery-same-day", Name: "delivery", Value: "same-day"}),
						html.Div(
							html.AClass("flex flex-col"),
							html.Span(html.AClass("font-medium"), html.Text("Same day")),
							html.Span(html.AClass("text-xs text-muted-foreground"), html.Text("Available in select cities")),
						),
					),
				),
			),
		),
	)
}

func renderSkeletonsContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Skeletons")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Display lightweight placeholders while content loads.")),
			),
			html.Div(
				html.AClass("grid gap-6 md:grid-cols-2"),
				html.Div(
					html.AClass("space-y-4 rounded-lg border p-6"),
					skeleton.Skeleton(skeleton.Props{Class: "h-8 w-1/2"}),
					html.Div(
						html.AClass("space-y-2"),
						skeleton.Skeleton(skeleton.Props{Class: "h-4 w-full"}),
						skeleton.Skeleton(skeleton.Props{Class: "h-4 w-2/3"}),
						skeleton.Skeleton(skeleton.Props{Class: "h-4 w-3/4"}),
					),
					html.Div(
						html.AClass("flex gap-3"),
						skeleton.Skeleton(skeleton.Props{Class: "h-9 w-24"}),
						skeleton.Skeleton(skeleton.Props{Class: "h-9 w-32"}),
					),
				),
				html.Div(
					html.AClass("space-y-4 rounded-lg border p-6"),
					skeleton.Skeleton(skeleton.Props{Class: "h-48 w-full rounded-md"}),
					html.Div(
						html.AClass("space-y-2"),
						skeleton.Skeleton(skeleton.Props{Class: "h-4 w-3/5"}),
						skeleton.Skeleton(skeleton.Props{Class: "h-4 w-4/5"}),
					),
				),
			),
		),
	)
}

func renderAspectRatiosContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Aspect Ratios")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Maintain consistent media sizing across breakpoints.")),
			),
			html.Div(
				html.AClass("grid gap-6 md:grid-cols-2"),
				aspectratio.AspectRatio(
					aspectratio.Props{Ratio: aspectratio.RatioVideo, Class: "overflow-hidden rounded-lg border"},
					html.Img(
						html.ASrc("https://images.unsplash.com/photo-1522199755839-a2bacb67c546?auto=format&fit=crop&w=800&q=80"),
						html.AAlt("Team working together"),
						html.AClass("h-full w-full object-cover"),
					),
				),
				aspectratio.AspectRatio(
					aspectratio.Props{Ratio: aspectratio.RatioSquare, Class: "overflow-hidden rounded-lg border bg-muted"},
					html.Div(
						html.AClass("flex h-full w-full items-center justify-center text-center"),
						html.Div(
							html.AClass("space-y-2"),
							html.H3(html.AClass("text-lg font-semibold"), html.Text("Square preview")),
							html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Use any content inside the ratio wrapper.")),
						),
					),
				),
			),
		),
	)
}

func renderBreadcrumbsContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Breadcrumbs")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Provide location context and quick navigation back to parent pages.")),
			),
			breadcrumb.Breadcrumb(
				breadcrumb.Props{Class: "rounded-lg border bg-card px-4 py-3"},
				breadcrumb.List(
					breadcrumb.ListProps{},
					breadcrumb.Item(
						breadcrumb.ItemProps{},
						breadcrumb.Link(breadcrumb.LinkProps{Href: "#"}, html.Text("Dashboard")),
						breadcrumb.Separator(breadcrumb.SeparatorProps{}),
					),
					breadcrumb.Item(
						breadcrumb.ItemProps{},
						breadcrumb.Link(breadcrumb.LinkProps{Href: "#"}, html.Text("Projects")),
						breadcrumb.Separator(breadcrumb.SeparatorProps{}),
					),
					breadcrumb.Item(
						breadcrumb.ItemProps{},
						breadcrumb.Page(breadcrumb.ItemProps{}, html.Text("Plain UI")),
					),
				),
			),
		),
	)
}

func renderCollapsibleContent() html.Node {
	triggerClasses := "flex items-center justify-between gap-3 rounded-md border border-border bg-card px-4 py-3 text-sm font-medium text-card-foreground cursor-pointer"
	contentClasses := "rounded-b-md border border-border border-t-0 bg-card px-4 py-3 text-sm text-muted-foreground"

	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Collapsible")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Progressively reveal supplemental details without leaving the current view.")),
			),
			html.Div(
				html.AClass("space-y-4"),
				collapsible.Collapsible(
					collapsible.Props{Class: "space-y-2"},
					collapsible.Trigger(
						collapsible.TriggerProps{Class: triggerClasses},
						html.Span(html.Text("What is Plain UI?")),
						lucide.ChevronDown(html.AClass("size-4 text-muted-foreground")),
					),
					collapsible.Content(
						collapsible.ContentProps{},
						html.Div(
							html.AClass(contentClasses),
							html.P(html.Text("Plain UI is a growing collection of composable components built on the plainkit html helpers.")),
						),
					),
				),
				collapsible.Collapsible(
					collapsible.Props{Class: "space-y-2", Open: true},
					collapsible.Trigger(
						collapsible.TriggerProps{Class: triggerClasses},
						html.Span(html.Text("Can I style it with Tailwind?")),
						lucide.ChevronDown(html.AClass("size-4 text-muted-foreground")),
					),
					collapsible.Content(
						collapsible.ContentProps{},
						html.Div(
							html.AClass(contentClasses),
							html.P(html.Text("Yes! Utility classes are merged automatically so you can override tokens per instance.")),
						),
					),
				),
			),
		),
	)
}

func renderFormContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Form Building Blocks")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Compose labels, descriptions, and validation messages around existing inputs.")),
			),
			html.Div(
				html.AClass("grid gap-6 lg:grid-cols-2"),
				html.Form(
					html.AClass("space-y-6 rounded-lg border bg-card p-6 shadow-xs"),
					form.Item(
						form.ItemProps{},
						form.Label(form.LabelProps{For: "full-name"}, html.Text("Full name")),
						input.Input(input.Props{
							ID:          "full-name",
							Name:        "full_name",
							Placeholder: "Ada Lovelace",
							Required:    true,
						}),
						form.Description(form.DescriptionProps{}, html.Text("We use your full name to personalize communications.")),
					),
					form.Item(
						form.ItemProps{},
						form.Label(form.LabelProps{For: "email-address"}, html.Text("Email address")),
						input.Input(input.Props{
							ID:          "email-address",
							Name:        "email",
							Type:        input.TypeEmail,
							Placeholder: "ada@example.com",
							Required:    true,
						}, html.AAutocomplete("email")),
						form.Description(form.DescriptionProps{}, html.Text("We'll only send important updates.")),
					),
					form.Item(
						form.ItemProps{},
						form.Label(form.LabelProps{For: "account-password"}, html.Text("Password")),
						input.Input(input.Props{
							ID:                 "account-password",
							Name:               "password",
							Type:               input.TypePassword,
							Placeholder:        "Create a strong password",
							Required:           true,
							ShowPasswordToggle: true,
							HasError:           true,
						}),
						form.Message(form.MessageProps{Variant: form.MessageVariantError}, html.Text("Use at least 8 characters with a number and symbol.")),
					),
					html.Div(
						html.AClass("flex items-center justify-end gap-3"),
						button.Button(button.Props{Variant: button.VariantOutline, Type: button.TypeButton}, html.Text("Cancel")),
						button.Button(button.Props{Type: button.TypeSubmit}, html.Text("Save changes")),
					),
				),
				html.Div(
					html.AClass("space-y-4 rounded-lg border bg-card p-6 shadow-xs"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Inline controls")),
					html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Use ItemFlex to align toggles and short helpers on a single row.")),
					form.ItemFlex(
						form.ItemProps{},
						html.Child(switchcomp.Switch(switchcomp.Props{
							ID:      "alerts-switch",
							Name:    "alerts",
							Checked: true,
						}, html.Text("Email notifications"))),
						html.Div(
							html.AClass("text-sm text-muted-foreground"),
							html.Text("Send me important product and billing updates."),
						),
					),
					form.ItemFlex(
						form.ItemProps{},
						html.Child(switchcomp.Switch(switchcomp.Props{
							ID:   "beta-switch",
							Name: "beta_program",
						}, html.Text("Join beta program"))),
						html.Div(
							html.AClass("text-sm text-muted-foreground"),
							html.Text("Access experimental features before they ship."),
						),
					),
				),
			),
		),
	)
}

func renderInputOTPContent() html.Node {
	buildSlots := func(start int, hasError bool) []html.DivArg {
		items := make([]html.DivArg, 0, 3)
		for i := 0; i < 3; i++ {
			items = append(items, html.Child(inputotp.Slot(inputotp.SlotProps{
				Index:       start + i,
				Placeholder: "0",
				HasError:    hasError,
			})))
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
						inputotp.Group(inputotp.GroupProps{}, buildSlots(0, true)...),
						inputotp.Separator(inputotp.SeparatorProps{}, html.Text("-")),
						inputotp.Group(inputotp.GroupProps{}, buildSlots(3, true)...),
					),
					form.Message(form.MessageProps{Variant: form.MessageVariantError}, html.Text("The verification code is incorrect.")),
				),
				html.Div(
					html.AClass("space-y-3 rounded-lg border bg-card p-6 shadow-xs"),
					html.Label(html.AClass("text-sm font-medium"), html.Text("SMS code")),
					inputotp.InputOTP(
						inputotp.Props{ID: "otp-secondary", Name: "sms_code", Value: "123456", Autofocus: true, Class: "flex-wrap gap-y-3"},
						inputotp.Group(inputotp.GroupProps{}, buildSlots(0, false)...),
						inputotp.Separator(inputotp.SeparatorProps{}, html.Text("-")),
						inputotp.Group(inputotp.GroupProps{}, buildSlots(3, false)...),
					),
					form.Description(form.DescriptionProps{}, html.Text("Autofocus jumps to the next slot as the user types.")),
				),
			),
		),
	)
}

func renderRatingsContent() html.Node {
	buildItems := func(style rating.Style) []html.DivArg {
		items := make([]html.DivArg, 0, 5)
		for i := 1; i <= 5; i++ {
			items = append(items, html.Child(rating.Item(rating.ItemProps{Value: i, Style: style})))
		}
		return items
	}

	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Ratings")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Support stars, hearts, or emoji reactions with keyboard and pointer interactivity.")),
			),
			html.Div(
				html.AClass("grid gap-6 md:grid-cols-2"),
				html.Div(
					html.AClass("space-y-4 rounded-lg border bg-card p-6 shadow-xs"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Half-star precision")),
					rating.Rating(
						rating.Props{ID: "rating-stars", Name: "product_rating", Precision: 0.5},
						rating.Group(rating.GroupProps{}, buildItems(rating.StyleStar)...),
					),
					html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Click or hover to preview partial fills before committing.")),
				),
				html.Div(
					html.AClass("space-y-4 rounded-lg border bg-card p-6 shadow-xs"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Heart feedback")),
					rating.Rating(
						rating.Props{ID: "rating-hearts", Precision: 1, OnlyInteger: true},
						rating.Group(rating.GroupProps{}, buildItems(rating.StyleHeart)...),
					),
					html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Perfect when you just need a simple like-meter.")),
				),
				html.Div(
					html.AClass("space-y-4 rounded-lg border bg-card p-6 shadow-xs md:col-span-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Emoji mood")),
					rating.Rating(
						rating.Props{ID: "rating-emoji", Value: 4.2, Precision: 0.5, ReadOnly: true},
						rating.Group(rating.GroupProps{}, buildItems(rating.StyleEmoji)...),
					),
					html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Read-only ratings show aggregate sentiment without interaction.")),
				),
			),
		),
	)
}

func renderPopoversContent() html.Node {
	const infoPopoverID = "profile-popover"
	const hoverPopoverID = "shortcut-popover"

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

func renderToastsContent() html.Node {
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
					toast.ToastTrigger(
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
					toast.ToastTrigger(
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
					toast.ToastTrigger(
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
					toast.ToastTrigger(
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

func renderPaginationContent() html.Node {
	data := pagination.CreatePagination(6, 20, 5)

	items := []html.UlArg{
		pagination.Item(
			pagination.ItemProps{},
			pagination.Previous(pagination.PreviousProps{
				Href:     previousHref(data),
				Disabled: !data.HasPrevious,
				Label:    "Previous",
			}),
		),
	}

	if first := data.Pages[0]; first > 1 {
		items = append(items,
			pagination.Item(
				pagination.ItemProps{},
				pagination.Link(pagination.LinkProps{Href: "?page=1", IsActive: data.CurrentPage == 1}, html.T("1")),
			),
		)
		if first > 2 {
			items = append(items, pagination.Item(pagination.ItemProps{}, pagination.Ellipsis()))
		}
	}

	for _, pageNumber := range data.Pages {
		label := strconv.Itoa(pageNumber)
		items = append(items,
			pagination.Item(
				pagination.ItemProps{},
				pagination.Link(
					pagination.LinkProps{
						Href:     "?page=" + label,
						IsActive: pageNumber == data.CurrentPage,
					},
					html.T(label),
				),
			),
		)
	}

	lastVisible := data.Pages[len(data.Pages)-1]
	if lastVisible < data.TotalPages {
		if lastVisible < data.TotalPages-1 {
			items = append(items, pagination.Item(pagination.ItemProps{}, pagination.Ellipsis()))
		}
		lastLabel := strconv.Itoa(data.TotalPages)
		items = append(items,
			pagination.Item(
				pagination.ItemProps{},
				pagination.Link(
					pagination.LinkProps{
						Href:     "?page=" + lastLabel,
						IsActive: data.CurrentPage == data.TotalPages,
					},
					html.T(lastLabel),
				),
			),
		)
	}

	items = append(items,
		pagination.Item(
			pagination.ItemProps{},
			pagination.Next(pagination.NextProps{
				Href:     nextHref(data),
				Disabled: !data.HasNext,
				Label:    "Next",
			}),
		),
	)

	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Pagination")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Navigate long collections page by page.")),
			),
			pagination.Pagination(
				pagination.Props{},
				pagination.Content(pagination.ContentProps{}, items...),
			),
		),
	)
}

func previousHref(data struct {
	CurrentPage int
	TotalPages  int
	Pages       []int
	HasPrevious bool
	HasNext     bool
}) string {
	if !data.HasPrevious {
		return ""
	}
	return "?page=" + strconv.Itoa(data.CurrentPage-1)
}

func nextHref(data struct {
	CurrentPage int
	TotalPages  int
	Pages       []int
	HasPrevious bool
	HasNext     bool
}) string {
	if !data.HasNext {
		return ""
	}
	return "?page=" + strconv.Itoa(data.CurrentPage+1)
}

func renderTablesContent() html.Node {
	type invoice struct {
		Number string
		Status string
		Method string
		Amount string
	}

	invoices := []invoice{
		{Number: "INV-001", Status: "Paid", Method: "Credit Card", Amount: "$2,500"},
		{Number: "INV-002", Status: "Processing", Method: "Bank Transfer", Amount: "$980"},
		{Number: "INV-003", Status: "Unpaid", Method: "Credit Card", Amount: "$1,430"},
		{Number: "INV-004", Status: "Paid", Method: "PayPal", Amount: "$3,200"},
	}

	bodyRows := make([]html.TbodyArg, 0, len(invoices))
	for idx, inv := range invoices {
		bodyRows = append(bodyRows,
			table.Row(
				table.RowProps{Selected: idx == 1},
				table.Cell(table.CellProps{}, html.Text(inv.Number)),
				table.Cell(table.CellProps{}, html.Text(inv.Status)),
				table.Cell(table.CellProps{}, html.Text(inv.Method)),
				table.Cell(table.CellProps{}, html.Text(inv.Amount)),
			),
		)
	}

	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Tables")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Present structured data with consistent spacing and hover affordances.")),
			),
			table.Table(
				table.Props{},
				table.Caption(table.CaptionProps{}, html.Text("Invoice history for Q3 revenue.")),
				table.Header(
					table.HeaderProps{},
					table.Row(
						table.RowProps{},
						table.Head(table.HeadProps{}, html.Text("Invoice")),
						table.Head(table.HeadProps{}, html.Text("Status")),
						table.Head(table.HeadProps{}, html.Text("Method")),
						table.Head(table.HeadProps{}, html.Text("Amount")),
					),
				),
				table.Body(table.BodyProps{}, bodyRows...),
				table.Footer(
					table.FooterProps{},
					table.Row(
						table.RowProps{},
						table.Cell(table.CellProps{Class: "font-medium"}, html.Text("Total")),
						table.Cell(table.CellProps{}, html.Text("2 paid")),
						table.Cell(table.CellProps{}, html.Text("2 pending")),
						table.Cell(table.CellProps{Class: "font-semibold"}, html.Text("$8,110")),
					),
				),
			),
		),
	)
}

func renderCodeContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Code")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Syntax-highlighted code blocks with optional copy functionality.")),
			),
			html.Div(
				html.AClass("grid gap-6 md:grid-cols-2"),
				html.Div(
					html.AClass("space-y-4"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("JavaScript Example")),
					code.Code(
						code.Props{
							Language:       "javascript",
							ShowCopyButton: true,
							Size:           code.SizeSm,
						},
						html.Text(`function fibonacci(n) {
  if (n <= 1) return n;
  return fibonacci(n - 1) + fibonacci(n - 2);
}

console.log(fibonacci(10)); // 55`),
					),
				),
				html.Div(
					html.AClass("space-y-4"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Go Example")),
					code.Code(
						code.Props{
							Language:       "go",
							ShowCopyButton: true,
							Size:           code.SizeLg,
						},
						html.Text(`package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        fmt.Printf("Hello, World! %d\n", i)
    }
}`),
					),
				),
			),
		),
	)
}

func renderDialogsContent() html.Node {
	const dialogID = "demo-dialog"

	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Dialogs")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Modal dialogs for confirmations, forms, and important communications.")),
			),
			html.Div(
				html.AClass("space-y-6"),
				html.Div(
					html.AClass("flex flex-wrap gap-3"),
					dialog.Trigger(
						dialog.TriggerProps{For: dialogID},
						button.Props{},
						html.Text("Open Dialog"),
					),
				),
				dialog.Content(
					dialog.ContentProps{ID: dialogID},
					dialog.Header(
						dialog.HeaderProps{},
						dialog.Title(dialog.TitleProps{}, html.Text("Confirm Action")),
						dialog.Description(dialog.DescriptionProps{}, html.Text("Are you sure you want to delete this item? This action cannot be undone.")),
					),
					dialog.Footer(
						dialog.FooterProps{},
						dialog.Close(dialog.CloseProps{For: dialogID},
							button.Button(button.Props{Variant: button.VariantOutline}, html.Text("Cancel"))),
						button.Button(button.Props{Variant: button.VariantDestructive}, html.Text("Delete")),
					),
				),
			),
		),
	)
}

func renderDropdownsContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Dropdowns")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Floating menus with items, groups, and keyboard navigation.")),
			),
			html.Div(
				html.AClass("flex flex-wrap gap-6"),
				html.Div(
					html.AClass("space-y-4"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Basic Menu")),
					dropdown.Dropdown(
						dropdown.Props{},
						dropdown.Trigger(
							dropdown.TriggerProps{For: "basic-dropdown"},
							button.Props{Variant: button.VariantOutline},
							html.Text("Options"),
							lucide.ChevronDown(html.AClass("ml-2 size-4")),
						),
						dropdown.Content(
							dropdown.ContentProps{ID: "basic-dropdown"},
							dropdown.Item(dropdown.ItemProps{}, html.Span(html.Text("Profile"))),
							dropdown.Item(dropdown.ItemProps{}, html.Span(html.Text("Settings"))),
							dropdown.Separator(dropdown.SeparatorProps{}),
							dropdown.Item(dropdown.ItemProps{}, html.Span(html.Text("Sign out"))),
						),
					),
				),
				html.Div(
					html.AClass("space-y-4"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("With Groups")),
					dropdown.Dropdown(
						dropdown.Props{},
						dropdown.Trigger(
							dropdown.TriggerProps{For: "grouped-dropdown"},
							button.Props{Variant: button.VariantOutline},
							html.Text("Actions"),
							lucide.ChevronDown(html.AClass("ml-2 size-4")),
						),
						dropdown.Content(
							dropdown.ContentProps{ID: "grouped-dropdown"},
							dropdown.Group(
								dropdown.GroupProps{},
								dropdown.Label(dropdown.LabelProps{}, html.Text("Account")),
								dropdown.Item(dropdown.ItemProps{}, html.Span(html.Text("Profile"))),
								dropdown.Item(dropdown.ItemProps{}, html.Span(html.Text("Billing"))),
							),
							dropdown.Separator(dropdown.SeparatorProps{}),
							dropdown.Group(
								dropdown.GroupProps{},
								dropdown.Label(dropdown.LabelProps{}, html.Text("Actions")),
								dropdown.Item(dropdown.ItemProps{}, html.Span(html.Text("New file")), dropdown.Shortcut(dropdown.ShortcutProps{}, html.Text("⌘N"))),
								dropdown.Item(dropdown.ItemProps{}, html.Span(html.Text("Save")), dropdown.Shortcut(dropdown.ShortcutProps{}, html.Text("⌘S"))),
							),
						),
					),
				),
			),
		),
	)
}

func renderSelectBoxesContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Select Boxes")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Dropdown selectors with search and multi-selection capabilities.")),
			),
			html.Div(
				html.AClass("grid gap-6 md:grid-cols-2"),
				html.Div(
					html.AClass("space-y-4"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Single Select")),
					selectbox.SelectBox(
						selectbox.Props{},
						selectbox.Trigger(
							selectbox.TriggerProps{Name: "framework"},
							"single-select-content",
							selectbox.Value(selectbox.ValueProps{Placeholder: "Select framework..."}),
						),
						selectbox.Content(
							selectbox.ContentProps{ID: "single-select-content"},
							selectbox.Group(
								selectbox.GroupProps{},
								selectbox.Label(selectbox.LabelProps{}, html.Text("Frontend")),
								selectbox.Item(selectbox.ItemProps{Value: "react"}, html.Text("React")),
								selectbox.Item(selectbox.ItemProps{Value: "vue"}, html.Text("Vue.js")),
								selectbox.Item(selectbox.ItemProps{Value: "angular"}, html.Text("Angular")),
							),
							selectbox.Group(
								selectbox.GroupProps{},
								selectbox.Label(selectbox.LabelProps{}, html.Text("Backend")),
								selectbox.Item(selectbox.ItemProps{Value: "go"}, html.Text("Go")),
								selectbox.Item(selectbox.ItemProps{Value: "node"}, html.Text("Node.js")),
								selectbox.Item(selectbox.ItemProps{Value: "python"}, html.Text("Python")),
							),
						),
					),
				),
				html.Div(
					html.AClass("space-y-4"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("With Search")),
					selectbox.SelectBox(
						selectbox.Props{},
						selectbox.Trigger(
							selectbox.TriggerProps{Name: "country"},
							"searchable-select-content",
							selectbox.Value(selectbox.ValueProps{Placeholder: "Select country..."}),
						),
						selectbox.Content(
							selectbox.ContentProps{
								ID:                "searchable-select-content",
								SearchPlaceholder: "Search countries...",
							},
							selectbox.Item(selectbox.ItemProps{Value: "us"}, html.Text("United States")),
							selectbox.Item(selectbox.ItemProps{Value: "uk"}, html.Text("United Kingdom")),
							selectbox.Item(selectbox.ItemProps{Value: "de"}, html.Text("Germany")),
							selectbox.Item(selectbox.ItemProps{Value: "fr"}, html.Text("France")),
							selectbox.Item(selectbox.ItemProps{Value: "jp"}, html.Text("Japan")),
						),
					),
				),
			),
		),
	)
}

func renderTooltipsContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Tooltips")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Contextual information on hover with flexible positioning.")),
			),
			html.Div(
				html.AClass("grid gap-6 md:grid-cols-2"),
				html.Div(
					html.AClass("space-y-4"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Basic Tooltips")),
					html.Div(
						html.AClass("flex flex-wrap gap-4"),
						tooltip.Trigger(
							tooltip.TriggerProps{For: "tooltip-top"},
							button.Button(button.Props{Variant: button.VariantOutline}, html.Text("Top")),
						),
						tooltip.Content(
							tooltip.ContentProps{
								ID:        "tooltip-top",
								Position:  tooltip.PositionTop,
								ShowArrow: true,
							},
							html.Text("This tooltip appears on top"),
						),
						tooltip.Trigger(
							tooltip.TriggerProps{For: "tooltip-right"},
							button.Button(button.Props{Variant: button.VariantOutline}, html.Text("Right")),
						),
						tooltip.Content(
							tooltip.ContentProps{
								ID:        "tooltip-right",
								Position:  tooltip.PositionRight,
								ShowArrow: true,
							},
							html.Text("This tooltip appears on the right"),
						),
					),
				),
				html.Div(
					html.AClass("space-y-4"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("With Delays")),
					html.Div(
						html.AClass("flex flex-wrap gap-4"),
						tooltip.Trigger(
							tooltip.TriggerProps{For: "tooltip-delayed"},
							button.Button(button.Props{Variant: button.VariantOutline}, html.Text("Hover me")),
						),
						tooltip.Content(
							tooltip.ContentProps{
								ID:            "tooltip-delayed",
								Position:      tooltip.PositionBottom,
								ShowArrow:     true,
								HoverDelay:    500,
								HoverOutDelay: 200,
							},
							html.Text("This tooltip has custom delays"),
						),
					),
				),
			),
		),
	)
}

func renderCalendarContent() html.Node {
	return card.Card(card.Props{},
		card.Header(card.HeaderProps{},
			card.Title(card.TitleProps{}, html.Text("Calendar")),
			card.Description(card.DescriptionProps{}, html.Text("A calendar component for date selection.")),
		),
		card.Content(card.ContentProps{},
			html.Div(
				html.AClass("space-y-8"),

				// Basic Calendar
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Basic Calendar")),
					html.Div(
						html.AClass("max-w-md"),
						calendar.Calendar(calendar.Props{
							ID:   "basic-calendar",
							Name: "selected-date",
						}),
					),
				),

				// Calendar with Initial Value
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Calendar with Initial Value")),
					html.Div(
						html.AClass("max-w-md"),
						calendar.Calendar(calendar.Props{
							ID:   "calendar-with-value",
							Name: "preset-date",
							// Value: time.Now(), // Would set current date
						}),
					),
				),

				// Calendar with Custom Locale
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Calendar with French Locale")),
					html.Div(
						html.AClass("max-w-md"),
						calendar.Calendar(calendar.Props{
							ID:        "calendar-french",
							Name:      "french-date",
							LocaleTag: calendar.LocaleTagFrench,
						}),
					),
				),

				// Calendar Starting on Sunday
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Calendar Starting on Sunday")),
					html.Div(
						html.AClass("max-w-md"),
						calendar.Calendar(calendar.Props{
							ID:          "calendar-sunday",
							Name:        "sunday-date",
							StartOfWeek: &[]calendar.Day{calendar.Sunday}[0],
						}),
					),
				),
			),
		),
	)
}

func renderCarouselContent() html.Node {
	return card.Card(card.Props{},
		card.Header(card.HeaderProps{},
			card.Title(card.TitleProps{}, html.Text("Carousel")),
			card.Description(card.DescriptionProps{}, html.Text("A carousel component for displaying sliding content.")),
		),
		card.Content(card.ContentProps{},
			html.Div(
				html.AClass("space-y-8"),

				// Basic Carousel
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Basic Carousel")),
					html.Div(
						html.AClass("relative"),
						carousel.Carousel(carousel.Props{
							ID: "basic-carousel",
						},
							carousel.Content(carousel.ContentProps{},
								carousel.Item(carousel.ItemProps{},
									html.Div(
										html.AClass("bg-gradient-to-r from-blue-500 to-purple-600 h-64 flex items-center justify-center text-white text-xl font-bold rounded-lg"),
										html.Text("Slide 1"),
									),
								),
								carousel.Item(carousel.ItemProps{},
									html.Div(
										html.AClass("bg-gradient-to-r from-green-500 to-blue-600 h-64 flex items-center justify-center text-white text-xl font-bold rounded-lg"),
										html.Text("Slide 2"),
									),
								),
								carousel.Item(carousel.ItemProps{},
									html.Div(
										html.AClass("bg-gradient-to-r from-red-500 to-pink-600 h-64 flex items-center justify-center text-white text-xl font-bold rounded-lg"),
										html.Text("Slide 3"),
									),
								),
							),
							carousel.Previous(carousel.PreviousProps{}),
							carousel.Next(carousel.NextProps{}),
							carousel.Indicators(carousel.IndicatorsProps{Count: 3}),
						),
					),
				),

				// Autoplay Carousel with Loop
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Autoplay Carousel")),
					html.Div(
						html.AClass("relative"),
						carousel.Carousel(carousel.Props{
							ID:       "autoplay-carousel",
							Autoplay: true,
							Interval: 3000,
							Loop:     true,
						},
							carousel.Content(carousel.ContentProps{},
								carousel.Item(carousel.ItemProps{},
									html.Div(
										html.AClass("bg-gradient-to-r from-yellow-400 to-orange-500 h-48 flex items-center justify-center text-white text-lg font-semibold rounded-lg"),
										html.Text("Auto Slide 1 - 3s intervals"),
									),
								),
								carousel.Item(carousel.ItemProps{},
									html.Div(
										html.AClass("bg-gradient-to-r from-purple-500 to-indigo-600 h-48 flex items-center justify-center text-white text-lg font-semibold rounded-lg"),
										html.Text("Auto Slide 2 - Loops infinitely"),
									),
								),
								carousel.Item(carousel.ItemProps{},
									html.Div(
										html.AClass("bg-gradient-to-r from-teal-500 to-cyan-600 h-48 flex items-center justify-center text-white text-lg font-semibold rounded-lg"),
										html.Text("Auto Slide 3 - Hover to pause"),
									),
								),
							),
							carousel.Previous(carousel.PreviousProps{}),
							carousel.Next(carousel.NextProps{}),
							carousel.Indicators(carousel.IndicatorsProps{Count: 3}),
						),
					),
				),
			),
		),
	)
}

func renderTagsInputContent() html.Node {
	return card.Card(card.Props{},
		card.Header(card.HeaderProps{},
			card.Title(card.TitleProps{}, html.Text("Tags Input")),
			card.Description(card.DescriptionProps{}, html.Text("An input component for entering multiple tags.")),
		),
		card.Content(card.ContentProps{},
			html.Div(
				html.AClass("space-y-8"),

				// Basic Tags Input
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Basic Tags Input")),
					html.Div(
						html.AClass("max-w-md"),
						tagsinput.TagsInput(tagsinput.Props{
							ID:          "basic-tags",
							Name:        "tags",
							Placeholder: "Type and press Enter or comma to add tags",
						}),
					),
					html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Press Enter or comma to add tags. Backspace to remove the last tag when input is empty.")),
				),

				// Pre-filled Tags Input
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Pre-filled Tags")),
					html.Div(
						html.AClass("max-w-md"),
						tagsinput.TagsInput(tagsinput.Props{
							ID:          "prefilled-tags",
							Name:        "prefilled-tags",
							Value:       []string{"React", "Go", "JavaScript", "TypeScript"},
							Placeholder: "Add more tags...",
						}),
					),
				),

				// Disabled Tags Input
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Disabled Tags Input")),
					html.Div(
						html.AClass("max-w-md"),
						tagsinput.TagsInput(tagsinput.Props{
							ID:          "disabled-tags",
							Name:        "disabled-tags",
							Value:       []string{"Read-only", "Disabled"},
							Placeholder: "Cannot add tags",
							Disabled:    true,
						}),
					),
				),

				// Tags Input with Error State
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Error State")),
					html.Div(
						html.AClass("max-w-md"),
						tagsinput.TagsInput(tagsinput.Props{
							ID:          "error-tags",
							Name:        "error-tags",
							Placeholder: "This field has an error",
							HasError:    true,
						}),
					),
					html.P(html.AClass("text-sm text-destructive"), html.Text("Please add at least one tag.")),
				),
			),
		),
	)
}

func renderTimePickerContent() html.Node {
	return card.Card(card.Props{},
		card.Header(card.HeaderProps{},
			card.Title(card.TitleProps{}, html.Text("Time Picker")),
			card.Description(card.DescriptionProps{}, html.Text("A time picker component for selecting hours and minutes.")),
		),
		card.Content(card.ContentProps{},
			html.Div(
				html.AClass("space-y-8"),

				// Basic Time Picker (24-hour)
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("24-Hour Format")),
					html.Div(
						html.AClass("max-w-xs"),
						timepicker.TimePicker(timepicker.Props{
							ID:          "time-24h",
							Name:        "time-24h",
							Use12Hours:  false,
							Placeholder: "Select time (24h)",
						}),
					),
				),

				// 12-Hour Time Picker
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("12-Hour Format")),
					html.Div(
						html.AClass("max-w-xs"),
						timepicker.TimePicker(timepicker.Props{
							ID:          "time-12h",
							Name:        "time-12h",
							Use12Hours:  true,
							Placeholder: "Select time (12h)",
						}),
					),
				),

				// Time Picker with 15-minute Steps
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("15-Minute Steps")),
					html.Div(
						html.AClass("max-w-xs"),
						timepicker.TimePicker(timepicker.Props{
							ID:          "time-15min",
							Name:        "time-15min",
							Use12Hours:  true,
							Step:        15,
							Placeholder: "Select time (15min steps)",
						}),
					),
					html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Minutes are shown in 15-minute increments.")),
				),

				// Disabled Time Picker
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Disabled")),
					html.Div(
						html.AClass("max-w-xs"),
						timepicker.TimePicker(timepicker.Props{
							ID:          "time-disabled",
							Name:        "time-disabled",
							Placeholder: "Disabled time picker",
							Disabled:    true,
						}),
					),
				),

				// Time Picker with Error State
				html.Div(
					html.AClass("space-y-2"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Error State")),
					html.Div(
						html.AClass("max-w-xs"),
						timepicker.TimePicker(timepicker.Props{
							ID:          "time-error",
							Name:        "time-error",
							Placeholder: "Time with error",
							HasError:    true,
							Required:    true,
						}),
					),
					html.P(html.AClass("text-sm text-destructive"), html.Text("Please select a time.")),
				),
			),
		),
	)
}
