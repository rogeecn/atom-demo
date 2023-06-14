package migrations

import (
	"time"

	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func init() {
	Migrations = append(Migrations, New20230613_193349CreateUserMigration)
}

type Migration20230613_193349CreateUser struct {
	id string
}

func New20230613_193349CreateUserMigration() contracts.Migration {
	return &Migration20230613_193349CreateUser{id: "20230613_193349_create_user"}
}

func (m *Migration20230613_193349CreateUser) ID() string {
	return m.id
}

func (m *Migration20230613_193349CreateUser) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230613_193349CreateUser) Down(tx *gorm.DB) error {
	return tx.Migrator().DropTable(m.table())
}

func (m *Migration20230613_193349CreateUser) table() interface{} {
	type User struct {
		gorm.Model
		Username string
		Age      uint
		Sex      string
		Birthday time.Time
		Status   string
	}

	return &User{}
}
