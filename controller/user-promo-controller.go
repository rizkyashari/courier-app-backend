package controller

import (
	"backend/dto"
	"backend/entity"
	"backend/helper"
	"backend/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserPromoController interface {
	All(context *gin.Context)
	AllUserPromos(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
}

type userPromoController struct {
	userPromoService service.UserPromoService
	jwtService       service.JWTService
}

func NewUserPromoController(userPromoServ service.UserPromoService, jwtServ service.JWTService) UserPromoController {
	return &userPromoController{
		userPromoService: userPromoServ,
		jwtService:       jwtServ,
	}
}

func (c *userPromoController) AllUserPromos(context *gin.Context) {
	var userPromos []entity.UserPromo = c.userPromoService.AllUserPromos()
	res := helper.BuildResponse(true, "OK", userPromos)
	context.JSON(http.StatusOK, res)
}

func (c *userPromoController) All(context *gin.Context) {

	// idUser, _ := utils.ExtractTokenID(context)
	idUser := context.GetString("user_id")
	userID, _ := strconv.Atoi(idUser)
	var userPromos []entity.UserPromo = c.userPromoService.All(uint64(userID))
	res := helper.BuildResponse(true, "OK", userPromos)
	context.JSON(http.StatusOK, res)
	fmt.Println("sasqqsd", idUser)
}

func (c *userPromoController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var userPromo entity.UserPromo = c.userPromoService.FindByID(id)
	if (userPromo == entity.UserPromo{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", userPromo)
		context.JSON(http.StatusOK, res)
	}
}

func (c *userPromoController) Insert(context *gin.Context) {
	var userPromoCreateDTO dto.UserPromoCreateDTO
	errDTO := context.ShouldBind(&userPromoCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			userPromoCreateDTO.UserID = convertedUserID
		}
		result := c.userPromoService.Insert(userPromoCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *userPromoController) Update(context *gin.Context) {
	var userPromoUpdateDTO dto.UserPromoUpdateDTO
	errDTO := context.ShouldBind(&userPromoUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	userPromoUpdateDTO.ID = id
	result := c.userPromoService.Update(userPromoUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

func (c *userPromoController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	return fmt.Sprintf("%v", claims["user_id"])
}
