package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area/controller"
	"area/middlewares"
	"area/schemas"
)

type UserApi struct {
	controller controller.UserController
}

func NewUserApi(controller controller.UserController, apiRoutes *gin.RouterGroup) *UserApi {
	apiRoutes = apiRoutes.Group("/user")
	api := UserApi{
		controller: controller,
	}
	api.Login(apiRoutes)
	api.Register(apiRoutes)
	apiRoutesInfo := apiRoutes.Group("/info", middlewares.AuthorizeJWT())
	api.GetUserInfo(apiRoutesInfo)
	api.GetUserAllInfo(apiRoutesInfo)
	return &api
}

// Login godoc
//
//	@Summary		Provides a JSON Web Token
//	@Description	Authenticates a user and provides a JWT to Authorize API calls
//	@Tags			User
//	@Consume		application/x-www-form-urlencoded
//	@Produce		json
//	@Param			username	formData	string	true	"Username"
//	@Param			password	formData	string	true	"Password"
//	@Success		200			{object}	schemas.JWT
//	@Failure		401			{object}	schemas.ErrorResponse
//	@Router			/user/login [post].
func (api *UserApi) Login(apiRoutes *gin.RouterGroup) {
	apiRoutes.POST("/login", func(ctx *gin.Context) {
		token, err := api.controller.Login(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, &schemas.ErrorResponse{
				Error: err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, &schemas.JWT{
			Token: token,
		})
	})
}

// Register godoc
//
//	@Summary		Provides a JSON Web Token
//	@Description	Authenticates a user and provides a JWT to Authorize API calls
//	@Tags			User
//	@Consume		application/x-www-form-urlencoded
//	@Produce		json
//	@Param			email		formData	string	true	"Email"
//	@Param			username	formData	string	true	"Username"
//	@Param			password	formData	string	true	"Password"
//	@Success		200			{object}	schemas.JWT
//	@Failure		401			{object}	schemas.ErrorResponse
//	@Router			/user/register [post].
func (api *UserApi) Register(apiRoutes *gin.RouterGroup) {
	apiRoutes.POST("/register", func(ctx *gin.Context) {
		token, err := api.controller.Register(ctx)
		if err != nil {
			ctx.JSON(http.StatusConflict, &schemas.ErrorResponse{
				Error: err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusCreated, &schemas.JWT{
			Token: token,
		})
	})
}

// GetUserInfo godoc
//
//	@Summary		give user info of user
//	@Description	give user info of user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		bearerAuth
//	@Success		200	{object}	schemas.UserCredentials
//	@Failure		401	{object}	schemas.ErrorResponse
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/user/info [get]
func (api *UserApi) GetUserInfo(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/", func(ctx *gin.Context) {
		usetInfo, err := api.controller.GetUserInfo(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{
				Error: err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, usetInfo)
		}
	})
}

// GetUserAllInfo godoc
//
//	@Summary		give user info of user
//	@Description	give user info of user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		bearerAuth
//	@Success		200	{object}	schemas.UserAllInfo
//	@Failure		401	{object}	schemas.ErrorResponse
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/user/info/all [get]
func (api *UserApi) GetUserAllInfo(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/all", func(ctx *gin.Context) {
		usetInfo, err := api.controller.GetUserAllInfo(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{
				Error: err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, usetInfo)
		}
	})
}
