export default function authGuard(to, from, next) {
  const isAuthenticated = sessionStorage.getItem("token");

  if (to.meta.requiresAuth && !isAuthenticated) {
    next("/login");
  } else if (to.meta.requiresGuest && isAuthenticated) {
    next("/dashboard");
  } else {
    next();
  }
}
