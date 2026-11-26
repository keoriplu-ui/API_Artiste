package models

type Concert struct {
    Id       int    `json:"id"`
    Date     string `json:"date"`
    ArtistId int    `json:"artist_id"`
    Artist   Artist `json:"-" gorm:"foreignKey:ArtistId;references:Id"`
}