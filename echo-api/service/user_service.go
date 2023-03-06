package user

import (
	"fmt"
	"github.com/labstack/echo/v4"

	"echo-api/db"
	"echo-api/entity"
)

// Service procides user's behavior
type Service struct{}

// User is alias of entity.User struct
type User entity.User

// GetAll is get all User
func (s Service) GetAll() ([]User, error) {
	db := db.GetDB()
	var u []User

	if err := db.Find(&u).Error; err != nil {
		fmt.Println("find error", err)
		return nil, err
	}

	return u, nil
}

// CreateModel is create User model
func (s Service) CreateModel(c echo.Context) (User, error) {
	db := db.GetDB()
	var u User

	if err := c.Bind(&u); err != nil {
		return u, err
	}

	if err := db.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// GetByID is get a User
func (s Service) GetByID(id string) (User, error) {
	db := db.GetDB()
	var u User

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// UpdateByID is update a User
func (s Service) UpdateByID(id string, c echo.Context) (User, error) {
	db := db.GetDB()
	var u User

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	if err := c.Bind(&u); err != nil {
		return u, err
	}

	db.Save(&u)

	return u, nil
}

// DeleteByID is delete a User
func (s Service) DeleteByID(id string) error {
	db := db.GetDB()
	var u User

	if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
		return err
	}

	return nil
}
