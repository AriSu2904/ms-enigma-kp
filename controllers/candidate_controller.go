package controllers

import (
	"awesomeProject/models"
	"awesomeProject/services"
	"awesomeProject/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CandidateController struct {
	candidateService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req models.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	authResponse, err := c.authService.Login(&req)
	if err != nil {
		utils.ErrorJSON(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	utils.SuccessJSON(ctx, http.StatusOK, "login successful", authResponse)
}
