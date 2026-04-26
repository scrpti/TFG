package models

import "time"

type Document struct {
	ID string `json:"id"`
	PatientID string `json:"patient_id"`
	DoctorID string `json:"doctor_id"`
	HospitalID string `json:"hospital_id"`
	DocumentType string `json:"document_type"`
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
	Hash string `json:"hash"`
	Status string `json:"status"`
	UploadedAt time.Time `json:"uploaded_at"`
}
