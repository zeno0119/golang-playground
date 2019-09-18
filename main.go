package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/timestamp"
)

var log

func main(){
	port := "8080"
	router := gin.Default()
	router.Static("/", "./client/dist")
	api := router.Group("/api"){
		api.GET("/post", post)
	}
	err := router.Run(":" + port)
	if err != nil {panic(err)}
	fmt.Printf("Server Listening at port %s \n", port)
}

func post (c *gin.Context){

}

type Chat struct{
	user string
	content string
	timestamp string
}