package development

import (
	"awesomeProject/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
