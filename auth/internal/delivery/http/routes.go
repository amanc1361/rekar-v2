package http

import (
	"github.com/gin-gonic/gin"
	"rekar.ir/v2/auth/internal/usecase"
)

// RegisterRoutes تعریف مسیرهای HTTP
func RegisterRoutes(router *gin.Engine, userUseCase *usecase.UserUseCase) {
	handler := NewUserHandler(userUseCase)
	router.POST("/login", handler.Login)
	router.POST("/register", handler.Register)
	router.GET("/user/:id", handler.GetUser)
	router.PUT("/user", handler.UpdateUser)
	router.DELETE("/user/:id", handler.DeleteUser)
}
