package pagination

import (
	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/button"
	"github.com/plainkit/ui/internal/styles"
)

type Props struct {
	ID    string
	Class string
	Attrs []html.Global
}

type ContentProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type ItemProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type LinkProps struct {
	ID       string
	Class    string
	Attrs    []html.Global
	Href     string
	IsActive bool
	Disabled bool
}

type PreviousProps struct {
	ID       string
	Class    string
	Attrs    []html.Global
	Href     string
	Disabled bool
	Label    string
}

type NextProps struct {
	ID       string
	Class    string
	Attrs    []html.Global
	Href     string
	Disabled bool
	Label    string
}

func navArgsFromProps(baseClass string, extra ...string) func(p Props) []html.NavArg {
	return func(p Props) []html.NavArg {
		args := []html.NavArg{
			html.AClass(html.ClassMerge(append([]string{baseClass}, append(extra, p.Class)...)...)),
			html.AAria("label", "Pagination"),
		}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p Props) ApplyNav(attrs *html.NavAttrs, children *[]html.Component) {
	for _, a := range navArgsFromProps(styles.SurfaceMuted("flex flex-wrap items-center justify-center gap-3 rounded-2xl p-3 sm:p-4"))(p) {
		a.ApplyNav(attrs, children)
	}
}

func Pagination(args ...html.NavArg) html.Node {
	var (
		props Props
		rest  []html.NavArg
	)

	for _, a := range args {
		if v, ok := a.(Props); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Nav(append([]html.NavArg{props}, rest...)...)
}

func ulArgsFromProps(baseClass string, extra ...string) func(p ContentProps) []html.UlArg {
	return func(p ContentProps) []html.UlArg {
		args := []html.UlArg{html.AClass(html.ClassMerge(append([]string{baseClass}, append(extra, p.Class)...)...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p ContentProps) ApplyUl(attrs *html.UlAttrs, children *[]html.Component) {
	for _, a := range ulArgsFromProps(styles.Surface("inline-flex items-center gap-2 rounded-full border-none bg-transparent p-1"))(p) {
		a.ApplyUl(attrs, children)
	}
}

func Content(args ...html.UlArg) html.Node {
	var (
		props ContentProps
		rest  []html.UlArg
	)

	for _, a := range args {
		if v, ok := a.(ContentProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Ul(append([]html.UlArg{props}, rest...)...)
}

func liArgsFromProps(baseClass string, extra ...string) func(p ItemProps) []html.LiArg {
	return func(p ItemProps) []html.LiArg {
		var args []html.LiArg

		classNames := append([]string{baseClass}, extra...)
		classNames = append(classNames, p.Class)

		className := html.ClassMerge(classNames...)
		if className != "" {
			args = append(args, html.AClass(className))
		}

		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p ItemProps) ApplyLi(attrs *html.LiAttrs, children *[]html.Component) {
	for _, a := range liArgsFromProps("")(p) {
		a.ApplyLi(attrs, children)
	}
}

func Item(args ...html.LiArg) html.Node {
	var (
		props ItemProps
		rest  []html.LiArg
	)

	for _, a := range args {
		if v, ok := a.(ItemProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Li(append([]html.LiArg{props}, rest...)...)
}

func buttonArgsFromProps(baseClass string, extra ...string) func(p LinkProps) []html.ButtonArg {
	return func(p LinkProps) []html.ButtonArg {
		args := []html.ButtonArg{html.AClass(html.ClassMerge(append([]string{baseClass}, append(extra, p.Class)...)...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p LinkProps) ApplyButton(attrs *html.ButtonAttrs, children *[]html.Component) {
	for _, a := range buttonArgsFromProps("")(p) {
		a.ApplyButton(attrs, children)
	}
}

func Link(args ...html.ButtonArg) html.Node {
	var (
		props LinkProps
		rest  []html.ButtonArg
	)

	for _, a := range args {
		if v, ok := a.(LinkProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	btnProps := button.Props{
		ID:    props.ID,
		Attrs: props.Attrs,
		Size:  button.SizeIcon,
	}

	btnProps.Class = html.ClassMerge(
		styles.InteractiveGhost("pagination-link inline-flex h-10 w-10 items-center justify-center rounded-xl text-sm font-medium"),
		props.Class,
	)

	if props.IsActive {
		btnProps.Class = html.ClassMerge(btnProps.Class, "bg-primary/15 text-primary-foreground ring-1 ring-primary/40")
	}

	btnProps.Variant = button.VariantGhost
	if props.Disabled {
		btnProps.Disabled = true
		btnProps.Variant = button.VariantGhost
	} else {
		btnProps.Href = props.Href
	}

	buttonArgs := append([]html.ButtonArg{btnProps}, rest...)

	return button.Button(buttonArgs...)
}

func previousButtonArgsFromProps(baseClass string, extra ...string) func(p PreviousProps) []html.ButtonArg {
	return func(p PreviousProps) []html.ButtonArg {
		classNames := append([]string{baseClass}, extra...)
		classNames = append(classNames, "gap-1", p.Class)

		args := []html.ButtonArg{html.AClass(html.ClassMerge(classNames...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p PreviousProps) ApplyButton(attrs *html.ButtonAttrs, children *[]html.Component) {
	for _, a := range previousButtonArgsFromProps("")(p) {
		a.ApplyButton(attrs, children)
	}
}

func Previous(args ...html.ButtonArg) html.Node {
	var (
		props PreviousProps
		rest  []html.ButtonArg
	)

	for _, a := range args {
		if v, ok := a.(PreviousProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	btnProps := button.Props{
		ID:       props.ID,
		Attrs:    props.Attrs,
		Variant:  button.VariantGhost,
		Href:     props.Href,
		Disabled: props.Disabled,
	}

	btnProps.Class = html.ClassMerge(
		styles.InteractiveGhost("gap-2 rounded-full px-4 py-2 text-sm font-medium"),
		props.Class,
	)

	buttonArgs := []html.ButtonArg{lucide.ChevronLeft(html.AClass("size-4"))}

	buttonArgs = append(buttonArgs, rest...)
	if props.Label != "" {
		buttonArgs = append(buttonArgs, html.Span(html.Text(props.Label)))
	}

	allArgs := append([]html.ButtonArg{btnProps}, buttonArgs...)

	return button.Button(allArgs...)
}

func nextButtonArgsFromProps(baseClass string, extra ...string) func(p NextProps) []html.ButtonArg {
	return func(p NextProps) []html.ButtonArg {
		classNames := append([]string{baseClass}, extra...)
		classNames = append(classNames, "gap-1", p.Class)

		args := []html.ButtonArg{html.AClass(html.ClassMerge(classNames...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p NextProps) ApplyButton(attrs *html.ButtonAttrs, children *[]html.Component) {
	for _, a := range nextButtonArgsFromProps("")(p) {
		a.ApplyButton(attrs, children)
	}
}

func Next(args ...html.ButtonArg) html.Node {
	var (
		props NextProps
		rest  []html.ButtonArg
	)

	for _, a := range args {
		if v, ok := a.(NextProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	btnProps := button.Props{
		ID:       props.ID,
		Attrs:    props.Attrs,
		Variant:  button.VariantGhost,
		Href:     props.Href,
		Disabled: props.Disabled,
	}

	btnProps.Class = html.ClassMerge(
		styles.InteractiveGhost("gap-2 rounded-full px-4 py-2 text-sm font-medium"),
		props.Class,
	)

	buttonArgs := make([]html.ButtonArg, 0, len(rest)+2)
	if props.Label != "" {
		buttonArgs = append(buttonArgs, html.Span(html.Text(props.Label)))
	}

	buttonArgs = append(buttonArgs, rest...)
	buttonArgs = append(buttonArgs, lucide.ChevronRight(html.AClass("size-4")))

	allArgs := append([]html.ButtonArg{btnProps}, buttonArgs...)

	return button.Button(allArgs...)
}

func Ellipsis(args ...html.SvgArg) html.Node {
	svgArgs := []html.SvgArg{html.AClass("size-4 text-muted-foreground/70")}
	svgArgs = append(svgArgs, args...)

	return lucide.Ellipsis(svgArgs...)
}

func CreatePagination(currentPage, totalPages, maxVisible int) struct {
	CurrentPage int
	TotalPages  int
	Pages       []int
	HasPrevious bool
	HasNext     bool
} {
	if currentPage < 1 {
		currentPage = 1
	}

	if totalPages < 1 {
		totalPages = 1
	}

	if currentPage > totalPages {
		currentPage = totalPages
	}

	if maxVisible < 1 {
		maxVisible = 5
	}

	start, end := calculateVisibleRange(currentPage, totalPages, maxVisible)

	pages := make([]int, 0, end-start+1)
	for i := start; i <= end; i++ {
		pages = append(pages, i)
	}

	return struct {
		CurrentPage int
		TotalPages  int
		Pages       []int
		HasPrevious bool
		HasNext     bool
	}{
		CurrentPage: currentPage,
		TotalPages:  totalPages,
		Pages:       pages,
		HasPrevious: currentPage > 1,
		HasNext:     currentPage < totalPages,
	}
}

func calculateVisibleRange(currentPage, totalPages, maxVisible int) (int, int) {
	if totalPages <= maxVisible {
		return 1, totalPages
	}

	half := maxVisible / 2
	start := currentPage - half
	end := currentPage + half

	if start < 1 {
		end += 1 - start
		start = 1
	}

	if end > totalPages {
		start -= end - totalPages
		if start < 1 {
			start = 1
		}

		end = totalPages
	}

	return start, end
}
