package alert

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/internal/classnames"
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

func h5ArgsFromProps(baseClass string, extra ...string) func(p TitleProps) []html.H5Arg {
	return func(p TitleProps) []html.H5Arg {
		args := []html.H5Arg{html.AClass(classnames.Merge(append([]string{baseClass}, append(extra, p.Class)...)...))}
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
		"relative w-full p-4",
		"[&>svg]:absolute [&>svg]:left-4 [&>svg]:top-4",
		"[&>svg+div]:translate-y-[-3px] [&:has(svg)]:pl-11",
		"rounded-lg border",
		variantClass(p.Variant),
	)(p)

	args = append([]html.DivArg{html.ACustom("role", "alert")}, args...)
	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

func (p TitleProps) ApplyH5(attrs *html.H5Attrs, children *[]html.Component) {
	for _, a := range h5ArgsFromProps("mb-1 font-medium leading-none tracking-tight")(p) {
		a.ApplyH5(attrs, children)
	}
}

func (p DescriptionProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	for _, a := range divArgsFromProps("[&_p]:leading-relaxed text-sm")(Props(p)) {
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
		return "border-destructive text-destructive"
	case VariantSuccess:
		return "border-green-200 text-green-900 dark:border-green-800 dark:text-green-100"
	case VariantWarning:
		return "border-yellow-200 text-yellow-900 dark:border-yellow-800 dark:text-yellow-100"
	case VariantInfo:
		return "border-blue-200 text-blue-900 dark:border-blue-800 dark:text-blue-100"
	case VariantDefault:
		return "border-border text-foreground"
	default:
		return "border-border text-foreground"
	}
}
