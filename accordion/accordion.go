package accordion

import (
	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/internal/classnames"
)

type Props struct {
	ID    string
	Class string
	Attrs []html.Global
}

type ItemProps Props

type TriggerProps Props

type ContentProps Props

func Accordion(props Props, args ...html.DivArg) html.Node {
	divArgs := []html.DivArg{html.AClass(props.Class)}
	if props.ID != "" {
		divArgs = append(divArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}

	divArgs = append(divArgs, args...)

	return html.Div(divArgs...)
}

func Item(props ItemProps, args ...html.DetailsArg) html.Node {
	detailsArgs := []html.DetailsArg{
		html.AClass(classnames.Merge("group border-b last:border-b-0 [&[open]>summary>svg]:rotate-180", props.Class)),
		html.AName("accordion"),
	}
	if props.ID != "" {
		detailsArgs = append(detailsArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		detailsArgs = append(detailsArgs, attr)
	}

	detailsArgs = append(detailsArgs, args...)

	return html.Details(detailsArgs...)
}

func Trigger(props TriggerProps, args ...html.SummaryArg) html.Node {
	summaryArgs := []html.SummaryArg{html.AClass(classnames.Merge(
		"flex flex-1 items-start justify-between gap-4 py-4",
		"text-left text-sm font-medium",
		"transition-all hover:underline cursor-pointer",
		"outline-none focus-visible:ring-[3px] focus-visible:ring-ring/50 focus-visible:border-ring rounded-md",
		"disabled:pointer-events-none disabled:opacity-50",
		"list-none [&::-webkit-details-marker]:hidden",
		props.Class,
	))}
	if props.ID != "" {
		summaryArgs = append(summaryArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		summaryArgs = append(summaryArgs, attr)
	}

	summaryArgs = append(summaryArgs, args...)
	summaryArgs = append(summaryArgs, lucide.ChevronDown(
		html.AClass("size-4 shrink-0 translate-y-0.5 transition-transform duration-200 text-muted-foreground pointer-events-none"),
	))

	return html.Summary(summaryArgs...)
}

func Content(props ContentProps, args ...html.DivArg) html.Node {
	divArgs := []html.DivArg{html.AClass(classnames.Merge("pt-0 pb-4 text-sm overflow-hidden", props.Class))}
	if props.ID != "" {
		divArgs = append(divArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}

	divArgs = append(divArgs, args...)

	return html.Div(divArgs...)
}
