package services

import (
	"Traffic-data-processing/internal/json"
	"Traffic-data-processing/internal/jsonToMq"
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
	mq := rabbitmq.NewRabbitMQSimple(jsonToMq.SHORT_MQ)
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
		logger.Error("UnMarshJsonShortRoad", zap.Error(err))
		return
	}
	err = mysql.InsertShortRoad(shortRoad)
	if err != nil{
		logger.Error("InsertShortRoad", zap.Error(err))
	}
	logger.Info("save short road", zap.String("detail", shortRoad.String()))
}