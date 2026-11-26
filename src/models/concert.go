package models

type Concert struct {
    Id       int    `json:"ID"`
    Date     string `json:"date"`
    ArtistId int    `json:"artist_id"`
}