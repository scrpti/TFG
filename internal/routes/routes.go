package routes

import (
	"github.com/gin-gonic/gin"
	"tfg/internal/handlers"
)

func RegisterRoutes(router *gin.Engine, patientHandler *handlers.PatientHandler) {
	// Ruta de health (/health)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status" : "ok",
		})
	})

	// Rutas de patients

	router.POST("/patients", patientHandler.Create)
	router.GET("/patients", patientHandler.GetAll)
	router.GET("/patients/:id", patientHandler.GetByID)
}