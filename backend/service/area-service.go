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
	InitArea(areaStartValue schemas.Area)
	AreaExist(id uint64) bool
}

type areaService struct {
	repository      repository.AreaRepository
	actionService   ActionService
	reactionService ReactionService
	serviceUser     UserService
	serviceService  ServiceService
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
		serviceService:  serviceService,
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
		ReactionOption: result.ReactionOption,
		Enable:         true,
		Action:         service.actionService.FindById(result.ActionId),
		Reaction:       service.reactionService.FindById(result.ReactionId),
	}
	id, error := service.repository.SaveArea(newArea)
	if error != nil {
		return "", fmt.Errorf("can't save area: %w", error)
	}
	newArea.Id = id
	service.InitArea(newArea)
	return "Area created successfully", nil
}

func (service *areaService) AreaExist(id uint64) bool {
	_, err := service.repository.FindById(id)
	if err != nil {
		return false
	}
	return true
}

func (service *areaService) InitArea(areaStartValue schemas.Area) {
	channelArea := make(chan string)
	println("go routine action")
	go func(areaStartValue schemas.Area, channelArea chan string) {
		// get the action with the id
		for service.AreaExist(areaStartValue.Id) {
			area, err := service.repository.FindById(areaStartValue.Id)
			if err != nil {
				println("error")
				return
			}
			// println(area.Action.Name)
			action := service.serviceService.FindActionbyName(area.Action.Name)
			if action == nil {
				println("action not found")
				return
			}
			if area.Enable {
				action(channelArea, area.ActionOption)
			}
		}
		println("clear")
		channelArea <- "response to clear"
	}(areaStartValue, channelArea)
	// area
	println("go routine area")
	go func(areaStartValue schemas.Area, channelArea chan string) {
		// check if the area is in the databse
		for service.AreaExist(areaStartValue.Id) {
			// check if the area is enable in the databse
			area, err := service.repository.FindById(areaStartValue.Id)
			if err != nil {
				return
			}
			reaction := service.serviceService.FindReactionbyName(area.Reaction.Name)
			if area.Enable {
				resultAction := <-channelArea
				reaction(area.ReactionOption)
				println(resultAction)
			}
		}
	}(areaStartValue, channelArea)
}
