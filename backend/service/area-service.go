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
	GetUserAreas(ctx *gin.Context) ([]schemas.Area, error)
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
	println("CreateArea Service")
	var result schemas.AreaMessage

	fmt.Printf("\n\nctx.Request.Body %+v\n\n\n", ctx.Request.Body)

	// respBody, _ := io.ReadAll(ctx.Request.Body)

	// fmt.Printf("\n\nrespBody %+v\n\n\n", respBody)

	err := json.NewDecoder(ctx.Request.Body).Decode(&result)
	if err != nil {
		println(fmt.Errorf("can't bind credentials: %w", err))
		return "", fmt.Errorf("can't bind credentials: %w", err)
	}

	fmt.Printf("\n\nresult %v\n\n\n", result)
	fmt.Printf("\n\nresult %+v\n\n\n", result)

	if result.ActionOption == "" {
		return "", fmt.Errorf("empty action empty: %w", err)
	}

	if result.ReactionOption == "" {
		return "", fmt.Errorf("empty reaction empty: %w", err)
	}

	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len("Bearer "):]

	user, err := service.serviceUser.GetUserInfo(tokenString)
	if err != nil {
		return "", fmt.Errorf("can't get user info: %w", err)
	}
	newArea := schemas.Area{
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
	return err == nil
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
				action(channelArea, area.ActionOption, area.Id)
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
				reaction(area.ReactionOption, area.Id)
				println(resultAction)
			}
		}
	}(areaStartValue, channelArea)
}

func (service *areaService) GetUserAreas(ctx *gin.Context) ([]schemas.Area, error) {
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len("Bearer "):]

	user, err := service.serviceUser.GetUserInfo(tokenString)
	if err != nil {
		return nil, fmt.Errorf("can't get user info: %w", err)
	}
	areas := service.repository.FindByUserId(user.Id)
	return areas, nil
}
