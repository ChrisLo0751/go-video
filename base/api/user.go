package api

import (
	"go-video/user-web/dto"
	"go-video/user-web/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	user, err := h.userService.GetUser(int(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	response := dto.UserResponse{
		ID:    user.ID,
		Phone: user.Phone,
	}
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.S().Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}

	user := &dto.CreateUserRequest{
		Phone:    req.Phone,
		Password: req.Password,
	}
	if err := h.userService.CreateUser(user); err != nil {
		zap.S().Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建用户失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, &dto.BaseResponese{
		Code: http.StatusCreated,
		Msg:  "成功创建用户",
	})
}
