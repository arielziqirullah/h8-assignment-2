package routers

import (
	"fmt"
	"h8-assignment-2/controller"
	"h8-assignment-2/database"
	"h8-assignment-2/helpers"
	"h8-assignment-2/repository"
	"h8-assignment-2/service"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db              *gorm.DB                   = database.SetUpDatabaseConnection()
	OrderRepository repository.OrderRepository = repository.NewOrderRepository(db)
	OrderService    service.OrderService       = service.NewOrderService(OrderRepository)
	OrderController controller.OrderController = controller.NewOrderController(OrderService)
)

func Run() *gin.Engine {
	defer database.CloseDatabaseConnection(db)
	helpers.LoadEnv()

	routes := gin.Default()

	order := routes.Group("api/orders")
	{
		// Read All
		order.GET("/", OrderController.FindAll)
		// Create
		order.POST("/", OrderController.CreateOrder)
		// Update
		order.PUT("/:id", OrderController.UpdateOrder)
		// Delete
		order.DELETE("/:id", OrderController.DeleteOrder)
	}

	routes.Run(fmt.Sprintf(":%s", os.Getenv("APPLICATION_PORT")))
	return routes
}
