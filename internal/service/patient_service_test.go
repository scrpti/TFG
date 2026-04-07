package service

import (
	"testing"

	"tfg/internal/models"
)

type mockPatientRepository struct {
	createFn  func(patient models.Patient) error
	getByIDFn func(id string) (*models.Patient, error)
	getAllFn  func() ([]models.Patient, error)
	updateFn  func(patient models.Patient) error
	deleteFn  func(id string) error
}

func (m *mockPatientRepository) Create(patient models.Patient) error {
	return m.createFn(patient)
}

func (m *mockPatientRepository) GetByID(id string) (*models.Patient, error) {
	return m.getByIDFn(id)
}

func (m *mockPatientRepository) GetAll() ([]models.Patient, error) {
	return m.getAllFn()
}

func (m *mockPatientRepository) Update(patient models.Patient) error {
	return m.updateFn(patient)
}

func (m *mockPatientRepository) Delete(id string) error {
	return m.deleteFn(id)
}

func TestPatientService_Create_AssignsIDAndCreatedAt(t *testing.T) {
	mockRepo := &mockPatientRepository{
		createFn: func(patient models.Patient) error {
			if patient.ID == "" {
				t.Fatal("expected patient ID to be generated")
			}
			if patient.CreatedAt.IsZero() {
				t.Fatal("expected CreatedAt to be assigned")
			}
			return nil
		},
		getByIDFn: func(id string) (*models.Patient, error) {
			return nil, nil
		},
		getAllFn: func() ([]models.Patient, error) {
			return nil, nil
		},
		updateFn: func(patient models.Patient) error {
			return nil
		},
		deleteFn: func(id string) error {
			return nil
		},
	}

	service := NewPatientService(mockRepo)

	input := models.Patient{
		FullName:   "Juan Perez",
		Identifier: "PAT-001",
	}

	result, err := service.Create(input)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.ID == "" {
		t.Fatal("expected returned patient to have ID")
	}

	if result.CreatedAt.IsZero() {
		t.Fatal("expected returned patient to have CreatedAt")
	}
}