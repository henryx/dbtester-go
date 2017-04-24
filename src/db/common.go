/*
Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
Project       dbtester
Description   A simple database tester
License       GPL version 2 (see GPL.txt for details)
*/

package db

import (
	"database/sql"
	"errors"
)

func OpenDB(engine, user, password, host, dbname string, port int) (*sql.DB, error) {
	var db *sql.DB
	var err error

	switch engine {
	case "mysql":
		db, err = openMySQLDB(user, password, host, dbname, port)
		return db, err
	default:
		return nil, errors.New("Unknown database engine: " + engine)
	}
}

func CheckStructure(db *sql.DB, dbname, engine string) (bool, error) {
	switch engine {
	case "mysql":
		return checkMySQLStructure(db, dbname), nil
	default:
		return false, errors.New("Unknown database engine: " + engine)
	}

}
