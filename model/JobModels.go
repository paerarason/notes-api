package model

import (
    "time"
   // "encoding/json"
)


type Job struct{
    id  int
    title string                 `json:"username"`
    experience_in_months  int    `json:"Experience"`
    description string           `json:"Description"`
    created_at  time.Time
    expires_at  time.Time    
    posted_user int           
    comapany_id  int     //`validate:required`
  }

