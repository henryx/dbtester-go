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
	cfgfile := flag.String("cfg", "", "Set the configuration file")

	cfg := readCfg(*cfgfile)
	sect := cfg.Section("mysql")
	port, err := sect.Key("port").Int()
	if err != nil {
		fmt.Println("Malformed port value in configuration file:", err)
		os.Exit(1)
	}

	fmt.Println("Hello World!")
}
