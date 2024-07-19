package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Todo struct {
	gorm.Model
	Title       string `json:"title"`
	Done        bool   `json:"done"`
	Explanation string `json:"explanation"`
}

var dsn string = "host=localhost user=user password=password dbname=todolist port=5433 sslmode=disable TimeZone=Asia/Tehran"

func pgConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	name := db.Migrator().CurrentDatabase()
	_ = name
	err = db.AutoMigrate(&Todo{})
	if err != nil {
		log.Fatal("Database migration failed : ", err)
	}

	return db
}

func retriveToDos(db *gorm.DB) []Todo {
	var Todos []Todo
	db.Find(&Todos)
	return Todos
}

func addToDo(db *gorm.DB, todo Todo) {
	db.Create(todo)
	return
}
