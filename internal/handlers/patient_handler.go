package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"tfg/internal/models"
	"tfg/internal/service"
)

type PatientHandler struct {
	service service.PatientService
}

func NewPatientHandler(service service.PatientService) *PatientHandler {
	return &PatientHandler{
		service: service,
	}
}

func (h *PatientHandler) Create(c *gin.Context) {
	var patient models.Patient

	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	createdPatient, err := h.service.Create(patient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create patient",
		})
		return
	}

	c.JSON(http.StatusCreated, createdPatient)
}

func (h *PatientHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	patient, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "patient not found",
		})
		return
	}

	c.JSON(http.StatusOK, patient)
}

func (h *PatientHandler) GetAll(c *gin.Context) {
	patients, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch patients",
		})
		return
	}

	c.JSON(http.StatusOK, patients)
}