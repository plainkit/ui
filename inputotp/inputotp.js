(function () {
  "use strict";
  function getSlots(container) {
    return Array.from(
      container.querySelectorAll("[data-pui-inputotp-slot]"),
    ).sort(function (a, b) {
      return (
        parseInt(a.getAttribute("data-pui-inputotp-index")) -
        parseInt(b.getAttribute("data-pui-inputotp-index"))
      );
    });
  }
  function focusSlot(slot) {
    if (!slot) return;
    slot.focus();
    setTimeout(function () {
      slot.select();
    }, 0);
  }
  function updateHiddenValue(container) {
    var hiddenInput = container.querySelector(
      "[data-pui-inputotp-value-target]",
    );
    var slots = getSlots(container);
    if (hiddenInput && slots.length) {
      hiddenInput.value = slots
        .map(function (s) {
          return s.value;
        })
        .join("");
    }
  }
  function findFirstEmptySlot(container) {
    var slots = getSlots(container);
    for (var i = 0; i < slots.length; i++) {
      if (!slots[i].value) return slots[i];
    }
    return null;
  }
  function getNextSlot(container, currentSlot) {
    var slots = getSlots(container);
    var index = slots.indexOf(currentSlot);
    return index >= 0 && index < slots.length - 1 ? slots[index + 1] : null;
  }
  function getPrevSlot(container, currentSlot) {
    var slots = getSlots(container);
    var index = slots.indexOf(currentSlot);
    return index > 0 ? slots[index - 1] : null;
  }
  document.addEventListener("input", function (e) {
    if (!e.target.matches("[data-pui-inputotp-slot]")) return;
    var slot = e.target;
    var container = slot.closest("[data-pui-inputotp]");
    if (!container) return;
    if (slot.value === " ") {
      slot.value = "";
      return;
    }
    if (slot.value.length > 1) {
      slot.value = slot.value.slice(-1);
    }
    if (slot.value) {
      var nextSlot = getNextSlot(container, slot);
      if (nextSlot) focusSlot(nextSlot);
    }
    updateHiddenValue(container);
  });
  document.addEventListener("keydown", function (e) {
    if (!e.target.matches("[data-pui-inputotp-slot]")) return;
    var slot = e.target;
    var container = slot.closest("[data-pui-inputotp]");
    if (!container) return;
    if (e.key === "Backspace") {
      e.preventDefault();
      if (slot.value) {
        slot.value = "";
        updateHiddenValue(container);
      } else {
        var prevSlot = getPrevSlot(container, slot);
        if (prevSlot) {
          prevSlot.value = "";
          updateHiddenValue(container);
          focusSlot(prevSlot);
        }
      }
    } else if (e.key === "ArrowLeft") {
      e.preventDefault();
      var prev = getPrevSlot(container, slot);
      if (prev) focusSlot(prev);
    } else if (e.key === "ArrowRight") {
      e.preventDefault();
      var next = getNextSlot(container, slot);
      if (next) focusSlot(next);
    }
  });
  document.addEventListener(
    "focus",
    function (e) {
      if (!e.target.matches("[data-pui-inputotp-slot]")) return;
      var slot = e.target;
      var container = slot.closest("[data-pui-inputotp]");
      if (!container) return;
      var firstEmpty = findFirstEmptySlot(container);
      if (firstEmpty && firstEmpty !== slot) {
        focusSlot(firstEmpty);
        return;
      }
      setTimeout(function () {
        slot.select();
      }, 0);
    },
    true,
  );
  document.addEventListener("paste", function (e) {
    var slot = e.target.closest("[data-pui-inputotp-slot]");
    if (!slot) return;
    e.preventDefault();
    var container = slot.closest("[data-pui-inputotp]");
    if (!container) return;
    var pasted = (e.clipboardData || window.clipboardData).getData("text");
    var chars = pasted.replace(/\s/g, "").split("");
    var slots = getSlots(container);
    var startIndex = slots.indexOf(slot);
    for (var i = 0; i < chars.length && startIndex + i < slots.length; i++) {
      slots[startIndex + i].value = chars[i];
    }
    updateHiddenValue(container);
    var nextEmpty = findFirstEmptySlot(container);
    focusSlot(
      nextEmpty || slots[Math.min(startIndex + chars.length, slots.length - 1)],
    );
  });
  document.addEventListener("click", function (e) {
    if (!e.target.matches("label[for]")) return;
    var targetId = e.target.getAttribute("for");
    var hiddenInput = document.getElementById(targetId);
    if (
      !hiddenInput ||
      !hiddenInput.matches("[data-pui-inputotp-value-target]")
    )
      return;
    e.preventDefault();
    var container = hiddenInput.closest("[data-pui-inputotp]");
    if (!container) return;
    var slots = getSlots(container);
    if (slots.length > 0) focusSlot(slots[0]);
  });
  document.addEventListener("reset", function (e) {
    if (!e.target.matches("form")) return;
    e.target
      .querySelectorAll("[data-pui-inputotp]")
      .forEach(function (container) {
        getSlots(container).forEach(function (slot) {
          slot.value = "";
        });
        updateHiddenValue(container);
      });
  });
  new MutationObserver(function () {
    document
      .querySelectorAll("[data-pui-inputotp]")
      .forEach(function (container) {
        var slots = getSlots(container);
        if (!slots.length) return;
        var initialValue = container.getAttribute("data-pui-inputotp-value");
        if (initialValue && !slots[0].value) {
          for (var i = 0; i < slots.length && i < initialValue.length; i++) {
            if (!slots[i].value) slots[i].value = initialValue[i];
          }
          updateHiddenValue(container);
        }
        if (
          container.hasAttribute("autofocus") &&
          !slots.some(function (s) {
            return s === document.activeElement;
          })
        ) {
          requestAnimationFrame(function () {
            if (
              slots[0] &&
              !slots.some(function (s) {
                return s === document.activeElement;
              })
            ) {
              focusSlot(slots[0]);
            }
          });
        }
      });
  }).observe(document.body, { childList: true, subtree: true });
})();
