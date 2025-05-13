package controllers

import (
	"awesomeProject/services"
)

type CandidateController struct {
}

func NewCandidateController(authService services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}
