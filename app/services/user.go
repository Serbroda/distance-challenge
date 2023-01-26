package services

import (
	"errors"
	"github.com/Serbroda/distance-challenge/models"
	"gorm.io/gorm"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)

type UserService struct {
	DB *gorm.DB
}

func (s *UserService) Get(id uint) (models.User, error) {
	var entity models.User
	res := s.DB.First(&entity, id)
	if res.Error != nil {
		return models.User{}, res.Error
	}
	return entity, nil
}

func (s *UserService) GetAll() []models.User {
	var entites []models.User
	s.DB.Find(&entites)
	return entites
}

func (s *UserService) GetByUsername(username string) (models.User, error) {
	var entity models.User
	res := s.DB.Where("lower(username) = lower(?)", username).First(&entity)
	if res.Error != nil {
		return models.User{}, res.Error
	}
	return entity, nil
}

func (s *UserService) ExistsByUsername(username string) bool {
	_, err := s.GetByUsername(username)
	return err == nil
}

func (s *UserService) Create(user models.User) (models.User, error) {
	if s.ExistsByUsername(user.Username) {
		return models.User{}, ErrUserAlreadyExists
	}
	res := s.DB.Create(&user)
	if res.Error != nil {
		return models.User{}, res.Error
	}
	return user, nil
}

func (s *UserService) Update(user models.User) (models.User, error) {
	res := s.DB.Model(&user).Updates(user)
	if res.Error != nil {
		return models.User{}, res.Error
	}
	return s.Get(user.ID)
}
