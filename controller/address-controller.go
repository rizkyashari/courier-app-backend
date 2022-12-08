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

type AddressController interface {
	All(context *gin.Context)
	AllAddresses(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type addressController struct {
	addressService service.AddressService
	jwtService     service.JWTService
}

func NewAddressController(addressServ service.AddressService, jwtServ service.JWTService) AddressController {
	return &addressController{
		addressService: addressServ,
		jwtService:     jwtServ,
	}
}

func (c *addressController) All(context *gin.Context) {
	idUser := context.GetString("user_id")
	userID, _ := strconv.Atoi(idUser)
	var addresses []entity.Address = c.addressService.AllByUserID(uint64(userID))
	res := helper.BuildResponse(true, "OK", addresses)
	context.JSON(http.StatusOK, res)
}

func (c *addressController) AllAddresses(context *gin.Context) {
	var addresses []entity.Address = c.addressService.All()
	res := helper.BuildResponse(true, "OK", addresses)
	context.JSON(http.StatusOK, res)
}

func (c *addressController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var address entity.Address = c.addressService.FindByID(id)
	if (address == entity.Address{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", address)
		context.JSON(http.StatusOK, res)
	}
}

func (c *addressController) Insert(context *gin.Context) {
	var addressCreateDTO dto.AddressCreateDTO
	errDTO := context.ShouldBind(&addressCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			addressCreateDTO.UserID = convertedUserID
		}
		result := c.addressService.Insert(addressCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *addressController) Update(context *gin.Context) {
	var addressUpdateDTO dto.AddressUpdateDTO
	errDTO := context.ShouldBind(&addressUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.addressService.IsAllowedToEdit(userID, addressUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			addressUpdateDTO.UserID = id
		}
		result := c.addressService.Update(addressUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You don't have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *addressController) Delete(context *gin.Context) {
	var address entity.Address
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	address.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.addressService.IsAllowedToEdit(userID, address.ID) {
		c.addressService.Delete(address)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You don't have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *addressController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	return fmt.Sprintf("%v", claims["user_id"])
}
