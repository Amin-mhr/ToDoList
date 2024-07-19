package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	db := pgConnection()
	_ = db
	router := gin.Default()

	router.GET("/todo", getAllToDos(db))
	router.POST("/todo", AddToDo(db))

	router.Run(":6060")
}

func AddToDo(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var todo Todo
		if err := c.Bind(&todo); err != nil {
			log.Print("binding todo failed : ", err)
		}
		addToDo()
	}

}

func getAllToDos(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		Todos := retriveToDos(db)
		c.IndentedJSON(http.StatusOK, Todos)
	}

}
