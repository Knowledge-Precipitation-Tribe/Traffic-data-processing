package model

import "fmt"

/**
* @Author: super
* @Date: 2020-09-15 20:09
* @Description:
**/

type ShortRoad struct {
	StopLongitude  string `gorm:"column:stop_longitude" json:"stop_longitude"`
	TtiID          int    `gorm:"column:tti_id" json:"tti_id"`
	TtiName        string `gorm:"column:tti_name" json:"tti_name"`
	ID             int    `gorm:"column:id;primary_key" json:"id"`
	StartLatitude  string `gorm:"column:start_latitude" json:"start_latitude"`
	StartLongitude string `gorm:"column:start_longitude" json:"start_longitude"`
	StopLatitude   string `gorm:"column:stop_latitude" json:"stop_latitude"`
}

// TableName sets the insert table name for this struct type
func (s *ShortRoad) TableName() string {
	return "shortRoads"
}

func (road ShortRoad) String() string {
	return fmt.Sprintf("TTIID: %d\n"+
		"TTINAME: %s\n"+
		"StartLongitude: %s\n"+
		"StartLatitude: %s\n"+
		"StopLongitude: %s\n"+
		"StopLatitude: %s\n",
		road.TtiID, road.TtiName,
		road.StartLongitude, road.StartLatitude,
		road.StopLongitude, road.StopLatitude)
}
