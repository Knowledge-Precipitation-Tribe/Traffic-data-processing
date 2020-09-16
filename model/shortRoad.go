package model

import "fmt"

/**
* @Author: super
* @Date: 2020-09-15 20:09
* @Description:
**/

type ShortRoad struct {
	TTiId          int    `json:"tti_id"`
	TTIName        string `json:"tti_name"`
	StartLongitude string `json:"start_longitude"`
	StartLatitude  string `json:"start_latitude"`
	StopLongitude  string `json:"stop_longitude"`
	StopLatitude   string `json:"stop_latitude"`
}

func (road ShortRoad) String() string{
	return fmt.Sprintf("TTIID: %d\n" +
		"TTINAME: %s\n" +
		"StartLongitude: %s\n" +
		"StartLatitude: %s\n" +
		"StopLongitude: %s\n" +
		"StopLatitude: %s\n",
		road.TTiId, road.TTIName,
		road.StartLongitude, road.StartLatitude,
		road.StopLongitude, road.StopLatitude)
}