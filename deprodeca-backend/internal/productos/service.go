package productos

import "context"

// Service contiene la lógica de negocio para el catálogo.
type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// ListarProductos obtiene productos paginados con filtros opcionales.
func (s *Service) ListarProductos(ctx context.Context, pagina, limite int, categoriaID int64, busqueda string) ([]Producto, error) {
	if pagina < 1 {
		pagina = 1
	}
	if limite < 1 || limite > 50 {
		limite = 20
	}

	return s.repo.ListarProductos(ctx, ListarParams{
		Pagina:      pagina,
		Limite:      limite,
		CategoriaID: categoriaID,
		Busqueda:    busqueda,
	})
}

// ObtenerProductoPorSlug busca un producto por su slug.
func (s *Service) ObtenerProductoPorSlug(ctx context.Context, slug string) (Producto, error) {
	return s.repo.ObtenerProductoPorSlug(ctx, slug)
}

// BuscarProductos busca productos por texto.
func (s *Service) BuscarProductos(ctx context.Context, query string, pagina, limite int) ([]Producto, error) {
	if pagina < 1 {
		pagina = 1
	}
	if limite < 1 || limite > 50 {
		limite = 20
	}

	return s.repo.BuscarProductos(ctx, query, pagina, limite)
}

// ListarCategorias devuelve todas las categorías activas.
func (s *Service) ListarCategorias(ctx context.Context) ([]Categoria, error) {
	return s.repo.ListarCategorias(ctx)
}

// ObtenerCategoriaPorSlug busca una categoría por su slug.
func (s *Service) ObtenerCategoriaPorSlug(ctx context.Context, slug string) (Categoria, error) {
	return s.repo.ObtenerCategoriaPorSlug(ctx, slug)
}

// CrearProducto crea un nuevo producto (solo admin).
func (s *Service) CrearProducto(ctx context.Context, input CrearProductoInput) (Producto, error) {
	return s.repo.CrearProducto(ctx, input)
}

// ActualizarProducto actualiza un producto existente (solo admin).
func (s *Service) ActualizarProducto(ctx context.Context, id int64, input CrearProductoInput) (Producto, error) {
	return s.repo.ActualizarProducto(ctx, id, input)
}

// EliminarProducto elimina un producto (solo admin).
func (s *Service) EliminarProducto(ctx context.Context, id int64) error {
	return s.repo.EliminarProducto(ctx, id)
}
