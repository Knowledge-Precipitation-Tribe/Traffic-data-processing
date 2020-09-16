package mysql

import (
	"Traffic-data-processing/model"
	"errors"
)

/**
* @Author: super
* @Date: 2020-09-16 20:38
* @Description:
**/

func SelectRoadById(id int) (*model.Road, error) {
	db := GetDB()
	Road := &model.Road{}
	result := db.Where("tti_id = ?", id).Find(Road)
	if result.RecordNotFound() {
		return nil, errors.New("wrong id")
	}
	return Road, nil
}

func InsertRoad(Road *model.Road) error {
	db := GetDB()

	result := db.Create(Road)
	if result.RowsAffected == int64(0) {
		return errors.New("insert error")
	}
	return nil
}

func UpdateRoad(Road *model.Road) error {
	db := GetDB()

	result := db.Model(Road).Where("tti_id = ?", Road.TtiID).Updates(Road)

	if result.RowsAffected == int64(0) {
		return errors.New("update error")
	}

	return nil
}

func DeleteRoad(id int) error {
	db := GetDB()

	result := db.Where("tti_id = ?", id).Delete(model.Road{})

	if result.RowsAffected == int64(0) {
		return errors.New("delete error")
	}

	return nil
}

func CountRoad() (int, error) {
	db := GetDB()

	var count int

	result := db.Model(&model.Road{}).Count(&count)

	if result.RecordNotFound() {
		return -1, errors.New("count error")
	}

	return count, nil
}