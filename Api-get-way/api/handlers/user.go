package handlers

import (
	"fmt"
	"net/http"
	"time"

	pb "Github.com/LocalEats/Api-get-way/genproto"
	"github.com/gin-gonic/gin"
)

// CreateKitchen Kitchen
// @Summary Create a new kitchen
// @Description Create a new kitchen based on the provided request
// @Tags Kitchen
// @Accept  json
// @Produce  json
// @Param input body genproto.CreateKitchenRequest true "Kitchen details to create"
// @Success 200 {object} genproto.CreateKitchenResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/user_service/kitchen/create [post]
func (h *Handler) CreateKitchen(ctx *gin.Context) {
	var request pb.CreateKitchenRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		BadRequest(ctx, fmt.Errorf("fild not true input"))
		return
	}

	response, err := h.UserService.CreateKitchen(ctx, &request)
	if err != nil {
		fmt.Println("+++++++", err)
		InternalServerError(ctx, fmt.Errorf("Failed to create kitchen"))
		return
	}

	ctx.JSON(200, response)
}

// UpdateKitchen godoc
// @Summary Update an existing kitchen
// @Description Update an existing kitchen based on the provided request
// @Tags Kitchen
// @Accept  json
// @Produce  json
// @Param kitchen_id path string true "Kitchen ID to update"
// @Param input body genproto.UpdateKitchenRequest true "Updated kitchen details"
// @Success 200 {object} genproto.UpdateKitchenResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/user_service/kitchen/{kitchen_id} [put]
func (h *Handler) UpdateKitchen(ctx *gin.Context) {
	var request pb.UpdateKitchenRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		BadRequest(ctx, err)
		return
	}

	request.Id = ctx.Param("kitchen_id")
	response, err := h.UserService.UpdateKitchen(ctx, &request)
	if err != nil {
		InternalServerError(ctx, fmt.Errorf("Failed to update kitchen"))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// GetKitchenById godoc
// @Summary Get a kitchen by ID
// @Description Retrieve a kitchen by its ID
// @Tags Kitchen
// @Accept  json
// @Produce  json
// @Param id path string true "Kitchen ID to fetch"
// @Success 200 {object} genproto.KitchenResponse
// @Failure 404 {object} string
// @Router /api/user_service/kitchen/{id} [get]
func (h *Handler) GetKitchenById(ctx *gin.Context) {
	id := ctx.Param("id")

	request := &pb.IdRequest{Id: id}
	response, err := h.UserService.GetByIdKitchen(ctx, request)
	if err != nil {
		InternalServerError(ctx, fmt.Errorf("kitchen not found"))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// GetAllKitchens godoc
// @Summary Get all kitchens
// @Description Retrieve a list of all kitchens with optional pagination
// @Tags Kitchen
// @Accept  json
// @Produce  json
// @Param limit query int false "Limit the number of results"
// @Param offset query int false "Offset for pagination"
// @Success 200 {object} genproto.KitchensResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/user_service/kitchen/get_all [get]
func (h *Handler) GetAllKitchens(ctx *gin.Context) {
	var request pb.LimitOffset
	if err := ctx.ShouldBindQuery(&request); err != nil {
		BadRequest(ctx, fmt.Errorf("limit offset error"))
		return
	}

	response, err := h.UserService.GetAll(ctx, &request)
	if err != nil {
		fmt.Println("+++++++", err)
		InternalServerError(ctx, fmt.Errorf("Failed to fetch kitchens"))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// SearchKitchens godoc
// @Summary Search kitchens
// @Description Search for kitchens based on various criteria
// @Tags Kitchen
// @Accept  json
// @Produce  json
// @Param name query string false "Name of the kitchen to search for"
// @Param rating query float64 false "Rating of the kitchen"
// @Param address query string false "Address of the kitchen"
// @Param total_orders query int false "Total orders for the kitchen"
// @Param phone_number query string false "Phone number of the kitchen"
// @Param cuisine_type query string false "Cuisine type of the kitchen"
// @Param description query string false "Description of the kitchen"
// @Param owner_id query string false "Owner ID of the kitchen"
// @Param limit query int false "Limit the number of results"
// @Param offset query int false "Offset for pagination"
// @Success 200 {object} genproto.KitchensResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/user_service/kitchen/search [get]
func (h *Handler) SearchKitchens(ctx *gin.Context) {
	var request pb.SearchKitchenRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		BadRequest(ctx, fmt.Errorf("query not found"))
		return
	}

	response, err := h.UserService.SearchKitchen(ctx, &request)
	if err != nil {
		InternalServerError(ctx, fmt.Errorf("Failed to search kitchens"))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// GetUserProfile godoc
// @Summary Get user profile by ID
// @Description Retrieves a user's profile information based on the provided ID
// @Tags UserProfile
// @ID get-user-profile
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} genproto.UserResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/user_service/users/{id}/profile [get]
func (h *Handler) GetUserProfile(ctx *gin.Context) {
	userId := ctx.Param("id")
	request := &pb.IdRequest{Id: userId}

	response, err := h.UserService.UserProfile(ctx, request)
	if err != nil {
		InternalServerError(ctx, fmt.Errorf("failed to fetch user profile"))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// UpdateUserProfile godoc
// @Summary Update user profile
// @Description Updates a user's profile information based on the provided data
// @Tags UserProfile
// @ID update-user-profile
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param request body genproto.UpdateUserProfileRequest true "Update User Profile Request"
// @Success 200 {object} genproto.UserResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/user_service/users/{id}/profile [put]
func (h *Handler) UpdateUserProfile(ctx *gin.Context) {
	var request pb.UpdateUserProfileRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		BadRequest(ctx, fmt.Errorf("invalid input data"))
		return
	}

	request.Id = ctx.Param("id")

	response, err := h.UserService.UpdateUserProfile(ctx, &request)
	if err != nil {
		InternalServerError(ctx, fmt.Errorf("failed to update user profile"))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// GetKitchenStatisticsHandler retrieves statistics for a kitchen within a specified date range.
// @Summary Retrieve kitchen statistics
// @Description Retrieves statistics for a kitchen based on the provided kitchen_id, start_date, and end_date.
// @ID getKitchenStatistics
// @Produce json
// @Tags qo'shimcha Api
// @Param kitchen_id path string true "Kitchen ID"
// @Param start_date query string true "Start date (YYYY-MM-DD)"
// @Param end_date query string true "End date (YYYY-MM-DD)"
// @Success 200 {object} genproto.KitchenStatisticsResponse "Statistics for the kitchen"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/user_service/kitchen/{kitchen_id}/statistics [get]
func (h *Handler) GetKitchenStatisticsHandler(ctx *gin.Context) {
	kitchenID := ctx.Param("kitchen_id")
	startDateStr := ctx.Query("start_date")
	endDateStr := ctx.Query("end_date")

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		BadRequest(ctx, fmt.Errorf("Invalid start_date format. Use YYYY-MM-DD"))
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		BadRequest(ctx, fmt.Errorf("Invalid start_date format. Use YYYY-MM-DD"))
		return
	}

	request := &pb.GetUserActivityRequest{
		UserId:    kitchenID,
		StartDate: startDate.String(),
		EndDate:   endDate.String(),
	}

	stats, err := h.UserService.KitchenStatistic(ctx, request)
	if err != nil {
		InternalServerError(ctx, fmt.Errorf("Failed to fetch kitchen statistics"))
		return
	}

	ctx.JSON(http.StatusOK, stats)
}

// GetUserActivityHandler retrieves activity statistics for a user within a specified date range.
// @Summary Retrieve user activity
// @Description Retrieves activity statistics for a user based on the provided start_date and end_date query parameters.
// @ID getUserActivity
// @Tags qo'shimcha Api
// @Produce json
// @Param start_date query string true "Start date (YYYY-MM-DD)"
// @Param end_date query string true "End date (YYYY-MM-DD)"
// @Success 200 {object} genproto.GetUserActivityResponse "Activity statistics for the user"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/user_service/users/activity [get]
func (h *Handler) GetUserActivityHandler(ctx *gin.Context) {
	var req pb.GetUserActivityRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		BadRequest(ctx, fmt.Errorf("Invalid request parameter"))
		return
	}

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		BadRequest(ctx, fmt.Errorf("Invalid start_date format. Use YYYY-MM-DD"))
		return
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		BadRequest(ctx, fmt.Errorf("Invalid start_date format. Use YYYY-MM-DD"))
		return
	}

	req.StartDate = startDate.String()
	req.EndDate = endDate.String()

	stats, err := h.UserService.ActivityUser(ctx, &req)
	if err != nil {
		InternalServerError(ctx, fmt.Errorf("Failed to fetch user activitys"))

		return
	}

	ctx.JSON(http.StatusOK, stats)
}

// UpdateWorkingHoursHandler updates the working hours of a kitchen.
// @Summary Update kitchen working hours
// @Description Updates the working hours of a kitchen based on the provided request payload.
// @ID updateWorkingHours
// @Tags qo'shimcha Api
// @Accept json
// @Produce json
// @Param request body genproto.UpdateWorkingHoursRequest true "Request payload"
// @Success 200 {object} genproto.UpdateWorkingHoursResponse "Updated working hours information"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/user_service/kitchen/update-working-hours [put]
func (h *Handler) UpdateWorkingHoursHandler(ctx *gin.Context) {
	var req pb.UpdateWorkingHoursRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		BadRequest(ctx, fmt.Errorf("Invalid request payload"))
		return
	}

	response, err := h.UserService.UpdateWorkingHours(ctx, &req)
	if err != nil {
		InternalServerError(ctx, fmt.Errorf("Failed to update working hours"))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
