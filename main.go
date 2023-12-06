package main

import (
	"github.com/gin-gonic/gin"
	"github.com/paerarason/notes-api/database"
	"github.com/paerarason/notes-api/middlewares"
	"github.com/paerarason/notes-api/api/controller/users"
	"log"
		"database/sql"
	
)
var db *sql.DB
func init(){
		dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable"
		db, err :=sql.Open("postgres",dsn)
		if err!=nil{
			log.Fatal(err)
		} 
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(5)
		log.Println("DATABASE Connection made....sucessfully")
}
	


func main(){
	    router := gin.Default()
		router.Use(middlewares.JWTAuthentication())
		router.Use(func(c *gin.Context) {
            c.Set("db", db)
            c.Next()
           })
	    defer router.Run(":8080")
		//migrations
		router.Use(middleware.JWTAuthentication())

		user:=router.Group("/api/user")
		{
         user.POST("/",users.UserCreateHandler())
		}
}
	