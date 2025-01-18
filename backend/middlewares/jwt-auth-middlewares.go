package middlewares

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"area/schemas"
	"area/service"
)

// AuthorizeJWT is a middleware function for authorizing JWT tokens in HTTP requests.
// It extracts the token from the "Authorization" header, validates it, and checks the claims.
// If the token is valid and the claims match the user information, the request is allowed to proceed.
// Otherwise, it responds with an unauthorized error and aborts the request.
//
// Parameters:
// - serviceUser: an instance of UserService to retrieve user information from the database.
//
// Returns:
// - gin.HandlerFunc: a function that handles the HTTP request and performs JWT authorization.
func AuthorizeJWT(serviceUser service.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if len(authHeader) <= len("Bearer ") {
			ctx.JSON(http.StatusUnauthorized, schemas.ErrorResponse{
				Error: "No token provided",
			})
			ctx.Abort()

			return
		}
		tokenString := authHeader[len("Bearer "):]

		token, err := service.NewJWTService().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			// log.Println("Claims: ", claims)
			// log.Println("Claims[Id]: ", claims["jti"])
			// log.Println("Claims[Name]: ", claims["name"])
			// log.Println("Claims[Admin]: ", claims["admin"])
			// log.Println("Claims[Issuer]: ", claims["iss"])
			// log.Println("Claims[IssuedAt]: ", claims["iat"])
			// log.Println("Claims[ExpiresAt]: ", claims["exp"])

			idString := claims["jti"].(string)
			idUint64, err := strconv.ParseUint(idString, 10, 64)
			if err != nil {
				ctx.JSON(http.StatusUnauthorized, schemas.ErrorResponse{
					Error: "Invalid token ID",
				})
				ctx.Abort()
				return
			}
			user, err := serviceUser.GetUserById(idUint64)
			if err != nil {
				ctx.JSON(http.StatusUnauthorized, schemas.ErrorResponse{
					Error: "Invalid token",
				})
				ctx.Abort()
			}
			Username := claims["name"].(string)
			if user.Username != Username {
				ctx.JSON(http.StatusUnauthorized, schemas.ErrorResponse{
					Error: "Invalid token",
				})
				ctx.Abort()
			}
		} else {
			log.Println(err)
			ctx.JSON(http.StatusUnauthorized, schemas.ErrorResponse{
				Error: "Invalid token",
			})
			ctx.Abort()

			return
		}
	}
}
