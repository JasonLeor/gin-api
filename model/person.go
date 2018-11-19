package model

import (
	"gindemo/db"
	"log"
)

type Person struct {
	Id   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Age  int    `json:"age" form:"age"`
}

func (p *Person) NewPerson() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO person(name, age) VALUES (?, ?)", p.Name, p.Age)
	if err != nil {
		log.Fatal(err.Error())
	}
	id, err = rs.LastInsertId()
	if err != nil {
		log.Fatal(err.Error())
	}
	return
}
func (Person) findOne(id int64) (person Person, err error) {
	rs, err := db.SqlDB.Query("select * from person where id =?1", id)
	defer rs.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	rs.Scan(&person)
	return
}
