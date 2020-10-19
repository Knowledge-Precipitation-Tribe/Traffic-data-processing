package pems

import (
	"Traffic-data-processing/model"
	"Traffic-data-processing/pkg/logging"
	"Traffic-data-processing/pkg/mysql"
	"bufio"
	"encoding/csv"
	"go.uber.org/zap"
	"io"
	"os"
	"strconv"
)

/**
* @Author: super
* @Date: 2020-10-19 14:25
* @Description:
**/

func ReadFileDirect(path string) {
	file, err := os.Open(path)
	if err != nil {
		logging.GetLogger().Error("open file error", zap.Error(err))
		return
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			logging.GetLogger().Error("read file error", zap.Error(err))
		}

		id, err := strconv.Atoi(line[1])
		if err != nil {
			logging.GetLogger().Error("parse id error", zap.Error(err))
		}
		number, err := strconv.Atoi(line[2])
		if err != nil {
			logging.GetLogger().Error("parse number error", zap.Error(err))
		}
		district, err := strconv.Atoi(line[4])
		if err != nil {
			logging.GetLogger().Error("parse district error", zap.Error(err))
		}
		countryId, err := strconv.Atoi(line[5])
		if err != nil {
			logging.GetLogger().Error("parse countryId error", zap.Error(err))
		}
		lane, err := strconv.Atoi(line[13])
		if err != nil {
			logging.GetLogger().Error("parse lane error", zap.Error(err))
		}

		metadata := &model.MetaData{
			ID:        id,
			Number:    number,
			Direction: line[3],
			District:  district,
			CountryID: countryId,
			City:      line[6],
			StatePm:   line[7],
			AbsPm:     line[8],
			Latitude:  line[9],
			Longitude: line[10],
			Length:    line[11],
			Type:      line[12],
			Lanes:     lane,
			Name:      line[14],
			Tag:       "",
		}
		err = mysql.InsertMetaData(metadata)
		if err != nil {
			logging.GetLogger().Error("Insert metadata", zap.Error(err))
		}
		logger.Info("save metadata", zap.String("detail", metadata.String()))
	}
}
