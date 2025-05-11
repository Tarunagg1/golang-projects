package service

import (
	"fmt"
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"
	"log"
)

type UserService struct{}

func (us *UserService) Signup(input dto.UserSignup) (string, error) {
	log.Println("Signing up user:", input)

	return "this is register", nil
}

func (us *UserService) findUserByEmail(email string) (domain.User, error) {
	fmt.Println("Finding user by email:", email)
	return domain.User{}, nil
}

func (us *UserService) Login(input any) (string, error) {
	_, err := us.findUserByEmail("tarun")

	if err != nil {
		return "", err
	}

	return "", nil
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
