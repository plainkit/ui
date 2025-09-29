package table

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/internal/classnames"
)

type Props struct {
	ID    string
	Class string
	Attrs []html.Global
}

type HeaderProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type BodyProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type FooterProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type RowProps struct {
	ID       string
	Class    string
	Attrs    []html.Global
	Selected bool
}

type HeadProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type CellProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type CaptionProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

func tableArgsFromProps(baseClass string, extra ...string) func(p Props) []html.TableArg {
	return func(p Props) []html.TableArg {
		args := []html.TableArg{html.AClass(classnames.Merge(append([]string{baseClass}, append(extra, p.Class)...)...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p Props) ApplyTable(attrs *html.TableAttrs, children *[]html.Component) {
	for _, a := range tableArgsFromProps("w-full caption-bottom text-sm")(p) {
		a.ApplyTable(attrs, children)
	}
}

func Table(args ...html.TableArg) html.Node {
	var (
		props Props
		rest  []html.TableArg
	)

	for _, a := range args {
		if v, ok := a.(Props); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Div(
		html.AClass("relative w-full overflow-auto"),
		html.Table(append([]html.TableArg{props}, rest...)...),
	)
}

func theadArgsFromProps(baseClass string, extra ...string) func(p HeaderProps) []html.TheadArg {
	return func(p HeaderProps) []html.TheadArg {
		args := []html.TheadArg{html.AClass(classnames.Merge(append([]string{baseClass}, append(extra, p.Class)...)...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p HeaderProps) ApplyThead(attrs *html.TheadAttrs, children *[]html.Component) {
	for _, a := range theadArgsFromProps("[&_tr]:border-b")(p) {
		a.ApplyThead(attrs, children)
	}
}

func Header(args ...html.TheadArg) html.Node {
	var (
		props HeaderProps
		rest  []html.TheadArg
	)

	for _, a := range args {
		if v, ok := a.(HeaderProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Thead(append([]html.TheadArg{props}, rest...)...)
}

func tbodyArgsFromProps(baseClass string, extra ...string) func(p BodyProps) []html.TbodyArg {
	return func(p BodyProps) []html.TbodyArg {
		args := []html.TbodyArg{html.AClass(classnames.Merge(append([]string{baseClass}, append(extra, p.Class)...)...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p BodyProps) ApplyTbody(attrs *html.TbodyAttrs, children *[]html.Component) {
	for _, a := range tbodyArgsFromProps("[&_tr:last-child]:border-0")(p) {
		a.ApplyTbody(attrs, children)
	}
}

func Body(args ...html.TbodyArg) html.Node {
	var (
		props BodyProps
		rest  []html.TbodyArg
	)

	for _, a := range args {
		if v, ok := a.(BodyProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Tbody(append([]html.TbodyArg{props}, rest...)...)
}

func tfootArgsFromProps(baseClass string, extra ...string) func(p FooterProps) []html.TfootArg {
	return func(p FooterProps) []html.TfootArg {
		args := []html.TfootArg{html.AClass(classnames.Merge(append([]string{baseClass}, append(extra, p.Class)...)...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p FooterProps) ApplyTfoot(attrs *html.TfootAttrs, children *[]html.Component) {
	for _, a := range tfootArgsFromProps("border-t bg-muted/50 font-medium [&>tr]:last:border-b-0")(p) {
		a.ApplyTfoot(attrs, children)
	}
}

func Footer(args ...html.TfootArg) html.Node {
	var (
		props FooterProps
		rest  []html.TfootArg
	)

	for _, a := range args {
		if v, ok := a.(FooterProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Tfoot(append([]html.TfootArg{props}, rest...)...)
}

func trArgsFromProps(baseClass string, extra ...string) func(p RowProps) []html.TrArg {
	return func(p RowProps) []html.TrArg {
		classNames := append([]string{baseClass}, extra...)
		if p.Selected {
			classNames = append(classNames, "data-[pui-table-state-selected]:bg-muted")
		}

		classNames = append(classNames, p.Class)

		args := []html.TrArg{html.AClass(classnames.Merge(classNames...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		if p.Selected {
			args = append(args, html.AData("pui-table-state-selected", ""))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p RowProps) ApplyTr(attrs *html.TrAttrs, children *[]html.Component) {
	for _, a := range trArgsFromProps("border-b transition-colors hover:bg-muted/50")(p) {
		a.ApplyTr(attrs, children)
	}
}

func Row(args ...html.TrArg) html.Node {
	var (
		props RowProps
		rest  []html.TrArg
	)

	for _, a := range args {
		if v, ok := a.(RowProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Tr(append([]html.TrArg{props}, rest...)...)
}

func thArgsFromProps(baseClass string, extra ...string) func(p HeadProps) []html.ThArg {
	return func(p HeadProps) []html.ThArg {
		args := []html.ThArg{html.AClass(classnames.Merge(append([]string{baseClass}, append(extra, p.Class)...)...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p HeadProps) ApplyTh(attrs *html.ThAttrs, children *[]html.Component) {
	for _, a := range thArgsFromProps(
		"h-10 px-2 text-left align-middle font-medium text-muted-foreground",
		"[&:has([role=checkbox])]:pr-0 [&>[role=checkbox]]:translate-y-[2px]",
	)(p) {
		a.ApplyTh(attrs, children)
	}
}

func Head(args ...html.ThArg) html.Node {
	var (
		props HeadProps
		rest  []html.ThArg
	)

	for _, a := range args {
		if v, ok := a.(HeadProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Th(append([]html.ThArg{props}, rest...)...)
}

func tdArgsFromProps(baseClass string, extra ...string) func(p CellProps) []html.TdArg {
	return func(p CellProps) []html.TdArg {
		args := []html.TdArg{html.AClass(classnames.Merge(append([]string{baseClass}, append(extra, p.Class)...)...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p CellProps) ApplyTd(attrs *html.TdAttrs, children *[]html.Component) {
	for _, a := range tdArgsFromProps(
		"p-2 align-middle",
		"[&:has([role=checkbox])]:pr-0 [&>[role=checkbox]]:translate-y-[2px]",
	)(p) {
		a.ApplyTd(attrs, children)
	}
}

func Cell(args ...html.TdArg) html.Node {
	var (
		props CellProps
		rest  []html.TdArg
	)

	for _, a := range args {
		if v, ok := a.(CellProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Td(append([]html.TdArg{props}, rest...)...)
}

func captionArgsFromProps(baseClass string, extra ...string) func(p CaptionProps) []html.CaptionArg {
	return func(p CaptionProps) []html.CaptionArg {
		args := []html.CaptionArg{html.AClass(classnames.Merge(append([]string{baseClass}, append(extra, p.Class)...)...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p CaptionProps) ApplyCaption(attrs *html.CaptionAttrs, children *[]html.Component) {
	for _, a := range captionArgsFromProps("mt-4 text-sm text-muted-foreground")(p) {
		a.ApplyCaption(attrs, children)
	}
}

func Caption(args ...html.CaptionArg) html.Node {
	var (
		props CaptionProps
		rest  []html.CaptionArg
	)

	for _, a := range args {
		if v, ok := a.(CaptionProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Caption(append([]html.CaptionArg{props}, rest...)...)
}
