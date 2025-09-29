package code

import (
	"crypto/rand"
	_ "embed"
	"encoding/hex"

	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/internal/classnames"
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

// Code renders a syntax-highlighted code block
func Code(props Props, args ...html.CodeArg) html.Node {
	id := props.ID
	if id == "" {
		id = "code-" + randomID("")
	}

	// Build code element classes
	codeClasses := classnames.Merge(
		"language-"+props.Language,
		"overflow-y-auto rounded-md block text-sm max-h-[501px]",
		func() string {
			switch props.Size {
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
		props.CodeClass,
	)

	// Create code element
	codeArgs := []html.CodeArg{
		html.AClass(codeClasses),
		html.AData("pui-code-block", ""),
	}
	codeArgs = append(codeArgs, args...)

	// Create pre element
	preElement := html.Pre(
		html.AClass("overflow-hidden"),
		html.Code(codeArgs...),
	)

	// Create container div args
	divArgs := []html.DivArg{
		html.AId(id),
		html.AClass(classnames.Merge("relative code-component", props.Class)),
		html.AData("pui-code-component", ""),
		preElement,
	}

	// Add copy button if requested
	if props.ShowCopyButton {
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
		divArgs = append(divArgs, copyButton)
	}

	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}

	// Include highlight.js CSS
	styleLink := html.Link(
		html.ARel("stylesheet"),
		html.AHref("https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/styles/pojoaque.min.css"),
	)

	return html.Div(divArgs...).WithAssets("", codeJS, "ui-code").WithAssets(html.Render(styleLink), "", "highlight-css")
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
