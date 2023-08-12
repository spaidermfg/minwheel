package test

import (
	"beegotest/models"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"testing"
)

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "data.db")
	orm.RegisterModel(new(models.HouseHood))
	orm.RunSyncdb("default", false, true)
}

func TestDel(t *testing.T) {
	house := &models.HouseHood{
		ID: 4,
	}
	models.DeleleHouse(house)
}

func TestUpdate(t *testing.T) {
	house := &models.HouseHood{
		ID:  4,
		Age: 900,
		Sex: true,
	}
	models.UpdateHouse(house)
}

func TestInsert(t *testing.T) {
	data := make([]*models.HouseHood, 0)
	house1 := models.NewHouse("big1", "aa@google.com", 10, true)
	house2 := models.NewHouse("big2", "aaa@google.com", 11, false)
	house3 := models.NewHouse("big3", "aaaa@google.com", 12, true)
	house4 := models.NewHouse("big4", "aaaaa@google.com", 13, false)
	house5 := models.NewHouse("big5", "aaaaaa@google.com", 14, true)
	house6 := models.NewHouse("big6", "aaaaaaa@google.com", 15, false)
	data = append(data, house1, house2, house3, house4, house5, house6)

	models.AddHouse(data)
}

func TestRead(t *testing.T) {
	house := &models.HouseHood{
		ID: 4,
	}

	models.QueryHouse(house)
	log.Println(house)
}
