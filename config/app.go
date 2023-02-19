package config

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

func Main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/tododb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

}
func GetDB() *sql.DB {
	return db
}

//var (
//	db *gorm.DB
//)
//
//func Connect() {
//	d, err := gorm.Open("mysql", "ope:opeyemi123/opetable?charset=uft8&parseTime=True&loc=local")
//	if err != nil {
//		panic(err)
//	}
//	db = d
//}
//func GetDB() *gorm.DB {
//	return db
//}
