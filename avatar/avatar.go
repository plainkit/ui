package avatar

import (
	"fmt"

	"github.com/plainkit/html"
	"github.com/plainkit/ui/internal/classnames"
)

type Size string
type GroupSpacing string

type Props struct {
	ID      string
	Class   string
	Attrs   []html.Global
	Size    Size
	InGroup bool
}

type ImageProps struct {
	ID    string
	Class string
	Attrs []html.Global
	Alt   string
	Src   string
}

type FallbackProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type GroupProps struct {
	ID      string
	Class   string
	Attrs   []html.Global
	Spacing GroupSpacing
}

const (
	SizeSm Size = "sm"
	SizeMd Size = "md"
	SizeLg Size = "lg"
)

const (
	GroupSpacingSm GroupSpacing = "sm"
	GroupSpacingMd GroupSpacing = "md"
	GroupSpacingLg GroupSpacing = "lg"
)

func Avatar(props Props, args ...html.DivArg) html.Node {
	className := classnames.Merge(
		"relative inline-flex shrink-0 items-center justify-center overflow-hidden",
		"size-12 text-base",
		"data-[pui-avatar-size=sm]:size-8 data-[pui-avatar-size=sm]:text-xs",
		"data-[pui-avatar-size=md]:size-12 data-[pui-avatar-size=md]:text-base",
		"data-[pui-avatar-size=lg]:size-16 data-[pui-avatar-size=lg]:text-xl",
		"rounded-full bg-muted",
		"data-[pui-avatar-in-group=true]:ring-2 data-[pui-avatar-in-group=true]:ring-background",
		props.Class,
	)

	divArgs := []html.DivArg{
		html.AClass(className),
		html.AData("pui-avatar", ""),
	}
	if props.ID != "" {
		divArgs = append(divArgs, html.AId(props.ID))
	}
	if props.Size != "" {
		divArgs = append(divArgs, html.AData("pui-avatar-size", string(props.Size)))
	}
	if props.InGroup {
		divArgs = append(divArgs, html.AData("pui-avatar-in-group", "true"))
	}
	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}
	divArgs = append(divArgs, args...)

	return html.Div(divArgs...)
}

func Image(props ImageProps, args ...html.ImgArg) html.Node {
	className := classnames.Merge(
		"absolute inset-0 w-full h-full",
		"object-cover",
		"z-10",
		props.Class,
	)

	imgArgs := []html.ImgArg{
		html.AClass(className),
		html.AData("pui-avatar-image", ""),
	}
	if props.ID != "" {
		imgArgs = append(imgArgs, html.AId(props.ID))
	}
	if props.Src != "" {
		imgArgs = append(imgArgs, html.ASrc(props.Src))
	}
	imgArgs = append(imgArgs, html.AAlt(props.Alt))
	for _, attr := range props.Attrs {
		imgArgs = append(imgArgs, attr)
	}
	imgArgs = append(imgArgs, args...)

	return html.Img(imgArgs...)
}

func Fallback(props FallbackProps, args ...html.SpanArg) html.Node {
	className := classnames.Merge(
		"font-medium text-muted-foreground",
		props.Class,
	)

	spanArgs := []html.SpanArg{
		html.AClass(className),
		html.AData("pui-avatar-fallback", ""),
	}
	if props.ID != "" {
		spanArgs = append(spanArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		spanArgs = append(spanArgs, attr)
	}
	spanArgs = append(spanArgs, args...)

	return html.Span(spanArgs...)
}

func Group(props GroupProps, args ...html.DivArg) html.Node {
	if props.Spacing == "" {
		props.Spacing = GroupSpacingMd
	}

	className := classnames.Merge(
		"flex items-center",
		"data-[pui-avatar-spacing=sm]:-space-x-1",
		"data-[pui-avatar-spacing=md]:-space-x-2",
		"data-[pui-avatar-spacing=lg]:-space-x-4",
		props.Class,
	)

	divArgs := []html.DivArg{
		html.AClass(className),
		html.AData("pui-avatar-spacing", string(props.Spacing)),
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

func GroupOverflow(count int, props Props, args ...html.DivArg) html.Node {
	className := classnames.Merge(
		"inline-flex items-center justify-center",
		"w-12 h-12 text-base",
		"data-[pui-avatar-size=sm]:w-8 data-[pui-avatar-size=sm]:h-8 data-[pui-avatar-size=sm]:text-xs",
		"data-[pui-avatar-size=md]:w-12 data-[pui-avatar-size=md]:h-12 data-[pui-avatar-size=md]:text-base",
		"data-[pui-avatar-size=lg]:w-16 data-[pui-avatar-size=lg]:h-16 data-[pui-avatar-size=lg]:text-xl",
		"rounded-full bg-muted ring-2 ring-background",
		props.Class,
	)

	dArgs := []html.DivArg{
		html.AClass(className),
	}
	if props.ID != "" {
		dArgs = append(dArgs, html.AId(props.ID))
	}
	if props.Size != "" {
		dArgs = append(dArgs, html.AData("pui-avatar-size", string(props.Size)))
	}
	for _, attr := range props.Attrs {
		dArgs = append(dArgs, attr)
	}
	dArgs = append(dArgs, args...)

	return html.Div(append(dArgs, html.Span(
		html.AClass("text-xs font-medium"),
		html.Text("+"+fmt.Sprint(count)),
	))...)
}
