package migrations

import (
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func init() {
	Migrations = append(Migrations, New20230614_092523AlterUserAddColumnStatusMigration)
}

type Migration20230614_092523AlterUserAddColumnStatus struct {
	id string
}

func New20230614_092523AlterUserAddColumnStatusMigration() contracts.Migration {
	return &Migration20230614_092523AlterUserAddColumnStatus{id: "20230614_092523_alter_user_add_column_status"}
}

func (m *Migration20230614_092523AlterUserAddColumnStatus) ID() string {
	return m.id
}

func (m *Migration20230614_092523AlterUserAddColumnStatus) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230614_092523AlterUserAddColumnStatus) Down(tx *gorm.DB) error {
	return tx.Migrator().DropColumn(m.table(), "state")
}

func (m *Migration20230614_092523AlterUserAddColumnStatus) table() interface{} {
	type User struct {
		State bool
	}

	return &User{}
}
