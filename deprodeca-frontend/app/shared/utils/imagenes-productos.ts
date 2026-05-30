// Imágenes locales del catálogo.

const asset = (file: string): string => new URL(`../../assets/img/${file}`, import.meta.url).href;

const imagenesPorProducto: Record<string, string> = {
  "menestra lentejas": asset("menestra-lentejas.jpg"),
  "leche gloria evap 400g": asset("leche-gloria-evap-400g.jpg"),
  "leche ideal entera": asset("leche-ideal-entera.jpg"),
  "leche ideal cremosita": asset("leche-ideal-cremosita.jpg"),
  "yogu yogu fresa 1l": asset("yogu-yogu-fresa-1l.jpg"),
  "yogu yogu vainilla 1l": asset("yogu-yogu-vainilla-1l.jpg"),
  "sublime clasico 30g": asset("sublime-clasico-30g.jpg"),
  "donofrio six pack 6 und": asset("donofrio-six-pack-6-und.jpg"),
  "caja sublime x24und": asset("caja-sublime-x24und.jpg"),
  "triangulo 36g": asset("triangulo-36g.jpg"),
  "morochas 36g": asset("morochas-36g.jpg"),
  "princesa 30g": asset("princesa-30g.jpg"),
  "nescafe tradicion 200g": asset("nescafe-tradicion-200g.jpg"),
  "nescafe capuccino 10 sobres": asset("nescafe-capuccino-10-sobres.jpg"),
  "milo listo 200ml": asset("milo-listo-200ml.jpg"),
  "nestea limon 1.5l": asset("nestea-limon-1-5l.jpg"),
  "nescafe dolca 170g": asset("nescafe-dolca-170g.jpg"),
  "milo 400g": asset("milo-400g.jpg"),
  "maggi cubito gallina": asset("maggicubitogallina.jpg"),
  "maggi sopa criolla": asset("maggisopacriolla.jpg"),
  "doña gusta": asset("doñagusta.png"),
  clorox: asset("lejiaclorox.jpg"),
  lejia: asset("lejiaclorox.jpg"),
  limpiador: asset("limpiadormultiusos.jpg"),
  multiusos: asset("limpiadormultiusos.jpg"),
  "jabon liquido": asset("jabonliquido.png"),
  protex: asset("jabonprotex.jpg"),
  shampoo: asset("shampoosavital.png"),
  savital: asset("shampoosavital.png"),
  pepsodent: asset("creampepsodent.png"),
  desodorante: asset("desodoroantespray.png"),
  spray: asset("desodoroantespray.png"),
  bolivar: asset("bolivar.jpg"),
};

export const imagenesProductos = {
  bebidas: [
    asset("doñagusta.png"),
    asset("maggicubitogallina.jpg"),
    asset("maggisopacriolla.jpg"),
    asset("jabonliquido.png"),
    asset("bolivar.jpg"),
    asset("shampoosavital.png"),
  ],
  snacks: [
    asset("doñagusta.png"),
    asset("maggicubitogallina.jpg"),
    asset("maggisopacriolla.jpg"),
    asset("creampepsodent.png"),
    asset("jabonprotex.jpg"),
    asset("desodoroantespray.png"),
  ],
  lacteos: [
    asset("creampepsodent.png"),
    asset("jabonliquido.png"),
    asset("bolivar.jpg"),
    asset("limpiadormultiusos.jpg"),
    asset("lejiaclorox.jpg"),
    asset("jabonprotex.jpg"),
  ],
  abarrotes: [
    asset("doñagusta.png"),
    asset("maggicubitogallina.jpg"),
    asset("maggisopacriolla.jpg"),
    asset("limpiadormultiusos.jpg"),
    asset("lejiaclorox.jpg"),
    asset("bolivar.jpg"),
  ],
  limpieza: [
    asset("lejiaclorox.jpg"),
    asset("limpiadormultiusos.jpg"),
    asset("jabonliquido.png"),
    asset("bolivar.jpg"),
    asset("jabonprotex.jpg"),
    asset("desodoroantespray.png"),
  ],
  "cuidado-personal": [
    asset("shampoosavital.png"),
    asset("desodoroantespray.png"),
    asset("jabonprotex.jpg"),
    asset("creampepsodent.png"),
    asset("jabonliquido.png"),
    asset("bolivar.jpg"),
  ],
} as const;

