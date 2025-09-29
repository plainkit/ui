(function () {
  if (typeof document === "undefined") return;
  document.addEventListener("click", function (event) {
    var btn = event.target.closest("[data-pui-input-toggle-password]");
    if (!btn) return;
    var inputId = btn.getAttribute("data-pui-input-toggle-password");
    if (!inputId) return;
    var input = document.getElementById(inputId);
    if (!input) return;
    var isPassword = input.getAttribute("type") === "password";
    input.setAttribute("type", isPassword ? "text" : "password");
    var openIcon = btn.querySelector(".icon-open");
    var closedIcon = btn.querySelector(".icon-closed");
    if (openIcon) {
      openIcon.classList.toggle("hidden", !isPassword);
      openIcon.classList.toggle("block", isPassword);
    }
    if (closedIcon) {
      closedIcon.classList.toggle("hidden", isPassword);
      closedIcon.classList.toggle("block", !isPassword);
    }
  });
})();
