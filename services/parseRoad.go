package services

import (
	"Traffic-data-processing/internal/jsonToMq"
	"Traffic-data-processing/model"
	"Traffic-data-processing/pkg/logging"
	"bufio"
	"encoding/csv"
	"go.uber.org/zap"
	"io"
	"os"
	"strconv"
	"time"
)

/**
* @Author: super
* @Date: 2020-09-17 07:35
* @Description:
**/

func ReadCSV(path string){
	file, err := os.Open(path)
	if err != nil{
		logging.GetLogger().Error("open file error",zap.Error(err))
		return
	}
	reader := csv.NewReader(bufio.NewReader(file))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			logging.GetLogger().Error("read file error", zap.Error(err))
		}
		id, err := strconv.Atoi(line[1])
		if err != nil{
			logging.GetLogger().Error("parse tti_id error", zap.Error(err))
		}
		road := model.Road{
			TtiID:id,
			TtiName:line[2],
			WktRoad:line[3],
		}
		go jsonToMq.MQRoad(road)
		time.Sleep(1 * time.Second)
	}
}