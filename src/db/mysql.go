/*
Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
Project       dbtester
Description   A simple database tester
License       GPL version 2 (see GPL.txt for details)
*/

package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
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

func openMySQLDB(user, password, host, dbname string, port int) (*sql.DB, error) {
	dsn := user + ":" + password + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + dbname
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}