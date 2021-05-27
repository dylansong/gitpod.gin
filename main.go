package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name" form:"name" bind:"required"`
	Age  int    `json:"age" form:"age" bind:"required,min=6,max=100"`
}

func main() {
	fmt.Println("hello world!!")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"name": "xiaosong"})
	})
	r.POST("/user", func(c *gin.Context) {
		var user User
		err := c.BindJSON(&user)

		if err !=  nil {
			log.Fatal(err.Error())
		}
		c.JSON(http.StatusOK, gin.H{"message": user.Name, "age": user.Age, "status":"sucess"})
	})
	r.Run()
}
