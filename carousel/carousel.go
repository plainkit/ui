package carousel

import (
	"crypto/rand"
	_ "embed"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/internal/styles"
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
		id = randomID("carousel")
	}

	interval := p.Interval
	if interval == 0 {
		interval = 5000
	}

	args := divArgsFromProps(styles.Surface("relative w-full overflow-hidden rounded-3xl"))(p)
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
	args := []html.DivArg{html.AClass(html.ClassMerge("flex h-full w-full gap-6 transition-transform duration-500 ease-in-out", p.Class))}
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
	args := []html.DivArg{html.AClass(html.ClassMerge("relative h-full w-full shrink-0", p.Class))}
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
		html.AClass(html.ClassMerge(styles.InteractiveGhost("absolute left-4 top-1/2 -translate-y-1/2 size-10 rounded-full bg-background/80 shadow-lg backdrop-blur", "hover:bg-background"), p.Class)),
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

	args = append(args, lucide.ChevronLeft(html.AClass("size-5")))

	for _, a := range args {
		a.ApplyButton(attrs, children)
	}
}

func (p NextProps) ApplyButton(attrs *html.ButtonAttrs, children *[]html.Component) {
	args := []html.ButtonArg{
		html.AClass(html.ClassMerge(styles.InteractiveGhost("absolute right-4 top-1/2 -translate-y-1/2 size-10 rounded-full bg-background/80 shadow-lg backdrop-blur", "hover:bg-background"), p.Class)),
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

	args = append(args, lucide.ChevronRight(html.AClass("size-5")))

	for _, a := range args {
		a.ApplyButton(attrs, children)
	}
}

func (p IndicatorsProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	args := []html.DivArg{html.AClass(html.ClassMerge("absolute bottom-5 left-1/2 flex -translate-x-1/2 gap-3", p.Class))}

	if p.ID != "" {
		args = append(args, html.AId(p.ID))
	}

	for _, a := range p.Attrs {
		args = append(args, a)
	}

	// Add indicator buttons
	indicatorButtons := make([]html.Component, 0, p.Count)
	for i := 0; i < p.Count; i++ {
		buttonClass := styles.InteractiveGhost(
			"size-3 rounded-full bg-foreground/30 p-0",
			"data-[pui-carousel-active=true]:bg-primary",
		)

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
