package main

import (
	"github.com/HoYaStudy/Go_Study/hcoin/cli"
	"github.com/HoYaStudy/Go_Study/hcoin/db"
)

func main() {
	defer db.Close()
	db.InitDB()
	cli.Start()
}
