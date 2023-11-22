package database

import (
    "gorm.io/gorm"
)

type User struct {
    ID           uint       `gorm:primarykey`
    Name         string
    Email        *string
    Age          uint8
    Birthday     *time.Time
    MemberNumber sql.NullString
    ActivatedAt  sql.NullTime
    CreatedAt    time.Time
    UpdatedAt    time.Time
  }

type Lists struct {
    task string 
    priority int32
    reviewer User
    assignee  []User   
}


type Board  struct {
    lists Lists
}


