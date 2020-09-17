package model

import "fmt"

/**
* @Author: super
* @Date: 2020-09-16 20:18
* @Description:
**/

type Road struct {
	ID      int    `gorm:"column:id;primary_key" json:"id"`
	TtiID   int    `gorm:"column:tti_id" json:"tti_id"`
	TtiName string `gorm:"column:tti_name" json:"tti_name"`
	WktRoad string `gorm:"column:wkt_road" json:"wkt_road"`
	Roads   []string
}

// TableName sets the insert table name for this struct type
func (r *Road) TableName() string {
	return "roads"
}

func (road Road) String() string {
	return fmt.Sprintf("TTIID: %d\n"+
		"TTINAME: %s\n"+
		"WktRoad: %s\n",
		road.TtiID, road.TtiName,
		road.WktRoad)
}
