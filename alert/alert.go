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
)

func Alert(props Props, args ...html.DivArg) html.Node {
	className := classnames.Merge(
		"relative w-full p-4",
		"[&>svg]:absolute [&>svg]:left-4 [&>svg]:top-4",
		"[&>svg+div]:translate-y-[-3px] [&:has(svg)]:pl-11",
		"rounded-lg border",
		variantClass(props.Variant),
		props.Class,
	)

	divArgs := []html.DivArg{
		html.AClass(className),
		html.ACustom("role", "alert"),
	}
	if props.ID != "" {
		divArgs = append(divArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}
	divArgs = append(divArgs, args...)

	return html.Div(divArgs...)
}

func Title(props TitleProps, args ...html.H5Arg) html.Node {
	className := classnames.Merge(
		"mb-1 font-medium leading-none tracking-tight",
		props.Class,
	)

	hArgs := []html.H5Arg{html.AClass(className)}
	if props.ID != "" {
		hArgs = append(hArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		hArgs = append(hArgs, attr)
	}
	hArgs = append(hArgs, args...)

	return html.H5(hArgs...)
}

func Description(props DescriptionProps, args ...html.DivArg) html.Node {
	className := classnames.Merge(
		"[&_p]:leading-relaxed text-sm",
		props.Class,
	)

	dArgs := []html.DivArg{html.AClass(className)}
	if props.ID != "" {
		dArgs = append(dArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		dArgs = append(dArgs, attr)
	}
	dArgs = append(dArgs, args...)

	return html.Div(dArgs...)
}

func variantClass(v Variant) string {
	switch v {
	case VariantDestructive:
		return "border-destructive text-destructive"
	default:
		return "border-border text-foreground"
	}
}
