package badge

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/internal/styles"
)

type Variant string

const (
	VariantDefault     Variant = "default"
	VariantSecondary   Variant = "secondary"
	VariantDestructive Variant = "destructive"
	VariantOutline     Variant = "outline"
)

type Props struct {
	ID      string
	Class   string
	Attrs   []html.Global
	Variant Variant
}

func spanArgsFromProps(baseClass string, extra ...string) func(p Props) []html.SpanArg {
	return func(p Props) []html.SpanArg {
		className := html.ClassMerge(
			append([]string{baseClass},
				append(extra,
					variantClasses(p.Variant),
					p.Class,
				)...)...)

		args := []html.SpanArg{html.AClass(className)}

		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

// ApplySpan implements the html.SpanArg interface for Props
func (p Props) ApplySpan(attrs *html.SpanAttrs, children *[]html.Component) {
	args := spanArgsFromProps(
		styles.Tag(
			"w-fit whitespace-nowrap shrink-0",
			"[&>svg]:size-3 [&>svg]:pointer-events-none",
			"transition-colors",
		),
		"aria-invalid:border-destructive aria-invalid:ring-destructive/30 dark:aria-invalid:ring-destructive/40",
	)(p)

	for _, a := range args {
		a.ApplySpan(attrs, children)
	}
}

// Badge renders a badge span using the composable pattern.
// Accepts variadic html.SpanArg arguments, with Props as an optional first argument.
func Badge(args ...html.SpanArg) html.Node {
	var (
		props Props
		rest  []html.SpanArg
	)

	// Separate Props from other arguments
	for _, a := range args {
		if v, ok := a.(Props); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Span(append([]html.SpanArg{props}, rest...)...)
}

func variantClasses(v Variant) string {
	switch v {
	case VariantDestructive:
		return "border-transparent bg-destructive/85 text-destructive-foreground [a&]:hover:bg-destructive focus-visible:ring-destructive/40"
	case VariantOutline:
		return "border-border/60 bg-transparent text-foreground/80 [a&]:hover:bg-muted/70 [a&]:hover:text-foreground"
	case VariantSecondary:
		return "border-border/50 bg-muted/70 text-foreground/80 [a&]:hover:bg-muted"
	default:
		return "border-transparent bg-primary/90 text-primary-foreground [a&]:hover:bg-primary"
	}
}
