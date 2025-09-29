package button

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/internal/classnames"
)

type Variant string
type Size string
type Type string

type Props struct {
	ID        string
	Class     string
	Attrs     []html.Global
	Variant   Variant
	Size      Size
	FullWidth bool
	Href      string
	Target    string
	Disabled  bool
	Type      Type
	Form      string
}

const (
	VariantDefault     Variant = "default"
	VariantDestructive Variant = "destructive"
	VariantOutline     Variant = "outline"
	VariantSecondary   Variant = "secondary"
	VariantGhost       Variant = "ghost"
	VariantLink        Variant = "link"
)

const (
	SizeDefault Size = "default"
	SizeSm      Size = "sm"
	SizeLg      Size = "lg"
	SizeIcon    Size = "icon"
)

const (
	TypeButton Type = "button"
	TypeReset  Type = "reset"
	TypeSubmit Type = "submit"
)

// Button renders a button (or link if Href is provided) using the given props and child arguments.
// Children should be html.ButtonArg values such as html.T("Label"), html.Span(...), etc.
func Button(props Props, args ...html.ButtonArg) html.Node {
	if props.Type == "" {
		props.Type = TypeButton
	}

	className := classnames.Merge(
		"inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-all",
		"disabled:pointer-events-none disabled:opacity-50 [&_svg]:pointer-events-none [&_svg:not([class*='size-'])]:size-4 shrink-0 [&_svg]:shrink-0",
		"outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px]",
		"aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive",
		"cursor-pointer",
		variantClass(props.Variant),
		sizeClass(props.Size),
		modifierClass(props.FullWidth),
		props.Class,
	)

	if props.Href != "" && !props.Disabled {
		return renderAnchor(props, className, args)
	}
	return renderButton(props, className, args)
}

func renderButton(props Props, className string, args []html.ButtonArg) html.Node {
	btnArgs := []html.ButtonArg{html.AClass(className)}

	if props.ID != "" {
		btnArgs = append(btnArgs, html.AId(props.ID))
	}
	if props.Type != "" {
		btnArgs = append(btnArgs, html.AType(string(props.Type)))
	}
	if props.Form != "" {
		btnArgs = append(btnArgs, html.AForm(props.Form))
	}
	if props.Disabled {
		btnArgs = append(btnArgs, html.ADisabled())
	}

	for _, attr := range props.Attrs {
		btnArgs = append(btnArgs, attr)
	}
	btnArgs = append(btnArgs, args...)

	return html.Button(btnArgs...)
}

func renderAnchor(props Props, className string, args []html.ButtonArg) html.Node {
	anchorArgs := []html.AArg{
		html.AHref(props.Href),
		html.AClass(className),
	}

	if props.ID != "" {
		anchorArgs = append(anchorArgs, html.AId(props.ID))
	}
	if props.Target != "" {
		anchorArgs = append(anchorArgs, html.ATarget(props.Target))
	}

	for _, attr := range props.Attrs {
		anchorArgs = append(anchorArgs, attr)
	}

	for _, arg := range args {
		if aArg, ok := arg.(html.AArg); ok {
			anchorArgs = append(anchorArgs, aArg)
		}
	}

	return html.A(anchorArgs...)
}

func variantClass(v Variant) string {
	switch v {
	case VariantDestructive:
		return "bg-destructive text-destructive-foreground shadow-xs hover:bg-destructive/90 focus-visible:ring-destructive/20 dark:focus-visible:ring-destructive/40 dark:bg-destructive/60"
	case VariantOutline:
		return "border bg-background shadow-xs hover:bg-accent hover:text-accent-foreground dark:bg-input/30 dark:border-input dark:hover:bg-input/50"
	case VariantSecondary:
		return "bg-secondary text-secondary-foreground shadow-xs hover:bg-secondary/80"
	case VariantGhost:
		return "hover:bg-accent hover:text-accent-foreground dark:hover:bg-accent/50"
	case VariantLink:
		return "text-primary underline-offset-4 hover:underline"
	default:
		return "bg-primary text-primary-foreground shadow-xs hover:bg-primary/90"
	}
}

func sizeClass(s Size) string {
	switch s {
	case SizeSm:
		return "h-8 rounded-md gap-1.5 px-3 has-[>svg]:px-2.5"
	case SizeLg:
		return "h-10 rounded-md px-6 has-[>svg]:px-4"
	case SizeIcon:
		return "size-9"
	default:
		return "h-9 px-4 py-2 has-[>svg]:px-3"
	}
}

func modifierClass(fullWidth bool) string {
	if fullWidth {
		return "w-full"
	}
	return ""
}
