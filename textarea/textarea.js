(function () {
  if (typeof document === "undefined") return;
  function resize(el) {
    el.style.height = "auto";
    el.style.height = el.scrollHeight + "px";
  }
  document.addEventListener("input", function (event) {
    var target = event.target;
    if (target && target.matches("[data-pui-textarea-auto-resize]")) {
      resize(target);
    }
  });
  document
    .querySelectorAll("[data-pui-textarea-auto-resize]")
    .forEach(function (el) {
      resize(el);
    });
})();
