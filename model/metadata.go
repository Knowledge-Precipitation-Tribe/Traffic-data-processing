package model

import "fmt"

/**
* @Author: super
* @Date: 2020-10-19 14:22
* @Description:
**/

type MetaData struct {
	CountryID int    `gorm:"column:country_id" json:"country_id"`
	Type      string `gorm:"column:type" json:"type"`
	Tag       string `gorm:"column:tag" json:"tag"`
	ID        int    `gorm:"column:id;primary_key" json:"id"`
	StatePm   string `gorm:"column:state_pm" json:"state_pm"`
	AbsPm     string `gorm:"column:abs_pm" json:"abs_pm"`
	Name      string `gorm:"column:name" json:"name"`
	Direction string `gorm:"column:direction" json:"direction"`
	Latitude  string `gorm:"column:latitude" json:"latitude"`
	Longitude string `gorm:"column:longitude" json:"longitude"`
	Length    string `gorm:"column:length" json:"length"`
	Number    int    `gorm:"column:number" json:"number"`
	Lanes     int    `gorm:"column:lanes" json:"lanes"`
	City      string `gorm:"column:city" json:"city"`
	District  int    `gorm:"column:district" json:"district"`
}

// TableName sets the insert table name for this struct type
func (metadata *MetaData) TableName() string {
	return "metadata"
}

func (metadata MetaData) String() string {
	return fmt.Sprintf("id: %d\n"+
		"latitude: %s\n"+
		"longitude: %s\n"+
		"type: %s\n"+
		"tag: %s\n"+
		"name: %s",
		metadata.ID, metadata.Latitude, metadata.Longitude,
		metadata.Type, metadata.Tag, metadata.Name)
}
