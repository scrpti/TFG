package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"

	"tfg/internal/models"
)

type mockPatientService struct {
	createFn  func(patient models.Patient) (*models.Patient, error)
	getByIDFn func(id string) (*models.Patient, error)
	getAllFn  func() ([]models.Patient, error)
}

func (m *mockPatientService) Create(patient models.Patient) (*models.Patient, error) {
	return m.createFn(patient)
}

func (m *mockPatientService) GetByID(id string) (*models.Patient, error) {
	return m.getByIDFn(id)
}

func (m *mockPatientService) GetAll() ([]models.Patient, error) {
	return m.getAllFn()
}

func TestPatientHandler_Create_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &mockPatientService{
		createFn: func(patient models.Patient) (*models.Patient, error) {
			patient.ID = "test-id"
			patient.CreatedAt = time.Now()
			return &patient, nil
		},
		getByIDFn: func(id string) (*models.Patient, error) {
			return nil, errors.New("not implemented")
		},
		getAllFn: func() ([]models.Patient, error) {
			return nil, errors.New("not implemented")
		},
	}

	handler := NewPatientHandler(mockService)

	router := gin.Default()
	router.POST("/patients", handler.Create)

	body := map[string]string{
		"full_name":  "Juan Perez",
		"identifier": "PAT-001",
	}

	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest(http.MethodPost, "/patients", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, rec.Code)
	}
}

func TestPatientHandler_Create_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &mockPatientService{
		createFn: func(patient models.Patient) (*models.Patient, error) {
			return nil, nil
		},
		getByIDFn: func(id string) (*models.Patient, error) {
			return nil, errors.New("not implemented")
		},
		getAllFn: func() ([]models.Patient, error) {
			return nil, errors.New("not implemented")
		},
	}

	handler := NewPatientHandler(mockService)

	router := gin.Default()
	router.POST("/patients", handler.Create)

	req, _ := http.NewRequest(http.MethodPost, "/patients", bytes.NewBufferString(`{invalid-json}`))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}