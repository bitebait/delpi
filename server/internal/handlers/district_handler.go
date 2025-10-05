package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"apiGo/internal/dto"
	"apiGo/internal/models"
	"apiGo/internal/services"
)

type DistrictHandler struct {
	districtService *services.DistrictService
}

func NewDistrictHandler(districtService *services.DistrictService) *DistrictHandler {
	return &DistrictHandler{districtService}
}

func (h *DistrictHandler) GetAll(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	sort := c.DefaultQuery("sort", "id asc")

	pagination, err := h.districtService.Find(limit, page, sort)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pagination)
}

func (h *DistrictHandler) GetAllByCity(c *gin.Context) {
	id := c.Param("ciudadID")

	cityID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de ciudad inválido"})
		return
	}

	barrios := h.districtService.FindByCityID(cityID)

	c.JSON(http.StatusOK, barrios)
}

func (h *DistrictHandler) Create(c *gin.Context) {
	id := c.Param("ciudadID")

	cityID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de ciudad inválido"})
		return
	}

	var requestBody dto.CreateStateRequest
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, body := range requestBody.Response {
		district := &models.Barrio{
			Nombre:   body.Nombre,
			CiudadID: cityID,
		}

		h.districtService.Save(district)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Barrios agregados con éxito"})
}
