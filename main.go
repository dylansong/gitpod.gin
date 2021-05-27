package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world!!")
	r := gin.Default()
	r.GET("/",func(c *gin.Context){
		c.String(http.StatusOK,"good")
	})
	r.Run()
}
