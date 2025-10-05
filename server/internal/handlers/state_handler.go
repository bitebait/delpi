package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"apiGo/internal/dto"
	"apiGo/internal/models"
	"apiGo/internal/services"
)

type StateHandler struct {
	stateService *services.StateService
}

func NewStateHandler(stateService *services.StateService) *StateHandler {
	return &StateHandler{stateService}
}

func (h *StateHandler) GetAll(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	sort := c.DefaultQuery("sort", "id asc")

	pagination, err := h.stateService.Find(limit, page, sort)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pagination)
}

func (h *StateHandler) Create(c *gin.Context) {
	var requestBody dto.CreateStateRequest

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _, body := range requestBody.Response {
		departamento := models.Departamento{
			Nombre: body.Nombre,
			GPS:    body.Gps,
			Dato:   body.Fact,
		}
		h.stateService.Save(&departamento)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Departamentos agregados con éxito"})

}
