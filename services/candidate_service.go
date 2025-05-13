package services

import (
	"awesomeProject/models"
	"awesomeProject/repositories"
	"errors"
)

type CandidateService interface {
	List(batch string) ([]*models.Candidate, error)
	Create(candidates []*models.Candidate) error
}

type candidateService struct {
	candidateRepository repositories.CandidateRepository
}

func NewCandidateService(candidateRepository repositories.CandidateRepository) CandidateService {
	return &candidateService{
		candidateRepository: candidateRepository,
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
