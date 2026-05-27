# CODEX.md — DEPRODECA

Instrucciones para Codex CLI. Este archivo se carga automáticamente en cada sesión.

## Idioma

- **TODO en español:** variables, funciones, comentarios, mensajes de commit, UI copy, logs, errores.
- Única excepción: nombres de paquetes/dependencias que ya están en inglés.

## Stack

| Capa | Tecnología |
|------|-----------|
| Frontend | Nuxt 4 + Vue 3 + Tailwind CSS v4 |
| Backend | Go + Fiber v2 + PostgreSQL (pgx) |
| Package manager | **pnpm** (NO npm, NO yarn) |
| Runtime | Node 22+, Go 1.23+ |

### Frontend: `deprodeca-frontend/`

- **pnpm** para todo (`pnpm install`, `pnpm dev`, `pnpm build`)
- Tailwind **v4** (CSS-first config con `@import "tailwindcss"`, NO `tailwind.config.js`)
- Componentes Vue 3 con `<script setup lang="ts">`
- Pages en `pages/`, components en `components/`, composables en `composables/`
- Rutas con Nuxt file-based routing
- Testing: `pnpm test` (vitest)

### Backend: `deprodeca-backend/`

- Estructura: `cmd/api/`, `internal/handler/`, `internal/model/`, `internal/repository/`, `internal/service/`
- PostgreSQL con pgx. Conexión vía `DATABASE_URL` en `.env`
- Migraciones en `migrations/`
- Compilar: `go build ./cmd/api/`
- Ejecutar: `go run ./cmd/api/`
- Variables de entorno en `.env`

## Diseño — Estilo Brutalista

**NO USAR:** sombras (`box-shadow`), bordes redondeados (`border-radius`), gradientes, blur, glassmorphism.

**USAR:**
- Esquinas rectas (0px radius)
- Gap pequeño (`gap-px`, `gap-0.5`)
- SVG geométricos únicos como elementos decorativos
- Alto contraste: fondo `#171717`, texto `#FAFAFA`, acento `#D4A017` (dorado)
- Hover **sutil**: `scale: 1.02` máximo, `transition: 500-700ms`, sin efectos agresivos
- Tipografía: Cinzel Decorative (headings), Big Noodle Titling (display)

## Commits

- Formato **Conventional Commits** estándar
- **Sin emojis** en los mensajes
- Ejemplos: `feat: validacion de stock en carrito`, `fix: correccion de overflow en navbar`, `refactor: extraer logica de auth a servicio`

## Reglas generales

1. **Nunca borres código sin preguntar** si no estás 100% seguro de que es basura.
2. **Respeta el estilo existente.** Si el archivo usa tabs, usá tabs. Si usa 2 espacios, 2 espacios.
3. **No reinventes.** Si ya hay un patrón en el proyecto (ej: cómo se hacen los handlers), seguilo.
4. **Commiteá al terminar** con mensaje en español y formato conventional commits.
5. **Antes de commit:** corré tests (`pnpm test` en frontend, `go test ./...` en backend).
6. **Si algo falla, debuggeá** — no pidas ayuda a la primera. Leé logs, seguí el stack trace.
7. **TypeScript estricto.** No usés `any` a menos que sea absolutamente necesario.

## Intentos de arreglo

Si encontrás un error, intentá arreglarlo **hasta 3 veces** por tu cuenta antes de reportarlo. En cada intento, probá un enfoque distinto.
