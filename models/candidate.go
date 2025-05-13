package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Candidate struct {
	ID            string    `json:"id" gorm:"type:uuid;primaryKey"`
	FullName      string    `json:"full_name"`
	Nickname      string    `json:"nickname"`
	DateOfBirth   time.Time `json:"date_of_birth"`
	Address       string    `json:"address"`
	PhoneNumber   string    `json:"phone_number"`
	Email         string    `json:"email"`
	Batch         string    `json:"batch"`
	Skills        string    `json:"skills"`
	Experience    int       `json:"experience"`
	SoftSkillTest float64   `json:"soft_skill_test"`
	MathTest      float64   `json:"math_test"`
	CodingTest    float64   `json:"coding_test"`
	TotalSkills   int       `json:"total_skills"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (s *Candidate) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New().String()
	return
}

type CandidateResponse struct {
	ID            string    `json:"id"`
	FullName      string    `json:"fullName"`
	Nickname      string    `json:"nickname"`
	DateOfBirth   time.Time `json:"dob"`
	Address       string    `json:"address"`
	PhoneNumber   string    `json:"phoneNumber"`
	Email         string    `json:"email"`
	Batch         string    `json:"batch"`
	Skills        string    `json:"skills"`
	Experience    int       `json:"experience"`
	SoftSkillTest float64   `json:"soft_skill_test"`
	MathTest      float64   `json:"math_test"`
	CodingTest    float64   `json:"coding_test"`
	Status        string    `json:"status"`
}
