package models

type Room struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"<-:create",unique_index`
}
