package main

import (
	"github.com/gin-gonic/gin"
	"github.com/paerarason/notes-api/database"
	"github.com/paerarason/notes-api/middlewares"
	
)


func main(){
	    router := gin.Default()
		router.Use(middlewares.JWTAuthentication())
		db:=database.SQL_Connect()
		router.Use(func(c *gin.Context) {
            c.Set("db", db)
            c.Next()
           })
	    defer router.Run(":8080")
		//migrations
		router.Use(middleware.JWTAuthentication())

		router.Group()
		{
         
		}
}
	