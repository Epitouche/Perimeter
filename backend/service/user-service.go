package service

import (
	"fmt"
	"strconv"

	"github.com/Epitouche/Perimeter/database"
	"github.com/Epitouche/Perimeter/repository"
	"github.com/Epitouche/Perimeter/schemas"
)

type UserService interface {
	Login(user schemas.User) (jwtToken string, userID uint64, err error)
	Register(newUser schemas.User) (jwtToken string, userID uint64, err error)
	GetUserInfo(token string) (userInfo schemas.User, err error)
	UpdateUserInfo(newUser schemas.User) (err error)
	GetUserById(userID uint64) (user schemas.User, err error)
	DeleteUser(newUser schemas.User) (err error)
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

func (service *userService) Login(
	newUser schemas.User,
) (jwtToken string, userID uint64, err error) {
	userWiththisUserName, err := service.repository.FindByUserName(newUser.Username)
	if err != nil {
		return "", 0, err
	}
	if len(userWiththisUserName) == 0 {
		return "", 0, schemas.ErrInvalidCredentials
	}
	// regular user
	for _, user := range userWiththisUserName {
		if database.DoPasswordsMatch(user.Password, newUser.Password) {
			return service.serviceJWT.GenerateToken(
				strconv.FormatUint(user.Id, 10),
				user.Username,
				false,
			), user.Id, nil
		}
	}

	// Oauth2.0 user
	for _, user := range userWiththisUserName {
		if user.Email == newUser.Email {
			if user.TokenId != 0 {
				return service.serviceJWT.GenerateToken(
					strconv.FormatUint(user.Id, 10),
					user.Username,
					false,
				), user.Id, nil
			}
		}
	}

	return "", 0, schemas.ErrUserNotFound
}

func (service *userService) Register(
	newUser schemas.User,
) (jwtToken string, userID uint64, err error) {
	userWiththisEmail, err := service.repository.FindByEmail(newUser.Email)
	if err != nil {
		return "", 0, err
	}
	fmt.Printf("%+v\n", userWiththisEmail)

	if len(userWiththisEmail) != 0 {
		// return service.Login(newUser)
		return "", 0, schemas.ErrEmailAlreadyExist
	}

	if newUser.Password != "" {
		hashedPassword, err := database.HashPassword(newUser.Password)
		if err != nil {
			return "", 0, schemas.ErrHashingPassword
		}
		newUser.Password = hashedPassword
	}

	service.repository.Save(newUser)

	userTemp, err := service.repository.FindByUserName(newUser.Username)
	if err != nil {
		return "", 0, err
	}
	newUser.Id = userTemp[0].Id

	return service.serviceJWT.GenerateToken(
		strconv.FormatUint(newUser.Id, 10),
		newUser.Username,
		false,
	), newUser.Id, nil
}

func (service *userService) GetUserInfo(token string) (userInfo schemas.User, err error) {
	userId, err := service.serviceJWT.GetUserIdfromJWTToken(token)
	if err != nil {
		return schemas.User{}, fmt.Errorf("unable to get user info because %w", err)
	}
	userInfo, err = service.repository.FindById(userId)
	if err != nil {
		return schemas.User{}, fmt.Errorf("unable to get user info because %w", err)
	}
	return userInfo, nil
}

func (service *userService) UpdateUserInfo(newUser schemas.User) (err error) {
	return service.repository.Update(newUser)
}

func (service *userService) GetUserById(userID uint64) (user schemas.User, err error) {
	return service.repository.FindById(userID)
}

func (service *userService) DeleteUser(newUser schemas.User) (err error) {
	return service.repository.Delete(newUser)
}
