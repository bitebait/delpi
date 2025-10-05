package services

import (
	"gorm.io/gorm"

	"apiGo/internal/models"
)

type StateService struct {
	db *gorm.DB
}

func NewStateService(db *gorm.DB) *StateService {
	return &StateService{db}
}

func (s *StateService) Find(limit int, page int, sort string) (models.Pagination, error) {
	var totalRows int64
	if err := s.db.Model(&models.Departamento{}).Count(&totalRows).Error; err != nil {
		return models.Pagination{}, err
	}

	offset := (page - 1) * limit

	var departamentos []models.Departamento
	if err := s.db.Order(sort).Limit(limit).Offset(offset).Find(&departamentos).Error; err != nil {
		return models.Pagination{}, err
	}

	totalPages := int((totalRows + int64(limit) - 1) / int64(limit))

	return models.Pagination{
		Limit:      limit,
		Page:       page,
		Sort:       sort,
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Data:       departamentos,
	}, nil
}

func (s *StateService) Save(state *models.Departamento) {
	s.db.Create(state)
}
