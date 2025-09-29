package form

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/label"
)

type MessageVariant string

const (
	MessageVariantError MessageVariant = "error"
	MessageVariantInfo  MessageVariant = "info"
)

type ItemProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type LabelProps struct {
	ID            string
	Class         string
	Attrs         []html.Global
	For           string
	DisabledClass string
}

type DescriptionProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type MessageProps struct {
	ID      string
	Class   string
	Attrs   []html.Global
	Variant MessageVariant
}

func itemDivArgsFromProps(baseClass string, extra ...string) func(p ItemProps) []html.DivArg {
	return func(p ItemProps) []html.DivArg {
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

func (p ItemProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	for _, a := range itemDivArgsFromProps("space-y-2")(p) {
		a.ApplyDiv(attrs, children)
	}
}

// Item wraps form controls with vertical spacing.
func Item(args ...html.DivArg) html.Node {
	var (
		props ItemProps
		rest  []html.DivArg
	)

	for _, a := range args {
		if v, ok := a.(ItemProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Div(append([]html.DivArg{props}, rest...)...)
}

// ItemFlex aligns form controls horizontally.
func ItemFlex(args ...html.DivArg) html.Node {
	var (
		props ItemProps
		rest  []html.DivArg
	)

	for _, a := range args {
		if v, ok := a.(ItemProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	// Create a new Props with the flex classes
	flexProps := ItemProps{
		ID:    props.ID,
		Class: html.ClassMerge("items-center flex space-x-2", props.Class),
		Attrs: props.Attrs,
	}

	return html.Div(append([]html.DivArg{flexProps}, rest...)...)
}

func labelArgsFromProps(baseClass string, extra ...string) func(p LabelProps) []html.LabelArg {
	return func(p LabelProps) []html.LabelArg {
		classNames := append([]string{baseClass}, extra...)
		classNames = append(classNames, p.Class, p.DisabledClass)

		args := []html.LabelArg{html.AClass(html.ClassMerge(classNames...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		if p.For != "" {
			args = append(args, html.AFor(p.For))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p LabelProps) ApplyLabel(attrs *html.LabelAttrs, children *[]html.Component) {
	for _, a := range labelArgsFromProps("")(p) {
		a.ApplyLabel(attrs, children)
	}
}

// Label proxies to the UI label component so everything stays consistent.
func Label(args ...html.LabelArg) html.Node {
	var (
		props LabelProps
		rest  []html.LabelArg
	)

	for _, a := range args {
		if v, ok := a.(LabelProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	labelProps := label.Props{
		ID:    props.ID,
		Class: html.ClassMerge(props.Class, props.DisabledClass),
		Attrs: props.Attrs,
		For:   props.For,
	}

	return label.Label(append([]html.LabelArg{labelProps}, rest...)...)
}

func pArgsFromProps(baseClass string, extra ...string) func(p DescriptionProps) []html.PArg {
	return func(p DescriptionProps) []html.PArg {
		args := []html.PArg{html.AClass(html.ClassMerge(append([]string{baseClass}, append(extra, p.Class)...)...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p DescriptionProps) ApplyP(attrs *html.PAttrs, children *[]html.Component) {
	for _, a := range pArgsFromProps("text-sm text-muted-foreground")(p) {
		a.ApplyP(attrs, children)
	}
}

// Description renders helper text below an input.
func Description(args ...html.PArg) html.Node {
	var (
		props DescriptionProps
		rest  []html.PArg
	)

	for _, a := range args {
		if v, ok := a.(DescriptionProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.P(append([]html.PArg{props}, rest...)...)
}

func messagePArgsFromProps(baseClass string, extra ...string) func(p MessageProps) []html.PArg {
	return func(p MessageProps) []html.PArg {
		classNames := append([]string{baseClass}, extra...)
		classNames = append(classNames, messageVariantClass(p.Variant), p.Class)

		args := []html.PArg{html.AClass(html.ClassMerge(classNames...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p MessageProps) ApplyP(attrs *html.PAttrs, children *[]html.Component) {
	for _, a := range messagePArgsFromProps("text-[0.8rem] font-medium")(p) {
		a.ApplyP(attrs, children)
	}
}

// Message displays validation feedback.
func Message(args ...html.PArg) html.Node {
	var (
		props MessageProps
		rest  []html.PArg
	)

	for _, a := range args {
		if v, ok := a.(MessageProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.P(append([]html.PArg{props}, rest...)...)
}

func messageVariantClass(variant MessageVariant) string {
	switch variant {
	case MessageVariantError:
		return "text-destructive"
	case MessageVariantInfo:
		return "text-blue-500"
	default:
		return ""
	}
}
