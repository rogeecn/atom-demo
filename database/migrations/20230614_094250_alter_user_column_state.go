package migrations

import (
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230614_094250AlterUserColumnState) table() interface{} {
	type User struct {
		State string
	}

	return &User{}
}

func (m *Migration20230614_094250AlterUserColumnState) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230614_094250AlterUserColumnState) Down(tx *gorm.DB) error {
	return tx.Migrator().DropTable(m.table())
	// return tx.Migrator().DropColumn(m.table(), "input_column_name")
}

// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
func init() {
	Migrations = append(Migrations, New20230614_094250AlterUserColumnStateMigration)
}

type Migration20230614_094250AlterUserColumnState struct {
	id string
}

func New20230614_094250AlterUserColumnStateMigration() contracts.Migration {
	return &Migration20230614_094250AlterUserColumnState{id: "20230614_094250_alter_user_column_state"}
}

func (m *Migration20230614_094250AlterUserColumnState) ID() string {
	return m.id
}
