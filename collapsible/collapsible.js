(function () {
  if (typeof document === "undefined") return;
  function toggle(trigger) {
    var root = trigger.closest('[data-pui-collapsible="root"]');
    if (!root) return;
    var isOpen = root.getAttribute("data-pui-collapsible-state") === "open";
    var newState = isOpen ? "closed" : "open";
    root.setAttribute("data-pui-collapsible-state", newState);
    trigger.setAttribute("aria-expanded", (!isOpen).toString());
  }
  document.addEventListener("click", function (event) {
    var trigger = event.target.closest('[data-pui-collapsible="trigger"]');
    if (trigger) {
      event.preventDefault();
      toggle(trigger);
    }
  });
  document.addEventListener("keydown", function (event) {
    if (event.key !== " " && event.key !== "Enter") return;
    var trigger = event.target.closest('[data-pui-collapsible="trigger"]');
    if (trigger) {
      event.preventDefault();
      toggle(trigger);
    }
  });
})();
