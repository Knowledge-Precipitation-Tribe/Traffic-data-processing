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

func SelectShortRoadByTTIId(id int) (*model.ShortRoad, error) {
	db := GetDB()
	ShortRoad := &model.ShortRoad{}
	result := db.Where("tti_id = ?", id).Find(ShortRoad)
	if result.RecordNotFound() {
		return nil, errors.New("wrong id")
	}
	return ShortRoad, nil
}

func InsertShortRoad(shortRoad *model.ShortRoad) error {
	db := GetDB()

	result := db.Create(shortRoad)
	if result.RowsAffected == int64(0) {
		return errors.New("insert error")
	}
	return nil
}

func UpdateShortRoad(shortRoad *model.ShortRoad) error {
	db := GetDB()

	result := db.Model(shortRoad).Where("tti_id = ?", shortRoad.TTiId).Updates(shortRoad)

	if result.RowsAffected == int64(0) {
		return errors.New("update error")
	}

	return nil
}

func DeleteShortRoad(id int) error {
	db := GetDB()

	result := db.Where("tti_id = ?", id).Delete(model.ShortRoad{})

	if result.RowsAffected == int64(0) {
		return errors.New("delete error")
	}

	return nil
}

func CountShortRoad() (int, error) {
	db := GetDB()

	var count int

	result := db.Model(&model.ShortRoad{}).Count(&count)

	if result.RecordNotFound() {
		return -1, errors.New("count error")
	}

	return count, nil
}