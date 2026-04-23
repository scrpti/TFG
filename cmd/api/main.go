//cargar variables de entorno
// leer config
// conectar con PostgreSQL
// crear repositorios
// crear servicios
// crear handlers
// registrar rutas
// arrancar servidor

//Imports de paquetes estandar
package main

//Import de gin
//Import de paquetes locales

import (
	"log"
	"fmt"

	"github.com/gin-gonic/gin"

	"tfg/internal/config"
	"tfg/internal/db"

	"tfg/internal/repository"
	"tfg/internal/service"
	"tfg/internal/handlers"
	"tfg/internal/routes"
)

func main() {
	// Carga de configuracion

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	fmt.Println(cfg)
	// Carga de base de datos

	database, err := db.NewPostgresConnection(cfg.DB)
	if err != nil {
		log.Fatalf("error loading db: %v", err)
	}
	defer database.Close()

	// Inyeccion de las capsulas de Patient
	patientRepo := repository.NewPatientRepository(database)
	patientService := service.NewPatientService(patientRepo)
	patientHandler := handlers.NewPatientHandler(patientService)

	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)



	router := gin.Default()
	//Ruta de patient
	routes.RegisterRoutes(router, patientHandler, userHandler)

	//Arrancar el servidor

	log.Printf("server running on port %s", cfg.AppPort)
	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}