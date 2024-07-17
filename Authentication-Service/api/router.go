package api

import (
	"Github.com/LocalEats/Authentication-Service/api/handlers"
	"github.com/gin-gonic/gin"
)

func RouterApi() *gin.Engine {
	router := gin.Default()
	h := handlers.Handler{} // Assuming Handler is defined in the handlers package

	auth := router.Group("/api/auth_service/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
		auth.PUT("/update_token", h.UpdateToken)
		auth.PUT("/update_password", h.) // Assuming you have a handler for updating password
	}

	return router
}
