package accordion

import (
	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/internal/styles"
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
	args := []html.DivArg{html.AClass(html.ClassMerge(styles.Surface("divide-y divide-border/40 overflow-hidden"), p.Class))}
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
		html.AClass(html.ClassMerge("group border-b border-border/40 last:border-b-0 [&[open]>summary>svg]:rotate-180", p.Class)),
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
	args := []html.SummaryArg{html.AClass(html.ClassMerge(
		styles.InteractiveGhost(
			"w-full justify-between gap-4 text-left",
			"rounded-none px-0 py-5 text-base font-medium",
		),
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
	args := []html.DivArg{html.AClass(html.ClassMerge(styles.SubtleText("overflow-hidden pb-5 pt-0"), p.Class))}
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
