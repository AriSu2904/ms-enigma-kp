package services

import (
	"awesomeProject/connector"
	"awesomeProject/models"
	"errors"
)

type PredictorService interface {
	MachineLearningPrediction(candidate []*models.Candidate) ([]*models.Candidate, error)
}

type predictorService struct {
	machineLearningConnector connector.MachineLearningConnector
}

func NewPredictorService(machineLearningConnector connector.MachineLearningConnector) PredictorService {
	return &predictorService{
		machineLearningConnector: machineLearningConnector,
	}
}

func (s *predictorService) MachineLearningPrediction(candidate []*models.Candidate) ([]*models.Candidate, error) {
	var modelRequests []*connector.ModelRequest
	for _, c := range candidate {
		modelRequests = append(modelRequests, &connector.ModelRequest{
			Id:         c.ID,
			FullName:   c.FullName,
			CodingTest: c.CodingTest,
			//SkillExperience:   c.SkillExperience,
			//MathSoftSkillTest: c.MathSoftSkillTest,
			Status: c.Status,
		})
	}

	modelResponses := s.machineLearningConnector.Predict(modelRequests)
	if modelResponses == nil {
		return nil, errors.New("failed to get predictions")
	}

	var candidates []*models.Candidate
	for _, response := range modelResponses {
		candidates = append(candidates, &models.Candidate{
			ID:       response.Id,
			FullName: response.Name,
			Status:   response.Status,
			// PredictedStatus: response.PredictedStatus,
		})
	}
	return candidates, nil
}
