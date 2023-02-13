package config

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
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

	//create table
	query := `CREATE TABLE IF NOT EXISTS todos (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		completed BOOLEAN NOT NULL DEFAULT false
	);`
	
	_, err = db.Query("SELECT * FROM todos")
	if err != nil {
		log.Fatal(err)
	}

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
