package services

import (
	"Traffic-data-processing/internal/json"
	"Traffic-data-processing/pkg/logging"
	"Traffic-data-processing/pkg/mysql"
	"go.uber.org/zap"
)

/**
* @Author: super
* @Date: 2020-09-17 08:29
* @Description:
**/

//func StorageRoad(){
//	mq := rabbitmq.NewRabbitMQSimple(ROAD_MQ)
//	messages := mq.GetMsgs()
//	var wg sync.WaitGroup
//	for d := range messages {
//		go SaveRoad(d.Body)
//	}
//	wg.Wait()
//	mq.Destroy()
//}

func SaveRoad(bytes []byte){
	road, err := json.UnMarshJsonRoad(bytes)
	if err != nil{
		logging.GetLogger().Error("UnMarshJsonRoad", zap.Error(err))
		return
	}
	err = mysql.InsertRoad(road)
	if err != nil{
		logging.GetLogger().Error("InsertRoad", zap.Error(err))
	}
	logger.Info("save road", zap.String("detail", road.String()))
}