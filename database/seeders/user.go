package seeders

import (
	"log"

	"atom/http/database/models"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

type UserSeeder struct {
}

func NewUserSeeder() contracts.Seeder {
	return &UserSeeder{}
}

func (s *UserSeeder) Run(faker *gofakeit.Faker, db *gorm.DB) {
	times := 10
	for i := 0; i < times; i++ {
		data := s.Generate(faker, i)
		if i == 0 {
			stmt := &gorm.Statement{DB: db}
			_ = stmt.Parse(&data)
			log.Printf("seeding %s for %d times", stmt.Schema.Table, times)
		}
		db.Create(&data)
	}
}

func (s *UserSeeder) Generate(faker *gofakeit.Faker, idx int) models.User {
	return models.User{
		Username: faker.Name(),
		Age:      faker.Int32(),
		Sex:      faker.RandomString([]string{"male", "female"}),
		Birthday: faker.Date(),
		Status:   faker.RandomString([]string{"enabled", "disabled"}),
	}
}
