// UI Store — estado global de la interfaz.
// Feature-Sliced Design: store compartida para toda la app.
import { defineStore } from "pinia"

export const useUIStore = defineStore("ui", () => {
  const mobileMenuOpen = ref(false)
  const toastMessage = ref("")
  const toastType = ref<"success" | "error" | "warning" | "info">("info")
  const toastVisible = ref(false)

  function toggleMobileMenu() {
    mobileMenuOpen.value = !mobileMenuOpen.value
  }

  function showToast(message: string, type: "success" | "error" | "warning" | "info" = "info") {
    toastMessage.value = message
    toastType.value = type
    toastVisible.value = true
    setTimeout(() => {
      toastVisible.value = false
    }, 4000)
  }

  return {
    mobileMenuOpen,
    toastMessage,
    toastType,
    toastVisible,
    toggleMobileMenu,
    showToast,
  }
})
