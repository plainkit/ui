package code

import (
	"crypto/rand"
	_ "embed"
	"encoding/hex"

	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
)

type Size string

const (
	SizeSm   Size = "sm"
	SizeLg   Size = "lg"
	SizeFull Size = "full"
)

type Props struct {
	ID             string
	Class          string
	Attrs          []html.Global
	Language       string
	ShowCopyButton bool
	Size           Size
	CodeClass      string
}

func divArgsFromProps(baseClass string, extra ...string) func(p Props) []html.DivArg {
	return func(p Props) []html.DivArg {
		args := []html.DivArg{html.AClass(html.ClassMerge(append([]string{baseClass}, append(extra, p.Class)...)...))}
		if p.ID != "" {
			args = append(args, html.AId(p.ID))
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p Props) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	id := p.ID
	if id == "" {
		id = "code-" + randomID("")
	}

	// Build code element classes
	codeClasses := html.ClassMerge(
		"language-"+p.Language,
		"overflow-y-auto rounded-md block text-sm max-h-[501px]",
		func() string {
			switch p.Size {
			case SizeSm:
				return "max-h-[250px]"
			case SizeLg:
				return "max-h-[1000px]"
			case SizeFull:
				return "max-h-full"
			default:
				return ""
			}
		}(),
		"hljs-target",
		p.CodeClass,
	)

	// Build container content
	containerContent := make([]html.Component, 0)

	// Create code element with children already passed in
	codeElement := html.Code(
		html.AClass(codeClasses),
		html.AData("pui-code-block", ""),
	)

	// Add code in pre element
	preElement := html.Pre(
		html.AClass("overflow-hidden"),
		codeElement,
	)

	containerContent = append(containerContent, preElement)

	// Add copy button if requested
	if p.ShowCopyButton {
		copyButton := html.Button(
			html.AType("button"),
			html.AClass("absolute top-2 right-2 hover:bg-gray-500 hover:bg-opacity-30 text-white p-2 rounded"),
			html.AData("pui-code-copy-button", ""),
			html.Span(
				html.AClass("hidden"),
				html.AData("pui-code-icon-check", ""),
				lucide.Check(html.AClass("size-3.5")),
			),
			html.Span(
				html.AData("pui-code-icon-clipboard", ""),
				lucide.Clipboard(html.AClass("size-3.5")),
			),
		)
		containerContent = append(containerContent, copyButton)
	}

	*children = append(*children, containerContent...)

	args := divArgsFromProps("relative code-component")(p)
	args = append([]html.DivArg{
		html.AId(id),
		html.AData("pui-code-component", ""),
	}, args...)

	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

// Code renders a syntax-highlighted code block
func Code(args ...html.DivArg) html.Node {
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

	// Include highlight.js CSS link directly in assets
	highlightCSS := `<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/styles/pojoaque.min.css">`

	return html.Div(append([]html.DivArg{props}, rest...)...).WithAssets(highlightCSS, codeJS, "ui-code")
}

func randomID(prefix string) string {
	buf := make([]byte, 4)
	if _, err := rand.Read(buf); err != nil {
		return prefix + "id"
	}

	return prefix + hex.EncodeToString(buf)
}

//go:embed code.js
var codeJS string
