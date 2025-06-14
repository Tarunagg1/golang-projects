package service

import (
	"fmt"
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/helper"
	"go-ecommerce-app/internal/repository"
	"log"
)

type UserService struct {
	Repo repository.UserRepository
	Auth helper.Auth
}

func (us *UserService) Signup(input dto.UserSignup) (string, error) {
	log.Println("Signing up user:", input)

	hPassword, err := us.Auth.CreateHashedPassword(input.Password)

	if err != nil {
		log.Println("Error hashing password:", err)
		return "", fmt.Errorf("error hashing password: %w", err)
	}

	// Check if user already exists
	existingUser, err := us.findUserByEmail(input.Email)

	if err == nil {
		log.Println("User already exists:", existingUser)
		return "", fmt.Errorf("user already exists")
	}

	user, err := us.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: hPassword,
		Phone:    input.Phone,
	})

	if err != nil {
		log.Println("Error creating user:", err)
		return "", err
	}

	return us.Auth.GenrateToken(user.ID, user.Email, "role")
}

func (us *UserService) findUserByEmail(email string) (*domain.User, error) {
	user, err := us.Repo.FindUser(email)

	return &user, err
}

func (us *UserService) Login(email string, password string) (string, error) {

	user, err := us.findUserByEmail(email)

	if err != nil {
		return "", err
	}

	// Check if password is correct
	err = us.Auth.VerifyPassword(password, user.Password)

	if err != nil {
		log.Println("Incorrect password for user:", email)
		return "", fmt.Errorf("incorrect password: %w", err)
	}

	return us.Auth.GenrateToken(user.ID, user.Email, "role")
}

func (us *UserService) GetVerificationCode(e domain.User) int {
	return 0
}

func (s UserService) VerifyCode(id uint, code string) error {
	return nil
}

func (s UserService) CreateProfile(id uint, input any) error {
	return nil
}

func (s UserService) GetProfile(id uint) (*domain.User, error) {
	return nil, nil
}

func (s UserService) UpdateProfile(id uint, input any) error {
	return nil
}

func (s UserService) BecomeSeller(id uint, input any) (string, error) {
	return "", nil
}

func (s UserService) FindCart(id uint) ([]interface{}, float64, error) {
	return nil, 0, nil
}

func (s UserService) CreateCart(input any, u any) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) CreateOrder(uId uint, orderRef string, pId string, amount float64) error {
	return nil
}

func (s UserService) GetOrders(u domain.User) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) GetOrderById(id uint, uId uint) (interface{}, error) {
	return nil, nil
}
