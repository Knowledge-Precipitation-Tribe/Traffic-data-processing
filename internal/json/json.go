package json

import "Traffic-data-processing/model"

/**
* @Author: super
* @Date: 2020-09-17 07:10
* @Description:
**/

func MarshJson(i interface{}) (string, error) {
	var result string
	var err error
	switch i.(type) {
	case model.ShortRoad:
		bytes, err := i.(model.ShortRoad).MarshalJSON()
		if err == nil{
			result = string(bytes)
		}
	case model.Road:
		bytes, err := i.(model.Road).MarshalJSON()
		if err == nil{
			result = string(bytes)
		}
	}
	return result, err
}

func UnMarshJsonRoad(bytes []byte) (*model.Road, error) {
	road := &model.Road{}
	err := road.UnmarshalJSON(bytes)
	if err != nil{
		return nil, err
	}
	return road, nil
}

func UnMarshJsonShortRoad(bytes []byte) (*model.ShortRoad, error) {
	road := &model.ShortRoad{}
	err := road.UnmarshalJSON(bytes)
	if err != nil{
		return nil, err
	}
	return road, nil
}