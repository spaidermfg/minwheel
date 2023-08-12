package main

import (
	"beegotest/models"
	_ "beegotest/routers"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/mattn/go-sqlite3"
	"log"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	orm.RegisterDriver("sqlite3", orm.DRSqlite)

	orm.RegisterDataBase("default", "sqlite3", "data.db")

	orm.RegisterModel(new(models.HouseHood))

	orm.RegisterModel(new(models.RiverMember))

	operationDB()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func operationDB() {
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Fatal(err)
	}
	//
	//newOrm := orm.NewOrm()
	//house := new(models.HouseHood)
	//house.ID = 1
	//house.Name = "big"
	//
	//insert, err := newOrm.Insert(house)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(insert)
	//
	//m := new(models.RiverMember)
	//m.Water = true
	//m.Fish = "fish1"
	//insert1, err := newOrm.Insert(m)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(insert1)
}
