package models
import (
    "gorm.io/gorm"
    "time"
    "encoding/json"
)


type Job struct{
    id  int
    title string                 `json:"username"`
    experience_in_months  int   `json:"username"`
    description string         `json:"username"`
    created_at  time
    expires_at  time      
    posted_user int           
    comapany_id  int     `validate:required`
  }




type User struct  {
    id int
    username string   `json:"username"`
    email string      `json:"email"`
    password string   `json:"password"`
    created_at time   
  }
  
type user_role struct  {
    id int 
    user_id int
    role_id int
  }
  
type Role struct  {
     user_role_id int
     name string
  }
  
type Application struct  {
     id int
     user_id  int
     Job_id int
  }


