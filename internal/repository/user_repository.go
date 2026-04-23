package repository

import (
	"database/sql"
	"tfg/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Funcion Create

func (r *UserRepository) Create(user models.User) error {
	query := `
		INSERT INTO users(id, full_name, role, hospital_id, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(
		query,
		user.ID,
		user.FullName,
		user.Role,
		user.HospitalID,
		user.CreatedAt,
	)
	return err
}

//Funcion GetByID

func (r *UserRepository) GetByID(id string) (*models.User, error) {
	query := `
		SELECT id, full_name, role, hospital_id, created_at
		FROM users
		WHERE id = $1
	`

	row := r.db.QueryRow(query, id)

	var user models.User

	err := row.Scan(
		&user.ID,
		&user.FullName,
		&user.Role,
		&user.HospitalID,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

//Funcion GetAll

func (r *UserRepository) GetAll() ([]models.User, error){
	query := `
		SELECT id, full_name, role, hospital_id, created_at
		FROM users 
		ORDER BY created_at DESC
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var users []models.User

	for rows.Next(){
		var user models.User

		err := rows.Scan(
			&user.ID,
			&user.FullName,
			&user.Role,
			&user.HospitalID,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}