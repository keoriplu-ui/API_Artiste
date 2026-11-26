package models

type Artist struct {
	Id int `json:"ID"`
	Name string `json:"name"`
	Genre string `json:"genre"`
}