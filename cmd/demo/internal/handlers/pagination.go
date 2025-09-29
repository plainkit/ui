package handlers

import (
	"strconv"

	"github.com/plainkit/html"
	"github.com/plainkit/ui/pagination"
)

func previousHref(data struct {
	CurrentPage int
	TotalPages  int
	Pages       []int
	HasPrevious bool
	HasNext     bool
}) string {
	if !data.HasPrevious {
		return ""
	}

	return "?page=" + strconv.Itoa(data.CurrentPage-1)
}

func nextHref(data struct {
	CurrentPage int
	TotalPages  int
	Pages       []int
	HasPrevious bool
	HasNext     bool
}) string {
	if !data.HasNext {
		return ""
	}

	return "?page=" + strconv.Itoa(data.CurrentPage+1)
}

func RenderPaginationContent() html.Node {
	data := pagination.CreatePagination(6, 20, 5)

	items := []html.UlArg{
		pagination.Item(
			pagination.ItemProps{},
			pagination.Previous(pagination.PreviousProps{
				Href:     previousHref(data),
				Disabled: !data.HasPrevious,
				Label:    "Previous",
			}),
		),
	}

	if first := data.Pages[0]; first > 1 {
		items = append(items,
			pagination.Item(
				pagination.ItemProps{},
				pagination.Link(pagination.LinkProps{Href: "?page=1", IsActive: data.CurrentPage == 1}, html.T("1")),
			),
		)
		if first > 2 {
			items = append(items, pagination.Item(pagination.ItemProps{}, pagination.Ellipsis()))
		}
	}

	for _, pageNumber := range data.Pages {
		label := strconv.Itoa(pageNumber)
		items = append(items,
			pagination.Item(
				pagination.ItemProps{},
				pagination.Link(
					pagination.LinkProps{
						Href:     "?page=" + label,
						IsActive: pageNumber == data.CurrentPage,
					},
					html.T(label),
				),
			),
		)
	}

	lastVisible := data.Pages[len(data.Pages)-1]
	if lastVisible < data.TotalPages {
		if lastVisible < data.TotalPages-1 {
			items = append(items, pagination.Item(pagination.ItemProps{}, pagination.Ellipsis()))
		}

		lastLabel := strconv.Itoa(data.TotalPages)
		items = append(items,
			pagination.Item(
				pagination.ItemProps{},
				pagination.Link(
					pagination.LinkProps{
						Href:     "?page=" + lastLabel,
						IsActive: data.CurrentPage == data.TotalPages,
					},
					html.T(lastLabel),
				),
			),
		)
	}

	items = append(items,
		pagination.Item(
			pagination.ItemProps{},
			pagination.Next(pagination.NextProps{
				Href:     nextHref(data),
				Disabled: !data.HasNext,
				Label:    "Next",
			}),
		),
	)

	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Pagination")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Navigate long collections page by page.")),
			),
			pagination.Pagination(
				pagination.Props{},
				pagination.Content(append([]html.UlArg{pagination.ContentProps{}}, items...)...),
			),
		),
	)
}
