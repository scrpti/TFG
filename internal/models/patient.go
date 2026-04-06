package models

import "time"

type Patient struct {
	ID string ´json:"id"´
	FullName string ´json:"full_name"´
	Identifier string ´json:"identifier"´
	CreatedAt time.Time ´json:"created_at"´
}

