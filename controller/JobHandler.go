package api 

import (
	"github.com/gin-gonic/gin"
	//"github.com/paerarason/notes-api/models"
    "github.com/go-playground/validator/v10"

)


func CreateJob_Handler() gin.HandlerFunc{
	   return func (c *gin.Context){
		db:=database.SQL_Connect()
		
		//validate the Given Values with struct
		if err := c.ShouldBindHeader(&body{}); err != nil {
            c.AbortWithStatus(400)
            return
		}

		validate := validator.New()
        err := validate.Struct(&body{})
        if err != nil {
            c.AbortWithStatus(400)
            return
        }
        user:=models.User{username:body.username,email:body.email,password:body.password,created_at:time.Now()}
		return db.Create(&user) 
	   }
	  
} 

//
func ReadJob_Handler() gin.HandlerFunc{
    

   return func (c *gin.Context){

   }
}


//
func UpdateJob_Handler() gin.HandlerFunc{

	
	return func (c *gin.Context){

	}
}

//
func DeleteJob_Handler() gin.HandlerFunc{
    return func (c *gin.Context){

	}
}