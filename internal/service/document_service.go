package service

import (
	"os"
	"time"
	"mime/multipart"
	"io"
	"path/filepath"

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
	Create(document models.Document, file *multipart.FileHeader) (*models.Document, error)
	GetByID(id string) (*models.Document, error)
	GetAll() ([]models.Document, error)
	GetByPatientID(id string) ([]models.Document, error)

}

type documentService struct{
	repo DocumentRepository
	uploadDir string
}

func NewDocumentService(repo DocumentRepository, uploadDir string) DocumentService {
	return &documentService{
		repo: repo,
		uploadDir: uploadDir,
	}
}

func (s *documentService) Create(document models.Document, file *multipart.FileHeader) (*models.Document, error) {
	document.ID = uuid.New().String()
	document.UploadedAt = time.Now()
	document.Status = "ACTIVE"

	filePath := filepath.Join(s.uploadDir, document.ID+"_"+file.Filename)

	if err := os.MkdirAll(s.uploadDir, os.ModePerm); err != nil {
		return nil, err
	}

	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return nil, err
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	document.FilePath = filePath
	document.FileName = file.Filename
	document.Hash = crypto.CalculateSHA256(data)

	if err := s.repo.Create(document); err != nil {
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