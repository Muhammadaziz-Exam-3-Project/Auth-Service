package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"Github.com/LocalEats/Api-get-way/genproto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateMealHandler handles the creation of a new meal item.
// @Summary Create Meal
// @Description Create a new meal
// @Tags Meal
// @Accept json
// @Produce json
// @Param Create body genproto.CreateMealRequest true "Create Menu"
// @Success 200 {object} genproto.MealResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/order_service/meal/create [post]
func (h *Handler) CreateMealHandler(ctx *gin.Context) {
	var request genproto.CreateMealRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		BadRequest(ctx, err)
		fmt.Println("_________", err)
		return
	}

	_, err := h.UserService.GetByIdKitchen(ctx, &genproto.IdRequest{Id: request.KitchenId})
	if err != nil {
		fmt.Println("_________", err)
		BadRequest(ctx, fmt.Errorf("kitchen_id mavjud emas"))
		return
	}

	response, err := h.OrderService.CreateMeal(ctx, &request)
	if err != nil {
		fmt.Println("++++++++++++", err)

		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(200, response)
}

// UpdateMealHandler handles the update of a meal item.
// @Summary Update Meal
// @Description Update an existing meal
// @Tags Meal
// @Accept json
// @Produce json
// @Param meal_id path string true "Meal ID"
// @Param Update body genproto.UpdateMealRequest true "Update Menu"
// @Success 200 {object} genproto.MealResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/order_service/meal/{meal_id} [put]
func (h *Handler) UpdateMealHandler(ctx *gin.Context) {
	var request genproto.UpdateMealRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		BadRequest(ctx, err)
		fmt.Println("_________", err)
		return
	}
	request.Id = ctx.Param("meal_id")

	response, err := h.OrderService.UpdateMeal(ctx, &request)
	if err != nil {
		fmt.Println("++++++++++++", err)

		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(200, response)
}

// DeleteMealHandler handles the deletion of a meal item.
// @Summary Delete Meal
// @Description Delete a meal
// @Tags Meal
// @Produce json
// @Param meal_id path string true "Meal ID"
// @Success 200 {object} map[string]string{"message": "Dish successfully deleted"}
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/order_service/meal/{meal_id} [delete]

func (h *Handler) DeleteMealHandler(ctx *gin.Context) {
	var request genproto.IdRequest

	request.Id = ctx.Param("meal_id")
	_, err := h.OrderService.Delete(ctx, &request)
	if err != nil {
		fmt.Println("++++++++++++", err)

		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Dish successfully deleted",
	})
}

// GetMealHandler retrieves meals based on query parameters.
// @Summary Get Meals
// @Description Get meals based on query parameters
// @Tags Meal
// @Produce json
// @Param name query string false "Name"
// @Param category query string false "Category"
// @Param available query string false "Available"
// @Param price query string false "Price"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Param kitchen_id path string true "Kitchen ID"
// @Success 200 {object} genproto.MealsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/order_service/{kitchen_id}/meals [get]
func (h *Handler) GetMealHandler(ctx *gin.Context) {
	request := genproto.GetAllMealRequest{
		Name:     ctx.Query("name"),
		Category: ctx.Query("category"),
	}
	price := ctx.Query("price")
	number, err := IsNumber(price)
	if err != nil {
		return
	}
	request.Price = float32(number)
	limit := ctx.Query("limit")
	limit1, err := IsNumber(limit)
	if err != nil {
		BadRequest(ctx, err)
		return
	}
	offset := ctx.Query("offset")
	offset1, err := IsNumber(offset)
	if err != nil {
		BadRequest(ctx, err)
		return
	}
	request.LimitOffset.Limit = int64(limit1)
	request.LimitOffset.Offset = int64(offset1)
	request.KitchenId = ctx.Param("kitchen_id")
	if !Parse(request.KitchenId) {
		fmt.Println("_________", err)
		BadRequest(ctx, err)
		return
	}

	response, err := h.OrderService.GetAllMeal(ctx, &request)
	if err != nil {
		fmt.Println("++++++++++++", err)

		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"meals": response,
	})
}

//. order repository

