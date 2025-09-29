package inputotp

import (
	_ "embed"
	"strconv"

	"github.com/plainkit/html"
	"github.com/plainkit/ui/internal/classnames"
)

type Props struct {
	ID        string
	Class     string
	Attrs     []html.Global
	Value     string
	Required  bool
	Name      string
	Form      string
	HasError  bool
	Autofocus bool
}

type GroupProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type SlotProps struct {
	ID          string
	Class       string
	Attrs       []html.Global
	Index       int
	Type        string
	Placeholder string
	Disabled    bool
	HasError    bool
}

type SeparatorProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

func InputOTP(props Props, args ...html.DivArg) html.Node {
	containerID := ""
	if props.ID != "" {
		containerID = props.ID + "-container"
	}

	divArgs := []html.DivArg{
		html.AClass(classnames.Merge("flex flex-row items-center gap-2 w-fit", props.Class)),
		html.AData("pui-inputotp", ""),
	}
	if containerID != "" {
		divArgs = append(divArgs, html.AId(containerID))
	}
	if props.Value != "" {
		divArgs = append(divArgs, html.AData("pui-inputotp-value", props.Value))
	}
	if props.Autofocus {
		divArgs = append(divArgs, html.AAutofocus())
	}
	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}

	hiddenArgs := []html.InputArg{
		html.AType("hidden"),
		html.AData("pui-inputotp-value-target", ""),
	}
	if props.ID != "" {
		hiddenArgs = append(hiddenArgs, html.AId(props.ID))
	}
	if props.Name != "" {
		hiddenArgs = append(hiddenArgs, html.AName(props.Name))
	}
	if props.Form != "" {
		hiddenArgs = append(hiddenArgs, html.AForm(props.Form))
	}
	if props.HasError {
		hiddenArgs = append(hiddenArgs, html.AAria("invalid", "true"))
	}
	if props.Required {
		hiddenArgs = append(hiddenArgs, html.ARequired())
	}

	hidden := html.Input(hiddenArgs...)
	divArgs = append(divArgs, hidden)
	divArgs = append(divArgs, args...)

	node := html.Div(divArgs...)
	return node.WithAssets("", inputOTPJS, "ui-inputotp")
}

func Group(props GroupProps, args ...html.DivArg) html.Node {
	groupArgs := []html.DivArg{html.AClass(classnames.Merge("flex gap-2", props.Class))}
	if props.ID != "" {
		groupArgs = append(groupArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		groupArgs = append(groupArgs, attr)
	}
	groupArgs = append(groupArgs, args...)
	return html.Div(groupArgs...)
}

func Slot(props SlotProps) html.Node {
	inputType := props.Type
	if inputType == "" {
		inputType = "text"
	}

	wrapperArgs := []html.DivArg{html.AClass("relative")}
	if props.ID != "" {
		wrapperArgs = append(wrapperArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		wrapperArgs = append(wrapperArgs, attr)
	}

	classes := []string{
		"w-10 h-12 text-center rounded-md border border-input bg-transparent text-base shadow-xs transition-[color,box-shadow] outline-none md:text-sm",
		"dark:bg-input/30",
		"selection:bg-primary selection:text-primary-foreground",
		"placeholder:text-muted-foreground",
		"focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px]",
		"disabled:pointer-events-none disabled:cursor-not-allowed disabled:opacity-50",
		"aria-invalid:ring-destructive/20 aria-invalid:border-destructive dark:aria-invalid:ring-destructive/40",
	}
	if props.HasError {
		classes = append(classes, "border-destructive ring-destructive/20 dark:ring-destructive/40")
	}
	classes = append(classes, props.Class)

	inputArgs := []html.InputArg{
		html.AType(inputType),
		html.AInputmode("numeric"),
		html.AMaxlength("1"),
		html.AClass(classnames.Merge(classes...)),
		html.AData("pui-inputotp-index", strconv.Itoa(props.Index)),
		html.AData("pui-inputotp-slot", ""),
	}
	if props.Placeholder != "" {
		inputArgs = append(inputArgs, html.APlaceholder(props.Placeholder))
	}
	if props.Disabled {
		inputArgs = append(inputArgs, html.ADisabled())
	}
	if props.HasError {
		inputArgs = append(inputArgs, html.AAria("invalid", "true"))
	}

	return html.Div(append(wrapperArgs, html.Input(inputArgs...))...)
}

func Separator(props SeparatorProps, args ...html.DivArg) html.Node {
	divArgs := []html.DivArg{html.AClass(classnames.Merge("flex items-center text-muted-foreground text-xl", props.Class))}
	if props.ID != "" {
		divArgs = append(divArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}
	if len(args) == 0 {
		divArgs = append(divArgs, html.Span(html.Text("-")))
	} else {
		divArgs = append(divArgs, args...)
	}
	return html.Div(divArgs...)
}

//go:embed inputotp.js
var inputOTPJS string
