package main

import (
	"github.com/gin-gonic/gin"
	"github.com/paerarason/notes-api/model"
	"github.com/paerarason/notes-api/api"
	"github.com/paerarason/notes-api/database"
	"github.com/paerarason/notes-api/middleware"
	
)


func main(){
	    router := gin.Default()
		router.Use(middleware.JWTAuthentication())
	    defer router.Run(":8080")
		//migrations
		db:=database.SQL_Connect()
        db.AutoMigrate(&model.Job{},&model.Application{},&model.User{},&model.Role{})
		api.ApplicationRouter(router)
		api.CompanyRouter(router)
		api.JobRouter(router)
		api.ReviewRouter(router)
		api.UserRouter(router)
}
	