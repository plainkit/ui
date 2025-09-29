package rating

import (
	_ "embed"
	"fmt"
	"strconv"

	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/internal/classnames"
)

type Style string

type Props struct {
	ID          string
	Class       string
	Attrs       []html.Global
	Value       float64
	ReadOnly    bool
	Precision   float64
	Name        string
	Form        string
	OnlyInteger bool
}

type GroupProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type ItemProps struct {
	ID    string
	Class string
	Attrs []html.Global
	Value int
	Style Style
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
	precision := p.Precision
	if precision <= 0 {
		precision = 1
	}

	dataset := []html.Global{
		html.AData("pui-rating-component", ""),
		html.AData("pui-rating-initial-value", fmt.Sprintf("%.2f", p.Value)),
		html.AData("pui-rating-precision", fmt.Sprintf("%.2f", precision)),
		html.AData("pui-rating-readonly", strconv.FormatBool(p.ReadOnly)),
		html.AData("pui-rating-onlyinteger", strconv.FormatBool(p.OnlyInteger)),
	}
	if p.Name != "" {
		dataset = append(dataset, html.AData("pui-rating-name", p.Name))
	}

	args := divArgsFromProps("flex flex-col items-start gap-1")(p)
	for _, data := range dataset {
		args = append(args, data)
	}

	if p.Name != "" {
		hiddenArgs := []html.InputArg{
			html.AType("hidden"),
			html.AName(p.Name),
			html.AValue(fmt.Sprintf("%.2f", p.Value)),
			html.AData("pui-rating-input", ""),
		}
		if p.Form != "" {
			hiddenArgs = append(hiddenArgs, html.AForm(p.Form))
		}

		hidden := html.Input(hiddenArgs...)
		*children = append(*children, hidden)
	}

	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

func (p GroupProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	args := []html.DivArg{html.AClass(classnames.Merge("flex flex-row items-center gap-1", p.Class))}
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
	style := p.Style
	if style == "" {
		style = StyleStar
	}

	args := []html.DivArg{
		html.AClass(classnames.Merge("relative transition-opacity cursor-pointer", colorClass(style), p.Class)),
		html.AData("pui-rating-item", ""),
		html.AData("pui-rating-value", strconv.Itoa(p.Value)),
	}
	if p.ID != "" {
		args = append(args, html.AId(p.ID))
	}

	for _, a := range p.Attrs {
		args = append(args, a)
	}

	background := html.Div(
		html.AClass("opacity-30"),
		ratingIcon(style, false, float64(p.Value)),
	)

	foreground := html.Div(
		html.AClass("absolute inset-0 overflow-hidden w-0"),
		html.AData("pui-rating-item-foreground", ""),
		ratingIcon(style, true, float64(p.Value)),
	)

	*children = append(*children, background, foreground)

	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

const (
	StyleStar  Style = "star"
	StyleHeart Style = "heart"
	StyleEmoji Style = "emoji"
)

// Rating wraps interactive rating items and manages hidden input value syncing.
func Rating(args ...html.DivArg) html.Node {
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

	return html.Div(append([]html.DivArg{props}, rest...)...).WithAssets("", ratingJS, "ui-rating")
}

// Group arranges rating items in a single row.
func Group(args ...html.DivArg) html.Node {
	var (
		props GroupProps
		rest  []html.DivArg
	)

	for _, a := range args {
		if v, ok := a.(GroupProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Div(append([]html.DivArg{props}, rest...)...)
}

// Item renders an individual rating icon with layered fill for partial states.
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

func colorClass(style Style) string {
	switch style {
	case StyleHeart:
		return "text-destructive"
	case StyleEmoji:
		return "text-yellow-500"
	default:
		return "text-yellow-400"
	}
}

func ratingIcon(style Style, filled bool, value float64) html.Node {
	if style == StyleEmoji {
		if filled {
			switch {
			case value <= 1:
				return lucide.Angry()
			case value <= 2:
				return lucide.Frown()
			case value <= 3:
				return lucide.Meh()
			case value <= 4:
				return lucide.Smile()
			default:
				return lucide.Laugh()
			}
		}

		return lucide.Meh()
	}

	iconArgs := []html.SvgArg{}
	if filled {
		iconArgs = append(iconArgs, html.AFill("currentColor"), html.AStroke("none"))
	}

	switch style {
	case StyleHeart:
		return lucide.Heart(iconArgs...)
	default:
		return lucide.Star(iconArgs...)
	}
}

//go:embed rating.js
var ratingJS string
