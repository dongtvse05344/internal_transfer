package app

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	db "github.com/internal_transfer/dal/db/sqlc"
)

const (
	dbDriver = "mysql"
	dbSource = "root:my_password@tcp(127.0.0.1:3306)/mysql"
)

func WithStore(app *App) error {
	dbConnection, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		return err
	}
	store := db.NewStore(dbConnection)
	app.Store = store
	return nil
}
