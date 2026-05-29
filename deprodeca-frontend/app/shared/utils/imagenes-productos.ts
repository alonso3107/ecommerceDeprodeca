export const imagenesProductos = {
  bebidas: [
    "https://images.unsplash.com/photo-1544145945-f90425340c7e?w=800&q=90",
    "https://images.unsplash.com/photo-1622483767028-3f66f32aef97?w=800&q=90",
    "https://images.unsplash.com/photo-1527960471264-932f39eb5846?w=800&q=90",
  ],
  snacks: [
    "https://images.unsplash.com/photo-1621939514649-280e2ee25f60?w=800&q=90",
    "https://images.unsplash.com/photo-1582058091505-f87a2e55a40f?w=800&q=90",
    "https://images.unsplash.com/photo-1606312619070-d48b4c652a52?w=800&q=90",
  ],
  lacteos: [
    "https://images.unsplash.com/photo-1550583724-b2692b85b150?w=800&q=90",
    "https://images.unsplash.com/photo-1576186726580-a9e9b9d9b0f3?w=800&q=90",
    "https://images.unsplash.com/photo-1563636619-e9143da7973b?w=800&q=90",
  ],
  abarrotes: [
    "https://images.unsplash.com/photo-1586201375761-83865001e31c?w=800&q=90",
    "https://images.unsplash.com/photo-1579113800032-c38bd7635818?w=800&q=90",
    "https://images.unsplash.com/photo-1534723452862-4c874018d66d?w=800&q=90",
  ],
  limpieza: [
    "https://images.unsplash.com/photo-1581578731548-c64695cc6952?w=800&q=90",
    "https://images.unsplash.com/photo-1596803244618-8dbee441d70b?w=800&q=90",
    "https://images.unsplash.com/photo-1603796846097-bee99e4a601f?w=800&q=90",
  ],
} as const

type ProductoComoCatalogo = {
  id?: number | string
  slug?: string
  nombre?: string
  descripcion?: string | null
  imagen_url?: string | null
  categoria_nombre?: string | null
  categoria_slug?: string | null
}

type ClaveCategoria = keyof typeof imagenesProductos

const aliasCategorias: Record<string, ClaveCategoria> = {
  bebidas: "bebidas",
  bebida: "bebidas",
  snacks: "snacks",
  snack: "snacks",
  lacteos: "lacteos",
  lacteo: "lacteos",
  abarrotes: "abarrotes",
  limpieza: "limpieza",
}

const imagenesGlobales = Object.values(imagenesProductos).flat()

function normalizarTexto(valor: string): string {
  return valor
    .normalize("NFD")
    .replace(/[\u0300-\u036f]/g, "")
    .toLowerCase()
    .trim()
}

function obtenerClaveCategoria(producto: ProductoComoCatalogo): ClaveCategoria | null {
  const claves = [producto.categoria_slug || "", producto.categoria_nombre || ""]
  for (const clave of claves) {
    const normalizada = normalizarTexto(clave)
    if (normalizada in aliasCategorias) {
      const categoria = aliasCategorias[normalizada as keyof typeof aliasCategorias]
      if (categoria) return categoria
    }
  }
  return null
}

export function obtenerKeywordImagen(producto: ProductoComoCatalogo): string {
  const texto = normalizarTexto(`${producto.nombre || ""} ${producto.descripcion || ""}`)

  if (texto.includes("chocolate") || texto.includes("galleta") || texto.includes("snack")) {
    return "chocolate+bar+premium+studio+food"
  }
  if (texto.includes("leche") || texto.includes("yogurt") || texto.includes("queso")) {
    return "dairy+product+premium+studio+packaging"
  }
  if (texto.includes("detergente") || texto.includes("limpieza") || texto.includes("jabon")) {
    return "cleaning+product+premium+studio+packaging"
  }
  if (texto.includes("arroz") || texto.includes("aceite") || texto.includes("azucar")) {
    return "grocery+packaging+premium+studio+food"
  }
  if (texto.includes("bebida") || texto.includes("agua") || texto.includes("refresco")) {
    return "beverage+bottle+premium+studio+product"
  }

  return "packaged+food+premium+studio+product"
}

function obtenerIndiceSemilla(producto: ProductoComoCatalogo, indice: number): number {
  const semilla = `${producto.slug || producto.id || producto.nombre || "producto"}-${indice}`
  let hash = 0
  for (let i = 0; i < semilla.length; i += 1) {
    hash = (hash * 31 + semilla.charCodeAt(i)) >>> 0
  }
  return hash
}

function crearFallbackUnico(producto: ProductoComoCatalogo, indice: number): string {
  const nombre = normalizarTexto(producto.nombre || producto.slug || `producto-${indice}`) || `producto-${indice}`
  return `data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='800' height='960' viewBox='0 0 800 960'%3E%3Crect width='800' height='960' fill='%23FAFAF9'/%3E%3Ctext x='56' y='116' fill='%231C1917' font-family='monospace' font-size='24' letter-spacing='2'%3E${encodeURIComponent(nombre.toUpperCase())}%3C/text%3E%3Cpath d='M72 840H728' stroke='%23D6D3D1' stroke-width='2'/%3E%3Cpath d='M640 760L728 672L800 744' stroke='%23A16207' stroke-opacity='.2' stroke-width='8' fill='none'/%3E%3C/svg%3E`
}

export function resolverImagenProducto(
  producto: ProductoComoCatalogo,
  indice: number,
  imagenesUsadas: Set<string>,
): string {
  const imagenBackend = producto.imagen_url?.trim()
  if (imagenBackend && !imagenesUsadas.has(imagenBackend)) {
    imagenesUsadas.add(imagenBackend)
    return imagenBackend
  }

  const claveCategoria = obtenerClaveCategoria(producto)
  const semilla = obtenerIndiceSemilla(producto, indice)
  const candidatos = [
    ...(claveCategoria ? imagenesProductos[claveCategoria] : []),
    ...imagenesGlobales,
  ]

  for (let i = 0; i < candidatos.length; i += 1) {
    const candidato = candidatos[(semilla + i) % candidatos.length]
    if (candidato && !imagenesUsadas.has(candidato)) {
      imagenesUsadas.add(candidato)
      return candidato
    }
  }

  return crearFallbackUnico(producto, indice)
}
