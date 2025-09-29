package carousel

import (
	"crypto/rand"
	_ "embed"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/internal/classnames"
)

type Props struct {
	ID       string
	Class    string
	Attrs    []html.Global
	Autoplay bool
	Interval int
	Loop     bool
}

type ContentProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type ItemProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type PreviousProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type NextProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type IndicatorsProps struct {
	ID    string
	Class string
	Attrs []html.Global
	Count int
}

// Carousel renders a carousel container for sliding content
func Carousel(props Props, args ...html.DivArg) html.Node {
	id := props.ID
	if id == "" {
		id = randomID("carousel")
	}

	interval := props.Interval
	if interval == 0 {
		interval = 5000
	}

	divArgs := []html.DivArg{
		html.AId(id),
		html.AClass(classnames.Merge("relative overflow-hidden w-full", props.Class)),
		html.AData("pui-carousel", ""),
		html.AData("pui-carousel-current", "0"),
		html.AData("pui-carousel-autoplay", strconv.FormatBool(props.Autoplay)),
		html.AData("pui-carousel-interval", fmt.Sprintf("%d", interval)),
		html.AData("pui-carousel-loop", strconv.FormatBool(props.Loop)),
	}

	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}

	divArgs = append(divArgs, args...)

	return html.Div(divArgs...).WithAssets("", carouselJS, "ui-carousel")
}

// Content creates the carousel track that contains the items
func Content(props ContentProps, args ...html.DivArg) html.Node {
	divArgs := []html.DivArg{
		html.AClass(classnames.Merge("flex h-full w-full transition-transform duration-500 ease-in-out cursor-grab", props.Class)),
		html.AData("pui-carousel-track", ""),
	}

	if props.ID != "" {
		divArgs = append(divArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}

	divArgs = append(divArgs, args...)

	return html.Div(divArgs...)
}

// Item creates a carousel item slide
func Item(props ItemProps, args ...html.DivArg) html.Node {
	divArgs := []html.DivArg{
		html.AClass(classnames.Merge("flex-shrink-0 w-full h-full relative", props.Class)),
		html.AData("pui-carousel-item", ""),
	}

	if props.ID != "" {
		divArgs = append(divArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}

	divArgs = append(divArgs, args...)

	return html.Div(divArgs...)
}

// Previous creates a previous navigation button
func Previous(props PreviousProps, args ...html.ButtonArg) html.Node {
	buttonArgs := []html.ButtonArg{
		html.AClass(classnames.Merge("absolute left-2 top-1/2 transform -translate-y-1/2 p-2 rounded-full bg-black/20 text-white hover:bg-black/40 focus:outline-none", props.Class)),
		html.AData("pui-carousel-prev", ""),
		html.AAria("label", "Previous slide"),
		html.AType("button"),
		lucide.ChevronLeft(html.AClass("h-4 w-4")),
	}

	if props.ID != "" {
		buttonArgs = append(buttonArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		buttonArgs = append(buttonArgs, attr)
	}

	buttonArgs = append(buttonArgs, args...)

	return html.Button(buttonArgs...)
}

// Next creates a next navigation button
func Next(props NextProps, args ...html.ButtonArg) html.Node {
	buttonArgs := []html.ButtonArg{
		html.AClass(classnames.Merge("absolute right-2 top-1/2 transform -translate-y-1/2 p-2 rounded-full bg-black/20 text-white hover:bg-black/40 focus:outline-none", props.Class)),
		html.AData("pui-carousel-next", ""),
		html.AAria("label", "Next slide"),
		html.AType("button"),
		lucide.ChevronRight(html.AClass("h-4 w-4")),
	}

	if props.ID != "" {
		buttonArgs = append(buttonArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		buttonArgs = append(buttonArgs, attr)
	}

	buttonArgs = append(buttonArgs, args...)

	return html.Button(buttonArgs...)
}

// Indicators creates carousel indicators for navigation
func Indicators(props IndicatorsProps, args ...html.DivArg) html.Node {
	divArgs := []html.DivArg{
		html.AClass(classnames.Merge("absolute bottom-4 left-1/2 transform -translate-x-1/2 flex gap-2", props.Class)),
	}

	if props.ID != "" {
		divArgs = append(divArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}

	// Add indicator buttons
	indicatorButtons := make([]html.DivArg, 0, props.Count)
	for i := 0; i < props.Count; i++ {
		buttonClass := "w-3 h-3 rounded-full bg-foreground/30 hover:bg-foreground/50 focus:outline-none transition-colors"
		if i == 0 {
			buttonClass = classnames.Merge(buttonClass, "bg-primary")
		}

		button := html.Button(
			html.AClass(buttonClass),
			html.AData("pui-carousel-indicator", strconv.Itoa(i)),
			html.AData("pui-carousel-active", func() string {
				if i == 0 {
					return "true"
				}

				return "false"
			}()),
			html.AAria("label", fmt.Sprintf("Go to slide %d", i+1)),
			html.AType("button"),
		)
		indicatorButtons = append(indicatorButtons, button)
	}

	divArgs = append(divArgs, indicatorButtons...)
	divArgs = append(divArgs, args...)

	return html.Div(divArgs...)
}

func randomID(prefix string) string {
	buf := make([]byte, 4)
	if _, err := rand.Read(buf); err != nil {
		return prefix + "-id"
	}

	return prefix + "-" + hex.EncodeToString(buf)
}

//go:embed carousel.js
var carouselJS string
