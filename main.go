package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/paerarason/notes-api/database"
	
)


type Message struct {
	Sender string `json:"sender"`
	Content string `json:"content"`
}

func main(){

	    router := gin.Default()
	    defer router.Run(":8080")
		//migrations
		db:=database.SQL_Connect()
        db.Automigrate(&Job,&Application,&users,&Role)
		
		//Routers 
		user:=router.Group("/user")
		{
			user.POST("/login")
			user.GET("/users")
			user.GET("/user/:id")
		}

		projects:=router.Group("/project")
		{
			projects.GET("/list")
	
        }
}
	