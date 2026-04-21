package controllers

import (
	"drinklytics/internal/models"
	"drinklytics/internal/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FinController struct {
	finService services.FinService
}

func NewFinController(finService services.FinService) *FinController {
	return &FinController{finService: finService}
}

func (h *FinController) CreateTransactionTypes(c *gin.Context) {
	var request models.TransactionTypes
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("Error Binding Json", err)
		c.JSON(http.StatusBadRequest, models.BuildErrorResponse("Bad Request", nil, err))
		return
	}

	result, err := h.finService.CreateTransactionTypes(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BuildErrorResponse("Bad Request", nil, err))
	} else {
		c.JSON(http.StatusOK, models.BuildResponse("Message", result, nil))
	}
}

func (h *FinController) GetAllTransactionTypes(c *gin.Context) {
	result, err := h.finService.GetAllTransactionTypes()

	if err != nil {
		c.JSON(http.StatusBadRequest, models.BuildErrorResponse("Bad Request", nil, err))
	} else {
		c.JSON(http.StatusOK, models.BuildResponse("Message", result, nil))
	}
}

func (h *FinController) UpdateTransactionTypes(c *gin.Context) {
	var request models.TransactionTypes
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("Error Binding Json", err)
		c.JSON(http.StatusBadRequest, models.BuildErrorResponse("Bad Request", nil, err))
		return
	}

	result, err := h.finService.UpdateTransactionTypes(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BuildErrorResponse("Bad Request", nil, err))
	} else {
		c.JSON(http.StatusOK, models.BuildResponse("Message", result, nil))
	}
}

func (h *FinController) DeleteTransactionTypes(c *gin.Context) {
	id := c.Param("id")
	var transactionTypeId int64 = -1
	var err error
	if id != "" {
		transactionTypeId, err = strconv.ParseInt(id, 10, 64)
	}
	if err != nil {
		log.Println("Error Parsing int to string", err)
		c.JSON(http.StatusBadRequest, models.BuildErrorResponse("Bad Request", nil, err))
	}

	err = h.finService.DeleteTransactionTypes(&transactionTypeId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BuildErrorResponse("Bad Request", nil, err))
	} else {
		c.JSON(http.StatusOK, models.BuildResponse("Message", "Successfully Deleted", nil))
	}
}

func (h *FinController) CreateTransaction(c *gin.Context) {
	var request models.TransactionRecords
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("Error Binding Json", err)
		c.JSON(http.StatusBadRequest, models.BuildErrorResponse("Bad Request", nil, err))
		return
	}

	result, err := h.finService.CreateTransaction(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BuildErrorResponse("Bad Request", nil, err))
	} else {
		c.JSON(http.StatusOK, models.BuildResponse("Message", result, nil))
	}
}

func (h *FinController) GetTransaction(c *gin.Context) {
	var request models.TransactionRequests
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("Error Binding Json", err)
		c.JSON(http.StatusBadRequest, models.BuildErrorResponse("Bad Request", nil, err))
		return
	}
	filters := map[string]string{
		"start_date":  c.Query("start_date"),
		"end_date":    c.Query("end_date"),
		"trx_type_id": c.Query("trx_type_id"),
		"category":    c.Query("category"),
	}

	result, err := h.finService.GetTransaction(&request, filters)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.BuildErrorResponse("Bad Request", nil, err))
	} else {
		c.JSON(http.StatusOK, models.BuildResponse("Message", result, nil))
	}
}

func (h *FinController) UpdateTransaction(c *gin.Context) {
	var request models.TransactionRecords
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("Error Binding Json", err)
		c.JSON(http.StatusBadRequest, models.BuildErrorResponse("Bad Request", nil, err))
		return
	}

	result, err := h.finService.UpdateTransaction(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BuildErrorResponse("Bad Request", nil, err))
	} else {
		c.JSON(http.StatusOK, models.BuildResponse("Message", result, nil))
	}
}

func (h *FinController) DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	var transactionTypeId int64 = -1
	var err error
	if id != "" {
		transactionTypeId, err = strconv.ParseInt(id, 10, 64)
	}
	if err != nil {
		log.Println("Error Parsing int to string", err)
		c.JSON(http.StatusBadRequest, models.BuildErrorResponse("Bad Request", nil, err))
	}

	err = h.finService.DeleteTransaction(&transactionTypeId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BuildErrorResponse("Bad Request", nil, err))
	} else {
		c.JSON(http.StatusOK, models.BuildResponse("Message", "Successfully Deleted", nil))
	}
}

// Dashboard
func (h *FinController) FinancialSummary(c *gin.Context) {
	var request models.FinancialSummaryRequests
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("Error Binding Json", err)
		c.JSON(http.StatusBadRequest, models.BuildErrorResponse("Bad Request", nil, err))
		return
	}
	filters := map[string]string{
		"start_date": c.Query("start_date"),
		"end_date":   c.Query("end_date"),
	}

	result, err := h.finService.FinancialSummary(&request, filters)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.BuildErrorResponse("Bad Request", nil, err))
	} else {
		c.JSON(http.StatusOK, models.BuildResponse("Message", result, nil))
	}
}

func (h *FinController) CategorySummary(c *gin.Context) {
	var request models.FinancialSummaryRequests
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("Error Binding Json", err)
		c.JSON(http.StatusBadRequest, models.BuildErrorResponse("Bad Request", nil, err))
		return
	}
	filters := map[string]string{
		"start_date": c.Query("start_date"),
		"end_date":   c.Query("end_date"),
		"category":   c.Query("category"),
	}

	result, err := h.finService.CategorySummary(&request, filters)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.BuildErrorResponse("Bad Request", nil, err))
	} else {
		c.JSON(http.StatusOK, models.BuildResponse("Message", result, nil))
	}
}

func (h *FinController) RecentActivity(c *gin.Context) {
	result, err := h.finService.RecentActivity()

	if err != nil {
		c.JSON(http.StatusBadRequest, models.BuildErrorResponse("Bad Request", nil, err))
	} else {
		c.JSON(http.StatusOK, models.BuildResponse("Message", result, nil))
	}
}

func (h *FinController) Trends(c *gin.Context) {
	var request models.FinancialSummaryRequests
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("Error Binding Json", err)
		c.JSON(http.StatusBadRequest, models.BuildErrorResponse("Bad Request", nil, err))
		return
	}
	trends := c.Query("trend")

	result, err := h.finService.Trends(&request, trends)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.BuildErrorResponse("Bad Request", nil, err))
	} else {
		c.JSON(http.StatusOK, models.BuildResponse("Message", result, nil))
	}
}
