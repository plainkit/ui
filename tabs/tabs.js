(function () {
  if (typeof document === "undefined") return;
  function activate(tabsId, value) {
    var triggers = document.querySelectorAll(
      '[data-pui-tabs-trigger][data-pui-tabs-id="' + tabsId + '"]',
    );
    var contents = document.querySelectorAll(
      '[data-pui-tabs-content][data-pui-tabs-id="' + tabsId + '"]',
    );
    triggers.forEach(function (trigger) {
      var isActive = trigger.getAttribute("data-pui-tabs-value") === value;
      trigger.setAttribute(
        "data-pui-tabs-state",
        isActive ? "active" : "inactive",
      );
    });
    contents.forEach(function (panel) {
      var isActive = panel.getAttribute("data-pui-tabs-value") === value;
      panel.setAttribute(
        "data-pui-tabs-state",
        isActive ? "active" : "inactive",
      );
      panel.classList.toggle("hidden", !isActive);
    });
  }
  document.addEventListener("click", function (event) {
    var target = event.target;
    if (target && target.nodeType === 3) target = target.parentElement;
    if (!(target instanceof Element)) return;
    var trigger = target.closest("[data-pui-tabs-trigger]");
    if (!trigger) return;
    var tabsId = trigger.getAttribute("data-pui-tabs-id");
    var value = trigger.getAttribute("data-pui-tabs-value");
    if (!tabsId || !value) return;
    activate(tabsId, value);
  });
})();
