package toast

import (
	"crypto/rand"
	"encoding/hex"
	"strconv"

	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/button"
	"github.com/plainkit/ui/internal/classnames"
)

type Variant string
type Position string

type Props struct {
	ID            string
	Class         string
	Attrs         []html.Global
	Title         string
	Description   string
	Variant       Variant
	Position      Position
	Duration      int
	Dismissible   bool
	ShowIndicator bool
	Icon          bool
}

type TriggerProps struct {
	ID    string
	Class string
	Attrs []html.Global
	Toast Props // The toast configuration to spawn
}

const (
	VariantDefault Variant = "default"
	VariantSuccess Variant = "success"
	VariantError   Variant = "error"
	VariantWarning Variant = "warning"
	VariantInfo    Variant = "info"
)

const (
	PositionTopRight     Position = "top-right"
	PositionTopLeft      Position = "top-left"
	PositionTopCenter    Position = "top-center"
	PositionBottomRight  Position = "bottom-right"
	PositionBottomLeft   Position = "bottom-left"
	PositionBottomCenter Position = "bottom-center"
)

// Toast renders an interactive toast notification container with optional auto-dismiss logic.
func Toast(props Props, args ...html.DivArg) html.Node {
	variant := props.Variant
	if variant == "" {
		variant = VariantDefault
	}
	position := props.Position
	if position == "" {
		position = PositionBottomRight
	}
	duration := props.Duration
	if duration == 0 {
		duration = 3000
	}

	id := props.ID
	if id == "" {
		id = randomID("toast")
	}

	divArgs := []html.DivArg{
		html.AId(id),
		html.AClass(classnames.Merge(
			"z-50 fixed pointer-events-auto p-4 w-full md:max-w-[420px]",
			"animate-in fade-in slide-in-from-bottom-4 duration-300",
			"data-[position=top-right]:top-0 data-[position=top-right]:right-0",
			"data-[position=top-left]:top-0 data-[position=top-left]:left-0",
			"data-[position=top-center]:top-0 data-[position=top-center]:left-1/2 data-[position=top-center]:-translate-x-1/2",
			"data-[position=bottom-right]:bottom-0 data-[position=bottom-right]:right-0",
			"data-[position=bottom-left]:bottom-0 data-[position=bottom-left]:left-0",
			"data-[position=bottom-center]:bottom-0 data-[position=bottom-center]:left-1/2 data-[position=bottom-center]:-translate-x-1/2",
			props.Class,
		)),
		html.AData("pui-toast", ""),
		html.AData("pui-toast-duration", strconv.Itoa(duration)),
		html.AData("position", string(position)),
		html.AData("variant", string(variant)),
	}
	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}

	innerDivArgs := []html.DivArg{}

	if props.ShowIndicator && duration > 0 {
		progressBar := html.Div(
			html.AClass("absolute top-0 left-0 right-0 h-1 overflow-hidden"),
			html.Div(
				html.AClass(classnames.Merge(
					"toast-progress h-full origin-left transition-transform ease-linear",
					"data-[variant=default]:bg-gray-500",
					"data-[variant=success]:bg-green-500",
					"data-[variant=error]:bg-red-500",
					"data-[variant=warning]:bg-yellow-500",
					"data-[variant=info]:bg-blue-500",
				)),
				html.AData("variant", string(variant)),
				html.AData("duration", strconv.Itoa(duration)),
			),
		)
		innerDivArgs = append(innerDivArgs, progressBar)
	}

	contentDivArgs := []html.DivArg{html.AClass("w-full bg-popover text-popover-foreground rounded-lg shadow-xs border pt-5 pb-4 px-4 flex items-center justify-center relative overflow-hidden group gap-3")}
	if props.Icon && variant != VariantDefault {
		iconNode := variantIcon(variant)
		contentDivArgs = append(contentDivArgs, iconNode)
	}

	textSpanArgs := []html.SpanArg{html.AClass("flex-1 min-w-0")}
	if props.Title != "" {
		textSpanArgs = append(textSpanArgs, html.P(
			html.AClass("text-sm font-semibold truncate"),
			html.Text(props.Title),
		))
	}
	if props.Description != "" {
		textSpanArgs = append(textSpanArgs, html.P(
			html.AClass("text-sm opacity-90 mt-1"),
			html.Text(props.Description),
		))
	}
	textContainer := html.Span(textSpanArgs...)
	contentDivArgs = append(contentDivArgs, textContainer)

	if props.Dismissible {
		btn := button.Button(button.Props{
			Variant: button.VariantGhost,
			Size:    button.SizeIcon,
			Attrs: []html.Global{
				html.AAria("label", "Close"),
				html.AData("pui-toast-dismiss", ""),
			},
		}, lucide.X(
			html.AClass("size-4 opacity-75 hover:opacity-100"),
		))
		contentDivArgs = append(contentDivArgs, btn)
	}

	contentDiv := html.Div(contentDivArgs...)
	innerDivArgs = append(innerDivArgs, contentDiv)

	divArgs = append(divArgs, innerDivArgs...)
	divArgs = append(divArgs, args...)

	node := html.Div(divArgs...)
	return node.WithAssets("", toastJS, "ui-toast")
}

