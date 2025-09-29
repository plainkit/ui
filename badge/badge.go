package badge

import (
	"github.com/plainkit/html"
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
		"inline-flex items-center justify-center rounded-md border px-2 py-0.5 text-xs font-medium w-fit whitespace-nowrap shrink-0 [&>svg]:size-3 gap-1 [&>svg]:pointer-events-none",
		"focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px]",
		"aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive",
		"transition-[color,box-shadow] overflow-hidden",
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
		return "border-transparent bg-destructive text-white [a&]:hover:bg-destructive/90 focus-visible:ring-destructive/20 dark:focus-visible:ring-destructive/40 dark:bg-destructive/60"
	case VariantOutline:
		return "text-foreground [a&]:hover:bg-accent [a&]:hover:text-accent-foreground"
	case VariantSecondary:
		return "border-transparent bg-secondary text-secondary-foreground [a&]:hover:bg-secondary/90"
	default:
		return "border-transparent bg-primary text-primary-foreground [a&]:hover:bg-primary/90"
	}
}
