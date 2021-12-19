function logOut() {
  Cookies.remove('auth-session');
}

window.onload = function () {
  var logoutBtn = document.getElementById("btn-logout");
  if (logoutBtn) {
    logoutBtn.addEventListener('click', logOut);
  }
}

window.onunload = function() {
  var logoutBtn = document.getElementById("btn-logout");
  if (logoutBtn) {
    logoutBtn.removeEventListener('click', logOut);
  }
}