package model
import (
  "time"
)

type Application struct  {
     id int
     user_id  int
     Job_id int
	   Applied_at time.Time
}
