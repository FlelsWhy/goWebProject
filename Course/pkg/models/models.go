package models

import "time"

type Users struct {
	Id             int       `json:"id"`
	Login          string    `json:"login"`
	DataOfRegister time.Time `json:"data of create accaunt"`
	Email          string    `json:"your email"`
}

type Spendings struct {
	Id         int
	Price      int
	Data       time.Time
	UserId     int
	CatigoryId int
}
