package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Machine struct {
	IP       string
	Port     string
	Arch     string
	Username string
	Password string
}

func main() {
	db, _ := sql.Open("sqlite3", "sqlite/test.db")
	sql := "select * from thirty_vlan where ;"
	r, _ := db.Query(sql)

	fmt.Println()
	for r.Next() {
		m := new(Machine)
		r.Scan(&m.IP, &m.Port, &m.Username, &m.Password, &m.Arch)
		fmt.Println(r)
		fmt.Println(m)
	}

}

/*


 */
