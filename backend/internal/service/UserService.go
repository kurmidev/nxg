package service

import (
	"errors"
	"fmt"
	"nxg/internal/domain"
	"nxg/internal/dto"
	"nxg/internal/helper"
	"nxg/internal/repository"
)

type UserService struct {
	Repo repository.UserRepository
	Auth helper.Auth
}

func (s UserService) FindUserByEmail(email string) (*domain.User, error) {
	user, err := s.Repo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s UserService) SignUp(input dto.UserSignup) (string, error) {
	fmt.Printf("create user %v", input)

	hashPassword, err := s.Auth.CreateHashedPassword(input.Password)
	if err != nil {
		return "", err
	}
	user, err := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: hashPassword,
		UserType: input.UserType,
	})

	if err != nil {
		return "", err
	}

	fmt.Println("User signed up", user)
	userinfo := fmt.Sprintf("%v, %v,%v", user.ID, user.Email, user.UserType)
	return userinfo, nil
}

func (s UserService) Login(email string, password string) (string, error) {
	user, err := s.FindUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	fmt.Printf("%+v", user)
	err = s.Auth.VerifyPassword(password, user.Password)
	if err != nil {
		return "", errors.New("invalid password")
	}
	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (s UserService) GetVerificationCode(input any) (int, error) {
	return 0, nil
}

func (s UserService) VerifyCode(id uint, code int) (string, error) {
	return "", nil
}

func (s UserService) GetProfile(id uint) (*domain.User, error) {
	return nil, nil
}

func (s UserService) CreateProfile(id uint, input any) error {
	return nil
}

func (s UserService) UpdateProfile(id uint, input any) error {
	return nil
}
