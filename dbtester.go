/*
Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
Project       dbtester
Description   A simple database tester
License       GPL version 2 (see GPL.txt for details)
*/

package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
	"flag"
	"db"
)

func readCfg(filename string) *ini.File {
	res, err := ini.Load([]byte{}, filename)
	if err != nil {
		fmt.Println("Error about reading config file:", err)
		os.Exit(1)
	}

	return res
}

func main() {
	var err error

	cfgfile := flag.String("cfg", "", "Set the configuration file")
	flag.Parse()

	cfg := readCfg(*cfgfile)
	sect := cfg.Section("mysql")
	port, err := sect.Key("port").Int()
	if err != nil {
		fmt.Println("Malformed port value in configuration file:", err)
		os.Exit(1)
	}

	user := sect.Key("user").String()
	password := sect.Key("password").String()
	host := sect.Key("host").String()
	dbname := sect.Key("dbname").String()

	dbconn, err := db.OpenDB("mysql", user, password, dbname, host, port)
	if err != nil {
		fmt.Println("Error opening database connection:", err)
		os.Exit(1)
	}
	defer dbconn.Close()

	if exist, err := db.CheckStructure("mysql", dbname, dbconn); err != nil || !exist {

	}

	fmt.Println("Hello World!")
}
