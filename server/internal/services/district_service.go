package services

import (
	"gorm.io/gorm"

	"apiGo/internal/models"
)

type DistrictService struct {
	db *gorm.DB
}

func NewDistrictService(db *gorm.DB) *DistrictService {
	return &DistrictService{db}
}

func (s *DistrictService) Find(limit int, page int, sort string) (models.Pagination, error) {
	var totalRows int64
	if err := s.db.Model(&models.Barrio{}).Count(&totalRows).Error; err != nil {
		return models.Pagination{}, err
	}

	offset := (page - 1) * limit

	var barrios []models.Barrio
	if err := s.db.Preload("Ciudad").Order(sort).Limit(limit).Offset(offset).Find(&barrios).Error; err != nil {
		return models.Pagination{}, err
	}

	totalPages := int((totalRows + int64(limit) - 1) / int64(limit))

	return models.Pagination{
		Limit:      limit,
		Page:       page,
		Sort:       sort,
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Data:       barrios,
	}, nil
}

func (s *DistrictService) FindByCityID(id int) []models.BarrioOnly {
	var barrios []models.BarrioOnly
	s.db.Where("ciudad_id = ?", id).Preload("Ciudad").Find(&barrios)
	return barrios
}

func (s *DistrictService) Save(district *models.Barrio) {
	s.db.Create(district)
}
