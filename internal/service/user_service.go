package service 

import (
	"time"

	"github.com/google/uuid"

	"tfg/internal/models"
)

type UserRepository interface {
	Create(user models.User) error
	GetByID(id string) (*models.User, error)
	GetAll() ([]models.User, error)
}

type UserService interface {
	Create(user models.User) (*models.User, error)
	GetByID(id string) (*models.User, error)
	GetAll() ([]models.User, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Create(user models.User) (*models.User, error) {
	user.ID = uuid.New().String()
	user.CreatedAt = time.Now()

	err := s.repo.Create(user)
	if err != nil{
		return nil, err
	}

	return &user, nil
}

func (s *userService) GetByID(id string) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
} 