package pems

import (
	"Traffic-data-processing/model"
	"Traffic-data-processing/pkg/logging"
	"Traffic-data-processing/pkg/mysql"
	"go.uber.org/zap"
)

/**
* @Author: super
* @Date: 2020-10-19 14:46
* @Description:
**/
var logger = logging.GetLogger()

func SaveMetaData(bytes []byte){
	metadata := &model.MetaData{}
	err := metadata.UnmarshalJSON(bytes)
	if err != nil{
		logger.Error("UnMarshJsonMetaData", zap.Error(err))
		return
	}
	err = mysql.InsertMetaData(metadata)
	if err != nil{
		logger.Error("InsertMetadata", zap.Error(err))
	}
	logger.Info("save metadata", zap.String("detail", metadata.String()))
}