package api

import (
	"fmt"

	_ "Github.com/LocalEats/Api-get-way/api/docs"
	"Github.com/LocalEats/Api-get-way/api/handlers"
	"Github.com/LocalEats/Api-get-way/genproto"
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger" // Import gin-swagger middleware
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

// RouterApi @title LocalEats
// @version 1.0
// @description API service
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func RouterApi(con1 *grpc.ClientConn, con2 *grpc.ClientConn) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))
	paymentCon := genproto.NewUserServiceClient(con1)
	reservationCon := genproto.NewOrderServiceClient(con2)
	h := handlers.NewHandler(paymentCon, reservationCon)
	fmt.Println(h)

	authRoutes := router.Group("/")
	//authRoutes.Use(middleware.MiddleWare())
	{
		meal := router.Group("/api/order_service/meal/")
		{
			meal.POST("/create", h.CreateMealHandler)
			meal.PUT("/:meal_id", h.UpdateMealHandler)
			meal.DELETE("/:meal_id", h.DeleteMealHandler)
			meal.GET("/:kitchen_id/meals", h.GetMealHandler)
		}

		restaurant := authRoutes.Group("/api/order_service/order")
		{
			restaurant.POST("/create", h.CreateOrderHandler)
			restaurant.PUT("/:order_id/status", h.UpdateOrderHandler)
			restaurant.GET("/chef/:kitchen_id ", h.GetOrdersForChefHandler)
			restaurant.GET("/customer/:user_id ", h.GetOrdersForCustomerHandler)
			restaurant.GET("/:id", h.GetOrderByIdHandler)
		}

		reservation := authRoutes.Group("/api/order_service/comment")
		{
			reservation.POST("/create", h.CreateCommentHandler)
			reservation.GET("/:kitchen_id", h.GetCommentHandler)
		}

		kitchen := authRoutes.Group("/api/user_service/kitchen")
		{
			kitchen.POST("/create", h.CreateKitchen)
			kitchen.PUT("/:kitchen_id", h.UpdateKitchen)
			kitchen.GET("/:id", h.GetKitchenById)
			kitchen.GET("/get_all", h.GetAllKitchens)
			kitchen.DELETE("/search", h.SearchKitchens)
		}

		userProfile := authRoutes.Group("/api/user_service/users/")
		{
			userProfile.GET("/:id/profile", h.CreateOrderHandler)
			userProfile.PUT("/:id/profile", h.UpdateUserProfile)
		}
		//
		//	payment := authRoutes.Group("/api/payment")
		//	{
		//		payment.POST("/create", h.CreatePaymentHandler)
		//		payment.GET("/get_id/:id", h.GetByIdPaymentHandler)
		//		payment.PUT("/update/:id", h.UpdatePaymentHandler)
		//		payment.GET("/get_all", h.GetAllPaymentHandler)
		//	}
		//}

		return router
	}
}
