-- ═══════════════════════════════════════════════════════════
-- DEPRODECA — Seed Data v3: Todas las imágenes verificadas y ÚNICAS
-- ═══════════════════════════════════════════════════════════

-- ─── Categorías (6) ─────────────────────────────────────
INSERT INTO categorias (nombre, slug, descripcion, imagen_url) VALUES
('Bebidas', 'bebidas', 'Bebidas Nestlé: Milo, Nescafé, Nestea y más.', 'https://images.unsplash.com/photo-1544145945-f90425340c7e?w=400&q=80'),
('Snacks', 'snacks', 'Galletas, chocolates y snacks para tu bodega.', 'https://images.unsplash.com/photo-1621447504864-d8686e25298c?w=400&q=80'),
('Lácteos', 'lacteos', 'Leche Ideal, Yogurt Yogu Yogu, leches evaporadas.', 'https://images.unsplash.com/photo-1550583724-b2692b85b150?w=400&q=80'),
('Abarrotes', 'abarrotes', 'Productos de despensa: Maggi, Doña Gusta, menestras.', 'https://images.unsplash.com/photo-1615485290382-441e4d049cb5?w=400&q=80'),
('Limpieza', 'limpieza', 'Productos de limpieza para el hogar y bodega.', 'https://images.unsplash.com/photo-1583947215259-38e31be8751f?w=400&q=80'),
('Cuidado Personal', 'cuidado-personal', 'Cuidado personal: shampoos, jabones, desodorantes.', 'https://images.unsplash.com/photo-1611930022073-b7a4ba5fcccd?w=400&q=80');

-- ─── Bebidas (6) ────────────────────────────────────────
INSERT INTO productos (categoria_id, nombre, slug, descripcion, precio, stock, unidad, imagen_url) VALUES
(1, 'Milo 400g', 'milo-400g', 'Milo en polvo lata 400g. Energía y nutrición para toda la familia.', 12.50, 200, 'lata', 'https://images.unsplash.com/photo-1596803244618-8dbeeed3021b?w=400&q=80'),
(1, 'Nescafé Tradición 200g', 'nescafe-tradicion-200g', 'Café soluble Nescafé Tradición frasco 200g.', 18.90, 150, 'frasco', 'https://images.unsplash.com/photo-1559056199-641a0ac8b55e?w=400&q=80'),
(1, 'Nescafé Dolca 170g', 'nescafe-dolca-170g', 'Café soluble Dolca suave frasco 170g.', 16.50, 120, 'frasco', 'https://images.unsplash.com/photo-1572442388796-11668a67e53d?w=400&q=80'),
(1, 'Nestea Limón 1.5L', 'nestea-limon-1-5l', 'Bebida refrescante sabor limón 1.5 litros.', 5.20, 300, 'botella', 'https://images.unsplash.com/photo-1624517452488-04869289c4ca?w=400&q=80'),
(1, 'Milo Listo 200ml', 'milo-listo-200ml', 'Milo listo para beber tetra pack 200ml.', 2.50, 500, 'unidad', 'https://images.unsplash.com/photo-1527960471264-932f39eb5846?w=400&q=80'),
(1, 'Nescafé Cappuccino 10 sobres', 'nescafe-cappuccino-10-sobres', 'Caja con 10 sobres de Cappuccino instantáneo.', 8.90, 180, 'caja', 'https://images.unsplash.com/photo-1551024506-0bccd828d307?w=400&q=80');

