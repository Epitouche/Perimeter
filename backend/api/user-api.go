package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area/controller"
	"area/schemas"
)

type UserApi struct {
	controller controller.UserController
}

func NewUserApi(controller controller.UserController) *UserApi {
	return &UserApi{
		controller: controller,
	}
}

// Paths Information

// Authenticate godoc
// @Summary Provides a JSON Web Token
// @Description Authenticates a user and provides a JWT to Authorize API calls
// @ID Authentication
// @Consume application/x-www-form-urlencoded
// @Produce json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {object} schemas.JWT
// @Failure 401 {object} schemas.Response
// @Router /auth/token [post].
func (api *UserApi) Login(ctx *gin.Context) {
	token, err := api.controller.Login(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, &schemas.Response{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &schemas.JWT{
		Token: token,
	})
}

func (api *UserApi) Register(ctx *gin.Context) {
	token, err := api.controller.Register(ctx)
	if err != nil {
		ctx.JSON(http.StatusConflict, &schemas.Response{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, &schemas.JWT{
		Token: token,
	})
}
