package handlers 

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"tfg/internal/models"
	"tfg/internal/service"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Create(c *gin.Context) {
	var user models.User

	//En caso de error al obtener el user via JSON mandamos al contexto bad request 
	// y le incluimos al gin.H error de invalid request body
	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	//Intentamos crear desde el service un usuario, en caso de que error no sea nulo mandamos un 
	// internal server error y adjuntamos gin.H que ha fallado la request
	createdUser, err := h.service.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create user"
		})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

func (h *UserHandler) GetByID(c *gin.Context) {
	//Obtenemos del gin.context el param "id"
	id := c.Param("id")

	user, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetAll(c *gin.Context) {
	users, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch users",
		})
		return
	}
	
	c.JSON(http.StatusOK, users)
}