func variantIcon(variant Variant) html.Node {
	switch variant {
	case VariantSuccess:
		return lucide.CircleCheck(html.AClass("size-5 text-green-500 flex-shrink-0"))
	case VariantError:
		return lucide.CircleX(html.AClass("size-5 text-destructive flex-shrink-0"))
	case VariantWarning:
		return lucide.TriangleAlert(html.AClass("size-5 text-yellow-500 flex-shrink-0"))
	case VariantInfo:
		return lucide.Info(html.AClass("size-5 text-blue-500 flex-shrink-0"))
	default:
		// Return an empty span as a placeholder
		return html.Span()
	}
}

func randomID(prefix string) string {
	buf := make([]byte, 4)
	if _, err := rand.Read(buf); err != nil {
		return prefix + "-id"
	}
	return prefix + "-" + hex.EncodeToString(buf)
}

// ToastTrigger creates a button that spawns a toast when clicked
func ToastTrigger(props TriggerProps, buttonProps button.Props, args ...html.ButtonArg) html.Node {
	id := props.ID
	if id == "" {
		id = randomID("toast-trigger")
	}

	// Encode toast configuration in data attributes
	attrs := buttonProps.Attrs
	if attrs == nil {
		attrs = []html.Global{}
	}
	attrs = append(attrs,
		html.AData("pui-toast-trigger", ""),
		html.AData("toast-title", props.Toast.Title),
		html.AData("toast-description", props.Toast.Description),
		html.AData("toast-variant", string(props.Toast.Variant)),
		html.AData("toast-position", string(props.Toast.Position)),
		html.AData("toast-duration", strconv.Itoa(props.Toast.Duration)),
		html.AData("toast-dismissible", strconv.FormatBool(props.Toast.Dismissible)),
		html.AData("toast-show-indicator", strconv.FormatBool(props.Toast.ShowIndicator)),
		html.AData("toast-icon", strconv.FormatBool(props.Toast.Icon)),
	)

	attrs = append(attrs, props.Attrs...)

	buttonProps.Attrs = attrs
	buttonProps.ID = id
	if props.Class != "" {
		buttonProps.Class = classnames.Merge(buttonProps.Class, props.Class)
	}

	return button.Button(buttonProps, args...).WithAssets("", toastJS, "ui-toast")
}

// ToastContainer creates a container for toasts to be spawned into
func ToastContainer(position Position) html.Node {
	if position == "" {
		position = PositionBottomRight
	}

	positionClasses := map[Position]string{
		PositionTopRight:     "top-0 right-0",
		PositionTopLeft:      "top-0 left-0",
		PositionTopCenter:    "top-0 left-1/2 -translate-x-1/2",
		PositionBottomRight:  "bottom-0 right-0",
		PositionBottomLeft:   "bottom-0 left-0",
		PositionBottomCenter: "bottom-0 left-1/2 -translate-x-1/2",
	}

	return html.Div(
		html.AId("toast-container-"+string(position)),
		html.AClass(classnames.Merge(
			"fixed z-50 pointer-events-none p-4",
			positionClasses[position],
		)),
		html.AData("pui-toast-container", string(position)),
	).WithAssets("", toastJS, "ui-toast")
}

