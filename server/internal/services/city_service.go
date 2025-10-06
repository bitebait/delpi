package services

import (
	"gorm.io/gorm"

	"apiGo/internal/models"
)

type CityService struct {
	db *gorm.DB
}

func NewCityService(db *gorm.DB) *CityService {
	return &CityService{db}
}

func (s *CityService) Find(limit int, page int, sort string) (models.Pagination, error) {
	var totalRows int64
	if err := s.db.Model(&models.Ciudad{}).Count(&totalRows).Error; err != nil {
		return models.Pagination{}, err
	}

	offset := (page - 1) * limit

	var ciudades []models.Ciudad
	if err := s.db.Preload("Departamento").Order(sort).Limit(limit).Offset(offset).Find(&ciudades).Error; err != nil {
		return models.Pagination{}, err
	}

	totalPages := int((totalRows + int64(limit) - 1) / int64(limit))

	return models.Pagination{
		Limit:      limit,
		Page:       page,
		Sort:       sort,
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Data:       ciudades,
	}, nil
}

func (s *CityService) FindByStateID(id int) []models.CiudadOnly {
	var cities []models.CiudadOnly
	s.db.Where("departamento_id = ?", id).Preload("Deparamento").Find(&cities)
	return cities
}

func (s *CityService) Save(city *models.Ciudad) {
	s.db.Create(city)
}
