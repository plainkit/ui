package button

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/internal/styles"
)

type Variant string
type Size string
type Type string

type Props struct {
	ID        string
	Class     string
	Attrs     []html.Global
	Variant   Variant
	Size      Size
	FullWidth bool
	Href      string
	Target    string
	Disabled  bool
	Type      Type
	Form      string
}

const (
	VariantDefault     Variant = "default"
	VariantDestructive Variant = "destructive"
	VariantOutline     Variant = "outline"
	VariantSecondary   Variant = "secondary"
	VariantGhost       Variant = "ghost"
	VariantLink        Variant = "link"
)

const (
	SizeDefault Size = "default"
	SizeSm      Size = "sm"
	SizeLg      Size = "lg"
	SizeIcon    Size = "icon"
)

const (
	TypeButton Type = "button"
	TypeReset  Type = "reset"
	TypeSubmit Type = "submit"
)

func buttonArgsFromProps(baseClass string, extra ...string) func(p Props) []html.ButtonArg {
	return func(p Props) []html.ButtonArg {
		className := html.ClassMerge(
			append([]string{baseClass},
				append(extra,
					variantClass(p.Variant),
					sizeClass(p.Size),
					modifierClass(p.FullWidth),
					p.Class,
				)...)...)

		args := []html.ButtonArg{html.AClass(className)}

		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		if p.Type != "" {
			args = append(args, html.AType(string(p.Type)))
		} else {
			args = append(args, html.AType(string(TypeButton)))
		}

		if p.Form != "" {
			args = append(args, html.AForm(p.Form))
		}

		if p.Disabled {
			args = append(args, html.ADisabled())
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func aArgsFromProps(baseClass string, extra ...string) func(p Props) []html.AArg {
	return func(p Props) []html.AArg {
		className := html.ClassMerge(
			append([]string{baseClass},
				append(extra,
					variantClass(p.Variant),
					sizeClass(p.Size),
					modifierClass(p.FullWidth),
					p.Class,
				)...)...)

		args := []html.AArg{
			html.AHref(p.Href),
			html.AClass(className),
		}

		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		if p.Target != "" {
			args = append(args, html.ATarget(p.Target))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

// ApplyButton implements the html.ButtonArg interface for Props
func (p Props) ApplyButton(attrs *html.ButtonAttrs, children *[]html.Component) {
	args := buttonArgsFromProps(
		styles.Interactive(
			"shadow-sm",
			"motion-reduce:transition-none motion-reduce:transform-none",
			"hover:-translate-y-0.5 hover:shadow-lg",
			"active:translate-y-0 active:shadow-sm",
			"ring-offset-background",
		),
		"[&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4",
		"aria-invalid:border-destructive aria-invalid:ring-destructive/30 dark:aria-invalid:ring-destructive/40",
	)(p)

	for _, a := range args {
		a.ApplyButton(attrs, children)
	}
}

// ApplyA implements the html.AArg interface for Props
func (p Props) ApplyA(attrs *html.AAttrs, children *[]html.Component) {
	if p.Href == "" || p.Disabled {
		// If no href or disabled, don't apply anchor attributes
		return
	}

	args := aArgsFromProps(
		styles.Interactive(
			"shadow-sm",
			"motion-reduce:transition-none motion-reduce:transform-none",
			"hover:-translate-y-0.5 hover:shadow-lg",
			"active:translate-y-0 active:shadow-sm",
			"ring-offset-background",
		),
		"[&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4",
		"aria-invalid:border-destructive aria-invalid:ring-destructive/30 dark:aria-invalid:ring-destructive/40",
	)(p)

	for _, a := range args {
		a.ApplyA(attrs, children)
	}
}

// Button renders a button (or link if Href is provided) using the composable pattern.
// Accepts variadic html.ButtonArg arguments, with Props as an optional first argument.
func Button(args ...html.ButtonArg) html.Node {
	var (
		props Props
		rest  []html.ButtonArg
	)

	// Separate Props from other arguments

	for _, a := range args {
		if v, ok := a.(Props); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	// Render as anchor if href is provided and not disabled
	if props.Href != "" && !props.Disabled {
		// Convert button args to anchor args for link rendering
		var anchorArgs []html.AArg

		anchorArgs = append(anchorArgs, props) // Props implements html.AArg

		for _, arg := range rest {
			if aArg, ok := arg.(html.AArg); ok {
				anchorArgs = append(anchorArgs, aArg)
			}
		}

		return html.A(anchorArgs...)
	}

	// Otherwise render as button
	return html.Button(append([]html.ButtonArg{props}, rest...)...)
}

func variantClass(v Variant) string {
	switch v {
	case VariantDestructive:
		return "bg-destructive text-destructive-foreground shadow-md hover:bg-destructive/90 focus-visible:ring-destructive/40 dark:bg-destructive/70"
	case VariantOutline:
		return "border-border/70 bg-background/80 text-foreground shadow-sm hover:bg-background/90 dark:bg-background/40 dark:border-border"
	case VariantSecondary:
		return "bg-secondary/80 text-secondary-foreground shadow-md hover:bg-secondary"
	case VariantGhost:
		return "bg-transparent text-foreground/80 hover:bg-muted/70 hover:text-foreground"
	case VariantLink:
		return "gap-1 border-transparent bg-transparent px-0 py-0 text-primary underline underline-offset-4 decoration-primary/60 hover:decoration-primary"
	default:
		return "bg-gradient-to-r from-primary via-primary/90 to-primary/80 text-primary-foreground shadow-lg hover:from-primary/95 hover:to-primary/90"
	}
}

func sizeClass(s Size) string {
	switch s {
	case SizeSm:
		return "h-8 rounded-md px-3 text-sm has-[>svg]:px-2.5"
	case SizeLg:
		return "h-12 rounded-lg px-6 text-base has-[>svg]:px-5"
	case SizeIcon:
		return "size-10 rounded-xl [&>svg]:size-5"
	default:
		return "h-10 rounded-lg px-4 py-2 has-[>svg]:px-3"
	}
}

func modifierClass(fullWidth bool) string {
	if fullWidth {
		return "w-full"
	}

	return ""
}
