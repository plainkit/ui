package inputotp

import (
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

const inputOTPJS = `(function(){
  'use strict';
  function getSlots(container){
    return Array.from(container.querySelectorAll('[data-pui-inputotp-slot]')).sort(function(a,b){
      return parseInt(a.getAttribute('data-pui-inputotp-index')) - parseInt(b.getAttribute('data-pui-inputotp-index'));
    });
  }
  function focusSlot(slot){
    if(!slot) return;
    slot.focus();
    setTimeout(function(){ slot.select(); },0);
  }
  function updateHiddenValue(container){
    var hiddenInput = container.querySelector('[data-pui-inputotp-value-target]');
    var slots = getSlots(container);
    if(hiddenInput && slots.length){
      hiddenInput.value = slots.map(function(s){ return s.value; }).join('');
    }
  }
  function findFirstEmptySlot(container){
    var slots = getSlots(container);
    for(var i=0;i<slots.length;i++){
      if(!slots[i].value) return slots[i];
    }
    return null;
  }
  function getNextSlot(container,currentSlot){
    var slots = getSlots(container);
    var index = slots.indexOf(currentSlot);
    return index >= 0 && index < slots.length - 1 ? slots[index + 1] : null;
  }
  function getPrevSlot(container,currentSlot){
    var slots = getSlots(container);
    var index = slots.indexOf(currentSlot);
    return index > 0 ? slots[index - 1] : null;
  }
  document.addEventListener('input', function(e){
    if(!e.target.matches('[data-pui-inputotp-slot]')) return;
    var slot = e.target;
    var container = slot.closest('[data-pui-inputotp]');
    if(!container) return;
    if(slot.value === ' '){
      slot.value = '';
      return;
    }
    if(slot.value.length > 1){
      slot.value = slot.value.slice(-1);
    }
    if(slot.value){
      var nextSlot = getNextSlot(container, slot);
      if(nextSlot) focusSlot(nextSlot);
    }
    updateHiddenValue(container);
  });
  document.addEventListener('keydown', function(e){
    if(!e.target.matches('[data-pui-inputotp-slot]')) return;
    var slot = e.target;
    var container = slot.closest('[data-pui-inputotp]');
    if(!container) return;
    if(e.key === 'Backspace'){
      e.preventDefault();
      if(slot.value){
        slot.value = '';
        updateHiddenValue(container);
      } else {
        var prevSlot = getPrevSlot(container, slot);
        if(prevSlot){
          prevSlot.value = '';
          updateHiddenValue(container);
          focusSlot(prevSlot);
        }
      }
    } else if(e.key === 'ArrowLeft'){
      e.preventDefault();
      var prev = getPrevSlot(container, slot);
      if(prev) focusSlot(prev);
    } else if(e.key === 'ArrowRight'){
      e.preventDefault();
      var next = getNextSlot(container, slot);
      if(next) focusSlot(next);
    }
  });
  document.addEventListener('focus', function(e){
    if(!e.target.matches('[data-pui-inputotp-slot]')) return;
    var slot = e.target;
    var container = slot.closest('[data-pui-inputotp]');
    if(!container) return;
    var firstEmpty = findFirstEmptySlot(container);
    if(firstEmpty && firstEmpty !== slot){
      focusSlot(firstEmpty);
      return;
    }
    setTimeout(function(){ slot.select(); },0);
  }, true);
  document.addEventListener('paste', function(e){
    var slot = e.target.closest('[data-pui-inputotp-slot]');
    if(!slot) return;
    e.preventDefault();
    var container = slot.closest('[data-pui-inputotp]');
    if(!container) return;
    var pasted = (e.clipboardData || window.clipboardData).getData('text');
    var chars = pasted.replace(/\s/g, '').split('');
    var slots = getSlots(container);
    var startIndex = slots.indexOf(slot);
    for(var i=0;i<chars.length && startIndex + i < slots.length;i++){
      slots[startIndex + i].value = chars[i];
    }
    updateHiddenValue(container);
    var nextEmpty = findFirstEmptySlot(container);
    focusSlot(nextEmpty || slots[Math.min(startIndex + chars.length, slots.length - 1)]);
  });
  document.addEventListener('click', function(e){
    if(!e.target.matches('label[for]')) return;
    var targetId = e.target.getAttribute('for');
    var hiddenInput = document.getElementById(targetId);
    if(!hiddenInput || !hiddenInput.matches('[data-pui-inputotp-value-target]')) return;
    e.preventDefault();
    var container = hiddenInput.closest('[data-pui-inputotp]');
    if(!container) return;
    var slots = getSlots(container);
    if(slots.length > 0) focusSlot(slots[0]);
  });
  document.addEventListener('reset', function(e){
    if(!e.target.matches('form')) return;
    e.target.querySelectorAll('[data-pui-inputotp]').forEach(function(container){
      getSlots(container).forEach(function(slot){ slot.value = ''; });
      updateHiddenValue(container);
    });
  });
  new MutationObserver(function(){
    document.querySelectorAll('[data-pui-inputotp]').forEach(function(container){
      var slots = getSlots(container);
      if(!slots.length) return;
      var initialValue = container.getAttribute('data-pui-inputotp-value');
      if(initialValue && !slots[0].value){
        for(var i=0;i<slots.length && i < initialValue.length;i++){
          if(!slots[i].value) slots[i].value = initialValue[i];
        }
        updateHiddenValue(container);
      }
      if(container.hasAttribute('autofocus') && !slots.some(function(s){ return s === document.activeElement; })){
        requestAnimationFrame(function(){
          if(slots[0] && !slots.some(function(s){ return s === document.activeElement; })){
            focusSlot(slots[0]);
          }
        });
      }
    });
  }).observe(document.body, { childList: true, subtree: true });
})();`
