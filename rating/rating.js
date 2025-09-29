(function () {
  "use strict";
  function getConfig(el) {
    return {
      value: parseFloat(el.getAttribute("data-pui-rating-initial-value")) || 0,
      precision: parseFloat(el.getAttribute("data-pui-rating-precision")) || 1,
      readonly: el.getAttribute("data-pui-rating-readonly") === "true",
      name: el.getAttribute("data-pui-rating-name") || "",
      onlyInteger: el.getAttribute("data-pui-rating-onlyinteger") === "true",
    };
  }
  function getCurrentValue(el) {
    var stored = parseFloat(el.getAttribute("data-pui-rating-current"));
    if (!isNaN(stored)) return stored;
    return parseFloat(el.getAttribute("data-pui-rating-initial-value")) || 0;
  }
  function setCurrentValue(el, value) {
    el.setAttribute("data-pui-rating-current", value);
    var hidden = el.querySelector("[data-pui-rating-input]");
    if (hidden) {
      hidden.value = value.toFixed(2);
      hidden.dispatchEvent(new Event("input", { bubbles: true }));
      hidden.dispatchEvent(new Event("change", { bubbles: true }));
    }
  }
  function updateItemStyles(el, displayValue) {
    var currentValue = getCurrentValue(el);
    var valueToCompare = displayValue > 0 ? displayValue : currentValue;
    el.querySelectorAll("[data-pui-rating-item]").forEach(function (item) {
      var itemValue = parseInt(item.getAttribute("data-pui-rating-value"), 10);
      if (isNaN(itemValue)) return;
      var foreground = item.querySelector("[data-pui-rating-item-foreground]");
      if (!foreground) return;
      var filled = itemValue <= Math.floor(valueToCompare);
      var partial =
        !filled && itemValue - 1 < valueToCompare && valueToCompare < itemValue;
      var percentage = partial
        ? (valueToCompare - Math.floor(valueToCompare)) * 100
        : 0;
      foreground.style.width = filled
        ? "100%"
        : partial
          ? percentage + "%"
          : "0%";
    });
  }
  function getMaxValue(el) {
    var max = 0;
    el.querySelectorAll("[data-pui-rating-item]").forEach(function (item) {
      var value = parseInt(item.getAttribute("data-pui-rating-value"), 10);
      if (!isNaN(value) && value > max) max = value;
    });
    return Math.max(1, max);
  }
  document.addEventListener("click", function (e) {
    var item = e.target.closest("[data-pui-rating-item]");
    if (!item) return;
    var el = item.closest("[data-pui-rating-component]");
    if (!el) return;
    var config = getConfig(el);
    if (config.readonly) return;
    var itemValue = parseInt(item.getAttribute("data-pui-rating-value"), 10);
    if (isNaN(itemValue)) return;
    var currentValue = getCurrentValue(el);
    var maxValue = getMaxValue(el);
    var newValue = itemValue;
    if (config.onlyInteger) {
      newValue = Math.round(newValue);
    } else {
      if (currentValue === newValue && newValue % 1 === 0) {
        newValue = Math.max(0, newValue - config.precision);
      } else {
        newValue = Math.round(newValue / config.precision) * config.precision;
      }
    }
    newValue = Math.max(0, Math.min(maxValue, newValue));
    setCurrentValue(el, newValue);
    updateItemStyles(el, 0);
    el.dispatchEvent(
      new CustomEvent("rating-change", {
        bubbles: true,
        detail: { name: config.name, value: newValue, maxValue: maxValue },
      }),
    );
  });
  document.addEventListener("mouseover", function (e) {
    var item = e.target.closest("[data-pui-rating-item]");
    if (!item) return;
    var el = item.closest("[data-pui-rating-component]");
    if (!el || getConfig(el).readonly) return;
    var previewValue = parseInt(item.getAttribute("data-pui-rating-value"), 10);
    if (!isNaN(previewValue)) updateItemStyles(el, previewValue);
  });
  document.addEventListener("mouseout", function (e) {
    var el = e.target.closest("[data-pui-rating-component]");
    if (!el || getConfig(el).readonly) return;
    if (!el.contains(e.relatedTarget)) updateItemStyles(el, 0);
  });
  document.addEventListener("reset", function (e) {
    if (!e.target.matches("form")) return;
    e.target
      .querySelectorAll("[data-pui-rating-component]")
      .forEach(function (el) {
        var config = getConfig(el);
        setCurrentValue(el, config.value);
        updateItemStyles(el, 0);
      });
  });
  new MutationObserver(function () {
    document
      .querySelectorAll("[data-pui-rating-component]")
      .forEach(function (el) {
        if (!el.hasAttribute("data-pui-rating-current")) {
          var config = getConfig(el);
          var maxValue = getMaxValue(el);
          var value = Math.max(0, Math.min(maxValue, config.value));
          var rounded = Math.round(value / config.precision) * config.precision;
          setCurrentValue(el, isFinite(rounded) ? rounded : 0);
        }
        updateItemStyles(el, 0);
        if (getConfig(el).readonly) {
          el.style.cursor = "default";
          el.querySelectorAll("[data-pui-rating-item]").forEach(
            function (item) {
              item.style.cursor = "default";
            },
          );
        }
      });
  }).observe(document.body, { childList: true, subtree: true });
})();
