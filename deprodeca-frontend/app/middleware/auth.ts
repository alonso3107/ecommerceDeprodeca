// Middleware de autenticación global.
// Redirige a /auth/login si la ruta requiere auth y no hay token.
export default defineNuxtRouteMiddleware((to) => {
  // Solo aplica si la página tiene middleware: "auth"
  if (to.meta.middleware !== "auth") return

  if (import.meta.client) {
    const token = localStorage.getItem("deprodeca_token")
    if (!token) {
      return navigateTo("/auth/login")
    }
  }
})
