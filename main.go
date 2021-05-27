package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitpod.gin/account"
)

func main() {
	fmt.Println("hello world!!")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"name ": "xiaosong"})
	})
	r.POST("/user", func(c *gin.Context) {
		var account account.Account
		err := c.BindJSON(&account)

		if err != nil {
			log.Fatal(err.Error())
		}
		c.JSON(http.StatusOK, gin.H{"message": account.Name, "age": account.Age, "status": "sucess"})
	})
	r.Run()
}
