package repository

import (
	"database/sql"
	"log"

	"github.com/brotigen23/gopherMart/internal/database"
	"github.com/brotigen23/gopherMart/internal/entity"
	_ "github.com/lib/pq"
)

type postgresRepository struct {
	db *sql.DB
}

const migrationPath = "migrations"

func NewPostgresUserRepository(driver string, stringConnection string) (Repository, error) {
	ret := &postgresRepository{}
	db, err := sql.Open(driver, stringConnection)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	migrator := database.MustGetNewMigrator(migrationPath)
	err = migrator.ApplyMigrations(db)
	if err != nil {
		return nil, err
	}
	db, err = sql.Open(driver, stringConnection)
	if err != nil {
		return nil, err
	}
	ret.db = db
	return ret, nil
}

func (r *postgresRepository) SaveWithdrawAndUpdateBalance(user *entity.User, sum float32, withdraw *entity.Withdraw) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	// Update balance
	queryToUpdateBalance := `	
	UPDATE Users 
	SET balance = balance + $1 WHERE id = $2`

	_, err = tx.Exec(queryToUpdateBalance, sum, user.ID)

	if err != nil {
		tx.Rollback()
		return err
	}

	queryToSaveWithdraw := `	
	INSERT INTO Withdrawals(user_id, "order", sum, processed_at) 
	VALUES($1, $2, $3, $4)`
	_, err = tx.Exec(queryToSaveWithdraw, user.ID, withdraw.Order, withdraw.Sum, withdraw.ProccessedAt)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}
