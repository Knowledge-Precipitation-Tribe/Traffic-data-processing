package main

import (
	"Traffic-data-processing/config"
	"Traffic-data-processing/pkg/logging"
	"Traffic-data-processing/pkg/mysql"
	"Traffic-data-processing/services/pems"
	"go.uber.org/zap"
)

/**
* @Author: super
* @Date: 2020-09-15 20:04
* @Description:
**/

func main() {
	//init config
	if err := config.Init("local"); err != nil {
		logging.GetLogger().Error("init config error", zap.Error(err))
		panic(err)
	}

	mysql.InitDB()
	//jsonToMq.InitMQ()

	//services.ReadCSVDirect("/Users/super/develop/Traffic-data-processing/data/boundry.csv")
	pems.ReadFileDirect("/Users/super/develop/Traffic-data-processing/data/pems/metadata.csv")
	//services.GetRoad()
	//services.StorageShortRoad()
}