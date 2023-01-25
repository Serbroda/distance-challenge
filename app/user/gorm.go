package user

import (
	"errors"
	"gorm.io/gorm"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)

type UserServiceGorm struct {
	DB *gorm.DB
}

func (s *UserServiceGorm) Get(id uint) (User, error) {
	var entity User
	res := s.DB.First(&entity, id)
	if res.Error != nil {
		return User{}, res.Error
	}
	return entity, nil
}

func (s *UserServiceGorm) GetByUsername(username string) (User, error) {
	var entity User
	res := s.DB.Where("lower(username) = lower(?)", username).First(&entity)
	if res.Error != nil {
		return User{}, res.Error
	}
	return entity, nil
}

func (s *UserServiceGorm) Create(user User) (User, error) {
	if _, err := s.GetByUsername(user.Username); err != nil {
		return User{}, ErrUserAlreadyExists
	}
	res := s.DB.Create(&user)
	if res.Error != nil {
		return User{}, res.Error
	}
	return user, nil
}

func (s *UserServiceGorm) Update(user User) (User, error) {
	res := s.DB.Model(&user).Updates(user)
	if res.Error != nil {
		return User{}, res.Error
	}
	return s.Get(user.ID)
}
