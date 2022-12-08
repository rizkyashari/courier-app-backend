package controller

import (
	"backend/dto"
	"backend/entity"
	"backend/helper"
	"backend/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PromoController interface {
	AllPromos(context *gin.Context)
	FindByID(context *gin.Context)
	Update(context *gin.Context)
}

type promoController struct {
	promoService service.PromoService
	jwtService   service.JWTService
}

func NewPromoController(promoServ service.PromoService, jwtServ service.JWTService) PromoController {
	return &promoController{
		promoService: promoServ,
		jwtService:   jwtServ,
	}
}

func (c *promoController) AllPromos(context *gin.Context) {
	var promos []entity.Promo = c.promoService.All()
	res := helper.BuildResponse(true, "OK", promos)
	context.JSON(http.StatusOK, res)
}

func (c *promoController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var promo entity.Promo = c.promoService.FindByID(id)
	if (promo == entity.Promo{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", promo)
		context.JSON(http.StatusOK, res)
	}
}

func (c *promoController) Update(context *gin.Context) {
	var promoUpdateDTO dto.Promo
	errDTO := context.ShouldBind(&promoUpdateDTO)
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
	promoUpdateDTO.ID = id
	result := c.promoService.Update(promoUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}
