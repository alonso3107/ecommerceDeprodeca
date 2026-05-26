// useToast — Composable para mostrar notificaciones toast.
// Usa la UI Store de Pinia para estado global.
import { useUIStore } from "~/middleware/stores/ui.store"

export function useToast() {
  const ui = useUIStore()

  function success(message: string) {
    ui.showToast(message, "success")
  }

  function error(message: string) {
    ui.showToast(message, "error")
  }

  function warning(message: string) {
    ui.showToast(message, "warning")
  }

  function info(message: string) {
    ui.showToast(message, "info")
  }

  return { success, error, warning, info, toast: ui }
}
