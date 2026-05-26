// Middleware guest: redirige al home si ya está autenticado.
// Usado en login y registro para evitar que un usuario logueado vea esas páginas.
export default defineNuxtRouteMiddleware(() => {
  if (import.meta.client) {
    const token = localStorage.getItem("deprodeca_token")
    if (token) {
      return navigateTo("/")
    }
  }
})
