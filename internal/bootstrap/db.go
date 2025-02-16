package bootstrap

import (
	"avito-shop-service/internal/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rubenv/sql-migrate"
	"log"
)

func InitDB(cfg *config.Config) (*sqlx.DB, error) {
	dest := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbSslMode,
	)

	db, err := sqlx.Open(cfg.DbDriver, dest)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.DbMaxIdleConn)
	db.SetMaxOpenConns(cfg.DbMaxOpenConn)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ApplyMigrations(cfg *config.Config, db *sqlx.DB) error {
	migrations := &migrate.FileMigrationSource{Dir: cfg.MigrationsDir}
	n, err := migrate.Exec(db.DB, cfg.DbDialect, migrations, migrate.Up)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Applied %d migrations!\n", n)

	return nil
}
