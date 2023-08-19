package migration

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go-starter-kit/pkg/config"
	"log"
)

func MigrateDown() {
	cfg := config.Get()

	dsn := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)

	m, err := migrate.New(
		"file://db/migrations",
		dsn,
	)

	if err != nil {
		log.Fatal("Cannot connect to database")
		return
	}

	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Cannot migrate down database")
		return
	}

	log.Println("Database migrated down successfully")
	return
}
