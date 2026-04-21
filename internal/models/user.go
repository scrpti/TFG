package models

import "time"

type User struct {
	ID string `json:"id"`
	FullName string `json:"full_name"`
	Role string `json:"role"`
	HospitalID string `json:"hospital_id"`
	CreatedAt time.Time `json:"created_at"`
}