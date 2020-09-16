package services

import (
	"Traffic-data-processing/internal/jsonToMq"
	"Traffic-data-processing/model"
	"strings"
)

/**
* @Author: super
* @Date: 2020-09-16 21:04
* @Description:
**/

func ParseGeom(road *model.Road) []string {
	if strings.HasPrefix(road.WktRoad, "MULTILINESTRING") {
		l := len(road.WktRoad)
		road.WktRoad = road.WktRoad[17 : l-2]
		road.Roads = strings.Split(road.WktRoad, "),(")
	}
	return nil
}

func GetRoads(road *model.Road) {
	if road == nil || len(road.Roads) == 0 {
		return
	}
	for i := 0; i < len(road.Roads); i++ {
		s := strings.Split(road.Roads[i], ",")
		for j := len(s) - 1; j > 0; j-- {
			start := strings.Split(s[j], " ")
			stop := strings.Split(s[j-1], " ")
			road := model.ShortRoad{
				TtiID:          road.TtiID,
				TtiName:        road.TtiName,
				StartLongitude: start[0],
				StartLatitude:  start[1],
				StopLongitude:  stop[0],
				StopLatitude:   stop[1],
			}
			go jsonToMq.JsonToMq(road)
		}
	}
}
