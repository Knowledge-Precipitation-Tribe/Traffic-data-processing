package mysql

import (
	"Traffic-data-processing/model"
	"errors"
)

/**
* @Author: super
* @Date: 2020-10-19 14:28
* @Description:
**/

func SelectMetaDataById(id int) (*model.MetaData, error) {
	db := GetDB()
	metadata := &model.MetaData{}
	result := db.Where("id = ?", id).Find(metadata)
	if result.RecordNotFound() {
		return nil, errors.New("wrong id")
	}
	return metadata, nil
}

func InsertMetaData(metaData *model.MetaData) error {
	db := GetDB()

	result := db.Create(metaData)
	if result.RowsAffected == int64(0) {
		return errors.New("insert error")
	}
	return nil
}

func UpdateMetaData(metaData *model.MetaData) error {
	db := GetDB()

	result := db.Model(metaData).Where("id = ?", metaData.ID).Updates(metaData)

	if result.RowsAffected == int64(0) {
		return errors.New("update error")
	}

	return nil
}

func DeleteMetaData(id int) error {
	db := GetDB()

	result := db.Where("id = ?", id).Delete(model.MetaData{})

	if result.RowsAffected == int64(0) {
		return errors.New("delete error")
	}

	return nil
}

func CountMetaData() (int, error) {
	db := GetDB()

	var count int

	result := db.Model(&model.MetaData{}).Count(&count)

	if result.RecordNotFound() {
		return -1, errors.New("count error")
	}

	return count, nil
}