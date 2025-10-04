package tabs

import (
	"crypto/rand"
	_ "embed"
	"encoding/hex"

	"github.com/plainkit/html"
	"github.com/plainkit/ui/internal/styles"
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
	for _, a := range tabsDivArgsFromProps("flex flex-col gap-4")(p) {
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
	for _, a := range listDivArgsFromProps(styles.SurfaceMuted("inline-flex h-11 w-fit items-center gap-1 rounded-full p-1.5 backdrop-blur-sm supports-[backdrop-filter]:bg-muted/70"))(p) {
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
		styles.InteractiveGhost(
			"flex-1 items-center justify-center gap-2 rounded-full px-4 py-2 text-sm font-medium",
			"transition-all",
			"data-[pui-tabs-state=active]:bg-background data-[pui-tabs-state=active]:text-foreground",
			"data-[pui-tabs-state=active]:shadow-sm",
		),
		"focus-visible:ring-2 focus-visible:ring-ring/40",
		"disabled:pointer-events-none disabled:opacity-60",
		"[&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4",
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
	for _, a := range contentDivArgsFromProps(styles.SurfaceMuted("flex-1 rounded-2xl border-none bg-background/40 p-6 outline-none"))(p) {
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
