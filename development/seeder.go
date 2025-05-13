package development

import (
	"awesomeProject/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type Seeder struct {
	db *gorm.DB
}

func NewSeeder(db *gorm.DB) *Seeder {
	return &Seeder{db: db}
}

func (s *Seeder) SeedAdminUser() error {
	// Cek apakah user admin sudah ada
	var count int64
	s.db.Model(&models.User{}).Where("email = ?", "admin@example.com").Count(&count)
	if count > 0 {
		return nil
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Buat user admin
	adminUser := &models.User{
		Name:     "Admin",
		Nik:      "220401010029",
		Email:    "arisu@enigma.com",
		Password: string(hashedPassword),
	}

	return s.db.Create(adminUser).Error
}

func (s *Seeder) SeedCandidate() error {
	candidates := []models.Candidate{
		{
			ID:            uuid.New().String(),
			FullName:      "John Doe",
			Nickname:      "Johnny",
			DateOfBirth:   time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Address:       "123 Main St",
			PhoneNumber:   "1234567890",
			Email:         "john.doe@example.com",
			Batch:         "2023A",
			Skills:        "Go, Python",
			Experience:    5,
			SoftSkillTest: 85.5,
			MathTest:      90.0,
			CodingTest:    92.0,
			TotalSkills:   2,
			Status:        "Active",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			ID:            "2",
			FullName:      "Jane Smith",
			Nickname:      "Janey",
			DateOfBirth:   time.Date(1992, 2, 2, 0, 0, 0, 0, time.UTC),
			Address:       "456 Elm St",
			PhoneNumber:   "9876543210",
			Email:         "jane.smith@example.com",
			Batch:         "2023B",
			Skills:        "Java, JavaScript",
			Experience:    3,
			SoftSkillTest: 80.0,
			MathTest:      85.0,
			CodingTest:    88.0,
			TotalSkills:   2,
			Status:        "Active",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			ID:            "3",
			FullName:      "Alice Johnson",
			Nickname:      "Ali",
			DateOfBirth:   time.Date(1995, 3, 3, 0, 0, 0, 0, time.UTC),
			Address:       "789 Oak St",
			PhoneNumber:   "5551234567",
			Email:         "alice.johnson@example.com",
			Batch:         "2023C",
			Skills:        "C++, Ruby",
			Experience:    4,
			SoftSkillTest: 78.0,
			MathTest:      82.0,
			CodingTest:    85.0,
			TotalSkills:   2,
			Status:        "Active",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			ID:            "4",
			FullName:      "Bob Brown",
			Nickname:      "Bobby",
			DateOfBirth:   time.Date(1988, 4, 4, 0, 0, 0, 0, time.UTC),
			Address:       "321 Pine St",
			PhoneNumber:   "4449876543",
			Email:         "bob.brown@example.com",
			Batch:         "2023D",
			Skills:        "PHP, SQL",
			Experience:    6,
			SoftSkillTest: 88.0,
			MathTest:      92.0,
			CodingTest:    95.0,
			TotalSkills:   2,
			Status:        "Active",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			ID:            "5",
			FullName:      "Charlie Davis",
			Nickname:      "Charlie",
			DateOfBirth:   time.Date(1993, 5, 5, 0, 0, 0, 0, time.UTC),
			Address:       "654 Cedar St",
			PhoneNumber:   "3336547890",
			Email:         "charlie.davis@example.com",
			Batch:         "2023E",
			Skills:        "Swift, Kotlin",
			Experience:    2,
			SoftSkillTest: 75.0,
			MathTest:      80.0,
			CodingTest:    84.0,
			TotalSkills:   2,
			Status:        "Active",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
	}

	for _, candidate := range candidates {
		if err := s.db.Create(&candidate).Error; err != nil {
			return err
		}
	}

	return nil
}