const toastJS = `(function(){
  "use strict";
  var toastTimers = new Map();

  // Initialize toast containers if not present
  function ensureToastContainer(position){
    position = position || 'bottom-right';
    var containerId = 'toast-container-' + position;
    var container = document.getElementById(containerId);
    if(!container){
      container = document.createElement('div');
      container.id = containerId;
      container.dataset.tuiToastContainer = position;
      var posClasses = {
        'top-right': 'top-0 right-0',
        'top-left': 'top-0 left-0',
        'top-center': 'top-0 left-1/2 -translate-x-1/2',
        'bottom-right': 'bottom-0 right-0',
        'bottom-left': 'bottom-0 left-0',
        'bottom-center': 'bottom-0 left-1/2 -translate-x-1/2'
      };
      container.className = 'fixed z-50 pointer-events-none p-4 ' + (posClasses[position] || posClasses['bottom-right']);
      document.body.appendChild(container);
    }
    return container;
  }

  // Create toast element from configuration
  function createToastElement(config){
    var toast = document.createElement('div');
    toast.dataset.tuiToast = '';
    toast.dataset.variant = config.variant || 'default';
    toast.dataset.position = config.position || 'bottom-right';
    toast.dataset.tuiToastDuration = config.duration || '3000';
    toast.className = 'pointer-events-auto p-4 w-full md:max-w-[420px] animate-in fade-in slide-in-from-bottom-4 duration-300 mb-4';

    var inner = document.createElement('div');
    inner.className = 'w-full bg-popover text-popover-foreground rounded-lg shadow-xs border pt-5 pb-4 px-4 flex items-center justify-center relative overflow-hidden group gap-3';

    // Add progress indicator
    if(config.showIndicator === 'true' && parseInt(config.duration) > 0){
      var progressWrapper = document.createElement('div');
      progressWrapper.className = 'absolute top-0 left-0 right-0 h-1 overflow-hidden';
      var progress = document.createElement('div');
      progress.className = 'toast-progress h-full origin-left transition-transform ease-linear ';
      var variantColors = {
        'default': 'bg-gray-500',
        'success': 'bg-green-500',
        'error': 'bg-red-500',
        'warning': 'bg-yellow-500',
        'info': 'bg-blue-500'
      };
      progress.className += variantColors[config.variant] || variantColors['default'];
      progress.dataset.variant = config.variant || 'default';
      progress.dataset.duration = config.duration || '3000';
      progressWrapper.appendChild(progress);
      inner.appendChild(progressWrapper);
    }

    // Add icon
    if(config.icon === 'true' && config.variant !== 'default'){
      var iconWrapper = document.createElement('div');
      iconWrapper.className = 'size-5 flex-shrink-0';
      var iconColors = {
        'success': 'text-green-500',
        'error': 'text-destructive',
        'warning': 'text-yellow-500',
        'info': 'text-blue-500'
      };
      iconWrapper.className += ' ' + (iconColors[config.variant] || '');
      // Simple icon representation (you'd need actual SVG icons here)
      iconWrapper.innerHTML = getIconSvg(config.variant);
      inner.appendChild(iconWrapper);
    }

    // Add text content
    var textContainer = document.createElement('span');
    textContainer.className = 'flex-1 min-w-0';
    if(config.title){
      var title = document.createElement('p');
      title.className = 'text-sm font-semibold truncate';
      title.textContent = config.title;
      textContainer.appendChild(title);
    }
    if(config.description){
      var desc = document.createElement('p');
      desc.className = 'text-sm opacity-90 mt-1';
      desc.textContent = config.description;
      textContainer.appendChild(desc);
    }
    inner.appendChild(textContainer);

    // Add dismiss button
    if(config.dismissible === 'true'){
      var btn = document.createElement('button');
      btn.type = 'button';
      btn.className = 'inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 hover:bg-accent hover:text-accent-foreground size-8 p-0';
      btn.setAttribute('aria-label', 'Close');
      btn.dataset.tuiToastDismiss = '';
      btn.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="opacity-75 hover:opacity-100"><path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg>';
      inner.appendChild(btn);
    }

    toast.appendChild(inner);
    return toast;
  }

  function getIconSvg(variant){
    var icons = {
      'success': '<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><path d="m9 12 2 2 4-4"/></svg>',
      'error': '<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><path d="m15 9-6 6"/><path d="m9 9 6 6"/></svg>',
      'warning': '<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m21.73 18-8-14a2 2 0 0 0-3.48 0l-8 14A2 2 0 0 0 4 21h16a2 2 0 0 0 1.73-3"/><path d="M12 9v4"/><path d="M12 17h.01"/></svg>',
      'info': '<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg>'
    };
    return icons[variant] || '';
  }

  function spawnToast(config){
    var container = ensureToastContainer(config.position);
    var toast = createToastElement(config);
    container.appendChild(toast);
    setupToast(toast);
  }

  function setupToast(toast){
    var duration = parseInt(toast.dataset.tuiToastDuration || "3000", 10);
    var progress = toast.querySelector(".toast-progress");
    var state = { timer: null, startTime: Date.now(), remaining: duration, paused: false };
    toastTimers.set(toast, state);
    if(progress && duration > 0){
      progress.style.transitionDuration = duration + "ms";
      requestAnimationFrame(function(){ progress.style.transform = "scaleX(0)"; });
    }
    if(duration > 0){
      state.timer = setTimeout(function(){ dismissToast(toast); }, duration);
    }
    toast.addEventListener("mouseenter", function(){
      var st = toastTimers.get(toast);
      if(!st || st.paused) return;
      clearTimeout(st.timer);
      st.remaining = st.remaining - (Date.now() - st.startTime);
      st.paused = true;
      if(progress){
        var computed = getComputedStyle(progress);
        progress.style.transitionDuration = "0ms";
        progress.style.transform = computed.transform;
      }
    });
    toast.addEventListener("mouseleave", function(){
      var st = toastTimers.get(toast);
      if(!st || !st.paused || st.remaining <= 0) return;
      st.startTime = Date.now();
      st.paused = false;
      st.timer = setTimeout(function(){ dismissToast(toast); }, st.remaining);
      if(progress){
        progress.style.transitionDuration = st.remaining + "ms";
        progress.style.transform = "scaleX(0)";
      }
    });
  }
  function dismissToast(toast){
    toastTimers.delete(toast);
    toast.style.transition = "opacity 300ms, transform 300ms";
    toast.style.opacity = "0";
    toast.style.transform = "translateY(1rem)";
    setTimeout(function(){ toast.remove(); }, 300);
  }
  document.addEventListener("click", function(e){
    var dismissBtn = e.target.closest('[data-pui-toast-dismiss]');
    if(dismissBtn){
      var toast = dismissBtn.closest('[data-pui-toast]');
      if(toast) dismissToast(toast);
    }
  });
  // Handle toast trigger clicks
  document.addEventListener('click', function(e){
    var trigger = e.target.closest('[data-pui-toast-trigger]');
    if(trigger){
      e.preventDefault();
      var config = {
        title: trigger.dataset.toastTitle,
        description: trigger.dataset.toastDescription,
        variant: trigger.dataset.toastVariant || 'default',
        position: trigger.dataset.toastPosition || 'bottom-right',
        duration: trigger.dataset.toastDuration || '3000',
        dismissible: trigger.dataset.toastDismissible || 'true',
        showIndicator: trigger.dataset.toastShowIndicator || 'true',
        icon: trigger.dataset.toastIcon || 'true'
      };
      spawnToast(config);
    }
  });

  // Setup any existing toasts on page load
  document.querySelectorAll('[data-pui-toast]').forEach(function(toast){
    if(!toast.hasAttribute('data-pui-toast-template')){
      setupToast(toast);
    }
  });
})();`
