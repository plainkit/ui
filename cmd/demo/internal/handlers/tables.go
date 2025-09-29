package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/table"
)

func RenderTablesContent() html.Node {
	type invoice struct {
		Number string
		Status string
		Method string
		Amount string
	}

	invoices := []invoice{
		{Number: "INV-001", Status: "Paid", Method: "Credit Card", Amount: "$2,500"},
		{Number: "INV-002", Status: "Processing", Method: "Bank Transfer", Amount: "$980"},
		{Number: "INV-003", Status: "Unpaid", Method: "Credit Card", Amount: "$1,430"},
		{Number: "INV-004", Status: "Paid", Method: "PayPal", Amount: "$3,200"},
	}

	bodyRows := make([]html.TbodyArg, 0, len(invoices))
	for idx, inv := range invoices {
		bodyRows = append(bodyRows,
			table.Row(
				table.RowProps{Selected: idx == 1},
				table.Cell(table.CellProps{}, html.Text(inv.Number)),
				table.Cell(table.CellProps{}, html.Text(inv.Status)),
				table.Cell(table.CellProps{}, html.Text(inv.Method)),
				table.Cell(table.CellProps{}, html.Text(inv.Amount)),
			),
		)
	}

	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Tables")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Present structured data with consistent spacing and hover affordances.")),
			),
			table.Table(
				table.Props{},
				table.Caption(table.CaptionProps{}, html.Text("Invoice history for Q3 revenue.")),
				table.Header(
					table.HeaderProps{},
					table.Row(
						table.RowProps{},
						table.Head(table.HeadProps{}, html.Text("Invoice")),
						table.Head(table.HeadProps{}, html.Text("Status")),
						table.Head(table.HeadProps{}, html.Text("Method")),
						table.Head(table.HeadProps{}, html.Text("Amount")),
					),
				),
				table.Body(table.BodyProps{}, bodyRows...),
				table.Footer(
					table.FooterProps{},
					table.Row(
						table.RowProps{},
						table.Cell(table.CellProps{Class: "font-medium"}, html.Text("Total")),
						table.Cell(table.CellProps{}, html.Text("2 paid")),
						table.Cell(table.CellProps{}, html.Text("2 pending")),
						table.Cell(table.CellProps{Class: "font-semibold"}, html.Text("$8,110")),
					),
				),
			),
		),
	)
}
