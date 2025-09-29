package tabs

import (
	"crypto/rand"
	_ "embed"
	"encoding/hex"

	"github.com/plainkit/html"
)

type Props struct {
	ID    string
	Class string
	Attrs []html.Global
}

type ListProps struct {
	ID     string
	Class  string
	Attrs  []html.Global
	TabsID string
}

type TriggerProps struct {
	ID       string
	Class    string
	Attrs    []html.Global
	Value    string
	IsActive bool
	TabsID   string
}

type ContentProps struct {
	ID       string
	Class    string
	Attrs    []html.Global
	Value    string
	IsActive bool
	TabsID   string
}

func tabsDivArgsFromProps(baseClass string, extra ...string) func(p Props) []html.DivArg {
	return func(p Props) []html.DivArg {
		if p.ID == "" {
			p.ID = randomID("tabs")
		}

		args := []html.DivArg{
			html.AClass(html.ClassMerge(append([]string{baseClass}, append(extra, p.Class)...)...)),
			html.AData("pui-tabs", ""),
			html.AData("pui-tabs-id", p.ID),
			html.AId(p.ID),
		}

		for _, a := range p.Attrs {
			args = append(args, a)
		}

		return args
	}
}

func (p Props) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	for _, a := range tabsDivArgsFromProps("flex flex-col gap-2")(p) {
		a.ApplyDiv(attrs, children)
	}
}

func Tabs(args ...html.DivArg) html.Node {
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

	node := html.Div(append([]html.DivArg{props}, rest...)...)

	return node.WithAssets("", tabsJS, "ui-tabs")
}

func listDivArgsFromProps(baseClass string, extra ...string) func(p ListProps) []html.DivArg {
	return func(p ListProps) []html.DivArg {
		tabsID := p.TabsID
		if tabsID == "" {
			tabsID = p.ID
		}

		args := []html.DivArg{
			html.AClass(html.ClassMerge(append([]string{baseClass}, append(extra, p.Class)...)...)),
			html.AData("pui-tabs-list", ""),
			html.AData("pui-tabs-id", tabsID),
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

func (p ListProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	for _, a := range listDivArgsFromProps("bg-muted text-muted-foreground inline-flex h-9 w-fit items-center justify-center rounded-lg p-[3px]")(p) {
		a.ApplyDiv(attrs, children)
	}
}

func List(args ...html.DivArg) html.Node {
	var (
		props ListProps
		rest  []html.DivArg
	)

	for _, a := range args {
		if v, ok := a.(ListProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	return html.Div(append([]html.DivArg{props}, rest...)...)
}

func triggerButtonArgsFromProps(baseClass string, extra ...string) func(p TriggerProps) []html.ButtonArg {
	return func(p TriggerProps) []html.ButtonArg {
		if p.TabsID == "" {
			p.TabsID = randomID("tabs")
		}

		args := []html.ButtonArg{
			html.AType("button"),
			html.AClass(html.ClassMerge(append([]string{baseClass}, append(extra, p.Class)...)...)),
			html.AData("pui-tabs-trigger", ""),
			html.AData("pui-tabs-id", p.TabsID),
			html.AData("pui-tabs-value", p.Value),
			html.AData("pui-tabs-state", stateAttr(p.IsActive)),
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

func (p TriggerProps) ApplyButton(attrs *html.ButtonAttrs, children *[]html.Component) {
	for _, a := range triggerButtonArgsFromProps(
		"data-[pui-tabs-state=active]:bg-background dark:data-[pui-tabs-state=active]:text-foreground",
		"focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:outline-ring dark:data-[pui-tabs-state=active]:border-input",
		"dark:data-[pui-tabs-state=active]:bg-input/30 text-foreground dark:text-muted-foreground",
		"inline-flex h-[calc(100%-1px)] flex-1 items-center justify-center gap-1.5",
		"rounded-md border border-transparent px-2 py-1 text-sm font-medium whitespace-nowrap transition-[color,box-shadow]",
		"focus-visible:ring-[3px] focus-visible:outline-1 disabled:pointer-events-none disabled:opacity-50",
		"data-[pui-tabs-state=active]:shadow-sm [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4",
	)(p) {
		a.ApplyButton(attrs, children)
	}
}

func Trigger(args ...html.ButtonArg) html.Node {
	var (
		props TriggerProps
		rest  []html.ButtonArg
	)

	for _, a := range args {
		if v, ok := a.(TriggerProps); ok {
			props = v
		} else {
			rest = append(rest, a)
		}
	}

	if props.Value == "" {
		return html.Span(html.AClass("text-xs text-destructive"), html.Text("tabs.Trigger requires Value"))
	}

	return html.Button(append([]html.ButtonArg{props}, rest...)...)
}

func contentDivArgsFromProps(baseClass string, extra ...string) func(p ContentProps) []html.DivArg {
	return func(p ContentProps) []html.DivArg {
		classNames := append([]string{baseClass}, extra...)
		classNames = append(classNames, hiddenClass(!p.IsActive), p.Class)

		args := []html.DivArg{
			html.AClass(html.ClassMerge(classNames...)),
			html.AData("pui-tabs-content", ""),
			html.AData("pui-tabs-value", p.Value),
			html.AData("pui-tabs-state", stateAttr(p.IsActive)),
		}
		if p.TabsID != "" {
			args = append(args, html.AData("pui-tabs-id", p.TabsID))
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

func (p ContentProps) ApplyDiv(attrs *html.DivAttrs, children *[]html.Component) {
	for _, a := range contentDivArgsFromProps("flex-1 outline-none")(p) {
		a.ApplyDiv(attrs, children)
	}
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

	if props.Value == "" {
		return html.Span(html.AClass("text-xs text-destructive"), html.Text("tabs.Content requires Value"))
	}

	return html.Div(append([]html.DivArg{props}, rest...)...)
}

func stateAttr(active bool) string {
	if active {
		return "active"
	}

	return "inactive"
}

func hiddenClass(hidden bool) string {
	if hidden {
		return "hidden"
	}

	return ""
}

func randomID(prefix string) string {
	buf := make([]byte, 4)
	if _, err := rand.Read(buf); err != nil {
		return prefix + "-id"
	}

	return prefix + "-" + hex.EncodeToString(buf)
}

//go:embed tabs.js
var tabsJS string
