(function () {
  "use strict";
  var toastTimers = new Map();

  // Initialize toast containers if not present
  function ensureToastContainer(position) {
    position = position || "bottom-right";
    var containerId = "toast-container-" + position;
    var container = document.getElementById(containerId);
    if (!container) {
      container = document.createElement("div");
      container.id = containerId;
      container.dataset.tuiToastContainer = position;
      var posClasses = {
        "top-right": "top-0 right-0",
        "top-left": "top-0 left-0",
        "top-center": "top-0 left-1/2 -translate-x-1/2",
        "bottom-right": "bottom-0 right-0",
        "bottom-left": "bottom-0 left-0",
        "bottom-center": "bottom-0 left-1/2 -translate-x-1/2",
      };
      container.className =
        "fixed z-50 pointer-events-none p-4 " +
        (posClasses[position] || posClasses["bottom-right"]);
      document.body.appendChild(container);
    }
    return container;
  }

  // Create toast element from configuration
  function createToastElement(config) {
    var toast = document.createElement("div");
    toast.dataset.tuiToast = "";
    toast.dataset.variant = config.variant || "default";
    toast.dataset.position = config.position || "bottom-right";
    toast.dataset.tuiToastDuration = config.duration || "3000";
    toast.className =
      "pointer-events-auto p-4 w-full md:max-w-[420px] animate-in fade-in slide-in-from-bottom-4 duration-300 mb-4";

    var inner = document.createElement("div");
    inner.className =
      "w-full bg-popover text-popover-foreground rounded-lg shadow-xs border pt-5 pb-4 px-4 flex items-center justify-center relative overflow-hidden group gap-3";

    // Add progress indicator
    if (config.showIndicator === "true" && parseInt(config.duration) > 0) {
      var progressWrapper = document.createElement("div");
      progressWrapper.className =
        "absolute top-0 left-0 right-0 h-1 overflow-hidden";
      var progress = document.createElement("div");
      progress.className =
        "toast-progress h-full origin-left transition-transform ease-linear ";
      var variantColors = {
        default: "bg-gray-500",
        success: "bg-green-500",
        error: "bg-red-500",
        warning: "bg-yellow-500",
        info: "bg-blue-500",
      };
      progress.className +=
        variantColors[config.variant] || variantColors["default"];
      progress.dataset.variant = config.variant || "default";
      progress.dataset.duration = config.duration || "3000";
      progressWrapper.appendChild(progress);
      inner.appendChild(progressWrapper);
    }

    // Add icon
    if (config.icon === "true" && config.variant !== "default") {
      var iconWrapper = document.createElement("div");
      iconWrapper.className = "size-5 flex-shrink-0";
      var iconColors = {
        success: "text-green-500",
        error: "text-destructive",
        warning: "text-yellow-500",
        info: "text-blue-500",
      };
      iconWrapper.className += " " + (iconColors[config.variant] || "");
      // Simple icon representation (you'd need actual SVG icons here)
      iconWrapper.innerHTML = getIconSvg(config.variant);
      inner.appendChild(iconWrapper);
    }

    // Add text content
    var textContainer = document.createElement("span");
    textContainer.className = "flex-1 min-w-0";
    if (config.title) {
      var title = document.createElement("p");
      title.className = "text-sm font-semibold truncate";
      title.textContent = config.title;
      textContainer.appendChild(title);
    }
    if (config.description) {
      var desc = document.createElement("p");
      desc.className = "text-sm opacity-90 mt-1";
      desc.textContent = config.description;
      textContainer.appendChild(desc);
    }
    inner.appendChild(textContainer);

    // Add dismiss button
    if (config.dismissible === "true") {
      var btn = document.createElement("button");
      btn.type = "button";
      btn.className =
        "inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 hover:bg-accent hover:text-accent-foreground size-8 p-0";
      btn.setAttribute("aria-label", "Close");
      btn.dataset.tuiToastDismiss = "";
      btn.innerHTML =
        '<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="opacity-75 hover:opacity-100"><path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg>';
      inner.appendChild(btn);
    }

    toast.appendChild(inner);
    return toast;
  }

  function getIconSvg(variant) {
    var icons = {
      success:
        '<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><path d="m9 12 2 2 4-4"/></svg>',
      error:
        '<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><path d="m15 9-6 6"/><path d="m9 9 6 6"/></svg>',
      warning:
        '<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m21.73 18-8-14a2 2 0 0 0-3.48 0l-8 14A2 2 0 0 0 4 21h16a2 2 0 0 0 1.73-3"/><path d="M12 9v4"/><path d="M12 17h.01"/></svg>',
      info: '<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg>',
    };
    return icons[variant] || "";
  }

  function spawnToast(config) {
    var container = ensureToastContainer(config.position);
    var toast = createToastElement(config);
    container.appendChild(toast);
    setupToast(toast);
  }

  function setupToast(toast) {
    var duration = parseInt(toast.dataset.tuiToastDuration || "3000", 10);
    var progress = toast.querySelector(".toast-progress");
    var state = {
      timer: null,
      startTime: Date.now(),
      remaining: duration,
      paused: false,
    };
    toastTimers.set(toast, state);
    if (progress && duration > 0) {
      progress.style.transitionDuration = duration + "ms";
      requestAnimationFrame(function () {
        progress.style.transform = "scaleX(0)";
      });
    }
    if (duration > 0) {
      state.timer = setTimeout(function () {
        dismissToast(toast);
      }, duration);
    }
    toast.addEventListener("mouseenter", function () {
      var st = toastTimers.get(toast);
      if (!st || st.paused) return;
      clearTimeout(st.timer);
      st.remaining = st.remaining - (Date.now() - st.startTime);
      st.paused = true;
      if (progress) {
        var computed = getComputedStyle(progress);
        progress.style.transitionDuration = "0ms";
        progress.style.transform = computed.transform;
      }
    });
    toast.addEventListener("mouseleave", function () {
      var st = toastTimers.get(toast);
      if (!st || !st.paused || st.remaining <= 0) return;
      st.startTime = Date.now();
      st.paused = false;
      st.timer = setTimeout(function () {
        dismissToast(toast);
      }, st.remaining);
      if (progress) {
        progress.style.transitionDuration = st.remaining + "ms";
        progress.style.transform = "scaleX(0)";
      }
    });
  }
  function dismissToast(toast) {
    toastTimers.delete(toast);
    toast.style.transition = "opacity 300ms, transform 300ms";
    toast.style.opacity = "0";
    toast.style.transform = "translateY(1rem)";
    setTimeout(function () {
      toast.remove();
    }, 300);
  }
  document.addEventListener("click", function (e) {
    var dismissBtn = e.target.closest("[data-pui-toast-dismiss]");
    if (dismissBtn) {
      var toast = dismissBtn.closest("[data-pui-toast]");
      if (toast) dismissToast(toast);
    }
  });
  // Handle toast trigger clicks
  document.addEventListener("click", function (e) {
    var trigger = e.target.closest("[data-pui-toast-trigger]");
    if (trigger) {
      e.preventDefault();
      var config = {
        title: trigger.dataset.toastTitle,
        description: trigger.dataset.toastDescription,
        variant: trigger.dataset.toastVariant || "default",
        position: trigger.dataset.toastPosition || "bottom-right",
        duration: trigger.dataset.toastDuration || "3000",
        dismissible: trigger.dataset.toastDismissible || "true",
        showIndicator: trigger.dataset.toastShowIndicator || "true",
        icon: trigger.dataset.toastIcon || "true",
      };
      spawnToast(config);
    }
  });

  // Setup any existing toasts on page load
  document.querySelectorAll("[data-pui-toast]").forEach(function (toast) {
    if (!toast.hasAttribute("data-pui-toast-template")) {
      setupToast(toast);
    }
  });
})();
