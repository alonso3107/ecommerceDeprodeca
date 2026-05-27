-- Ejecutar en Supabase SQL Editor (https://oyapnvkunpxixxdcxics.supabase.co)
-- 1. Ir a SQL Editor en el dashboard
-- 2. Pegar este script completo
-- 3. Ejecutar (Ctrl+Enter o botón Run)

-- Tabla de consultas del formulario de contacto
CREATE TABLE IF NOT EXISTS consultas (
  id          UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  nombre      TEXT NOT NULL,
  email       TEXT NOT NULL,
  telefono    TEXT NOT NULL,
  motivo      TEXT NOT NULL,
  mensaje     TEXT NOT NULL,
  created_at  TIMESTAMPTZ DEFAULT now() NOT NULL
);

-- Índice para búsqueda por fecha
CREATE INDEX IF NOT EXISTS idx_consultas_created_at ON consultas (created_at DESC);

-- Políticas de seguridad (RLS)
ALTER TABLE consultas ENABLE ROW LEVEL SECURITY;

-- Permitir INSERT público (para el formulario de contacto)
CREATE POLICY "Permitir insertar consultas"
  ON consultas FOR INSERT
  TO anon
  WITH CHECK (true);

-- Permitir SELECT solo al usuario autenticado (admin)
CREATE POLICY "Admin puede ver consultas"
  ON consultas FOR SELECT
  TO authenticated
  USING (true);
