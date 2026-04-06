package repository

import (
	"database/sql"
	"tfg/internal/models"
)

type PatientRepository struct {
	db *sql.DB
}

func NewPatientRepository(db *sql.DB) *PatientRepository {
	return &PatientRepository{
		db: db
	}
}

//Funcion Create

func (r *PatientRepository) Create(patient models.Patient) error {
	query := `
		INSERT INTO patients (id, full_name, identifier, created_at)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.Exec(
		query,
		patient.ID,
		patient.FullName,
		patient.Identifier,
		patient.CreatedAt
	)
}

//Funcion GetByID

func (r *PatientRepository) GetByID(id string) (*models.Patient, error){
	query := `
		SELECT id, full_name, identifier, created_at
		FROM patients
		WHERE id = $1
	`

	row := r.db.QueryRow(query, id)

	var patient models.Patient

	err := row.Scan(
		&patient.ID,
		&patient.FullName,
		&patient.Identifier,
		&patient.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &patient, nil
}
