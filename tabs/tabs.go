package tabs

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

func Tabs(props Props, args ...html.DivArg) html.Node {
	if props.ID == "" {
		props.ID = randomID("tabs")
	}

	divArgs := []html.DivArg{
		html.AClass(classnames.Merge("flex flex-col gap-2", props.Class)),
		html.AData("pui-tabs", ""),
		html.AData("pui-tabs-id", props.ID),
	}

	divArgs = append(divArgs, html.AId(props.ID))
	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}

	divArgs = append(divArgs, args...)

	node := html.Div(divArgs...)

	return node.WithAssets("", tabsJS, "ui-tabs")
}

func List(props ListProps, args ...html.DivArg) html.Node {
	tabsID := props.TabsID
	if tabsID == "" {
		tabsID = props.ID
	}

	divArgs := []html.DivArg{
		html.AClass(classnames.Merge("bg-muted text-muted-foreground inline-flex h-9 w-fit items-center justify-center rounded-lg p-[3px]", props.Class)),
		html.AData("pui-tabs-list", ""),
		html.AData("pui-tabs-id", tabsID),
	}
	if props.ID != "" {
		divArgs = append(divArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}

	divArgs = append(divArgs, args...)

	return html.Div(divArgs...)
}

func Trigger(props TriggerProps, args ...html.ButtonArg) html.Node {
	if props.Value == "" {
		return html.Span(html.AClass("text-xs text-destructive"), html.Text("tabs.Trigger requires Value"))
	}

	if props.TabsID == "" {
		props.TabsID = randomID("tabs")
	}

	buttonArgs := []html.ButtonArg{
		html.AType("button"),
		html.AClass(classnames.Merge(
			"data-[pui-tabs-state=active]:bg-background dark:data-[pui-tabs-state=active]:text-foreground",
			"focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:outline-ring dark:data-[pui-tabs-state=active]:border-input",
			"dark:data-[pui-tabs-state=active]:bg-input/30 text-foreground dark:text-muted-foreground",
			"inline-flex h-[calc(100%-1px)] flex-1 items-center justify-center gap-1.5",
			"rounded-md border border-transparent px-2 py-1 text-sm font-medium whitespace-nowrap transition-[color,box-shadow]",
			"focus-visible:ring-[3px] focus-visible:outline-1 disabled:pointer-events-none disabled:opacity-50",
			"data-[pui-tabs-state=active]:shadow-sm [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4",
			props.Class,
		)),
		html.AData("pui-tabs-trigger", ""),
		html.AData("pui-tabs-id", props.TabsID),
		html.AData("pui-tabs-value", props.Value),
		html.AData("pui-tabs-state", stateAttr(props.IsActive)),
	}
	if props.ID != "" {
		buttonArgs = append(buttonArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		buttonArgs = append(buttonArgs, attr)
	}

	buttonArgs = append(buttonArgs, args...)

	return html.Button(buttonArgs...)
}

func Content(props ContentProps, args ...html.DivArg) html.Node {
	if props.Value == "" {
		return html.Span(html.AClass("text-xs text-destructive"), html.Text("tabs.Content requires Value"))
	}

	divArgs := []html.DivArg{
		html.AClass(classnames.Merge("flex-1 outline-none", hiddenClass(!props.IsActive), props.Class)),
		html.AData("pui-tabs-content", ""),
		html.AData("pui-tabs-value", props.Value),
		html.AData("pui-tabs-state", stateAttr(props.IsActive)),
	}
	if props.TabsID != "" {
		divArgs = append(divArgs, html.AData("pui-tabs-id", props.TabsID))
	}

	if props.ID != "" {
		divArgs = append(divArgs, html.AId(props.ID))
	}

	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}

	divArgs = append(divArgs, args...)

	return html.Div(divArgs...)
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
