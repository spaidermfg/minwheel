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
	orm.Debug = true
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
		ID:      6,
		Version: "v2.2.2",
	}
	models.UpdateHouse(house)
}

func TestInsert(t *testing.T) {
	data := make([]*models.HouseHood, 0)
	house1 := models.NewHouse("big1", "aa@google.com", "Tony", "v1.0.0", 10, true)
	house2 := models.NewHouse("big2", "aaa@google.com", "Tony", "v1.0.0", 11, false)
	house3 := models.NewHouse("big3", "aaaa@google.com", "Tony", "v1.0.0", 12, true)
	house4 := models.NewHouse("big4", "aaaaa@google.com", "Tony", "v1.0.0", 13, false)
	house5 := models.NewHouse("big5", "aaaaaa@google.com", "Tony", "v1.0.0", 14, true)
	house6 := models.NewHouse("big6", "aaaaaaa@google.com", "Tony", "v1.0.0", 151, false)
	house7 := models.NewHouse("big7", "aaaaaaab@google.com", "Star", "v2.0.0", 151, false)
	house8 := models.NewHouse("big8", "aaaaaaac@google.com", "Star", "v2.0.0", 152, false)
	house9 := models.NewHouse("big9", "aaaaaaad@google.com", "Star", "v2.0.0", 153, true)
	house10 := models.NewHouse("big10", "aaaaaaae@google.com", "Star", "v2.0.0", 154, false)
	house11 := models.NewHouse("big11", "aaaaaaaf@google.com", "Star", "v2.0.0", 155, false)
	house12 := models.NewHouse("big14", "aaaaaaaf@google.com", "Star", "v2.0.0", 155, false)
	data = append(data, house1, house2, house3, house4, house5, house6)
	data = append(data, house7, house8, house9, house10, house11)
	data = append(data, house12)

	models.AddHouse(data)
}

func TestRead(t *testing.T) {
	house := &models.HouseHood{
		ID: 6,
	}

	models.QueryHouse(house)
	log.Println(house)
}

func TestQuery(t *testing.T) {
	q := &models.HouseHood{
		//Name:    "",
		//Age: 10,
		//Sex: true,
		//Email: "b",
		Street: "z",
		//Version: "v",
	}

	houses := models.QueryHouses(1, 20, q)
	log.Println("len:", len(houses))
	if len(houses) > 0 {
		for k, v := range houses {
			log.Println(k, v)
		}
	}
}
