package service

import (
	"encoding/json"
	"fmt"
	"reflect"

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

// compareMaps compares two maps for equality
func compareMaps(map1, map2 map[string]interface{}) bool {
	if len(map1) != len(map2) {
		return false
	}
	for key, value1 := range map1 {
		value2, ok := map2[key]
		if !ok || reflect.TypeOf(value1) != reflect.TypeOf(value2) {
			return false
		}
	}
	return true
}

type areaService struct {
	repository        repository.AreaRepository
	actionService     ActionService
	reactionService   ReactionService
	serviceUser       UserService
	serviceService    ServiceService
	areaResultService AreaResultService
}

func NewAreaService(
	repository repository.AreaRepository,
	serviceService ServiceService,
	actionService ActionService,
	reactionService ReactionService,
	serviceUser UserService,
	areaResultService AreaResultService,
) AreaService {
	newService := areaService{
		repository:        repository,
		actionService:     actionService,
		reactionService:   reactionService,
		serviceUser:       serviceUser,
		serviceService:    serviceService,
		areaResultService: areaResultService,
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
		println(fmt.Errorf("can't bind credentials: %w", err))
		return "", fmt.Errorf("can't bind credentials: %w", err)
	}

	var actionOption, reactionOption json.RawMessage

	if err := json.Unmarshal(result.ActionOption, &actionOption); err != nil {
		return "", fmt.Errorf("can't unmarshal action option: %w", err)
	}

	if err := json.Unmarshal(result.ReactionOption, &reactionOption); err != nil {
		return "", fmt.Errorf("can't unmarshal reaction option: %w", err)
	}

	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len("Bearer "):]

	user, err := service.serviceUser.GetUserInfo(tokenString)
	if err != nil {
		return "", fmt.Errorf("can't get user info: %w", err)
	}

	areaAction := service.actionService.FindById(result.ActionId)
	areaReaction := service.reactionService.FindById(result.ReactionId)

	// check if the json key are the same as default areaAction.Option, json value can be different
	var defaultActionOption, providedActionOption map[string]interface{}
	if err := json.Unmarshal(areaAction.Option, &defaultActionOption); err != nil {
		return "", fmt.Errorf("can't unmarshal default action option: %w", err)
	}
	if err := json.Unmarshal(result.ActionOption, &providedActionOption); err != nil {
		return "", fmt.Errorf("can't unmarshal provided action option: %w", err)
	}
	if !compareMaps(defaultActionOption, providedActionOption) {
		return "", fmt.Errorf("action option does not match default option type")
	}

	var defaultReactionOption, providedReactionOption map[string]interface{}
	if err := json.Unmarshal(areaReaction.Option, &defaultReactionOption); err != nil {
		return "", fmt.Errorf("can't unmarshal default reaction option: %w", err)
	}
	if err := json.Unmarshal(result.ReactionOption, &providedReactionOption); err != nil {
		return "", fmt.Errorf("can't unmarshal provided reaction option: %w", err)
	}
	if !compareMaps(defaultReactionOption, providedReactionOption) {
		return "", fmt.Errorf("reaction option does not match default option type")
	}

	newArea := schemas.Area{
		User:           user,
		ActionOption:   result.ActionOption,
		ReactionOption: result.ReactionOption,
		Enable:         true,
		Action:         areaAction,
		Reaction:       areaReaction,
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
	println("go routine action " + areaStartValue.Action.Name)
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
				resultReaction := reaction(area.ReactionOption, area.Id)
				service.areaResultService.Save(schemas.AreaResult{
					Area:   area,
					Result: resultReaction,
				})
				println(resultAction)
				println(resultReaction)
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
