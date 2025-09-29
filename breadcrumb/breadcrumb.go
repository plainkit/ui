package breadcrumb

import (
	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/internal/classnames"
)

type Props struct {
	ID    string
	Class string
	Attrs []html.Global
}

type ListProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type ItemProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type LinkProps struct {
	ID       string
	Class    string
	Attrs    []html.Global
	Href     string
	IsActive bool
	Disabled bool
}

type SeparatorProps struct {
	ID        string
	Class     string
	Attrs     []html.Global
	UseCustom bool
}

func navArgsFromProps(baseClass string, extra ...string) func(p Props) []html.NavArg {
	return func(p Props) []html.NavArg {
		args := []html.NavArg{
			html.AClass(classnames.Merge(append([]string{baseClass}, append(extra, p.Class)...)...)),
			html.AAria("label", "Breadcrumb"),
		}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p Props) ApplyNav(attrs *html.NavAttrs, children *[]html.Component) {
	for _, a := range navArgsFromProps("flex")(p) {
		a.ApplyNav(attrs, children)
	}
}

func Breadcrumb(args ...html.NavArg) html.Node {
	var (
		props Props
		rest  []html.NavArg
	)

	for _, a := range args {
		if v, ok := a.(Props); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Nav(append([]html.NavArg{props}, rest...)...)
}

func olArgsFromProps(baseClass string, extra ...string) func(p ListProps) []html.OlArg {
	return func(p ListProps) []html.OlArg {
		args := []html.OlArg{html.AClass(classnames.Merge(append([]string{baseClass}, append(extra, p.Class)...)...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p ListProps) ApplyOl(attrs *html.OlAttrs, children *[]html.Component) {
	for _, a := range olArgsFromProps("flex items-center flex-wrap gap-1 text-sm")(p) {
		a.ApplyOl(attrs, children)
	}
}

func List(args ...html.OlArg) html.Node {
	var (
		props ListProps
		rest  []html.OlArg
	)

	for _, a := range args {
		if v, ok := a.(ListProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Ol(append([]html.OlArg{props}, rest...)...)
}

func liArgsFromProps(baseClass string, extra ...string) func(p ItemProps) []html.LiArg {
	return func(p ItemProps) []html.LiArg {
		args := []html.LiArg{html.AClass(classnames.Merge(append([]string{baseClass}, append(extra, p.Class)...)...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p ItemProps) ApplyLi(attrs *html.LiAttrs, children *[]html.Component) {
	for _, a := range liArgsFromProps("flex items-center")(p) {
		a.ApplyLi(attrs, children)
	}
}

func Item(args ...html.LiArg) html.Node {
	var (
		props ItemProps
		rest  []html.LiArg
	)

	for _, a := range args {
		if v, ok := a.(ItemProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Li(append([]html.LiArg{props}, rest...)...)
}

func aArgsFromProps(baseClass string, extra ...string) func(p LinkProps) []html.AArg {
	return func(p LinkProps) []html.AArg {
		classNames := append([]string{baseClass}, extra...)
		if p.IsActive {
			classNames = append(classNames, "text-foreground")
		}

		if p.Disabled {
			classNames = append(classNames, "pointer-events-none opacity-60")
		}

		classNames = append(classNames, p.Class)

		args := []html.AArg{html.AClass(classnames.Merge(classNames...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		if p.Href != "" {
			args = append(args, html.AHref(p.Href))
		}

		if p.Disabled {
			args = append(args, html.AAria("disabled", "true"), html.ATabindex(-1))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p LinkProps) ApplyA(attrs *html.AAttrs, children *[]html.Component) {
	for _, a := range aArgsFromProps("text-muted-foreground hover:text-foreground hover:underline flex items-center gap-1.5 transition-colors")(p) {
		a.ApplyA(attrs, children)
	}
}

func Link(args ...html.AArg) html.Node {
	var (
		props LinkProps
		rest  []html.AArg
	)

	for _, a := range args {
		if v, ok := a.(LinkProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.A(append([]html.AArg{props}, rest...)...)
}

func spanArgsFromProps(baseClass string, extra ...string) func(p SeparatorProps) []html.SpanArg {
	return func(p SeparatorProps) []html.SpanArg {
		args := []html.SpanArg{html.AClass(classnames.Merge(append([]string{baseClass}, append(extra, p.Class)...)...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p SeparatorProps) ApplySpan(attrs *html.SpanAttrs, children *[]html.Component) {
	for _, a := range spanArgsFromProps("mx-2 text-muted-foreground")(p) {
		a.ApplySpan(attrs, children)
	}
}

func Separator(args ...html.SpanArg) html.Node {
	var (
		props SeparatorProps
		rest  []html.SpanArg
	)

	for _, a := range args {
		if v, ok := a.(SeparatorProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	if props.UseCustom {
		return html.Span(append([]html.SpanArg{props}, rest...)...)
	} else {
		return html.Span(append([]html.SpanArg{props}, lucide.ChevronRight(html.AClass("size-3.5 text-muted-foreground")))...)
	}
}

func pageSpanArgsFromProps(baseClass string, extra ...string) func(p ItemProps) []html.SpanArg {
	return func(p ItemProps) []html.SpanArg {
		args := []html.SpanArg{
			html.AClass(classnames.Merge(append([]string{baseClass}, append(extra, p.Class)...)...)),
			html.AAria("current", "page"),
		}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p ItemProps) ApplySpan(attrs *html.SpanAttrs, children *[]html.Component) {
	for _, a := range pageSpanArgsFromProps("font-medium text-foreground flex items-center gap-1.5")(p) {
		a.ApplySpan(attrs, children)
	}
}

func Page(args ...html.SpanArg) html.Node {
	var (
		props ItemProps
		rest  []html.SpanArg
	)

	for _, a := range args {
		if v, ok := a.(ItemProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Span(append([]html.SpanArg{props}, rest...)...)
}

func Ellipsis(args ...html.SvgArg) html.Node {
	svgArgs := []html.SvgArg{html.AClass("size-3.5 text-muted-foreground")}
	svgArgs = append(svgArgs, args...)

	return lucide.Ellipsis(svgArgs...)
}
