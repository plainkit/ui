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

func (p Props) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	args := []html.DivArg{html.AClass(classnames.Merge(p.Class))}
	if p.ID != "" {
		args = append(args, html.AId(p.ID))
	}

	for _, a := range p.Attrs {
		args = append(args, a)
	}

	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

func Accordion(args ...html.DivArg) html.Node {
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

func (p ItemProps) ApplyDetails(attrs *html.DetailsAttrs, children *[]html.Component) {
	args := []html.DetailsArg{
		html.AClass(classnames.Merge("group border-b last:border-b-0 [&[open]>summary>svg]:rotate-180", p.Class)),
		html.AName("accordion"),
	}
	if p.ID != "" {
		args = append(args, html.AId(p.ID))
	}

	for _, a := range p.Attrs {
		args = append(args, a)
	}

	for _, a := range args {
		a.ApplyDetails(attrs, children)
	}
}

func Item(args ...html.DetailsArg) html.Node {
	var (
		props ItemProps
		rest  []html.DetailsArg
	)

	for _, a := range args {
		if v, ok := a.(ItemProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Details(append([]html.DetailsArg{props}, rest...)...)
}

func (p TriggerProps) ApplySummary(attrs *html.SummaryAttrs, children *[]html.Component) {
	args := []html.SummaryArg{html.AClass(classnames.Merge(
		"flex flex-1 items-start justify-between gap-4 py-4",
		"text-left text-sm font-medium",
		"transition-all hover:underline cursor-pointer",
		"outline-none focus-visible:ring-[3px] focus-visible:ring-ring/50 focus-visible:border-ring rounded-md",
		"disabled:pointer-events-none disabled:opacity-50",
		"list-none [&::-webkit-details-marker]:hidden",
		p.Class,
	))}
	if p.ID != "" {
		args = append(args, html.AId(p.ID))
	}

	for _, a := range p.Attrs {
		args = append(args, a)
	}

	for _, a := range args {
		a.ApplySummary(attrs, children)
	}
}

func Trigger(args ...html.SummaryArg) html.Node {
	var (
		props TriggerProps
		rest  []html.SummaryArg
	)

	for _, a := range args {
		if v, ok := a.(TriggerProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	// Add the chevron icon at the end
	rest = append(rest, lucide.ChevronDown(
		html.AClass("size-4 shrink-0 translate-y-0.5 transition-transform duration-200 text-muted-foreground pointer-events-none"),
	))

	return html.Summary(append([]html.SummaryArg{props}, rest...)...)
}

func (p ContentProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	args := []html.DivArg{html.AClass(classnames.Merge("pt-0 pb-4 text-sm overflow-hidden", p.Class))}
	if p.ID != "" {
		args = append(args, html.AId(p.ID))
	}

	for _, a := range p.Attrs {
		args = append(args, a)
	}

	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

func Content(args ...html.DivArg) html.Node {
	var (
		props ContentProps
		rest  []html.DivArg
	)

	for _, a := range args {
		if v, ok := a.(ContentProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Div(append([]html.DivArg{props}, rest...)...)
}
