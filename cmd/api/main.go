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

	router := gin.Default()

	// Ruta de health (/health)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status" : "ok",
		})
	})

	//Arrancar el servidor

	log.Printf("server running on port %s", cfg.AppPort)
	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}