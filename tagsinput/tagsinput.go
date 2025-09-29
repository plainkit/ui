package tagsinput

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/badge"
	"github.com/plainkit/ui/input"
	"github.com/plainkit/ui/internal/classnames"
)

type Props struct {
	ID          string
	Name        string
	Value       []string
	Form        string
	Placeholder string
	Class       string
	Attrs       []html.Global
	HasError    bool
	Disabled    bool
	Readonly    bool
}

// TagsInput renders a tags input component for entering multiple tags
func TagsInput(props Props, args ...html.DivArg) html.Node {
	id := props.ID
	if id == "" {
		id = randomID("tagsinput")
	}

	containerArgs := []html.DivArg{
		html.AId(id + "-container"),
		html.AClass(classnames.Merge(
			// Base styles
			"flex items-center flex-wrap gap-2 p-2 rounded-md border border-input bg-transparent shadow-xs transition-[color,box-shadow] outline-none",
			// Dark mode background
			"dark:bg-input/30",
			// Focus styles
			"focus-within:border-ring focus-within:ring-ring/50 focus-within:ring-[3px]",
			// Disabled styles
			func() string {
				if props.Disabled {
					return "opacity-50 cursor-not-allowed"
				}
				return ""
			}(),
			// Width
			"w-full",
			// Error/Invalid styles
			func() string {
				if props.HasError {
					return "border-destructive ring-destructive/20 dark:ring-destructive/40"
				}
				return ""
			}(),
			props.Class,
		)),
		html.AData("pui-tagsinput", ""),
		html.AData("pui-tagsinput-name", props.Name),
		html.AData("pui-tagsinput-form", props.Form),
	}

	for _, attr := range props.Attrs {
		containerArgs = append(containerArgs, attr)
	}

	// Add existing tags as children
	tagChildren := make([]html.DivArg, 0, len(props.Value))
	for _, tag := range props.Value {
		tagBadge := badge.Badge(badge.Props{
			Attrs: []html.Global{
				html.AData("pui-tagsinput-chip", ""),
			},
		},
			html.Span(html.Text(tag)),
			html.Button(
				html.AType("button"),
				html.AClass("ml-1 text-current hover:text-destructive disabled:opacity-50 disabled:cursor-not-allowed cursor-pointer"),
				html.AData("pui-tagsinput-remove", ""),
				func() html.ButtonArg {
					if props.Disabled {
						return html.ADisabled()
					}
					return html.AAria("", "")
				}(),
				lucide.X(html.AClass("h-3 w-3 pointer-events-none")),
			),
		)
		tagChildren = append(tagChildren, tagBadge)
	}

	// Build the full tagsContainer args
	tagsContainerArgs := []html.DivArg{
		html.AClass("flex items-center flex-wrap gap-2"),
		html.AData("pui-tagsinput-container", ""),
	}
	tagsContainerArgs = append(tagsContainerArgs, tagChildren...)
	tagsContainer := html.Div(tagsContainerArgs...)

	// Text input
	textInput := input.Input(input.Props{
		ID:          id,
		Class:       "border-0 shadow-none focus-visible:ring-0 h-auto py-0 px-0 bg-transparent rounded-none min-h-0 disabled:opacity-100 dark:bg-transparent",
		Type:        input.TypeText,
		Placeholder: props.Placeholder,
		Disabled:    props.Disabled,
		Readonly:    props.Readonly,
		Attrs: []html.Global{
			html.AData("pui-tagsinput-text-input", ""),
		},
	})

	// Add existing hidden inputs
	hiddenInputChildren := make([]html.DivArg, 0, len(props.Value))
	for _, tag := range props.Value {
		hiddenInput := html.Input(
			html.AType("hidden"),
			html.AName(props.Name),
			html.AValue(tag),
		)
		hiddenInputChildren = append(hiddenInputChildren, hiddenInput)
	}

	// Build the full hiddenInputsContainer args
	hiddenInputsArgs := []html.DivArg{
		html.AData("pui-tagsinput-hidden-inputs", ""),
	}
	hiddenInputsArgs = append(hiddenInputsArgs, hiddenInputChildren...)
	hiddenInputsContainer := html.Div(hiddenInputsArgs...)

	containerArgs = append(containerArgs,
		tagsContainer,
		textInput,
		hiddenInputsContainer,
	)
	containerArgs = append(containerArgs, args...)

	return html.Div(containerArgs...).WithAssets("", tagsinputJS, "ui-tagsinput")
}

func randomID(prefix string) string {
	buf := make([]byte, 4)
	if _, err := rand.Read(buf); err != nil {
		return prefix + "-id"
	}
	return prefix + "-" + hex.EncodeToString(buf)
}

