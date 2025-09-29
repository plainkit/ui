package pagination

import (
	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/button"
	"github.com/plainkit/ui/internal/classnames"
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

func Pagination(props Props, args ...html.NavArg) html.Node {
	navArgs := []html.NavArg{
		html.AClass(classnames.Merge("flex flex-wrap justify-center", props.Class)),
		html.AAria("label", "Pagination"),
	}
	if props.ID != "" {
		navArgs = append(navArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		navArgs = append(navArgs, attr)
	}
	navArgs = append(navArgs, args...)
	return html.Nav(navArgs...)
}

func Content(props ContentProps, args ...html.UlArg) html.Node {
	ulArgs := []html.UlArg{html.AClass(classnames.Merge("flex flex-row items-center gap-1", props.Class))}
	if props.ID != "" {
		ulArgs = append(ulArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		ulArgs = append(ulArgs, attr)
	}
	ulArgs = append(ulArgs, args...)
	return html.Ul(ulArgs...)
}

func Item(props ItemProps, args ...html.LiArg) html.Node {
	liArgs := []html.LiArg{}
	if props.Class != "" {
		liArgs = append(liArgs, html.AClass(props.Class))
	}
	if props.ID != "" {
		liArgs = append(liArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		liArgs = append(liArgs, attr)
	}
	liArgs = append(liArgs, args...)
	return html.Li(liArgs...)
}

func Link(props LinkProps, args ...html.ButtonArg) html.Node {
	btnProps := button.Props{
		ID:      props.ID,
		Class:   props.Class,
		Attrs:   props.Attrs,
		Size:    button.SizeIcon,
		Variant: buttonVariant(props.IsActive),
	}
	if props.Disabled {
		btnProps.Disabled = true
		btnProps.Variant = button.VariantGhost
	} else {
		btnProps.Href = props.Href
	}
	return button.Button(btnProps, args...)
}

func Previous(props PreviousProps, args ...html.ButtonArg) html.Node {
	btnProps := button.Props{
		ID:       props.ID,
		Attrs:    props.Attrs,
		Class:    classnames.Merge("gap-1", props.Class),
		Variant:  button.VariantGhost,
		Href:     props.Href,
		Disabled: props.Disabled,
	}

	buttonArgs := []html.ButtonArg{lucide.ChevronLeft(html.AClass("size-4"))}
	buttonArgs = append(buttonArgs, args...)
	if props.Label != "" {
		buttonArgs = append(buttonArgs, html.Span(html.Text(props.Label)))
	}

	return button.Button(btnProps, buttonArgs...)
}

func Next(props NextProps, args ...html.ButtonArg) html.Node {
	btnProps := button.Props{
		ID:       props.ID,
		Attrs:    props.Attrs,
		Class:    classnames.Merge("gap-1", props.Class),
		Variant:  button.VariantGhost,
		Href:     props.Href,
		Disabled: props.Disabled,
	}

	buttonArgs := make([]html.ButtonArg, 0, len(args)+2)
	if props.Label != "" {
		buttonArgs = append(buttonArgs, html.Span(html.Text(props.Label)))
	}
	buttonArgs = append(buttonArgs, args...)
	buttonArgs = append(buttonArgs, lucide.ChevronRight(html.AClass("size-4")))

	return button.Button(btnProps, buttonArgs...)
}

func Ellipsis(args ...html.SvgArg) html.Node {
	svgArgs := []html.SvgArg{html.AClass("size-4 text-muted-foreground")}
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

func buttonVariant(isActive bool) button.Variant {
	if isActive {
		return button.VariantOutline
	}
	return button.VariantGhost
}
