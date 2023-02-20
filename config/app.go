package config

import (
	"database/sql"
	"fmt"
	"log"
)

func Main() *sql.DB {
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
	//create db
	_, err = db.Exec(`CREATE Database todo`)
	if err != nil {
		log.Fatal(err)
	}
	//query db
	_, err = db.Exec(`USE todo`)
	if err != nil {
		log.Fatal(err)
	}
	// create table
	_, err = db.Exec(`
		CREATE TABLE todos (
		    id INT AUTO_INCREMENT,
		    title TEXT NOT NULL,
		    completed BOOLEAN DEFAULT FALSE,
		    PRIMARY KEY (id)
		);
	`)

	if err != nil {
		fmt.Println(err)
	}
	return db
}
