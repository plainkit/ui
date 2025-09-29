(function () {
  if (typeof document === "undefined") return;
  function updateValue(input) {
    var wrapper = input.closest("[data-pui-slider-wrapper]");
    if (!wrapper) return;
    var valueEls = wrapper.querySelectorAll(
      '[data-pui-slider-value][data-pui-slider-value-for="' + input.id + '"]',
    );
    valueEls.forEach(function (el) {
      el.textContent = input.value;
    });
  }
  document.addEventListener("input", function (event) {
    var target = event.target;
    if (target && target.matches("[data-pui-slider-input]")) {
      updateValue(target);
    }
  });
  document.querySelectorAll("[data-pui-slider-input]").forEach(updateValue);
})();
