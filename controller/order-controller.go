package controller

import (
	"h8-assignment-2/helpers"
	"h8-assignment-2/models"
	"h8-assignment-2/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController interface {
	FindAll(ctx *gin.Context)
	CreateOrder(ctx *gin.Context)
	UpdateOrder(ctx *gin.Context)
	DeleteOrder(ctx *gin.Context)
}

type orderController struct {
	service service.OrderService
}

func NewOrderController(service service.OrderService) OrderController {
	return &orderController{
		service: service,
	}
}

func (c *orderController) FindAll(ctx *gin.Context) {

	orders, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else if len(orders) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No orders found"})
		return
	}
	response := helpers.ResponseFindAll{
		Data: orders,
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *orderController) CreateOrder(ctx *gin.Context) {
	var order models.Orders

	err := ctx.ShouldBindJSON(&order)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdOrder, err := c.service.CreateOrder(&order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, createdOrder)
}

func (c *orderController) UpdateOrder(ctx *gin.Context) {
	var reqOrder helpers.RequestOrder

	err := ctx.ShouldBindJSON(&reqOrder)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idOrder := ctx.Param("id")
	i, err := strconv.ParseUint(idOrder, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	myUint := uint(i)
	reqOrder.OrderID = myUint

	updatedOrder, err := c.service.UpdateOrder(&reqOrder)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := helpers.ResponseUpdateOrder{
		Data:    updatedOrder,
		Message: "Update data success",
		Success: true,
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *orderController) DeleteOrder(ctx *gin.Context) {
	var order models.Orders
	idOrder := ctx.Param("id")
	i, err := strconv.ParseUint(idOrder, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	myUint := uint(i)
	order.ID = myUint

	err = c.service.DeleteOrder(&order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := helpers.ResponseDeleteOrder{
		Message: "Delete data success",
		Success: true,
	}

	ctx.JSON(http.StatusOK, response)
}
