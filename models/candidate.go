package models

import "time"

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
	SkExp         float64   `json:"sk_exp"`
	BasicTest     float64   `json:"basic_test"`
	CodingTest    float64   `json:"coding_test"`
	TotalSkills   int       `json:"total_skills"`
	TotalScore    float64   `json:"total_score"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
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
	TotalScore    float64   `json:"total_score"`
	Status        string    `json:"status"`
}
