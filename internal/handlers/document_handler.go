package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"tfg/internal/models"
	"tfg/internal/service"
)

type DocumentHandler struct {
	service service.DocumentService
}

func NewDocumentHandler(service service.DocumentService) *DocumentHandler {
	return &DocumentHandler{
		service: service,
	}
}

func (h *DocumentHandler) Create(c *gin.Context) {
	var document models.Document

	if err := c.ShouldBindJSON(&document); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	createdDocument, err := h.service.Create(document)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : "failed to create document",
		})
		return
	}

	c.JSON(http.StatusCreated, createdDocument)
}

func (h *DocumentHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	document, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "document not found",
		})
		return
	}

	c.JSON(http.StatusOK, document)
}

func (h *DocumentHandler) GetAll(c *gin.Context) {
	documents, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch documents",
		})
		return
	}
	c.JSON(http.StatusOK, documents)
}

func (h *DocumentHandler) GetByPatientID(c *gin.Context) {
	id := c.Param("patient_id")

	documents, err := h.service.GetByPatientID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch documents",
		})
		return
	}
	c.JSON(http.StatusOK, documents)
}