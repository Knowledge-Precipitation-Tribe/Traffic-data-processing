package jsonToMq

import (
	"Traffic-data-processing/internal/json"
	"Traffic-data-processing/pkg/logging"
	"Traffic-data-processing/pkg/rabbitmq"
	"go.uber.org/zap"
)

/**
* @Author: super
* @Date: 2020-09-17 07:17
* @Description:
**/

var logger = logging.GetLogger()

func MQRoad(i interface{}){
	s, err := json.MarshJson(i)
	if err != nil{
		logging.GetLogger().Error("marshJson error", zap.Error(err))
		return
	}
	logger.Info("road", zap.String("detail", s))
	road := rabbitmq.NewRabbitMQSimple(ROAD_MQ)
	road.PublishSimple(string(s))
	road.Destroy()
}

func ShortRoad(i interface{}){
	s, err := json.MarshJson(i)
	if err != nil{
		logging.GetLogger().Error("marshJson error", zap.Error(err))
		return
	}
	logger.Info("short road", zap.String("detail", s))
	shortRoad := rabbitmq.NewRabbitMQSimple(SHORT_MQ)
	shortRoad.PublishSimple(string(s))
	shortRoad.Destroy()
}