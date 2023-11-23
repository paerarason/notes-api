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


