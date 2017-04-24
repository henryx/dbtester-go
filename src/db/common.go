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

func OpenDB(engine, user, password, dbname, host string, port int) (*sql.DB, error) {
	var conn *sql.DB
	var err error

	switch engine {
	case "mysql":
		conn, err = openMySQLDB(user, password, host, dbname, port)
		return conn, err
	default:
		return nil, errors.New("Unknown database engine: " + engine)
	}
}

func CheckStructure(engine, dbname string, conn *sql.DB) (bool, error) {
	switch engine {
	case "mysql":
		return checkMySQLStructure(conn, dbname), nil
	default:
		return false, errors.New("Unknown database engine: " + engine)
	}
}

func CreateStructure(engine string, conn *sql.DB) error {
	switch engine {
	case "mysql":
		return createMySQLStructure(conn)
	default:
		return errors.New("Unknown database engine: " + engine)
	}
}