type ProductoComoCatalogo = {
  id?: number | string;
  slug?: string;
  nombre?: string;
  categoria_nombre?: string | null;
  categoria_slug?: string | null;
};

const aliasCategorias: Record<string, keyof typeof imagenesProductos> = {
  bebidas: "bebidas",
  bebida: "bebidas",
  snacks: "snacks",
  snack: "snacks",
  lacteos: "lacteos",
  lacteo: "lacteos",
  abarrotes: "abarrotes",
  limpieza: "limpieza",
  "cuidado-personal": "cuidado-personal",
  "cuidado personal": "cuidado-personal",
  cuidado_personal: "cuidado-personal",
};

const imagenesGlobales = Object.values(imagenesProductos).flat();

function normalizarTexto(valor: string): string {
  return valor
    .normalize("NFD")
    .replace(/[\u0300-\u036f]/g, "")
    .replace(/[.,/\\_\-]+/g, " ")
    .toLowerCase()
    .trim();
}

function resolverImagenExacta(producto: ProductoComoCatalogo): string | null {
  const claves = [producto.nombre || "", producto.slug || ""];
  for (const clave of claves) {
    const normalizada = normalizarTexto(clave);
    if (normalizada in imagenesPorProducto) {
      return imagenesPorProducto[normalizada];
    }
  }
  return null;
}

function obtenerClaveCategoria(
  producto: ProductoComoCatalogo,
): keyof typeof imagenesProductos | null {
  const claves = [
    producto.categoria_slug || "",
    producto.categoria_nombre || "",
  ];
  for (const clave of claves) {
    const normalizada = normalizarTexto(clave);
    if (normalizada in aliasCategorias) {
      return aliasCategorias[normalizada];
    }
  }
  return null;
}

function obtenerIndiceSemilla(
  producto: ProductoComoCatalogo,
  indice: number,
): number {
  const semilla = `${producto.slug || producto.id || producto.nombre || "producto"}-${indice}`;
  let hash = 0;
  for (let i = 0; i < semilla.length; i += 1) {
    hash = (hash * 31 + semilla.charCodeAt(i)) >>> 0;
  }
  return hash;
}

function crearFallbackUnico(
  producto: ProductoComoCatalogo,
  indice: number,
): string {
  const nombre =
    normalizarTexto(producto.nombre || producto.slug || `producto-${indice}`) ||
    `producto-${indice}`;
  return `data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='800' height='600' viewBox='0 0 800 600'%3E%3Crect width='800' height='600' fill='%23F5F0E8'/%3E%3Ctext x='48' y='84' fill='%231C1917' font-family='monospace' font-size='24' letter-spacing='2'%3E${encodeURIComponent(nombre.toUpperCase())}%3C/text%3E%3Cpath d='M72 468H728' stroke='%23C5BFB5' stroke-width='2'/%3E%3Cpath d='M562 422L712 272L800 360' stroke='%23A16207' stroke-opacity='.18' stroke-width='10' fill='none'/%3E%3C/svg%3E`;
}

export function resolverImagenProducto(
  producto: ProductoComoCatalogo,
  indice: number,
  imagenesUsadas: Set<string>,
): string {
  const imagenExacta = resolverImagenExacta(producto);
  if (imagenExacta) {
    imagenesUsadas.add(imagenExacta);
    return imagenExacta;
  }

  const claveCategoria = obtenerClaveCategoria(producto);
  const semilla = obtenerIndiceSemilla(producto, indice);
  const candidatos = [
    ...(claveCategoria ? imagenesProductos[claveCategoria] : []),
    ...imagenesGlobales,
  ];

  for (let i = 0; i < candidatos.length; i += 1) {
    const candidato = candidatos[(semilla + i) % candidatos.length];
    if (!imagenesUsadas.has(candidato)) {
      imagenesUsadas.add(candidato);
      return candidato;
    }
  }

  if (candidatos.length > 0) {
    const candidato = candidatos[semilla % candidatos.length];
    imagenesUsadas.add(candidato);
    return candidato;
  }

  return crearFallbackUnico(producto, indice);
}
