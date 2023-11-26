package api 
import (
	"github.com/gin-gonic/gin"
	"github.com/paerarason/notes-api/database"
)

//Grt the Jobs
func CreateJob_Handler() gin.HandlerFunc{
	   return func (c *gin.Context){
		db:=database.SQL_Connect()
		
	   }
	  
}

func ReadJob_Handler() gin.HandlerFunc{
    
   return func (c *gin.Context){

   }
}

func UpdateJob_Handler() gin.HandlerFunc{

	return func (c *gin.Context){

	}
}

func DeleteJob_Handler() gin.HandlerFunc{
    return func (c *gin.Context){

	}
}