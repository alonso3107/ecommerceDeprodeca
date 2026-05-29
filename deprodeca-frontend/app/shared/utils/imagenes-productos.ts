// Imágenes premium verificadas vía Pexels API — Mayo 2026
// Cada URL carga con HTTP 200. Calidad: ?w=800&q=90

export const imagenesProductos = {
  bebidas: [
    'https://images.pexels.com/photos/1634074/pexels-photo-1634074.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/28268141/pexels-photo-28268141.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/3651045/pexels-photo-3651045.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/17077385/pexels-photo-17077385.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/5995769/pexels-photo-5995769.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/13871766/pexels-photo-13871766.jpeg?w=800&q=90',
  ],
  snacks: [
    'https://images.pexels.com/photos/31323236/pexels-photo-31323236.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/4617834/pexels-photo-4617834.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/6485538/pexels-photo-6485538.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/25098396/pexels-photo-25098396.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/29387144/pexels-photo-29387144.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/37304947/pexels-photo-37304947.jpeg?w=800&q=90',
  ],
  lacteos: [
    'https://images.pexels.com/photos/36183642/pexels-photo-36183642.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/7190366/pexels-photo-7190366.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/29463558/pexels-photo-29463558.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/31899745/pexels-photo-31899745.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/35740586/pexels-photo-35740586.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/18258533/pexels-photo-18258533.jpeg?w=800&q=90',
  ],
  abarrotes: [
    'https://images.pexels.com/photos/343871/pexels-photo-343871.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/11025882/pexels-photo-11025882.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/20586595/pexels-photo-20586595.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/15597774/pexels-photo-15597774.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/35134317/pexels-photo-35134317.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/7606063/pexels-photo-7606063.jpeg?w=800&q=90',
  ],
  limpieza: [
    'https://images.pexels.com/photos/12997255/pexels-photo-12997255.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/12997254/pexels-photo-12997254.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/28921817/pexels-photo-28921817.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/34686413/pexels-photo-34686413.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/9230357/pexels-photo-9230357.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/4107271/pexels-photo-4107271.jpeg?w=800&q=90',
  ],
  'cuidado-personal': [
    'https://images.pexels.com/photos/18066458/pexels-photo-18066458.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/29502137/pexels-photo-29502137.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/3750640/pexels-photo-3750640.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/8842706/pexels-photo-8842706.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/27561830/pexels-photo-27561830.jpeg?w=800&q=90',
    'https://images.pexels.com/photos/10760874/pexels-photo-10760874.jpeg?w=800&q=90',
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
  'cuidado-personal': 'cuidado-personal',
  'cuidado personal': 'cuidado-personal',
  cuidado_personal: 'cuidado-personal',
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
  return `data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='800' height='600' viewBox='0 0 800 600'%3E%3Crect width='800' height='600' fill='%23F5F0E8'/%3E%3Ctext x='48' y='84' fill='%231C1917' font-family='monospace' font-size='24' letter-spacing='2'%3E${encodeURIComponent(nombre.toUpperCase())}%3C/text%3E%3Cpath d='M72 468H728' stroke='%23C5BFB5' stroke-width='2'/%3E%3Cpath d='M562 422L712 272L800 360' stroke='%23A16207' stroke-opacity='.18' stroke-width='10' fill='none'/%3E%3C/svg%3E`
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

export function obtenerKeywordImagen(producto: ProductoComoCatalogo): string {
  const nombre = normalizarTexto(producto.nombre || '')
  const cat = producto.categoria_nombre || ''
  // Mapeo simple de palabras comunes a inglés
  const mapa: Record<string, string> = {
    leche: 'milk', bebida: 'drink', galleta: 'cookie', chocolate: 'chocolate',
    arroz: 'rice', aceite: 'oil', azucar: 'sugar', fideos: 'pasta',
    detergente: 'detergent', jabon: 'soap', shampoo: 'shampoo',
    cafe: 'coffee', yogurt: 'yogurt', queso: 'cheese', mantequilla: 'butter',
    pan: 'bread', huevo: 'egg', atun: 'tuna', conserva: 'canned',
  }
  const palabras = nombre.split(/\s+/)
  const keywords = palabras
    .map(p => mapa[normalizarTexto(p)] || p)
    .filter(p => p.length > 2)
  return `${keywords.join('+')}+${cat}+product+studio`
}
