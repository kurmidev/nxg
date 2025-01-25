package repository

import (
	"errors"
	"fmt"
	"log"
	"nxg/internal/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateUser(u domain.User) (domain.User, error)
	FindUserByEmail(email string) (domain.User, error)
	FindUserById(id string) (domain.User, error)
	UpdateUser(id int, u domain.User) (domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u userRepository) CreateUser(user domain.User) (domain.User, error) {
	fmt.Printf("user details is %v\n", user)
	err := u.db.Create(&user).Error
	if err != nil {
		log.Printf("create user error %v", err)
		return domain.User{}, errors.New("failed to create user")
	}
	return user, nil
}

func (u userRepository) FindUserByEmail(email string) (domain.User, error) {
	var user domain.User
	err := u.db.First(&user, "email=?", email).Error
	if err != nil {
		log.Printf("find user by email error %v", err)
		return domain.User{}, errors.New("failed to find user by email")
	}
	return user, nil
}
func (u userRepository) FindUserById(id string) (domain.User, error) {
	var user domain.User
	err := u.db.First(&user, id).Error
	if err != nil {
		log.Printf("find user by id error %v", err)
		return domain.User{}, errors.New("failed to find user by id")
	}
	return user, nil
}
func (u userRepository) UpdateUser(id int, usr domain.User) (domain.User, error) {
	var user domain.User
	err := u.db.Model(&user).Clauses(clause.Returning{}).Where("id=?", id).Updates(usr).Error
	if err != nil {
		log.Printf("update user error %v", err)
		return domain.User{}, errors.New("failed to update user")
	}
	return user, nil
}
