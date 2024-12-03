package service

import (
	"fmt"
	"strconv"

	"area/database"
	"area/repository"
	"area/schemas"
)

type UserService interface {
	Login(user schemas.User) (JWTtoken string, err error)
	Register(newUser schemas.User) (JWTtoken string, userId uint64, err error)
	GetUserInfo(token string) (userInfo schemas.User, err error)
	UpdateUserInfo(newUser schemas.User) (err error)
	GetUserById(userId uint64) schemas.User
}

type userService struct {
	authorizedUsername string
	authorizedPassword string
	repository         repository.UserRepository
	serviceJWT         JWTService
}

func NewUserService(userRepository repository.UserRepository, serviceJWT JWTService) UserService {
	return &userService{
		authorizedUsername: "root",
		authorizedPassword: "password",
		repository:         userRepository,
		serviceJWT:         serviceJWT,
	}
}

func (service *userService) Login(newUser schemas.User) (JWTtoken string, err error) {
	userWiththisUserName := service.repository.FindByUserName(newUser.Username)
	if len(userWiththisUserName) == 0 {
		return "", fmt.Errorf("invalid credentials")
	}
	// regular user
	for _, user := range userWiththisUserName {
		if database.DoPasswordsMatch(user.Password, newUser.Password) {
			return service.serviceJWT.GenerateToken(
				strconv.FormatUint(user.Id, 10),
				user.Username,
				false,
			), nil
		}
	}

	// Oauth2.0 user
	for _, user := range userWiththisUserName {
		if user.Email == newUser.Email {
			if newUser.TokenId != 0 {
				return service.serviceJWT.GenerateToken(
					strconv.FormatUint(user.Id, 10),
					user.Username,
					false,
				), nil
			}
		}
	}

	return "", fmt.Errorf("invalid credentials")
}

func (service *userService) Register(newUser schemas.User) (JWTtoken string, userId uint64, err error) {
	userWiththisEmail := service.repository.FindByEmail(newUser.Email)
	if len(userWiththisEmail) != 0 {
		return "", 0, fmt.Errorf("email already in use")
	}

	if newUser.Password != "" {
		hashedPassword, err := database.HashPassword(newUser.Password)
		if err != nil {
			return "", 0, fmt.Errorf("error while hashing the password")
		}
		newUser.Password = hashedPassword
	}

	service.repository.Save(newUser)
	return service.serviceJWT.GenerateToken(fmt.Sprint(newUser.Id), newUser.Username, false), newUser.Id, nil
}

func (service *userService) GetUserInfo(token string) (userInfo schemas.User, err error) {
	userId, err := service.serviceJWT.GetUserIdfromJWTToken(token)
	if err != nil {
		return schemas.User{}, fmt.Errorf("unable to get user info because %w", err)
	}
	userInfo = service.repository.FindById(userId)
	return userInfo, nil
}

func (service *userService) UpdateUserInfo(newUser schemas.User) (err error) {
	service.repository.Update(newUser)
	return nil
}

func (service *userService) GetUserById(userId uint64) schemas.User {
	return service.repository.FindById(userId)
}
