package database

type User struct { 
    name string 
    password string
    role string 
    team string
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


