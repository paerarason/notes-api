package models

import (
    "time"
)


type Job struct{
    id  int
    title string                 `json:"username"`
    experience_in_months  int    `json:"username"`
    description string           `json:"username"`
    created_at  time.Time
    expires_at  time.Time    
    posted_user int           
    comapany_id  int     `validate:required`
  }

