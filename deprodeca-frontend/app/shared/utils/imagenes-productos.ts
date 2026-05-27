export const imagenesProductos = {
  bebidas: [
    'https://images.unsplash.com/photo-1613510214658-f1e56b099811?w=600&q=80',
    'https://images.unsplash.com/photo-1761281254116-82a6bded6c30?w=600&q=80',
    'https://images.unsplash.com/photo-1650547001322-145ff2c0bed7?w=600&q=80',
  ],
  snacks: [
    'https://images.unsplash.com/photo-1708746333890-8e775f97f0a6?w=600&q=80',
    'https://images.unsplash.com/photo-1768475011729-80f491cd1c86?w=600&q=80',
    'https://images.unsplash.com/photo-1748765968965-7e18d4f7192b?w=600&q=80',
  ],
  lacteos: [
    'https://images.unsplash.com/photo-1760273464017-4bb7dfa42d91?w=600&q=80',
    'https://images.unsplash.com/photo-1771503937636-c0e37bd23517?w=600&q=80',
    'https://images.unsplash.com/photo-1723367456186-8f6c7d3436fe?w=600&q=80',
  ],
  abarrotes: [
    'https://images.unsplash.com/photo-1754391853169-88024d46a1d8?w=600&q=80',
    'https://images.unsplash.com/photo-1762926627902-c0846725e69e?w=600&q=80',
    'https://images.unsplash.com/photo-1761281254116-82a6bded6c30?w=600&q=80',
  ],
  limpieza: [
    'https://images.unsplash.com/photo-1762549564478-7a65b07f939f?w=600&q=80',
    'https://images.unsplash.com/photo-1754391853169-88024d46a1d8?w=600&q=80',
    'https://images.unsplash.com/photo-1732580984109-5d77c0b5bc73?w=600&q=80',
  ],
} as const

type ProductoComoCatalogo = {
  id?: number | string
  slug?: string
  nombre?: string
  imagen_url?: string | null
  categoria_nombre?: string | null
  categoria_slug?: string | null
}

const aliasCategorias: Record<string, keyof typeof imagenesProductos> = {
  bebidas: 'bebidas',
  bebida: 'bebidas',
  snacks: 'snacks',
  snack: 'snacks',
  lacteos: 'lacteos',
  lacteo: 'lacteos',
  abarrotes: 'abarrotes',
  limpieza: 'limpieza',
}

const imagenesGlobales = Object.values(imagenesProductos).flat()

function normalizarTexto(valor: string): string {
  return valor
    .normalize('NFD')
    .replace(/[\u0300-\u036f]/g, '')
    .toLowerCase()
    .trim()
}

function obtenerClaveCategoria(producto: ProductoComoCatalogo): keyof typeof imagenesProductos | null {
  const claves = [producto.categoria_slug || '', producto.categoria_nombre || '']
  for (const clave of claves) {
    const normalizada = normalizarTexto(clave)
    if (normalizada in aliasCategorias) {
      return aliasCategorias[normalizada]
    }
  }
  return null
}

function obtenerIndiceSemilla(producto: ProductoComoCatalogo, indice: number): number {
  const semilla = `${producto.slug || producto.id || producto.nombre || 'producto'}-${indice}`
  let hash = 0
  for (let i = 0; i < semilla.length; i += 1) {
    hash = (hash * 31 + semilla.charCodeAt(i)) >>> 0
  }
  return hash
}

function crearFallbackUnico(producto: ProductoComoCatalogo, indice: number): string {
  const nombre = normalizarTexto(producto.nombre || producto.slug || `producto-${indice}`) || `producto-${indice}`
  return `data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='800' height='600' viewBox='0 0 800 600'%3E%3Crect width='800' height='600' fill='%23FAFAFA'/%3E%3Ctext x='48' y='84' fill='%23171717' font-family='monospace' font-size='24' letter-spacing='2'%3E${encodeURIComponent(nombre.toUpperCase())}%3C/text%3E%3Cpath d='M72 468H728' stroke='%23e5e5e5' stroke-width='2'/%3E%3Cpath d='M562 422L712 272L800 360' stroke='%23D4A017' stroke-opacity='.18' stroke-width='10' fill='none'/%3E%3Cpath d='M612 512L720 404L784 468' stroke='%23D4A017' stroke-opacity='.18' stroke-width='10' fill='none'/%3E%3Cpath d='M654 150H734V230Z' fill='none' stroke='%23D4A017' stroke-opacity='.14' stroke-width='8'/%3E%3C/svg%3E`
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
    if (!imagenesUsadas.has(candidato)) {
      imagenesUsadas.add(candidato)
      return candidato
    }
  }

  return crearFallbackUnico(producto, indice)
}
