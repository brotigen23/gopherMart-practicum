package repository

import (
	"database/sql"
	"log"

	"github.com/brotigen23/gopherMart/internal/database"
	_ "github.com/lib/pq"
)

type postgresUserRepository struct {
	db *sql.DB
}

const migrationPath = "migrations"

func NewPostgresUserRepository(driver string, stringConnection string) (UserRepository, error) {
	ret := &postgresUserRepository{}
	db, err := sql.Open(driver, stringConnection)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	migrator := database.MustGetNewMigrator(migrationPath)
	migrator.ApplyMigrations(db)
	db, err = sql.Open(driver, stringConnection)
	if err != nil {
		return nil, err
	}
	ret.db = db
	return ret, nil
}
