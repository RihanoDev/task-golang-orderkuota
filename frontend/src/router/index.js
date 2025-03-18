import { createRouter, createWebHistory } from "vue-router";
import Login from "../views/LoginView.vue";
import Register from "../views/RegisterView.vue";
import Dashboard from "../views/DashboardView.vue";
import authGuard from "./authGuard";

const routes = [
  { path: "/", redirect: "/login" },
  { path: "/login", name: "Login", component: Login, meta: { requiresGuest: true } },
  { path: "/register", name: "Register", component: Register, meta: { requiresGuest: true } },
  { path: "/dashboard", name: "Dashboard", component: Dashboard, meta: { requiresAuth: true } },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach(authGuard);

export default router;
