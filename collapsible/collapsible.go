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

func Collapsible(props Props, args ...html.DivArg) html.Node {
	if props.ID == "" {
		props.ID = randomID("collapsible")
	}

	state := "closed"
	if props.Open {
		state = "open"
	}

	divArgs := []html.DivArg{
		html.AId(props.ID),
		html.AClass(classnames.Merge("", props.Class)),
		html.AData("pui-collapsible", "root"),
		html.AData("pui-collapsible-state", state),
	}
	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}
	divArgs = append(divArgs, args...)

	node := html.Div(divArgs...)
	return node.WithAssets("", collapsibleJS, "ui-collapsible")
}

func Trigger(props TriggerProps, args ...html.DivArg) html.Node {
	divArgs := []html.DivArg{
		html.AClass(classnames.Merge("", props.Class)),
		html.AData("pui-collapsible", "trigger"),
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

func Content(props ContentProps, args ...html.DivArg) html.Node {
	wrapperArgs := []html.DivArg{
		html.AClass(classnames.Merge(
			"grid grid-rows-[0fr] transition-[grid-template-rows] duration-200 ease-out [[data-pui-collapsible-state=open]_&]:grid-rows-[1fr]",
			props.Class,
		)),
		html.AData("pui-collapsible", "content"),
	}
	if props.ID != "" {
		wrapperArgs = append(wrapperArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		wrapperArgs = append(wrapperArgs, attr)
	}

	innerArgs := []html.DivArg{html.AClass("overflow-hidden")}
	innerArgs = append(innerArgs, args...)

	wrapperArgs = append(wrapperArgs, html.Div(innerArgs...))
	return html.Div(wrapperArgs...)
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
