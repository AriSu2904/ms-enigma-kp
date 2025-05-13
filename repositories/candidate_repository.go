package repositories

import (
	"awesomeProject/models"

	"gorm.io/gorm"
)

type CandidateRepository interface {
	Create(candidates []*models.Candidate) error
	FindByBatch(batch string) ([]*models.Candidate, error)
	FindById(id string) (*models.Candidate, error)
}

type candidateRepository struct {
	db *gorm.DB
}

func NewCandidateRepository(db *gorm.DB) CandidateRepository {
	return &candidateRepository{db: db}
}
func (r *candidateRepository) Create(candidates []*models.Candidate) error {
	return r.db.Create(candidates).Error
}

func (r *candidateRepository) FindByBatch(batch string) ([]*models.Candidate, error) {
	var candidates []*models.Candidate

	if batch == "" {
		err := r.db.Find(&candidates).Error
		if err != nil {
			return nil, err
		}
		return candidates, nil
	}

	err := r.db.Where("batch = ?", batch).Find(&candidates).Error
	if err != nil {
		return nil, err
	}
	return candidates, nil
}

func (r *candidateRepository) FindById(id string) (*models.Candidate, error) {
	var candidate models.Candidate
	err := r.db.Where("id = ?", id).First(&candidate).Error
	if err != nil {
		return nil, err
	}
	return &candidate, nil
}
