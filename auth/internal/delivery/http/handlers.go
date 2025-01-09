package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rekar.ir/v2/auth/internal/domain/models"
	"rekar.ir/v2/auth/internal/usecase"
)

// UserHandler مدیریت درخواست‌های HTTP برای کاربران
type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

// NewUserHandler ایجاد یک نمونه جدید از UserHandler
func NewUserHandler(uc *usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: uc,
	}
}

// Login مدیریت درخواست‌های لاگین
func (h *UserHandler) Login(c *gin.Context) {
	var credentials struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	user, err := h.userUseCase.Authenticate(credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    user,
	})
}

// Register مدیریت درخواست‌های ثبت‌نام
func (h *UserHandler) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.userUseCase.Register(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
