package auth

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"deprodeca-backend/pkg/hash"
	pkgjwt "deprodeca-backend/pkg/jwt"

	"github.com/jackc/pgx/v5"
)

// Service contiene la lógica de negocio para autenticación.
// Recibe una interfaz Repository, NO una implementación concreta (DIP de SOLID).
type Service struct {
	repo         Repository
	jwtSecret    string
	jwtExpHoras  int
}

// NewService crea el servicio de auth con sus dependencias inyectadas.
func NewService(repo Repository, jwtSecret string, jwtExpHoras int) *Service {
	return &Service{
		repo:         repo,
		jwtSecret:    jwtSecret,
		jwtExpHoras:  jwtExpHoras,
	}
}

// Registrar crea un nuevo usuario proveedor y devuelve el token JWT.
func (s *Service) Registrar(ctx context.Context, input RegistroInput) (AuthResponse, error) {
	// Normalizar email
	input.Email = strings.ToLower(strings.TrimSpace(input.Email))

	// Validar campos obligatorios
	if err := validarRegistro(input); err != nil {
		return AuthResponse{}, err
	}

	// Verificar si el email ya existe
	_, err := s.repo.ObtenerUsuarioPorEmail(ctx, input.Email)
	if err == nil {
		return AuthResponse{}, errors.New("el email ya está registrado")
	}

	// Hashear contraseña
	passwordHash, err := hash.Hashear(input.Password)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("error al hashear contraseña: %w", err)
	}

	// Crear usuario
	usuario, err := s.repo.CrearUsuario(ctx, input, passwordHash)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("error al crear usuario: %w", err)
	}

	// Generar JWT
	return s.generarRespuesta(usuario)
}

// Login autentica a un usuario con email y contraseña.
func (s *Service) Login(ctx context.Context, input LoginInput) (AuthResponse, error) {
	input.Email = strings.ToLower(strings.TrimSpace(input.Email))

	if input.Email == "" || input.Password == "" {
		return AuthResponse{}, errors.New("email y contraseña son obligatorios")
	}

	// Buscar usuario
	usuario, err := s.repo.ObtenerUsuarioPorEmail(ctx, input.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return AuthResponse{}, errors.New("email o contraseña incorrectos")
		}
		return AuthResponse{}, fmt.Errorf("error al buscar usuario: %w", err)
	}

	// Verificar contraseña
	if !hash.Verificar(input.Password, usuario.PasswordHash) {
		return AuthResponse{}, errors.New("email o contraseña incorrectos")
	}

	return s.generarRespuesta(usuario)
}

// generarRespuesta crea el token JWT y construye la AuthResponse.
func (s *Service) generarRespuesta(usuario Usuario) (AuthResponse, error) {
	token, expira, err := pkgjwt.GenerarToken(s.jwtSecret, s.jwtExpHoras, usuario.ID, usuario.Email, usuario.Rol)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("error al generar token: %w", err)
	}

	return AuthResponse{
		Token:     token,
		ExpiraEn:  expira.Format("2006-01-02T15:04:05Z07:00"),
		UsuarioID: usuario.ID,
		Email:     usuario.Email,
		Nombre:    usuario.Nombre,
		Rol:       usuario.Rol,
	}, nil
}

func validarRegistro(input RegistroInput) error {
	if strings.TrimSpace(input.Email) == "" {
		return errors.New("email es obligatorio")
	}
	if len(input.Password) < 6 {
		return errors.New("la contraseña debe tener al menos 6 caracteres")
	}
	if strings.TrimSpace(input.Nombre) == "" {
		return errors.New("nombre es obligatorio")
	}
	if strings.TrimSpace(input.Empresa) == "" {
		return errors.New("empresa es obligatorio")
	}
	if strings.TrimSpace(input.RUC) == "" {
		return errors.New("RUC es obligatorio")
	}
	return nil
}
