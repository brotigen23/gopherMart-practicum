package repository

import (
	"database/sql"

	"github.com/brotigen23/gopherMart/internal/database"
	"github.com/brotigen23/gopherMart/internal/entity"
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
		return nil, err
	}
	ret.db = db
	migrator := database.MustGetNewMigrator(migrationPath)
	migrator.ApplyMigrations(db)
	return ret, nil
}

func (r *postgresUserRepository) GetUserByID() (*entity.User, error)    { return nil, nil }
func (r *postgresUserRepository) GetUserByLogin() (*entity.User, error) { return nil, nil }

func (r *postgresUserRepository) Save(*entity.User) error { return nil }
