// Middleware de admin.
// Redirige a / si el usuario no tiene rol "admin".
export default defineNuxtRouteMiddleware(() => {
  if (import.meta.client) {
    const userStr = localStorage.getItem("deprodeca_user")
    if (userStr) {
      const user = JSON.parse(userStr)
      if (user.rol === "admin") return
    }
    return navigateTo("/")
  }
})
