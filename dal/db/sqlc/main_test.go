package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var (
	testQueries *Queries
	testStore   *Store
)

const (
	dbDriver = "mysql"
	dbSource = "root:my_password@tcp(127.0.0.1:3306)/mysql"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	defer conn.Close()
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(conn)
	testStore = NewStore(conn)
	os.Exit(m.Run())
}
