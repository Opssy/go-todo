package models

type Todo struct {
	Id        string `gorm: ""json:"id"`
	Title     string `gorm: ""json:"title"`
	Completed bool   `gorm: ""json:"completed"`
}
