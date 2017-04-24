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

func openMySQLDB(user, password, host, dbname string, port int) (*sql.DB, error) {
	var conn *sql.DB
	var err error

	dsn := user + ":" + password + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + dbname
	conn, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return conn, nil
}

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

func createMySQLStructure(conn *sql.DB) error {
	return nil
}