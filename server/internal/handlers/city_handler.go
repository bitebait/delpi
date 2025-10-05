package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"apiGo/internal/dto"
	"apiGo/internal/models"
	"apiGo/internal/services"
)

type CityHandler struct {
	cityService *services.CityService
}

func NewCityHandler(cityService *services.CityService) *CityHandler {
	return &CityHandler{cityService}
}

func (h *CityHandler) GetAll(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	sort := c.DefaultQuery("sort", "id asc")

	pagination, err := h.cityService.Find(limit, page, sort)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pagination)

}

func (h *CityHandler) GetAllByState(c *gin.Context) {
	stateID, err := strconv.Atoi(c.Param("departamentoID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de estado inválido"})
		return
	}

	cities := h.cityService.FindByStateID(stateID)
	c.JSON(http.StatusOK, cities)
}

func (h *CityHandler) Create(c *gin.Context) {
	id := c.Param("departamentoID")

	departamentoID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de estado inválido"})
		return
	}

	var requestBody dto.CreateStateRequest
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, body := range requestBody.Response {
		ciudad := &models.Ciudad{
			Nombre:         body.Nombre,
			DepartamentoID: departamentoID,
		}
		h.cityService.Save(ciudad)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ciudades agregadas con éxito"})
}
