package service

import (
	"os"
	"time"

	"github.com/google/uuid"

	"tfg/internal/models"
	"tfg/internal/crypto"
)

type DocumentRepository interface {
	Create(document models.Document) error
	GetByID(id string) (*models.Document, error)
	GetAll() ([]models.Document, error)
	GetByPatientID(id string) ([]models.Document, error)

}

type DocumentService interface {
	Create(document models.Document) (*models.Document, error)
	GetByID(id string) (*models.Document, error)
	GetAll() ([]models.Document, error)
	GetByPatientID(id string) ([]models.Document, error)

}

type documentService struct{
	repo DocumentRepository
}

func NewDocumentService(repo DocumentRepository) DocumentService {
	return &documentService{
		repo: repo,
	}
}

func (s *documentService) Create(document models.Document) (*models.Document, error) {
	data, err := os.ReadFile(document.FilePath)
	if err != nil {
		return nil, err
	}
	
	document.ID = uuid.New().String()
	document.UploadedAt = time.Now()
	document.Hash = crypto.CalculateSHA256(data)
	document.Status = "Active"

	err = s.repo.Create(document)
	if err != nil {
		return nil, err
	}

	return &document, nil
}

func (s *documentService) GetByID(id string) (*models.Document, error) {
	return s.repo.GetByID(id)
}

func (s *documentService) GetAll() ([]models.Document, error) {
	return s.repo.GetAll()
}

func (s *documentService) GetByPatientID(id string) ([]models.Document, error) {
	return s.repo.GetByPatientID(id)
}