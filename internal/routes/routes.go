package routes

import (
	"github.com/gin-gonic/gin"
	"tfg/internal/handlers"
)

func RegisterRoutes(router *gin.Engine, patientHandler *handlers.PatientHandler, 
					userHandler *handlers.UserHandler, documentHandler *handlers.DocumentHandler) {
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

	// Rutas de users

	router.POST("/users", userHandler.Create)
	router.GET("/users", userHandler.GetAll)
	router.GET("/users/:id", userHandler.GetByID)

	// Rutas de documents

	router.POST("/documents", documentHandler.Create)
	router.GET("/documents", documentHandler.GetAll)
	router.GET("/documents/:id", documentHandler.GetByID)
	router.GET("/documents/patient/:id", documentHandler.GetByPatientID)
	router.GET("/documents/:id/verify", documentHandler.Verify)
}