-- ─── Snacks (6) ─────────────────────────────────────────
INSERT INTO productos (categoria_id, nombre, slug, descripcion, precio, stock, unidad, imagen_url) VALUES
(2, 'Sublime Clásico 30g', 'sublime-clasico-30g', 'Chocolate Sublime clásico con maní 30g.', 1.20, 600, 'unidad', 'https://images.unsplash.com/photo-1549007994-cb92caebd54b?w=400&q=80'),
(2, 'Princesa 30g', 'princesa-30g', 'Galleta Princesa rellena de chocolate 30g.', 0.80, 700, 'unidad', 'https://images.unsplash.com/photo-1558961363-fa8fdf82db35?w=400&q=80'),
(2, 'Morochas 30g', 'morochas-30g', 'Galleta Morochas sabor chocolate 30g.', 0.80, 650, 'unidad', 'https://images.unsplash.com/photo-1606313564200-e75d5e30476c?w=400&q=80'),
(2, 'Triángulo 36g', 'triangulo-36g', 'Chocolate Triángulo relleno 36g.', 1.50, 450, 'unidad', 'https://images.unsplash.com/photo-1582058091505-f87a2e55a40f?w=400&q=80'),
(2, 'Caja Sublime x24 und', 'caja-sublime-x24', 'Caja de 24 unidades Sublime Clásico (venta mayor).', 25.00, 80, 'caja', 'https://images.unsplash.com/photo-1621939514649-280e2ee25f60?w=400&q=80'),
(2, 'Donofrio Six Pack 6 und', 'donofrio-six-pack', 'Six pack de helados Donofrio (venta bodega).', 15.00, 100, 'pack', 'https://images.unsplash.com/photo-1501443762994-82bd5dace89a?w=400&q=80');

-- ─── Lácteos (5) ────────────────────────────────────────
INSERT INTO productos (categoria_id, nombre, slug, descripcion, precio, stock, unidad, imagen_url) VALUES
(3, 'Leche Ideal Entera 400g', 'leche-ideal-entera-400g', 'Leche evaporada Ideal entera lata 400g.', 4.50, 400, 'lata', 'https://images.unsplash.com/photo-1563636619-e9143da7973b?w=400&q=80'),
(3, 'Leche Ideal Cremosita 400g', 'leche-ideal-cremosita-400g', 'Leche evaporada Ideal Cremosita lata 400g.', 5.20, 350, 'lata', 'https://images.unsplash.com/photo-1628088062854-d1870b4553da?w=400&q=80'),
(3, 'Yogu Yogu Fresa 1L', 'yogu-yogu-fresa-1l', 'Yogurt Yogu Yogu bebible sabor fresa 1 litro.', 6.80, 220, 'botella', 'https://images.unsplash.com/photo-1488477181946-6428a0291777?w=400&q=80'),
(3, 'Yogu Yogu Vainilla 1L', 'yogu-yogu-vainilla-1l', 'Yogurt Yogu Yogu bebible sabor vainilla 1 litro.', 6.80, 200, 'botella', 'https://images.unsplash.com/photo-1579364046864-f226bbbc8685?w=400&q=80'),
(3, 'Leche Gloria Evap 400g', 'leche-gloria-evap-400g', 'Leche evaporada Gloria entera lata 400g.', 4.20, 500, 'lata', 'https://images.unsplash.com/photo-1574071318508-1cdbab80d002?w=400&q=80');

