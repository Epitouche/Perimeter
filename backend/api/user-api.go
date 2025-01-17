package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area/controller"
	"area/middlewares"
	"area/schemas"
	"area/service"
)

// UserApi provides an API layer for user-related operations.
// It interacts with the UserController to handle user data and actions.
type UserApi struct {
	controller controller.UserController
}

// NewUserApi initializes a new UserApi instance, sets up the necessary routes, and returns the instance.
// It configures the following routes under the "/user" group:
// - POST /login: handled by api.Login
// - POST /register: handled by api.Register
// - GET /info: handled by api.GetUserInfo (protected by JWT authorization middleware)
// - GET /info/all: handled by api.GetUserAllInfo (protected by JWT authorization middleware)
// - PUT /info: handled by api.UpdateUserInfo (protected by JWT authorization middleware)
// - DELETE /info: handled by api.DeleteUserInfo (protected by JWT authorization middleware)
//
// Parameters:
// - controller: an instance of UserController to handle user-related operations.
// - apiRoutes: a gin.RouterGroup to define the API routes.
// - serviceUser: an instance of UserService used for JWT authorization middleware.
//
// Returns:
// - A pointer to the initialized UserApi instance.
func NewUserApi(
	controller controller.UserController,
	apiRoutes *gin.RouterGroup,
	serviceUser service.UserService,
) *UserApi {
	apiRoutes = apiRoutes.Group("/user")
	api := UserApi{
		controller: controller,
	}
	api.Login(apiRoutes)
	api.Register(apiRoutes)
	apiRoutesInfo := apiRoutes.Group("/info", middlewares.AuthorizeJWT(serviceUser))
	api.GetUserInfo(apiRoutesInfo)
	api.GetUserAllInfo(apiRoutesInfo)
	api.UpdateUserInfo(apiRoutesInfo)
	api.DeleteUserInfo(apiRoutesInfo)
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
//	@Failure		400			{object}	schemas.ErrorResponse
//	@Router			/user/login [post].
func (api *UserApi) Login(apiRoutes *gin.RouterGroup) {
	apiRoutes.POST("/login", func(ctx *gin.Context) {
		token, err := api.controller.Login(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, &schemas.ErrorResponse{
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
//	@Success		201			{object}	schemas.JWT
//	@Failure		400			{object}	schemas.ErrorResponse
//	@Router			/user/register [post].
func (api *UserApi) Register(apiRoutes *gin.RouterGroup) {
	apiRoutes.POST("/register", func(ctx *gin.Context) {
		token, err := api.controller.Register(ctx)
		if err != nil {
			switch err {
			case schemas.ErrEmailTooShort,
				schemas.ErrUsernameTooShort,
				schemas.ErrPasswordTooShort,
				schemas.ErrInvalidEmail:
				ctx.JSON(http.StatusBadRequest, &schemas.ErrorResponse{
					Error: err.Error(),
				})
				return
			case schemas.ErrEmailAlreadyExist:
				ctx.JSON(http.StatusConflict, &schemas.ErrorResponse{
					Error: err.Error(),
				})
			default:
				ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{
					Error: err.Error(),
				})
				return
			}
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
		usetInfo, err := api.controller.GetUser(ctx)
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

// GetUserAllInfo godoc
//
//	@Summary		give user info of user
//	@Description	give user info of user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		bearerAuth
//	@Success		200	{object}	schemas.User
//	@Failure		401	{object}	schemas.ErrorResponse
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/user/info/ [put]
func (api *UserApi) UpdateUserInfo(apiRoutes *gin.RouterGroup) {
	apiRoutes.PUT("/", func(ctx *gin.Context) {
		usetInfo, err := api.controller.UpdateUser(ctx)
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
//	@Success		200	{object}	schemas.User
//	@Failure		401	{object}	schemas.ErrorResponse
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/user/info/ [delete]
func (api *UserApi) DeleteUserInfo(apiRoutes *gin.RouterGroup) {
	apiRoutes.DELETE("/", func(ctx *gin.Context) {
		usetInfo, err := api.controller.DeleteUser(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{
				Error: err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, usetInfo)
		}
	})
}
