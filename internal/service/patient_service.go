package service

import (
	"tfg/internal/models"
	"tfg/internal/repository"
	"time"
	"github.com/google/uuid"
)

type PatientService struct {
	repo *repository.PatientRepository
}

func NewPatientService(repo *repository.PatientRepository) *PatientService {
	return &PatientService{
		repo: repo,
	}
}

func (s *PatientService) Create(patient models.Patient) (*models.Patient, error) {
	patient.ID = uuid.New().String()
	patient.CreatedAt = time.now()

	err := s.repo.Create(patient)
	if err != nil {
		return nil, err
	}

	return &patient, nil
}

func (s *PatientService) GetByID(id string) (*models.Patient, error) {
	return s.repo.GetByID(id)
}

func (s *PatientService) GetAll() ([]models.Patient, error) {
	return s.repo.getAll()
}