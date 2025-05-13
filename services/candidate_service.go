package services

import (
	"awesomeProject/models"
	"awesomeProject/repositories"
	"errors"
)

type CandidateService interface {
	List(batch string) ([]*models.Candidate, error)
	Create(candidates []*models.Candidate) error
	PredictByBatch(batch string) ([]*models.Candidate, error)
	PredictById(id string) ([]*models.Candidate, error)
}

type candidateService struct {
	candidateRepository repositories.CandidateRepository
	predictorService    PredictorService
}

func NewCandidateService(candidateRepository repositories.CandidateRepository, pr PredictorService) CandidateService {
	return &candidateService{
		candidateRepository: candidateRepository,
		predictorService:    pr,
	}
}

func (s *candidateService) List(batch string) ([]*models.Candidate, error) {
	candidates, err := s.candidateRepository.FindByBatch(batch)
	if err != nil {
		return nil, err
	}
	return candidates, nil
}

func (s *candidateService) Create(candidates []*models.Candidate) error {
	if len(candidates) == 0 {
		return errors.New("no candidates to create")
	}

	return s.candidateRepository.Create(candidates)
}

func (s *candidateService) PredictByBatch(batch string) ([]*models.Candidate, error) {
	candidates, err := s.candidateRepository.FindByBatch(batch)
	if err != nil {
		return nil, err
	}
	return candidates, nil
}

func (s *candidateService) PredictById(id string) ([]*models.Candidate, error) {
	candidate, err := s.candidateRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	if candidate == nil {
		return nil, errors.New("candidate not found")
	}

	prediction, err := s.predictorService.MachineLearningPrediction([]*models.Candidate{candidate})
	if err != nil {
		return nil, err
	}
	return prediction, nil
}
