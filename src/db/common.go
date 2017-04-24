/*
Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
Project       dbtester
Description   A simple database tester
License       GPL version 2 (see GPL.txt for details)
*/

package db

import (
	"database/sql"
	"strings"
	"errors"
)

func checkMySQLStructure(db *sql.DB, dbname string) bool {
	var counted int
	query := strings.Join([]string{"SELECT count(*)",
				       "FROM information_schema.tables",
				       "WHERE table_schema = $1"}, " ")

	db.QueryRow(query, dbname).Scan(&counted)
	if counted > 0 {
		return true
	} else {
		return false
	}
}

func CheckStructure(db *sql.DB, dbname, engine string) (bool, error) {
	switch engine {
	case "mysql":
		return checkMySQLStructure(db, dbname), nil
	default:
		return nil, errors.New("Unknown database engine: " + engine)
	}

}
