package db

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pouyam79i/simple_quera/authX/config"
	"github.com/pouyam79i/simple_quera/authX/syserror"
)

var db *sql.DB = nil
var table string = ""

func Connect(dbi config.DataBaseInfo) error {
	fmt.Println("Go MySQL Tutorial")

	connectionInfo := dbi.Username + ":" + dbi.Password + "@tcp(" + dbi.IP + ":" + dbi.Port + ")/" + dbi.Database
	fmt.Println("Connecting to DB:\n", connectionInfo)
	var err error
	db, err = sql.Open("mysql", connectionInfo)

	// if there is an error opening the connection, handle it
	if err != nil {
		return errors.New(syserror.DatabaseConnectionFailure)
	}

	table = dbi.Table

	return nil
}

func Close() {
	if db != nil {
		db.Close()
		table = ""
	}
}

func DB() (*sql.DB, error) {
	if db == nil {
		return nil, errors.New(syserror.DatabaseFailure)
	}
	return db, nil
}

func Table() string {
	return table
}
