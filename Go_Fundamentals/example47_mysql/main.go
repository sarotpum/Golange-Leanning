package main

import (
	"encoding/json"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
 
func main() {
	var db *gorm.DB = connectDatabase()
	todoList := todov1{Username: "admin", Title: "Angular homework", Message: "Study about ngIf"}

	// Insert
	fmt.Println("----------- Create -------------")
	db.Create(&todoList)

	// Query
	fmt.Println("----------- Query -------------")
	query(db)

	// Update
	fmt.Println("----------- Update -------------")
	var tmp todov1
	db.First(&tmp, 1)
	update(db, tmp)
	query(db)

	// Delete
	fmt.Println("----------- Delete -------------")
	var deleteTmp []todov1
	db.Find(&deleteTmp, "message like ?", "%For")
	delete(db, deleteTmp)
	query(db)
}

func connectDatabase() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	database, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	database.AutoMigrate(&todov1{})
	return database
}

func query(_db *gorm.DB) {
	var todos []todov1
	_db.Find(&todos)
	// fmt.Println(todos)
	printPretty(todos)
}

func update(_db *gorm.DB, todo todov1) {
	_db.Model(&todo).Update("Message", "ngFor")
}

func delete(_db *gorm.DB, todos []todov1) {
	// _db.Unscoped().Delete(&todos) Soft

	_db.Unscoped().Delete(&todos) // Hard
}


func printPretty(data []todov1) {
	json, _ := json.MarshalIndent(data, "", " ")
	fmt.Println(string(json))
}

type todov1 struct {
	gorm.Model
	Username string
	Title    string
	Message  string
}

type todov2 struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Title    string `json:"title"`
	Message  string `json:"message"`
}
