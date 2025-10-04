package alert

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/internal/styles"
)

type Variant string

type Props struct {
	ID      string
	Class   string
	Attrs   []html.Global
	Variant Variant
}

type TitleProps Props
type DescriptionProps Props

const (
	VariantDefault     Variant = "default"
	VariantDestructive Variant = "destructive"
	VariantSuccess     Variant = "success"
	VariantError       Variant = "error"
	VariantWarning     Variant = "warning"
	VariantInfo        Variant = "info"
)

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

func h5ArgsFromProps(baseClass string, extra ...string) func(p TitleProps) []html.H5Arg {
	return func(p TitleProps) []html.H5Arg {
		args := []html.H5Arg{html.AClass(html.ClassMerge(append([]string{baseClass}, append(extra, p.Class)...)...))}
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
	args := divArgsFromProps(
		styles.SurfaceMuted("relative w-full p-5"),
		"[&>svg]:absolute [&>svg]:left-5 [&>svg]:top-5",
		"[&>svg+div]:translate-y-[-2px] [&:has(svg)]:pl-14",
		variantClass(p.Variant),
	)(p)

	args = append([]html.DivArg{html.ACustom("role", "alert")}, args...)
	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

func (p TitleProps) ApplyH5(attrs *html.H5Attrs, children *[]html.Component) {
	for _, a := range h5ArgsFromProps(styles.DisplayHeading("mb-1 text-lg"))(p) {
		a.ApplyH5(attrs, children)
	}
}

func (p DescriptionProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	for _, a := range divArgsFromProps(styles.SubtleText("[&_p]:leading-relaxed"))(Props(p)) {
		a.ApplyDiv(attrs, children)
	}
}

func Alert(args ...html.DivArg) html.Node {
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

	return html.Div(append([]html.DivArg{props}, rest...)...)
}

func Title(args ...html.H5Arg) html.Node {
	var (
		props TitleProps
		rest  []html.H5Arg
	)

	for _, a := range args {
		if v, ok := a.(TitleProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.H5(append([]html.H5Arg{props}, rest...)...)
}

func Description(args ...html.DivArg) html.Node {
	var (
		props DescriptionProps
		rest  []html.DivArg
	)

	for _, a := range args {
		if v, ok := a.(DescriptionProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Div(append([]html.DivArg{props}, rest...)...)
}

func variantClass(v Variant) string {
	switch v {
	case VariantDestructive, VariantError:
		return "border-destructive/60 bg-destructive/15 text-destructive"
	case VariantSuccess:
		return "border-emerald-400/50 bg-emerald-500/15 text-emerald-900 dark:border-emerald-500/60 dark:bg-emerald-500/20 dark:text-emerald-100"
	case VariantWarning:
		return "border-amber-400/60 bg-amber-500/15 text-amber-900 dark:border-amber-500/50 dark:bg-amber-400/20 dark:text-amber-50"
	case VariantInfo:
		return "border-sky-400/60 bg-sky-500/15 text-sky-900 dark:border-sky-500/60 dark:bg-sky-500/20 dark:text-sky-100"
	case VariantDefault:
		return "border-border/60 bg-card/70 text-card-foreground"
	default:
		return "border-border/60 bg-card/70 text-card-foreground"
	}
}
