package models

type Customer struct {
    Id int `json:"id" gorm:"primaryKey"`
    First_Name string `json:"first_name"`
	Last_Name string `json:"last_name"`
	Email string `json:"email"`
	Password string `json:"password"`
}