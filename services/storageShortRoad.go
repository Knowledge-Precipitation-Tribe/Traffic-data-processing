package services

import (
	"Traffic-data-processing/internal/json"
	"Traffic-data-processing/pkg/logging"
	"Traffic-data-processing/pkg/mysql"
	"Traffic-data-processing/pkg/rabbitmq"
	"go.uber.org/zap"
	"sync"
)

/**
* @Author: super
* @Date: 2020-09-17 07:36
* @Description:
**/

func StorageShortRoad(){
	mq := rabbitmq.NewRabbitMQSimple(SHORT_MQ)
	messages := mq.GetMsgs()
	var wg sync.WaitGroup
	for d := range messages {
		go SaveShortRoad(d.Body)
	}
	wg.Wait()
	mq.Destroy()
}

func SaveShortRoad(bytes []byte){
	shortRoad, err := json.UnMarshJsonShortRoad(bytes)
	if err != nil{
		logging.GetLogger().Error("UnMarshJsonShortRoad", zap.Error(err))
		return
	}
	err = mysql.InsertShortRoad(shortRoad)
	if err != nil{
		logging.GetLogger().Error("InsertShortRoad", zap.Error(err))
	}
}