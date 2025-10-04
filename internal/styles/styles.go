package styles

import "github.com/plainkit/html"

const (
	surfaceBase          = "rounded-2xl border border-border/60 bg-card/95 text-card-foreground shadow-lg transition-colors supports-[backdrop-filter]:bg-card/80 backdrop-blur-md"
	surfaceMutedBase     = "rounded-xl border border-border/40 bg-muted/80 text-muted-foreground shadow-sm"
	panelBase            = "rounded-2xl border border-border/60 bg-popover/95 text-popover-foreground shadow-xl supports-[backdrop-filter]:bg-popover/80 backdrop-blur-md"
	interactiveBase      = "inline-flex items-center justify-center gap-2 rounded-lg border border-transparent text-sm font-medium transition-all duration-200 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring/50 focus-visible:ring-offset-2 focus-visible:ring-offset-background disabled:pointer-events-none disabled:opacity-60"
	interactiveGhostBase = "inline-flex items-center justify-center gap-2 rounded-lg border border-transparent text-sm font-medium text-foreground/80 transition-all duration-200 hover:bg-muted/70 hover:text-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring/50 focus-visible:ring-offset-2 focus-visible:ring-offset-background disabled:pointer-events-none disabled:opacity-60"
	interactiveSoftBase  = "inline-flex items-center justify-center gap-2 rounded-lg border border-border/60 bg-muted/60 text-foreground/80 transition-all duration-200 hover:bg-muted focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring/50 focus-visible:ring-offset-2 focus-visible:ring-offset-background disabled:pointer-events-none disabled:opacity-60"
	inputBase            = "flex h-10 w-full min-w-0 rounded-lg border border-input/60 bg-background/60 px-3 py-2 text-sm shadow-xs transition-[border-color,box-shadow,background-color] focus-visible:border-ring focus-visible:ring-2 focus-visible:ring-ring/40 focus-visible:ring-offset-1 focus-visible:ring-offset-background placeholder:text-muted-foreground disabled:pointer-events-none disabled:opacity-60 selection:bg-primary/10 selection:text-foreground"
	controlBase          = "peer inline-flex size-4 shrink-0 cursor-pointer items-center justify-center rounded-md border border-border/60 bg-background/70 transition-all duration-200 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring/50 focus-visible:ring-offset-1 focus-visible:ring-offset-background disabled:cursor-not-allowed disabled:opacity-50"
	labelBase            = "text-sm font-medium text-muted-foreground"
	subtleTextBase       = "text-sm text-muted-foreground/80"
	tagBase              = "inline-flex items-center gap-1 rounded-full border border-border/40 bg-muted/60 px-3 py-1 text-xs font-medium tracking-wide text-muted-foreground/80"
	displayHeadingBase   = "text-2xl font-semibold tracking-tight text-foreground"
	subHeadingBase       = "text-base font-medium text-muted-foreground"
)

func merge(base string, extra ...string) string {
	return html.ClassMerge(append([]string{base}, extra...)...)
}

// Surface returns a high-emphasis surface treatment with depth and subtle blur.
func Surface(extra ...string) string {
	return merge(surfaceBase, extra...)
}

// SurfaceMuted returns a softer surface for secondary content.
func SurfaceMuted(extra ...string) string {
	return merge(surfaceMutedBase, extra...)
}

// Panel returns a floating surface style suitable for popovers, dialogs, and dropdown content.
func Panel(extra ...string) string {
	return merge(panelBase, extra...)
}

// Interactive returns the default interactive control styling (buttons, triggers).
func Interactive(extra ...string) string {
	return merge(interactiveBase, extra...)
}

// InteractiveGhost returns a low-emphasis interactive style.
func InteractiveGhost(extra ...string) string {
	return merge(interactiveGhostBase, extra...)
}

// InteractiveSoft returns a soft, outlined interactive treatment.
func InteractiveSoft(extra ...string) string {
	return merge(interactiveSoftBase, extra...)
}

// Input returns the shared styling for form inputs and pseudo-input controls.
func Input(extra ...string) string {
	return merge(inputBase, extra...)
}

// Control returns styling for binary controls such as checkboxes and radios.
func Control(extra ...string) string {
	return merge(controlBase, extra...)
}

// Label returns consistent label typography.
func Label(extra ...string) string {
	return merge(labelBase, extra...)
}

// SubtleText returns muted supporting copy styling.
func SubtleText(extra ...string) string {
	return merge(subtleTextBase, extra...)
}

// Tag returns the styling for pill-like metadata, used by badges or status chips.
func Tag(extra ...string) string {
	return merge(tagBase, extra...)
}

// DisplayHeading returns the styling for prominent headings within components.
func DisplayHeading(extra ...string) string {
	return merge(displayHeadingBase, extra...)
}

// SubHeading returns the styling for supporting headings.
func SubHeading(extra ...string) string {
	return merge(subHeadingBase, extra...)
}
