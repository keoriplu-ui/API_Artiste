package models

type Concert struct {
	Id int `json:"id"`
	Date string `json:"date"`
	Artist_id int `json:"artiste_id"`
}

var ListConcerts []Concert