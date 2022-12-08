package controller

import (
	"backend/dto"
	"backend/entity"
	"backend/helper"
	"backend/service"
	"backend/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ShippingController interface {
	All(context *gin.Context)
	AllShippings(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type shippingController struct {
	shippingService service.ShippingService
	jwtService      service.JWTService
}

func NewShippingController(shippingServ service.ShippingService, jwtServ service.JWTService) ShippingController {
	return &shippingController{
		shippingService: shippingServ,
		jwtService:      jwtServ,
	}
}

func (c *shippingController) AllShippings(context *gin.Context) {
	var shippings []entity.Shipping = c.shippingService.AllShippings()
	res := helper.BuildResponse(true, "OK", shippings)
	context.JSON(http.StatusOK, res)
}

func (c *shippingController) All(context *gin.Context) {

	// idUser := context.GetString("user_id")
	// userID, _ := strconv.Atoi(idUser)
	// var shippings []entity.Shipping = c.shippingService.All(uint64(userID))
	// res := helper.BuildResponse(true, "OK", shippings)
	// context.JSON(http.StatusOK, res)
	// fmt.Println("sasqqsd", idUser)

	query := &dto.ShippingRequestQuery{}
	err := context.ShouldBindQuery(query)
	if err != nil {
		if err != nil {
			errors := utils.FormatValidationError(err)
			response := utils.ErrorResponse("get shippings failed", http.StatusUnprocessableEntity, errors)
			context.JSON(http.StatusUnprocessableEntity, response)
			return
		}
	}
	query = dto.FormatShippingsQuery(query)

	user := context.MustGet("user").(*entity.User)

	// shippings := c.shippingService.All(uint64(user.ID), query)
	shippings := c.shippingService.All(uint64(user.ID), query)

	totalShippings, err := c.shippingService.CountShipping(uint64(user.ID))
	if err != nil {
		response := utils.ErrorResponse("get shippings failed", http.StatusInternalServerError, err.Error())
		context.JSON(http.StatusInternalServerError, response)
		return
	}

	formattedShipping := dto.FormatShippings(shippings)
	metadata := utils.Metadata{Resource: "shipping", TotalAll: int(totalShippings), TotalNow: len(shippings), Page: query.Page, Limit: query.Limit, Sort: query.Sort}
	response := utils.ResponseWithPagination("get shipping success", http.StatusOK, formattedShipping, metadata)
	context.JSON(http.StatusOK, response)

}

func (c *shippingController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var shipping entity.Shipping = c.shippingService.FindByID(id)
	if (shipping == entity.Shipping{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", shipping)
		context.JSON(http.StatusOK, res)
	}
}

func (c *shippingController) Insert(context *gin.Context) {
	var shippingCreateDTO dto.ShippingCreateDTO
	errDTO := context.ShouldBind(&shippingCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			shippingCreateDTO.UserID = convertedUserID
		}
		result := c.shippingService.Insert(shippingCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *shippingController) Update(context *gin.Context) {
	var shippingUpdateDTO dto.ShippingUpdateDTO
	errDTO := context.ShouldBind(&shippingUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	// authHeader := context.GetHeader("Authorization")
	// token, errToken := c.jwtService.ValidateToken(authHeader)
	// if errToken != nil {
	// 	panic(errToken.Error())
	// }
	// claims := token.Claims.(jwt.MapClaims)
	// userID := fmt.Sprintf("%v", claims["user_id"])
	// if c.shippingService.IsAllowedToEdit(userID, shippingUpdateDTO.ID) {
	// 	id, errID := strconv.ParseUint(userID, 10, 64)
	// 	if errID == nil {
	// 		shippingUpdateDTO.UserID = id
	// 	}
	// 	result := c.shippingService.Update(shippingUpdateDTO)
	// 	response := helper.BuildResponse(true, "OK", result)
	// 	context.JSON(http.StatusOK, response)
	// } else {
	// 	response := helper.BuildErrorResponse("You don't have permission", "You are not the owner", helper.EmptyObj{})
	// 	context.JSON(http.StatusForbidden, response)
	// }

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	shippingUpdateDTO.ID = id
	result := c.shippingService.Update(shippingUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

func (c *shippingController) Delete(context *gin.Context) {
	var shipping entity.Shipping
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	shipping.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.shippingService.IsAllowedToEdit(userID, shipping.ID) {
		c.shippingService.Delete(shipping)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You don't have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *shippingController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	return fmt.Sprintf("%v", claims["user_id"])
}
