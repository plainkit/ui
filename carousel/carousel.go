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

func divArgsFromProps(baseClass string, extra ...string) func(p Props) []html.DivArg {
	return func(p Props) []html.DivArg {
		args := []html.DivArg{html.AClass(classnames.Merge(append([]string{baseClass}, append(extra, p.Class)...)...))}
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
		id = randomID("carousel")
	}

	interval := p.Interval
	if interval == 0 {
		interval = 5000
	}

	args := divArgsFromProps("relative overflow-hidden w-full")(p)
	args = append([]html.DivArg{
		html.AId(id),
		html.AData("pui-carousel", ""),
		html.AData("pui-carousel-current", "0"),
		html.AData("pui-carousel-autoplay", strconv.FormatBool(p.Autoplay)),
		html.AData("pui-carousel-interval", fmt.Sprintf("%d", interval)),
		html.AData("pui-carousel-loop", strconv.FormatBool(p.Loop)),
	}, args...)

	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

func (p ContentProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	args := []html.DivArg{html.AClass(classnames.Merge("flex h-full w-full transition-transform duration-500 ease-in-out cursor-grab", p.Class))}
	args = append(args, html.AData("pui-carousel-track", ""))

	if p.ID != "" {
		args = append(args, html.AId(p.ID))
	}

	for _, a := range p.Attrs {
		args = append(args, a)
	}

	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

func (p ItemProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	args := []html.DivArg{html.AClass(classnames.Merge("flex-shrink-0 w-full h-full relative", p.Class))}
	args = append(args, html.AData("pui-carousel-item", ""))

	if p.ID != "" {
		args = append(args, html.AId(p.ID))
	}

	for _, a := range p.Attrs {
		args = append(args, a)
	}

	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

func (p PreviousProps) ApplyButton(attrs *html.ButtonAttrs, children *[]html.Component) {
	args := []html.ButtonArg{
		html.AClass(classnames.Merge("absolute left-2 top-1/2 transform -translate-y-1/2 p-2 rounded-full bg-black/20 text-white hover:bg-black/40 focus:outline-none", p.Class)),
		html.AData("pui-carousel-prev", ""),
		html.AAria("label", "Previous slide"),
		html.AType("button"),
	}

	if p.ID != "" {
		args = append(args, html.AId(p.ID))
	}

	for _, a := range p.Attrs {
		args = append(args, a)
	}

	args = append(args, lucide.ChevronLeft(html.AClass("h-4 w-4")))

	for _, a := range args {
		a.ApplyButton(attrs, children)
	}
}

func (p NextProps) ApplyButton(attrs *html.ButtonAttrs, children *[]html.Component) {
	args := []html.ButtonArg{
		html.AClass(classnames.Merge("absolute right-2 top-1/2 transform -translate-y-1/2 p-2 rounded-full bg-black/20 text-white hover:bg-black/40 focus:outline-none", p.Class)),
		html.AData("pui-carousel-next", ""),
		html.AAria("label", "Next slide"),
		html.AType("button"),
	}

	if p.ID != "" {
		args = append(args, html.AId(p.ID))
	}

	for _, a := range p.Attrs {
		args = append(args, a)
	}

	args = append(args, lucide.ChevronRight(html.AClass("h-4 w-4")))

	for _, a := range args {
		a.ApplyButton(attrs, children)
	}
}

func (p IndicatorsProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	args := []html.DivArg{html.AClass(classnames.Merge("absolute bottom-4 left-1/2 transform -translate-x-1/2 flex gap-2", p.Class))}

	if p.ID != "" {
		args = append(args, html.AId(p.ID))
	}

	for _, a := range p.Attrs {
		args = append(args, a)
	}

	// Add indicator buttons
	indicatorButtons := make([]html.Component, 0, p.Count)
	for i := 0; i < p.Count; i++ {
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

	*children = append(*children, indicatorButtons...)

	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

// Carousel renders a carousel container for sliding content
func Carousel(args ...html.DivArg) html.Node {
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

	return html.Div(append([]html.DivArg{props}, rest...)...).WithAssets("", carouselJS, "ui-carousel")
}

// Content creates the carousel track that contains the items
func Content(args ...html.DivArg) html.Node {
	var (
		props ContentProps
		rest  []html.DivArg
	)

	for _, a := range args {
		if v, ok := a.(ContentProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Div(append([]html.DivArg{props}, rest...)...)
}

// Item creates a carousel item slide
func Item(args ...html.DivArg) html.Node {
	var (
		props ItemProps
		rest  []html.DivArg
	)

	for _, a := range args {
		if v, ok := a.(ItemProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Div(append([]html.DivArg{props}, rest...)...)
}

// Previous creates a previous navigation button
func Previous(args ...html.ButtonArg) html.Node {
	var (
		props PreviousProps
		rest  []html.ButtonArg
	)

	for _, a := range args {
		if v, ok := a.(PreviousProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Button(append([]html.ButtonArg{props}, rest...)...)
}

// Next creates a next navigation button
func Next(args ...html.ButtonArg) html.Node {
	var (
		props NextProps
		rest  []html.ButtonArg
	)

	for _, a := range args {
		if v, ok := a.(NextProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Button(append([]html.ButtonArg{props}, rest...)...)
}

// Indicators creates carousel indicators for navigation
func Indicators(args ...html.DivArg) html.Node {
	var (
		props IndicatorsProps
		rest  []html.DivArg
	)

	for _, a := range args {
		if v, ok := a.(IndicatorsProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Div(append([]html.DivArg{props}, rest...)...)
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
