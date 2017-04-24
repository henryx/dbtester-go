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

func checkMySQLStructure(db *sql.DB, dbname string) bool {
	return true
}

func CheckStructure(db *sql.DB, dbname, engine string) (bool, error) {
	switch engine {
	case "mysql":
		return checkMySQLStructure(db, dbname), nil
	default:
		return nil, errors.New("Unknown database engine: " + engine)
	}

}