// CreateOrderHandler handles the creation of a new order.
// @Summary Create Order
// @Description Create a new order
// @Tags Order
// @Accept json
// @Produce json
// @Param Create body genproto.CreateOrderRequest true "Create Order"
// @Success 200 {object} genproto.OrderResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/order_service/order/create [post]
func (h *Handler) CreateOrderHandler(ctx *gin.Context) {
	var request genproto.CreateOrderRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		BadRequest(ctx, err)
		return
	}

	if _, err := uuid.Parse(request.KitchenId); err != nil {
		BadRequest(ctx, errors.New("invalid KitchenId format"))
		return
	}

	_, err := h.UserService.GetByIdKitchen(ctx, &genproto.IdRequest{Id: request.KitchenId})
	if err != nil {
		BadRequest(ctx, errors.New("kitchen_id not found"))
		return
	}

	response, err := h.OrderService.CreateOrder(ctx, &request)
	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(200, response)
}

// UpdateOrderHandler handles the update of an order's status.
// @Summary Update Order Status
// @Description Update an order's status
// @Tags Order
// @Accept json
// @Produce json
// @Param order_id path string true "Order ID"
// @Param Update body genproto.UpdateOrderStatusRequest true "Update Order Status"
// @Success 200 {object} genproto.OrderStatusResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/order_service/order/{order_id}/status [put]
func (h *Handler) UpdateOrderHandler(ctx *gin.Context) {
	var request genproto.UpdateOrderStatusRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		BadRequest(ctx, err)
		fmt.Println("_________", err)
		return
	}
	request.Id = ctx.Param("order_id")
	response, err := h.OrderService.UpdateOrderStatus(ctx, &request)
	if err != nil {
		fmt.Println("++++++++++++", err)

		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(200, response)
}

