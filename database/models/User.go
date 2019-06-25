package models

const USERS = "members"

type User struct {
    Name     string `json:"name"`
    Age      int    `json:"age"`
    Company  string `json:"company"`
    IsActive bool   `json:"is_active"`
}

type UserPost struct {
    Name    string `json:"name"`
    Age     int    `json:"age"`
    Company string `json:"company"`
}

type UserPatch struct {
    Id  string `json:"id"`
    Age int    `json:"age"`
}

type UserApi struct {
    Id string `json:"id"`
    User
}
