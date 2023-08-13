package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"log"
)

type HouseHood struct {
	ID      int32  `orm:"column(id);auto"`
	Name    string `orm:"column(name);unique"`
	Age     int64  `orm:"column(age)"`
	Sex     bool   `orm:"column(sex)"`
	Email   string `orm:"column(email)"`
	Street  string `orm:"column(street)"`
	Version string `orm:"column(version)"`
}

type SqliteSequence struct {
	Name string `orm:"column(name)"`
	Seq  int32  `orm:"column(seq)"`
}

func NewHouse(name, email, street, version string, age int64, sex bool) *HouseHood {
	return &HouseHood{
		Name:    name,
		Age:     age,
		Sex:     sex,
		Email:   email,
		Street:  street,
		Version: version,
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

	if house.Street != "" {
		data.Street = house.Street
	}

	if house.Version != "" {
		data.Version = house.Version
	}

	if house.Age != 0 {
		data.Age = house.Age
	}

	data.Sex = house.Sex
	res, err := orm.NewOrm().Update(data)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("[------------update------------]:", res, err)
}

func AddHouse(house []*HouseHood) {
	db := orm.NewOrm()
	//i, err := db.QueryTable("house_hood").Delete()
	//if err != nil {
	//	log.Fatal("--------------------", i, err)
	//}

	begin, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	begin.Raw("delete from house_hood").Exec()

	exec, err := begin.Raw(fmt.Sprintf("UPDATE `sqlite_sequence` SET `seq` = 0 WHERE `name` = '%v'", "house_hood")).Exec()
	if err != nil {
		log.Fatal(exec, err)
	}

	res, err := begin.InsertMulti(len(house), house)
	if err != nil {
		log.Fatal(err)
	}

	if err = begin.Commit(); err != nil {
		log.Fatal(err)
	}

	logs.Info("[-------------insert-----------]:", res, err)
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

func QueryHouses(page, limit int32, house *HouseHood) []*HouseHood {
	db := orm.NewOrm()
	q := db.QueryTable(&HouseHood{})
	//condition := orm.NewCondition()

	if house.Street != "" {
		//condition = condition.And("street__icontains", house.Street)
		q = q.Filter("street__icontains", house.Street)
	}

	if house.Version != "" {
		//condition = condition.And("version__icontains", house.Version)
		q = q.Filter("version__icontains", house.Version)
	}

	if house.Name != "" {
		q = q.Filter("name_icontains", house.Name)
	}

	if house.Sex {
		q = q.Filter("sex", house.Sex)
	}

	if house.Email != "" {
		q = q.Filter("email__icontains", house.Email)
	}

	if house.Age != 0 {
		q = q.Filter("age", house.Age)
	}

	var result []*HouseHood
	all, err := q.Limit(limit).Offset((page - 1) * limit).OrderBy("-name").All(&result)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("[-------------read-----------]:", all, err)
	return result
}
