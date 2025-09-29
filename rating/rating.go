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

const (
	StyleStar  Style = "star"
	StyleHeart Style = "heart"
	StyleEmoji Style = "emoji"
)

// Rating wraps interactive rating items and manages hidden input value syncing.
func Rating(props Props, args ...html.DivArg) html.Node {
	precision := props.Precision
	if precision <= 0 {
		precision = 1
	}

	dataset := []html.Global{
		html.AData("pui-rating-component", ""),
		html.AData("pui-rating-initial-value", fmt.Sprintf("%.2f", props.Value)),
		html.AData("pui-rating-precision", fmt.Sprintf("%.2f", precision)),
		html.AData("pui-rating-readonly", strconv.FormatBool(props.ReadOnly)),
		html.AData("pui-rating-onlyinteger", strconv.FormatBool(props.OnlyInteger)),
	}
	if props.Name != "" {
		dataset = append(dataset, html.AData("pui-rating-name", props.Name))
	}

	divArgs := []html.DivArg{html.AClass(classnames.Merge("flex flex-col items-start gap-1", props.Class))}
	if props.ID != "" {
		divArgs = append(divArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}

	for _, data := range dataset {
		divArgs = append(divArgs, data)
	}

	if props.Name != "" {
		hiddenArgs := []html.InputArg{
			html.AType("hidden"),
			html.AName(props.Name),
			html.AValue(fmt.Sprintf("%.2f", props.Value)),
			html.AData("pui-rating-input", ""),
		}
		if props.Form != "" {
			hiddenArgs = append(hiddenArgs, html.AForm(props.Form))
		}

		hidden := html.Input(hiddenArgs...)
		divArgs = append(divArgs, hidden)
	}

	divArgs = append(divArgs, args...)

	node := html.Div(divArgs...)

	return node.WithAssets("", ratingJS, "ui-rating")
}

// Group arranges rating items in a single row.
func Group(props GroupProps, args ...html.DivArg) html.Node {
	groupArgs := []html.DivArg{html.AClass(classnames.Merge("flex flex-row items-center gap-1", props.Class))}
	if props.ID != "" {
		groupArgs = append(groupArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		groupArgs = append(groupArgs, attr)
	}

	groupArgs = append(groupArgs, args...)

	return html.Div(groupArgs...)
}

// Item renders an individual rating icon with layered fill for partial states.
func Item(props ItemProps) html.Node {
	style := props.Style
	if style == "" {
		style = StyleStar
	}

	itemArgs := []html.DivArg{
		html.AClass(classnames.Merge("relative transition-opacity cursor-pointer", colorClass(style), props.Class)),
		html.AData("pui-rating-item", ""),
		html.AData("pui-rating-value", strconv.Itoa(props.Value)),
	}
	if props.ID != "" {
		itemArgs = append(itemArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		itemArgs = append(itemArgs, attr)
	}

	background := html.Div(
		html.AClass("opacity-30"),
		ratingIcon(style, false, float64(props.Value)),
	)

	foreground := html.Div(
		html.AClass("absolute inset-0 overflow-hidden w-0"),
		html.AData("pui-rating-item-foreground", ""),
		ratingIcon(style, true, float64(props.Value)),
	)

	itemArgs = append(itemArgs, background, foreground)

	return html.Div(itemArgs...)
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
