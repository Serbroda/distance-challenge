package services

import (
	"github.com/Serbroda/distance-challenge/models"
	"gorm.io/gorm"
)

type RunService struct {
	DB *gorm.DB
}

func (s *RunService) Get(id uint) (models.Run, error) {
	var entity models.Run
	res := s.DB.First(&entity, id)
	if res.Error != nil {
		return models.Run{}, res.Error
	}
	return entity, nil
}

func (s *RunService) GetAll() []models.Run {
	var entites []models.Run
	s.DB.Find(&entites)
	return entites
}

func (s *RunService) GetByUser(userId uint) ([]models.Run, error) {
	var entities []models.Run
	res := s.DB.Where("user_id = ?", userId).Find(&entities)
	if res.Error != nil {
		return []models.Run{}, res.Error
	}
	return entities, nil
}

func (s *RunService) Create(run models.Run) (models.Run, error) {
	res := s.DB.Create(&run)
	if res.Error != nil {
		return models.Run{}, res.Error
	}
	return run, nil
}

func (s *RunService) Update(run models.Run) (models.Run, error) {
	res := s.DB.Model(&run).Updates(run)
	if res.Error != nil {
		return models.Run{}, res.Error
	}
	return s.Get(run.ID)
}

func (s *RunService) DeleteRun(id uint) error {
	res := s.DB.Delete(&models.Run{}, id)
	return res.Error
}
