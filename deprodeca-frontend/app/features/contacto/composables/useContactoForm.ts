type CampoFormulario = "nombre" | "email" | "telefono" | "mensaje"

type ContactCard = {
  title: string
  value: string
  description: string
  href: string | null
  action: string
  icon: "phone" | "mail" | "pin"
}

export function useContactoForm() {
  const supabase = useSupabase()

  const formulario = reactive({
    nombre: "",
    email: "",
    telefono: "",
    motivo: "",
    mensaje: "",
  })

  const touched = reactive<Record<CampoFormulario, boolean>>({
    nombre: false,
    email: false,
    telefono: false,
    mensaje: false,
  })

  const submitting = ref(false)
  const submitted = ref(false)
  const attemptedSubmit = ref(false)
  const formError = ref("")
  const copied = ref(false)
  const isReady = ref(false)

  const phoneDisplay = "+51 900 000 000"
  const phoneDigits = "51900000000"
  const whatsappUrl = `https://wa.me/${phoneDigits}`
  const emailAddress = "contacto@deprodeca.pe"

  const motivoOptions = [
    "Quiero registrarme como bodega",
    "Consulta sobre pedidos",
    "Soporte técnico",
    "Devolución o reclamo",
    "Otro",
  ] as const

  const contactCards: ContactCard[] = [
    {
      title: "WhatsApp",
      value: phoneDisplay,
      description: "Respuesta prioritaria para pedidos y coordinación comercial.",
      href: whatsappUrl,
      action: "Abrir chat",
      icon: "phone",
    },
    {
      title: "Correo",
      value: emailAddress,
      description: "Para consultas, soporte y seguimiento de solicitudes.",
      href: `mailto:${emailAddress}`,
      action: "Enviar mail",
      icon: "mail",
    },
    {
      title: "Cobertura",
      value: "Ica, Chincha, Pisco, Nazca",
      description: "Atendemos bodegas y negocios con atención regional.",
      href: null,
      action: "",
      icon: "pin",
    },
  ]

  const horarios = [
    { dia: "Lunes - Viernes", hora: "8:00 - 18:00" },
    { dia: "Sábado", hora: "9:00 - 13:00" },
    { dia: "Domingo", hora: "Cerrado" },
  ] as const

  const responseStats = [
    { label: "Tiempo de respuesta", value: "< 24h" },
    { label: "Canales activos", value: "3" },
    { label: "Cobertura", value: "Regional" },
  ] as const

  const formState = computed<Record<CampoFormulario, string | null>>(() => {
    const errors: Record<CampoFormulario, string | null> = {
      nombre: null,
      email: null,
      telefono: null,
      mensaje: null,
    }

    const nombre = formulario.nombre.trim()
    const email = formulario.email.trim()
    const telefono = formulario.telefono.trim()
    const mensaje = formulario.mensaje.trim()

    if (nombre.length > 0 && nombre.length < 2) errors.nombre = "Mínimo 2 caracteres"
    if (nombre.length === 0) errors.nombre = "El nombre es obligatorio"

    if (email.length === 0) {
      errors.email = "El email es obligatorio"
    } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)) {
      errors.email = "Formato de email inválido"
    }

    if (telefono.length > 0 && !/^\d{9}$/.test(telefono)) {
      errors.telefono = "Usa 9 dígitos, sin espacios"
    }

    if (mensaje.length === 0) {
      errors.mensaje = "El mensaje es obligatorio"
    } else if (mensaje.length < 10) {
      errors.mensaje = "Mínimo 10 caracteres"
    }

    return errors
  })

  const canSubmit = computed(() => Object.values(formState.value).every((value) => value === null))

  const showFieldFeedback = (field: CampoFormulario) => touched[field] || attemptedSubmit.value

  const fieldStatus = (field: CampoFormulario) => {
    if (!showFieldFeedback(field)) return "idle"
    return formState.value[field] ? "error" : "valid"
  }

  function markTouched(field: CampoFormulario) {
    touched[field] = true
  }

  function sanitizePhone(event: Event) {
    const input = event.target as HTMLInputElement | null
    if (!input) return
    const clean = input.value.replace(/\D/g, "").slice(0, 9)
    if (clean !== formulario.telefono) formulario.telefono = clean
  }

  async function copyPhoneNumber() {
    if (!import.meta.client || !navigator.clipboard) return
    try {
      await navigator.clipboard.writeText(phoneDisplay)
      copied.value = true
      window.setTimeout(() => {
        copied.value = false
      }, 1800)
    } catch {
      copied.value = false
    }
  }

  function resetForm() {
    formulario.nombre = ""
    formulario.email = ""
    formulario.telefono = ""
    formulario.motivo = ""
    formulario.mensaje = ""
    touched.nombre = false
    touched.email = false
    touched.telefono = false
    touched.mensaje = false
  }

  async function enviarFormulario() {
    attemptedSubmit.value = true
    submitted.value = false
    formError.value = ""

    touched.nombre = true
    touched.email = true
    touched.telefono = true
    touched.mensaje = true

    if (!canSubmit.value || submitting.value) {
      formError.value = "Revisá los campos marcados antes de enviar."
      return
    }

    submitting.value = true

    try {
      const { error } = await supabase.from("consultas").insert({
        nombre: formulario.nombre.trim(),
        email: formulario.email.trim(),
        telefono: formulario.telefono.trim(),
        motivo: formulario.motivo || "No especificado",
        mensaje: formulario.mensaje.trim(),
      })

      if (error) throw error

      submitted.value = true
      attemptedSubmit.value = false
      resetForm()
      window.setTimeout(() => {
        submitted.value = false
      }, 5000)
    } catch (error: any) {
      formError.value = error?.message || "No pudimos enviar tu mensaje. Intentá otra vez."
    } finally {
      submitting.value = false
    }
  }

  onMounted(() => {
    window.setTimeout(() => {
      isReady.value = true
    }, 60)
  })

  return {
    formulario,
    formState,
    formError,
    submitted,
    submitting,
    copied,
    isReady,
    canSubmit,
    contactCards,
    horarios,
    responseStats,
    motivoOptions,
    whatsappUrl,
    copyPhoneNumber,
    markTouched,
    sanitizePhone,
    fieldStatus,
    enviarFormulario,
  }
}
