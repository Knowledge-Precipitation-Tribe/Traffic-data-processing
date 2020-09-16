package model

/**
* @Author: super
* @Date: 2020-09-16 20:18
* @Description:
**/

type Road struct {
	Id      int    `json:"id"`
	TTiId   int    `json:"tti_id"`
	TTIName string `json:"tti_name"`
	WKTRoad string `json:"wkt_road"`
}
