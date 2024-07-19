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

func retriveToDos(db *gorm.DB) ([]Todo, error) {
	var Todos []Todo
	result := db.Find(&Todos)
	if result.Error != nil {
		log.Fatal("Database migration failed : ", result.Error)
		return nil, result.Error
	}
	return Todos, nil
}

func addToDo(db *gorm.DB, Todo Todo) error {
	result := db.Create(&Todo)
	if result.Error != nil {
		log.Fatal("Database migration failed : ", result.Error)
		return result.Error
	}
	return nil
}
