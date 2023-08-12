package models

import (
	"github.com/beego/beego/v2/client/orm"
	"log"
)

type HouseHood struct {
	ID    int32  `orm:"column(id);auto"`
	Name  string `orm:"column(name);unique"`
	Age   int64  `orm:"column(age)"`
	Sex   bool   `orm:"column(sex)"`
	Email string `orm:"column(email)"`
}

func NewHouse(name, email string, age int64, sex bool) *HouseHood {
	return &HouseHood{
		Name:  name,
		Age:   age,
		Sex:   sex,
		Email: email,
	}
}

func DeleleHouse(house *HouseHood) {
	res, err := orm.NewOrm().Delete(house)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("[------------delete------------]:", res, err)
}

func UpdateHouse(house *HouseHood) {
	data := &HouseHood{ID: house.ID}
	err := orm.NewOrm().Read(data)
	if err != nil {
		log.Fatal(err)
	}

	if house.Email != "" {
		data.Email = house.Email
	}

	data.Sex = house.Sex
	data.Age = house.Age
	res, err := orm.NewOrm().Update(data)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("[------------update------------]:", res, err)
}

func AddHouse(house []*HouseHood) {
	res, err := orm.NewOrm().InsertMulti(len(house), house)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("[-------------insert-----------]:", res, err)
}

func QueryHouse(house *HouseHood) {
	err := orm.NewOrm().Read(house)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("[-------------read-----------]:", err)
	//data := req.GetOnlineApp()
	//db := orm.NewOrm().QueryTable("online_app")
	//if data.Name != "" {
	//	db.Filter("name", data.Name)
	//}
	//
	//if data.Type != "" {
	//	db.Filter("type", data.Type)
	//}
	//
	//if data.Version != "" {
	//	db.Filter("version", data.Version)
	//}
	//
	//if data.Update {
	//	db.Filter("update", data.Update)
	//}
	//
	//if data.Delete {
	//	db.Filter("delete", data.Delete)
	//}
	//
	//if data.Download {
	//	db.Filter("download", data.Download)
	//}
	//
	//var app []*online.OnlineApp
	//num, err := db.All(&app)
	//if err != nil {
	//	return nil, err
	//}
	//
	//res := &online.QueryRes{
	//	Total:     int32(num),
	//	OnlineApp: app,
	//}
	//return res, err
}

func QueryHouses(house []*HouseHood) {
	err := orm.NewOrm().Read(house, "age")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("[-------------read-----------]:", err)
}
