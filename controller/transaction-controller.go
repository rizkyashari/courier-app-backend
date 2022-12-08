package controller

import (
	"backend/dto"
	"backend/entity"
	"backend/service"
	"net/http"

	"backend/utils"

	"github.com/gin-gonic/gin"
)

type TransactionController interface {
	GetTransactions(context *gin.Context)
	TopUp(context *gin.Context)
	Payment(context *gin.Context)
}

type transactionController struct {
	transactionService service.TransactionService
	jwtService         service.JWTService
}

func NewTransactionController(transactionServ service.TransactionService, jwtServ service.JWTService) TransactionController {
	return &transactionController{
		transactionService: transactionServ,
		jwtService:         jwtServ,
	}
}

func (c *transactionController) GetTransactions(context *gin.Context) {
	query := &dto.TransactionRequestQuery{}
	err := context.ShouldBindQuery(query)
	if err != nil {
		if err != nil {
			errors := utils.FormatValidationError(err)
			response := utils.ErrorResponse("get transaction failed", http.StatusUnprocessableEntity, errors)
			context.JSON(http.StatusUnprocessableEntity, response)
			return
		}
	}
	query = dto.FormatQuery(query)

	user := context.MustGet("user").(*entity.User)
	// var user *entity.User
	transactions, err := c.transactionService.GetTransactions(uint64(user.ID), query)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("get transactions failed", statusCode, err.Error())
		context.JSON(statusCode, response)
		return
	}
	totalTransactions, err := c.transactionService.CountTransaction(uint64(user.ID))
	if err != nil {
		response := utils.ErrorResponse("get transactions failed", http.StatusInternalServerError, err.Error())
		context.JSON(http.StatusInternalServerError, response)
		return
	}

	formattedTransaction := dto.FormatTransactions(transactions)
	metadata := utils.Metadata{Resource: "transactions", TotalAll: int(totalTransactions), TotalNow: len(transactions), Page: query.Page, Limit: query.Limit, Sort: query.Sort}
	response := utils.ResponseWithPagination("get transaction success", http.StatusOK, formattedTransaction, metadata)
	context.JSON(http.StatusOK, response)
}

func (c *transactionController) TopUp(context *gin.Context) {
	input := &dto.TopUpRequestBody{}
	err := context.ShouldBindJSON(input)
	if err != nil {
		if err != nil {
			errors := utils.FormatValidationError(err)
			response := utils.ErrorResponse("top up failed", http.StatusUnprocessableEntity, errors)
			context.JSON(http.StatusUnprocessableEntity, response)
			return
		}
	}

	user := context.MustGet("user").(*entity.User)
	// var user *entity.User
	input.User = user
	transaction, err := c.transactionService.TopUp(input)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("top up failed", statusCode, err.Error())
		context.JSON(statusCode, response)
		return
	}

	formattedTransaction := dto.FormatTopUp(transaction)
	response := utils.SuccessResponse("top up success", http.StatusOK, formattedTransaction)
	context.JSON(http.StatusOK, response)
}

func (c *transactionController) Payment(context *gin.Context) {
	input := &dto.PaymentRequestBody{}
	err := context.ShouldBindJSON(input)
	if err != nil {
		if err != nil {
			errors := utils.FormatValidationError(err)
			response := utils.ErrorResponse("payment failed", http.StatusUnprocessableEntity, errors)
			context.JSON(http.StatusUnprocessableEntity, response)
			return
		}
	}

	user := context.MustGet("user").(*entity.User)
	// var user *entity.User
	input.User = user
	transaction, err := c.transactionService.Payment(input)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("payment failed", statusCode, err.Error())
		context.JSON(statusCode, response)
		return
	}

	formattedTransaction := dto.FormatPayment(transaction)
	response := utils.SuccessResponse("payment success", http.StatusOK, formattedTransaction)
	context.JSON(http.StatusOK, response)
}
