	package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

var log []Chat

func main(){
	port := os.Getenv("PORT")
	if port == "" {port = "8080"}
	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("chatSession", store))

	api := router.Group("/api")
	{
		api.POST("/login", login)
		api.POST("/post", post)
		api.GET("/get", get)
	}

	router.GET("/", func(c *gin.Context){
		c.File("./client/dist/index.html")
	})


	router.Static("/_nuxt", "./client/dist/_nuxt")


	err := router.Run(":" + port)
	if err != nil {panic(err)}
	fmt.Printf("Server Listening at port %s \n", port)
}

func post (c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("user")
	log = append(log, Chat{ user, c.Request.FormValue("content"), c.Request.FormValue("timestamp")})
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON( http.StatusOK, log)
}

func get (c *gin.Context){
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, log)
}

func login (c *gin.Context){
	session := sessions.Default(c)
	session.Set("user", c.Request.FormValue("user"))
	err := session.Save()
	if err != nil {panic(err)}
	print(session.Get("user").(string))
	c.Header("Access-Control-Allow-Origin", "*")
	c.Status(http.StatusOK)
}

type Chat struct{
	User interface{} `json:"user"`
	Content string `json:"content"`
	Timestamp string `json:"timestamp"`
}