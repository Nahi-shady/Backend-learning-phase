package controllers

import (
	"net/http"

	"task-manager/usecases"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase *usecases.UserUsecase
}

func NewUserController(userUsecase *usecases.UserUsecase) *UserController {
	return &UserController{
		userUsecase: userUsecase,
	}
}

func (c *UserController) Register(ctx *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if err := c.userUsecase.Register(req.Username, req.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (c *UserController) Login(ctx *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	token, err := c.userUsecase.LoginUser(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
