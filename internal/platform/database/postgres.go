package database

import (
	"database/sql"
	"fmt"

	"github.com/arthurhenrique-Dev/multi-tenant-api/internal/platform/bootstrap/boot"
	"github.com/arthurhenrique-Dev/multi-tenant-api/internal/platform/bootstrap/dependency"
	"github.com/arthurhenrique-Dev/multi-tenant-api/internal/platform/config"
	"github.com/arthurhenrique-Dev/multi-tenant-api/pkg"
)

func connect(config *config.Config) (*sql.DB, error) {
	dbConfig := config.Database
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Name,
	)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func startDatabase(cfg *config.Config) (any, error) {
	return connect(cfg)
}
func init() {
	boot.Dependencies.Register(
		*dependency.NewInitConfig(
			"database",
			startDatabase,
			pkg.NewSet[string]()),
	)
}
