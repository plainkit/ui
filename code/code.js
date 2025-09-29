(function () {
  "use strict";

  // Load highlight.js if not already loaded
  if (typeof hljs === "undefined") {
    const script = document.createElement("script");
    script.src =
      "https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/highlight.min.js";
    script.onload = function () {
      initializeCode();
    };
    document.head.appendChild(script);
  } else {
    initializeCode();
  }

  function initializeCode() {
    // Highlight all code blocks
    document
      .querySelectorAll("[data-pui-code-block]")
      .forEach(function (block) {
        if (!block.dataset.highlighted) {
          hljs.highlightElement(block);
          block.dataset.highlighted = "true";
        }
      });

    // Setup copy functionality
    document
      .querySelectorAll("[data-pui-code-copy-button]")
      .forEach(function (button) {
        if (button.dataset.initialized) return;
        button.dataset.initialized = "true";

        button.addEventListener("click", function () {
          const codeComponent = button.closest("[data-pui-code-component]");
          const codeBlock = codeComponent.querySelector(
            "[data-pui-code-block]",
          );
          const checkIcon = button.querySelector("[data-pui-code-icon-check]");
          const clipboardIcon = button.querySelector(
            "[data-pui-code-icon-clipboard]",
          );

          if (codeBlock) {
            navigator.clipboard
              .writeText(codeBlock.textContent)
              .then(function () {
                // Show check icon
                clipboardIcon.style.display = "none";
                checkIcon.style.display = "inline";

                // Reset after 2 seconds
                setTimeout(function () {
                  clipboardIcon.style.display = "inline";
                  checkIcon.style.display = "none";
                }, 2000);
              })
              .catch(function (err) {
                console.error("Failed to copy code: ", err);
              });
          }
        });
      });
  }

  // Initialize on DOM content loaded
  if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", initializeCode);
  } else {
    initializeCode();
  }
})();
