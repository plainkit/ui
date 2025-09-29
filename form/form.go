package form

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/internal/classnames"
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

// Item wraps form controls with vertical spacing.
func Item(props ItemProps, args ...html.DivArg) html.Node {
	divArgs := []html.DivArg{html.AClass(classnames.Merge("space-y-2", props.Class))}
	if props.ID != "" {
		divArgs = append(divArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}

	divArgs = append(divArgs, args...)

	return html.Div(divArgs...)
}

// ItemFlex aligns form controls horizontally.
func ItemFlex(props ItemProps, args ...html.DivArg) html.Node {
	divArgs := []html.DivArg{html.AClass(classnames.Merge("items-center flex space-x-2", props.Class))}
	if props.ID != "" {
		divArgs = append(divArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}

	divArgs = append(divArgs, args...)

	return html.Div(divArgs...)
}

// Label proxies to the UI label component so everything stays consistent.
func Label(props LabelProps, args ...html.LabelArg) html.Node {
	labelProps := label.Props{
		ID:    props.ID,
		Class: classnames.Merge(props.Class, props.DisabledClass),
		Attrs: props.Attrs,
		For:   props.For,
	}

	return label.Label(labelProps, args...)
}

// Description renders helper text below an input.
func Description(props DescriptionProps, args ...html.PArg) html.Node {
	pArgs := []html.PArg{html.AClass(classnames.Merge("text-sm text-muted-foreground", props.Class))}
	if props.ID != "" {
		pArgs = append(pArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		pArgs = append(pArgs, attr)
	}

	pArgs = append(pArgs, args...)

	return html.P(pArgs...)
}

// Message displays validation feedback.
func Message(props MessageProps, args ...html.PArg) html.Node {
	pArgs := []html.PArg{html.AClass(classnames.Merge("text-[0.8rem] font-medium", messageVariantClass(props.Variant), props.Class))}
	if props.ID != "" {
		pArgs = append(pArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		pArgs = append(pArgs, attr)
	}

	pArgs = append(pArgs, args...)

	return html.P(pArgs...)
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
