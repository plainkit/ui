package rating

import (
	"fmt"
	"strconv"

	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/internal/classnames"
)

type Style string

type Props struct {
	ID          string
	Class       string
	Attrs       []html.Global
	Value       float64
	ReadOnly    bool
	Precision   float64
	Name        string
	Form        string
	OnlyInteger bool
}

type GroupProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type ItemProps struct {
	ID    string
	Class string
	Attrs []html.Global
	Value int
	Style Style
}

const (
	StyleStar  Style = "star"
	StyleHeart Style = "heart"
	StyleEmoji Style = "emoji"
)

// Rating wraps interactive rating items and manages hidden input value syncing.
func Rating(props Props, args ...html.DivArg) html.Node {
	precision := props.Precision
	if precision <= 0 {
		precision = 1
	}

	dataset := []html.Global{
		html.AData("pui-rating-component", ""),
		html.AData("pui-rating-initial-value", fmt.Sprintf("%.2f", props.Value)),
		html.AData("pui-rating-precision", fmt.Sprintf("%.2f", precision)),
		html.AData("pui-rating-readonly", strconv.FormatBool(props.ReadOnly)),
		html.AData("pui-rating-onlyinteger", strconv.FormatBool(props.OnlyInteger)),
	}
	if props.Name != "" {
		dataset = append(dataset, html.AData("pui-rating-name", props.Name))
	}

	divArgs := []html.DivArg{html.AClass(classnames.Merge("flex flex-col items-start gap-1", props.Class))}
	if props.ID != "" {
		divArgs = append(divArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}
	for _, data := range dataset {
		divArgs = append(divArgs, data)
	}

	if props.Name != "" {
		hiddenArgs := []html.InputArg{
			html.AType("hidden"),
			html.AName(props.Name),
			html.AValue(fmt.Sprintf("%.2f", props.Value)),
			html.AData("pui-rating-input", ""),
		}
		if props.Form != "" {
			hiddenArgs = append(hiddenArgs, html.AForm(props.Form))
		}
		hidden := html.Input(hiddenArgs...)
		divArgs = append(divArgs, hidden)
	}

	divArgs = append(divArgs, args...)

	node := html.Div(divArgs...)
	return node.WithAssets("", ratingJS, "ui-rating")
}

// Group arranges rating items in a single row.
func Group(props GroupProps, args ...html.DivArg) html.Node {
	groupArgs := []html.DivArg{html.AClass(classnames.Merge("flex flex-row items-center gap-1", props.Class))}
	if props.ID != "" {
		groupArgs = append(groupArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		groupArgs = append(groupArgs, attr)
	}
	groupArgs = append(groupArgs, args...)
	return html.Div(groupArgs...)
}

// Item renders an individual rating icon with layered fill for partial states.
func Item(props ItemProps) html.Node {
	style := props.Style
	if style == "" {
		style = StyleStar
	}

	itemArgs := []html.DivArg{
		html.AClass(classnames.Merge("relative transition-opacity cursor-pointer", colorClass(style), props.Class)),
		html.AData("pui-rating-item", ""),
		html.AData("pui-rating-value", strconv.Itoa(props.Value)),
	}
	if props.ID != "" {
		itemArgs = append(itemArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		itemArgs = append(itemArgs, attr)
	}

	background := html.Div(
		html.AClass("opacity-30"),
		ratingIcon(style, false, float64(props.Value)),
	)

	foreground := html.Div(
		html.AClass("absolute inset-0 overflow-hidden w-0"),
		html.AData("pui-rating-item-foreground", ""),
		ratingIcon(style, true, float64(props.Value)),
	)

	itemArgs = append(itemArgs, background, foreground)
	return html.Div(itemArgs...)
}

func colorClass(style Style) string {
	switch style {
	case StyleHeart:
		return "text-destructive"
	case StyleEmoji:
		return "text-yellow-500"
	default:
		return "text-yellow-400"
	}
}

func ratingIcon(style Style, filled bool, value float64) html.Node {
	if style == StyleEmoji {
		if filled {
			switch {
			case value <= 1:
				return lucide.Angry()
			case value <= 2:
				return lucide.Frown()
			case value <= 3:
				return lucide.Meh()
			case value <= 4:
				return lucide.Smile()
			default:
				return lucide.Laugh()
			}
		}
		return lucide.Meh()
	}

	iconArgs := []html.SvgArg{}
	if filled {
		iconArgs = append(iconArgs, html.AFill("currentColor"), html.AStroke("none"))
	}

	switch style {
	case StyleHeart:
		return lucide.Heart(iconArgs...)
	default:
		return lucide.Star(iconArgs...)
	}
}

const ratingJS = `(function(){
  'use strict';
  function getConfig(el){
    return {
      value: parseFloat(el.getAttribute('data-pui-rating-initial-value')) || 0,
      precision: parseFloat(el.getAttribute('data-pui-rating-precision')) || 1,
      readonly: el.getAttribute('data-pui-rating-readonly') === 'true',
      name: el.getAttribute('data-pui-rating-name') || '',
      onlyInteger: el.getAttribute('data-pui-rating-onlyinteger') === 'true'
    };
  }
  function getCurrentValue(el){
    var stored = parseFloat(el.getAttribute('data-pui-rating-current'));
    if(!isNaN(stored)) return stored;
    return parseFloat(el.getAttribute('data-pui-rating-initial-value')) || 0;
  }
  function setCurrentValue(el, value){
    el.setAttribute('data-pui-rating-current', value);
    var hidden = el.querySelector('[data-pui-rating-input]');
    if(hidden){
      hidden.value = value.toFixed(2);
      hidden.dispatchEvent(new Event('input', { bubbles: true }));
      hidden.dispatchEvent(new Event('change', { bubbles: true }));
    }
  }
  function updateItemStyles(el, displayValue){
    var currentValue = getCurrentValue(el);
    var valueToCompare = displayValue > 0 ? displayValue : currentValue;
    el.querySelectorAll('[data-pui-rating-item]').forEach(function(item){
      var itemValue = parseInt(item.getAttribute('data-pui-rating-value'), 10);
      if(isNaN(itemValue)) return;
      var foreground = item.querySelector('[data-pui-rating-item-foreground]');
      if(!foreground) return;
      var filled = itemValue <= Math.floor(valueToCompare);
      var partial = !filled && itemValue - 1 < valueToCompare && valueToCompare < itemValue;
      var percentage = partial ? (valueToCompare - Math.floor(valueToCompare)) * 100 : 0;
      foreground.style.width = filled ? '100%' : (partial ? percentage + '%' : '0%');
    });
  }
  function getMaxValue(el){
    var max = 0;
    el.querySelectorAll('[data-pui-rating-item]').forEach(function(item){
      var value = parseInt(item.getAttribute('data-pui-rating-value'), 10);
      if(!isNaN(value) && value > max) max = value;
    });
    return Math.max(1, max);
  }
  document.addEventListener('click', function(e){
    var item = e.target.closest('[data-pui-rating-item]');
    if(!item) return;
    var el = item.closest('[data-pui-rating-component]');
    if(!el) return;
    var config = getConfig(el);
    if(config.readonly) return;
    var itemValue = parseInt(item.getAttribute('data-pui-rating-value'), 10);
    if(isNaN(itemValue)) return;
    var currentValue = getCurrentValue(el);
    var maxValue = getMaxValue(el);
    var newValue = itemValue;
    if(config.onlyInteger){
      newValue = Math.round(newValue);
    } else {
      if(currentValue === newValue && newValue % 1 === 0){
        newValue = Math.max(0, newValue - config.precision);
      } else {
        newValue = Math.round(newValue / config.precision) * config.precision;
      }
    }
    newValue = Math.max(0, Math.min(maxValue, newValue));
    setCurrentValue(el, newValue);
    updateItemStyles(el, 0);
    el.dispatchEvent(new CustomEvent('rating-change', {
      bubbles: true,
      detail: { name: config.name, value: newValue, maxValue: maxValue }
    }));
  });
  document.addEventListener('mouseover', function(e){
    var item = e.target.closest('[data-pui-rating-item]');
    if(!item) return;
    var el = item.closest('[data-pui-rating-component]');
    if(!el || getConfig(el).readonly) return;
    var previewValue = parseInt(item.getAttribute('data-pui-rating-value'), 10);
    if(!isNaN(previewValue)) updateItemStyles(el, previewValue);
  });
  document.addEventListener('mouseout', function(e){
    var el = e.target.closest('[data-pui-rating-component]');
    if(!el || getConfig(el).readonly) return;
    if(!el.contains(e.relatedTarget)) updateItemStyles(el, 0);
  });
  document.addEventListener('reset', function(e){
    if(!e.target.matches('form')) return;
    e.target.querySelectorAll('[data-pui-rating-component]').forEach(function(el){
      var config = getConfig(el);
      setCurrentValue(el, config.value);
      updateItemStyles(el, 0);
    });
  });
  new MutationObserver(function(){
    document.querySelectorAll('[data-pui-rating-component]').forEach(function(el){
      if(!el.hasAttribute('data-pui-rating-current')){
        var config = getConfig(el);
        var maxValue = getMaxValue(el);
        var value = Math.max(0, Math.min(maxValue, config.value));
        var rounded = Math.round(value / config.precision) * config.precision;
        setCurrentValue(el, isFinite(rounded) ? rounded : 0);
      }
      updateItemStyles(el, 0);
      if(getConfig(el).readonly){
        el.style.cursor = 'default';
        el.querySelectorAll('[data-pui-rating-item]').forEach(function(item){
          item.style.cursor = 'default';
        });
      }
    });
  }).observe(document.body, { childList: true, subtree: true });
})();`
