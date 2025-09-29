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

func Breadcrumb(props Props, args ...html.NavArg) html.Node {
	navArgs := []html.NavArg{
		html.AClass(classnames.Merge("flex", props.Class)),
		html.AAria("label", "Breadcrumb"),
	}
	if props.ID != "" {
		navArgs = append(navArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		navArgs = append(navArgs, attr)
	}
	navArgs = append(navArgs, args...)
	return html.Nav(navArgs...)
}

func List(props ListProps, args ...html.OlArg) html.Node {
	olArgs := []html.OlArg{html.AClass(classnames.Merge("flex items-center flex-wrap gap-1 text-sm", props.Class))}
	if props.ID != "" {
		olArgs = append(olArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		olArgs = append(olArgs, attr)
	}
	olArgs = append(olArgs, args...)
	return html.Ol(olArgs...)
}

func Item(props ItemProps, args ...html.LiArg) html.Node {
	liArgs := []html.LiArg{html.AClass(classnames.Merge("flex items-center", props.Class))}
	if props.ID != "" {
		liArgs = append(liArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		liArgs = append(liArgs, attr)
	}
	liArgs = append(liArgs, args...)
	return html.Li(liArgs...)
}

func Link(props LinkProps, args ...html.AArg) html.Node {
	classes := []string{
		"text-muted-foreground hover:text-foreground hover:underline flex items-center gap-1.5 transition-colors",
	}
	if props.IsActive {
		classes = append(classes, "text-foreground")
	}
	if props.Disabled {
		classes = append(classes, "pointer-events-none opacity-60")
	}
	classes = append(classes, props.Class)
	className := classnames.Merge(classes...)

	anchorArgs := []html.AArg{html.AClass(className)}
	if props.ID != "" {
		anchorArgs = append(anchorArgs, html.AId(props.ID))
	}
	if props.Href != "" {
		anchorArgs = append(anchorArgs, html.AHref(props.Href))
	}
	for _, attr := range props.Attrs {
		anchorArgs = append(anchorArgs, attr)
	}
	if props.Disabled {
		anchorArgs = append(anchorArgs, html.AAria("disabled", "true"), html.ATabindex(-1))
	}
	anchorArgs = append(anchorArgs, args...)

	return html.A(anchorArgs...)
}

func Separator(props SeparatorProps, args ...html.SpanArg) html.Node {
	spanArgs := []html.SpanArg{html.AClass(classnames.Merge("mx-2 text-muted-foreground", props.Class))}
	if props.ID != "" {
		spanArgs = append(spanArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		spanArgs = append(spanArgs, attr)
	}
	if props.UseCustom {
		spanArgs = append(spanArgs, args...)
	} else {
		spanArgs = append(spanArgs, lucide.ChevronRight(html.AClass("size-3.5 text-muted-foreground")))
	}
	return html.Span(spanArgs...)
}

func Page(props ItemProps, args ...html.SpanArg) html.Node {
	spanArgs := []html.SpanArg{html.AClass(classnames.Merge("font-medium text-foreground flex items-center gap-1.5", props.Class)), html.AAria("current", "page")}
	if props.ID != "" {
		spanArgs = append(spanArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		spanArgs = append(spanArgs, attr)
	}
	spanArgs = append(spanArgs, args...)
	return html.Span(spanArgs...)
}

func Ellipsis(args ...html.SvgArg) html.Node {
	svgArgs := []html.SvgArg{html.AClass("size-3.5 text-muted-foreground")}
	svgArgs = append(svgArgs, args...)
	return lucide.Ellipsis(svgArgs...)
}
