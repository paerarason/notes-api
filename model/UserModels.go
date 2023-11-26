package models

import (
  "time"
)




type User struct  {
    id int
    username string   `json:"username"`
    email string      `json:"email"`
    password string   `json:"password"`
    created_at time.Time 
  }