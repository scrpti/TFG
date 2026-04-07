package service

import (
	"time"

	"github.com/google/uuid"

	"tfg/internal/models"
)

type PatientRepository interface {
	Create(patient models.Patient) error
	GetByID(id string) (*models.Patient, error)
	GetAll() ([]models.Patient, error)
}

type PatientService interface {
	Create(patient models.Patient) (*models.Patient, error)
	GetByID(id string) (*models.Patient, error)
	GetAll() ([]models.Patient, error)
}

type patientService struct {
	repo PatientRepository
}

func NewPatientService(repo PatientRepository) PatientService {
	return &patientService{
		repo: repo,
	}
}

func (s *patientService) Create(patient models.Patient) (*models.Patient, error) {
	patient.ID = uuid.New().String()
	patient.CreatedAt = time.Now()

	err := s.repo.Create(patient)
	if err != nil {
		return nil, err
	}

	return &patient, nil
}

func (s *patientService) GetByID(id string) (*models.Patient, error) {
	return s.repo.GetByID(id)
}

func (s *patientService) GetAll() ([]models.Patient, error) {
	return s.repo.GetAll()
}