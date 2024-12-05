package service

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"

	"area/repository"
	"area/schemas"
)

type AreaService interface {
	FindAll() []schemas.Area
	CreateArea(ctx *gin.Context) (string, error)
}

type areaService struct {
	repository      repository.AreaRepository
	actionService   ActionService
	reactionService ReactionService
	serviceUser     UserService
}

func NewAreaService(
	repository repository.AreaRepository,
	serviceService ServiceService,
	actionService ActionService,
	reactionService ReactionService,
	serviceUser UserService,
) AreaService {
	newService := areaService{
		repository:      repository,
		actionService:   actionService,
		reactionService: reactionService,
		serviceUser:     serviceUser,
	}
	return &newService
}

func (service *areaService) FindAll() []schemas.Area {
	return service.repository.FindAll()
}

func (service *areaService) CreateArea(ctx *gin.Context) (string, error) {
	var result schemas.AreaMessage
	err := json.NewDecoder(ctx.Request.Body).Decode(&result)
	if err != nil {
		return "", fmt.Errorf("can't bind credentials: %w", err)
	}
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len("Bearer "):]

	user, err := service.serviceUser.GetUserInfo(tokenString)
	if err != nil {
		return "", fmt.Errorf("can't get user info: %w", err)
	}
	newArea := schemas.Area{
		UserId:         result.UserId,
		User:           user,
		ActionOption:   result.ActionOption,
		ActionId:       result.ActionId,
		ReactionOption: result.ReactionOption,
		ReactionId:     result.ReactionId,
		Enable:         true,
		Action:         service.actionService.FindById(result.ActionId),
		Reaction:       service.reactionService.FindById(result.ReactionId),
	}
	// TODO
	service.repository.Save(newArea)
	return "Area created successfully", nil
}
