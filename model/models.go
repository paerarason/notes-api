package models
import (
    "gorm.io/gorm"
    "time"
)


type Job struct{
    id  int  [primary key ]
    title string                `validate:required`
    experience_in_months  int   `validate:required`
    description string          `validate:required`
    created_at time
    expires_at time      
    posted_user int           
    comapany_id  int     `validate:required`
  }
  

type User struct  {
    id int
    username string   `validate:required`
    email string      `validate:required`
    password string   `validate:required`
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


