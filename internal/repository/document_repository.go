package repository

import (
	"database/sql"
	"tfg/internal/models"
)

type DocumentRepository struct {
	db *sql.DB
}

func NewDocumentRepository(db *sql.DB) *DocumentRepository {
	return &DocumentRepository{
		db: db,
	}
}

//Funcion Create

func (r *DocumentRepository) Create(document models.Document) error {
	query := `
		INSERT INTO documents (id, patient_id, doctor_id, hospital_id, document_type, file_name, file_path, hash, status, uploaded_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err := r.db.Exec(
		query,
		document.ID,
		document.PatientID,
		document.DoctorID,
		document.HospitalID,
		document.DocumentType,
		document.FileName,
		document.FilePath,
		document.Hash,
		document.Status,
		document.UploadedAt,
	)
	return err
}

//Funcion GetByID

func (r *DocumentRepository) GetByID(id string) (*models.Document, error) {
	query := `
		SELECT id, patient_id, doctor_id, hospital_id, document_type, file_name, file_path, hash, status, uploaded_at
		FROM documents
		WHERE id = $1
	`

	row := r.db.QueryRow(query, id)

	var document models.Document

	err := row.Scan(
		&document.ID,
		&document.PatientID,
		&document.DoctorID,
		&document.HospitalID,
		&document.DocumentType,
		&document.FileName,
		&document.FilePath,
		&document.Hash,
		&document.Status,
		&document.UploadedAt,
	)

	if err != nil {
		return nil, err
	}

	return &document, nil
}

//Funcion GetAll
 
func (r *DocumentRepository) GetAll() ([]models.Document, error){
	query :=`
		SELECT id, patient_id, doctor_id, hospital_id, document_type, file_name, file_path, hash, status, uploaded_at
		FROM documents
		ORDER BY uploaded_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var documents []models.Document

	for rows.Next(){
		var document models.Document

		err := rows.Scan(
			&document.ID,
			&document.PatientID,
			&document.DoctorID,
			&document.HospitalID,
			&document.DocumentType,
			&document.FileName,
			&document.FilePath,
			&document.Hash,
			&document.Status,
			&document.UploadedAt,
		)
		if err != nil {
			return nil, err
		}

		documents = append(documents, document)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return documents, nil
}

//Funcion GetByPatientID

func (r *DocumentRepository) GetByPatientID(patientID string) ([]models.Document, error) {
	query :=`
		SELECT id, patient_id, doctor_id, hospital_id, document_type, file_name, file_path, hash, status, uploaded_at
		FROM documents
		WHERE patient_id = $1
		ORDER BY uploaded_at DESC
	`

	rows, err := r.db.Query(query, patientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var documents []models.Document

	for rows.Next(){
		var document models.Document

		err := rows.Scan(
			&document.ID,
			&document.PatientID,
			&document.DoctorID,
			&document.HospitalID,
			&document.DocumentType,
			&document.FileName,
			&document.FilePath,
			&document.Hash,
			&document.Status,
			&document.UploadedAt,
		)
		if err != nil {
			return nil, err
		}

		documents = append(documents, document)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return documents, nil
}