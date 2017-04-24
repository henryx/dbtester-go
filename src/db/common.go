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
	"strings"
	"fmt"
	"strconv"
)

func tables() []string {
	var res []string
	for i := 1; i <= 10; i++ {
		res = append(res,
			fmt.Sprintf(
				strings.Join([]string{
					"CREATE TABLE test%s(",
					"ins INTEGER,",
					"date_ins INTEGER,",
					"hour_ins INTEGER,",
					"txt VARCHAR(30))",
				}, " "), strconv.Itoa(i)))
	}
	return res
}

func views() []string {
	var res []string

	res = append(res, strings.Join([]string{
		"CREATE OR REPLACE VIEW test AS",
		"SELECT * FROM test1 UNION",
		"SELECT * FROM test2 UNION",
		"SELECT * FROM test3 UNION",
		"SELECT * FROM test4 UNION",
		"SELECT * FROM test5 UNION",
		"SELECT * FROM test6 UNION",
		"SELECT * FROM test7 UNION",
		"SELECT * FROM test8 UNION",
		"SELECT * FROM test9 UNION",
		"SELECT * FROM test10"},
		" "))

	return res
}

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
