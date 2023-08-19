package migrate

import (
	"go-starter-kit/internal/database/migration"
	"go-starter-kit/pkg/config"
	"go-starter-kit/pkg/database"
)

func Up() {
	config.Set()

	database.Connect()

	migration.MigrateUp()
}
