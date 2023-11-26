package database
import (
    "gorm.io/gorm"
    "time"
)


type Job struct{
    id  int  [primary key ]
    title string 
    experience_in_months  int
    description string 
    created_at time
    expires_at time 
    posted_user int
    comapany_id  int
  }
  
type users struct  {
    id int [primary key]
    username string
    password string  
    created_at time
  }
  
type user_role struct  {
    id int [primary key ]
    user_id int
    role_id int
  }
  
type Role struct  {
     user_role_id int
     name string  [primary key]
  }
  
type Application struct  {
     id int [primary key ]
     user_id  int
     Job_id int
  }