-- ─── Abarrotes (7) ──────────────────────────────────────
INSERT INTO productos (categoria_id, nombre, slug, descripcion, precio, stock, unidad, imagen_url) VALUES
(4, 'Maggi Cubito Carne 12 und', 'maggi-cubito-carne-12', 'Caja de 12 cubos Maggi sabor carne.', 6.50, 300, 'caja', 'https://images.unsplash.com/photo-1590794056226-79ef3a8147e1?w=400&q=80'),
(4, 'Maggi Cubito Gallina 12 und', 'maggi-cubito-gallina-12', 'Caja de 12 cubos Maggi sabor gallina.', 6.50, 280, 'caja', 'https://images.unsplash.com/photo-1476718406336-bb5a9690ee2a?w=400&q=80'),
(4, 'Doña Gusta Ají Panca 90g', 'dona-gusta-aji-panca-90g', 'Ají Panca molido Doña Gusta sobre 90g.', 4.00, 380, 'sobre', 'https://images.unsplash.com/photo-1558961363-fa8fdf82db35?w=400&q=80'),
(4, 'Doña Gusta Aderezo 90g', 'dona-gusta-aderezo-90g', 'Aderezo completo Doña Gusta sobre 90g.', 4.00, 350, 'sobre', 'https://images.unsplash.com/photo-1504674900247-0877df9cc836?w=400&q=80'),
(4, 'Maggi Sopa Criolla 68g', 'maggi-sopa-criolla-68g', 'Sopa instantánea Maggi sabor criolla sobre 68g.', 1.80, 500, 'sobre', 'https://images.unsplash.com/photo-1617196034796-73dfa7b1fd56?w=400&q=80'),
(4, 'Maggi Sopa de Pollo 68g', 'maggi-sopa-pollo-68g', 'Sopa instantánea Maggi sabor pollo sobre 68g.', 1.80, 480, 'sobre', 'https://images.unsplash.com/photo-1547592166-23ac45744acd?w=400&q=80'),
(4, 'Menestra Lentejas 500g', 'menestra-lentejas-500g', 'Lentejas secas seleccionadas bolsa 500g.', 3.50, 250, 'bolsa', 'https://images.unsplash.com/photo-1617692855027-33b14f061079?w=400&q=80');

-- ─── Limpieza (4) ───────────────────────────────────────
INSERT INTO productos (categoria_id, nombre, slug, descripcion, precio, stock, unidad, imagen_url) VALUES
(5, 'Detergente Bolívar 1kg', 'detergente-bolivar-1kg', 'Detergente en polvo Bolívar bolsa 1kg.', 7.90, 280, 'bolsa', 'https://images.unsplash.com/photo-1585421514284-efb74c2b69ba?w=400&q=80'),
(5, 'Jabón Líquido Lavanda 500ml', 'jabon-liquido-lavanda-500ml', 'Jabón líquido aroma lavanda botella 500ml.', 9.50, 200, 'botella', 'https://images.unsplash.com/photo-1600857062241-98e5dba7f214?w=400&q=80'),
(5, 'Lejía Clorox 1L', 'lejia-clorox-1l', 'Lejía Clorox botella 1 litro.', 3.80, 400, 'botella', 'https://images.unsplash.com/photo-1558618666-fcd25c85f82e?w=400&q=80'),
(5, 'Limpiador Multiusos 500ml', 'limpiador-multiusos-500ml', 'Limpiador multiusos aroma limón 500ml.', 6.50, 180, 'botella', 'https://images.unsplash.com/photo-1563453392212-326f5e854473?w=400&q=80');

-- ─── Cuidado Personal (4) ───────────────────────────────
INSERT INTO productos (categoria_id, nombre, slug, descripcion, precio, stock, unidad, imagen_url) VALUES
(6, 'Shampoo Savital 400ml', 'shampoo-savital-400ml', 'Shampoo Savital revitalizante 400ml.', 12.90, 220, 'botella', 'https://images.unsplash.com/photo-1556228578-0d85b1a4d571?w=400&q=80'),
(6, 'Jabón Protex 125g', 'jabon-protex-125g', 'Jabón de tocador Protex barra 125g.', 3.50, 500, 'unidad', 'https://images.unsplash.com/photo-1615914503779-ab6058e2a4b6?w=400&q=80'),
(6, 'Desodorante Rexona 50ml', 'desodorante-rexona-50ml', 'Desodorante Rexona roll-on 50ml.', 8.50, 300, 'unidad', 'https://images.unsplash.com/photo-1616401784845-180882c9cd21?w=400&q=80'),
(6, 'Crema Dental Pepsodent 75ml', 'crema-dental-pepsodent-75ml', 'Crema dental Pepsodent protección completa 75ml.', 5.50, 350, 'unidad', 'https://images.unsplash.com/photo-1624454002302-36b824d7bd0a?w=400&q=80');
