package input

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/button"
	"github.com/plainkit/ui/internal/classnames"
)

type Type string

const (
	TypeText     Type = "text"
	TypePassword Type = "password"
	TypeEmail    Type = "email"
	TypeNumber   Type = "number"
	TypeTel      Type = "tel"
	TypeURL      Type = "url"
	TypeSearch   Type = "search"
	TypeDate     Type = "date"
	TypeDateTime Type = "datetime-local"
	TypeTime     Type = "time"
	TypeFile     Type = "file"
	TypeColor    Type = "color"
	TypeWeek     Type = "week"
	TypeMonth    Type = "month"
)

type Props struct {
	ID                 string
	Class              string
	Attrs              []html.Global
	Name               string
	Type               Type
	Form               string
	Placeholder        string
	Value              string
	Disabled           bool
	Readonly           bool
	Required           bool
	FileAccept         string
	HasError           bool
	ShowPasswordToggle bool
}

// Input renders a styled input control with optional password toggle button.
// Additional input arguments can be provided after props (e.g. html.AAutocomplete("email")).
func Input(props Props, extra ...html.InputArg) html.Node {
	if props.Type == "" {
		props.Type = TypeText
	}
	if props.ID == "" {
		props.ID = randomID()
	}

	inputArgs := []html.InputArg{
		html.AId(props.ID),
		html.AType(string(props.Type)),
		html.AClass(inputClass(props)),
	}
	if props.Name != "" {
		inputArgs = append(inputArgs, html.AName(props.Name))
	}
	if props.Placeholder != "" {
		inputArgs = append(inputArgs, html.APlaceholder(props.Placeholder))
	}
	if props.Value != "" {
		inputArgs = append(inputArgs, html.AValue(props.Value))
	}
	if props.Type == TypeFile && props.FileAccept != "" {
		inputArgs = append(inputArgs, html.AAccept(props.FileAccept))
	}
	if props.Form != "" {
		inputArgs = append(inputArgs, html.AForm(props.Form))
	}
	if props.Disabled {
		inputArgs = append(inputArgs, html.ADisabled())
	}
	if props.Readonly {
		inputArgs = append(inputArgs, html.AReadonly())
	}
	if props.Required {
		inputArgs = append(inputArgs, html.ARequired())
	}
	if props.HasError {
		inputArgs = append(inputArgs, html.AAria("invalid", "true"))
	}
	for _, attr := range props.Attrs {
		inputArgs = append(inputArgs, attr)
	}
	inputArgs = append(inputArgs, extra...)

	children := []html.Component{html.Input(inputArgs...)}

	if props.Type == TypePassword && props.ShowPasswordToggle {
		children = append(children, passwordToggleButton(props.ID))
	}

	divArgs := []html.DivArg{html.AClass("relative w-full")}
	for _, child := range children {
		divArgs = append(divArgs, html.Child(child))
	}
	node := html.Div(divArgs...)
	if props.Type == TypePassword && props.ShowPasswordToggle {
		node = node.WithAssets("", passwordToggleJS, "ui-input-toggle")
	}
	return node
}

func inputClass(props Props) string {
	extraPadding := ""
	if props.Type == TypePassword && props.ShowPasswordToggle {
		extraPadding = "pr-8"
	}
	errorClass := ""
	if props.HasError {
		errorClass = "border-destructive ring-destructive/20 dark:ring-destructive/40"
	}

	return classnames.Merge(
		"flex h-9 w-full min-w-0 rounded-md border border-input bg-transparent px-3 py-1 text-base shadow-xs transition-[color,box-shadow] outline-none md:text-sm",
		"dark:bg-input/30",
		"selection:bg-primary selection:text-primary-foreground",
		"placeholder:text-muted-foreground",
		"file:inline-flex file:h-7 file:border-0 file:bg-transparent file:text-sm file:font-medium file:text-foreground",
		"focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px]",
		"disabled:pointer-events-none disabled:cursor-not-allowed disabled:opacity-50",
		"aria-invalid:ring-destructive/20 aria-invalid:border-destructive dark:aria-invalid:ring-destructive/40",
		errorClass,
		extraPadding,
		props.Class,
	)
}

func passwordToggleButton(inputID string) html.Node {
	openIcon := html.Span(
		html.AClass("icon-open block"),
		lucide.Eye(html.AClass("size-4")),
	)
	closedIcon := html.Span(
		html.AClass("icon-closed hidden"),
		lucide.EyeOff(html.AClass("size-4")),
	)

	return button.Button(
		button.Props{
			Size:    button.SizeIcon,
			Variant: button.VariantGhost,
			Class:   "absolute right-0 top-1/2 -translate-y-1/2 opacity-50 cursor-pointer",
			Attrs: []html.Global{
				html.AData("pui-input-toggle-password", inputID),
			},
		},
		openIcon,
		closedIcon,
	)
}

func randomID() string {
	bytes := make([]byte, 6)
	if _, err := rand.Read(bytes); err != nil {
		return "input-id"
	}
	return "input-" + hex.EncodeToString(bytes)
}

const passwordToggleJS = `(function(){
  if(typeof document === 'undefined') return;
  document.addEventListener('click',function(event){
    var btn = event.target.closest('[data-pui-input-toggle-password]');
    if(!btn) return;
    var inputId = btn.getAttribute('data-pui-input-toggle-password');
    if(!inputId) return;
    var input = document.getElementById(inputId);
    if(!input) return;
    var isPassword = input.getAttribute('type') === 'password';
    input.setAttribute('type', isPassword ? 'text' : 'password');
    var openIcon = btn.querySelector('.icon-open');
    var closedIcon = btn.querySelector('.icon-closed');
    if(openIcon){
      openIcon.classList.toggle('hidden', !isPassword);
      openIcon.classList.toggle('block', isPassword);
    }
    if(closedIcon){
      closedIcon.classList.toggle('hidden', isPassword);
      closedIcon.classList.toggle('block', !isPassword);
    }
  });
})();`
