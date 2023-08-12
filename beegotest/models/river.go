package models

type RiverMember struct {
	ID    int32  `orm:"column(id);auto"`
	Fish  string `orm:"column(fish);unique"`
	Water bool   `orm:"column(water)"`
}
