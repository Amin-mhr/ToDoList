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
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //more compact than indentedJson but debug is easier in indentedJson
		}
		err := addToDo(db, todo)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.IndentedJSON(http.StatusOK, todo)

	}

}

func getAllToDos(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		Todos, err := retriveToDos(db)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.IndentedJSON(http.StatusOK, Todos)
	}

}
