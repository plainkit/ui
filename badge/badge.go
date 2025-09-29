package badge

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/internal/classnames"
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

// Badge renders a badge span with optional props and children supplied as span arguments.
// Pass a Props struct (zero value is fine) followed by `html.T(...)`, `html.Span(...)`, etc.
func Badge(props Props, args ...html.SpanArg) html.Node {
	className := classnames.Merge(
		"inline-flex items-center justify-center rounded-md border px-2 py-0.5 text-xs font-medium w-fit whitespace-nowrap shrink-0 [&>svg]:size-3 gap-1 [&>svg]:pointer-events-none",
		"focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px]",
		"aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive",
		"transition-[color,box-shadow] overflow-hidden",
		variantClasses(props.Variant),
		props.Class,
	)

	spanArgs := []html.SpanArg{html.AClass(className)}
	if props.ID != "" {
		spanArgs = append(spanArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		spanArgs = append(spanArgs, attr)
	}
	spanArgs = append(spanArgs, args...)

	return html.Span(spanArgs...)
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
