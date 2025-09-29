package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	democss "github.com/plainkit/ui/cmd/demo/internal/css"
	"github.com/plainkit/ui/cmd/demo/internal/handlers"
)

type page struct {
	Path    string
	Label   string
	Content func() html.Node
}

var pages = []page{
	{Path: "/accordion", Label: "Accordion", Content: handlers.RenderAccordionContent},
	{Path: "/alerts", Label: "Alerts", Content: handlers.RenderAlertsContent},
	{Path: "/aspect-ratios", Label: "Aspect Ratios", Content: handlers.RenderAspectRatiosContent},
	{Path: "/avatars", Label: "Avatars", Content: handlers.RenderAvatarsContent},
	{Path: "/badges", Label: "Badges", Content: handlers.RenderBadgesContent},
	{Path: "/breadcrumbs", Label: "Breadcrumbs", Content: handlers.RenderBreadcrumbsContent},
	{Path: "/buttons", Label: "Buttons", Content: handlers.RenderButtonsContent},
	{Path: "/calendar", Label: "Calendar", Content: handlers.RenderCalendarContent},
	{Path: "/cards", Label: "Cards", Content: handlers.RenderCardsContent},
	{Path: "/carousel", Label: "Carousel", Content: handlers.RenderCarouselContent},
	{Path: "/checkboxes", Label: "Checkboxes", Content: handlers.RenderCheckboxesContent},
	{Path: "/code", Label: "Code", Content: handlers.RenderCodeContent},
	{Path: "/collapsible", Label: "Collapsible", Content: handlers.RenderCollapsibleContent},
	{Path: "/dialogs", Label: "Dialogs", Content: handlers.RenderDialogsContent},
	{Path: "/dropdowns", Label: "Dropdowns", Content: handlers.RenderDropdownsContent},
	{Path: "/forms", Label: "Form Helpers", Content: handlers.RenderFormContent},
	{Path: "/input-otp", Label: "Input OTP", Content: handlers.RenderInputOTPContent},
	{Path: "/inputs", Label: "Inputs", Content: handlers.RenderInputsContent},
	{Path: "/labels", Label: "Labels", Content: handlers.RenderLabelsContent},
	{Path: "/pagination", Label: "Pagination", Content: handlers.RenderPaginationContent},
	{Path: "/popovers", Label: "Popovers", Content: handlers.RenderPopoversContent},
	{Path: "/progress", Label: "Progress", Content: handlers.RenderProgressContent},
	{Path: "/radios", Label: "Radios", Content: handlers.RenderRadiosContent},
	{Path: "/ratings", Label: "Ratings", Content: handlers.RenderRatingsContent},
	{Path: "/selectboxes", Label: "Select Boxes", Content: handlers.RenderSelectBoxesContent},
	{Path: "/separators", Label: "Separators", Content: handlers.RenderSeparatorsContent},
	{Path: "/skeletons", Label: "Skeletons", Content: handlers.RenderSkeletonsContent},
	{Path: "/sliders", Label: "Sliders", Content: handlers.RenderSlidersContent},
	{Path: "/switches", Label: "Switches", Content: handlers.RenderSwitchesContent},
	{Path: "/tables", Label: "Tables", Content: handlers.RenderTablesContent},
	{Path: "/tabs", Label: "Tabs", Content: handlers.RenderTabsContent},
	{Path: "/tags-input", Label: "Tags Input", Content: handlers.RenderTagsInputContent},
	{Path: "/textareas", Label: "Textareas", Content: handlers.RenderTextareasContent},
	{Path: "/timepicker", Label: "Time Picker", Content: handlers.RenderTimePickerContent},
	{Path: "/toasts", Label: "Toasts", Content: handlers.RenderToastsContent},
	{Path: "/tooltips", Label: "Tooltips", Content: handlers.RenderTooltipsContent},
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/assets/styles.css", cssHandler)
	mux.HandleFunc("/robots.txt", robotsHandler)

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

func robotsHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=86400") // Cache for 24 hours
	robotsContent := `User-agent: *
Allow: /
Crawl-delay: 0
Disallow:
`

	if _, err := w.Write([]byte(robotsContent)); err != nil {
		log.Printf("write robots.txt: %v", err)
	}
}

func renderPage(activePath string, body html.Node) string {
	assets := html.NewAssets()
	assets.Collect(body)

	headChildren := []html.HeadArg{
		html.Child(html.Meta(html.ACharset("utf-8"))),
		html.Child(html.Meta(html.AName("viewport"), html.AContent("width=device-width, initial-scale=1"))),
		html.Child(html.Title(html.Text("Plain UI - Modern Go Components for Web Development"))),

		// SEO Meta Tags
		html.Child(html.Meta(html.AName("description"), html.AContent("Plain UI is a comprehensive collection of modern, accessible, and customizable UI components built with Go and HTMX for building beautiful web applications."))),
		html.Child(html.Meta(html.AName("keywords"), html.AContent("go ui components, htmx components, web components, golang ui library, plainkit ui, tailwind components, accessible components, ui kit"))),
		html.Child(html.Meta(html.AName("author"), html.AContent("Plain UI Team"))),
		html.Child(html.Meta(html.AName("robots"), html.AContent("index, follow"))),

		// Open Graph Meta Tags
		html.Child(html.Meta(html.AName("og:title"), html.AContent("Plain UI - Modern Go Components for Web Development"))),
		html.Child(html.Meta(html.AName("og:description"), html.AContent("Build beautiful, accessible web applications with Plain UI's comprehensive collection of Go components featuring HTMX integration."))),
		html.Child(html.Meta(html.AName("og:type"), html.AContent("website"))),
		html.Child(html.Meta(html.AName("og:site_name"), html.AContent("Plain UI"))),
		html.Child(html.Meta(html.AName("og:locale"), html.AContent("en_US"))),

		// Twitter Card Meta Tags
		html.Child(html.Meta(html.AName("twitter:card"), html.AContent("summary_large_image"))),
		html.Child(html.Meta(html.AName("twitter:title"), html.AContent("Plain UI - Modern Go Components for Web Development"))),
		html.Child(html.Meta(html.AName("twitter:description"), html.AContent("Build beautiful, accessible web applications with Plain UI's comprehensive collection of Go components."))),

		// Additional SEO
		html.Child(html.Meta(html.AHttpEquiv("X-UA-Compatible"), html.AContent("IE=edge"))),
		html.Child(html.Meta(html.AName("theme-color"), html.AContent("#000000"))),
		html.Child(html.Link(html.ARel("canonical"), html.AHref("https://plainui.com"))),

		// Stylesheet
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
		html.ALang("en"),
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
