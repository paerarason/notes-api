package database

import (
    "gorm.io/gorm"
    "time"
)

type User struct {

    gorm.Model
    ID           uint       
    Name         string     `gorm:"serializer:json"`
    Email        *string    `gorm:"serializer:json"`
    Age          uint8      `gorm:"serializer:json"`
    CreatedAt    time.Time  `gorm:"serializer:json"`
    UpdatedAt    time.Time  `gorm:"serializer:json"`

    Roles
  }

type Lists struct {

    gorm.Model
    task         string   `gorm:"serializer:json"`
    priority     int32    `gorm:"serializer:json"`
    reviewer     User     `gorm:"serializer:json"`
    assignee     []User   `gorm:"foreignKey:serializer:json"`
}


type Board  struct {
    lists Lists       `gorm:"serializer:json"`
}

//

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


  