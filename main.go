package main

import (
	_ "bulletin_board/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"fmt"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	dataSource := "dbname=bulletin_board user=bulletin_board password=1234567890 host=localhost port=5432 sslmode=disable"
	orm.RegisterDataBase("default", "postgres", dataSource)
}

func main() {
	// Database alias.
	name := "default"
	// Drop table and re-create.
	force := false
	// Print log.
	verbose := true
	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}

	beego.Run()
}