const tagsinputJS = `(function () {
  'use strict';

  // Create tag chip element
  function createTagChip(tagValue, isDisabled) {
    const tagChip = document.createElement('div');
    tagChip.setAttribute('data-pui-tagsinput-chip', '');
    tagChip.className = 'inline-flex items-center gap-2 rounded-md border px-2.5 py-0.5 text-xs font-semibold transition-colors focus:outline-hidden focus:ring-2 focus:ring-ring focus:ring-offset-2 border-transparent bg-primary text-primary-foreground';

    tagChip.innerHTML =
      '<span>' + tagValue + '</span>' +
      '<button type="button"' +
              ' class="ml-1 text-current hover:text-destructive disabled:opacity-50 disabled:cursor-not-allowed cursor-pointer"' +
              ' data-pui-tagsinput-remove=""' +
              (isDisabled ? ' disabled' : '') + '>' +
        '<svg xmlns="http://www.w3.org/2000/svg" class="h-3 w-3 pointer-events-none" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">' +
          '<path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />' +
        '</svg>' +
      '</button>';

    return tagChip;
  }

  // Add tag
  function addTag(container, value) {
    const textInput = container.querySelector('[data-pui-tagsinput-text-input]');
    if (textInput && textInput.hasAttribute('disabled')) return;

    const tagValue = value.trim();
    if (!tagValue) return;

    const hiddenInputsContainer = container.querySelector('[data-pui-tagsinput-hidden-inputs]');
    const tagsContainer = container.querySelector('[data-pui-tagsinput-container]');
    const name = container.getAttribute('data-pui-tagsinput-name');
    const form = container.getAttribute('data-pui-tagsinput-form');

    // Check for duplicates
    const existingTags = hiddenInputsContainer.querySelectorAll('input[type="hidden"]');
    for (const t of existingTags) {
      if (t.value.toLowerCase() === tagValue.toLowerCase()) {
        textInput.value = '';
        return;
      }
    }

    // Add tag chip and hidden input
    const tagChip = createTagChip(tagValue, textInput && textInput.hasAttribute('disabled'));
    tagsContainer.appendChild(tagChip);

    const hiddenInput = document.createElement('input');
    hiddenInput.type = 'hidden';
    hiddenInput.name = name;
    hiddenInput.value = tagValue;
    if (form !== null && form !== "") {
      hiddenInput.setAttribute("form", form);
    }

    hiddenInputsContainer.appendChild(hiddenInput);

    textInput.value = '';
  }

  // Remove tag
  function removeTag(button) {
    const tagChip = button.closest('[data-pui-tagsinput-chip]');
    if (!tagChip) return;

    const container = tagChip.closest('[data-pui-tagsinput]');
    const tagValue = tagChip.querySelector('span').textContent.trim();
    const hiddenInputsContainer = container.querySelector('[data-pui-tagsinput-hidden-inputs]');

    const hiddenInput = hiddenInputsContainer.querySelector('input[type="hidden"][value="' + tagValue + '"]');
    if (hiddenInput) hiddenInput.remove();

    tagChip.remove();
  }

  // Event delegation
  document.addEventListener('keydown', (e) => {
    const textInput = e.target.closest('[data-pui-tagsinput-text-input]');
    if (!textInput) return;

    const container = textInput.closest('[data-pui-tagsinput]');
    if (!container) return;

    if (e.key === 'Enter' || e.key === ',') {
      e.preventDefault();
      addTag(container, textInput.value);
    } else if (e.key === 'Backspace' && textInput.value === '') {
      e.preventDefault();
      const lastChip = container.querySelector('[data-pui-tagsinput-chip]:last-child');
      const removeButton = lastChip && lastChip.querySelector('[data-pui-tagsinput-remove]');
      if (removeButton && !removeButton.disabled) {
        removeTag(removeButton);
      }
    }
  });

  document.addEventListener('click', (e) => {
    // Handle remove button clicks
    const removeButton = e.target.closest('[data-pui-tagsinput-remove]');
    if (removeButton && !removeButton.disabled) {
      e.preventDefault();
      e.stopPropagation();
      removeTag(removeButton);
      return;
    }

    // Focus input when clicking container
    const container = e.target.closest('[data-pui-tagsinput]');
    if (container && !e.target.closest('input')) {
      const textInput = container.querySelector('[data-pui-tagsinput-text-input]');
      if (textInput) textInput.focus();
    }
  });

  // Form reset
  document.addEventListener('reset', (e) => {
    if (!e.target.matches('form')) return;

    e.target.querySelectorAll('[data-pui-tagsinput]').forEach(container => {
      container.querySelectorAll('[data-pui-tagsinput-chip]').forEach(chip => chip.remove());
      container.querySelectorAll('[data-pui-tagsinput-hidden-inputs] input[type="hidden"]').forEach(input => input.remove());
      const textInput = container.querySelector('[data-pui-tagsinput-text-input]');
      if (textInput) textInput.value = '';
    });
  });
})();`
