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


func JsonToMq(i interface{}){
	s, err := json.MarshJson(i)
	if err != nil{
		logging.GetLogger().Error("marshJson error", zap.Error(err))
		return
	}
	shortRoad := rabbitmq.NewRabbitMQSimple(SHORT_MQ)
	shortRoad.PublishSimple(string(s))
}