// GetOrdersForChefHandler retrieves orders for a chef based on query parameters.
// @Summary Get Orders for Chef
// @Description Get orders for a chef based on query parameters
// @Tags Order
// @Produce json
// @Param kitchen_id path string true "Kitchen ID"
// @Param status query string false "Status"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} genproto.GetOrdersResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/order_service/orders/chef/{kitchen_id} [get]
func (h *Handler) GetOrdersForChefHandler(ctx *gin.Context) {
	request := genproto.GetOrdersRequest{
		Status: ctx.Query("status"),
	}
	request.KitchenId = ctx.Param("kitchen_id")
	if err := ctx.ShouldBindQuery(&request.LimitOffset); err != nil {
		BadRequest(ctx, fmt.Errorf("Invalid request is limit offset "))
		return
	}
	_, err := h.UserService.GetByIdKitchen(ctx, &genproto.IdRequest{Id: request.KitchenId})
	if err != nil {
		fmt.Println("_________", err)
		BadRequest(ctx, fmt.Errorf("kitchen id found not in kichen table"))
	}

	orders, err := h.OrderService.GetOrdersForChef(ctx, &request)
	if err != nil {
		fmt.Println("++++++++++++", err)

		InternalServerError(ctx, fmt.Errorf("Failed to retrieve orders"))
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

// GetOrdersForCustomerHandler retrieves orders for a customer based on query parameters.
// @Summary Get Orders for Customer
// @Description Get orders for a customer based on query parameters
// @Tags Order
// @Produce json
// @Param user_id path string true "User ID"
// @Param status query string false "Status"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} genproto.GetOrdersResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/order_service/orders/customer/{user_id} [get]
func (h *Handler) GetOrdersForCustomerHandler(ctx *gin.Context) {
	request := genproto.GetOrdersRequest{
		Status: ctx.Query("status"),
	}
	request.UserId = ctx.Param("user_id")

	if err := ctx.ShouldBindQuery(&request.LimitOffset); err != nil {
		fmt.Println("_________", err)
		BadRequest(ctx, fmt.Errorf("Invalid request"))
		return
	}

	orders, err := h.OrderService.GetOrdersForCustomer(ctx, &request)
	if err != nil {
		fmt.Println("++++++++++++", err)

		InternalServerError(ctx, fmt.Errorf("Failed to retrieve orders"))
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

// GetOrderByIdHandler retrieves an order by its ID.
// @Summary Get Order by ID
// @Description Get an order by its ID
// @Tags Order
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} genproto.OrderResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/order_service/order/{id} [get]
func (h *Handler) GetOrderByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if !Parse(id) {
		fmt.Println("_________")
		BadRequest(ctx, fmt.Errorf("Invalid order ID"))
		return
	}

	request := &genproto.GetOrderRequest{Id: id}
	order, err := h.OrderService.GetOrderById(ctx, request)
	if err != nil {
		fmt.Println("++++++++++++", err)

		InternalServerError(ctx, fmt.Errorf("Failed to retrieve orders"))

		return
	}
	ctx.JSON(http.StatusOK, order)
}

// CreateCommentHandler handles the creation of a new comment.
// @Summary Create Comment
// @Description Create a new comment
// @Tags Comment
// @Accept json
// @Produce json
// @Param Create body genproto.CreateReviewRequest true "Create Comment"
// @Success 200 {object} genproto.ReviewResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/order_service/comment/create [post]
func (h *Handler) CreateCommentHandler(ctx *gin.Context) {
	request := genproto.CreateReviewRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		fmt.Println("_________", err)
		BadRequest(ctx, fmt.Errorf("sholud not blind json "))
		return
	}
	_, err = h.UserService.GetByIdKitchen(ctx, &genproto.IdRequest{
		Id: request.KitchenId,
	})
	if err != nil {
		fmt.Println("_________", err)
		BadRequest(ctx, fmt.Errorf("kitchen id not found kitchen table"))
	}

	order, err := h.OrderService.CreateReview(ctx, &request)
	if err != nil {
		fmt.Println("++++++++++++", err)

		InternalServerError(ctx, fmt.Errorf("Failed to retrieve comment"))
		return
	}
	ctx.JSON(http.StatusOK, order)
}

// GetCommentHandler retrieves comments for a kitchen.
// @Summary Get Comments
// @Description Get comments for a kitchen
// @Tags Comment
// @Produce json
// @Param kitchen_id path string true "Kitchen ID"
// @Success 200 {object} genproto.GetReviewsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/order_service/comments/{kitchen_id} [get]
func (h *Handler) GetCommentHandler(ctx *gin.Context) {
	request := genproto.GetReviewsRequest{}
	request.KitchenId = ctx.Param("kitchen_id")
	_, err := h.UserService.GetByIdKitchen(ctx, &genproto.IdRequest{
		Id: request.KitchenId,
	})
	if err != nil {
		BadRequest(ctx, fmt.Errorf("kitchen id not found kitchen table"))
	}

	comments, err := h.OrderService.GetReviews(ctx, &request)
	if err != nil {
		fmt.Println("++++++++++++", err)
		InternalServerError(ctx, fmt.Errorf("Failed to retrieve comment"))
		return
	}
	ctx.JSON(http.StatusOK, comments)
}

// CreatePaymentHandler handles the creation of payments.
// @Summary Create Payment
// @Description Create a new payment
// @Tags Payment
// @Accept json
// @Produce json
// @Param payment body genproto.CreatePaymentRequest true "Payment Request"
// @Success 200 {object} genproto.CreatePaymentResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/order_service/payments [post]
func (h *Handler) CreatePaymentHandler(ctx *gin.Context) {
	var request genproto.CreatePaymentRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		BadRequest(ctx, fmt.Errorf("Invalid request body"))
		return
	}
	response, err := h.OrderService.CreatePayment(ctx, &request)
	if err != nil {
		fmt.Println("++++++++++++", err)

		InternalServerError(ctx, fmt.Errorf("payment not created"))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// UpdateNutritionInfoHandler updates the nutrition information of a dish.
// @Summary Update nutrition information of a dish
// @Description Updates the nutrition information of a dish based on the provided request payload.
// @ID updateNutritionInfo
// @Tags qo'shimcha Api
// @Accept json
// @Produce json
// @Param request body genproto.UpdateNutritionInfoRequest true "Request payload"
// @Success 200 {object} genproto.Dish "Updated dish information"
// @Router /api/order_service/meal/update-nutrition-info [put]
func (h *Handler) UpdateNutritionInfoHandler(ctx *gin.Context) {
	var req genproto.UpdateNutritionInfoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		BadRequest(ctx, fmt.Errorf("Invalid request payload"))
		return
	}
	dish, err := h.OrderService.UpdateNutrition(ctx, &req)
	if err != nil {
		InternalServerError(ctx, fmt.Errorf("Failed to update nutrition info"))
		return
	}

	ctx.JSON(http.StatusOK, dish)
}
