package card

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/internal/classnames"
)

type Props struct {
	ID    string
	Class string
	Attrs []html.Global
}

type HeaderProps Props

type TitleProps Props

type DescriptionProps Props

type ContentProps Props

type FooterProps Props

func Card(props Props, args ...html.DivArg) html.Node {
	divArgs := []html.DivArg{html.AClass(defaultClasses("w-full rounded-lg border bg-card text-card-foreground shadow-xs", props.Class))}
	if props.ID != "" {
		divArgs = append(divArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}
	divArgs = append(divArgs, args...)
	return html.Div(divArgs...)
}

func Header(props HeaderProps, args ...html.DivArg) html.Node {
	divArgs := []html.DivArg{html.AClass(defaultClasses("flex flex-col space-y-1.5 p-6 pb-0", props.Class))}
	if props.ID != "" {
		divArgs = append(divArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}
	divArgs = append(divArgs, args...)
	return html.Div(divArgs...)
}

func Title(props TitleProps, args ...html.H3Arg) html.Node {
	hargs := []html.H3Arg{html.AClass(defaultClasses("text-lg font-semibold leading-none tracking-tight", props.Class))}
	if props.ID != "" {
		hargs = append(hargs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		hargs = append(hargs, attr)
	}
	hargs = append(hargs, args...)
	return html.H3(hargs...)
}

func Description(props DescriptionProps, args ...html.PArg) html.Node {
	pargs := []html.PArg{html.AClass(defaultClasses("text-sm text-muted-foreground", props.Class))}
	if props.ID != "" {
		pargs = append(pargs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		pargs = append(pargs, attr)
	}
	pargs = append(pargs, args...)
	return html.P(pargs...)
}

func Content(props ContentProps, args ...html.DivArg) html.Node {
	divArgs := []html.DivArg{html.AClass(defaultClasses("p-6", props.Class))}
	if props.ID != "" {
		divArgs = append(divArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}
	divArgs = append(divArgs, args...)
	return html.Div(divArgs...)
}

func Footer(props FooterProps, args ...html.DivArg) html.Node {
	divArgs := []html.DivArg{html.AClass(defaultClasses("flex items-center p-6 pt-0", props.Class))}
	if props.ID != "" {
		divArgs = append(divArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}
	divArgs = append(divArgs, args...)
	return html.Div(divArgs...)
}

func defaultClasses(base, extra string) string {
	return classnames.Merge(base, extra)
}
