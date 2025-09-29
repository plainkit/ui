package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/code"
)

func RenderCodeContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Code")),
				html.P(html.AClass("text-sm text-muted-foreground"), html.Text("Syntax-highlighted code blocks with optional copy functionality.")),
			),
			html.Div(
				html.AClass("grid gap-6 md:grid-cols-2"),
				html.Div(
					html.AClass("space-y-4"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("JavaScript Example")),
					code.Code(
						code.Props{
							Language:       "javascript",
							ShowCopyButton: true,
							Size:           code.SizeSm,
						},
						html.Text(`function fibonacci(n) {
  if (n <= 1) return n;
  return fibonacci(n - 1) + fibonacci(n - 2);
}

console.log(fibonacci(10)); // 55`),
					),
				),
				html.Div(
					html.AClass("space-y-4"),
					html.H3(html.AClass("text-lg font-semibold"), html.Text("Go Example")),
					code.Code(
						code.Props{
							Language:       "go",
							ShowCopyButton: true,
							Size:           code.SizeLg,
						},
						html.Text(`package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        fmt.Printf("Hello, World! %d\n", i)
    }
}`),
					),
				),
			),
		),
	)
}
