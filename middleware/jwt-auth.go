package middleware

import (
	"backend/dto"
	"backend/helper"
	"backend/service"
	"backend/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthorizeJWT validates the token user givem return 401 if not valid
func AuthorizeJWT(jwtService service.JWTService, userService service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims, ok := token.Claims.(jwt.MapClaims)
			// payload, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				response := utils.ErrorResponse("Unauthorized", http.StatusUnauthorized, "not a valid bearer token")
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}
			userID := claims["user_id"].(string)
			userID2, _ := strconv.ParseUint(userID, 10, 64)

			params := &dto.UserRequestParams{}
			params.UserID = userID2
			user, err := userService.GetUser(params)
			if err != nil {
				response := utils.ErrorResponse("Unauthorized", http.StatusUnauthorized, err.Error())
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}

			c.Set("user", user)
			c.Set("user_id", userID)
			c.Next()
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[issuer]: ", claims["issuer"])
		} else {
			log.Println(err)
			response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}

	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type,access-control-allow-headers,access-control-allow-methods,access-control-allow-origin,Access-Control-Allow-Origin, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, DELETE,OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
