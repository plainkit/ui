(function () {
  "use strict";

  let openDialogs = new Set();

  // Open dialog
  function openDialog(dialogId) {
    // Find backdrop and content by instance ID
    const backdrop = document.querySelector(
      '[data-pui-dialog-backdrop][data-dialog-instance="' + dialogId + '"]',
    );
    const content = document.querySelector(
      '[data-pui-dialog-content][data-dialog-instance="' + dialogId + '"]',
    );

    if (!backdrop || !content) return;

    openDialogs.add(dialogId);

    // First, remove hidden state to make visible (but still in closed position)
    backdrop.removeAttribute("data-pui-dialog-hidden");
    content.removeAttribute("data-pui-dialog-hidden");

    // Then trigger the open animation after a frame
    requestAnimationFrame(() => {
      backdrop.setAttribute("data-pui-dialog-open", "true");
      content.setAttribute("data-pui-dialog-open", "true");
      document.body.style.overflow = "hidden";

      // Update triggers
      document
        .querySelectorAll(
          '[data-pui-dialog-trigger][data-dialog-instance="' + dialogId + '"]',
        )
        .forEach((trigger) => {
          trigger.setAttribute("data-pui-dialog-trigger-open", "true");
        });

      // Focus first focusable element
      setTimeout(() => {
        const focusable = content.querySelector(
          'button, [href], input, select, textarea, [tabindex]:not([tabindex="-1"])',
        );
        focusable?.focus();
      }, 50);
    });
  }

  // Close dialog
  function closeDialog(dialogId) {
    // Find backdrop and content by instance ID
    const backdrop = document.querySelector(
      '[data-pui-dialog-backdrop][data-dialog-instance="' + dialogId + '"]',
    );
    const content = document.querySelector(
      '[data-pui-dialog-content][data-dialog-instance="' + dialogId + '"]',
    );

    if (!backdrop || !content) return;

    // Start close animation
    backdrop.setAttribute("data-pui-dialog-open", "false");
    content.setAttribute("data-pui-dialog-open", "false");

    // Update triggers
    document
      .querySelectorAll(
        '[data-pui-dialog-trigger][data-dialog-instance="' + dialogId + '"]',
      )
      .forEach((trigger) => {
        trigger.setAttribute("data-pui-dialog-trigger-open", "false");
      });

    // Wait for animation to complete before hiding
    setTimeout(() => {
      backdrop.setAttribute("data-pui-dialog-hidden", "true");
      content.setAttribute("data-pui-dialog-hidden", "true");
    }, 300);

    openDialogs.delete(dialogId);

    // Restore body overflow if no dialogs are open
    if (openDialogs.size === 0) {
      document.body.style.overflow = "";
    }
  }

  // Get dialog instance from element
  function getDialogInstance(element) {
    // Try to get from data attribute
    const instance = element.getAttribute("data-dialog-instance");
    if (instance) return instance;

    // Try to find parent dialog
    const parentDialog = element.closest("[data-pui-dialog]");
    if (parentDialog) {
      return parentDialog.getAttribute("data-dialog-instance");
    }

    return null;
  }

  // Helper function for checking dialog state
  function isDialogOpen(dialogId) {
    const content = document.querySelector(
      '[data-pui-dialog-content][data-dialog-instance="' + dialogId + '"]',
    );
    return content?.getAttribute("data-pui-dialog-open") === "true" || false;
  }

  // Helper function for toggling dialog
  function toggleDialog(dialogId) {
    isDialogOpen(dialogId) ? closeDialog(dialogId) : openDialog(dialogId);
  }

  // Event delegation
  document.addEventListener("click", (e) => {
    // Handle trigger clicks
    const trigger = e.target.closest("[data-pui-dialog-trigger]");
    if (
      trigger &&
      !trigger.hasAttribute("disabled") &&
      !trigger.classList.contains("opacity-50")
    ) {
      const dialogId = trigger.getAttribute("data-dialog-instance");
      if (!dialogId) return;

      toggleDialog(dialogId);
      return;
    }

    // Handle close button clicks
    const closeBtn = e.target.closest("[data-pui-dialog-close]");
    if (closeBtn) {
      // First check if the close button has a For value (dialog ID specified)
      const forValue = closeBtn.getAttribute("data-pui-dialog-close");
      const dialogId = forValue || getDialogInstance(closeBtn);

      if (dialogId) {
        closeDialog(dialogId);
      }
      return;
    }

    // Handle click away - close when clicking on backdrop
    const backdrop = e.target.closest("[data-pui-dialog-backdrop]");
    if (backdrop) {
      const dialogId = backdrop.getAttribute("data-dialog-instance");
      if (!dialogId) return;

      // Check if click away is disabled
      const wrapper = document.querySelector(
        '[data-pui-dialog][data-dialog-instance="' + dialogId + '"]',
      );
      const content = document.querySelector(
        '[data-pui-dialog-content][data-dialog-instance="' + dialogId + '"]',
      );

      const isDisabled =
        wrapper?.hasAttribute("data-pui-dialog-disable-click-away") ||
        content?.hasAttribute("data-pui-dialog-disable-click-away");

      if (!isDisabled) {
        closeDialog(dialogId);
      }
    }
  });

  // ESC key handler
  document.addEventListener("keydown", (e) => {
    if (e.key === "Escape" && openDialogs.size > 0) {
      // Close the most recently opened dialog
      const dialogId = Array.from(openDialogs).pop();

      // Check if ESC is disabled
      const wrapper = document.querySelector(
        '[data-pui-dialog][data-dialog-instance="' + dialogId + '"]',
      );
      const content = document.querySelector(
        '[data-pui-dialog-content][data-dialog-instance="' + dialogId + '"]',
      );

      const isDisabled =
        wrapper?.hasAttribute("data-pui-dialog-disable-esc") ||
        content?.hasAttribute("data-pui-dialog-disable-esc");

      if (!isDisabled) {
        closeDialog(dialogId);
      }
    }
  });

  // Initialize dialogs that should be open on load
  document.addEventListener("DOMContentLoaded", () => {
    // Find all dialogs that should be initially open
    document
      .querySelectorAll(
        '[data-pui-dialog-content][data-pui-dialog-open="true"]',
      )
      .forEach((content) => {
        const dialogId = content.getAttribute("data-dialog-instance");
        if (dialogId) {
          openDialogs.add(dialogId);
          document.body.style.overflow = "hidden";
        }
      });
  });

  // Expose public API
  window.tui = window.tui || {};
  window.tui.dialog = {
    open: openDialog,
    close: closeDialog,
    toggle: toggleDialog,
    isOpen: isDialogOpen,
  };
})();
