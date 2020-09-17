package services

import (
	"Traffic-data-processing/internal/json"
	"Traffic-data-processing/internal/jsonToMq"
	"Traffic-data-processing/model"
	"Traffic-data-processing/pkg/logging"
	"Traffic-data-processing/pkg/rabbitmq"
	"go.uber.org/zap"
	"strings"
	"sync"
)

/**
* @Author: super
* @Date: 2020-09-16 21:04
* @Description:
**/

var logger = logging.GetLogger()

func GetRoad(){
	mq := rabbitmq.NewRabbitMQSimple(jsonToMq.ROAD_MQ)
	messages := mq.GetMsgs()
	var wg sync.WaitGroup
	for d := range messages {
		//将道路信息转换为道路详细信息
		go RoadToShortRoad(d.Body)
		//将road信息存储到数据库
		go SaveRoad(d.Body)
	}
	wg.Wait()
	mq.Destroy()
}

func RoadToShortRoad(bytes []byte){
	road, err := json.UnMarshJsonRoad(bytes)
	logger.Error("UnMarshJsonRoad", zap.Error(err))
	ParseGeom(road)
}

func ParseGeom(road *model.Road) {
	if strings.HasPrefix(road.WktRoad, "MULTILINESTRING") {
		l := len(road.WktRoad)
		road.WktRoad = road.WktRoad[17 : l-2]
		roadToShort := &model.RoadToShort{}
		roadToShort.Road = road
		roadToShort.Roads = strings.Split(road.WktRoad, "),(")
		GetRoads(roadToShort)
	}
}

func GetRoads(road *model.RoadToShort) {
	if road == nil || len(road.Roads) == 0 {
		return
	}
	for i := 0; i < len(road.Roads); i++ {
		s := strings.Split(road.Roads[i], ",")
		for j := len(s) - 1; j > 0; j-- {
			start := strings.Split(s[j], " ")
			stop := strings.Split(s[j-1], " ")
			road := model.ShortRoad{
				TtiID:          road.Road.TtiID,
				TtiName:        road.Road.TtiName,
				StartLongitude: start[0],
				StartLatitude:  start[1],
				StopLongitude:  stop[0],
				StopLatitude:   stop[1],
			}
			go jsonToMq.ShortRoad(road)
		}
	}
}
