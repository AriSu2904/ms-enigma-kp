package controllers

import (
	"awesomeProject/models"
	"awesomeProject/services"
	"awesomeProject/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CandidateController struct {
	candidateService services.CandidateService
}

func NewCandidateController(service services.CandidateService) *CandidateController {
	return &CandidateController{candidateService: service}
}

func (c *CandidateController) List(ctx *gin.Context) {
	var batch string

	if err := ctx.ShouldBindQuery(&batch); err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "invalid Query Params")

		return
	}

	candidates, err := c.candidateService.List(batch)
	if err != nil {
		utils.ErrorJSON(ctx, http.StatusInternalServerError, "failed to fetch candidates")
		return
	}

	var candidateResponses []*models.CandidateResponse

	for _, candidate := range candidates {
		candidateResponse := &models.CandidateResponse{
			ID:            candidate.ID,
			FullName:      candidate.FullName,
			Nickname:      candidate.Nickname,
			DateOfBirth:   candidate.DateOfBirth,
			Address:       candidate.Address,
			PhoneNumber:   candidate.PhoneNumber,
			Email:         candidate.Email,
			Batch:         candidate.Batch,
			Skills:        candidate.Skills,
			Experience:    candidate.Experience,
			SoftSkillTest: candidate.SoftSkillTest,
			MathTest:      candidate.MathTest,
			CodingTest:    candidate.CodingTest,
			Status:        candidate.Status,
		}
		candidateResponses = append(candidateResponses, candidateResponse)
	}

	utils.SuccessJSON(ctx, http.StatusOK, "success", candidateResponses)
}

func (c *CandidateController) PredictById(ctx *gin.Context) {
	id := ctx.Param("id")

	candidate, err := c.candidateService.PredictById(id)
	if err != nil {
		utils.ErrorJSON(ctx, http.StatusInternalServerError, "failed to fetch candidate")
		return
	}

	if candidate == nil {
		utils.ErrorJSON(ctx, http.StatusNotFound, "candidate not found")
		return
	}

	candidateResponse := &models.CandidateResponse{
		ID:            candidate.ID,
		FullName:      candidate.FullName,
		Nickname:      candidate.Nickname,
		DateOfBirth:   candidate.DateOfBirth,
		Address:       candidate.Address,
		PhoneNumber:   candidate.PhoneNumber,
		Email:         candidate.Email,
		Batch:         candidate.Batch,
		Skills:        candidate.Skills,
		Experience:    candidate.Experience,
		SoftSkillTest: candidate.SoftSkillTest,
		MathTest:      candidate.MathTest,
		CodingTest:    candidate.CodingTest,
		Status:        candidate.Status,
	}

	utils.SuccessJSON(ctx, http.StatusOK, "success", candidateResponse)
}
