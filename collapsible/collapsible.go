package collapsible

import (
	"crypto/rand"
	_ "embed"
	"encoding/hex"

	"github.com/plainkit/html"
	"github.com/plainkit/ui/internal/classnames"
)

type Props struct {
	ID    string
	Class string
	Attrs []html.Global
	Open  bool
}

type TriggerProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type ContentProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

func divArgsFromProps(baseClass string, extra ...string) func(p Props) []html.DivArg {
	return func(p Props) []html.DivArg {
		args := []html.DivArg{html.AClass(classnames.Merge(append([]string{baseClass}, append(extra, p.Class)...)...))}
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
		id = randomID("collapsible")
	}

	state := "closed"
	if p.Open {
		state = "open"
	}

	args := divArgsFromProps("")(p)
	args = append([]html.DivArg{
		html.AId(id),
		html.AData("pui-collapsible", "root"),
		html.AData("pui-collapsible-state", state),
	}, args...)

	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

func (p TriggerProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	args := []html.DivArg{
		html.AClass(classnames.Merge("", p.Class)),
		html.AData("pui-collapsible", "trigger"),
	}
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

func (p ContentProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	args := []html.DivArg{
		html.AClass(classnames.Merge(
			"grid grid-rows-[0fr] transition-[grid-template-rows] duration-200 ease-out [[data-pui-collapsible-state=open]_&]:grid-rows-[1fr]",
			p.Class,
		)),
		html.AData("pui-collapsible", "content"),
	}
	if p.ID != "" {
		args = append(args, html.AId(p.ID))
	}

	for _, a := range p.Attrs {
		args = append(args, a)
	}

	// Create inner div with overflow hidden
	innerDiv := html.Div(html.AClass("overflow-hidden"))
	*children = append(*children, innerDiv)

	for _, a := range args {
		a.ApplyDiv(attrs, children)
	}
}

func Collapsible(args ...html.DivArg) html.Node {
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

	return html.Div(append([]html.DivArg{props}, rest...)...).WithAssets("", collapsibleJS, "ui-collapsible")
}

func Trigger(args ...html.DivArg) html.Node {
	var (
		props TriggerProps
		rest  []html.DivArg
	)

	for _, a := range args {
		if v, ok := a.(TriggerProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Div(append([]html.DivArg{props}, rest...)...)
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

	// Create inner div with overflow hidden and rest of content
	innerDiv := html.Div(append([]html.DivArg{html.AClass("overflow-hidden")}, rest...)...)

	return html.Div(props, innerDiv)
}

func randomID(prefix string) string {
	buf := make([]byte, 4)
	if _, err := rand.Read(buf); err != nil {
		return prefix + "-id"
	}

	return prefix + "-" + hex.EncodeToString(buf)
}

//go:embed collapsible.js
var collapsibleJS string
