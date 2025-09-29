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

func Table(props Props, args ...html.TableArg) html.Node {
	tableArgs := []html.TableArg{html.AClass(classnames.Merge("w-full caption-bottom text-sm", props.Class))}
	if props.ID != "" {
		tableArgs = append(tableArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		tableArgs = append(tableArgs, attr)
	}

	tableArgs = append(tableArgs, args...)

	return html.Div(
		html.AClass("relative w-full overflow-auto"),
		html.Table(tableArgs...),
	)
}

func Header(props HeaderProps, args ...html.TheadArg) html.Node {
	theadArgs := []html.TheadArg{html.AClass(classnames.Merge("[&_tr]:border-b", props.Class))}
	if props.ID != "" {
		theadArgs = append(theadArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		theadArgs = append(theadArgs, attr)
	}

	theadArgs = append(theadArgs, args...)

	return html.Thead(theadArgs...)
}

func Body(props BodyProps, args ...html.TbodyArg) html.Node {
	tbodyArgs := []html.TbodyArg{html.AClass(classnames.Merge("[&_tr:last-child]:border-0", props.Class))}
	if props.ID != "" {
		tbodyArgs = append(tbodyArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		tbodyArgs = append(tbodyArgs, attr)
	}

	tbodyArgs = append(tbodyArgs, args...)

	return html.Tbody(tbodyArgs...)
}

func Footer(props FooterProps, args ...html.TfootArg) html.Node {
	tfootArgs := []html.TfootArg{html.AClass(classnames.Merge("border-t bg-muted/50 font-medium [&>tr]:last:border-b-0", props.Class))}
	if props.ID != "" {
		tfootArgs = append(tfootArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		tfootArgs = append(tfootArgs, attr)
	}

	tfootArgs = append(tfootArgs, args...)

	return html.Tfoot(tfootArgs...)
}

func Row(props RowProps, args ...html.TrArg) html.Node {
	baseClass := "border-b transition-colors hover:bg-muted/50"
	if props.Selected {
		baseClass = classnames.Merge(baseClass, "data-[pui-table-state-selected]:bg-muted")
	}

	className := classnames.Merge(baseClass, props.Class)

	trArgs := []html.TrArg{html.AClass(className)}
	if props.ID != "" {
		trArgs = append(trArgs, html.AId(props.ID))
	}

	if props.Selected {
		trArgs = append(trArgs, html.AData("pui-table-state-selected", ""))
	}

	for _, attr := range props.Attrs {
		trArgs = append(trArgs, attr)
	}

	trArgs = append(trArgs, args...)

	return html.Tr(trArgs...)
}

func Head(props HeadProps, args ...html.ThArg) html.Node {
	thArgs := []html.ThArg{html.AClass(classnames.Merge(
		"h-10 px-2 text-left align-middle font-medium text-muted-foreground",
		"[&:has([role=checkbox])]:pr-0 [&>[role=checkbox]]:translate-y-[2px]",
		props.Class,
	))}
	if props.ID != "" {
		thArgs = append(thArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		thArgs = append(thArgs, attr)
	}

	thArgs = append(thArgs, args...)

	return html.Th(thArgs...)
}

func Cell(props CellProps, args ...html.TdArg) html.Node {
	tdArgs := []html.TdArg{html.AClass(classnames.Merge(
		"p-2 align-middle",
		"[&:has([role=checkbox])]:pr-0 [&>[role=checkbox]]:translate-y-[2px]",
		props.Class,
	))}
	if props.ID != "" {
		tdArgs = append(tdArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		tdArgs = append(tdArgs, attr)
	}

	tdArgs = append(tdArgs, args...)

	return html.Td(tdArgs...)
}

func Caption(props CaptionProps, args ...html.CaptionArg) html.Node {
	captionArgs := []html.CaptionArg{html.AClass(classnames.Merge("mt-4 text-sm text-muted-foreground", props.Class))}
	if props.ID != "" {
		captionArgs = append(captionArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		captionArgs = append(captionArgs, attr)
	}

	captionArgs = append(captionArgs, args...)

	return html.Caption(captionArgs...)
}
